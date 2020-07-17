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

//func NewVisibilityNotifierFromPointer(ptr gdnative.Pointer) VisibilityNotifier {
func newVisibilityNotifierFromPointer(ptr gdnative.Pointer) VisibilityNotifier {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisibilityNotifier{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The VisibilityNotifier detects when it is visible on the screen. It also notifies when its bounding rectangle enters or exits the screen or a [Camera]'s view. [b]Note:[/b] VisibilityNotifier uses an approximate heuristic for performance reasons. It doesn't take walls and other occlusion into account. If you need exact visibility checking, use another method such as adding an [Area] node as a child of a [Camera] node.
*/
type VisibilityNotifier struct {
	Spatial
	owner gdnative.Object
}

func (o *VisibilityNotifier) BaseClass() string {
	return "VisibilityNotifier"
}

/*
        Undocumented
	Args: [], Returns: AABB
*/
func (o *VisibilityNotifier) GetAabb() gdnative.Aabb {
	//log.Println("Calling VisibilityNotifier.GetAabb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisibilityNotifier", "get_aabb")

	// Call the parent method.
	// AABB
	retPtr := gdnative.NewEmptyAabb()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewAabbFromPointer(retPtr)
	return ret
}

/*
        If [code]true[/code], the bounding box is on the screen. [b]Note:[/b] It takes one frame for the node's visibility to be assessed once added to the scene tree, so this method will return [code]false[/code] right after it is instantiated, even if it will be on screen in the draw pass.
	Args: [], Returns: bool
*/
func (o *VisibilityNotifier) IsOnScreen() gdnative.Bool {
	//log.Println("Calling VisibilityNotifier.IsOnScreen()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisibilityNotifier", "is_on_screen")

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
	Args: [{ false rect AABB}], Returns: void
*/
func (o *VisibilityNotifier) SetAabb(rect gdnative.Aabb) {
	//log.Println("Calling VisibilityNotifier.SetAabb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromAabb(rect)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisibilityNotifier", "set_aabb")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisibilityNotifierImplementer is an interface that implements the methods
// of the VisibilityNotifier class.
type VisibilityNotifierImplementer interface {
	SpatialImplementer
	GetAabb() gdnative.Aabb
	IsOnScreen() gdnative.Bool
	SetAabb(rect gdnative.Aabb)
}
