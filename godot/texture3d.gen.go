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

//func NewTexture3DFromPointer(ptr gdnative.Pointer) Texture3D {
func newTexture3DFromPointer(ptr gdnative.Pointer) Texture3D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Texture3D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Texture3D is a 3-dimensional texture that has a width, height, and depth.
*/
type Texture3D struct {
	TextureLayered
	owner gdnative.Object
}

func (o *Texture3D) BaseClass() string {
	return "Texture3D"
}

// Texture3DImplementer is an interface that implements the methods
// of the Texture3D class.
type Texture3DImplementer interface {
	TextureLayeredImplementer
}
