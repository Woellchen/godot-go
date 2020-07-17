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

//func NewGradientTextureFromPointer(ptr gdnative.Pointer) GradientTexture {
func newGradientTextureFromPointer(ptr gdnative.Pointer) GradientTexture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := GradientTexture{}
	obj.SetBaseObject(owner)

	return obj
}

/*
GradientTexture uses a [Gradient] to fill the texture data. The gradient will be filled from left to right using colors obtained from the gradient. This means the texture does not necessarily represent an exact copy of the gradient, but instead an interpolation of samples obtained from the gradient at fixed steps (see [member width]).
*/
type GradientTexture struct {
	Texture
	owner gdnative.Object
}

func (o *GradientTexture) BaseClass() string {
	return "GradientTexture"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *GradientTexture) X_Update() {
	//log.Println("Calling GradientTexture.X_Update()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GradientTexture", "_update")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Gradient
*/
func (o *GradientTexture) GetGradient() GradientImplementer {
	//log.Println("Calling GradientTexture.GetGradient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GradientTexture", "get_gradient")

	// Call the parent method.
	// Gradient
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newGradientFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(GradientImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Gradient" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(GradientImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false gradient Gradient}], Returns: void
*/
func (o *GradientTexture) SetGradient(gradient GradientImplementer) {
	//log.Println("Calling GradientTexture.SetGradient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(gradient.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GradientTexture", "set_gradient")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false width int}], Returns: void
*/
func (o *GradientTexture) SetWidth(width gdnative.Int) {
	//log.Println("Calling GradientTexture.SetWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(width)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GradientTexture", "set_width")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// GradientTextureImplementer is an interface that implements the methods
// of the GradientTexture class.
type GradientTextureImplementer interface {
	TextureImplementer
	X_Update()
	GetGradient() GradientImplementer
	SetGradient(gradient GradientImplementer)
	SetWidth(width gdnative.Int)
}
