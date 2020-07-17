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

//func NewVisualShaderNodeCubeMapUniformFromPointer(ptr gdnative.Pointer) VisualShaderNodeCubeMapUniform {
func newVisualShaderNodeCubeMapUniformFromPointer(ptr gdnative.Pointer) VisualShaderNodeCubeMapUniform {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeCubeMapUniform{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Translated to [code]uniform samplerCube[/code] in the shader language. The output value can be used as port for [VisualShaderNodeCubeMap].
*/
type VisualShaderNodeCubeMapUniform struct {
	VisualShaderNodeTextureUniform
	owner gdnative.Object
}

func (o *VisualShaderNodeCubeMapUniform) BaseClass() string {
	return "VisualShaderNodeCubeMapUniform"
}

// VisualShaderNodeCubeMapUniformImplementer is an interface that implements the methods
// of the VisualShaderNodeCubeMapUniform class.
type VisualShaderNodeCubeMapUniformImplementer interface {
	VisualShaderNodeTextureUniformImplementer
}
