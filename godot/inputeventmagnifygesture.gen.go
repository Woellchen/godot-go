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

//func NewInputEventMagnifyGestureFromPointer(ptr gdnative.Pointer) InputEventMagnifyGesture {
func newInputEventMagnifyGestureFromPointer(ptr gdnative.Pointer) InputEventMagnifyGesture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputEventMagnifyGesture{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type InputEventMagnifyGesture struct {
	InputEventGesture
	owner gdnative.Object
}

func (o *InputEventMagnifyGesture) BaseClass() string {
	return "InputEventMagnifyGesture"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *InputEventMagnifyGesture) GetFactor() gdnative.Real {
	// log.Println("Calling InputEventMagnifyGesture.GetFactor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventMagnifyGesture", "get_factor")

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
	Args: [{ false factor float}], Returns: void
*/
func (o *InputEventMagnifyGesture) SetFactor(factor gdnative.Real) {
	// log.Println("Calling InputEventMagnifyGesture.SetFactor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(factor)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventMagnifyGesture", "set_factor")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// InputEventMagnifyGestureImplementer is an interface that implements the methods
// of the InputEventMagnifyGesture class.
type InputEventMagnifyGestureImplementer interface {
	InputEventGestureImplementer
	GetFactor() gdnative.Real
	SetFactor(factor gdnative.Real)
}
