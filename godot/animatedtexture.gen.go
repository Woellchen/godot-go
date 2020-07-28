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

//func NewAnimatedTextureFromPointer(ptr gdnative.Pointer) AnimatedTexture {
func newAnimatedTextureFromPointer(ptr gdnative.Pointer) AnimatedTexture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimatedTexture{}
	obj.SetBaseObject(owner)

	return obj
}

/*
[AnimatedTexture] is a resource format for frame-based animations, where multiple textures can be chained automatically with a predefined delay for each frame. Unlike [AnimationPlayer] or [AnimatedSprite], it isn't a [Node], but has the advantage of being usable anywhere a [Texture] resource can be used, e.g. in a [TileSet]. The playback of the animation is controlled by the [member fps] property as well as each frame's optional delay (see [method set_frame_delay]). The animation loops, i.e. it will restart at frame 0 automatically after playing the last frame. [AnimatedTexture] currently requires all frame textures to have the same size, otherwise the bigger ones will be cropped to match the smallest one. Also, it doesn't support [AtlasTexture]. Each frame needs to be separate image.
*/
type AnimatedTexture struct {
	Texture
	owner gdnative.Object
}

func (o *AnimatedTexture) BaseClass() string {
	return "AnimatedTexture"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AnimatedTexture) X_UpdateProxy() {
	// log.Println("Calling AnimatedTexture.X_UpdateProxy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "_update_proxy")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *AnimatedTexture) GetCurrentFrame() gdnative.Int {
	// log.Println("Calling AnimatedTexture.GetCurrentFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "get_current_frame")

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
	Args: [], Returns: float
*/
func (o *AnimatedTexture) GetFps() gdnative.Real {
	// log.Println("Calling AnimatedTexture.GetFps()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "get_fps")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the given frame's delay value.
	Args: [{ false frame int}], Returns: float
*/
func (o *AnimatedTexture) GetFrameDelay(frame gdnative.Int) gdnative.Real {
	// log.Println("Calling AnimatedTexture.GetFrameDelay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(frame)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "get_frame_delay")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the given frame's [Texture].
	Args: [{ false frame int}], Returns: Texture
*/
func (o *AnimatedTexture) GetFrameTexture(frame gdnative.Int) TextureImplementer {
	// log.Println("Calling AnimatedTexture.GetFrameTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(frame)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "get_frame_texture")

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
	Args: [], Returns: int
*/
func (o *AnimatedTexture) GetFrames() gdnative.Int {
	// log.Println("Calling AnimatedTexture.GetFrames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "get_frames")

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
func (o *AnimatedTexture) GetOneshot() gdnative.Bool {
	// log.Println("Calling AnimatedTexture.GetOneshot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "get_oneshot")

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
func (o *AnimatedTexture) GetPause() gdnative.Bool {
	// log.Println("Calling AnimatedTexture.GetPause()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "get_pause")

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
	Args: [{ false frame int}], Returns: void
*/
func (o *AnimatedTexture) SetCurrentFrame(frame gdnative.Int) {
	// log.Println("Calling AnimatedTexture.SetCurrentFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(frame)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "set_current_frame")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false fps float}], Returns: void
*/
func (o *AnimatedTexture) SetFps(fps gdnative.Real) {
	// log.Println("Calling AnimatedTexture.SetFps()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(fps)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "set_fps")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets an additional delay (in seconds) between this frame and the next one, that will be added to the time interval defined by [member fps]. By default, frames have no delay defined. If a delay value is defined, the final time interval between this frame and the next will be [code]1.0 / fps + delay[/code]. For example, for an animation with 3 frames, 2 FPS and a frame delay on the second frame of 1.2, the resulting playback will be: [codeblock] Frame 0: 0.5 s (1 / fps) Frame 1: 1.7 s (1 / fps + 1.2) Frame 2: 0.5 s (1 / fps) Total duration: 2.7 s [/codeblock]
	Args: [{ false frame int} { false delay float}], Returns: void
*/
func (o *AnimatedTexture) SetFrameDelay(frame gdnative.Int, delay gdnative.Real) {
	// log.Println("Calling AnimatedTexture.SetFrameDelay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(frame)
	ptrArguments[1] = gdnative.NewPointerFromReal(delay)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "set_frame_delay")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Assigns a [Texture] to the given frame. Frame IDs start at 0, so the first frame has ID 0, and the last frame of the animation has ID [member frames] - 1. You can define any number of textures up to [constant MAX_FRAMES], but keep in mind that only frames from 0 to [member frames] - 1 will be part of the animation.
	Args: [{ false frame int} { false texture Texture}], Returns: void
*/
func (o *AnimatedTexture) SetFrameTexture(frame gdnative.Int, texture TextureImplementer) {
	// log.Println("Calling AnimatedTexture.SetFrameTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(frame)
	ptrArguments[1] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "set_frame_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false frames int}], Returns: void
*/
func (o *AnimatedTexture) SetFrames(frames gdnative.Int) {
	// log.Println("Calling AnimatedTexture.SetFrames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(frames)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "set_frames")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false oneshot bool}], Returns: void
*/
func (o *AnimatedTexture) SetOneshot(oneshot gdnative.Bool) {
	// log.Println("Calling AnimatedTexture.SetOneshot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(oneshot)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "set_oneshot")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pause bool}], Returns: void
*/
func (o *AnimatedTexture) SetPause(pause gdnative.Bool) {
	// log.Println("Calling AnimatedTexture.SetPause()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(pause)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedTexture", "set_pause")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimatedTextureImplementer is an interface that implements the methods
// of the AnimatedTexture class.
type AnimatedTextureImplementer interface {
	TextureImplementer
	X_UpdateProxy()
	GetCurrentFrame() gdnative.Int
	GetFps() gdnative.Real
	GetFrameDelay(frame gdnative.Int) gdnative.Real
	GetFrameTexture(frame gdnative.Int) TextureImplementer
	GetFrames() gdnative.Int
	GetOneshot() gdnative.Bool
	GetPause() gdnative.Bool
	SetCurrentFrame(frame gdnative.Int)
	SetFps(fps gdnative.Real)
	SetFrameDelay(frame gdnative.Int, delay gdnative.Real)
	SetFrameTexture(frame gdnative.Int, texture TextureImplementer)
	SetFrames(frames gdnative.Int)
	SetOneshot(oneshot gdnative.Bool)
	SetPause(pause gdnative.Bool)
}
