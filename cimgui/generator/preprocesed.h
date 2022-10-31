
 bool ImGui_ImplOpenGL3_Init(const char* glsl_version = nullptr);
 void ImGui_ImplOpenGL3_Shutdown();
 void ImGui_ImplOpenGL3_NewFrame();
 void ImGui_ImplOpenGL3_RenderDrawData(ImDrawData* draw_data);
 bool ImGui_ImplOpenGL3_CreateFontsTexture();
 void ImGui_ImplOpenGL3_DestroyFontsTexture();
 bool ImGui_ImplOpenGL3_CreateDeviceObjects();
 void ImGui_ImplOpenGL3_DestroyDeviceObjects();