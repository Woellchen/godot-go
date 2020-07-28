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

//func NewStreamTextureFromPointer(ptr gdnative.Pointer) StreamTexture {
func newStreamTextureFromPointer(ptr gdnative.Pointer) StreamTexture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := StreamTexture{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A texture that is loaded from a [code].stex[/code] file.
*/
type StreamTexture struct {
	Texture
	owner gdnative.Object
}

func (o *StreamTexture) BaseClass() string {
	return "StreamTexture"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *StreamTexture) GetLoadPath() gdnative.String {
	// log.Println("Calling StreamTexture.GetLoadPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StreamTexture", "get_load_path")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Loads the texture from the given path.
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *StreamTexture) Load(path gdnative.String) gdnative.Error {
	// log.Println("Calling StreamTexture.Load()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StreamTexture", "load")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

// StreamTextureImplementer is an interface that implements the methods
// of the StreamTexture class.
type StreamTextureImplementer interface {
	TextureImplementer
	GetLoadPath() gdnative.String
}
