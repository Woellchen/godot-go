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

//func NewInputEventGestureFromPointer(ptr gdnative.Pointer) InputEventGesture {
func newInputEventGestureFromPointer(ptr gdnative.Pointer) InputEventGesture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputEventGesture{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type InputEventGesture struct {
	InputEventWithModifiers
	owner gdnative.Object
}

func (o *InputEventGesture) BaseClass() string {
	return "InputEventGesture"
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *InputEventGesture) GetPosition() gdnative.Vector2 {
	// log.Println("Calling InputEventGesture.GetPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventGesture", "get_position")

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
	Args: [{ false position Vector2}], Returns: void
*/
func (o *InputEventGesture) SetPosition(position gdnative.Vector2) {
	// log.Println("Calling InputEventGesture.SetPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventGesture", "set_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// InputEventGestureImplementer is an interface that implements the methods
// of the InputEventGesture class.
type InputEventGestureImplementer interface {
	InputEventWithModifiersImplementer
	GetPosition() gdnative.Vector2
	SetPosition(position gdnative.Vector2)
}
