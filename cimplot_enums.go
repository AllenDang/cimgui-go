package cimgui

type ImPlotItemFlags int

const (
	ImPlotItemFlags_None     = 0
	ImPlotItemFlags_NoLegend = 1
	ImPlotItemFlags_NoFit    = 2
)

type ImPlotPieChartFlags int

const (
	ImPlotPieChartFlags_None      = 0
	ImPlotPieChartFlags_Normalize = 1024
)

type ImPlotTimeFmt int

const (
	ImPlotTimeFmt_None     = 0
	ImPlotTimeFmt_Us       = 1
	ImPlotTimeFmt_SUs      = 2
	ImPlotTimeFmt_SMs      = 3
	ImPlotTimeFmt_S        = 4
	ImPlotTimeFmt_HrMinSMs = 5
	ImPlotTimeFmt_HrMinS   = 6
	ImPlotTimeFmt_HrMin    = 7
	ImPlotTimeFmt_Hr       = 8
)

type ImPlotColormapScaleFlags int

const (
	ImPlotColormapScaleFlags_None     = 0
	ImPlotColormapScaleFlags_NoLabel  = 1
	ImPlotColormapScaleFlags_Opposite = 2
	ImPlotColormapScaleFlags_Invert   = 4
)

type ImPlotDigitalFlags int

const (
	ImPlotDigitalFlags_None = 0
)

type ImPlotErrorBarsFlags int

const (
	ImPlotErrorBarsFlags_None       = 0
	ImPlotErrorBarsFlags_Horizontal = 1024
)

type ImPlotFlags int

const (
	ImPlotFlags_None        = 0
	ImPlotFlags_NoTitle     = 1
	ImPlotFlags_NoLegend    = 2
	ImPlotFlags_NoMouseText = 4
	ImPlotFlags_NoInputs    = 8
	ImPlotFlags_NoMenus     = 16
	ImPlotFlags_NoBoxSelect = 32
	ImPlotFlags_NoChild     = 64
	ImPlotFlags_NoFrame     = 128
	ImPlotFlags_Equal       = 256
	ImPlotFlags_Crosshairs  = 512
	ImPlotFlags_CanvasOnly  = 55
)

type ImPlotHistogramFlags int

const (
	ImPlotHistogramFlags_None       = 0
	ImPlotHistogramFlags_Horizontal = 1024
	ImPlotHistogramFlags_Cumulative = 2048
	ImPlotHistogramFlags_Density    = 4096
	ImPlotHistogramFlags_NoOutliers = 8192
	ImPlotHistogramFlags_ColMajor   = 16384
)

type ImPlotScatterFlags int

const (
	ImPlotScatterFlags_None   = 0
	ImPlotScatterFlags_NoClip = 1024
)

type ImPlotStairsFlags int

const (
	ImPlotStairsFlags_None    = 0
	ImPlotStairsFlags_PreStep = 1024
)

type ImPlotAxisFlags int

const (
	ImPlotAxisFlags_None          = 0
	ImPlotAxisFlags_NoLabel       = 1
	ImPlotAxisFlags_NoGridLines   = 2
	ImPlotAxisFlags_NoTickMarks   = 4
	ImPlotAxisFlags_NoTickLabels  = 8
	ImPlotAxisFlags_NoInitialFit  = 16
	ImPlotAxisFlags_NoMenus       = 32
	ImPlotAxisFlags_Opposite      = 64
	ImPlotAxisFlags_Foreground    = 128
	ImPlotAxisFlags_Invert        = 256
	ImPlotAxisFlags_AutoFit       = 512
	ImPlotAxisFlags_RangeFit      = 1024
	ImPlotAxisFlags_PanStretch    = 2048
	ImPlotAxisFlags_LockMin       = 4096
	ImPlotAxisFlags_LockMax       = 8192
	ImPlotAxisFlags_Lock          = 12288
	ImPlotAxisFlags_NoDecorations = 15
	ImPlotAxisFlags_AuxDefault    = 66
)

type ImPlotCol int

const (
	ImPlotCol_Line          = 0
	ImPlotCol_Fill          = 1
	ImPlotCol_MarkerOutline = 2
	ImPlotCol_MarkerFill    = 3
	ImPlotCol_ErrorBar      = 4
	ImPlotCol_FrameBg       = 5
	ImPlotCol_PlotBg        = 6
	ImPlotCol_PlotBorder    = 7
	ImPlotCol_LegendBg      = 8
	ImPlotCol_LegendBorder  = 9
	ImPlotCol_LegendText    = 10
	ImPlotCol_TitleText     = 11
	ImPlotCol_InlayText     = 12
	ImPlotCol_AxisText      = 13
	ImPlotCol_AxisGrid      = 14
	ImPlotCol_AxisTick      = 15
	ImPlotCol_AxisBg        = 16
	ImPlotCol_AxisBgHovered = 17
	ImPlotCol_AxisBgActive  = 18
	ImPlotCol_Selection     = 19
	ImPlotCol_Crosshairs    = 20
	ImPlotCol_COUNT         = 21
)

type ImPlotCond int

const (
	ImPlotCond_None   = 0
	ImPlotCond_Always = 1
	ImPlotCond_Once   = 2
)

type ImPlotBin int

const (
	ImPlotBin_Sqrt    = -1
	ImPlotBin_Sturges = -2
	ImPlotBin_Rice    = -3
	ImPlotBin_Scott   = -4
)

type ImPlotDateFmt int

const (
	ImPlotDateFmt_None    = 0
	ImPlotDateFmt_DayMo   = 1
	ImPlotDateFmt_DayMoYr = 2
	ImPlotDateFmt_MoYr    = 3
	ImPlotDateFmt_Mo      = 4
	ImPlotDateFmt_Yr      = 5
)

type ImPlotInfLinesFlags int

const (
	ImPlotInfLinesFlags_None       = 0
	ImPlotInfLinesFlags_Horizontal = 1024
)

type ImPlotLegendFlags int

const (
	ImPlotLegendFlags_None            = 0
	ImPlotLegendFlags_NoButtons       = 1
	ImPlotLegendFlags_NoHighlightItem = 2
	ImPlotLegendFlags_NoHighlightAxis = 4
	ImPlotLegendFlags_NoMenus         = 8
	ImPlotLegendFlags_Outside         = 16
	ImPlotLegendFlags_Horizontal      = 32
)

type ImPlotStemsFlags int

const (
	ImPlotStemsFlags_None       = 0
	ImPlotStemsFlags_Horizontal = 1024
)

type ImPlotBarsFlags int

const (
	ImPlotBarsFlags_None       = 0
	ImPlotBarsFlags_Horizontal = 1024
)

type ImPlotColormap int

const (
	ImPlotColormap_Deep     = 0
	ImPlotColormap_Dark     = 1
	ImPlotColormap_Pastel   = 2
	ImPlotColormap_Paired   = 3
	ImPlotColormap_Viridis  = 4
	ImPlotColormap_Plasma   = 5
	ImPlotColormap_Hot      = 6
	ImPlotColormap_Cool     = 7
	ImPlotColormap_Pink     = 8
	ImPlotColormap_Jet      = 9
	ImPlotColormap_Twilight = 10
	ImPlotColormap_RdBu     = 11
	ImPlotColormap_BrBG     = 12
	ImPlotColormap_PiYG     = 13
	ImPlotColormap_Spectral = 14
	ImPlotColormap_Greys    = 15
)

type ImPlotDummyFlags int

const (
	ImPlotDummyFlags_None = 0
)

type ImAxis int

const (
	ImAxis_X1    = 0
	ImAxis_X2    = 1
	ImAxis_X3    = 2
	ImAxis_Y1    = 3
	ImAxis_Y2    = 4
	ImAxis_Y3    = 5
	ImAxis_COUNT = 6
)

type ImPlotMouseTextFlags int

const (
	ImPlotMouseTextFlags_None       = 0
	ImPlotMouseTextFlags_NoAuxAxes  = 1
	ImPlotMouseTextFlags_NoFormat   = 2
	ImPlotMouseTextFlags_ShowAlways = 4
)

type ImPlotTimeUnit int

const (
	ImPlotTimeUnit_Us    = 0
	ImPlotTimeUnit_Ms    = 1
	ImPlotTimeUnit_S     = 2
	ImPlotTimeUnit_Min   = 3
	ImPlotTimeUnit_Hr    = 4
	ImPlotTimeUnit_Day   = 5
	ImPlotTimeUnit_Mo    = 6
	ImPlotTimeUnit_Yr    = 7
	ImPlotTimeUnit_COUNT = 8
)

type ImPlotSubplotFlags int

const (
	ImPlotSubplotFlags_None       = 0
	ImPlotSubplotFlags_NoTitle    = 1
	ImPlotSubplotFlags_NoLegend   = 2
	ImPlotSubplotFlags_NoMenus    = 4
	ImPlotSubplotFlags_NoResize   = 8
	ImPlotSubplotFlags_NoAlign    = 16
	ImPlotSubplotFlags_ShareItems = 32
	ImPlotSubplotFlags_LinkRows   = 64
	ImPlotSubplotFlags_LinkCols   = 128
	ImPlotSubplotFlags_LinkAllX   = 256
	ImPlotSubplotFlags_LinkAllY   = 512
	ImPlotSubplotFlags_ColMajor   = 1024
)

type ImPlotTextFlags int

const (
	ImPlotTextFlags_None     = 0
	ImPlotTextFlags_Vertical = 1024
)

type ImPlotDragToolFlags int

const (
	ImPlotDragToolFlags_None      = 0
	ImPlotDragToolFlags_NoCursors = 1
	ImPlotDragToolFlags_NoFit     = 2
	ImPlotDragToolFlags_NoInputs  = 4
	ImPlotDragToolFlags_Delayed   = 8
)

type ImPlotHeatmapFlags int

const (
	ImPlotHeatmapFlags_None     = 0
	ImPlotHeatmapFlags_ColMajor = 1024
)

type ImPlotMarker int

const (
	ImPlotMarker_None     = -1
	ImPlotMarker_Circle   = 0
	ImPlotMarker_Square   = 1
	ImPlotMarker_Diamond  = 2
	ImPlotMarker_Up       = 3
	ImPlotMarker_Down     = 4
	ImPlotMarker_Left     = 5
	ImPlotMarker_Right    = 6
	ImPlotMarker_Cross    = 7
	ImPlotMarker_Plus     = 8
	ImPlotMarker_Asterisk = 9
	ImPlotMarker_COUNT    = 10
)

type ImPlotBarGroupsFlags int

const (
	ImPlotBarGroupsFlags_None       = 0
	ImPlotBarGroupsFlags_Horizontal = 1024
	ImPlotBarGroupsFlags_Stacked    = 2048
)

type ImPlotStyleVar int

const (
	ImPlotStyleVar_LineWeight         = 0
	ImPlotStyleVar_Marker             = 1
	ImPlotStyleVar_MarkerSize         = 2
	ImPlotStyleVar_MarkerWeight       = 3
	ImPlotStyleVar_FillAlpha          = 4
	ImPlotStyleVar_ErrorBarSize       = 5
	ImPlotStyleVar_ErrorBarWeight     = 6
	ImPlotStyleVar_DigitalBitHeight   = 7
	ImPlotStyleVar_DigitalBitGap      = 8
	ImPlotStyleVar_PlotBorderSize     = 9
	ImPlotStyleVar_MinorAlpha         = 10
	ImPlotStyleVar_MajorTickLen       = 11
	ImPlotStyleVar_MinorTickLen       = 12
	ImPlotStyleVar_MajorTickSize      = 13
	ImPlotStyleVar_MinorTickSize      = 14
	ImPlotStyleVar_MajorGridSize      = 15
	ImPlotStyleVar_MinorGridSize      = 16
	ImPlotStyleVar_PlotPadding        = 17
	ImPlotStyleVar_LabelPadding       = 18
	ImPlotStyleVar_LegendPadding      = 19
	ImPlotStyleVar_LegendInnerPadding = 20
	ImPlotStyleVar_LegendSpacing      = 21
	ImPlotStyleVar_MousePosPadding    = 22
	ImPlotStyleVar_AnnotationPadding  = 23
	ImPlotStyleVar_FitPadding         = 24
	ImPlotStyleVar_PlotDefaultSize    = 25
	ImPlotStyleVar_PlotMinSize        = 26
	ImPlotStyleVar_COUNT              = 27
)

type ImPlotScale int

const (
	ImPlotScale_Linear = 0
	ImPlotScale_Time   = 1
	ImPlotScale_Log10  = 2
	ImPlotScale_SymLog = 3
)

type ImPlotShadedFlags int

const (
	ImPlotShadedFlags_None = 0
)

type ImPlotImageFlags int

const (
	ImPlotImageFlags_None = 0
)

type ImPlotLineFlags int

const (
	ImPlotLineFlags_None     = 0
	ImPlotLineFlags_Segments = 1024
	ImPlotLineFlags_Loop     = 2048
	ImPlotLineFlags_SkipNaN  = 4096
	ImPlotLineFlags_NoClip   = 8192
)

type ImPlotLocation int

const (
	ImPlotLocation_Center    = 0
	ImPlotLocation_North     = 1
	ImPlotLocation_South     = 2
	ImPlotLocation_West      = 4
	ImPlotLocation_East      = 8
	ImPlotLocation_NorthWest = 5
	ImPlotLocation_NorthEast = 9
	ImPlotLocation_SouthWest = 6
	ImPlotLocation_SouthEast = 10
)
