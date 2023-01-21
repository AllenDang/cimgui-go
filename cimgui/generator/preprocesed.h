
<<<<<<< HEAD
               bool ImGui_ImplOpenGL3_Init(const char* glsl_version = nullptr);
               void ImGui_ImplOpenGL3_Shutdown();
               void ImGui_ImplOpenGL3_NewFrame();
               void ImGui_ImplOpenGL3_RenderDrawData(ImDrawData* draw_data);
               bool ImGui_ImplOpenGL3_CreateFontsTexture();
               void ImGui_ImplOpenGL3_DestroyFontsTexture();
               bool ImGui_ImplOpenGL3_CreateDeviceObjects();
               void ImGui_ImplOpenGL3_DestroyDeviceObjects();
=======
struct SDL_Window;
struct SDL_Renderer;
typedef union SDL_Event SDL_Event;
 bool ImGui_ImplSDL2_InitForOpenGL(SDL_Window* window, void* sdl_gl_context);
 bool ImGui_ImplSDL2_InitForVulkan(SDL_Window* window);
 bool ImGui_ImplSDL2_InitForD3D(SDL_Window* window);
 bool ImGui_ImplSDL2_InitForMetal(SDL_Window* window);
 bool ImGui_ImplSDL2_InitForSDLRenderer(SDL_Window* window, SDL_Renderer* renderer);
 void ImGui_ImplSDL2_Shutdown();
 void ImGui_ImplSDL2_NewFrame();
 bool ImGui_ImplSDL2_ProcessEvent(const SDL_Event* event);
>>>>>>> aff58f6 (code: udpate to the lates version of cimgui and cimplot)
