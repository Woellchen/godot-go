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

// VisualShaderNodeCompareComparisonType is an enum for ComparisonType values.
type VisualShaderNodeCompareComparisonType int

const (
	VisualShaderNodeCompareCtypeBoolean   VisualShaderNodeCompareComparisonType = 2
	VisualShaderNodeCompareCtypeScalar    VisualShaderNodeCompareComparisonType = 0
	VisualShaderNodeCompareCtypeTransform VisualShaderNodeCompareComparisonType = 3
	VisualShaderNodeCompareCtypeVector    VisualShaderNodeCompareComparisonType = 1
)

// VisualShaderNodeCompareCondition is an enum for Condition values.
type VisualShaderNodeCompareCondition int

const (
	VisualShaderNodeCompareCondAll VisualShaderNodeCompareCondition = 0
	VisualShaderNodeCompareCondAny VisualShaderNodeCompareCondition = 1
)

// VisualShaderNodeCompareFunction is an enum for Function values.
type VisualShaderNodeCompareFunction int

const (
	VisualShaderNodeCompareFuncEqual            VisualShaderNodeCompareFunction = 0
	VisualShaderNodeCompareFuncGreaterThan      VisualShaderNodeCompareFunction = 2
	VisualShaderNodeCompareFuncGreaterThanEqual VisualShaderNodeCompareFunction = 3
	VisualShaderNodeCompareFuncLessThan         VisualShaderNodeCompareFunction = 4
	VisualShaderNodeCompareFuncLessThanEqual    VisualShaderNodeCompareFunction = 5
	VisualShaderNodeCompareFuncNotEqual         VisualShaderNodeCompareFunction = 1
)

//func NewVisualShaderNodeCompareFromPointer(ptr gdnative.Pointer) VisualShaderNodeCompare {
func newVisualShaderNodeCompareFromPointer(ptr gdnative.Pointer) VisualShaderNodeCompare {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeCompare{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Compares [code]a[/code] and [code]b[/code] of [member type] by [member function]. Returns a boolean scalar. Translates to [code]if[/code] instruction in shader code.
*/
type VisualShaderNodeCompare struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeCompare) BaseClass() string {
	return "VisualShaderNodeCompare"
}

/*
        Undocumented
	Args: [], Returns: enum.VisualShaderNodeCompare::ComparisonType
*/
func (o *VisualShaderNodeCompare) GetComparisonType() VisualShaderNodeCompareComparisonType {
	//log.Println("Calling VisualShaderNodeCompare.GetComparisonType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeCompare", "get_comparison_type")

	// Call the parent method.
	// enum.VisualShaderNodeCompare::ComparisonType
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualShaderNodeCompareComparisonType(ret)
}

/*
        Undocumented
	Args: [], Returns: enum.VisualShaderNodeCompare::Condition
*/
func (o *VisualShaderNodeCompare) GetCondition() VisualShaderNodeCompareCondition {
	//log.Println("Calling VisualShaderNodeCompare.GetCondition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeCompare", "get_condition")

	// Call the parent method.
	// enum.VisualShaderNodeCompare::Condition
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualShaderNodeCompareCondition(ret)
}

/*
        Undocumented
	Args: [], Returns: enum.VisualShaderNodeCompare::Function
*/
func (o *VisualShaderNodeCompare) GetFunction() VisualShaderNodeCompareFunction {
	//log.Println("Calling VisualShaderNodeCompare.GetFunction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeCompare", "get_function")

	// Call the parent method.
	// enum.VisualShaderNodeCompare::Function
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualShaderNodeCompareFunction(ret)
}

/*
        Undocumented
	Args: [{ false type int}], Returns: void
*/
func (o *VisualShaderNodeCompare) SetComparisonType(aType gdnative.Int) {
	//log.Println("Calling VisualShaderNodeCompare.SetComparisonType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeCompare", "set_comparison_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false condition int}], Returns: void
*/
func (o *VisualShaderNodeCompare) SetCondition(condition gdnative.Int) {
	//log.Println("Calling VisualShaderNodeCompare.SetCondition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(condition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeCompare", "set_condition")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false func int}], Returns: void
*/
func (o *VisualShaderNodeCompare) SetFunction(function gdnative.Int) {
	//log.Println("Calling VisualShaderNodeCompare.SetFunction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(function)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeCompare", "set_function")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderNodeCompareImplementer is an interface that implements the methods
// of the VisualShaderNodeCompare class.
type VisualShaderNodeCompareImplementer interface {
	VisualShaderNodeImplementer
	SetComparisonType(aType gdnative.Int)
	SetCondition(condition gdnative.Int)
	SetFunction(function gdnative.Int)
}
