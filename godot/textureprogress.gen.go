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

// TextureProgressFillMode is an enum for FillMode values.
type TextureProgressFillMode int

const (
	TextureProgressFillBilinearLeftAndRight         TextureProgressFillMode = 6
	TextureProgressFillBilinearTopAndBottom         TextureProgressFillMode = 7
	TextureProgressFillBottomToTop                  TextureProgressFillMode = 3
	TextureProgressFillClockwise                    TextureProgressFillMode = 4
	TextureProgressFillClockwiseAndCounterClockwise TextureProgressFillMode = 8
	TextureProgressFillCounterClockwise             TextureProgressFillMode = 5
	TextureProgressFillLeftToRight                  TextureProgressFillMode = 0
	TextureProgressFillRightToLeft                  TextureProgressFillMode = 1
	TextureProgressFillTopToBottom                  TextureProgressFillMode = 2
)

//func NewTextureProgressFromPointer(ptr gdnative.Pointer) TextureProgress {
func newTextureProgressFromPointer(ptr gdnative.Pointer) TextureProgress {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := TextureProgress{}
	obj.SetBaseObject(owner)

	return obj
}

/*
TextureProgress works like [ProgressBar], but uses up to 3 textures instead of Godot's [Theme] resource. It can be used to create horizontal, vertical and radial progress bars.
*/
type TextureProgress struct {
	Range
	owner gdnative.Object
}

func (o *TextureProgress) BaseClass() string {
	return "TextureProgress"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *TextureProgress) GetFillDegrees() gdnative.Real {
	// log.Println("Calling TextureProgress.GetFillDegrees()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_fill_degrees")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *TextureProgress) GetFillMode() gdnative.Int {
	// log.Println("Calling TextureProgress.GetFillMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_fill_mode")

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
	Args: [], Returns: bool
*/
func (o *TextureProgress) GetNinePatchStretch() gdnative.Bool {
	// log.Println("Calling TextureProgress.GetNinePatchStretch()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_nine_patch_stretch")

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
	Args: [], Returns: Texture
*/
func (o *TextureProgress) GetOverTexture() TextureImplementer {
	// log.Println("Calling TextureProgress.GetOverTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_over_texture")

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
	Args: [], Returns: Texture
*/
func (o *TextureProgress) GetProgressTexture() TextureImplementer {
	// log.Println("Calling TextureProgress.GetProgressTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_progress_texture")

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
	Args: [], Returns: Vector2
*/
func (o *TextureProgress) GetRadialCenterOffset() gdnative.Vector2 {
	// log.Println("Calling TextureProgress.GetRadialCenterOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_radial_center_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *TextureProgress) GetRadialInitialAngle() gdnative.Real {
	// log.Println("Calling TextureProgress.GetRadialInitialAngle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_radial_initial_angle")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false margin int}], Returns: int
*/
func (o *TextureProgress) GetStretchMargin(margin gdnative.Int) gdnative.Int {
	// log.Println("Calling TextureProgress.GetStretchMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_stretch_margin")

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
	Args: [], Returns: Color
*/
func (o *TextureProgress) GetTintOver() gdnative.Color {
	// log.Println("Calling TextureProgress.GetTintOver()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_tint_over")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *TextureProgress) GetTintProgress() gdnative.Color {
	// log.Println("Calling TextureProgress.GetTintProgress()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_tint_progress")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *TextureProgress) GetTintUnder() gdnative.Color {
	// log.Println("Calling TextureProgress.GetTintUnder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_tint_under")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *TextureProgress) GetUnderTexture() TextureImplementer {
	// log.Println("Calling TextureProgress.GetUnderTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "get_under_texture")

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
	Args: [{ false mode float}], Returns: void
*/
func (o *TextureProgress) SetFillDegrees(mode gdnative.Real) {
	// log.Println("Calling TextureProgress.SetFillDegrees()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_fill_degrees")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *TextureProgress) SetFillMode(mode gdnative.Int) {
	// log.Println("Calling TextureProgress.SetFillMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_fill_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false stretch bool}], Returns: void
*/
func (o *TextureProgress) SetNinePatchStretch(stretch gdnative.Bool) {
	// log.Println("Calling TextureProgress.SetNinePatchStretch()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(stretch)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_nine_patch_stretch")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false tex Texture}], Returns: void
*/
func (o *TextureProgress) SetOverTexture(tex TextureImplementer) {
	// log.Println("Calling TextureProgress.SetOverTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(tex.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_over_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false tex Texture}], Returns: void
*/
func (o *TextureProgress) SetProgressTexture(tex TextureImplementer) {
	// log.Println("Calling TextureProgress.SetProgressTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(tex.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_progress_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode Vector2}], Returns: void
*/
func (o *TextureProgress) SetRadialCenterOffset(mode gdnative.Vector2) {
	// log.Println("Calling TextureProgress.SetRadialCenterOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_radial_center_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode float}], Returns: void
*/
func (o *TextureProgress) SetRadialInitialAngle(mode gdnative.Real) {
	// log.Println("Calling TextureProgress.SetRadialInitialAngle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_radial_initial_angle")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false margin int} { false value int}], Returns: void
*/
func (o *TextureProgress) SetStretchMargin(margin gdnative.Int, value gdnative.Int) {
	// log.Println("Calling TextureProgress.SetStretchMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)
	ptrArguments[1] = gdnative.NewPointerFromInt(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_stretch_margin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false tint Color}], Returns: void
*/
func (o *TextureProgress) SetTintOver(tint gdnative.Color) {
	// log.Println("Calling TextureProgress.SetTintOver()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(tint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_tint_over")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false tint Color}], Returns: void
*/
func (o *TextureProgress) SetTintProgress(tint gdnative.Color) {
	// log.Println("Calling TextureProgress.SetTintProgress()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(tint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_tint_progress")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false tint Color}], Returns: void
*/
func (o *TextureProgress) SetTintUnder(tint gdnative.Color) {
	// log.Println("Calling TextureProgress.SetTintUnder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(tint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_tint_under")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false tex Texture}], Returns: void
*/
func (o *TextureProgress) SetUnderTexture(tex TextureImplementer) {
	// log.Println("Calling TextureProgress.SetUnderTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(tex.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureProgress", "set_under_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// TextureProgressImplementer is an interface that implements the methods
// of the TextureProgress class.
type TextureProgressImplementer interface {
	RangeImplementer
	GetFillDegrees() gdnative.Real
	GetFillMode() gdnative.Int
	GetNinePatchStretch() gdnative.Bool
	GetOverTexture() TextureImplementer
	GetProgressTexture() TextureImplementer
	GetRadialCenterOffset() gdnative.Vector2
	GetRadialInitialAngle() gdnative.Real
	GetStretchMargin(margin gdnative.Int) gdnative.Int
	GetTintOver() gdnative.Color
	GetTintProgress() gdnative.Color
	GetTintUnder() gdnative.Color
	GetUnderTexture() TextureImplementer
	SetFillDegrees(mode gdnative.Real)
	SetFillMode(mode gdnative.Int)
	SetNinePatchStretch(stretch gdnative.Bool)
	SetOverTexture(tex TextureImplementer)
	SetProgressTexture(tex TextureImplementer)
	SetRadialCenterOffset(mode gdnative.Vector2)
	SetRadialInitialAngle(mode gdnative.Real)
	SetStretchMargin(margin gdnative.Int, value gdnative.Int)
	SetTintOver(tint gdnative.Color)
	SetTintProgress(tint gdnative.Color)
	SetTintUnder(tint gdnative.Color)
	SetUnderTexture(tex TextureImplementer)
}
