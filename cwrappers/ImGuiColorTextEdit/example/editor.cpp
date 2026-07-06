//	TextEditor - A syntax highlighting text editor for ImGui
//	Copyright (c) 2024-2026 Johan A. Goossens. All rights reserved.
//
//	This work is licensed under the terms of the MIT license.
//	For a copy, see <https://opensource.org/licenses/MIT>.


//
//	Include files
//

#include <cstdio>
#include <exception>
#include <fstream>

#include "imgui.h"
#include "ImGuiFileDialog.h"

#include "editor.h"


//
//	Constants
//

#if __APPLE__
#define SHORTCUT "Cmd-"
#else
#define SHORTCUT "Ctrl-"
#endif

static const char* demo =
"// Demo C++ Code\n"
"\n"
"#include <iostream>\n"
"#include <random>\n"
"#include <vector>\n"
"\n"
"int main(int, char**) {\n"
"	std::random_device rd;\n"
"	std::mt19937 gen(rd());\n"
"	std::uniform_int_distribution<> distrib(0, 1000);\n"
"	std::vector<int> numbers;\n"
"\n"
"	for (auto i = 0; i < 100; i++) {\n"
"		numbers.emplace_back(distrib(gen));\n"
"	}\n"
"\n"
"	for (auto n : numbers) {\n"
"		std::cout << n << std::endl;\n"
"	}\n"
"\n"
"	return 0;\n"
"}\n";


//
//	Editor::Editor
//

Editor::Editor() {
	// setup text editor with demo text
	originalText = demo;
	editor.SetText(demo);
	editor.SetLanguage(TextEditor::Language::Cpp());
	version = editor.GetUndoIndex();
	filename = "untitled";
}


//
//	Editor::newFile
//

void Editor::newFile() {
	if (isDirty()) {
		showConfirmClose([this]() {
			originalText.clear();
			editor.SetText("");
			version = editor.GetUndoIndex();
			filename = "untitled";
			buildAutocompleteTrie();
		});

	} else {
		originalText.clear();
		editor.SetText("");
		version = editor.GetUndoIndex();
		filename = "untitled";
		buildAutocompleteTrie();
	}
}


//
//	Editor::openFile
//

void Editor::openFile() {
	if (isDirty()) {
		showConfirmClose([this]() {
			showFileOpen();
		});

	} else {
		showFileOpen();
	}
}

void Editor::openFile(const std::string& path) {
	try {
		std::ifstream stream(path.c_str());
		std::string text;

		stream.seekg(0, std::ios::end);
		text.reserve(stream.tellg());
		stream.seekg(0, std::ios::beg);

		text.assign((std::istreambuf_iterator<char>(stream)), std::istreambuf_iterator<char>());
		stream.close();

		originalText = text;
		editor.SetText(text);
		version = editor.GetUndoIndex();
		filename = path;

		buildAutocompleteTrie();

	} catch (std::exception& e) {
		showError(e.what());
	}
}


//
//	Editor::saveFile
//

void Editor::saveFile() {
	try {
		editor.StripTrailingWhitespaces();
		std::ofstream stream(filename.c_str());
		stream << editor.GetText();
		stream.close();
		version = editor.GetUndoIndex();

	} catch (std::exception& e) {
		showError(e.what());
	}
}


//
//	Editor::render
//

void Editor::render() {
	// create the outer window
	ImGui::SetNextWindowPos(ImVec2(0.0f, 0.0f));
	ImGui::SetNextWindowSize(ImGui::GetMainViewport()->Size);
	ImGui::PushStyleVar(ImGuiStyleVar_WindowRounding, 0.0f);
	ImGui::Begin("MainWindow", nullptr, ImGuiWindowFlags_NoDecoration | ImGuiWindowFlags_MenuBar);

	// add a menubar
	renderMenuBar();

	// render the text editor widget
	auto area = ImGui::GetContentRegionAvail();
	auto& style = ImGui::GetStyle();
	auto statusBarHeight = ImGui::GetFrameHeight() + 2.0f * style.WindowPadding.y;
	auto editorSize = ImVec2(0.0f, area.y - style.ItemSpacing.y - statusBarHeight);
	ImGui::PushFont(nullptr, fontSize);
	editor.Render("TextEditor", editorSize);
	ImGui::PopFont();

	// render a statusbar
	ImGui::Spacing();
	renderStatusBar();

	if (state ==State::diff) {
		renderDiff();

	} else if (state == State::openFile) {
		renderFileOpen();

	} else if (state == State::saveFileAs) {
		renderSaveAs();

	} else if (state == State::confirmClose) {
		renderConfirmClose();

	} else if (state == State::confirmQuit) {
		renderConfirmQuit();

	} else if (state == State::confirmError) {
		renderConfirmError();
	}

	ImGui::End();
	ImGui::PopStyleVar();
}


//
//	Editor::tryToQuit
//

void Editor::tryToQuit() {
	if (isDirty()) {
		showConfirmQuit();

	} else {
		done = true;
	}
}


//
//	Editor::renderMenuBar
//

void Editor::renderMenuBar() {
	// create menubar
	if (ImGui::BeginMenuBar()) {
		if (ImGui::BeginMenu("File")) {
			if (ImGui::MenuItem("New", SHORTCUT "N")) { newFile(); }
			if (ImGui::MenuItem("Open...", SHORTCUT "O")) { openFile(); }
			ImGui::Separator();

			if (ImGui::MenuItem("Save", SHORTCUT "S", nullptr, isSavable())) { saveFile(); }
			if (ImGui::MenuItem("Save As...")) { showSaveFileAs(); }
			ImGui::EndMenu();
		}

		if (ImGui::BeginMenu("Edit")) {
			if (ImGui::MenuItem("Undo", " " SHORTCUT "Z", nullptr, editor.CanUndo())) { editor.Undo(); }
#if __APPLE__
			if (ImGui::MenuItem("Redo", "^" SHORTCUT "Z", nullptr, editor.CanRedo())) { editor.Redo(); }
#else
			if (ImGui::MenuItem("Redo", " " SHORTCUT "Y", nullptr, editor.CanRedo())) { editor.Redo(); }
#endif

			ImGui::Separator();
			if (ImGui::MenuItem("Cut", " " SHORTCUT "X", nullptr, editor.AnyCursorHasSelection())) { editor.Cut(); }
			if (ImGui::MenuItem("Copy", " " SHORTCUT "C", nullptr, editor.AnyCursorHasSelection())) { editor.Copy(); }
			if (ImGui::MenuItem("Paste", " " SHORTCUT "V", nullptr, ImGui::GetClipboardText() != nullptr)) { editor.Paste(); }

			ImGui::Separator();
			bool flag;
			flag = editor.IsInsertSpacesOnTabs(); if (ImGui::MenuItem("Insert Spaces on Tabs", nullptr, &flag)) { editor.SetInsertSpacesOnTabs(flag); };

			if (ImGui::MenuItem("Tabs To Spaces")) { editor.TabsToSpaces(); }
			if (ImGui::MenuItem("Spaces To Tabs", nullptr, nullptr, !editor.IsInsertSpacesOnTabs())) { editor.SpacesToTabs(); }
			if (ImGui::MenuItem("Strip Trailing Whitespaces")) { editor.StripTrailingWhitespaces(); }

			ImGui::EndMenu();
		}

		if (ImGui::BeginMenu("Selection")) {
			if (ImGui::MenuItem("Select All", " " SHORTCUT "A", nullptr, !editor.IsEmpty())) { editor.SelectAll(); }
			ImGui::Separator();

			if (ImGui::MenuItem("Indent Line(s)", " " SHORTCUT "]", nullptr, !editor.IsEmpty())) { editor.IndentLines(); }
			if (ImGui::MenuItem("Deindent Line(s)", " " SHORTCUT "[", nullptr, !editor.IsEmpty())) { editor.DeindentLines(); }
			if (ImGui::MenuItem("Move Line(s) Up", nullptr, nullptr, !editor.IsEmpty())) { editor.MoveUpLines(); }
			if (ImGui::MenuItem("Move Line(s) Down", nullptr, nullptr, !editor.IsEmpty())) { editor.MoveDownLines(); }
			if (ImGui::MenuItem("Toggle Comments", " " SHORTCUT "/", nullptr, editor.HasLanguage())) { editor.ToggleComments(); }
			ImGui::Separator();

			if (ImGui::MenuItem("To Uppercase", nullptr, nullptr, editor.AnyCursorHasSelection())) { editor.SelectionToUpperCase(); }
			if (ImGui::MenuItem("To Lowercase", nullptr, nullptr, editor.AnyCursorHasSelection())) { editor.SelectionToLowerCase(); }
			ImGui::Separator();

			if (ImGui::MenuItem("Add Next Occurrence", " " SHORTCUT "D", nullptr, editor.CurrentCursorHasSelection())) { editor.AddNextOccurrence(); }
			if (ImGui::MenuItem("Select All Occurrences", "^" SHORTCUT "D", nullptr, editor.CurrentCursorHasSelection())) { editor.SelectAllOccurrences(); }

			ImGui::EndMenu();
		}

		if (ImGui::BeginMenu("View")) {
			if (ImGui::MenuItem("Zoom In", " " SHORTCUT "+")) { increaseFontSIze(); }
			if (ImGui::MenuItem("Zoom Out", " " SHORTCUT "-")) { decreaseFontSIze(); }
			ImGui::Separator();

			bool flag;
			if (ImGui::MenuItem("Autocomplete", nullptr, &autocomplete)) { setAutocompleteMode(autocomplete); }
			flag = editor.IsShowWhitespacesEnabled(); if (ImGui::MenuItem("Show Whitespaces", nullptr, &flag)) { editor.SetShowWhitespacesEnabled(flag); };
			flag = editor.IsShowSpacesEnabled(); if (ImGui::MenuItem("Show Spaces", nullptr, &flag)) { editor.SetShowSpacesEnabled(flag); };
			flag = editor.IsShowTabsEnabled(); if (ImGui::MenuItem("Show Tabs", nullptr, &flag)) { editor.SetShowTabsEnabled(flag); };
			flag = editor.IsShowLineNumbersEnabled(); if (ImGui::MenuItem("Show Line Numbers", nullptr, &flag)) { editor.SetShowLineNumbersEnabled(flag); };
			flag = editor.IsShowingMatchingBrackets(); if (ImGui::MenuItem("Show Matching Brackets", nullptr, &flag)) { editor.SetShowMatchingBrackets(flag); };
			flag = editor.IsCompletingPairedGlyphs(); if (ImGui::MenuItem("Complete Matching Glyphs", nullptr, &flag)) { editor.SetCompletePairedGlyphs(flag); };
			flag = editor.IsShowPanScrollIndicatorEnabled(); if (ImGui::MenuItem("Show Pan/Scroll Indicator", nullptr, &flag)) { editor.SetShowPanScrollIndicatorEnabled(flag); };
			flag = editor.IsMiddleMousePanMode(); if (ImGui::MenuItem("Middle Mouse Pan Mode", nullptr, &flag)) { if (flag) editor.SetMiddleMousePanMode(); else editor.SetMiddleMouseScrollMode(); };

			ImGui::Separator();
			if (ImGui::MenuItem("Show Diff", " " SHORTCUT "I")) { showDiff(); }

			ImGui::EndMenu();
		}

		if (ImGui::BeginMenu("Find")) {
			if (ImGui::MenuItem("Find", " " SHORTCUT "F")) { editor.OpenFindReplaceWindow(); }
			if (ImGui::MenuItem("Find Next", " " SHORTCUT "G", nullptr, editor.HasFindString())) { editor.FindNext(); }
			if (ImGui::MenuItem("Find All", "^" SHORTCUT "G", nullptr, editor.HasFindString())) { editor.FindAll(); }
			ImGui::Separator();
			ImGui::EndMenu();
		}

		ImGui::EndMenuBar();
	}

	// handle keyboard shortcuts
	if (ImGui::IsWindowFocused(ImGuiFocusedFlags_RootAndChildWindows) && !ImGui::GetIO().WantCaptureKeyboard) {
		if (ImGui::IsKeyDown(ImGuiMod_Ctrl)) {
			if (ImGui::IsKeyPressed(ImGuiKey_N)) { newFile(); }
			else if (ImGui::IsKeyPressed(ImGuiKey_O)) { openFile(); }
			else if (ImGui::IsKeyPressed(ImGuiKey_S)) { if (filename == "untitled") { showSaveFileAs(); } else { saveFile(); } }
			else if (ImGui::IsKeyPressed(ImGuiKey_I)) { showDiff(); }
			else if (ImGui::IsKeyPressed(ImGuiKey_Equal)) { increaseFontSIze(); }
			else if (ImGui::IsKeyPressed(ImGuiKey_Minus)) { decreaseFontSIze(); }
		}
	}
}


//
//	Editor::renderStatusBar
//

void Editor::renderStatusBar() {
	// language support
	static const char* languages[] = {"C", "C++", "Cs", "AngelScript", "Lua", "Python", "GLSL", "HLSL",  "JSON", "Markdown", "SQL"};

	static const TextEditor::Language* definitions[] = {
		TextEditor::Language::C(),
		TextEditor::Language::Cpp(),
		TextEditor::Language::Cs(),
		TextEditor::Language::AngelScript(),
		TextEditor::Language::Lua(),
		TextEditor::Language::Python(),
		TextEditor::Language::Glsl(),
		TextEditor::Language::Hlsl(),
		TextEditor::Language::Json(),
		TextEditor::Language::Markdown(),
		TextEditor::Language::Sql()
	};

	std::string language = editor.GetLanguageName();

	// create a statusbar window
	ImGui::PushStyleColor(ImGuiCol_ChildBg, ImVec4(0.15f, 0.15f, 0.15f, 1.0f));
	ImGui::BeginChild("StatusBar", ImVec2(0.0f, 0.0f), ImGuiChildFlags_Borders);
	ImGui::SetNextItemWidth(120.0f);

	// allow user to select language for colorizing
	if (ImGui::BeginCombo("##LanguageSelector", language.c_str())) {
		for (int n = 0; n < IM_ARRAYSIZE(languages); n++) {
			bool selected = (language == languages[n]);

			if (ImGui::Selectable(languages[n], selected)) {
				editor.SetLanguage(definitions[n]);
				buildAutocompleteTrie();
			}

			if (selected) {
				ImGui::SetItemDefaultFocus();
			}
		}

		ImGui::EndCombo();
	}

	// determine horizontal gap so the rest is right aligned
	ImGui::SameLine(0.0f, 0.0f);
	ImGui::AlignTextToFramePadding();

	int line;
	int column;
	int tabSize = editor.GetTabSize();
	float glyphWidth = ImGui::CalcTextSize("#").x;
	editor.GetCurrentCursor(line, column);

	// determine status message
	char status[256];

	auto statusSize = std::snprintf(
		status,
		sizeof(status),
		"Ln %d, Col %d  Tab Size: %d  File: %s",
		line + 1,
		column + 1,
		tabSize,
		filename.c_str()
	);

	auto size = glyphWidth * (statusSize + 3);
	auto width = ImGui::GetContentRegionAvail().x;

	ImGui::SameLine(0.0f, width - size);
	ImGui::TextUnformatted(status);

	// render "text dirty" indicator
	ImGui::SameLine(0.0f, glyphWidth * 1.0f);
	auto drawlist = ImGui::GetWindowDrawList();
	auto pos = ImGui::GetCursorScreenPos();
	auto offset = ImGui::GetFrameHeight() * 0.5f;
	auto radius = offset * 0.6f;
	auto color = isDirty() ? IM_COL32(164, 0, 0, 255) : IM_COL32(164, 164, 164, 255);
	drawlist->AddCircleFilled(ImVec2(pos.x + offset, pos.y + offset), radius, color);

	ImGui::EndChild();
	ImGui::PopStyleColor();
}


//
//	Editor::showDiff
//

void Editor::showDiff() {
	diff.SetLanguage(editor.GetLanguage());
	diff.SetText(originalText, editor.GetText());
	state = State::diff;
}


//
//	Editor::showFileOpen
//

void Editor::showFileOpen() {
	// open a file selector dialog
	IGFD::FileDialogConfig config;
	config.countSelectionMax = 1;

	config.flags =
		ImGuiFileDialogFlags_Modal |
		ImGuiFileDialogFlags_DontShowHiddenFiles |
		ImGuiFileDialogFlags_ReadOnlyFileNameField;

	ImGuiFileDialog::Instance()->OpenDialog("file-open", "Select File to Open...", ".*", config);
	state = State::openFile;
}


//
//	Editor::showSaveFileAs
//

void Editor::showSaveFileAs() {
	IGFD::FileDialogConfig config;
	config.countSelectionMax = 1;

	config.flags =
		ImGuiFileDialogFlags_Modal |
		ImGuiFileDialogFlags_DontShowHiddenFiles |
		ImGuiFileDialogFlags_ConfirmOverwrite;

	ImGuiFileDialog::Instance()->OpenDialog("file-saveas", "Save File as...", "*", config);
	state = State::saveFileAs;
}


//
//	Editor::showConfirmClose
//

void Editor::showConfirmClose(std::function<void()> callback) {
	onConfirmClose = callback;
	state = State::confirmClose;
}


//
//	Editor::showConfirmQuit
//

void Editor::showConfirmQuit() {
	state = State::confirmQuit;
}


//
//	Editor::showError
//

void Editor::showError(const std::string &message) {
	errorMessage = message;
	state = State::confirmError;
}


//
//	Editor::renderDiff
//

void Editor::renderDiff() {
	ImGui::OpenPopup("Changes since Opening File##diff");
	auto viewport = ImGui::GetMainViewport();
	ImVec2 center = viewport->GetCenter();
	ImGui::SetNextWindowPos(center, ImGuiCond_Appearing, ImVec2(0.5f, 0.5f));

	if (ImGui::BeginPopupModal("Changes since Opening File##diff", NULL, ImGuiWindowFlags_AlwaysAutoResize)) {
		diff.Render("diff", viewport->Size * 0.8f, true);

		ImGui::Separator();
		static constexpr float buttonWidth = 80.0f;
		auto buttonOffset = ImGui::GetContentRegionAvail().x - buttonWidth;
		bool sideBySide = diff.GetSideBySideMode();

		if (ImGui::Checkbox("Show side-by-side", &sideBySide)) {
			diff.SetSideBySideMode(sideBySide);
		}

		ImGui::SameLine();
		ImGui::Indent(buttonOffset);

		if (ImGui::Button("OK", ImVec2(buttonWidth, 0.0f)) || ImGui::IsKeyPressed(ImGuiKey_Escape, false)) {
			state = State::edit;
			ImGui::CloseCurrentPopup();
		}

		ImGui::EndPopup();
	}
}


//
//	Editor::renderFileOpen
//

void Editor::renderFileOpen() {
	// handle file open dialog
	ImVec2 maxSize = ImGui::GetMainViewport()->Size;
	ImVec2 minSize = maxSize * 0.5f;
	auto dialog = ImGuiFileDialog::Instance();

	ImVec2 center = ImGui::GetMainViewport()->GetCenter();
	ImGui::SetNextWindowPos(center, ImGuiCond_Always, ImVec2(0.5f, 0.5f));

	if (dialog->Display("file-open", ImGuiWindowFlags_NoCollapse, minSize, maxSize)) {
		// open selected file (if required)
		if (dialog->IsOk()) {
			openFile(dialog->GetFilePathName());
			state = State::edit;
		}

		// close dialog
		dialog->Close();
	}
}


//
//	Editor::renderSaveAs
//

void Editor::renderSaveAs() {
	// handle saveas dialog
	ImVec2 maxSize = ImGui::GetMainViewport()->Size;
	ImVec2 minSize = maxSize * 0.5f;
	auto dialog = ImGuiFileDialog::Instance();

	ImVec2 center = ImGui::GetMainViewport()->GetCenter();
	ImGui::SetNextWindowPos(center, ImGuiCond_Always, ImVec2(0.5f, 0.5f));

	if (dialog->Display("file-saveas", ImGuiWindowFlags_NoCollapse, minSize, maxSize)) {
		// open selected file if required
		if (dialog->IsOk()) {
			filename = dialog->GetFilePathName();
			saveFile();
			state = State::edit;

		} else {
			state = State::edit;
		}

		// close dialog
		dialog->Close();
	}
}


//
//	Editor::renderConfirmClose
//

void Editor::renderConfirmClose() {
	ImGui::OpenPopup("Confirm Close");
	ImVec2 center = ImGui::GetMainViewport()->GetCenter();
	ImGui::SetNextWindowPos(center, ImGuiCond_Appearing, ImVec2(0.5f, 0.5f));

	if (ImGui::BeginPopupModal("Confirm Close", NULL, ImGuiWindowFlags_AlwaysAutoResize)) {
		ImGui::Text("This file has changed!\nDo you really want to delete it?\n\n");
		ImGui::Separator();

		static constexpr float buttonWidth = 80.0f;
		ImGui::Indent(ImGui::GetContentRegionAvail().x - buttonWidth * 2.0f - ImGui::GetStyle().ItemSpacing.x);

		if (ImGui::Button("OK", ImVec2(buttonWidth, 0.0f))) {
			state = State::edit;
			onConfirmClose();
			ImGui::CloseCurrentPopup();
		}

		ImGui::SameLine();

		if (ImGui::Button("Cancel", ImVec2(buttonWidth, 0.0f)) || ImGui::IsKeyPressed(ImGuiKey_Escape, false)) {
			state = State::edit;
			ImGui::CloseCurrentPopup();
		}

		ImGui::EndPopup();
	}
}


//
//	Editor::renderConfirmQuit
//

void Editor::renderConfirmQuit() {
	ImGui::OpenPopup("Quit Editor?");
	ImVec2 center = ImGui::GetMainViewport()->GetCenter();
	ImGui::SetNextWindowPos(center, ImGuiCond_Appearing, ImVec2(0.5f, 0.5f));

	if (ImGui::BeginPopupModal("Quit Editor?", NULL, ImGuiWindowFlags_AlwaysAutoResize)) {
		ImGui::Text("Your text has changed and is not saved!\nDo you really want to quit?\n\n");
		ImGui::Separator();

		static constexpr float buttonWidth = 80.0f;
		ImGui::Indent(ImGui::GetContentRegionAvail().x - buttonWidth * 2.0f - ImGui::GetStyle().ItemSpacing.x);

		if (ImGui::Button("OK", ImVec2(buttonWidth, 0.0f))) {
			done = true;
			state = State::edit;
			ImGui::CloseCurrentPopup();
		}

		ImGui::SameLine();

		if (ImGui::Button("Cancel", ImVec2(buttonWidth, 0.0f)) || ImGui::IsKeyPressed(ImGuiKey_Escape, false)) {
			state = State::edit;
			ImGui::CloseCurrentPopup();
		}

		ImGui::EndPopup();
	}
}


//
//	Editor::renderConfirmError
//

void Editor::renderConfirmError() {
	ImGui::OpenPopup("Error");
	ImVec2 center = ImGui::GetMainViewport()->GetCenter();
	ImGui::SetNextWindowPos(center, ImGuiCond_Appearing, ImVec2(0.5f, 0.5f));

	if (ImGui::BeginPopupModal("Error", NULL, ImGuiWindowFlags_AlwaysAutoResize)) {
		ImGui::Text("%s\n", errorMessage.c_str());
		ImGui::Separator();

		static constexpr float buttonWidth = 80.0f;
		ImGui::Indent(ImGui::GetContentRegionAvail().x - buttonWidth);

		if (ImGui::Button("OK", ImVec2(buttonWidth, 0.0f)) || ImGui::IsKeyPressed(ImGuiKey_Escape, false)) {
			errorMessage.clear();
			state = State::edit;
			ImGui::CloseCurrentPopup();
		}

		ImGui::EndPopup();
	}
}


//
//	Editor::setAutocompleteMode
//

void Editor::setAutocompleteMode(bool flag) {
	// see we are turning autocomplete on or off
	if (flag) {
		// rebuild word list
		buildAutocompleteTrie();

		// setup autocomplete by submitting a new configuration
		TextEditor::AutoCompleteConfig config;

		config.callback = [this](TextEditor::AutoCompleteState& state) {
			trie.findSuggestions(state.suggestions, state.searchTerm);
		};

		editor.SetAutoCompleteConfig(&config);

		// enable change tracking
		// we don't track every keystroke, callbacks can be delayed up to 3000 milliseconds
		// if you want live tracking, change the 3000 to 0 (performance hit will be minimal for small documents)
		editor.SetChangeCallback([this]() {
			buildAutocompleteTrie();
		}, 3000);

	} else {
		// disable autocomplete and change tracking
		editor.SetAutoCompleteConfig(nullptr);
		editor.SetChangeCallback(nullptr);
	}
}


//
//	Editor::buildAutocompleteTrie
//

void Editor::buildAutocompleteTrie() {
	// empty list first
	trie.clear();

	// add language words (if required)
	auto language = editor.GetLanguage();

	if (language) {
		for (auto& word : language->keywords) { trie.insert(word); }
		for (auto& word : language->declarations) { trie.insert(word); }
		for (auto& word : language->identifiers) { trie.insert(word); }
	}

	// add all identifiers in current document
	editor.IterateIdentifiers([this](const std::string& identifier) {
		trie.insert(identifier);
	});
}
