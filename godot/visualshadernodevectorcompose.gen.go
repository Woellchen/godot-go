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

//func NewVisualShaderNodeVectorComposeFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorCompose {
func newVisualShaderNodeVectorComposeFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorCompose {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeVectorCompose{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Creates a [code]vec3[/code] using three scalar values that can be provided from separate inputs.
*/
type VisualShaderNodeVectorCompose struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeVectorCompose) BaseClass() string {
	return "VisualShaderNodeVectorCompose"
}

// VisualShaderNodeVectorComposeImplementer is an interface that implements the methods
// of the VisualShaderNodeVectorCompose class.
type VisualShaderNodeVectorComposeImplementer interface {
	VisualShaderNodeImplementer
}
