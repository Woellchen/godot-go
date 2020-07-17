package godot

import (
	"github.com/Woellchen/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewVisualScriptComposeArrayFromPointer(ptr gdnative.Pointer) VisualScriptComposeArray {
func newVisualScriptComposeArrayFromPointer(ptr gdnative.Pointer) VisualScriptComposeArray {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptComposeArray{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptComposeArray struct {
	VisualScriptLists
	owner gdnative.Object
}

func (o *VisualScriptComposeArray) BaseClass() string {
	return "VisualScriptComposeArray"
}

// VisualScriptComposeArrayImplementer is an interface that implements the methods
// of the VisualScriptComposeArray class.
type VisualScriptComposeArrayImplementer interface {
	VisualScriptListsImplementer
}
