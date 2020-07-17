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

//func NewExternalTextureFromPointer(ptr gdnative.Pointer) ExternalTexture {
func newExternalTextureFromPointer(ptr gdnative.Pointer) ExternalTexture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ExternalTexture{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Enable support for the OpenGL ES external texture extension as defined by [url=https://www.khronos.org/registry/OpenGL/extensions/OES/OES_EGL_image_external.txt]OES_EGL_image_external[/url]. [b]Note:[/b] This is only supported for Android platforms.
*/
type ExternalTexture struct {
	Texture
	owner gdnative.Object
}

func (o *ExternalTexture) BaseClass() string {
	return "ExternalTexture"
}

/*
        Returns the external texture name.
	Args: [], Returns: int
*/
func (o *ExternalTexture) GetExternalTextureId() gdnative.Int {
	//log.Println("Calling ExternalTexture.GetExternalTextureId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ExternalTexture", "get_external_texture_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false size Vector2}], Returns: void
*/
func (o *ExternalTexture) SetSize(size gdnative.Vector2) {
	//log.Println("Calling ExternalTexture.SetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ExternalTexture", "set_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ExternalTextureImplementer is an interface that implements the methods
// of the ExternalTexture class.
type ExternalTextureImplementer interface {
	TextureImplementer
	GetExternalTextureId() gdnative.Int
	SetSize(size gdnative.Vector2)
}