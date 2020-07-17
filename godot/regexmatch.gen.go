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

//func NewRegExMatchFromPointer(ptr gdnative.Pointer) RegExMatch {
func newRegExMatchFromPointer(ptr gdnative.Pointer) RegExMatch {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := RegExMatch{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type RegExMatch struct {
	Reference
	owner gdnative.Object
}

func (o *RegExMatch) BaseClass() string {
	return "RegExMatch"
}

/*
        Undocumented
	Args: [{0 true name Variant}], Returns: int
*/
func (o *RegExMatch) GetEnd(name gdnative.Variant) gdnative.Int {
	//log.Println("Calling RegExMatch.GetEnd()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVariant(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RegExMatch", "get_end")

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
	Args: [], Returns: int
*/
func (o *RegExMatch) GetGroupCount() gdnative.Int {
	//log.Println("Calling RegExMatch.GetGroupCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RegExMatch", "get_group_count")

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
	Args: [], Returns: Dictionary
*/
func (o *RegExMatch) GetNames() gdnative.Dictionary {
	//log.Println("Calling RegExMatch.GetNames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RegExMatch", "get_names")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{0 true name Variant}], Returns: int
*/
func (o *RegExMatch) GetStart(name gdnative.Variant) gdnative.Int {
	//log.Println("Calling RegExMatch.GetStart()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVariant(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RegExMatch", "get_start")

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
	Args: [{0 true name Variant}], Returns: String
*/
func (o *RegExMatch) GetString(name gdnative.Variant) gdnative.String {
	//log.Println("Calling RegExMatch.GetString()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVariant(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RegExMatch", "get_string")

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
	Args: [], Returns: Array
*/
func (o *RegExMatch) GetStrings() gdnative.Array {
	//log.Println("Calling RegExMatch.GetStrings()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RegExMatch", "get_strings")

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
	Args: [], Returns: String
*/
func (o *RegExMatch) GetSubject() gdnative.String {
	//log.Println("Calling RegExMatch.GetSubject()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RegExMatch", "get_subject")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

// RegExMatchImplementer is an interface that implements the methods
// of the RegExMatch class.
type RegExMatchImplementer interface {
	ReferenceImplementer
	GetEnd(name gdnative.Variant) gdnative.Int
	GetGroupCount() gdnative.Int
	GetNames() gdnative.Dictionary
	GetStart(name gdnative.Variant) gdnative.Int
	GetString(name gdnative.Variant) gdnative.String
	GetStrings() gdnative.Array
	GetSubject() gdnative.String
}
