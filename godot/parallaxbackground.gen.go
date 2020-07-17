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

//func NewParallaxBackgroundFromPointer(ptr gdnative.Pointer) ParallaxBackground {
func newParallaxBackgroundFromPointer(ptr gdnative.Pointer) ParallaxBackground {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ParallaxBackground{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A ParallaxBackground uses one or more [ParallaxLayer] child nodes to create a parallax effect. Each [ParallaxLayer] can move at a different speed using [member ParallaxLayer.motion_offset]. This creates an illusion of depth in a 2D game. If not used with a [Camera2D], you must manually calculate the [member scroll_offset].
*/
type ParallaxBackground struct {
	CanvasLayer
	owner gdnative.Object
}

func (o *ParallaxBackground) BaseClass() string {
	return "ParallaxBackground"
}

/*
        Undocumented
	Args: [{ false arg0 Transform2D} { false arg1 Vector2}], Returns: void
*/
func (o *ParallaxBackground) X_CameraMoved(arg0 gdnative.Transform2D, arg1 gdnative.Vector2) {
	//log.Println("Calling ParallaxBackground.X_CameraMoved()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromTransform2D(arg0)
	ptrArguments[1] = gdnative.NewPointerFromVector2(arg1)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "_camera_moved")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *ParallaxBackground) GetLimitBegin() gdnative.Vector2 {
	//log.Println("Calling ParallaxBackground.GetLimitBegin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "get_limit_begin")

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
	Args: [], Returns: Vector2
*/
func (o *ParallaxBackground) GetLimitEnd() gdnative.Vector2 {
	//log.Println("Calling ParallaxBackground.GetLimitEnd()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "get_limit_end")

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
	Args: [], Returns: Vector2
*/
func (o *ParallaxBackground) GetScrollBaseOffset() gdnative.Vector2 {
	//log.Println("Calling ParallaxBackground.GetScrollBaseOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "get_scroll_base_offset")

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
	Args: [], Returns: Vector2
*/
func (o *ParallaxBackground) GetScrollBaseScale() gdnative.Vector2 {
	//log.Println("Calling ParallaxBackground.GetScrollBaseScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "get_scroll_base_scale")

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
	Args: [], Returns: Vector2
*/
func (o *ParallaxBackground) GetScrollOffset() gdnative.Vector2 {
	//log.Println("Calling ParallaxBackground.GetScrollOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "get_scroll_offset")

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
	Args: [], Returns: bool
*/
func (o *ParallaxBackground) IsIgnoreCameraZoom() gdnative.Bool {
	//log.Println("Calling ParallaxBackground.IsIgnoreCameraZoom()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "is_ignore_camera_zoom")

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
	Args: [{ false ignore bool}], Returns: void
*/
func (o *ParallaxBackground) SetIgnoreCameraZoom(ignore gdnative.Bool) {
	//log.Println("Calling ParallaxBackground.SetIgnoreCameraZoom()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(ignore)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "set_ignore_camera_zoom")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ofs Vector2}], Returns: void
*/
func (o *ParallaxBackground) SetLimitBegin(ofs gdnative.Vector2) {
	//log.Println("Calling ParallaxBackground.SetLimitBegin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(ofs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "set_limit_begin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ofs Vector2}], Returns: void
*/
func (o *ParallaxBackground) SetLimitEnd(ofs gdnative.Vector2) {
	//log.Println("Calling ParallaxBackground.SetLimitEnd()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(ofs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "set_limit_end")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ofs Vector2}], Returns: void
*/
func (o *ParallaxBackground) SetScrollBaseOffset(ofs gdnative.Vector2) {
	//log.Println("Calling ParallaxBackground.SetScrollBaseOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(ofs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "set_scroll_base_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false scale Vector2}], Returns: void
*/
func (o *ParallaxBackground) SetScrollBaseScale(scale gdnative.Vector2) {
	//log.Println("Calling ParallaxBackground.SetScrollBaseScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(scale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "set_scroll_base_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ofs Vector2}], Returns: void
*/
func (o *ParallaxBackground) SetScrollOffset(ofs gdnative.Vector2) {
	//log.Println("Calling ParallaxBackground.SetScrollOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(ofs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParallaxBackground", "set_scroll_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ParallaxBackgroundImplementer is an interface that implements the methods
// of the ParallaxBackground class.
type ParallaxBackgroundImplementer interface {
	CanvasLayerImplementer
	X_CameraMoved(arg0 gdnative.Transform2D, arg1 gdnative.Vector2)
	GetLimitBegin() gdnative.Vector2
	GetLimitEnd() gdnative.Vector2
	GetScrollBaseOffset() gdnative.Vector2
	GetScrollBaseScale() gdnative.Vector2
	GetScrollOffset() gdnative.Vector2
	IsIgnoreCameraZoom() gdnative.Bool
	SetIgnoreCameraZoom(ignore gdnative.Bool)
	SetLimitBegin(ofs gdnative.Vector2)
	SetLimitEnd(ofs gdnative.Vector2)
	SetScrollBaseOffset(ofs gdnative.Vector2)
	SetScrollBaseScale(scale gdnative.Vector2)
	SetScrollOffset(ofs gdnative.Vector2)
}
