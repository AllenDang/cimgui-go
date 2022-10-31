
struct ImPlotContext;
typedef int ImAxis;
typedef int ImPlotFlags;
typedef int ImPlotAxisFlags;
typedef int ImPlotSubplotFlags;
typedef int ImPlotLegendFlags;
typedef int ImPlotMouseTextFlags;
typedef int ImPlotDragToolFlags;
typedef int ImPlotColormapScaleFlags;
typedef int ImPlotItemFlags;
typedef int ImPlotLineFlags;
typedef int ImPlotScatterFlags;
typedef int ImPlotStairsFlags;
typedef int ImPlotShadedFlags;
typedef int ImPlotBarsFlags;
typedef int ImPlotBarGroupsFlags;
typedef int ImPlotErrorBarsFlags;
typedef int ImPlotStemsFlags;
typedef int ImPlotInfLinesFlags;
typedef int ImPlotPieChartFlags;
typedef int ImPlotHeatmapFlags;
typedef int ImPlotHistogramFlags;
typedef int ImPlotDigitalFlags;
typedef int ImPlotImageFlags;
typedef int ImPlotTextFlags;
typedef int ImPlotDummyFlags;
typedef int ImPlotCond;
typedef int ImPlotCol;
typedef int ImPlotStyleVar;
typedef int ImPlotScale;
typedef int ImPlotMarker;
typedef int ImPlotColormap;
typedef int ImPlotLocation;
typedef int ImPlotBin;
enum ImAxis_ {
    ImAxis_X1 = 0,
    ImAxis_X2,
    ImAxis_X3,
    ImAxis_Y1,
    ImAxis_Y2,
    ImAxis_Y3,
    ImAxis_COUNT
};
enum ImPlotFlags_ {
    ImPlotFlags_None = 0,
    ImPlotFlags_NoTitle = 1 << 0,
    ImPlotFlags_NoLegend = 1 << 1,
    ImPlotFlags_NoMouseText = 1 << 2,
    ImPlotFlags_NoInputs = 1 << 3,
    ImPlotFlags_NoMenus = 1 << 4,
    ImPlotFlags_NoBoxSelect = 1 << 5,
    ImPlotFlags_NoChild = 1 << 6,
    ImPlotFlags_NoFrame = 1 << 7,
    ImPlotFlags_Equal = 1 << 8,
    ImPlotFlags_Crosshairs = 1 << 9,
    ImPlotFlags_CanvasOnly = ImPlotFlags_NoTitle | ImPlotFlags_NoLegend | ImPlotFlags_NoMenus | ImPlotFlags_NoBoxSelect | ImPlotFlags_NoMouseText
};
enum ImPlotAxisFlags_ {
    ImPlotAxisFlags_None = 0,
    ImPlotAxisFlags_NoLabel = 1 << 0,
    ImPlotAxisFlags_NoGridLines = 1 << 1,
    ImPlotAxisFlags_NoTickMarks = 1 << 2,
    ImPlotAxisFlags_NoTickLabels = 1 << 3,
    ImPlotAxisFlags_NoInitialFit = 1 << 4,
    ImPlotAxisFlags_NoMenus = 1 << 5,
    ImPlotAxisFlags_NoSideSwitch = 1 << 6,
    ImPlotAxisFlags_NoHighlight = 1 << 7,
    ImPlotAxisFlags_Opposite = 1 << 8,
    ImPlotAxisFlags_Foreground = 1 << 9,
    ImPlotAxisFlags_Invert = 1 << 10,
    ImPlotAxisFlags_AutoFit = 1 << 11,
    ImPlotAxisFlags_RangeFit = 1 << 12,
    ImPlotAxisFlags_PanStretch = 1 << 13,
    ImPlotAxisFlags_LockMin = 1 << 14,
    ImPlotAxisFlags_LockMax = 1 << 15,
    ImPlotAxisFlags_Lock = ImPlotAxisFlags_LockMin | ImPlotAxisFlags_LockMax,
    ImPlotAxisFlags_NoDecorations = ImPlotAxisFlags_NoLabel | ImPlotAxisFlags_NoGridLines | ImPlotAxisFlags_NoTickMarks | ImPlotAxisFlags_NoTickLabels,
    ImPlotAxisFlags_AuxDefault = ImPlotAxisFlags_NoGridLines | ImPlotAxisFlags_Opposite
};
enum ImPlotSubplotFlags_ {
    ImPlotSubplotFlags_None = 0,
    ImPlotSubplotFlags_NoTitle = 1 << 0,
    ImPlotSubplotFlags_NoLegend = 1 << 1,
    ImPlotSubplotFlags_NoMenus = 1 << 2,
    ImPlotSubplotFlags_NoResize = 1 << 3,
    ImPlotSubplotFlags_NoAlign = 1 << 4,
    ImPlotSubplotFlags_ShareItems = 1 << 5,
    ImPlotSubplotFlags_LinkRows = 1 << 6,
    ImPlotSubplotFlags_LinkCols = 1 << 7,
    ImPlotSubplotFlags_LinkAllX = 1 << 8,
    ImPlotSubplotFlags_LinkAllY = 1 << 9,
    ImPlotSubplotFlags_ColMajor = 1 << 10
};
enum ImPlotLegendFlags_ {
    ImPlotLegendFlags_None = 0,
    ImPlotLegendFlags_NoButtons = 1 << 0,
    ImPlotLegendFlags_NoHighlightItem = 1 << 1,
    ImPlotLegendFlags_NoHighlightAxis = 1 << 2,
    ImPlotLegendFlags_NoMenus = 1 << 3,
    ImPlotLegendFlags_Outside = 1 << 4,
    ImPlotLegendFlags_Horizontal = 1 << 5,
    ImPlotLegendFlags_Sort = 1 << 6,
};
enum ImPlotMouseTextFlags_ {
    ImPlotMouseTextFlags_None = 0,
    ImPlotMouseTextFlags_NoAuxAxes = 1 << 0,
    ImPlotMouseTextFlags_NoFormat = 1 << 1,
    ImPlotMouseTextFlags_ShowAlways = 1 << 2,
};
enum ImPlotDragToolFlags_ {
    ImPlotDragToolFlags_None = 0,
    ImPlotDragToolFlags_NoCursors = 1 << 0,
    ImPlotDragToolFlags_NoFit = 1 << 1,
    ImPlotDragToolFlags_NoInputs = 1 << 2,
    ImPlotDragToolFlags_Delayed = 1 << 3,
};
enum ImPlotColormapScaleFlags_ {
    ImPlotColormapScaleFlags_None = 0,
    ImPlotColormapScaleFlags_NoLabel = 1 << 0,
    ImPlotColormapScaleFlags_Opposite = 1 << 1,
    ImPlotColormapScaleFlags_Invert = 1 << 2,
};
enum ImPlotItemFlags_ {
    ImPlotItemFlags_None = 0,
    ImPlotItemFlags_NoLegend = 1 << 0,
    ImPlotItemFlags_NoFit = 1 << 1,
};
enum ImPlotLineFlags_ {
    ImPlotLineFlags_None = 0,
    ImPlotLineFlags_Segments = 1 << 10,
    ImPlotLineFlags_Loop = 1 << 11,
    ImPlotLineFlags_SkipNaN = 1 << 12,
    ImPlotLineFlags_NoClip = 1 << 13,
    ImPlotLineFlags_Shaded = 1 << 14,
};
enum ImPlotScatterFlags_ {
    ImPlotScatterFlags_None = 0,
    ImPlotScatterFlags_NoClip = 1 << 10,
};
enum ImPlotStairsFlags_ {
    ImPlotStairsFlags_None = 0,
    ImPlotStairsFlags_PreStep = 1 << 10,
    ImPlotStairsFlags_Shaded = 1 << 11
};
enum ImPlotShadedFlags_ {
    ImPlotShadedFlags_None = 0
};
enum ImPlotBarsFlags_ {
    ImPlotBarsFlags_None = 0,
    ImPlotBarsFlags_Horizontal = 1 << 10,
};
enum ImPlotBarGroupsFlags_ {
    ImPlotBarGroupsFlags_None = 0,
    ImPlotBarGroupsFlags_Horizontal = 1 << 10,
    ImPlotBarGroupsFlags_Stacked = 1 << 11,
};
enum ImPlotErrorBarsFlags_ {
    ImPlotErrorBarsFlags_None = 0,
    ImPlotErrorBarsFlags_Horizontal = 1 << 10,
};
enum ImPlotStemsFlags_ {
    ImPlotStemsFlags_None = 0,
    ImPlotStemsFlags_Horizontal = 1 << 10,
};
enum ImPlotInfLinesFlags_ {
    ImPlotInfLinesFlags_None = 0,
    ImPlotInfLinesFlags_Horizontal = 1 << 10
};
enum ImPlotPieChartFlags_ {
    ImPlotPieChartFlags_None = 0,
    ImPlotPieChartFlags_Normalize = 1 << 10
};
enum ImPlotHeatmapFlags_ {
    ImPlotHeatmapFlags_None = 0,
    ImPlotHeatmapFlags_ColMajor = 1 << 10,
};
enum ImPlotHistogramFlags_ {
    ImPlotHistogramFlags_None = 0,
    ImPlotHistogramFlags_Horizontal = 1 << 10,
    ImPlotHistogramFlags_Cumulative = 1 << 11,
    ImPlotHistogramFlags_Density = 1 << 12,
    ImPlotHistogramFlags_NoOutliers = 1 << 13,
    ImPlotHistogramFlags_ColMajor = 1 << 14
};
enum ImPlotDigitalFlags_ {
    ImPlotDigitalFlags_None = 0
};
enum ImPlotImageFlags_ {
    ImPlotImageFlags_None = 0
};
enum ImPlotTextFlags_ {
    ImPlotTextFlags_None = 0,
    ImPlotTextFlags_Vertical = 1 << 10
};
enum ImPlotDummyFlags_ {
    ImPlotDummyFlags_None = 0
};
enum ImPlotCond_
{
    ImPlotCond_None = ImGuiCond_None,
    ImPlotCond_Always = ImGuiCond_Always,
    ImPlotCond_Once = ImGuiCond_Once,
};
enum ImPlotCol_ {
    ImPlotCol_Line,
    ImPlotCol_Fill,
    ImPlotCol_MarkerOutline,
    ImPlotCol_MarkerFill,
    ImPlotCol_ErrorBar,
    ImPlotCol_FrameBg,
    ImPlotCol_PlotBg,
    ImPlotCol_PlotBorder,
    ImPlotCol_LegendBg,
    ImPlotCol_LegendBorder,
    ImPlotCol_LegendText,
    ImPlotCol_TitleText,
    ImPlotCol_InlayText,
    ImPlotCol_AxisText,
    ImPlotCol_AxisGrid,
    ImPlotCol_AxisTick,
    ImPlotCol_AxisBg,
    ImPlotCol_AxisBgHovered,
    ImPlotCol_AxisBgActive,
    ImPlotCol_Selection,
    ImPlotCol_Crosshairs,
    ImPlotCol_COUNT
};
enum ImPlotStyleVar_ {
    ImPlotStyleVar_LineWeight,
    ImPlotStyleVar_Marker,
    ImPlotStyleVar_MarkerSize,
    ImPlotStyleVar_MarkerWeight,
    ImPlotStyleVar_FillAlpha,
    ImPlotStyleVar_ErrorBarSize,
    ImPlotStyleVar_ErrorBarWeight,
    ImPlotStyleVar_DigitalBitHeight,
    ImPlotStyleVar_DigitalBitGap,
    ImPlotStyleVar_PlotBorderSize,
    ImPlotStyleVar_MinorAlpha,
    ImPlotStyleVar_MajorTickLen,
    ImPlotStyleVar_MinorTickLen,
    ImPlotStyleVar_MajorTickSize,
    ImPlotStyleVar_MinorTickSize,
    ImPlotStyleVar_MajorGridSize,
    ImPlotStyleVar_MinorGridSize,
    ImPlotStyleVar_PlotPadding,
    ImPlotStyleVar_LabelPadding,
    ImPlotStyleVar_LegendPadding,
    ImPlotStyleVar_LegendInnerPadding,
    ImPlotStyleVar_LegendSpacing,
    ImPlotStyleVar_MousePosPadding,
    ImPlotStyleVar_AnnotationPadding,
    ImPlotStyleVar_FitPadding,
    ImPlotStyleVar_PlotDefaultSize,
    ImPlotStyleVar_PlotMinSize,
    ImPlotStyleVar_COUNT
};
enum ImPlotScale_ {
    ImPlotScale_Linear = 0,
    ImPlotScale_Time,
    ImPlotScale_Log10,
    ImPlotScale_SymLog,
};
enum ImPlotMarker_ {
    ImPlotMarker_None = -1,
    ImPlotMarker_Circle,
    ImPlotMarker_Square,
    ImPlotMarker_Diamond,
    ImPlotMarker_Up,
    ImPlotMarker_Down,
    ImPlotMarker_Left,
    ImPlotMarker_Right,
    ImPlotMarker_Cross,
    ImPlotMarker_Plus,
    ImPlotMarker_Asterisk,
    ImPlotMarker_COUNT
};
enum ImPlotColormap_ {
    ImPlotColormap_Deep = 0,
    ImPlotColormap_Dark = 1,
    ImPlotColormap_Pastel = 2,
    ImPlotColormap_Paired = 3,
    ImPlotColormap_Viridis = 4,
    ImPlotColormap_Plasma = 5,
    ImPlotColormap_Hot = 6,
    ImPlotColormap_Cool = 7,
    ImPlotColormap_Pink = 8,
    ImPlotColormap_Jet = 9,
    ImPlotColormap_Twilight = 10,
    ImPlotColormap_RdBu = 11,
    ImPlotColormap_BrBG = 12,
    ImPlotColormap_PiYG = 13,
    ImPlotColormap_Spectral = 14,
    ImPlotColormap_Greys = 15,
};
enum ImPlotLocation_ {
    ImPlotLocation_Center = 0,
    ImPlotLocation_North = 1 << 0,
    ImPlotLocation_South = 1 << 1,
    ImPlotLocation_West = 1 << 2,
    ImPlotLocation_East = 1 << 3,
    ImPlotLocation_NorthWest = ImPlotLocation_North | ImPlotLocation_West,
    ImPlotLocation_NorthEast = ImPlotLocation_North | ImPlotLocation_East,
    ImPlotLocation_SouthWest = ImPlotLocation_South | ImPlotLocation_West,
    ImPlotLocation_SouthEast = ImPlotLocation_South | ImPlotLocation_East
};
enum ImPlotBin_ {
    ImPlotBin_Sqrt = -1,
    ImPlotBin_Sturges = -2,
    ImPlotBin_Rice = -3,
    ImPlotBin_Scott = -4,
};
struct ImPlotPoint {
    double x, y;
    ImPlotPoint() { x = y = 0.0; }
    ImPlotPoint(double _x, double _y) { x = _x; y = _y; }
    ImPlotPoint(const ImVec2& p) { x = p.x; y = p.y; }
    double operator[] (size_t idx) const { return (&x)[idx]; }
    double& operator[] (size_t idx) { return (&x)[idx]; }
};
struct ImPlotRange {
    double Min, Max;
    ImPlotRange() { Min = 0; Max = 0; }
    ImPlotRange(double _min, double _max) { Min = _min; Max = _max; }
    bool Contains(double value) const { return value >= Min && value <= Max; }
    double Size() const { return Max - Min; }
    double Clamp(double value) const { return (value < Min) ? Min : (value > Max) ? Max : value; }
};
struct ImPlotRect {
    ImPlotRange X, Y;
    ImPlotRect() { }
    ImPlotRect(double x_min, double x_max, double y_min, double y_max) { X.Min = x_min; X.Max = x_max; Y.Min = y_min; Y.Max = y_max; }
    bool Contains(const ImPlotPoint& p) const { return Contains(p.x, p.y); }
    bool Contains(double x, double y) const { return X.Contains(x) && Y.Contains(y); }
    ImPlotPoint Size() const { return ImPlotPoint(X.Size(), Y.Size()); }
    ImPlotPoint Clamp(const ImPlotPoint& p) { return Clamp(p.x, p.y); }
    ImPlotPoint Clamp(double x, double y) { return ImPlotPoint(X.Clamp(x),Y.Clamp(y)); }
    ImPlotPoint Min() const { return ImPlotPoint(X.Min, Y.Min); }
    ImPlotPoint Max() const { return ImPlotPoint(X.Max, Y.Max); }
};
struct ImPlotStyle {
    float LineWeight;
    int Marker;
    float MarkerSize;
    float MarkerWeight;
    float FillAlpha;
    float ErrorBarSize;
    float ErrorBarWeight;
    float DigitalBitHeight;
    float DigitalBitGap;
    float PlotBorderSize;
    float MinorAlpha;
    ImVec2 MajorTickLen;
    ImVec2 MinorTickLen;
    ImVec2 MajorTickSize;
    ImVec2 MinorTickSize;
    ImVec2 MajorGridSize;
    ImVec2 MinorGridSize;
    ImVec2 PlotPadding;
    ImVec2 LabelPadding;
    ImVec2 LegendPadding;
    ImVec2 LegendInnerPadding;
    ImVec2 LegendSpacing;
    ImVec2 MousePosPadding;
    ImVec2 AnnotationPadding;
    ImVec2 FitPadding;
    ImVec2 PlotDefaultSize;
    ImVec2 PlotMinSize;
    ImVec4 Colors[ImPlotCol_COUNT];
    ImPlotColormap Colormap;
    bool UseLocalTime;
    bool UseISO8601;
    bool Use24HourClock;
    ImPlotStyle();
};
struct ImPlotInputMap {
    ImGuiMouseButton Pan;
    int PanMod;
    ImGuiMouseButton Fit;
    ImGuiMouseButton Select;
    ImGuiMouseButton SelectCancel;
    int SelectMod;
    int SelectHorzMod;
    int SelectVertMod;
    ImGuiMouseButton Menu;
    int OverrideMod;
    int ZoomMod;
    float ZoomRate;
    ImPlotInputMap();
};
typedef int (*ImPlotFormatter)(double value, char* buff, int size, void* user_data);
typedef ImPlotPoint (*ImPlotGetter)(int idx, void* user_data);
typedef double (*ImPlotTransform)(double value, void* user_data);
namespace ImPlot {
 ImPlotContext* CreateContext();
 void DestroyContext(ImPlotContext* ctx =                                                    ((void *)0)                                                       );
 ImPlotContext* GetCurrentContext();
 void SetCurrentContext(ImPlotContext* ctx);
 void SetImGuiContext(ImGuiContext* ctx);
 bool BeginPlot(const char* title_id, const ImVec2& size=ImVec2(-1,0), ImPlotFlags flags=0);
 void EndPlot();
 bool BeginSubplots(const char* title_id,
                             int rows,
                             int cols,
                             const ImVec2& size,
                             ImPlotSubplotFlags flags = 0,
                             float* row_ratios =                                                        ((void *)0)                                                           ,
                             float* col_ratios =                                                        ((void *)0)                                                           );
 void EndSubplots();
 void SetupAxis(ImAxis axis, const char* label=                                                        ((void *)0)                                                            , ImPlotAxisFlags flags=0);
 void SetupAxisLimits(ImAxis axis, double v_min, double v_max, ImPlotCond cond = ImPlotCond_Once);
 void SetupAxisLinks(ImAxis axis, double* link_min, double* link_max);
 void SetupAxisFormat(ImAxis axis, const char* fmt);
 void SetupAxisFormat(ImAxis axis, ImPlotFormatter formatter, void* data=                                                                                  ((void *)0)                                                                                      );
 void SetupAxisTicks(ImAxis axis, const double* values, int n_ticks, const char* const labels[]=                                                                                                         ((void *)0)                                                                                                             , bool keep_default=false);
 void SetupAxisTicks(ImAxis axis, double v_min, double v_max, int n_ticks, const char* const labels[]=                                                                                                               ((void *)0)                                                                                                                   , bool keep_default=false);
 void SetupAxisScale(ImAxis axis, ImPlotScale scale);
 void SetupAxisScale(ImAxis axis, ImPlotTransform forward, ImPlotTransform inverse, void* data=                                                                                                        ((void *)0)                                                                                                            );
 void SetupAxisLimitsConstraints(ImAxis axis, double v_min, double v_max);
 void SetupAxisZoomConstraints(ImAxis axis, double z_min, double z_max);
 void SetupAxes(const char* x_label, const char* y_label, ImPlotAxisFlags x_flags=0, ImPlotAxisFlags y_flags=0);
 void SetupAxesLimits(double x_min, double x_max, double y_min, double y_max, ImPlotCond cond = ImPlotCond_Once);
 void SetupLegend(ImPlotLocation location, ImPlotLegendFlags flags=0);
 void SetupMouseText(ImPlotLocation location, ImPlotMouseTextFlags flags=0);
 void SetupFinish();
 void SetNextAxisLimits(ImAxis axis, double v_min, double v_max, ImPlotCond cond = ImPlotCond_Once);
 void SetNextAxisLinks(ImAxis axis, double* link_min, double* link_max);
 void SetNextAxisToFit(ImAxis axis);
 void SetNextAxesLimits(double x_min, double x_max, double y_min, double y_max, ImPlotCond cond = ImPlotCond_Once);
 void SetNextAxesToFit();
template <typename T> void PlotLine(const char* label_id, const T* values, int count, double xscale=1, double xstart=0, ImPlotLineFlags flags=0, int offset=0, int stride=sizeof(T));
template <typename T> void PlotLine(const char* label_id, const T* xs, const T* ys, int count, ImPlotLineFlags flags=0, int offset=0, int stride=sizeof(T));
 void PlotLineG(const char* label_id, ImPlotGetter getter, void* data, int count, ImPlotLineFlags flags=0);
template <typename T> void PlotScatter(const char* label_id, const T* values, int count, double xscale=1, double xstart=0, ImPlotScatterFlags flags=0, int offset=0, int stride=sizeof(T));
template <typename T> void PlotScatter(const char* label_id, const T* xs, const T* ys, int count, ImPlotScatterFlags flags=0, int offset=0, int stride=sizeof(T));
 void PlotScatterG(const char* label_id, ImPlotGetter getter, void* data, int count, ImPlotScatterFlags flags=0);
template <typename T> void PlotStairs(const char* label_id, const T* values, int count, double xscale=1, double xstart=0, ImPlotStairsFlags flags=0, int offset=0, int stride=sizeof(T));
template <typename T> void PlotStairs(const char* label_id, const T* xs, const T* ys, int count, ImPlotStairsFlags flags=0, int offset=0, int stride=sizeof(T));
 void PlotStairsG(const char* label_id, ImPlotGetter getter, void* data, int count, ImPlotStairsFlags flags=0);
template <typename T> void PlotShaded(const char* label_id, const T* values, int count, double yref=0, double xscale=1, double xstart=0, ImPlotShadedFlags flags=0, int offset=0, int stride=sizeof(T));
template <typename T> void PlotShaded(const char* label_id, const T* xs, const T* ys, int count, double yref=0, ImPlotShadedFlags flags=0, int offset=0, int stride=sizeof(T));
template <typename T> void PlotShaded(const char* label_id, const T* xs, const T* ys1, const T* ys2, int count, ImPlotShadedFlags flags=0, int offset=0, int stride=sizeof(T));
 void PlotShadedG(const char* label_id, ImPlotGetter getter1, void* data1, ImPlotGetter getter2, void* data2, int count, ImPlotShadedFlags flags=0);
template <typename T> void PlotBars(const char* label_id, const T* values, int count, double bar_size=0.67, double shift=0, ImPlotBarsFlags flags=0, int offset=0, int stride=sizeof(T));
template <typename T> void PlotBars(const char* label_id, const T* xs, const T* ys, int count, double bar_size, ImPlotBarsFlags flags=0, int offset=0, int stride=sizeof(T));
 void PlotBarsG(const char* label_id, ImPlotGetter getter, void* data, int count, double bar_size, ImPlotBarsFlags flags=0);
template <typename T> void PlotBarGroups(const char* const label_ids[], const T* values, int item_count, int group_count, double group_size=0.67, double shift=0, ImPlotBarGroupsFlags flags=0);
template <typename T> void PlotErrorBars(const char* label_id, const T* xs, const T* ys, const T* err, int count, ImPlotErrorBarsFlags flags=0, int offset=0, int stride=sizeof(T));
template <typename T> void PlotErrorBars(const char* label_id, const T* xs, const T* ys, const T* neg, const T* pos, int count, ImPlotErrorBarsFlags flags=0, int offset=0, int stride=sizeof(T));
template <typename T> void PlotStems(const char* label_id, const T* values, int count, double ref=0, double scale=1, double start=0, ImPlotStemsFlags flags=0, int offset=0, int stride=sizeof(T));
template <typename T> void PlotStems(const char* label_id, const T* xs, const T* ys, int count, double ref=0, ImPlotStemsFlags flags=0, int offset=0, int stride=sizeof(T));
template <typename T> void PlotInfLines(const char* label_id, const T* values, int count, ImPlotInfLinesFlags flags=0, int offset=0, int stride=sizeof(T));
template <typename T> void PlotPieChart(const char* const label_ids[], const T* values, int count, double x, double y, double radius, const char* label_fmt="%.1f", double angle0=90, ImPlotPieChartFlags flags=0);
template <typename T> void PlotHeatmap(const char* label_id, const T* values, int rows, int cols, double scale_min=0, double scale_max=0, const char* label_fmt="%.1f", const ImPlotPoint& bounds_min=ImPlotPoint(0,0), const ImPlotPoint& bounds_max=ImPlotPoint(1,1), ImPlotHeatmapFlags flags=0);
template <typename T> double PlotHistogram(const char* label_id, const T* values, int count, int bins=ImPlotBin_Sturges, double bar_scale=1.0, ImPlotRange range=ImPlotRange(), ImPlotHistogramFlags flags=0);
template <typename T> double PlotHistogram2D(const char* label_id, const T* xs, const T* ys, int count, int x_bins=ImPlotBin_Sturges, int y_bins=ImPlotBin_Sturges, ImPlotRect range=ImPlotRect(), ImPlotHistogramFlags flags=0);
template <typename T> void PlotDigital(const char* label_id, const T* xs, const T* ys, int count, ImPlotDigitalFlags flags=0, int offset=0, int stride=sizeof(T));
 void PlotDigitalG(const char* label_id, ImPlotGetter getter, void* data, int count, ImPlotDigitalFlags flags=0);
 void PlotImage(const char* label_id, ImTextureID user_texture_id, const ImPlotPoint& bounds_min, const ImPlotPoint& bounds_max, const ImVec2& uv0=ImVec2(0,0), const ImVec2& uv1=ImVec2(1,1), const ImVec4& tint_col=ImVec4(1,1,1,1), ImPlotImageFlags flags=0);
 void PlotText(const char* text, double x, double y, const ImVec2& pix_offset=ImVec2(0,0), ImPlotTextFlags flags=0);
 void PlotDummy(const char* label_id, ImPlotDummyFlags flags=0);
 bool DragPoint(int id, double* x, double* y, const ImVec4& col, float size = 4, ImPlotDragToolFlags flags=0);
 bool DragLineX(int id, double* x, const ImVec4& col, float thickness = 1, ImPlotDragToolFlags flags=0);
 bool DragLineY(int id, double* y, const ImVec4& col, float thickness = 1, ImPlotDragToolFlags flags=0);
 bool DragRect(int id, double* x1, double* y1, double* x2, double* y2, const ImVec4& col, ImPlotDragToolFlags flags=0);
 void Annotation(double x, double y, const ImVec4& col, const ImVec2& pix_offset, bool clamp, bool round = false);
 void Annotation(double x, double y, const ImVec4& col, const ImVec2& pix_offset, bool clamp, const char* fmt, ...) __attribute__((format(printf, 6, 6 +1)));
 void AnnotationV(double x, double y, const ImVec4& col, const ImVec2& pix_offset, bool clamp, const char* fmt, va_list args) __attribute__((format(printf, 6, 0)));
 void TagX(double x, const ImVec4& col, bool round = false);
 void TagX(double x, const ImVec4& col, const char* fmt, ...) __attribute__((format(printf, 3, 3 +1)));
 void TagXV(double x, const ImVec4& col, const char* fmt, va_list args) __attribute__((format(printf, 3, 0)));
 void TagY(double y, const ImVec4& col, bool round = false);
 void TagY(double y, const ImVec4& col, const char* fmt, ...) __attribute__((format(printf, 3, 3 +1)));
 void TagYV(double y, const ImVec4& col, const char* fmt, va_list args) __attribute__((format(printf, 3, 0)));
 void SetAxis(ImAxis axis);
 void SetAxes(ImAxis x_axis, ImAxis y_axis);
 ImPlotPoint PixelsToPlot(const ImVec2& pix, ImAxis x_axis = -1, ImAxis y_axis = -1);
 ImPlotPoint PixelsToPlot(float x, float y, ImAxis x_axis = -1, ImAxis y_axis = -1);
 ImVec2 PlotToPixels(const ImPlotPoint& plt, ImAxis x_axis = -1, ImAxis y_axis = -1);
 ImVec2 PlotToPixels(double x, double y, ImAxis x_axis = -1, ImAxis y_axis = -1);
 ImVec2 GetPlotPos();
 ImVec2 GetPlotSize();
 ImPlotPoint GetPlotMousePos(ImAxis x_axis = -1, ImAxis y_axis = -1);
 ImPlotRect GetPlotLimits(ImAxis x_axis = -1, ImAxis y_axis = -1);
 bool IsPlotHovered();
 bool IsAxisHovered(ImAxis axis);
 bool IsSubplotsHovered();
 bool IsPlotSelected();
 ImPlotRect GetPlotSelection(ImAxis x_axis = -1, ImAxis y_axis = -1);
 void CancelPlotSelection();
 void HideNextItem(bool hidden = true, ImPlotCond cond = ImPlotCond_Once);
 bool BeginAlignedPlots(const char* group_id, bool vertical = true);
 void EndAlignedPlots();
 bool BeginLegendPopup(const char* label_id, ImGuiMouseButton mouse_button=1);
 void EndLegendPopup();
 bool IsLegendEntryHovered(const char* label_id);
 bool BeginDragDropTargetPlot();
 bool BeginDragDropTargetAxis(ImAxis axis);
 bool BeginDragDropTargetLegend();
 void EndDragDropTarget();
 bool BeginDragDropSourcePlot(ImGuiDragDropFlags flags=0);
 bool BeginDragDropSourceAxis(ImAxis axis, ImGuiDragDropFlags flags=0);
 bool BeginDragDropSourceItem(const char* label_id, ImGuiDragDropFlags flags=0);
 void EndDragDropSource();
 ImPlotStyle& GetStyle();
 void StyleColorsAuto(ImPlotStyle* dst =                                                   ((void *)0)                                                      );
 void StyleColorsClassic(ImPlotStyle* dst =                                                      ((void *)0)                                                         );
 void StyleColorsDark(ImPlotStyle* dst =                                                   ((void *)0)                                                      );
 void StyleColorsLight(ImPlotStyle* dst =                                                    ((void *)0)                                                       );
 void PushStyleColor(ImPlotCol idx, ImU32 col);
 void PushStyleColor(ImPlotCol idx, const ImVec4& col);
 void PopStyleColor(int count = 1);
 void PushStyleVar(ImPlotStyleVar idx, float val);
 void PushStyleVar(ImPlotStyleVar idx, int val);
 void PushStyleVar(ImPlotStyleVar idx, const ImVec2& val);
 void PopStyleVar(int count = 1);
 void SetNextLineStyle(const ImVec4& col = ImVec4(0,0,0,-1), float weight = -1);
 void SetNextFillStyle(const ImVec4& col = ImVec4(0,0,0,-1), float alpha_mod = -1);
 void SetNextMarkerStyle(ImPlotMarker marker = -1, float size = -1, const ImVec4& fill = ImVec4(0,0,0,-1), float weight = -1, const ImVec4& outline = ImVec4(0,0,0,-1));
 void SetNextErrorBarStyle(const ImVec4& col = ImVec4(0,0,0,-1), float size = -1, float weight = -1);
 ImVec4 GetLastItemColor();
 const char* GetStyleColorName(ImPlotCol idx);
 const char* GetMarkerName(ImPlotMarker idx);
 ImPlotColormap AddColormap(const char* name, const ImVec4* cols, int size, bool qual=true);
 ImPlotColormap AddColormap(const char* name, const ImU32* cols, int size, bool qual=true);
 int GetColormapCount();
 const char* GetColormapName(ImPlotColormap cmap);
 ImPlotColormap GetColormapIndex(const char* name);
 void PushColormap(ImPlotColormap cmap);
 void PushColormap(const char* name);
 void PopColormap(int count = 1);
 ImVec4 NextColormapColor();
 int GetColormapSize(ImPlotColormap cmap = -1);
 ImVec4 GetColormapColor(int idx, ImPlotColormap cmap = -1);
 ImVec4 SampleColormap(float t, ImPlotColormap cmap = -1);
 void ColormapScale(const char* label, double scale_min, double scale_max, const ImVec2& size = ImVec2(0,0), const char* format = "%g", ImPlotColormapScaleFlags flags = 0, ImPlotColormap cmap = -1);
 bool ColormapSlider(const char* label, float* t, ImVec4* out =                                                                          ((void *)0)                                                                             , const char* format = "", ImPlotColormap cmap = -1);
 bool ColormapButton(const char* label, const ImVec2& size = ImVec2(0,0), ImPlotColormap cmap = -1);
 void BustColorCache(const char* plot_title_id =                                                           ((void *)0)                                                              );
 ImPlotInputMap& GetInputMap();
 void MapInputDefault(ImPlotInputMap* dst =                                                      ((void *)0)                                                         );
 void MapInputReverse(ImPlotInputMap* dst =                                                      ((void *)0)                                                         );
 void ItemIcon(const ImVec4& col);
 void ItemIcon(ImU32 col);
 void ColormapIcon(ImPlotColormap cmap);
 ImDrawList* GetPlotDrawList();
 void PushPlotClipRect(float expand=0);
 void PopPlotClipRect();
 bool ShowStyleSelector(const char* label);
 bool ShowColormapSelector(const char* label);
 bool ShowInputMapSelector(const char* label);
 void ShowStyleEditor(ImPlotStyle* ref =                                                   ((void *)0)                                                      );
 void ShowUserGuide();
 void ShowMetricsWindow(bool* p_popen =                                                  ((void *)0)                                                     );
 void ShowDemoWindow(bool* p_open =                                              ((void *)0)                                                 );
}
struct ImPlotTick;
struct ImPlotAxis;
struct ImPlotAxisColor;
struct ImPlotItem;
struct ImPlotLegend;
struct ImPlotPlot;
struct ImPlotNextPlotData;
struct ImPlotTicker;
extern ImPlotContext* GImPlot;
static inline float ImLog10(float x) { return log10f(x); }
static inline double ImLog10(double x) { return log10(x); }
static inline float ImSinh(float x) { return sinhf(x); }
static inline double ImSinh(double x) { return sinh(x); }
static inline float ImAsinh(float x) { return asinhf(x); }
static inline double ImAsinh(double x) { return asinh(x); }
template <typename TSet, typename TFlag>
static inline bool ImHasFlag(TSet set, TFlag flag) { return (set & flag) == flag; }
template <typename TSet, typename TFlag>
static inline void ImFlipFlag(TSet& set, TFlag flag) { ImHasFlag(set, flag) ? set &= ~flag : set |= flag; }
template <typename T>
static inline T ImRemap(T x, T x0, T x1, T y0, T y1) { return y0 + (x - x0) * (y1 - y0) / (x1 - x0); }
template <typename T>
static inline T ImRemap01(T x, T x0, T x1) { return (x - x0) / (x1 - x0); }
static inline int ImPosMod(int l, int r) { return (l % r + r) % r; }
static inline bool ImNan(double val) { return                                              __builtin_isnan (                                             val                                             )                                                       ; }
static inline bool ImNanOrInf(double val) { return !(val >= -((double)1.79769313486231570814527423731704357e+308L) && val <= ((double)1.79769313486231570814527423731704357e+308L)) || ImNan(val); }
static inline double ImConstrainNan(double val) { return ImNan(val) ? 0 : val; }
static inline double ImConstrainInf(double val) { return val >= ((double)1.79769313486231570814527423731704357e+308L) ? ((double)1.79769313486231570814527423731704357e+308L) : val <= -((double)1.79769313486231570814527423731704357e+308L) ? - ((double)1.79769313486231570814527423731704357e+308L) : val; }
static inline double ImConstrainLog(double val) { return val <= 0 ? 0.001f : val; }
static inline double ImConstrainTime(double val) { return val < 0 ? 0 : (val > 32503680000 ? 32503680000 : val); }
static inline bool ImAlmostEqual(double v1, double v2, int ulp = 2) { return ImAbs(v1-v2) < ((double)2.22044604925031308084726333618164062e-16L) * ImAbs(v1+v2) * ulp || ImAbs(v1-v2) < ((double)2.22507385850720138309023271733240406e-308L); }
template <typename T>
static inline T ImMinArray(const T* values, int count) { T m = values[0]; for (int i = 1; i < count; ++i) { if (values[i] < m) { m = values[i]; } } return m; }
template <typename T>
static inline T ImMaxArray(const T* values, int count) { T m = values[0]; for (int i = 1; i < count; ++i) { if (values[i] > m) { m = values[i]; } } return m; }
template <typename T>
static inline void ImMinMaxArray(const T* values, int count, T* min_out, T* max_out) {
    T Min = values[0]; T Max = values[0];
    for (int i = 1; i < count; ++i) {
        if (values[i] < Min) { Min = values[i]; }
        if (values[i] > Max) { Max = values[i]; }
    }
    *min_out = Min; *max_out = Max;
}
template <typename T>
static inline T ImSum(const T* values, int count) {
    T sum = 0;
    for (int i = 0; i < count; ++i)
        sum += values[i];
    return sum;
}
template <typename T>
static inline double ImMean(const T* values, int count) {
    double den = 1.0 / count;
    double mu = 0;
    for (int i = 0; i < count; ++i)
        mu += (double)values[i] * den;
    return mu;
}
template <typename T>
static inline double ImStdDev(const T* values, int count) {
    double den = 1.0 / (count - 1.0);
    double mu = ImMean(values, count);
    double x = 0;
    for (int i = 0; i < count; ++i)
        x += ((double)values[i] - mu) * ((double)values[i] - mu) * den;
    return sqrt(x);
}
static inline ImU32 ImMixU32(ImU32 a, ImU32 b, ImU32 s) {
    const ImU32 af = 256-s;
    const ImU32 bf = s;
    const ImU32 al = (a & 0x00ff00ff);
    const ImU32 ah = (a & 0xff00ff00) >> 8;
    const ImU32 bl = (b & 0x00ff00ff);
    const ImU32 bh = (b & 0xff00ff00) >> 8;
    const ImU32 ml = (al * af + bl * bf);
    const ImU32 mh = (ah * af + bh * bf);
    return (mh & 0xff00ff00) | ((ml & 0xff00ff00) >> 8);
}
static inline ImU32 ImLerpU32(const ImU32* colors, int size, float t) {
    int i1 = (int)((size - 1 ) * t);
    int i2 = i1 + 1;
    if (i2 == size || size == 1)
        return colors[i1];
    float den = 1.0f / (size - 1);
    float t1 = i1 * den;
    float t2 = i2 * den;
    float tr = ImRemap01(t, t1, t2);
    return ImMixU32(colors[i1], colors[i2], (ImU32)(tr*256));
}
static inline ImU32 ImAlphaU32(ImU32 col, float alpha) {
    return col & ~((ImU32)((1.0f-alpha)*255)<<24);
}
template <typename T>
static inline bool ImOverlaps(T min_a, T max_a, T min_b, T max_b) {
    return min_a <= max_b && min_b <= max_a;
}
typedef int ImPlotTimeUnit;
typedef int ImPlotDateFmt;
typedef int ImPlotTimeFmt;
enum ImPlotTimeUnit_ {
    ImPlotTimeUnit_Us,
    ImPlotTimeUnit_Ms,
    ImPlotTimeUnit_S,
    ImPlotTimeUnit_Min,
    ImPlotTimeUnit_Hr,
    ImPlotTimeUnit_Day,
    ImPlotTimeUnit_Mo,
    ImPlotTimeUnit_Yr,
    ImPlotTimeUnit_COUNT
};
enum ImPlotDateFmt_ {
    ImPlotDateFmt_None = 0,
    ImPlotDateFmt_DayMo,
    ImPlotDateFmt_DayMoYr,
    ImPlotDateFmt_MoYr,
    ImPlotDateFmt_Mo,
    ImPlotDateFmt_Yr
};
enum ImPlotTimeFmt_ {
    ImPlotTimeFmt_None = 0,
    ImPlotTimeFmt_Us,
    ImPlotTimeFmt_SUs,
    ImPlotTimeFmt_SMs,
    ImPlotTimeFmt_S,
    ImPlotTimeFmt_MinSMs,
    ImPlotTimeFmt_HrMinSMs,
    ImPlotTimeFmt_HrMinS,
    ImPlotTimeFmt_HrMin,
    ImPlotTimeFmt_Hr
};
typedef void (*ImPlotLocator)(ImPlotTicker& ticker, const ImPlotRange& range, float pixels, bool vertical, ImPlotFormatter formatter, void* formatter_data);
struct ImPlotDateTimeSpec {
    ImPlotDateTimeSpec() {}
    ImPlotDateTimeSpec(ImPlotDateFmt date_fmt, ImPlotTimeFmt time_fmt, bool use_24_hr_clk = false, bool use_iso_8601 = false) {
        Date = date_fmt;
        Time = time_fmt;
        UseISO8601 = use_iso_8601;
        Use24HourClock = use_24_hr_clk;
    }
    ImPlotDateFmt Date;
    ImPlotTimeFmt Time;
    bool UseISO8601;
    bool Use24HourClock;
};
struct ImPlotTime {
    time_t S;
    int Us;
    ImPlotTime() { S = 0; Us = 0; }
    ImPlotTime(time_t s, int us = 0) { S = s + us / 1000000; Us = us % 1000000; }
    void RollOver() { S = S + Us / 1000000; Us = Us % 1000000; }
    double ToDouble() const { return (double)S + (double)Us / 1000000.0; }
    static ImPlotTime FromDouble(double t) { return ImPlotTime((time_t)t, (int)(t * 1000000 - floor(t) * 1000000)); }
};
static inline ImPlotTime operator+(const ImPlotTime& lhs, const ImPlotTime& rhs)
{ return ImPlotTime(lhs.S + rhs.S, lhs.Us + rhs.Us); }
static inline ImPlotTime operator-(const ImPlotTime& lhs, const ImPlotTime& rhs)
{ return ImPlotTime(lhs.S - rhs.S, lhs.Us - rhs.Us); }
static inline bool operator==(const ImPlotTime& lhs, const ImPlotTime& rhs)
{ return lhs.S == rhs.S && lhs.Us == rhs.Us; }
static inline bool operator<(const ImPlotTime& lhs, const ImPlotTime& rhs)
{ return lhs.S == rhs.S ? lhs.Us < rhs.Us : lhs.S < rhs.S; }
static inline bool operator>(const ImPlotTime& lhs, const ImPlotTime& rhs)
{ return rhs < lhs; }
static inline bool operator<=(const ImPlotTime& lhs, const ImPlotTime& rhs)
{ return lhs < rhs || lhs == rhs; }
static inline bool operator>=(const ImPlotTime& lhs, const ImPlotTime& rhs)
{ return lhs > rhs || lhs == rhs; }
struct ImPlotColormapData {
    ImVector<ImU32> Keys;
    ImVector<int> KeyCounts;
    ImVector<int> KeyOffsets;
    ImVector<ImU32> Tables;
    ImVector<int> TableSizes;
    ImVector<int> TableOffsets;
    ImGuiTextBuffer Text;
    ImVector<int> TextOffsets;
    ImVector<bool> Quals;
    ImGuiStorage Map;
    int Count;
    ImPlotColormapData() { Count = 0; }
    int Append(const char* name, const ImU32* keys, int count, bool qual) {
        if (GetIndex(name) != -1)
            return -1;
        KeyOffsets.push_back(Keys.size());
        KeyCounts.push_back(count);
        Keys.reserve(Keys.size()+count);
        for (int i = 0; i < count; ++i)
            Keys.push_back(keys[i]);
        TextOffsets.push_back(Text.size());
        Text.append(name, name + strlen(name) + 1);
        Quals.push_back(qual);
        ImGuiID id = ImHashStr(name);
        int idx = Count++;
        Map.SetInt(id,idx);
        _AppendTable(idx);
        return idx;
    }
    void _AppendTable(ImPlotColormap cmap) {
        int key_count = GetKeyCount(cmap);
        const ImU32* keys = GetKeys(cmap);
        int off = Tables.size();
        TableOffsets.push_back(off);
        if (IsQual(cmap)) {
            Tables.reserve(key_count);
            for (int i = 0; i < key_count; ++i)
                Tables.push_back(keys[i]);
            TableSizes.push_back(key_count);
        }
        else {
            int max_size = 255 * (key_count-1) + 1;
            Tables.reserve(off + max_size);
            for (int i = 0; i < key_count-1; ++i) {
                for (int s = 0; s < 255; ++s) {
                    ImU32 a = keys[i];
                    ImU32 b = keys[i+1];
                    ImU32 c = ImMixU32(a,b,s);
                        Tables.push_back(c);
                }
            }
            ImU32 c = keys[key_count-1];
                Tables.push_back(c);
            TableSizes.push_back(max_size);
        }
    }
    void RebuildTables() {
        Tables.resize(0);
        TableSizes.resize(0);
        TableOffsets.resize(0);
        for (int i = 0; i < Count; ++i)
            _AppendTable(i);
    }
    inline bool IsQual(ImPlotColormap cmap) const { return Quals[cmap]; }
    inline const char* GetName(ImPlotColormap cmap) const { return cmap < Count ? Text.Buf.Data + TextOffsets[cmap] :                                                                                                                                             ((void *)0)                                                                                                                                                ; }
    inline ImPlotColormap GetIndex(const char* name) const { ImGuiID key = ImHashStr(name); return Map.GetInt(key,-1); }
    inline const ImU32* GetKeys(ImPlotColormap cmap) const { return &Keys[KeyOffsets[cmap]]; }
    inline int GetKeyCount(ImPlotColormap cmap) const { return KeyCounts[cmap]; }
    inline ImU32 GetKeyColor(ImPlotColormap cmap, int idx) const { return Keys[KeyOffsets[cmap]+idx]; }
    inline void SetKeyColor(ImPlotColormap cmap, int idx, ImU32 value) { Keys[KeyOffsets[cmap]+idx] = value; RebuildTables(); }
    inline const ImU32* GetTable(ImPlotColormap cmap) const { return &Tables[TableOffsets[cmap]]; }
    inline int GetTableSize(ImPlotColormap cmap) const { return TableSizes[cmap]; }
    inline ImU32 GetTableColor(ImPlotColormap cmap, int idx) const { return Tables[TableOffsets[cmap]+idx]; }
    inline ImU32 LerpTable(ImPlotColormap cmap, float t) const {
        int off = TableOffsets[cmap];
        int siz = TableSizes[cmap];
        int idx = Quals[cmap] ? ImClamp((int)(siz*t),0,siz-1) : (int)((siz - 1) * t + 0.5f);
        return Tables[off + idx];
    }
};
struct ImPlotPointError {
    double X, Y, Neg, Pos;
    ImPlotPointError(double x, double y, double neg, double pos) {
        X = x; Y = y; Neg = neg; Pos = pos;
    }
};
struct ImPlotAnnotation {
    ImVec2 Pos;
    ImVec2 Offset;
    ImU32 ColorBg;
    ImU32 ColorFg;
    int TextOffset;
    bool Clamp;
    ImPlotAnnotation() {
        ColorBg = ColorFg = 0;
        TextOffset = 0;
        Clamp = false;
    }
};
struct ImPlotAnnotationCollection {
    ImVector<ImPlotAnnotation> Annotations;
    ImGuiTextBuffer TextBuffer;
    int Size;
    ImPlotAnnotationCollection() { Reset(); }
    void AppendV(const ImVec2& pos, const ImVec2& off, ImU32 bg, ImU32 fg, bool clamp, const char* fmt, va_list args) __attribute__((format(printf, 7, 0))) {
        ImPlotAnnotation an;
        an.Pos = pos; an.Offset = off;
        an.ColorBg = bg; an.ColorFg = fg;
        an.TextOffset = TextBuffer.size();
        an.Clamp = clamp;
        Annotations.push_back(an);
        TextBuffer.appendfv(fmt, args);
        const char nul[] = "";
        TextBuffer.append(nul,nul+1);
        Size++;
    }
    void Append(const ImVec2& pos, const ImVec2& off, ImU32 bg, ImU32 fg, bool clamp, const char* fmt, ...) __attribute__((format(printf, 7, 7 +1))) {
        va_list args;
       __builtin_va_start(       args       ,       fmt       )                          ;
        AppendV(pos, off, bg, fg, clamp, fmt, args);
       __builtin_va_end(       args       )                   ;
    }
    const char* GetText(int idx) {
        return TextBuffer.Buf.Data + Annotations[idx].TextOffset;
    }
    void Reset() {
        Annotations.shrink(0);
        TextBuffer.Buf.shrink(0);
        Size = 0;
    }
};
struct ImPlotTag {
    ImAxis Axis;
    double Value;
    ImU32 ColorBg;
    ImU32 ColorFg;
    int TextOffset;
};
struct ImPlotTagCollection {
    ImVector<ImPlotTag> Tags;
    ImGuiTextBuffer TextBuffer;
    int Size;
    ImPlotTagCollection() { Reset(); }
    void AppendV(ImAxis axis, double value, ImU32 bg, ImU32 fg, const char* fmt, va_list args) __attribute__((format(printf, 6, 0))) {
        ImPlotTag tag;
        tag.Axis = axis;
        tag.Value = value;
        tag.ColorBg = bg;
        tag.ColorFg = fg;
        tag.TextOffset = TextBuffer.size();
        Tags.push_back(tag);
        TextBuffer.appendfv(fmt, args);
        const char nul[] = "";
        TextBuffer.append(nul,nul+1);
        Size++;
    }
    void Append(ImAxis axis, double value, ImU32 bg, ImU32 fg, const char* fmt, ...) __attribute__((format(printf, 6, 6 +1))) {
        va_list args;
       __builtin_va_start(       args       ,       fmt       )                          ;
        AppendV(axis, value, bg, fg, fmt, args);
       __builtin_va_end(       args       )                   ;
    }
    const char* GetText(int idx) {
        return TextBuffer.Buf.Data + Tags[idx].TextOffset;
    }
    void Reset() {
        Tags.shrink(0);
        TextBuffer.Buf.shrink(0);
        Size = 0;
    }
};
struct ImPlotTick
{
    double PlotPos;
    float PixelPos;
    ImVec2 LabelSize;
    int TextOffset;
    bool Major;
    bool ShowLabel;
    int Level;
    int Idx;
    ImPlotTick(double value, bool major, int level, bool show_label) {
        PixelPos = 0;
        PlotPos = value;
        Major = major;
        ShowLabel = show_label;
        Level = level;
        TextOffset = -1;
    }
};
struct ImPlotTicker {
    ImVector<ImPlotTick> Ticks;
    ImGuiTextBuffer TextBuffer;
    ImVec2 MaxSize;
    ImVec2 LateSize;
    int Levels;
    ImPlotTicker() {
        Reset();
    }
    ImPlotTick& AddTick(double value, bool major, int level, bool show_label, const char* label) {
        ImPlotTick tick(value, major, level, show_label);
        if (show_label && label !=                                   ((void *)0)                                      ) {
            tick.TextOffset = TextBuffer.size();
            TextBuffer.append(label, label + strlen(label) + 1);
            tick.LabelSize = ImGui::CalcTextSize(TextBuffer.Buf.Data + tick.TextOffset);
        }
        return AddTick(tick);
    }
    ImPlotTick& AddTick(double value, bool major, int level, bool show_label, ImPlotFormatter formatter, void* data) {
        ImPlotTick tick(value, major, level, show_label);
        if (show_label && formatter !=                                       ((void *)0)                                          ) {
            char buff[32];
            tick.TextOffset = TextBuffer.size();
            formatter(tick.PlotPos, buff, sizeof(buff), data);
            TextBuffer.append(buff, buff + strlen(buff) + 1);
            tick.LabelSize = ImGui::CalcTextSize(TextBuffer.Buf.Data + tick.TextOffset);
        }
        return AddTick(tick);
    }
    inline ImPlotTick& AddTick(ImPlotTick tick) {
        if (tick.ShowLabel) {
            MaxSize.x = tick.LabelSize.x > MaxSize.x ? tick.LabelSize.x : MaxSize.x;
            MaxSize.y = tick.LabelSize.y > MaxSize.y ? tick.LabelSize.y : MaxSize.y;
        }
        tick.Idx = Ticks.size();
        Ticks.push_back(tick);
        return Ticks.back();
    }
    const char* GetText(int idx) const {
        return TextBuffer.Buf.Data + Ticks[idx].TextOffset;
    }
    const char* GetText(const ImPlotTick& tick) {
        return GetText(tick.Idx);
    }
    void OverrideSizeLate(const ImVec2& size) {
        LateSize.x = size.x > LateSize.x ? size.x : LateSize.x;
        LateSize.y = size.y > LateSize.y ? size.y : LateSize.y;
    }
    void Reset() {
        Ticks.shrink(0);
        TextBuffer.Buf.shrink(0);
        MaxSize = LateSize;
        LateSize = ImVec2(0,0);
        Levels = 1;
    }
    int TickCount() const {
        return Ticks.Size;
    }
};
struct ImPlotAxis
{
    ImGuiID ID;
    ImPlotAxisFlags Flags;
    ImPlotAxisFlags PreviousFlags;
    ImPlotRange Range;
    ImPlotCond RangeCond;
    ImPlotScale Scale;
    ImPlotRange FitExtents;
    ImPlotAxis* OrthoAxis;
    ImPlotRange ConstraintRange;
    ImPlotRange ConstraintZoom;
    ImPlotTicker Ticker;
    ImPlotFormatter Formatter;
    void* FormatterData;
    char FormatSpec[16];
    ImPlotLocator Locator;
    double* LinkedMin;
    double* LinkedMax;
    int PickerLevel;
    ImPlotTime PickerTimeMin, PickerTimeMax;
    ImPlotTransform TransformForward;
    ImPlotTransform TransformInverse;
    void* TransformData;
    float PixelMin, PixelMax;
    double ScaleMin, ScaleMax;
    double ScaleToPixel;
    float Datum1, Datum2;
    ImRect HoverRect;
    int LabelOffset;
    ImU32 ColorMaj, ColorMin, ColorTick, ColorTxt, ColorBg, ColorHov, ColorAct, ColorHiLi;
    bool Enabled;
    bool Vertical;
    bool FitThisFrame;
    bool HasRange;
    bool HasFormatSpec;
    bool ShowDefaultTicks;
    bool Hovered;
    bool Held;
    ImPlotAxis() {
        ID = 0;
        Flags = PreviousFlags = ImPlotAxisFlags_None;
        Range.Min = 0;
        Range.Max = 1;
        Scale = ImPlotScale_Linear;
        TransformForward = TransformInverse =                                              ((void *)0)                                                 ;
        TransformData =                           ((void *)0)                              ;
        FitExtents.Min =                           (__builtin_huge_val ())                                  ;
        FitExtents.Max = -                           (__builtin_huge_val ())                                   ;
        OrthoAxis =                           ((void *)0)                              ;
        ConstraintRange = ImPlotRange(-                                       (__builtin_inff ())                                               ,                                                (__builtin_inff ())                                                        );
        ConstraintZoom = ImPlotRange(((double)2.22507385850720138309023271733240406e-308L),                                              (__builtin_inff ())                                                      );
        LinkedMin = LinkedMax =                                       ((void *)0)                                          ;
        PickerLevel = 0;
        Datum1 = Datum2 = 0;
        PixelMin = PixelMax = 0;
        LabelOffset = -1;
        ColorMaj = ColorMin = ColorTick = ColorTxt = ColorBg = ColorHov = ColorAct = 0;
        ColorHiLi = (((ImU32)(0)<<24) | ((ImU32)(0)<<16) | ((ImU32)(0)<<8) | ((ImU32)(0)<<0));
        Formatter =                           ((void *)0)                              ;
        FormatterData =                           ((void *)0)                              ;
        Locator =                           ((void *)0)                              ;
        Enabled = Hovered = Held = FitThisFrame = HasRange = HasFormatSpec = false;
        ShowDefaultTicks = true;
    }
    inline void Reset() {
        Enabled = false;
        Scale = ImPlotScale_Linear;
        TransformForward = TransformInverse =                                              ((void *)0)                                                 ;
        TransformData =                           ((void *)0)                              ;
        LabelOffset = -1;
        HasFormatSpec = false;
        Formatter =                           ((void *)0)                              ;
        FormatterData =                           ((void *)0)                              ;
        Locator =                           ((void *)0)                              ;
        ShowDefaultTicks = true;
        FitThisFrame = false;
        FitExtents.Min =                           (__builtin_huge_val ())                                  ;
        FitExtents.Max = -                           (__builtin_huge_val ())                                   ;
        OrthoAxis =                           ((void *)0)                              ;
        ConstraintRange = ImPlotRange(-                                       (__builtin_inff ())                                               ,                                                (__builtin_inff ())                                                        );
        ConstraintZoom = ImPlotRange(((double)2.22507385850720138309023271733240406e-308L),                                              (__builtin_inff ())                                                      );
        Ticker.Reset();
    }
    inline bool SetMin(double _min, bool force=false) {
        if (!force && IsLockedMin())
            return false;
        _min = ImConstrainNan(ImConstrainInf(_min));
        if (_min < ConstraintRange.Min)
            _min = ConstraintRange.Min;
        double z = Range.Max - _min;
        if (z < ConstraintZoom.Min)
            _min = Range.Max - ConstraintZoom.Min;
        if (z > ConstraintZoom.Max)
            _min = Range.Max - ConstraintZoom.Max;
        if (_min >= Range.Max)
            return false;
        Range.Min = _min;
        PickerTimeMin = ImPlotTime::FromDouble(Range.Min);
        UpdateTransformCache();
        return true;
    };
    inline bool SetMax(double _max, bool force=false) {
        if (!force && IsLockedMax())
            return false;
        _max = ImConstrainNan(ImConstrainInf(_max));
        if (_max > ConstraintRange.Max)
            _max = ConstraintRange.Max;
        double z = _max - Range.Min;
        if (z < ConstraintZoom.Min)
            _max = Range.Min + ConstraintZoom.Min;
        if (z > ConstraintZoom.Max)
            _max = Range.Min + ConstraintZoom.Max;
        if (_max <= Range.Min)
            return false;
        Range.Max = _max;
        PickerTimeMax = ImPlotTime::FromDouble(Range.Max);
        UpdateTransformCache();
        return true;
    };
    inline void SetRange(double v1, double v2) {
        Range.Min = ImMin(v1,v2);
        Range.Max = ImMax(v1,v2);
        Constrain();
        PickerTimeMin = ImPlotTime::FromDouble(Range.Min);
        PickerTimeMax = ImPlotTime::FromDouble(Range.Max);
        UpdateTransformCache();
    }
    inline void SetRange(const ImPlotRange& range) {
        SetRange(range.Min, range.Max);
    }
    inline void SetAspect(double unit_per_pix) {
        double new_size = unit_per_pix * PixelSize();
        double delta = (new_size - Range.Size()) * 0.5;
        if (IsLocked())
            return;
        else if (IsLockedMin() && !IsLockedMax())
            SetRange(Range.Min, Range.Max + 2*delta);
        else if (!IsLockedMin() && IsLockedMax())
            SetRange(Range.Min - 2*delta, Range.Max);
        else
            SetRange(Range.Min - delta, Range.Max + delta);
    }
    inline float PixelSize() const { return ImAbs(PixelMax - PixelMin); }
    inline double GetAspect() const { return Range.Size() / PixelSize(); }
    inline void Constrain() {
        Range.Min = ImConstrainNan(ImConstrainInf(Range.Min));
        Range.Max = ImConstrainNan(ImConstrainInf(Range.Max));
        if (Range.Min < ConstraintRange.Min)
            Range.Min = ConstraintRange.Min;
        if (Range.Max > ConstraintRange.Max)
            Range.Max = ConstraintRange.Max;
        double z = Range.Size();
        if (z < ConstraintZoom.Min) {
            double delta = (ConstraintZoom.Min - z) * 0.5;
            Range.Min -= delta;
            Range.Max += delta;
        }
        if (z > ConstraintZoom.Max) {
            double delta = (z - ConstraintZoom.Max) * 0.5;
            Range.Min += delta;
            Range.Max -= delta;
        }
        if (Range.Max <= Range.Min)
            Range.Max = Range.Min + ((double)2.22044604925031308084726333618164062e-16L);
    }
    inline void UpdateTransformCache() {
        ScaleToPixel = (PixelMax - PixelMin) / Range.Size();
        if (TransformForward !=                                ((void *)0)                                   ) {
            ScaleMin = TransformForward(Range.Min, TransformData);
            ScaleMax = TransformForward(Range.Max, TransformData);
        }
        else {
            ScaleMin = Range.Min;
            ScaleMax = Range.Max;
        }
    }
    inline float PlotToPixels(double plt) const {
        if (TransformForward !=                                ((void *)0)                                   ) {
            double s = TransformForward(plt, TransformData);
            double t = (s - ScaleMin) / (ScaleMax - ScaleMin);
            plt = Range.Min + Range.Size() * t;
        }
        return (float)(PixelMin + ScaleToPixel * (plt - Range.Min));
    }
    inline double PixelsToPlot(float pix) const {
        double plt = (pix - PixelMin) / ScaleToPixel + Range.Min;
        if (TransformInverse !=                                ((void *)0)                                   ) {
            double t = (plt - Range.Min) / Range.Size();
            double s = t * (ScaleMax - ScaleMin) + ScaleMin;
            plt = TransformInverse(s, TransformData);
        }
        return plt;
    }
    inline void ExtendFit(double v) {
        if (!ImNanOrInf(v) && v >= ConstraintRange.Min && v <= ConstraintRange.Max) {
            FitExtents.Min = v < FitExtents.Min ? v : FitExtents.Min;
            FitExtents.Max = v > FitExtents.Max ? v : FitExtents.Max;
        }
    }
    inline void ExtendFitWith(ImPlotAxis& alt, double v, double v_alt) {
        if (ImHasFlag(Flags, ImPlotAxisFlags_RangeFit) && !alt.Range.Contains(v_alt))
            return;
        if (!ImNanOrInf(v) && v >= ConstraintRange.Min && v <= ConstraintRange.Max) {
            FitExtents.Min = v < FitExtents.Min ? v : FitExtents.Min;
            FitExtents.Max = v > FitExtents.Max ? v : FitExtents.Max;
        }
    }
    inline void ApplyFit(float padding) {
        const double ext_size = FitExtents.Size() * 0.5;
        FitExtents.Min -= ext_size * padding;
        FitExtents.Max += ext_size * padding;
        if (!IsLockedMin() && !ImNanOrInf(FitExtents.Min))
            Range.Min = FitExtents.Min;
        if (!IsLockedMax() && !ImNanOrInf(FitExtents.Max))
            Range.Max = FitExtents.Max;
        if (ImAlmostEqual(Range.Min, Range.Max)) {
            Range.Max += 0.5;
            Range.Min -= 0.5;
        }
        Constrain();
        UpdateTransformCache();
    }
    inline bool HasLabel() const { return LabelOffset != -1 && !ImHasFlag(Flags, ImPlotAxisFlags_NoLabel); }
    inline bool HasGridLines() const { return !ImHasFlag(Flags, ImPlotAxisFlags_NoGridLines); }
    inline bool HasTickLabels() const { return !ImHasFlag(Flags, ImPlotAxisFlags_NoTickLabels); }
    inline bool HasTickMarks() const { return !ImHasFlag(Flags, ImPlotAxisFlags_NoTickMarks); }
    inline bool WillRender() const { return Enabled && (HasGridLines() || HasTickLabels() || HasTickMarks()); }
    inline bool IsOpposite() const { return ImHasFlag(Flags, ImPlotAxisFlags_Opposite); }
    inline bool IsInverted() const { return ImHasFlag(Flags, ImPlotAxisFlags_Invert); }
    inline bool IsForeground() const { return ImHasFlag(Flags, ImPlotAxisFlags_Foreground); }
    inline bool IsAutoFitting() const { return ImHasFlag(Flags, ImPlotAxisFlags_AutoFit); }
    inline bool CanInitFit() const { return !ImHasFlag(Flags, ImPlotAxisFlags_NoInitialFit) && !HasRange && !LinkedMin && !LinkedMax; }
    inline bool IsRangeLocked() const { return HasRange && RangeCond == ImPlotCond_Always; }
    inline bool IsLockedMin() const { return !Enabled || IsRangeLocked() || ImHasFlag(Flags, ImPlotAxisFlags_LockMin); }
    inline bool IsLockedMax() const { return !Enabled || IsRangeLocked() || ImHasFlag(Flags, ImPlotAxisFlags_LockMax); }
    inline bool IsLocked() const { return IsLockedMin() && IsLockedMax(); }
    inline bool IsInputLockedMin() const { return IsLockedMin() || IsAutoFitting(); }
    inline bool IsInputLockedMax() const { return IsLockedMax() || IsAutoFitting(); }
    inline bool IsInputLocked() const { return IsLocked() || IsAutoFitting(); }
    inline bool HasMenus() const { return !ImHasFlag(Flags, ImPlotAxisFlags_NoMenus); }
    inline bool IsPanLocked(bool increasing) {
        if (ImHasFlag(Flags, ImPlotAxisFlags_PanStretch)) {
            return IsInputLocked();
        }
        else {
            if (IsLockedMin() || IsLockedMax() || IsAutoFitting())
                return false;
            if (increasing)
                return Range.Max == ConstraintRange.Max;
            else
                return Range.Min == ConstraintRange.Min;
        }
    }
    void PushLinks() {
        if (LinkedMin) { *LinkedMin = Range.Min; }
        if (LinkedMax) { *LinkedMax = Range.Max; }
    }
    void PullLinks() {
        if (LinkedMin) { SetMin(*LinkedMin,true); }
        if (LinkedMax) { SetMax(*LinkedMax,true); }
    }
};
struct ImPlotAlignmentData {
    bool Vertical;
    float PadA;
    float PadB;
    float PadAMax;
    float PadBMax;
    ImPlotAlignmentData() {
        Vertical = true;
        PadA = PadB = PadAMax = PadBMax = 0;
    }
    void Begin() { PadAMax = PadBMax = 0; }
    void Update(float& pad_a, float& pad_b, float& delta_a, float& delta_b) {
        float bak_a = pad_a; float bak_b = pad_b;
        if (PadAMax < pad_a) { PadAMax = pad_a; }
        if (PadBMax < pad_b) { PadBMax = pad_b; }
        if (pad_a < PadA) { pad_a = PadA; delta_a = pad_a - bak_a; } else { delta_a = 0; }
        if (pad_b < PadB) { pad_b = PadB; delta_b = pad_b - bak_b; } else { delta_b = 0; }
    }
    void End() { PadA = PadAMax; PadB = PadBMax; }
    void Reset() { PadA = PadB = PadAMax = PadBMax = 0; }
};
struct ImPlotItem
{
    ImGuiID ID;
    ImU32 Color;
    ImRect LegendHoverRect;
    int NameOffset;
    bool Show;
    bool LegendHovered;
    bool SeenThisFrame;
    ImPlotItem() {
        ID = 0;
        Color = (((ImU32)(255)<<24) | ((ImU32)(255)<<16) | ((ImU32)(255)<<8) | ((ImU32)(255)<<0));
        NameOffset = -1;
        Show = true;
        SeenThisFrame = false;
        LegendHovered = false;
    }
    ~ImPlotItem() { ID = 0; }
};
struct ImPlotLegend
{
    ImPlotLegendFlags Flags;
    ImPlotLegendFlags PreviousFlags;
    ImPlotLocation Location;
    ImPlotLocation PreviousLocation;
    ImVector<int> Indices;
    ImGuiTextBuffer Labels;
    ImRect Rect;
    bool Hovered;
    bool Held;
    bool CanGoInside;
    ImPlotLegend() {
        Flags = PreviousFlags = ImPlotLegendFlags_None;
        CanGoInside = true;
        Hovered = Held = false;
        Location = PreviousLocation = ImPlotLocation_NorthWest;
    }
    void Reset() { Indices.shrink(0); Labels.Buf.shrink(0); }
};
struct ImPlotItemGroup
{
    ImGuiID ID;
    ImPlotLegend Legend;
    ImPool<ImPlotItem> ItemPool;
    int ColormapIdx;
    ImPlotItemGroup() { ID = 0; ColormapIdx = 0; }
    int GetItemCount() const { return ItemPool.GetBufSize(); }
    ImGuiID GetItemID(const char* label_id) { return ImGui::GetID(label_id); }
    ImPlotItem* GetItem(ImGuiID id) { return ItemPool.GetByKey(id); }
    ImPlotItem* GetItem(const char* label_id) { return GetItem(GetItemID(label_id)); }
    ImPlotItem* GetOrAddItem(ImGuiID id) { return ItemPool.GetOrAddByKey(id); }
    ImPlotItem* GetItemByIndex(int i) { return ItemPool.GetByIndex(i); }
    int GetItemIndex(ImPlotItem* item) { return ItemPool.GetIndex(item); }
    int GetLegendCount() const { return Legend.Indices.size(); }
    ImPlotItem* GetLegendItem(int i) { return ItemPool.GetByIndex(Legend.Indices[i]); }
    const char* GetLegendLabel(int i) { return Legend.Labels.Buf.Data + GetLegendItem(i)->NameOffset; }
    void Reset() { ItemPool.Clear(); Legend.Reset(); ColormapIdx = 0; }
};
struct ImPlotPlot
{
    ImGuiID ID;
    ImPlotFlags Flags;
    ImPlotFlags PreviousFlags;
    ImPlotLocation MouseTextLocation;
    ImPlotMouseTextFlags MouseTextFlags;
    ImPlotAxis Axes[ImAxis_COUNT];
    ImGuiTextBuffer TextBuffer;
    ImPlotItemGroup Items;
    ImAxis CurrentX;
    ImAxis CurrentY;
    ImRect FrameRect;
    ImRect CanvasRect;
    ImRect PlotRect;
    ImRect AxesRect;
    ImRect SelectRect;
    ImVec2 SelectStart;
    int TitleOffset;
    bool JustCreated;
    bool Initialized;
    bool SetupLocked;
    bool FitThisFrame;
    bool Hovered;
    bool Held;
    bool Selecting;
    bool Selected;
    bool ContextLocked;
    ImPlotPlot() {
        Flags = PreviousFlags = ImPlotFlags_None;
        for (int i = 0; i < ImAxis_Y1; ++i)
            XAxis(i).Vertical = false;
        for (int i = 0; i < (ImAxis_COUNT - ImAxis_Y1); ++i)
            YAxis(i).Vertical = true;
        SelectStart = ImVec2(0,0);
        CurrentX = ImAxis_X1;
        CurrentY = ImAxis_Y1;
        MouseTextLocation = ImPlotLocation_South | ImPlotLocation_East;
        MouseTextFlags = ImPlotMouseTextFlags_None;
        TitleOffset = -1;
        JustCreated = true;
        Initialized = SetupLocked = FitThisFrame = false;
        Hovered = Held = Selected = Selecting = ContextLocked = false;
    }
    inline bool IsInputLocked() const {
        for (int i = 0; i < ImAxis_Y1; ++i) {
            if (!XAxis(i).IsInputLocked())
                return false;
        }
        for (int i = 0; i < (ImAxis_COUNT - ImAxis_Y1); ++i) {
            if (!YAxis(i).IsInputLocked())
                return false;
        }
        return true;
    }
    inline void ClearTextBuffer() { TextBuffer.Buf.shrink(0); }
    inline void SetTitle(const char* title) {
        if (title && ImGui::FindRenderedTextEnd(title,                                                       ((void *)0)                                                          ) != title) {
            TitleOffset = TextBuffer.size();
            TextBuffer.append(title, title + strlen(title) + 1);
        }
        else {
            TitleOffset = -1;
        }
    }
    inline bool HasTitle() const { return TitleOffset != -1 && !ImHasFlag(Flags, ImPlotFlags_NoTitle); }
    inline const char* GetTitle() const { return TextBuffer.Buf.Data + TitleOffset; }
    inline ImPlotAxis& XAxis(int i) { return Axes[ImAxis_X1 + i]; }
    inline const ImPlotAxis& XAxis(int i) const { return Axes[ImAxis_X1 + i]; }
    inline ImPlotAxis& YAxis(int i) { return Axes[ImAxis_Y1 + i]; }
    inline const ImPlotAxis& YAxis(int i) const { return Axes[ImAxis_Y1 + i]; }
    inline int EnabledAxesX() {
        int cnt = 0;
        for (int i = 0; i < ImAxis_Y1; ++i)
            cnt += XAxis(i).Enabled;
        return cnt;
    }
    inline int EnabledAxesY() {
        int cnt = 0;
        for (int i = 0; i < (ImAxis_COUNT - ImAxis_Y1); ++i)
            cnt += YAxis(i).Enabled;
        return cnt;
    }
    inline void SetAxisLabel(ImPlotAxis& axis, const char* label) {
        if (label && ImGui::FindRenderedTextEnd(label,                                                       ((void *)0)                                                          ) != label) {
            axis.LabelOffset = TextBuffer.size();
            TextBuffer.append(label, label + strlen(label) + 1);
        }
        else {
            axis.LabelOffset = -1;
        }
    }
    inline const char* GetAxisLabel(const ImPlotAxis& axis) const { return TextBuffer.Buf.Data + axis.LabelOffset; }
};
struct ImPlotSubplot {
    ImGuiID ID;
    ImPlotSubplotFlags Flags;
    ImPlotSubplotFlags PreviousFlags;
    ImPlotItemGroup Items;
    int Rows;
    int Cols;
    int CurrentIdx;
    ImRect FrameRect;
    ImRect GridRect;
    ImVec2 CellSize;
    ImVector<ImPlotAlignmentData> RowAlignmentData;
    ImVector<ImPlotAlignmentData> ColAlignmentData;
    ImVector<float> RowRatios;
    ImVector<float> ColRatios;
    ImVector<ImPlotRange> RowLinkData;
    ImVector<ImPlotRange> ColLinkData;
    float TempSizes[2];
    bool FrameHovered;
    bool HasTitle;
    ImPlotSubplot() {
        ID = 0;
        Flags = PreviousFlags = ImPlotSubplotFlags_None;
        Rows = Cols = CurrentIdx = 0;
        FrameHovered = false;
        Items.Legend.Location = ImPlotLocation_North;
        Items.Legend.Flags = ImPlotLegendFlags_Horizontal|ImPlotLegendFlags_Outside;
        Items.Legend.CanGoInside = false;
        TempSizes[0] = TempSizes[1] = 0;
        FrameHovered = false;
        HasTitle = false;
    }
};
struct ImPlotNextPlotData
{
    ImPlotCond RangeCond[ImAxis_COUNT];
    ImPlotRange Range[ImAxis_COUNT];
    bool HasRange[ImAxis_COUNT];
    bool Fit[ImAxis_COUNT];
    double* LinkedMin[ImAxis_COUNT];
    double* LinkedMax[ImAxis_COUNT];
    ImPlotNextPlotData() { Reset(); }
    void Reset() {
        for (int i = 0; i < ImAxis_COUNT; ++i) {
            HasRange[i] = false;
            Fit[i] = false;
            LinkedMin[i] = LinkedMax[i] =                                          ((void *)0)                                             ;
        }
    }
};
struct ImPlotNextItemData {
    ImVec4 Colors[5];
    float LineWeight;
    ImPlotMarker Marker;
    float MarkerSize;
    float MarkerWeight;
    float FillAlpha;
    float ErrorBarSize;
    float ErrorBarWeight;
    float DigitalBitHeight;
    float DigitalBitGap;
    bool RenderLine;
    bool RenderFill;
    bool RenderMarkerLine;
    bool RenderMarkerFill;
    bool HasHidden;
    bool Hidden;
    ImPlotCond HiddenCond;
    ImPlotNextItemData() { Reset(); }
    void Reset() {
        for (int i = 0; i < 5; ++i)
            Colors[i] = ImVec4(0,0,0,-1);
        LineWeight = MarkerSize = MarkerWeight = FillAlpha = ErrorBarSize = ErrorBarWeight = DigitalBitHeight = DigitalBitGap = -1;
        Marker = -1;
        HasHidden = Hidden = false;
    }
};
struct ImPlotContext {
    ImPool<ImPlotPlot> Plots;
    ImPool<ImPlotSubplot> Subplots;
    ImPlotPlot* CurrentPlot;
    ImPlotSubplot* CurrentSubplot;
    ImPlotItemGroup* CurrentItems;
    ImPlotItem* CurrentItem;
    ImPlotItem* PreviousItem;
    ImPlotTicker CTicker;
    ImPlotAnnotationCollection Annotations;
    ImPlotTagCollection Tags;
    bool ChildWindowMade;
    ImPlotStyle Style;
    ImVector<ImGuiColorMod> ColorModifiers;
    ImVector<ImGuiStyleMod> StyleModifiers;
    ImPlotColormapData ColormapData;
    ImVector<ImPlotColormap> ColormapModifiers;
    tm Tm;
    ImVector<double> TempDouble1, TempDouble2;
    ImVector<int> TempInt1;
    int DigitalPlotItemCnt;
    int DigitalPlotOffset;
    ImPlotNextPlotData NextPlotData;
    ImPlotNextItemData NextItemData;
    ImPlotInputMap InputMap;
    bool OpenContextThisFrame;
    ImGuiTextBuffer MousePosStringBuilder;
    ImPlotItemGroup* SortItems;
    ImPool<ImPlotAlignmentData> AlignmentData;
    ImPlotAlignmentData* CurrentAlignmentH;
    ImPlotAlignmentData* CurrentAlignmentV;
};
namespace ImPlot {
 void Initialize(ImPlotContext* ctx);
 void ResetCtxForNextPlot(ImPlotContext* ctx);
 void ResetCtxForNextAlignedPlots(ImPlotContext* ctx);
 void ResetCtxForNextSubplot(ImPlotContext* ctx);
 ImPlotPlot* GetPlot(const char* title);
 ImPlotPlot* GetCurrentPlot();
 void BustPlotCache();
 void ShowPlotContextMenu(ImPlotPlot& plot);
static inline void SetupLock() {
    if (!GImPlot->CurrentPlot->SetupLocked)
        SetupFinish();
    GImPlot->CurrentPlot->SetupLocked = true;
}
 void SubplotNextCell();
 void ShowSubplotsContextMenu(ImPlotSubplot& subplot);
 bool BeginItem(const char* label_id, ImPlotItemFlags flags=0, ImPlotCol recolor_from=-1);
template <typename _Fitter>
bool BeginItemEx(const char* label_id, const _Fitter& fitter, ImPlotItemFlags flags=0, ImPlotCol recolor_from=-1) {
    if (BeginItem(label_id, flags, recolor_from)) {
        ImPlotPlot& plot = *GetCurrentPlot();
        if (plot.FitThisFrame && !ImHasFlag(flags, ImPlotItemFlags_NoFit))
            fitter.Fit(plot.Axes[plot.CurrentX], plot.Axes[plot.CurrentY]);
        return true;
    }
    return false;
}
 void EndItem();
 ImPlotItem* RegisterOrGetItem(const char* label_id, ImPlotItemFlags flags, bool* just_created =                                                                                                           ((void *)0)                                                                                                              );
 ImPlotItem* GetItem(const char* label_id);
 ImPlotItem* GetCurrentItem();
 void BustItemCache();
static inline bool AnyAxesInputLocked(ImPlotAxis* axes, int count) {
    for (int i = 0; i < count; ++i) {
        if (axes[i].Enabled && axes[i].IsInputLocked())
            return true;
    }
    return false;
}
static inline bool AllAxesInputLocked(ImPlotAxis* axes, int count) {
    for (int i = 0; i < count; ++i) {
        if (axes[i].Enabled && !axes[i].IsInputLocked())
            return false;
    }
    return true;
}
static inline bool AnyAxesHeld(ImPlotAxis* axes, int count) {
    for (int i = 0; i < count; ++i) {
        if (axes[i].Enabled && axes[i].Held)
            return true;
    }
    return false;
}
static inline bool AnyAxesHovered(ImPlotAxis* axes, int count) {
    for (int i = 0; i < count; ++i) {
        if (axes[i].Enabled && axes[i].Hovered)
            return true;
    }
    return false;
}
static inline bool FitThisFrame() {
    return GImPlot->CurrentPlot->FitThisFrame;
}
static inline void FitPointX(double x) {
    ImPlotPlot& plot = *GetCurrentPlot();
    ImPlotAxis& x_axis = plot.Axes[plot.CurrentX];
    x_axis.ExtendFit(x);
}
static inline void FitPointY(double y) {
    ImPlotPlot& plot = *GetCurrentPlot();
    ImPlotAxis& y_axis = plot.Axes[plot.CurrentY];
    y_axis.ExtendFit(y);
}
static inline void FitPoint(const ImPlotPoint& p) {
    ImPlotPlot& plot = *GetCurrentPlot();
    ImPlotAxis& x_axis = plot.Axes[plot.CurrentX];
    ImPlotAxis& y_axis = plot.Axes[plot.CurrentY];
    x_axis.ExtendFitWith(y_axis, p.x, p.y);
    y_axis.ExtendFitWith(x_axis, p.y, p.x);
}
static inline bool RangesOverlap(const ImPlotRange& r1, const ImPlotRange& r2)
{ return r1.Min <= r2.Max && r2.Min <= r1.Max; }
 void ShowAxisContextMenu(ImPlotAxis& axis, ImPlotAxis* equal_axis, bool time_allowed = false);
 ImVec2 GetLocationPos(const ImRect& outer_rect, const ImVec2& inner_size, ImPlotLocation location, const ImVec2& pad = ImVec2(0,0));
 ImVec2 CalcLegendSize(ImPlotItemGroup& items, const ImVec2& pad, const ImVec2& spacing, bool vertical);
 bool ShowLegendEntries(ImPlotItemGroup& items, const ImRect& legend_bb, bool interactable, const ImVec2& pad, const ImVec2& spacing, bool vertical, ImDrawList& DrawList);
 void ShowAltLegend(const char* title_id, bool vertical = true, const ImVec2 size = ImVec2(0,0), bool interactable = true);
 bool ShowLegendContextMenu(ImPlotLegend& legend, bool visible);
 void LabelAxisValue(const ImPlotAxis& axis, double value, char* buff, int size, bool round = false);
static inline const ImPlotNextItemData& GetItemData() { return GImPlot->NextItemData; }
static inline bool IsColorAuto(const ImVec4& col) { return col.w == -1; }
static inline bool IsColorAuto(ImPlotCol idx) { return IsColorAuto(GImPlot->Style.Colors[idx]); }
 ImVec4 GetAutoColor(ImPlotCol idx);
static inline ImVec4 GetStyleColorVec4(ImPlotCol idx) { return IsColorAuto(idx) ? GetAutoColor(idx) : GImPlot->Style.Colors[idx]; }
static inline ImU32 GetStyleColorU32(ImPlotCol idx) { return ImGui::ColorConvertFloat4ToU32(GetStyleColorVec4(idx)); }
 void AddTextVertical(ImDrawList *DrawList, ImVec2 pos, ImU32 col, const char* text_begin, const char* text_end =                                                                                                                            ((void *)0)                                                                                                                               );
 void AddTextCentered(ImDrawList* DrawList, ImVec2 top_center, ImU32 col, const char* text_begin, const char* text_end =                                                                                                                                   ((void *)0)                                                                                                                                      );
static inline ImVec2 CalcTextSizeVertical(const char *text) {
    ImVec2 sz = ImGui::CalcTextSize(text);
    return ImVec2(sz.y, sz.x);
}
static inline ImU32 CalcTextColor(const ImVec4& bg) { return (bg.x * 0.299f + bg.y * 0.587f + bg.z * 0.114f) > 0.5f ? (((ImU32)(255)<<24) | ((ImU32)(0)<<16) | ((ImU32)(0)<<8) | ((ImU32)(0)<<0)) : (((ImU32)(255)<<24) | ((ImU32)(255)<<16) | ((ImU32)(255)<<8) | ((ImU32)(255)<<0)); }
static inline ImU32 CalcTextColor(ImU32 bg) { return CalcTextColor(ImGui::ColorConvertU32ToFloat4(bg)); }
static inline ImU32 CalcHoverColor(ImU32 col) { return ImMixU32(col, CalcTextColor(col), 32); }
static inline ImVec2 ClampLabelPos(ImVec2 pos, const ImVec2& size, const ImVec2& Min, const ImVec2& Max) {
    if (pos.x < Min.x) pos.x = Min.x;
    if (pos.y < Min.y) pos.y = Min.y;
    if ((pos.x + size.x) > Max.x) pos.x = Max.x - size.x;
    if ((pos.y + size.y) > Max.y) pos.y = Max.y - size.y;
    return pos;
}
 ImU32 GetColormapColorU32(int idx, ImPlotColormap cmap);
 ImU32 NextColormapColorU32();
 ImU32 SampleColormapU32(float t, ImPlotColormap cmap);
 void RenderColorBar(const ImU32* colors, int size, ImDrawList& DrawList, const ImRect& bounds, bool vert, bool reversed, bool continuous);
 double NiceNum(double x, bool round);
static inline int OrderOfMagnitude(double val) { return val == 0 ? 0 : (int)(floor(log10(fabs(val)))); }
static inline int OrderToPrecision(int order) { return order > 0 ? 0 : 1 - order; }
static inline int Precision(double val) { return OrderToPrecision(OrderOfMagnitude(val)); }
static inline double RoundTo(double val, int prec) { double p = pow(10,(double)prec); return floor(val*p+0.5)/p; }
static inline ImVec2 Intersection(const ImVec2& a1, const ImVec2& a2, const ImVec2& b1, const ImVec2& b2) {
    float v1 = (a1.x * a2.y - a1.y * a2.x); float v2 = (b1.x * b2.y - b1.y * b2.x);
    float v3 = ((a1.x - a2.x) * (b1.y - b2.y) - (a1.y - a2.y) * (b1.x - b2.x));
    return ImVec2((v1 * (b1.x - b2.x) - v2 * (a1.x - a2.x)) / v3, (v1 * (b1.y - b2.y) - v2 * (a1.y - a2.y)) / v3);
}
template <typename T>
void FillRange(ImVector<T>& buffer, int n, T vmin, T vmax) {
    buffer.resize(n);
    T step = (vmax - vmin) / (n - 1);
    for (int i = 0; i < n; ++i) {
        buffer[i] = vmin + i * step;
    }
}
template <typename T>
static inline void CalculateBins(const T* values, int count, ImPlotBin meth, const ImPlotRange& range, int& bins_out, double& width_out) {
    switch (meth) {
        case ImPlotBin_Sqrt:
            bins_out = (int)ceil(sqrt(count));
            break;
        case ImPlotBin_Sturges:
            bins_out = (int)ceil(1.0 + log2(count));
            break;
        case ImPlotBin_Rice:
            bins_out = (int)ceil(2 * cbrt(count));
            break;
        case ImPlotBin_Scott:
            width_out = 3.49 * ImStdDev(values, count) / cbrt(count);
            bins_out = (int)round(range.Size() / width_out);
            break;
    }
    width_out = range.Size() / bins_out;
}
static inline bool IsLeapYear(int year) {
    return year % 4 == 0 && (year % 100 != 0 || year % 400 == 0);
}
static inline int GetDaysInMonth(int year, int month) {
    static const int days[12] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
    return days[month] + (int)(month == 1 && IsLeapYear(year));
}
 ImPlotTime MkGmtTime(struct tm *ptm);
 tm* GetGmtTime(const ImPlotTime& t, tm* ptm);
 ImPlotTime MkLocTime(struct tm *ptm);
 tm* GetLocTime(const ImPlotTime& t, tm* ptm);
 ImPlotTime MakeTime(int year, int month = 0, int day = 1, int hour = 0, int min = 0, int sec = 0, int us = 0);
 int GetYear(const ImPlotTime& t);
 ImPlotTime AddTime(const ImPlotTime& t, ImPlotTimeUnit unit, int count);
 ImPlotTime FloorTime(const ImPlotTime& t, ImPlotTimeUnit unit);
 ImPlotTime CeilTime(const ImPlotTime& t, ImPlotTimeUnit unit);
 ImPlotTime RoundTime(const ImPlotTime& t, ImPlotTimeUnit unit);
 ImPlotTime CombineDateTime(const ImPlotTime& date_part, const ImPlotTime& time_part);
 int FormatTime(const ImPlotTime& t, char* buffer, int size, ImPlotTimeFmt fmt, bool use_24_hr_clk);
 int FormatDate(const ImPlotTime& t, char* buffer, int size, ImPlotDateFmt fmt, bool use_iso_8601);
 int FormatDateTime(const ImPlotTime& t, char* buffer, int size, ImPlotDateTimeSpec fmt);
 bool ShowDatePicker(const char* id, int* level, ImPlotTime* t, const ImPlotTime* t1 =                                                                                                 ((void *)0)                                                                                                    , const ImPlotTime* t2 =                                                                                                                              ((void *)0)                                                                                                                                 );
 bool ShowTimePicker(const char* id, ImPlotTime* t);
static inline double TransformForward_Log10(double v, void*) {
    v = v <= 0.0 ? ((double)2.22507385850720138309023271733240406e-308L) : v;
    return ImLog10(v);
}
static inline double TransformInverse_Log10(double v, void*) {
    return ImPow(10, v);
}
static inline double TransformForward_SymLog(double v, void*) {
    return 2.0 * ImAsinh(v / 2.0);
}
static inline double TransformInverse_SymLog(double v, void*) {
    return 2.0 * ImSinh(v / 2.0);
}
static inline double TransformForward_Logit(double v, void*) {
    v = ImClamp(v, ((double)2.22507385850720138309023271733240406e-308L), 1.0 - ((double)2.22044604925031308084726333618164062e-16L));
    return ImLog10(v / (1 - v));
}
static inline double TransformInverse_Logit(double v, void*) {
    return 1.0 / (1.0 + ImPow(10,-v));
}
static inline int Formatter_Default(double value, char* buff, int size, void* data) {
    char* fmt = (char*)data;
    return ImFormatString(buff, size, fmt, value);
}
static inline int Formatter_Logit(double value, char* buff, int size, void*) {
    if (value == 0.5)
        return ImFormatString(buff,size,"1/2");
    else if (value < 0.5)
        return ImFormatString(buff,size,"%g", value);
    else
        return ImFormatString(buff,size,"1 - %g", 1 - value);
}
struct Formatter_Time_Data {
    ImPlotTime Time;
    ImPlotDateTimeSpec Spec;
    ImPlotFormatter UserFormatter;
    void* UserFormatterData;
};
static inline int Formatter_Time(double, char* buff, int size, void* data) {
    Formatter_Time_Data* ftd = (Formatter_Time_Data*)data;
    return FormatDateTime(ftd->Time, buff, size, ftd->Spec);
}
void Locator_Default(ImPlotTicker& ticker, const ImPlotRange& range, float pixels, bool vertical, ImPlotFormatter formatter, void* formatter_data);
void Locator_Time(ImPlotTicker& ticker, const ImPlotRange& range, float pixels, bool vertical, ImPlotFormatter formatter, void* formatter_data);
void Locator_Log10(ImPlotTicker& ticker, const ImPlotRange& range, float pixels, bool vertical, ImPlotFormatter formatter, void* formatter_data);
void Locator_SymLog(ImPlotTicker& ticker, const ImPlotRange& range, float pixels, bool vertical, ImPlotFormatter formatter, void* formatter_data);
}