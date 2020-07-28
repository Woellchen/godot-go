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

// Camera2DAnchorMode is an enum for AnchorMode values.
type Camera2DAnchorMode int

const (
	Camera2DAnchorModeDragCenter   Camera2DAnchorMode = 1
	Camera2DAnchorModeFixedTopLeft Camera2DAnchorMode = 0
)

// Camera2DCamera2DProcessMode is an enum for Camera2DProcessMode values.
type Camera2DCamera2DProcessMode int

const (
	Camera2DCamera2DProcessIdle    Camera2DCamera2DProcessMode = 1
	Camera2DCamera2DProcessPhysics Camera2DCamera2DProcessMode = 0
)

//func NewCamera2DFromPointer(ptr gdnative.Pointer) Camera2D {
func newCamera2DFromPointer(ptr gdnative.Pointer) Camera2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Camera2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Camera node for 2D scenes. It forces the screen (current layer) to scroll following this node. This makes it easier (and faster) to program scrollable scenes than manually changing the position of [CanvasItem]-based nodes. This node is intended to be a simple helper to get things going quickly and it may happen that more functionality is desired to change how the camera works. To make your own custom camera node, inherit from [Node2D] and change the transform of the canvas by setting [member Viewport.canvas_transform] in [Viewport] (you can obtain the current [Viewport] by using [method Node.get_viewport]). Note that the [Camera2D] node's [code]position[/code] doesn't represent the actual position of the screen, which may differ due to applied smoothing or limits. You can use [method get_camera_screen_center] to get the real position.
*/
type Camera2D struct {
	Node2D
	owner gdnative.Object
}

func (o *Camera2D) BaseClass() string {
	return "Camera2D"
}

/*
        Undocumented
	Args: [{ false arg0 Object}], Returns: void
*/
func (o *Camera2D) X_MakeCurrent(arg0 ObjectImplementer) {
	// log.Println("Calling Camera2D.X_MakeCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "_make_current")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false current bool}], Returns: void
*/
func (o *Camera2D) X_SetCurrent(current gdnative.Bool) {
	// log.Println("Calling Camera2D.X_SetCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(current)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "_set_current")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false follow_smoothing float}], Returns: void
*/
func (o *Camera2D) X_SetOldSmoothing(followSmoothing gdnative.Real) {
	// log.Println("Calling Camera2D.X_SetOldSmoothing()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(followSmoothing)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "_set_old_smoothing")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Camera2D) X_UpdateScroll() {
	// log.Println("Calling Camera2D.X_UpdateScroll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "_update_scroll")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Aligns the camera to the tracked node.
	Args: [], Returns: void
*/
func (o *Camera2D) Align() {
	// log.Println("Calling Camera2D.Align()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "align")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes any [Camera2D] from the ancestor [Viewport]'s internal currently-assigned camera.
	Args: [], Returns: void
*/
func (o *Camera2D) ClearCurrent() {
	// log.Println("Calling Camera2D.ClearCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "clear_current")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Forces the camera to update scroll immediately.
	Args: [], Returns: void
*/
func (o *Camera2D) ForceUpdateScroll() {
	// log.Println("Calling Camera2D.ForceUpdateScroll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "force_update_scroll")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: enum.Camera2D::AnchorMode
*/
func (o *Camera2D) GetAnchorMode() Camera2DAnchorMode {
	// log.Println("Calling Camera2D.GetAnchorMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_anchor_mode")

	// Call the parent method.
	// enum.Camera2D::AnchorMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return Camera2DAnchorMode(ret)
}

/*
        Returns the camera position.
	Args: [], Returns: Vector2
*/
func (o *Camera2D) GetCameraPosition() gdnative.Vector2 {
	// log.Println("Calling Camera2D.GetCameraPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_camera_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns the location of the [Camera2D]'s screen-center, relative to the origin.
	Args: [], Returns: Vector2
*/
func (o *Camera2D) GetCameraScreenCenter() gdnative.Vector2 {
	// log.Println("Calling Camera2D.GetCameraScreenCenter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_camera_screen_center")

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
	Args: [], Returns: Node
*/
func (o *Camera2D) GetCustomViewport() NodeImplementer {
	// log.Println("Calling Camera2D.GetCustomViewport()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_custom_viewport")

	// Call the parent method.
	// Node
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newNodeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(NodeImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Node" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(NodeImplementer)
	}

	return &ret
}

/*
        Returns the specified margin. See also [member drag_margin_bottom], [member drag_margin_top], [member drag_margin_left], and [member drag_margin_right].
	Args: [{ false margin int}], Returns: float
*/
func (o *Camera2D) GetDragMargin(margin gdnative.Int) gdnative.Real {
	// log.Println("Calling Camera2D.GetDragMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_drag_margin")

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
	Args: [], Returns: float
*/
func (o *Camera2D) GetFollowSmoothing() gdnative.Real {
	// log.Println("Calling Camera2D.GetFollowSmoothing()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_follow_smoothing")

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
	Args: [], Returns: float
*/
func (o *Camera2D) GetHOffset() gdnative.Real {
	// log.Println("Calling Camera2D.GetHOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_h_offset")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the specified camera limit. See also [member limit_bottom], [member limit_top], [member limit_left], and [member limit_right].
	Args: [{ false margin int}], Returns: int
*/
func (o *Camera2D) GetLimit(margin gdnative.Int) gdnative.Int {
	// log.Println("Calling Camera2D.GetLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_limit")

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
	Args: [], Returns: Vector2
*/
func (o *Camera2D) GetOffset() gdnative.Vector2 {
	// log.Println("Calling Camera2D.GetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_offset")

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
	Args: [], Returns: enum.Camera2D::Camera2DProcessMode
*/
func (o *Camera2D) GetProcessMode() Camera2DCamera2DProcessMode {
	// log.Println("Calling Camera2D.GetProcessMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_process_mode")

	// Call the parent method.
	// enum.Camera2D::Camera2DProcessMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return Camera2DCamera2DProcessMode(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Camera2D) GetVOffset() gdnative.Real {
	// log.Println("Calling Camera2D.GetVOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_v_offset")

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
	Args: [], Returns: Vector2
*/
func (o *Camera2D) GetZoom() gdnative.Vector2 {
	// log.Println("Calling Camera2D.GetZoom()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "get_zoom")

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
func (o *Camera2D) IsCurrent() gdnative.Bool {
	// log.Println("Calling Camera2D.IsCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "is_current")

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
func (o *Camera2D) IsFollowSmoothingEnabled() gdnative.Bool {
	// log.Println("Calling Camera2D.IsFollowSmoothingEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "is_follow_smoothing_enabled")

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
func (o *Camera2D) IsHDragEnabled() gdnative.Bool {
	// log.Println("Calling Camera2D.IsHDragEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "is_h_drag_enabled")

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
func (o *Camera2D) IsLimitDrawingEnabled() gdnative.Bool {
	// log.Println("Calling Camera2D.IsLimitDrawingEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "is_limit_drawing_enabled")

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
func (o *Camera2D) IsLimitSmoothingEnabled() gdnative.Bool {
	// log.Println("Calling Camera2D.IsLimitSmoothingEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "is_limit_smoothing_enabled")

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
func (o *Camera2D) IsMarginDrawingEnabled() gdnative.Bool {
	// log.Println("Calling Camera2D.IsMarginDrawingEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "is_margin_drawing_enabled")

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
func (o *Camera2D) IsRotating() gdnative.Bool {
	// log.Println("Calling Camera2D.IsRotating()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "is_rotating")

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
func (o *Camera2D) IsScreenDrawingEnabled() gdnative.Bool {
	// log.Println("Calling Camera2D.IsScreenDrawingEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "is_screen_drawing_enabled")

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
func (o *Camera2D) IsVDragEnabled() gdnative.Bool {
	// log.Println("Calling Camera2D.IsVDragEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "is_v_drag_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Make this the current 2D camera for the scene (viewport and layer), in case there are many cameras in the scene.
	Args: [], Returns: void
*/
func (o *Camera2D) MakeCurrent() {
	// log.Println("Calling Camera2D.MakeCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "make_current")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the camera's position immediately to its current smoothing destination. This has no effect if smoothing is disabled.
	Args: [], Returns: void
*/
func (o *Camera2D) ResetSmoothing() {
	// log.Println("Calling Camera2D.ResetSmoothing()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "reset_smoothing")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false anchor_mode int}], Returns: void
*/
func (o *Camera2D) SetAnchorMode(anchorMode gdnative.Int) {
	// log.Println("Calling Camera2D.SetAnchorMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(anchorMode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_anchor_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false viewport Node}], Returns: void
*/
func (o *Camera2D) SetCustomViewport(viewport NodeImplementer) {
	// log.Println("Calling Camera2D.SetCustomViewport()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(viewport.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_custom_viewport")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the specified margin. See also [member drag_margin_bottom], [member drag_margin_top], [member drag_margin_left], and [member drag_margin_right].
	Args: [{ false margin int} { false drag_margin float}], Returns: void
*/
func (o *Camera2D) SetDragMargin(margin gdnative.Int, dragMargin gdnative.Real) {
	// log.Println("Calling Camera2D.SetDragMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)
	ptrArguments[1] = gdnative.NewPointerFromReal(dragMargin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_drag_margin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false follow_smoothing bool}], Returns: void
*/
func (o *Camera2D) SetEnableFollowSmoothing(followSmoothing gdnative.Bool) {
	// log.Println("Calling Camera2D.SetEnableFollowSmoothing()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(followSmoothing)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_enable_follow_smoothing")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false follow_smoothing float}], Returns: void
*/
func (o *Camera2D) SetFollowSmoothing(followSmoothing gdnative.Real) {
	// log.Println("Calling Camera2D.SetFollowSmoothing()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(followSmoothing)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_follow_smoothing")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Camera2D) SetHDragEnabled(enabled gdnative.Bool) {
	// log.Println("Calling Camera2D.SetHDragEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_h_drag_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ofs float}], Returns: void
*/
func (o *Camera2D) SetHOffset(ofs gdnative.Real) {
	// log.Println("Calling Camera2D.SetHOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(ofs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_h_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the specified camera limit. See also [member limit_bottom], [member limit_top], [member limit_left], and [member limit_right].
	Args: [{ false margin int} { false limit int}], Returns: void
*/
func (o *Camera2D) SetLimit(margin gdnative.Int, limit gdnative.Int) {
	// log.Println("Calling Camera2D.SetLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)
	ptrArguments[1] = gdnative.NewPointerFromInt(limit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_limit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false limit_drawing_enabled bool}], Returns: void
*/
func (o *Camera2D) SetLimitDrawingEnabled(limitDrawingEnabled gdnative.Bool) {
	// log.Println("Calling Camera2D.SetLimitDrawingEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(limitDrawingEnabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_limit_drawing_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false limit_smoothing_enabled bool}], Returns: void
*/
func (o *Camera2D) SetLimitSmoothingEnabled(limitSmoothingEnabled gdnative.Bool) {
	// log.Println("Calling Camera2D.SetLimitSmoothingEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(limitSmoothingEnabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_limit_smoothing_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false margin_drawing_enabled bool}], Returns: void
*/
func (o *Camera2D) SetMarginDrawingEnabled(marginDrawingEnabled gdnative.Bool) {
	// log.Println("Calling Camera2D.SetMarginDrawingEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(marginDrawingEnabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_margin_drawing_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false offset Vector2}], Returns: void
*/
func (o *Camera2D) SetOffset(offset gdnative.Vector2) {
	// log.Println("Calling Camera2D.SetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *Camera2D) SetProcessMode(mode gdnative.Int) {
	// log.Println("Calling Camera2D.SetProcessMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_process_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rotating bool}], Returns: void
*/
func (o *Camera2D) SetRotating(rotating gdnative.Bool) {
	// log.Println("Calling Camera2D.SetRotating()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(rotating)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_rotating")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false screen_drawing_enabled bool}], Returns: void
*/
func (o *Camera2D) SetScreenDrawingEnabled(screenDrawingEnabled gdnative.Bool) {
	// log.Println("Calling Camera2D.SetScreenDrawingEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(screenDrawingEnabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_screen_drawing_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Camera2D) SetVDragEnabled(enabled gdnative.Bool) {
	// log.Println("Calling Camera2D.SetVDragEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_v_drag_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ofs float}], Returns: void
*/
func (o *Camera2D) SetVOffset(ofs gdnative.Real) {
	// log.Println("Calling Camera2D.SetVOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(ofs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_v_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false zoom Vector2}], Returns: void
*/
func (o *Camera2D) SetZoom(zoom gdnative.Vector2) {
	// log.Println("Calling Camera2D.SetZoom()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(zoom)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera2D", "set_zoom")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// Camera2DImplementer is an interface that implements the methods
// of the Camera2D class.
type Camera2DImplementer interface {
	Node2DImplementer
	X_MakeCurrent(arg0 ObjectImplementer)
	X_SetCurrent(current gdnative.Bool)
	X_SetOldSmoothing(followSmoothing gdnative.Real)
	X_UpdateScroll()
	Align()
	ClearCurrent()
	ForceUpdateScroll()
	GetCameraPosition() gdnative.Vector2
	GetCameraScreenCenter() gdnative.Vector2
	GetCustomViewport() NodeImplementer
	GetDragMargin(margin gdnative.Int) gdnative.Real
	GetFollowSmoothing() gdnative.Real
	GetHOffset() gdnative.Real
	GetLimit(margin gdnative.Int) gdnative.Int
	GetOffset() gdnative.Vector2
	GetVOffset() gdnative.Real
	GetZoom() gdnative.Vector2
	IsCurrent() gdnative.Bool
	IsFollowSmoothingEnabled() gdnative.Bool
	IsHDragEnabled() gdnative.Bool
	IsLimitDrawingEnabled() gdnative.Bool
	IsLimitSmoothingEnabled() gdnative.Bool
	IsMarginDrawingEnabled() gdnative.Bool
	IsRotating() gdnative.Bool
	IsScreenDrawingEnabled() gdnative.Bool
	IsVDragEnabled() gdnative.Bool
	MakeCurrent()
	ResetSmoothing()
	SetAnchorMode(anchorMode gdnative.Int)
	SetCustomViewport(viewport NodeImplementer)
	SetDragMargin(margin gdnative.Int, dragMargin gdnative.Real)
	SetEnableFollowSmoothing(followSmoothing gdnative.Bool)
	SetFollowSmoothing(followSmoothing gdnative.Real)
	SetHDragEnabled(enabled gdnative.Bool)
	SetHOffset(ofs gdnative.Real)
	SetLimit(margin gdnative.Int, limit gdnative.Int)
	SetLimitDrawingEnabled(limitDrawingEnabled gdnative.Bool)
	SetLimitSmoothingEnabled(limitSmoothingEnabled gdnative.Bool)
	SetMarginDrawingEnabled(marginDrawingEnabled gdnative.Bool)
	SetOffset(offset gdnative.Vector2)
	SetProcessMode(mode gdnative.Int)
	SetRotating(rotating gdnative.Bool)
	SetScreenDrawingEnabled(screenDrawingEnabled gdnative.Bool)
	SetVDragEnabled(enabled gdnative.Bool)
	SetVOffset(ofs gdnative.Real)
	SetZoom(zoom gdnative.Vector2)
}
