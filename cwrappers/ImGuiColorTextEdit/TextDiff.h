//	TextDiff - A syntax highlighting text diff widget for Dear ImGui.
//	Copyright (c) 2024-2026 Johan A. Goossens. All rights reserved.
//
//	This work is licensed under the terms of the MIT license.
//	For a copy, see <https://opensource.org/licenses/MIT>.


#pragma once


//
//	Include files
//

#include "TextEditor.h"


//
//	TextDiff
//

class IMGUI_API TextDiff : public TextEditor {
public:
	// constructor
	TextDiff();

	// specify visual mode (combined is default)
	inline void SetSideBySideMode(bool flag) { sideBySideMode = flag; }
	inline bool GetSideBySideMode() const { return sideBySideMode; }

	// specify the text to be compared (using UTF-8 encoded strings)
	void SetText(const std::string_view& left, const std::string_view& right);

	// specify a new language
	void SetLanguage(const Language* l);

	// specify the background color for added/deleted lines
	inline void SetColors(ImU32 ac, ImU32 dc) { addedColor = ac; deletedColor = dc; }

	// render the text editor in a Dear ImGui context
	void Render(const char* title, const ImVec2& size=ImVec2(), bool border=false);

private:
	// rendering mode
	bool sideBySideMode = false;

	// the two documents being compared
	TextEditor::Document leftDocument;
	TextEditor::Document rightDocument;

	// number of digits in line numbers
	int leftLineNumberDigits;
	int rightLineNumberDigits;

	// status of text line
	enum class LineStatus {
		common,
		added,
		deleted
	};

	// information about a single visible line
	class LineInfo {
	public:
		LineInfo(size_t l, size_t r, LineStatus s) : leftLine(l), rightLine(r), status(s) {}

		size_t leftLine;
		size_t rightLine;
		LineStatus status;
	};

	std::vector<LineInfo> lineInfo;
	bool updated = false;

	// colors
	ImU32 addedColor = IM_COL32(0, 150, 32, 128);
	ImU32 deletedColor = IM_COL32(180, 0, 32, 128);

	// side-by-side rendering parameters
	float leftLineNumberWidth;
	float rightLineNumberWidth;
	float textColumnWidth;
	float leftLineNumberPos;
	float leftTextPos;
	float rightLineNumberPos;
	float rightTextPos;
	float rightTextEnd;
	float textScroll;

	// split string into lines
	void splitLines(std::vector<std::string_view>& result, const std::string_view& text);

	// create a combined view of the two documents and their differences
	void createCombinedView();

	// line decorator
	void decorateLine(TextEditor::Decorator& decorator);

	// render the side-by-side mode and its parts
	void renderSideBySide(const char* title, const ImVec2& size, bool border);
	void renderSideBySideBackground();
	void renderSideBySideText();
	void renderSideBySideLine(float x, float y, TextEditor::Line& line);
	void renderSideBySideTextScrollbars();
	void renderSideBySideMiniMap();
};
