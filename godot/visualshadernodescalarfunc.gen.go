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

// VisualShaderNodeScalarFuncFunction is an enum for Function values.
type VisualShaderNodeScalarFuncFunction int

const (
	VisualShaderNodeScalarFuncFuncAbs         VisualShaderNodeScalarFuncFunction = 12
	VisualShaderNodeScalarFuncFuncAcos        VisualShaderNodeScalarFuncFunction = 4
	VisualShaderNodeScalarFuncFuncAcosh       VisualShaderNodeScalarFuncFunction = 20
	VisualShaderNodeScalarFuncFuncAsin        VisualShaderNodeScalarFuncFunction = 3
	VisualShaderNodeScalarFuncFuncAsinh       VisualShaderNodeScalarFuncFunction = 21
	VisualShaderNodeScalarFuncFuncAtan        VisualShaderNodeScalarFuncFunction = 5
	VisualShaderNodeScalarFuncFuncAtanh       VisualShaderNodeScalarFuncFunction = 22
	VisualShaderNodeScalarFuncFuncCeil        VisualShaderNodeScalarFuncFunction = 16
	VisualShaderNodeScalarFuncFuncCos         VisualShaderNodeScalarFuncFunction = 1
	VisualShaderNodeScalarFuncFuncCosh        VisualShaderNodeScalarFuncFunction = 7
	VisualShaderNodeScalarFuncFuncDegrees     VisualShaderNodeScalarFuncFunction = 23
	VisualShaderNodeScalarFuncFuncExp         VisualShaderNodeScalarFuncFunction = 10
	VisualShaderNodeScalarFuncFuncExp2        VisualShaderNodeScalarFuncFunction = 24
	VisualShaderNodeScalarFuncFuncFloor       VisualShaderNodeScalarFuncFunction = 14
	VisualShaderNodeScalarFuncFuncFrac        VisualShaderNodeScalarFuncFunction = 17
	VisualShaderNodeScalarFuncFuncInverseSqrt VisualShaderNodeScalarFuncFunction = 25
	VisualShaderNodeScalarFuncFuncLog         VisualShaderNodeScalarFuncFunction = 9
	VisualShaderNodeScalarFuncFuncLog2        VisualShaderNodeScalarFuncFunction = 26
	VisualShaderNodeScalarFuncFuncNegate      VisualShaderNodeScalarFuncFunction = 19
	VisualShaderNodeScalarFuncFuncOneminus    VisualShaderNodeScalarFuncFunction = 31
	VisualShaderNodeScalarFuncFuncRadians     VisualShaderNodeScalarFuncFunction = 27
	VisualShaderNodeScalarFuncFuncReciprocal  VisualShaderNodeScalarFuncFunction = 28
	VisualShaderNodeScalarFuncFuncRound       VisualShaderNodeScalarFuncFunction = 15
	VisualShaderNodeScalarFuncFuncRoundeven   VisualShaderNodeScalarFuncFunction = 29
	VisualShaderNodeScalarFuncFuncSaturate    VisualShaderNodeScalarFuncFunction = 18
	VisualShaderNodeScalarFuncFuncSign        VisualShaderNodeScalarFuncFunction = 13
	VisualShaderNodeScalarFuncFuncSin         VisualShaderNodeScalarFuncFunction = 0
	VisualShaderNodeScalarFuncFuncSinh        VisualShaderNodeScalarFuncFunction = 6
	VisualShaderNodeScalarFuncFuncSqrt        VisualShaderNodeScalarFuncFunction = 11
	VisualShaderNodeScalarFuncFuncTan         VisualShaderNodeScalarFuncFunction = 2
	VisualShaderNodeScalarFuncFuncTanh        VisualShaderNodeScalarFuncFunction = 8
	VisualShaderNodeScalarFuncFuncTrunc       VisualShaderNodeScalarFuncFunction = 30
)

//func NewVisualShaderNodeScalarFuncFromPointer(ptr gdnative.Pointer) VisualShaderNodeScalarFunc {
func newVisualShaderNodeScalarFuncFromPointer(ptr gdnative.Pointer) VisualShaderNodeScalarFunc {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeScalarFunc{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNodeScalarFunc struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeScalarFunc) BaseClass() string {
	return "VisualShaderNodeScalarFunc"
}

/*
        Undocumented
	Args: [], Returns: enum.VisualShaderNodeScalarFunc::Function
*/
func (o *VisualShaderNodeScalarFunc) GetFunction() VisualShaderNodeScalarFuncFunction {
	// log.Println("Calling VisualShaderNodeScalarFunc.GetFunction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeScalarFunc", "get_function")

	// Call the parent method.
	// enum.VisualShaderNodeScalarFunc::Function
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualShaderNodeScalarFuncFunction(ret)
}

/*
        Undocumented
	Args: [{ false func int}], Returns: void
*/
func (o *VisualShaderNodeScalarFunc) SetFunction(function gdnative.Int) {
	// log.Println("Calling VisualShaderNodeScalarFunc.SetFunction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(function)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeScalarFunc", "set_function")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderNodeScalarFuncImplementer is an interface that implements the methods
// of the VisualShaderNodeScalarFunc class.
type VisualShaderNodeScalarFuncImplementer interface {
	VisualShaderNodeImplementer
	SetFunction(function gdnative.Int)
}
