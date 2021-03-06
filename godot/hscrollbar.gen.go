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

//func NewHScrollBarFromPointer(ptr gdnative.Pointer) HScrollBar {
func newHScrollBarFromPointer(ptr gdnative.Pointer) HScrollBar {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HScrollBar{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Horizontal version of [ScrollBar], which goes from left (min) to right (max).
*/
type HScrollBar struct {
	ScrollBar
	owner gdnative.Object
}

func (o *HScrollBar) BaseClass() string {
	return "HScrollBar"
}

// HScrollBarImplementer is an interface that implements the methods
// of the HScrollBar class.
type HScrollBarImplementer interface {
	ScrollBarImplementer
}
