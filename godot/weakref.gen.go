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

//func NewWeakRefFromPointer(ptr gdnative.Pointer) WeakRef {
func newWeakRefFromPointer(ptr gdnative.Pointer) WeakRef {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := WeakRef{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A weakref can hold a [Reference], without contributing to the reference counter. A weakref can be created from an [Object] using [method @GDScript.weakref]. If this object is not a reference, weakref still works, however, it does not have any effect on the object. Weakrefs are useful in cases where multiple classes have variables that refer to each other. Without weakrefs, using these classes could lead to memory leaks, since both references keep each other from being released. Making part of the variables a weakref can prevent this cyclic dependency, and allows the references to be released.
*/
type WeakRef struct {
	Reference
	owner gdnative.Object
}

func (o *WeakRef) BaseClass() string {
	return "WeakRef"
}

/*
        Returns the [Object] this weakref is referring to.
	Args: [], Returns: Variant
*/
func (o *WeakRef) GetRef() gdnative.Variant {
	//log.Println("Calling WeakRef.GetRef()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WeakRef", "get_ref")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

// WeakRefImplementer is an interface that implements the methods
// of the WeakRef class.
type WeakRefImplementer interface {
	ReferenceImplementer
	GetRef() gdnative.Variant
}
