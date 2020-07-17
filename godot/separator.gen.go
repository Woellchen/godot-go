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

//func NewSeparatorFromPointer(ptr gdnative.Pointer) Separator {
func newSeparatorFromPointer(ptr gdnative.Pointer) Separator {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Separator{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Separator is a [Control] used for separating other controls. It's purely a visual decoration. Horizontal ([HSeparator]) and Vertical ([VSeparator]) versions are available.
*/
type Separator struct {
	Control
	owner gdnative.Object
}

func (o *Separator) BaseClass() string {
	return "Separator"
}

// SeparatorImplementer is an interface that implements the methods
// of the Separator class.
type SeparatorImplementer interface {
	ControlImplementer
}
