package class

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewVSliderFromPointer(ptr gdnative.Pointer) VSlider {
func NewVSliderFromPointer(ptr gdnative.Pointer) VSlider {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VSlider{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Vertical slider. See [Slider]. This one goes from left (min) to right (max).
*/
type VSlider struct {
	Slider
	owner gdnative.Object
}

func (o *VSlider) BaseClass() string {
	return "VSlider"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VSlider) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *VSlider) GetBaseObject() gdnative.Object {
	return o.owner
}