#if defined(CIMGUI_GO_USE_METAL)

#include "../../cwrappers/imgui/imgui.h"
#include "metal_backend.h"
#include "../../cwrappers/imgui/backends/imgui_impl_metal.h"
#include "../../cwrappers/imgui/backends/imgui_impl_osx.h"

#import <Cocoa/Cocoa.h>
#import <Metal/Metal.h>
#import <MetalKit/MetalKit.h>

#include <stdlib.h>
#include <string.h>

extern "C" {
extern void dropCallback(void* wnd, int count, char** names);
extern void closeCallback(void* wnd);
extern void sizeCallback(void* wnd, int w, int h);
}

@class MetalView;
static void igMetalResetTimer(struct MetalBackendContext* ctx);

struct MetalBackendContext {
    __strong NSWindow* window;
    __strong MetalView* view;
    __strong id<MTLDevice> device;
    __strong id<MTLCommandQueue> commandQueue;
    __strong id viewDelegate;
    __strong id appDelegate;
    __strong NSTimer* timer;
    __strong id timerTarget;
    ImVec4 clearColor;
    unsigned int targetFPS;
    bool hasCloseCallback;
    bool hasDropCallback;
    bool hasSizeCallback;

    VoidCallback loop;
    VoidCallback beforeRender;
    VoidCallback afterRender;
    VoidCallback beforeDestroyContext;
};

@interface MetalView : MTKView<NSDraggingDestination>
@property (nonatomic, assign) MetalBackendContext* ctx;
@end

@implementation MetalView

- (NSDragOperation)draggingEntered:(id<NSDraggingInfo>)sender
{
    if (self.ctx == nullptr || !self.ctx->hasDropCallback)
        return NSDragOperationNone;
    return NSDragOperationCopy;
}

- (BOOL)performDragOperation:(id<NSDraggingInfo>)sender
{
    if (self.ctx == nullptr || !self.ctx->hasDropCallback)
        return NO;

    NSPasteboard* pasteboard = [sender draggingPasteboard];
    NSDictionary* options = @{NSPasteboardURLReadingFileURLsOnlyKey: @YES};
    NSArray<NSURL*>* files = [pasteboard readObjectsForClasses:@[[NSURL class]] options:options];
    if (![files isKindOfClass:[NSArray class]])
        return NO;

    int count = (int)files.count;
    if (count <= 0)
        return NO;

    char** names = (char**)calloc((size_t)count, sizeof(char*));
    if (names == NULL)
        return NO;

    for (int i = 0; i < count; i++)
    {
        NSString* path = files[i].path;
        const char* utf8 = [path UTF8String];
        if (utf8 != NULL)
            names[i] = strdup(utf8);
        else
            names[i] = strdup("");
    }

    dropCallback((__bridge void*)self.ctx->window, count, names);

    for (int i = 0; i < count; i++)
    {
        if (names[i] != NULL)
            free(names[i]);
    }
    free(names);

    return YES;
}

@end

@interface MetalViewDelegate : NSObject<MTKViewDelegate, NSWindowDelegate>
@property (nonatomic, assign) MetalBackendContext* ctx;
@end

@implementation MetalViewDelegate

- (void)drawInMTKView:(MTKView*)view
{
    MetalBackendContext* ctx = self.ctx;
    if (ctx == nullptr)
        return;
    if (view.window == nil || !NSApp.isActive)
        return;
    if ((view.window.occlusionState & NSWindowOcclusionStateVisible) == 0)
        return;

    if (ctx->beforeRender != NULL)
        ctx->beforeRender();

    ImGuiIO& io = ImGui::GetIO();
    io.DisplaySize.x = view.bounds.size.width;
    io.DisplaySize.y = view.bounds.size.height;

    CGFloat framebufferScale = view.window.screen.backingScaleFactor ?: NSScreen.mainScreen.backingScaleFactor;
    io.DisplayFramebufferScale.x = (float)framebufferScale;
    io.DisplayFramebufferScale.y = (float)framebufferScale;

    id<MTLCommandBuffer> commandBuffer = [ctx->commandQueue commandBuffer];

    MTLRenderPassDescriptor* renderPassDescriptor = view.currentRenderPassDescriptor;
    if (renderPassDescriptor == nil)
    {
        [commandBuffer commit];
        return;
    }

    ImGui_ImplMetal_NewFrame(renderPassDescriptor);
    ImGui_ImplOSX_NewFrame(view);
    ImGui::NewFrame();

    if (ctx->loop != NULL)
        ctx->loop();

    ImGui::Render();
    ImDrawData* draw_data = ImGui::GetDrawData();

    renderPassDescriptor.colorAttachments[0].clearColor = MTLClearColorMake(
        ctx->clearColor.x * ctx->clearColor.w,
        ctx->clearColor.y * ctx->clearColor.w,
        ctx->clearColor.z * ctx->clearColor.w,
        ctx->clearColor.w);

    id<MTLRenderCommandEncoder> renderEncoder = [commandBuffer renderCommandEncoderWithDescriptor:renderPassDescriptor];
    [renderEncoder pushDebugGroup:@"Dear ImGui rendering"];
    ImGui_ImplMetal_RenderDrawData(draw_data, commandBuffer, renderEncoder);

    if (io.ConfigFlags & ImGuiConfigFlags_ViewportsEnable)
    {
        ImGui::UpdatePlatformWindows();
        ImGui::RenderPlatformWindowsDefault();
    }

    [renderEncoder popDebugGroup];
    [renderEncoder endEncoding];

    [commandBuffer presentDrawable:view.currentDrawable];
    [commandBuffer commit];

    if (ctx->afterRender != NULL)
        ctx->afterRender();
}

- (void)mtkView:(MTKView*)view drawableSizeWillChange:(CGSize)size
{
    MetalBackendContext* ctx = self.ctx;
    if (ctx == nullptr)
        return;

    if (ctx->hasSizeCallback)
        sizeCallback((__bridge void*)ctx->window, (int)size.width, (int)size.height);
}

- (void)windowWillClose:(NSNotification*)notification
{
    MetalBackendContext* ctx = self.ctx;
    if (ctx == nullptr)
        return;

    if (ctx->hasCloseCallback)
        closeCallback((__bridge void*)ctx->window);
}

- (void)windowDidBecomeKey:(NSNotification*)notification
{
    MetalBackendContext* ctx = self.ctx;
    if (ctx == nullptr)
        return;

    // After app reactivation, the timer may still be coalesced at a low rate.
    // Recreating the timer and forcing one frame avoids delayed first repaint.
    igMetalResetTimer(ctx);
    if (ctx->view != nil)
        [ctx->view setNeedsDisplay:YES];
}

@end

@interface MetalAppDelegate : NSObject<NSApplicationDelegate>
@end

@implementation MetalAppDelegate

- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication*)sender
{
    return YES;
}

@end

@interface MetalTimerTarget : NSObject
@property (nonatomic, assign) MetalBackendContext* ctx;
- (void)tick:(NSTimer*)timer;
@end

@implementation MetalTimerTarget

- (void)tick:(NSTimer*)timer
{
    if (self.ctx == nullptr || self.ctx->view == nil)
        return;
    if (!NSApp.isActive)
        return;
    if (self.ctx->window == nil || (self.ctx->window.occlusionState & NSWindowOcclusionStateVisible) == 0)
        return;

    [self.ctx->view setNeedsDisplay:YES];
}

@end

static void igMetalResetTimer(MetalBackendContext* ctx)
{
    if (ctx == nullptr)
        return;

    if (ctx->timer != nil)
    {
        [ctx->timer invalidate];
        ctx->timer = nil;
    }

    if (ctx->timerTarget == nil)
    {
        MetalTimerTarget* target = [[MetalTimerTarget alloc] init];
        target.ctx = ctx;
        ctx->timerTarget = target;
    }

    unsigned int fps = ctx->targetFPS == 0 ? 60 : ctx->targetFPS;
    NSTimeInterval interval = 1.0 / (double)fps;
    ctx->timer = [NSTimer timerWithTimeInterval:interval
                                         target:ctx->timerTarget
                                       selector:@selector(tick:)
                                       userInfo:nil
                                        repeats:YES];
    [[NSRunLoop mainRunLoop] addTimer:ctx->timer forMode:NSRunLoopCommonModes];
}

int igInitMetal(void)
{
    @autoreleasepool
    {
        [NSApplication sharedApplication];
        [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
    }
    return 1;
}

MetalBackendContext* igCreateMetalWindow(const char* title, int width, int height)
{
    MetalBackendContext* ctx = (MetalBackendContext*)calloc(1, sizeof(MetalBackendContext));
    if (ctx == nullptr)
        return nullptr;

    ctx->targetFPS = 60;
    ctx->clearColor.x = 0.45f;
    ctx->clearColor.y = 0.55f;
    ctx->clearColor.z = 0.60f;
    ctx->clearColor.w = 1.00f;

    ctx->device = MTLCreateSystemDefaultDevice();
    if (ctx->device == nil)
    {
        free(ctx);
        return nullptr;
    }

    ctx->commandQueue = [ctx->device newCommandQueue];

    NSRect frame = NSMakeRect(0, 0, width, height);
    ctx->view = [[MetalView alloc] initWithFrame:frame device:ctx->device];
    ctx->view.ctx = ctx;
    ctx->view.enableSetNeedsDisplay = YES;
    ctx->view.paused = YES;

    MetalViewDelegate* viewDelegate = [[MetalViewDelegate alloc] init];
    viewDelegate.ctx = ctx;
    ctx->viewDelegate = viewDelegate;
    ctx->view.delegate = viewDelegate;

    NSWindowStyleMask style = NSWindowStyleMaskTitled | NSWindowStyleMaskClosable | NSWindowStyleMaskResizable | NSWindowStyleMaskMiniaturizable;
    ctx->window = [[NSWindow alloc] initWithContentRect:frame
                                              styleMask:style
                                                backing:NSBackingStoreBuffered
                                                  defer:NO];

    ctx->window.contentView = ctx->view;
    ctx->window.delegate = viewDelegate;
    [ctx->window center];

    if (title != NULL)
        ctx->window.title = [NSString stringWithUTF8String:title];

    MetalAppDelegate* appDelegate = [[MetalAppDelegate alloc] init];
    ctx->appDelegate = appDelegate;
    [NSApp setDelegate:appDelegate];

    [ctx->window makeKeyAndOrderFront:nil];
    [NSApp activateIgnoringOtherApps:YES];

    IMGUI_CHECKVERSION();
    ImGui::CreateContext();

    ImGuiIO& io = ImGui::GetIO();
    io.ConfigFlags |= ImGuiConfigFlags_NavEnableKeyboard;
    io.ConfigFlags |= ImGuiConfigFlags_NavEnableGamepad;
    io.ConfigFlags |= ImGuiConfigFlags_DockingEnable;
    io.ConfigFlags |= ImGuiConfigFlags_ViewportsEnable;
    io.IniFilename = "";

    ImGui::StyleColorsDark();

    ImGuiStyle& styleRef = ImGui::GetStyle();
    if (io.ConfigFlags & ImGuiConfigFlags_ViewportsEnable)
    {
        styleRef.WindowRounding = 0.0f;
        styleRef.Colors[ImGuiCol_WindowBg].w = 1.0f;
    }

    ImGui_ImplMetal_Init(ctx->device);
    ImGui_ImplOSX_Init(ctx->view);

    return ctx;
}

void igMetalRunLoop(MetalBackendContext* ctx, VoidCallback loop, VoidCallback beforeRender, VoidCallback afterRender,
                    VoidCallback beforeDestroyContext)
{
    if (ctx == nullptr)
        return;

    ctx->loop = loop;
    ctx->beforeRender = beforeRender;
    ctx->afterRender = afterRender;
    ctx->beforeDestroyContext = beforeDestroyContext;

    igMetalResetTimer(ctx);

    [NSApp run];

    if (ctx->timer != nil)
    {
        [ctx->timer invalidate];
        ctx->timer = nil;
    }
    ctx->timerTarget = nil;

    ImGui_ImplMetal_Shutdown();
    ImGui_ImplOSX_Shutdown();

    if (ctx->beforeDestroyContext != NULL)
        ctx->beforeDestroyContext();

    ImGui::DestroyContext();

    ctx->view.delegate = nil;
    ctx->window.delegate = nil;

    ctx->view = nil;
    ctx->window = nil;
    ctx->commandQueue = nil;
    ctx->device = nil;
    ctx->viewDelegate = nil;
    ctx->appDelegate = nil;

    free(ctx);
}

void igMetalSetBgColor(MetalBackendContext* ctx, ImVec4 color)
{
    if (ctx == nullptr)
        return;
    ctx->clearColor = color;
}

void igMetalSetTargetFPS(MetalBackendContext* ctx, unsigned int fps)
{
    if (ctx == nullptr)
        return;
    ctx->targetFPS = fps;
    if (ctx->timer != nil)
        igMetalResetTimer(ctx);
}

void igMetalRefresh(MetalBackendContext* ctx)
{
    if (ctx == nullptr || ctx->view == nil)
        return;
    [ctx->view setNeedsDisplay:YES];
}

void igMetalWindow_GetDisplaySize(MetalBackendContext* ctx, int* width, int* height)
{
    if (ctx == nullptr || ctx->view == nil || width == NULL || height == NULL)
        return;
    NSSize size = ctx->view.bounds.size;
    *width = (int)size.width;
    *height = (int)size.height;
}

void igMetalWindow_GetContentScale(MetalBackendContext* ctx, float* width, float* height)
{
    if (ctx == nullptr || ctx->window == nil || width == NULL || height == NULL)
        return;
    CGFloat scale = ctx->window.backingScaleFactor;
    *width = (float)scale;
    *height = (float)scale;
}

void igMetalWindow_SetWindowPos(MetalBackendContext* ctx, int x, int y)
{
    if (ctx == nullptr || ctx->window == nil)
        return;
    [ctx->window setFrameOrigin:NSMakePoint(x, y)];
}

void igMetalWindow_GetWindowPos(MetalBackendContext* ctx, int* x, int* y)
{
    if (ctx == nullptr || ctx->window == nil || x == NULL || y == NULL)
        return;
    NSPoint pos = ctx->window.frame.origin;
    *x = (int)pos.x;
    *y = (int)pos.y;
}

void igMetalWindow_SetSize(MetalBackendContext* ctx, int width, int height)
{
    if (ctx == nullptr || ctx->window == nil)
        return;
    [ctx->window setContentSize:NSMakeSize(width, height)];
}

void igMetalWindow_SetSizeLimits(MetalBackendContext* ctx, int minWidth, int minHeight, int maxWidth, int maxHeight)
{
    if (ctx == nullptr || ctx->window == nil)
        return;

    if (minWidth >= 0 && minHeight >= 0)
        ctx->window.contentMinSize = NSMakeSize(minWidth, minHeight);

    if (maxWidth >= 0 && maxHeight >= 0)
        ctx->window.contentMaxSize = NSMakeSize(maxWidth, maxHeight);
}

void igMetalWindow_SetTitle(MetalBackendContext* ctx, const char* title)
{
    if (ctx == nullptr || ctx->window == nil || title == NULL)
        return;
    ctx->window.title = [NSString stringWithUTF8String:title];
}

void igMetalWindow_SetShouldClose(MetalBackendContext* ctx, bool value)
{
    if (ctx == nullptr || ctx->window == nil)
        return;
    if (value)
        [ctx->window performClose:nil];
}

void igMetalWindow_SetCloseCallback(MetalBackendContext* ctx)
{
    if (ctx == nullptr)
        return;
    ctx->hasCloseCallback = true;
}

void igMetalWindow_SetDropCallback(MetalBackendContext* ctx)
{
    if (ctx == nullptr || ctx->view == nil)
        return;
    ctx->hasDropCallback = true;
    [ctx->view registerForDraggedTypes:@[NSPasteboardTypeFileURL]];
}

void igMetalWindow_SetSizeCallback(MetalBackendContext* ctx)
{
    if (ctx == nullptr)
        return;
    ctx->hasSizeCallback = true;
}

void igMetalWindow_SetFlag(MetalBackendContext* ctx, MetalWindowFlag flag, int value)
{
    if (ctx == nullptr || ctx->window == nil)
        return;

    NSWindowStyleMask mask = ctx->window.styleMask;
    switch (flag)
    {
        case MetalWindowFlagResizable:
            if (value)
                mask |= NSWindowStyleMaskResizable;
            else
                mask &= ~NSWindowStyleMaskResizable;
            break;
        case MetalWindowFlagTitled:
            if (value)
                mask |= NSWindowStyleMaskTitled;
            else
                mask &= ~NSWindowStyleMaskTitled;
            break;
        case MetalWindowFlagClosable:
            if (value)
                mask |= NSWindowStyleMaskClosable;
            else
                mask &= ~NSWindowStyleMaskClosable;
            break;
        case MetalWindowFlagMiniaturizable:
            if (value)
                mask |= NSWindowStyleMaskMiniaturizable;
            else
                mask &= ~NSWindowStyleMaskMiniaturizable;
            break;
        case MetalWindowFlagFullSizeContentView:
            if (value)
                mask |= NSWindowStyleMaskFullSizeContentView;
            else
                mask &= ~NSWindowStyleMaskFullSizeContentView;
            break;
        case MetalWindowFlagTransparent:
            ctx->window.opaque = !value;
            ctx->window.backgroundColor = value ? [NSColor clearColor] : [NSColor windowBackgroundColor];
            break;
        default:
            break;
    }

    ctx->window.styleMask = mask;
}

ImTextureID igMetalCreateTexture(MetalBackendContext* ctx, unsigned char* pixels, int width, int height)
{
    if (ctx == nullptr || ctx->device == nil || pixels == NULL || width <= 0 || height <= 0)
        return (ImTextureID)0;

    MTLTextureDescriptor* textureDescriptor = [MTLTextureDescriptor texture2DDescriptorWithPixelFormat:MTLPixelFormatRGBA8Unorm
                                                                                                  width:(NSUInteger)width
                                                                                                 height:(NSUInteger)height
                                                                                              mipmapped:NO];
    textureDescriptor.usage = MTLTextureUsageShaderRead;
#if TARGET_OS_OSX || TARGET_OS_MACCATALYST
    textureDescriptor.storageMode = MTLStorageModeManaged;
#else
    textureDescriptor.storageMode = MTLStorageModeShared;
#endif

    id<MTLTexture> texture = [ctx->device newTextureWithDescriptor:textureDescriptor];
    if (texture == nil)
        return (ImTextureID)0;

    [texture replaceRegion:MTLRegionMake2D(0, 0, (NSUInteger)width, (NSUInteger)height)
               mipmapLevel:0
                 withBytes:pixels
               bytesPerRow:(NSUInteger)width * 4];

    return (ImTextureID)(intptr_t)(__bridge_retained void*)texture;
}

void igMetalDeleteTexture(ImTextureID texture_id)
{
    if (texture_id == 0)
        return;

    id<MTLTexture> texture = (__bridge_transfer id<MTLTexture>)(void*)(intptr_t)texture_id;
    texture = nil;
}

#endif // CIMGUI_GO_USE_METAL
