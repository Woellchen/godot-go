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

//func NewVisualScriptDeconstructFromPointer(ptr gdnative.Pointer) VisualScriptDeconstruct {
func newVisualScriptDeconstructFromPointer(ptr gdnative.Pointer) VisualScriptDeconstruct {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptDeconstruct{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptDeconstruct struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptDeconstruct) BaseClass() string {
	return "VisualScriptDeconstruct"
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *VisualScriptDeconstruct) X_GetElemCache() gdnative.Array {
	// log.Println("Calling VisualScriptDeconstruct.X_GetElemCache()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptDeconstruct", "_get_elem_cache")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false _cache Array}], Returns: void
*/
func (o *VisualScriptDeconstruct) X_SetElemCache(cache gdnative.Array) {
	// log.Println("Calling VisualScriptDeconstruct.X_SetElemCache()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(cache)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptDeconstruct", "_set_elem_cache")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: enum.Variant::Type
*/
func (o *VisualScriptDeconstruct) GetDeconstructType() gdnative.VariantType {
	// log.Println("Calling VisualScriptDeconstruct.GetDeconstructType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptDeconstruct", "get_deconstruct_type")

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
	Args: [{ false type int}], Returns: void
*/
func (o *VisualScriptDeconstruct) SetDeconstructType(aType gdnative.Int) {
	// log.Println("Calling VisualScriptDeconstruct.SetDeconstructType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptDeconstruct", "set_deconstruct_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptDeconstructImplementer is an interface that implements the methods
// of the VisualScriptDeconstruct class.
type VisualScriptDeconstructImplementer interface {
	VisualScriptNodeImplementer
	X_GetElemCache() gdnative.Array
	X_SetElemCache(cache gdnative.Array)
	SetDeconstructType(aType gdnative.Int)
}
