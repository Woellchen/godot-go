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

//func NewPosition2DFromPointer(ptr gdnative.Pointer) Position2D {
func newPosition2DFromPointer(ptr gdnative.Pointer) Position2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Position2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Generic 2D position hint for editing. It's just like a plain [Node2D], but it displays as a cross in the 2D editor at all times. You can set cross' visual size by using the gizmo in the 2D editor while the node is selected.
*/
type Position2D struct {
	Node2D
	owner gdnative.Object
}

func (o *Position2D) BaseClass() string {
	return "Position2D"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Position2D) X_GetGizmoExtents() gdnative.Real {
	// log.Println("Calling Position2D.X_GetGizmoExtents()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Position2D", "_get_gizmo_extents")

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
	Args: [{ false extents float}], Returns: void
*/
func (o *Position2D) X_SetGizmoExtents(extents gdnative.Real) {
	// log.Println("Calling Position2D.X_SetGizmoExtents()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(extents)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Position2D", "_set_gizmo_extents")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// Position2DImplementer is an interface that implements the methods
// of the Position2D class.
type Position2DImplementer interface {
	Node2DImplementer
	X_GetGizmoExtents() gdnative.Real
	X_SetGizmoExtents(extents gdnative.Real)
}
