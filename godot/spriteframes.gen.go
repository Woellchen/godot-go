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

//func NewSpriteFramesFromPointer(ptr gdnative.Pointer) SpriteFrames {
func newSpriteFramesFromPointer(ptr gdnative.Pointer) SpriteFrames {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SpriteFrames{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Sprite frame library for [AnimatedSprite]. Contains frames and animation data for playback.
*/
type SpriteFrames struct {
	Resource
	owner gdnative.Object
}

func (o *SpriteFrames) BaseClass() string {
	return "SpriteFrames"
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *SpriteFrames) X_GetAnimations() gdnative.Array {
	// log.Println("Calling SpriteFrames.X_GetAnimations()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "_get_animations")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *SpriteFrames) X_GetFrames() gdnative.Array {
	// log.Println("Calling SpriteFrames.X_GetFrames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "_get_frames")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false arg0 Array}], Returns: void
*/
func (o *SpriteFrames) X_SetAnimations(arg0 gdnative.Array) {
	// log.Println("Calling SpriteFrames.X_SetAnimations()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "_set_animations")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 Array}], Returns: void
*/
func (o *SpriteFrames) X_SetFrames(arg0 gdnative.Array) {
	// log.Println("Calling SpriteFrames.X_SetFrames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "_set_frames")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a new animation to the library.
	Args: [{ false anim String}], Returns: void
*/
func (o *SpriteFrames) AddAnimation(anim gdnative.String) {
	// log.Println("Calling SpriteFrames.AddAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "add_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a frame to the given animation.
	Args: [{ false anim String} { false frame Texture} {-1 true at_position int}], Returns: void
*/
func (o *SpriteFrames) AddFrame(anim gdnative.String, frame TextureImplementer, atPosition gdnative.Int) {
	// log.Println("Calling SpriteFrames.AddFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)
	ptrArguments[1] = gdnative.NewPointerFromObject(frame.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromInt(atPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "add_frame")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all frames from the given animation.
	Args: [{ false anim String}], Returns: void
*/
func (o *SpriteFrames) Clear(anim gdnative.String) {
	// log.Println("Calling SpriteFrames.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all animations. A "default" animation will be created.
	Args: [], Returns: void
*/
func (o *SpriteFrames) ClearAll() {
	// log.Println("Calling SpriteFrames.ClearAll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "clear_all")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If [code]true[/code], the given animation will loop.
	Args: [{ false anim String}], Returns: bool
*/
func (o *SpriteFrames) GetAnimationLoop(anim gdnative.String) gdnative.Bool {
	// log.Println("Calling SpriteFrames.GetAnimationLoop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "get_animation_loop")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns an array containing the names associated to each animation. Values are placed in alphabetical order.
	Args: [], Returns: PoolStringArray
*/
func (o *SpriteFrames) GetAnimationNames() gdnative.PoolStringArray {
	// log.Println("Calling SpriteFrames.GetAnimationNames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "get_animation_names")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        The animation's speed in frames per second.
	Args: [{ false anim String}], Returns: float
*/
func (o *SpriteFrames) GetAnimationSpeed(anim gdnative.String) gdnative.Real {
	// log.Println("Calling SpriteFrames.GetAnimationSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "get_animation_speed")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the animation's selected frame.
	Args: [{ false anim String} { false idx int}], Returns: Texture
*/
func (o *SpriteFrames) GetFrame(anim gdnative.String, idx gdnative.Int) TextureImplementer {
	// log.Println("Calling SpriteFrames.GetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)
	ptrArguments[1] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "get_frame")

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
        Returns the number of frames in the animation.
	Args: [{ false anim String}], Returns: int
*/
func (o *SpriteFrames) GetFrameCount(anim gdnative.String) gdnative.Int {
	// log.Println("Calling SpriteFrames.GetFrameCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "get_frame_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        If [code]true[/code], the named animation exists.
	Args: [{ false anim String}], Returns: bool
*/
func (o *SpriteFrames) HasAnimation(anim gdnative.String) gdnative.Bool {
	// log.Println("Calling SpriteFrames.HasAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "has_animation")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes the given animation.
	Args: [{ false anim String}], Returns: void
*/
func (o *SpriteFrames) RemoveAnimation(anim gdnative.String) {
	// log.Println("Calling SpriteFrames.RemoveAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "remove_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes the animation's selected frame.
	Args: [{ false anim String} { false idx int}], Returns: void
*/
func (o *SpriteFrames) RemoveFrame(anim gdnative.String, idx gdnative.Int) {
	// log.Println("Calling SpriteFrames.RemoveFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)
	ptrArguments[1] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "remove_frame")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Changes the animation's name to [code]newname[/code].
	Args: [{ false anim String} { false newname String}], Returns: void
*/
func (o *SpriteFrames) RenameAnimation(anim gdnative.String, newname gdnative.String) {
	// log.Println("Calling SpriteFrames.RenameAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)
	ptrArguments[1] = gdnative.NewPointerFromString(newname)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "rename_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If [code]true[/code], the animation will loop.
	Args: [{ false anim String} { false loop bool}], Returns: void
*/
func (o *SpriteFrames) SetAnimationLoop(anim gdnative.String, loop gdnative.Bool) {
	// log.Println("Calling SpriteFrames.SetAnimationLoop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)
	ptrArguments[1] = gdnative.NewPointerFromBool(loop)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "set_animation_loop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        The animation's speed in frames per second.
	Args: [{ false anim String} { false speed float}], Returns: void
*/
func (o *SpriteFrames) SetAnimationSpeed(anim gdnative.String, speed gdnative.Real) {
	// log.Println("Calling SpriteFrames.SetAnimationSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)
	ptrArguments[1] = gdnative.NewPointerFromReal(speed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "set_animation_speed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the texture of the given frame.
	Args: [{ false anim String} { false idx int} { false txt Texture}], Returns: void
*/
func (o *SpriteFrames) SetFrame(anim gdnative.String, idx gdnative.Int, txt TextureImplementer) {
	// log.Println("Calling SpriteFrames.SetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)
	ptrArguments[1] = gdnative.NewPointerFromInt(idx)
	ptrArguments[2] = gdnative.NewPointerFromObject(txt.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteFrames", "set_frame")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// SpriteFramesImplementer is an interface that implements the methods
// of the SpriteFrames class.
type SpriteFramesImplementer interface {
	ResourceImplementer
	X_GetAnimations() gdnative.Array
	X_GetFrames() gdnative.Array
	X_SetAnimations(arg0 gdnative.Array)
	X_SetFrames(arg0 gdnative.Array)
	AddAnimation(anim gdnative.String)
	AddFrame(anim gdnative.String, frame TextureImplementer, atPosition gdnative.Int)
	Clear(anim gdnative.String)
	ClearAll()
	GetAnimationLoop(anim gdnative.String) gdnative.Bool
	GetAnimationNames() gdnative.PoolStringArray
	GetAnimationSpeed(anim gdnative.String) gdnative.Real
	GetFrame(anim gdnative.String, idx gdnative.Int) TextureImplementer
	GetFrameCount(anim gdnative.String) gdnative.Int
	HasAnimation(anim gdnative.String) gdnative.Bool
	RemoveAnimation(anim gdnative.String)
	RemoveFrame(anim gdnative.String, idx gdnative.Int)
	RenameAnimation(anim gdnative.String, newname gdnative.String)
	SetAnimationLoop(anim gdnative.String, loop gdnative.Bool)
	SetAnimationSpeed(anim gdnative.String, speed gdnative.Real)
	SetFrame(anim gdnative.String, idx gdnative.Int, txt TextureImplementer)
}
