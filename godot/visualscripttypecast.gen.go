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

//func NewVisualScriptTypeCastFromPointer(ptr gdnative.Pointer) VisualScriptTypeCast {
func newVisualScriptTypeCastFromPointer(ptr gdnative.Pointer) VisualScriptTypeCast {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptTypeCast{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptTypeCast struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptTypeCast) BaseClass() string {
	return "VisualScriptTypeCast"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptTypeCast) GetBaseScript() gdnative.String {
	// log.Println("Calling VisualScriptTypeCast.GetBaseScript()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptTypeCast", "get_base_script")

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
	Args: [], Returns: String
*/
func (o *VisualScriptTypeCast) GetBaseType() gdnative.String {
	// log.Println("Calling VisualScriptTypeCast.GetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptTypeCast", "get_base_type")

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
	Args: [{ false path String}], Returns: void
*/
func (o *VisualScriptTypeCast) SetBaseScript(path gdnative.String) {
	// log.Println("Calling VisualScriptTypeCast.SetBaseScript()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptTypeCast", "set_base_script")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false type String}], Returns: void
*/
func (o *VisualScriptTypeCast) SetBaseType(aType gdnative.String) {
	// log.Println("Calling VisualScriptTypeCast.SetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptTypeCast", "set_base_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptTypeCastImplementer is an interface that implements the methods
// of the VisualScriptTypeCast class.
type VisualScriptTypeCastImplementer interface {
	VisualScriptNodeImplementer
	GetBaseScript() gdnative.String
	GetBaseType() gdnative.String
	SetBaseScript(path gdnative.String)
	SetBaseType(aType gdnative.String)
}
