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

// AnimationNodeStateMachineTransitionSwitchMode is an enum for SwitchMode values.
type AnimationNodeStateMachineTransitionSwitchMode int

const (
	AnimationNodeStateMachineTransitionSwitchModeAtEnd     AnimationNodeStateMachineTransitionSwitchMode = 2
	AnimationNodeStateMachineTransitionSwitchModeImmediate AnimationNodeStateMachineTransitionSwitchMode = 0
	AnimationNodeStateMachineTransitionSwitchModeSync      AnimationNodeStateMachineTransitionSwitchMode = 1
)

//func NewAnimationNodeStateMachineTransitionFromPointer(ptr gdnative.Pointer) AnimationNodeStateMachineTransition {
func newAnimationNodeStateMachineTransitionFromPointer(ptr gdnative.Pointer) AnimationNodeStateMachineTransition {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNodeStateMachineTransition{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AnimationNodeStateMachineTransition struct {
	Resource
	owner gdnative.Object
}

func (o *AnimationNodeStateMachineTransition) BaseClass() string {
	return "AnimationNodeStateMachineTransition"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AnimationNodeStateMachineTransition) GetAdvanceCondition() gdnative.String {
	//log.Println("Calling AnimationNodeStateMachineTransition.GetAdvanceCondition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "get_advance_condition")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *AnimationNodeStateMachineTransition) GetPriority() gdnative.Int {
	//log.Println("Calling AnimationNodeStateMachineTransition.GetPriority()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "get_priority")

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
	Args: [], Returns: enum.AnimationNodeStateMachineTransition::SwitchMode
*/
func (o *AnimationNodeStateMachineTransition) GetSwitchMode() AnimationNodeStateMachineTransitionSwitchMode {
	//log.Println("Calling AnimationNodeStateMachineTransition.GetSwitchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "get_switch_mode")

	// Call the parent method.
	// enum.AnimationNodeStateMachineTransition::SwitchMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return AnimationNodeStateMachineTransitionSwitchMode(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AnimationNodeStateMachineTransition) GetXfadeTime() gdnative.Real {
	//log.Println("Calling AnimationNodeStateMachineTransition.GetXfadeTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "get_xfade_time")

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
	Args: [], Returns: bool
*/
func (o *AnimationNodeStateMachineTransition) HasAutoAdvance() gdnative.Bool {
	//log.Println("Calling AnimationNodeStateMachineTransition.HasAutoAdvance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "has_auto_advance")

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
func (o *AnimationNodeStateMachineTransition) IsDisabled() gdnative.Bool {
	//log.Println("Calling AnimationNodeStateMachineTransition.IsDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "is_disabled")

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
	Args: [{ false name String}], Returns: void
*/
func (o *AnimationNodeStateMachineTransition) SetAdvanceCondition(name gdnative.String) {
	//log.Println("Calling AnimationNodeStateMachineTransition.SetAdvanceCondition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "set_advance_condition")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false auto_advance bool}], Returns: void
*/
func (o *AnimationNodeStateMachineTransition) SetAutoAdvance(autoAdvance gdnative.Bool) {
	//log.Println("Calling AnimationNodeStateMachineTransition.SetAutoAdvance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(autoAdvance)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "set_auto_advance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false disabled bool}], Returns: void
*/
func (o *AnimationNodeStateMachineTransition) SetDisabled(disabled gdnative.Bool) {
	//log.Println("Calling AnimationNodeStateMachineTransition.SetDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(disabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "set_disabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false priority int}], Returns: void
*/
func (o *AnimationNodeStateMachineTransition) SetPriority(priority gdnative.Int) {
	//log.Println("Calling AnimationNodeStateMachineTransition.SetPriority()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(priority)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "set_priority")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *AnimationNodeStateMachineTransition) SetSwitchMode(mode gdnative.Int) {
	//log.Println("Calling AnimationNodeStateMachineTransition.SetSwitchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "set_switch_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false secs float}], Returns: void
*/
func (o *AnimationNodeStateMachineTransition) SetXfadeTime(secs gdnative.Real) {
	//log.Println("Calling AnimationNodeStateMachineTransition.SetXfadeTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(secs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachineTransition", "set_xfade_time")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimationNodeStateMachineTransitionImplementer is an interface that implements the methods
// of the AnimationNodeStateMachineTransition class.
type AnimationNodeStateMachineTransitionImplementer interface {
	ResourceImplementer
	GetAdvanceCondition() gdnative.String
	GetPriority() gdnative.Int
	GetXfadeTime() gdnative.Real
	HasAutoAdvance() gdnative.Bool
	IsDisabled() gdnative.Bool
	SetAdvanceCondition(name gdnative.String)
	SetAutoAdvance(autoAdvance gdnative.Bool)
	SetDisabled(disabled gdnative.Bool)
	SetPriority(priority gdnative.Int)
	SetSwitchMode(mode gdnative.Int)
	SetXfadeTime(secs gdnative.Real)
}
