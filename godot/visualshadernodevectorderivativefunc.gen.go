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

// VisualShaderNodeVectorDerivativeFuncFunction is an enum for Function values.
type VisualShaderNodeVectorDerivativeFuncFunction int

const (
	VisualShaderNodeVectorDerivativeFuncFuncSum VisualShaderNodeVectorDerivativeFuncFunction = 0
	VisualShaderNodeVectorDerivativeFuncFuncX   VisualShaderNodeVectorDerivativeFuncFunction = 1
	VisualShaderNodeVectorDerivativeFuncFuncY   VisualShaderNodeVectorDerivativeFuncFunction = 2
)

//func NewVisualShaderNodeVectorDerivativeFuncFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorDerivativeFunc {
func newVisualShaderNodeVectorDerivativeFuncFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorDerivativeFunc {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeVectorDerivativeFunc{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This node is only available in [code]Fragment[/code] and [code]Light[/code] visual shaders.
*/
type VisualShaderNodeVectorDerivativeFunc struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeVectorDerivativeFunc) BaseClass() string {
	return "VisualShaderNodeVectorDerivativeFunc"
}

/*
        Undocumented
	Args: [], Returns: enum.VisualShaderNodeVectorDerivativeFunc::Function
*/
func (o *VisualShaderNodeVectorDerivativeFunc) GetFunction() VisualShaderNodeVectorDerivativeFuncFunction {
	// log.Println("Calling VisualShaderNodeVectorDerivativeFunc.GetFunction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeVectorDerivativeFunc", "get_function")

	// Call the parent method.
	// enum.VisualShaderNodeVectorDerivativeFunc::Function
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualShaderNodeVectorDerivativeFuncFunction(ret)
}

/*
        Undocumented
	Args: [{ false func int}], Returns: void
*/
func (o *VisualShaderNodeVectorDerivativeFunc) SetFunction(function gdnative.Int) {
	// log.Println("Calling VisualShaderNodeVectorDerivativeFunc.SetFunction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(function)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeVectorDerivativeFunc", "set_function")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderNodeVectorDerivativeFuncImplementer is an interface that implements the methods
// of the VisualShaderNodeVectorDerivativeFunc class.
type VisualShaderNodeVectorDerivativeFuncImplementer interface {
	VisualShaderNodeImplementer
	SetFunction(function gdnative.Int)
}
