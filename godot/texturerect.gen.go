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

// TextureRectStretchMode is an enum for StretchMode values.
type TextureRectStretchMode int

const (
	TextureRectStretchKeep               TextureRectStretchMode = 3
	TextureRectStretchKeepAspect         TextureRectStretchMode = 5
	TextureRectStretchKeepAspectCentered TextureRectStretchMode = 6
	TextureRectStretchKeepAspectCovered  TextureRectStretchMode = 7
	TextureRectStretchKeepCentered       TextureRectStretchMode = 4
	TextureRectStretchScale              TextureRectStretchMode = 1
	TextureRectStretchScaleOnExpand      TextureRectStretchMode = 0
	TextureRectStretchTile               TextureRectStretchMode = 2
)

//func NewTextureRectFromPointer(ptr gdnative.Pointer) TextureRect {
func newTextureRectFromPointer(ptr gdnative.Pointer) TextureRect {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := TextureRect{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Used to draw icons and sprites in a user interface. The texture's placement can be controlled with the [member stretch_mode] property. It can scale, tile, or stay centered inside its bounding rectangle.
*/
type TextureRect struct {
	Control
	owner gdnative.Object
}

func (o *TextureRect) BaseClass() string {
	return "TextureRect"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *TextureRect) X_TextureChanged() {
	// log.Println("Calling TextureRect.X_TextureChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureRect", "_texture_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: enum.TextureRect::StretchMode
*/
func (o *TextureRect) GetStretchMode() TextureRectStretchMode {
	// log.Println("Calling TextureRect.GetStretchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureRect", "get_stretch_mode")

	// Call the parent method.
	// enum.TextureRect::StretchMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return TextureRectStretchMode(ret)
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *TextureRect) GetTexture() TextureImplementer {
	// log.Println("Calling TextureRect.GetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureRect", "get_texture")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *TextureRect) HasExpand() gdnative.Bool {
	// log.Println("Calling TextureRect.HasExpand()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureRect", "has_expand")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *TextureRect) IsFlippedH() gdnative.Bool {
	// log.Println("Calling TextureRect.IsFlippedH()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureRect", "is_flipped_h")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *TextureRect) IsFlippedV() gdnative.Bool {
	// log.Println("Calling TextureRect.IsFlippedV()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureRect", "is_flipped_v")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *TextureRect) SetExpand(enable gdnative.Bool) {
	// log.Println("Calling TextureRect.SetExpand()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureRect", "set_expand")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *TextureRect) SetFlipH(enable gdnative.Bool) {
	// log.Println("Calling TextureRect.SetFlipH()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureRect", "set_flip_h")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *TextureRect) SetFlipV(enable gdnative.Bool) {
	// log.Println("Calling TextureRect.SetFlipV()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureRect", "set_flip_v")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false stretch_mode int}], Returns: void
*/
func (o *TextureRect) SetStretchMode(stretchMode gdnative.Int) {
	// log.Println("Calling TextureRect.SetStretchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(stretchMode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureRect", "set_stretch_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *TextureRect) SetTexture(texture TextureImplementer) {
	// log.Println("Calling TextureRect.SetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureRect", "set_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// TextureRectImplementer is an interface that implements the methods
// of the TextureRect class.
type TextureRectImplementer interface {
	ControlImplementer
	X_TextureChanged()
	GetTexture() TextureImplementer
	HasExpand() gdnative.Bool
	IsFlippedH() gdnative.Bool
	IsFlippedV() gdnative.Bool
	SetExpand(enable gdnative.Bool)
	SetFlipH(enable gdnative.Bool)
	SetFlipV(enable gdnative.Bool)
	SetStretchMode(stretchMode gdnative.Int)
	SetTexture(texture TextureImplementer)
}
