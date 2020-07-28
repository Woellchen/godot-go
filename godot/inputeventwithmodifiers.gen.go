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

//func NewInputEventWithModifiersFromPointer(ptr gdnative.Pointer) InputEventWithModifiers {
func newInputEventWithModifiersFromPointer(ptr gdnative.Pointer) InputEventWithModifiers {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputEventWithModifiers{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Contains keys events information with modifiers support like [code]Shift[/code] or [code]Alt[/code]. See [method Node._input].
*/
type InputEventWithModifiers struct {
	InputEvent
	owner gdnative.Object
}

func (o *InputEventWithModifiers) BaseClass() string {
	return "InputEventWithModifiers"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *InputEventWithModifiers) GetAlt() gdnative.Bool {
	// log.Println("Calling InputEventWithModifiers.GetAlt()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventWithModifiers", "get_alt")

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
func (o *InputEventWithModifiers) GetCommand() gdnative.Bool {
	// log.Println("Calling InputEventWithModifiers.GetCommand()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventWithModifiers", "get_command")

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
func (o *InputEventWithModifiers) GetControl() gdnative.Bool {
	// log.Println("Calling InputEventWithModifiers.GetControl()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventWithModifiers", "get_control")

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
func (o *InputEventWithModifiers) GetMetakey() gdnative.Bool {
	// log.Println("Calling InputEventWithModifiers.GetMetakey()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventWithModifiers", "get_metakey")

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
func (o *InputEventWithModifiers) GetShift() gdnative.Bool {
	// log.Println("Calling InputEventWithModifiers.GetShift()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventWithModifiers", "get_shift")

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
	Args: [{ false enable bool}], Returns: void
*/
func (o *InputEventWithModifiers) SetAlt(enable gdnative.Bool) {
	// log.Println("Calling InputEventWithModifiers.SetAlt()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventWithModifiers", "set_alt")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *InputEventWithModifiers) SetCommand(enable gdnative.Bool) {
	// log.Println("Calling InputEventWithModifiers.SetCommand()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventWithModifiers", "set_command")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *InputEventWithModifiers) SetControl(enable gdnative.Bool) {
	// log.Println("Calling InputEventWithModifiers.SetControl()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventWithModifiers", "set_control")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *InputEventWithModifiers) SetMetakey(enable gdnative.Bool) {
	// log.Println("Calling InputEventWithModifiers.SetMetakey()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventWithModifiers", "set_metakey")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *InputEventWithModifiers) SetShift(enable gdnative.Bool) {
	// log.Println("Calling InputEventWithModifiers.SetShift()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventWithModifiers", "set_shift")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// InputEventWithModifiersImplementer is an interface that implements the methods
// of the InputEventWithModifiers class.
type InputEventWithModifiersImplementer interface {
	InputEventImplementer
	GetAlt() gdnative.Bool
	GetCommand() gdnative.Bool
	GetControl() gdnative.Bool
	GetMetakey() gdnative.Bool
	GetShift() gdnative.Bool
	SetAlt(enable gdnative.Bool)
	SetCommand(enable gdnative.Bool)
	SetControl(enable gdnative.Bool)
	SetMetakey(enable gdnative.Bool)
	SetShift(enable gdnative.Bool)
}
