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

// HingeJointFlag is an enum for Flag values.
type HingeJointFlag int

const (
	HingeJointFlagEnableMotor HingeJointFlag = 1
	HingeJointFlagMax         HingeJointFlag = 2
	HingeJointFlagUseLimit    HingeJointFlag = 0
)

// HingeJointParam is an enum for Param values.
type HingeJointParam int

const (
	HingeJointParamBias                HingeJointParam = 0
	HingeJointParamLimitBias           HingeJointParam = 3
	HingeJointParamLimitLower          HingeJointParam = 2
	HingeJointParamLimitRelaxation     HingeJointParam = 5
	HingeJointParamLimitSoftness       HingeJointParam = 4
	HingeJointParamLimitUpper          HingeJointParam = 1
	HingeJointParamMax                 HingeJointParam = 8
	HingeJointParamMotorMaxImpulse     HingeJointParam = 7
	HingeJointParamMotorTargetVelocity HingeJointParam = 6
)

//func NewHingeJointFromPointer(ptr gdnative.Pointer) HingeJoint {
func newHingeJointFromPointer(ptr gdnative.Pointer) HingeJoint {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HingeJoint{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A HingeJoint normally uses the Z axis of body A as the hinge axis, another axis can be specified when adding it manually though.
*/
type HingeJoint struct {
	Joint
	owner gdnative.Object
}

func (o *HingeJoint) BaseClass() string {
	return "HingeJoint"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *HingeJoint) X_GetLowerLimit() gdnative.Real {
	// log.Println("Calling HingeJoint.X_GetLowerLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "_get_lower_limit")

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
func (o *HingeJoint) X_GetUpperLimit() gdnative.Real {
	// log.Println("Calling HingeJoint.X_GetUpperLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "_get_upper_limit")

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
	Args: [{ false lower_limit float}], Returns: void
*/
func (o *HingeJoint) X_SetLowerLimit(lowerLimit gdnative.Real) {
	// log.Println("Calling HingeJoint.X_SetLowerLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(lowerLimit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "_set_lower_limit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false upper_limit float}], Returns: void
*/
func (o *HingeJoint) X_SetUpperLimit(upperLimit gdnative.Real) {
	// log.Println("Calling HingeJoint.X_SetUpperLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(upperLimit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "_set_upper_limit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the value of the specified flag.
	Args: [{ false flag int}], Returns: bool
*/
func (o *HingeJoint) GetFlag(flag gdnative.Int) gdnative.Bool {
	// log.Println("Calling HingeJoint.GetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "get_flag")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns the value of the specified parameter.
	Args: [{ false param int}], Returns: float
*/
func (o *HingeJoint) GetParam(param gdnative.Int) gdnative.Real {
	// log.Println("Calling HingeJoint.GetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "get_param")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        If [code]true[/code], enables the specified flag.
	Args: [{ false flag int} { false enabled bool}], Returns: void
*/
func (o *HingeJoint) SetFlag(flag gdnative.Int, enabled gdnative.Bool) {
	// log.Println("Calling HingeJoint.SetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)
	ptrArguments[1] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "set_flag")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the value of the specified parameter.
	Args: [{ false param int} { false value float}], Returns: void
*/
func (o *HingeJoint) SetParam(param gdnative.Int, value gdnative.Real) {
	// log.Println("Calling HingeJoint.SetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromReal(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "set_param")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// HingeJointImplementer is an interface that implements the methods
// of the HingeJoint class.
type HingeJointImplementer interface {
	JointImplementer
	X_GetLowerLimit() gdnative.Real
	X_GetUpperLimit() gdnative.Real
	X_SetLowerLimit(lowerLimit gdnative.Real)
	X_SetUpperLimit(upperLimit gdnative.Real)
	GetFlag(flag gdnative.Int) gdnative.Bool
	GetParam(param gdnative.Int) gdnative.Real
	SetFlag(flag gdnative.Int, enabled gdnative.Bool)
	SetParam(param gdnative.Int, value gdnative.Real)
}
