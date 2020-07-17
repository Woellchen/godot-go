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

// VisualShaderNodeColorFuncFunction is an enum for Function values.
type VisualShaderNodeColorFuncFunction int

const (
	VisualShaderNodeColorFuncFuncGrayscale VisualShaderNodeColorFuncFunction = 0
	VisualShaderNodeColorFuncFuncSepia     VisualShaderNodeColorFuncFunction = 1
)

//func NewVisualShaderNodeColorFuncFromPointer(ptr gdnative.Pointer) VisualShaderNodeColorFunc {
func newVisualShaderNodeColorFuncFromPointer(ptr gdnative.Pointer) VisualShaderNodeColorFunc {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeColorFunc{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Accept a [Color] to the input port and transform it according to [member function].
*/
type VisualShaderNodeColorFunc struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeColorFunc) BaseClass() string {
	return "VisualShaderNodeColorFunc"
}

/*
        Undocumented
	Args: [], Returns: enum.VisualShaderNodeColorFunc::Function
*/
func (o *VisualShaderNodeColorFunc) GetFunction() VisualShaderNodeColorFuncFunction {
	//log.Println("Calling VisualShaderNodeColorFunc.GetFunction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeColorFunc", "get_function")

	// Call the parent method.
	// enum.VisualShaderNodeColorFunc::Function
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualShaderNodeColorFuncFunction(ret)
}

/*
        Undocumented
	Args: [{ false func int}], Returns: void
*/
func (o *VisualShaderNodeColorFunc) SetFunction(function gdnative.Int) {
	//log.Println("Calling VisualShaderNodeColorFunc.SetFunction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(function)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeColorFunc", "set_function")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderNodeColorFuncImplementer is an interface that implements the methods
// of the VisualShaderNodeColorFunc class.
type VisualShaderNodeColorFuncImplementer interface {
	VisualShaderNodeImplementer
	SetFunction(function gdnative.Int)
}
