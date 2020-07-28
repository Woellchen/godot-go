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

//func NewVisualScriptBasicTypeConstantFromPointer(ptr gdnative.Pointer) VisualScriptBasicTypeConstant {
func newVisualScriptBasicTypeConstantFromPointer(ptr gdnative.Pointer) VisualScriptBasicTypeConstant {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptBasicTypeConstant{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptBasicTypeConstant struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptBasicTypeConstant) BaseClass() string {
	return "VisualScriptBasicTypeConstant"
}

/*
        Undocumented
	Args: [], Returns: enum.Variant::Type
*/
func (o *VisualScriptBasicTypeConstant) GetBasicType() gdnative.VariantType {
	// log.Println("Calling VisualScriptBasicTypeConstant.GetBasicType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptBasicTypeConstant", "get_basic_type")

	// Call the parent method.
	// enum.Variant::Type
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.VariantType(ret)
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptBasicTypeConstant) GetBasicTypeConstant() gdnative.String {
	// log.Println("Calling VisualScriptBasicTypeConstant.GetBasicTypeConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptBasicTypeConstant", "get_basic_type_constant")

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
	Args: [{ false name int}], Returns: void
*/
func (o *VisualScriptBasicTypeConstant) SetBasicType(name gdnative.Int) {
	// log.Println("Calling VisualScriptBasicTypeConstant.SetBasicType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptBasicTypeConstant", "set_basic_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false name String}], Returns: void
*/
func (o *VisualScriptBasicTypeConstant) SetBasicTypeConstant(name gdnative.String) {
	// log.Println("Calling VisualScriptBasicTypeConstant.SetBasicTypeConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptBasicTypeConstant", "set_basic_type_constant")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptBasicTypeConstantImplementer is an interface that implements the methods
// of the VisualScriptBasicTypeConstant class.
type VisualScriptBasicTypeConstantImplementer interface {
	VisualScriptNodeImplementer
	GetBasicTypeConstant() gdnative.String
	SetBasicType(name gdnative.Int)
	SetBasicTypeConstant(name gdnative.String)
}
