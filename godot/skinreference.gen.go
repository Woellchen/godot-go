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

//func NewSkinReferenceFromPointer(ptr gdnative.Pointer) SkinReference {
func newSkinReferenceFromPointer(ptr gdnative.Pointer) SkinReference {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SkinReference{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type SkinReference struct {
	Reference
	owner gdnative.Object
}

func (o *SkinReference) BaseClass() string {
	return "SkinReference"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *SkinReference) X_SkinChanged() {
	// log.Println("Calling SkinReference.X_SkinChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkinReference", "_skin_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: RID
*/
func (o *SkinReference) GetSkeleton() gdnative.Rid {
	// log.Println("Calling SkinReference.GetSkeleton()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkinReference", "get_skeleton")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: Skin
*/
func (o *SkinReference) GetSkin() SkinImplementer {
	// log.Println("Calling SkinReference.GetSkin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkinReference", "get_skin")

	// Call the parent method.
	// Skin
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newSkinFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(SkinImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Skin" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(SkinImplementer)
	}

	return &ret
}

// SkinReferenceImplementer is an interface that implements the methods
// of the SkinReference class.
type SkinReferenceImplementer interface {
	ReferenceImplementer
	X_SkinChanged()
	GetSkeleton() gdnative.Rid
	GetSkin() SkinImplementer
}
