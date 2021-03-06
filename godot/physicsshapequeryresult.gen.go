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

//func NewPhysicsShapeQueryResultFromPointer(ptr gdnative.Pointer) PhysicsShapeQueryResult {
func newPhysicsShapeQueryResultFromPointer(ptr gdnative.Pointer) PhysicsShapeQueryResult {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PhysicsShapeQueryResult{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The result of a 3D shape query in [PhysicsServer]. See also [PhysicsShapeQueryParameters].
*/
type PhysicsShapeQueryResult struct {
	Reference
	owner gdnative.Object
}

func (o *PhysicsShapeQueryResult) BaseClass() string {
	return "PhysicsShapeQueryResult"
}

/*
        Returns the number of objects that intersected with the shape.
	Args: [], Returns: int
*/
func (o *PhysicsShapeQueryResult) GetResultCount() gdnative.Int {
	// log.Println("Calling PhysicsShapeQueryResult.GetResultCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsShapeQueryResult", "get_result_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the [Object] that intersected with the shape at index [code]idx[/code].
	Args: [{ false idx int}], Returns: Object
*/
func (o *PhysicsShapeQueryResult) GetResultObject(idx gdnative.Int) ObjectImplementer {
	// log.Println("Calling PhysicsShapeQueryResult.GetResultObject()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsShapeQueryResult", "get_result_object")

	// Call the parent method.
	// Object
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newObjectFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ObjectImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Object" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ObjectImplementer)
	}

	return &ret
}

/*
        Returns the instance ID of the [Object] that intersected with the shape at index [code]idx[/code].
	Args: [{ false idx int}], Returns: int
*/
func (o *PhysicsShapeQueryResult) GetResultObjectId(idx gdnative.Int) gdnative.Int {
	// log.Println("Calling PhysicsShapeQueryResult.GetResultObjectId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsShapeQueryResult", "get_result_object_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the child index of the object's [Shape] that intersected with the shape at index [code]idx[/code].
	Args: [{ false idx int}], Returns: int
*/
func (o *PhysicsShapeQueryResult) GetResultObjectShape(idx gdnative.Int) gdnative.Int {
	// log.Println("Calling PhysicsShapeQueryResult.GetResultObjectShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsShapeQueryResult", "get_result_object_shape")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the [RID] of the object that intersected with the shape at index [code]idx[/code].
	Args: [{ false idx int}], Returns: RID
*/
func (o *PhysicsShapeQueryResult) GetResultRid(idx gdnative.Int) gdnative.Rid {
	// log.Println("Calling PhysicsShapeQueryResult.GetResultRid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsShapeQueryResult", "get_result_rid")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)
	return ret
}

// PhysicsShapeQueryResultImplementer is an interface that implements the methods
// of the PhysicsShapeQueryResult class.
type PhysicsShapeQueryResultImplementer interface {
	ReferenceImplementer
	GetResultCount() gdnative.Int
	GetResultObject(idx gdnative.Int) ObjectImplementer
	GetResultObjectId(idx gdnative.Int) gdnative.Int
	GetResultObjectShape(idx gdnative.Int) gdnative.Int
	GetResultRid(idx gdnative.Int) gdnative.Rid
}
