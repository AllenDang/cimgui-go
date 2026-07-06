//	TextDiff - A syntax highlighting text diff widget for Dear ImGui.
//	Copyright (c) 2024-2026 Johan A. Goossens. All rights reserved.
//
//	This work is licensed under the terms of the MIT license.
//	For a copy, see <https://opensource.org/licenses/MIT>.


//
//	Include files
//

#include <cmath>

#include "imgui_internal.h"

#include "dtl.h"
#include "TextDiff.h"


//
//	TextDiff::TextDiff
//

TextDiff::TextDiff() {
	readOnly = true;
	showLineNumbers = false;
	showMatchingBrackets = false;
}


//
//	TextDiff::SetText
//

void TextDiff::SetText(const std::string_view& left, const std::string_view& right) {
	// split text into lines
	std::vector<std::string_view> leftLines;
	std::vector<std::string_view> rightLines;
	splitLines(leftLines, left);
	splitLines(rightLines, right);

	// create two documents and colorize them
	leftDocument.setText(leftLines);
	rightDocument.setText(rightLines);

	colorizer.updateEntireDocument(leftDocument, language);
	colorizer.updateEntireDocument(rightDocument, language);

	// setup line number decoration
	leftLineNumberDigits = static_cast<int>(std::log10(leftDocument.lineCount() + 1) + 1.0f);
	rightLineNumberDigits = static_cast<int>(std::log10(leftDocument.lineCount() + 1) + 1.0f);
	decoratorWidth = -(leftLineNumberDigits + rightLineNumberDigits + 4.0f);

	decoratorCallback = [this](TextEditor::Decorator& decorator) {
		this->decorateLine(decorator);
	};

	// calculate the difference
	dtl::Diff<std::string_view> diff(leftLines, rightLines);
	diff.compose();
	auto ranges = diff.getSes().getSequence();

	// determine status of each line in diff view
	lineInfo.clear();
	size_t leftIndex = 0;
	size_t rightIndex = 0;

	for (auto& [text, info] : ranges) {
		switch (info.type) {
			case dtl::SES_COMMON:
				lineInfo.emplace_back(leftIndex++, rightIndex++, LineStatus::common);
				break;

			case dtl::SES_ADD:
				lineInfo.emplace_back(leftIndex, rightIndex++, LineStatus::added);
				break;

			case dtl::SES_DELETE:
				lineInfo.emplace_back(leftIndex++, rightIndex, LineStatus::deleted);
				break;
		}
	}

	// set flag
	updated = true;
}


//
//	TextDiff::SetLanguage
//

void TextDiff::SetLanguage(const Language* l) {
	language = l;
	colorizer.updateEntireDocument(leftDocument, language);
	colorizer.updateEntireDocument(rightDocument, language);
	updated = true;
}


//
//	TextDiff::Render
//

void TextDiff::Render(const char* title, const ImVec2& size, bool border) {
	if (sideBySideMode) {
		// render a custom side-by-side view
		renderSideBySide(title, size, border);

	} else {
		// create a combined view (if required)
		if (updated) {
			createCombinedView();
			updated = false;
		}

		// render combined view as a normal read-only text editor
		TextEditor::render(title, size, border);
	}
}


//
//	TextDiff::splitLines
//

void TextDiff::splitLines(std::vector<std::string_view>& result, const std::string_view& text) {
	size_t prev = CodePoint::skipBOM(text.begin(), text.end()) - text.begin();
	size_t pos;

	while ((pos = text.find('\n', prev)) != std::string_view::npos) {
		result.push_back(text.substr(prev, pos - prev));
		prev = pos + 1;
	}

	result.push_back(text.substr(prev));
}


//
//	TextDiff::createCombinedView
//

void TextDiff::createCombinedView() {
	document.clear();
	cursors.clearAll();
	clearMarkers();

	for (size_t i = 0; i < lineInfo.size(); i++) {
		auto& line = lineInfo[i];

		switch(line.status) {
			case LineStatus::common:
				document.emplace_back(leftDocument[line.leftLine]);
				break;

			case LineStatus::added:
				document.emplace_back(rightDocument[line.rightLine]);
				addMarker(static_cast<int>(i), 0, addedColor, "", "");
				break;

			case LineStatus::deleted:
				document.emplace_back(leftDocument[line.leftLine]);
				addMarker(static_cast<int>(i), 0, deletedColor, "", "");
				break;
		}
	}

	document.updateMaximumColumn(0, document.lineCount() - 1);
}


//
//	TextDiff::decorateLine
//

void TextDiff::decorateLine(TextEditor::Decorator& decorator) {
	auto& line = lineInfo[decorator.line];
	auto lineLeft = static_cast<int>(line.leftLine + 1);
	auto lineRight = static_cast<int>(line.rightLine + 1);
	auto color = ImGui::ColorConvertU32ToFloat4(palette.get(Color::lineNumber));

	switch(line.status) {
		case LineStatus::common:
			ImGui::TextColored(color, " %*d %*d  ", leftLineNumberDigits, lineLeft, rightLineNumberDigits, lineRight);
			break;

		case LineStatus::added:
			ImGui::TextColored(color, " %*s %*d +", leftLineNumberDigits, "", rightLineNumberDigits, lineRight);
			break;

		case LineStatus::deleted:
			ImGui::TextColored(color, " %*d %*s -", leftLineNumberDigits, lineLeft, rightLineNumberDigits, "");
			break;
	}
}


//
//	TextDiff::renderSideBySide
//

void TextDiff::renderSideBySide(const char* title, const ImVec2& size, bool border) {
	// update color palette (if required)
	auto& style = ImGui::GetStyle();

	if (paletteAlpha != style.Alpha) {
		updatePalette();
	}

	// get font information
	font = ImGui::GetFont();
	fontSize = ImGui::GetFontSize();
	glyphSize = ImVec2(ImGui::CalcTextSize("#").x, ImGui::GetTextLineHeightWithSpacing() * lineSpacing);

	// scroll to specified line (if required)
	if (scrollToLineNumber >= 0) {
		scrollToLineNumber = std::min(scrollToLineNumber, static_cast<int>(lineInfo.size()));

		switch (scrollToAlignment) {
			case Scroll::alignTop:
				ImGui::SetNextWindowScroll(ImVec2(0.0f, std::max(0.0f, static_cast<float>(scrollToLineNumber) * glyphSize.y)));
				break;

			case Scroll::alignMiddle:
				ImGui::SetNextWindowScroll(ImVec2(0.0f, std::max(0.0f, static_cast<float>(scrollToLineNumber - visibleLines / 2) * glyphSize.y)));
				break;

			case Scroll::alignBottom:
				ImGui::SetNextWindowScroll(ImVec2(0.0f, std::max(0.0f, static_cast<float>(scrollToLineNumber - (visibleLines - 1)) * glyphSize.y)));
				break;
		}

		scrollToLineNumber = -1;
	}

	// ensure diff has focus (if required)
	if (focusOnEditor) {
		ImGui::SetNextWindowFocus();
		focusOnEditor = false;
	}

	// start rendering the widget
	ImGui::SetNextWindowContentSize(ImVec2(0.0f, glyphSize.y * lineInfo.size()));
	ImGui::PushStyleVar(ImGuiStyleVar_ItemSpacing, ImVec2(0.0f, 0.0f));
	ImGui::PushStyleColor(ImGuiCol_ChildBg, ImGui::ColorConvertU32ToFloat4(palette.get(Color::background)));
	ImGui::BeginChild(title, size, border, ImGuiWindowFlags_NoMove | ImGuiWindowFlags_HorizontalScrollbar | ImGuiWindowFlags_NoNavInputs);

	auto cursorScreenPos = ImGui::GetCursorScreenPos();
	auto visibleSize = ImGui::GetCurrentWindow()->Rect().GetSize();

	// determine view parameters
	leftLineNumberWidth = glyphSize.x * (leftLineNumberDigits + 4);
	rightLineNumberWidth = glyphSize.x * (rightLineNumberDigits + 4);
	textColumnWidth = static_cast<float>(visibleSize.x - leftLineNumberWidth - rightLineNumberWidth) / 2.0f;

	leftLineNumberPos = cursorScreenPos.x;
	leftTextPos = leftLineNumberPos + leftLineNumberWidth;
	rightLineNumberPos = leftTextPos + textColumnWidth;
	rightTextPos = rightLineNumberPos + rightLineNumberWidth;
	rightTextEnd = rightTextPos + textColumnWidth;

	visibleLines = std::max(static_cast<int>(std::ceil(visibleSize.y / glyphSize.y)), 0);
	visibleColumns = std::max(static_cast<int>(std::ceil(textColumnWidth / glyphSize.x)), 0);

	firstVisibleColumn = std::max(static_cast<int>(std::floor(textScroll / glyphSize.x)), 0);
	lastVisibleColumn = static_cast<int>(std::floor((textScroll + textColumnWidth) / glyphSize.x));
	firstVisibleLine = std::max(static_cast<int>(std::floor(ImGui::GetScrollY() / glyphSize.y)), 0);
	lastVisibleLine = std::min(static_cast<int>(std::floor((ImGui::GetScrollY() + visibleSize.y) / glyphSize.y)), static_cast<int>(lineInfo.size() - 1));

	renderSideBySideBackground();
	renderSideBySideText();
	renderSideBySideTextScrollbars();
	renderSideBySideMiniMap();

	ImGui::EndChild();
	ImGui::PopStyleColor();
	ImGui::PopStyleVar();
}


//
//	TextDiff::renderSideBySideBackground
//

void TextDiff::renderSideBySideBackground() {
	// render line numbers and text backgrounds
	auto drawList = ImGui::GetWindowDrawList();
	auto y = ImGui::GetCursorScreenPos().y + firstVisibleLine * glyphSize.y;
	char buffer[32];

	for (auto i = firstVisibleLine; i <= lastVisibleLine; i++) {
		auto& line = lineInfo[i];
		auto lineLeft = static_cast<int>(line.leftLine + 1);
		auto lineRight = static_cast<int>(line.rightLine + 1);

		switch(line.status) {
			case LineStatus::common:
				snprintf(buffer, sizeof(buffer), " %*d", leftLineNumberDigits, lineLeft);
				drawList->AddText(ImVec2(leftLineNumberPos, y), palette.get(Color::lineNumber), buffer);
				snprintf(buffer, sizeof(buffer), " %*d", rightLineNumberDigits, lineRight);
				drawList->AddText(ImVec2(rightLineNumberPos, y), palette.get(Color::lineNumber), buffer);
				break;

			case LineStatus::added:
				snprintf(buffer, sizeof(buffer), " %*d +", rightLineNumberDigits, lineRight);
				drawList->AddText(ImVec2(rightLineNumberPos, y), palette.get(Color::lineNumber), buffer);

				drawList->AddRectFilled(
					ImVec2(rightTextPos, y),
					ImVec2(rightTextEnd, y + glyphSize.y),
					addedColor);

				break;

			case LineStatus::deleted:
				snprintf(buffer, sizeof(buffer), " %*d -", leftLineNumberDigits, lineLeft);
				drawList->AddText(ImVec2(leftLineNumberPos, y), palette.get(Color::lineNumber), buffer);

				drawList->AddRectFilled(
					ImVec2(leftTextPos, y),
					ImVec2(rightLineNumberPos, y + glyphSize.y),
					deletedColor);

				break;
		}

		y += glyphSize.y;
	}
}


//
//	TextDiff::renderSideBySideText
//

void TextDiff::renderSideBySideText() {
	// setup rendering
	auto drawList = ImGui::GetWindowDrawList();
	auto cursorScreenPos = ImGui::GetCursorScreenPos();
	auto yTop = drawList->GetClipRectMin().y;
	auto yBottom = drawList->GetClipRectMax().y;

	// render left text
	drawList->PushClipRect(ImVec2(leftTextPos, yTop), ImVec2(rightLineNumberPos, yBottom), false);

	for (auto i = firstVisibleLine; i <= lastVisibleLine; i++) {
		auto& line = lineInfo[i];
		auto y = cursorScreenPos.y + i * glyphSize.y;

		switch(line.status) {
			case LineStatus::common:
			case LineStatus::deleted:
				renderSideBySideLine(leftTextPos, y, leftDocument[lineInfo[i].leftLine]);
				break;

			case LineStatus::added:
				break;
		}
	}

	drawList->PopClipRect();

	// render right text
	drawList->PushClipRect(ImVec2(rightTextPos, yTop), ImVec2(rightTextEnd, yBottom), false);

	for (auto i = firstVisibleLine; i <= lastVisibleLine; i++) {
		auto& line = lineInfo[i];
		auto y = cursorScreenPos.y + i * glyphSize.y;

		switch(line.status) {
			case LineStatus::common:
			case LineStatus::added:
				renderSideBySideLine(rightTextPos, y, rightDocument[lineInfo[i].rightLine]);
				break;

			case LineStatus::deleted:
				break;
		}
	}

	drawList->PopClipRect();
}


//
//	TextDiff::renderSideBySideLine
//

void TextDiff::renderSideBySideLine(float x, float y, TextEditor::Line& line) {
	// draw colored glyphs for specified line
	auto drawList = ImGui::GetWindowDrawList();
	auto tabSize = document.getTabSize();
	auto firstRenderableColumn = (firstVisibleColumn / tabSize) * tabSize;

	auto column = firstRenderableColumn;
	auto index = document.getIndex(line, column);
	auto lineSize = line.size();

	while (index < lineSize && column <= lastVisibleColumn) {
		auto& glyph = line[index];
		auto codepoint = glyph.codepoint;
		ImVec2 glyphPos(x + column * glyphSize.x - textScroll, y);

		if (codepoint == '\t') {
			if (showTabs) {
				const auto x1 = glyphPos.x + glyphSize.x * 0.3f;
				const auto y1 = glyphPos.y + fontSize * 0.5f;
				const auto x2 = glyphPos.x + glyphSize.x;

				ImVec2 p1, p2, p3, p4;
				p1 = ImVec2(x1, y1);
				p2 = ImVec2(x2, y1);
				p3 = ImVec2(x2 - fontSize * 0.16f, y1 - fontSize * 0.16f);
				p4 = ImVec2(x2 - fontSize * 0.16f, y1 + fontSize * 0.16f);

				drawList->AddLine(p1, p2, palette.get(Color::whitespace));
				drawList->AddLine(p2, p3, palette.get(Color::whitespace));
				drawList->AddLine(p2, p4, palette.get(Color::whitespace));
			}

		} else if (codepoint == ' ') {
			if (showSpaces) {
				const auto x1 = glyphPos.x + glyphSize.x * 0.5f;
				const auto y1 = glyphPos.y + fontSize * 0.5f;
				drawList->AddCircleFilled(ImVec2(x1, y1), 1.5f, palette.get(Color::whitespace), 4);
			}

		} else {
			font->RenderChar(drawList, fontSize, glyphPos, palette.get(glyph.color), codepoint);
		}

		index++;
		column += (codepoint == '\t') ? tabSize - (column % tabSize) : 1;
	}
}


//
//	TextDiff::renderSideBySideTextScrollbars
//

void TextDiff::renderSideBySideTextScrollbars() {
	auto maxColumnsWidth = std::max(leftDocument.getMaxColumn(), rightDocument.getMaxColumn()) * glyphSize.x;
	auto visibleColumnsWidth = rightLineNumberPos - leftTextPos;

	if (maxColumnsWidth > visibleColumnsWidth) {
		auto window = ImGui::GetCurrentWindow();
		const ImRect outerRect = window->Rect();
		auto borderSize = std::round(window->WindowBorderSize * 0.5f);
		auto scrollbarSize = ImGui::GetStyle().ScrollbarSize;

		auto textScrollbarTop = std::max(outerRect.Min.y + borderSize, outerRect.Max.y - borderSize - scrollbarSize);
		ImRect leftScrollbarFrame(leftTextPos, textScrollbarTop, rightLineNumberPos, textScrollbarTop + scrollbarSize);
		ImRect rightScrollbarFrame(rightTextPos, textScrollbarTop, rightTextEnd, textScrollbarTop + scrollbarSize);
		ImS64 scroll = static_cast<ImS64>(textScroll);

		if (ImGui::ScrollbarEx(
			leftScrollbarFrame,
			ImGui::GetID("leftTextScroll"),
			ImGuiAxis_X,
			&scroll,
			static_cast<ImS64>(visibleColumnsWidth),
			static_cast<ImS64>(maxColumnsWidth),
			ImDrawFlags_RoundCornersAll)) { textScroll = static_cast<float>(scroll); }

		if (ImGui::ScrollbarEx(
			rightScrollbarFrame,
			ImGui::GetID("rightTextScroll"),
			ImGuiAxis_X,
			&scroll,
			static_cast<ImS64>(visibleColumnsWidth),
			static_cast<ImS64>(maxColumnsWidth),
			ImDrawFlags_RoundCornersAll)) { textScroll = static_cast<float>(scroll); }

		if (ImGui::IsWindowHovered()) {
			textScroll = std::clamp(textScroll - ImGui::GetIO().MouseWheelH * ImGui::GetFontSize(), 0.0f, maxColumnsWidth - visibleColumnsWidth);

			if (ImGui::IsKeyPressed(ImGuiKey_LeftArrow)) {
				textScroll = std::max(textScroll - glyphSize.x, 0.0f);

			} else if (ImGui::IsKeyPressed(ImGuiKey_RightArrow)) {
				textScroll = std::min(textScroll + glyphSize.x, maxColumnsWidth - visibleColumnsWidth);

			} else if (ImGui::IsKeyPressed(ImGuiKey_Home)) {
				textScroll = 0.0f;

			} else if (ImGui::IsKeyPressed(ImGuiKey_End)) {
				textScroll = maxColumnsWidth - visibleColumnsWidth;
			}
		}
	}

	if (ImGui::IsWindowHovered()) {
		if (ImGui::IsKeyPressed(ImGuiKey_UpArrow)) {
			if (ImGui::IsKeyDown(ImGuiMod_Ctrl)) {
				ImGui::SetScrollY(0.0f);

			} else {
				ImGui::SetScrollY(std::max(ImGui::GetScrollY() - glyphSize.y, 0.0f));
			}

		} else if (ImGui::IsKeyPressed(ImGuiKey_DownArrow)) {
			if (ImGui::IsKeyDown(ImGuiMod_Ctrl)) {
				ImGui::SetScrollY(ImGui::GetScrollMaxY());

			} else {
				ImGui::SetScrollY(std::min(ImGui::GetScrollY() + glyphSize.y, ImGui::GetScrollMaxY()));
			}

		} else if (ImGui::IsKeyPressed(ImGuiKey_PageUp)) {
			ImGui::SetScrollY(std::max(ImGui::GetScrollY() - (visibleLines - 2) * glyphSize.y, 0.0f));

		} else if (ImGui::IsKeyPressed(ImGuiKey_PageDown)) {
			ImGui::SetScrollY(std::min(ImGui::GetScrollY() + (visibleLines - 2) * glyphSize.y, ImGui::GetScrollMaxY()));
		}
	}
}


//
//	TextDiff::renderSideBySideMiniMap
//

void TextDiff::renderSideBySideMiniMap() {
	// based on https://github.com/ocornut/imgui/issues/3114
	if (showScrollbarMiniMap) {
		auto window = ImGui::GetCurrentWindow();

		if (window->ScrollbarY) {
			auto drawList = ImGui::GetWindowDrawList();
			auto rect = ImGui::GetWindowScrollbarRect(window, ImGuiAxis_Y);
			auto lineHeight = rect.GetHeight() / static_cast<float>(document.size());
			auto offset = (rect.Max.x - rect.Min.x) * 0.3f;
			auto left = rect.Min.x + offset;
			auto right = rect.Max.x - offset;

			drawList->PushClipRect(rect.Min, rect.Max, false);

			// render diff locations
			for (size_t i = 0; i < lineInfo.size(); i++) {
				auto& line = lineInfo[i];

				if (line.status != LineStatus::common) {
					auto color = (line.status == LineStatus::added) ? addedColor : deletedColor;
					auto ly = std::round(rect.Min.y + i * lineHeight);
					drawList->AddRectFilled(ImVec2(left, ly), ImVec2(right, ly + lineHeight), color);
				}
			}

			drawList->PopClipRect();
		}
	}
}
