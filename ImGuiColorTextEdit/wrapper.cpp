// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

#include "wrapper.h"
#include "../cwrappers/cimCTE.h"

void wrap_TextEditor_Redo(TextEditor* self) { TextEditor_Redo(self,1); }
bool wrap_TextEditor_Render(TextEditor* self,const char* aTitle) { return TextEditor_Render(self,aTitle,false,(ImVec2){},false); }
void wrap_TextEditor_SelectAllOccurrencesOf(TextEditor* self,const char* aText,int aTextSize) { TextEditor_SelectAllOccurrencesOf(self,aText,aTextSize,true); }
void wrap_TextEditor_SelectNextOccurrenceOf(TextEditor* self,const char* aText,int aTextSize) { TextEditor_SelectNextOccurrenceOf(self,aText,aTextSize,true); }
void wrap_TextEditor_Undo(TextEditor* self) { TextEditor_Undo(self,1); }
