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

//func NewInputEventJoypadMotionFromPointer(ptr gdnative.Pointer) InputEventJoypadMotion {
func newInputEventJoypadMotionFromPointer(ptr gdnative.Pointer) InputEventJoypadMotion {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputEventJoypadMotion{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Stores information about joystick motions. One [InputEventJoypadMotion] represents one axis at a time.
*/
type InputEventJoypadMotion struct {
	InputEvent
	owner gdnative.Object
}

func (o *InputEventJoypadMotion) BaseClass() string {
	return "InputEventJoypadMotion"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *InputEventJoypadMotion) GetAxis() gdnative.Int {
	//log.Println("Calling InputEventJoypadMotion.GetAxis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventJoypadMotion", "get_axis")

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
func (o *InputEventJoypadMotion) GetAxisValue() gdnative.Real {
	//log.Println("Calling InputEventJoypadMotion.GetAxisValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventJoypadMotion", "get_axis_value")

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
	Args: [{ false axis int}], Returns: void
*/
func (o *InputEventJoypadMotion) SetAxis(axis gdnative.Int) {
	//log.Println("Calling InputEventJoypadMotion.SetAxis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(axis)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventJoypadMotion", "set_axis")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false axis_value float}], Returns: void
*/
func (o *InputEventJoypadMotion) SetAxisValue(axisValue gdnative.Real) {
	//log.Println("Calling InputEventJoypadMotion.SetAxisValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(axisValue)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventJoypadMotion", "set_axis_value")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// InputEventJoypadMotionImplementer is an interface that implements the methods
// of the InputEventJoypadMotion class.
type InputEventJoypadMotionImplementer interface {
	InputEventImplementer
	GetAxis() gdnative.Int
	GetAxisValue() gdnative.Real
	SetAxis(axis gdnative.Int)
	SetAxisValue(axisValue gdnative.Real)
}
