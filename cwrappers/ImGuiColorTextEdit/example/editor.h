//	TextEditor - A syntax highlighting text editor for ImGui
//	Copyright (c) 2024-2026 Johan A. Goossens. All rights reserved.
//
//	This work is licensed under the terms of the MIT license.
//	For a copy, see <https://opensource.org/licenses/MIT>.


#pragma once


//
//	Include files
//

#include <algorithm>
#include <functional>
#include <string>

#include "../TextEditor.h"
#include "../TextDiff.h"


//
//	Editor
//

class Editor {
public:
	// constructor
	Editor();

	// file related functions
	void newFile();
	void openFile();
	void openFile(const std::string& path);
	void saveFile();

	// manage program exit
	void tryToQuit();
	inline bool isDone() const { return done; }

	// render the editor
	void render();

	private:
	// private functions
	void renderMenuBar();
	void renderStatusBar();

	void showDiff();
	void showFileOpen();
	void showSaveFileAs();
	void showConfirmClose(std::function<void()> callback);
	void showConfirmQuit();
	void showError(const std::string& message);

	void renderDiff();
	void renderFileOpen();
	void renderSaveAs();
	void renderConfirmClose();
	void renderConfirmQuit();
	void renderConfirmError();

	void setAutocompleteMode(bool flag);
	void buildAutocompleteTrie();

	inline bool isDirty() const { return editor.GetUndoIndex() != version; }
	inline bool isSavable() const { return isDirty() && filename != "untitled"; }

	// properties
	std::string originalText;
	TextEditor editor;
	TextDiff diff;
	std::string filename;
	size_t version;
	bool done = false;
	std::string errorMessage;
	std::function<void()> onConfirmClose;

	float fontSize = 17.0f;
	inline void increaseFontSIze() { fontSize = std::clamp(fontSize + 1.0f, 8.0f, 24.0f); }
	inline void decreaseFontSIze() { fontSize = std::clamp(fontSize - 1.0f, 8.0f, 24.0f); }

	bool autocomplete = false;
	TextEditor::Trie trie;

	// editor state
	enum class State {
		edit,
		diff,
		newFile,
		openFile,
		saveFileAs,
		confirmClose,
		confirmQuit,
		confirmError
	} state = State::edit;
};
