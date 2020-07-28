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

//func NewGDNativeFromPointer(ptr gdnative.Pointer) GDNative {
func newGDNativeFromPointer(ptr gdnative.Pointer) GDNative {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := GDNative{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type GDNative struct {
	Reference
	owner gdnative.Object
}

func (o *GDNative) BaseClass() string {
	return "GDNative"
}

/*
        Undocumented
	Args: [{ false calling_type String} { false procedure_name String} { false arguments Array}], Returns: Variant
*/
func (o *GDNative) CallNative(callingType gdnative.String, procedureName gdnative.String, arguments gdnative.Array) gdnative.Variant {
	// log.Println("Calling GDNative.CallNative()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(callingType)
	ptrArguments[1] = gdnative.NewPointerFromString(procedureName)
	ptrArguments[2] = gdnative.NewPointerFromArray(arguments)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNative", "call_native")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: GDNativeLibrary
*/
func (o *GDNative) GetLibrary() GDNativeLibraryImplementer {
	// log.Println("Calling GDNative.GetLibrary()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNative", "get_library")

	// Call the parent method.
	// GDNativeLibrary
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newGDNativeLibraryFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(GDNativeLibraryImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "GDNativeLibrary" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(GDNativeLibraryImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *GDNative) Initialize() gdnative.Bool {
	// log.Println("Calling GDNative.Initialize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNative", "initialize")

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
	Args: [{ false library GDNativeLibrary}], Returns: void
*/
func (o *GDNative) SetLibrary(library GDNativeLibraryImplementer) {
	// log.Println("Calling GDNative.SetLibrary()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(library.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNative", "set_library")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *GDNative) Terminate() gdnative.Bool {
	// log.Println("Calling GDNative.Terminate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNative", "terminate")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

// GDNativeImplementer is an interface that implements the methods
// of the GDNative class.
type GDNativeImplementer interface {
	ReferenceImplementer
	CallNative(callingType gdnative.String, procedureName gdnative.String, arguments gdnative.Array) gdnative.Variant
	GetLibrary() GDNativeLibraryImplementer
	Initialize() gdnative.Bool
	SetLibrary(library GDNativeLibraryImplementer)
	Terminate() gdnative.Bool
}
