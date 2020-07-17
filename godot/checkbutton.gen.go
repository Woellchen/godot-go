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

//func NewCheckButtonFromPointer(ptr gdnative.Pointer) CheckButton {
func newCheckButtonFromPointer(ptr gdnative.Pointer) CheckButton {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CheckButton{}
	obj.SetBaseObject(owner)

	return obj
}

/*
CheckButton is a toggle button displayed as a check field. It's similar to [CheckBox] in functionality, but it has a different apperance. To follow established UX patterns, it's recommended to use CheckButton when toggling it has an [b]immediate[/b] effect on something. For instance, it should be used if toggling it enables/disables a setting without requiring the user to press a confirmation button.
*/
type CheckButton struct {
	Button
	owner gdnative.Object
}

func (o *CheckButton) BaseClass() string {
	return "CheckButton"
}

// CheckButtonImplementer is an interface that implements the methods
// of the CheckButton class.
type CheckButtonImplementer interface {
	ButtonImplementer
}
