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

//func NewPhysicsBodyFromPointer(ptr gdnative.Pointer) PhysicsBody {
func newPhysicsBodyFromPointer(ptr gdnative.Pointer) PhysicsBody {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PhysicsBody{}
	obj.SetBaseObject(owner)

	return obj
}

/*
PhysicsBody is an abstract base class for implementing a physics body. All *Body types inherit from it.
*/
type PhysicsBody struct {
	CollisionObject
	owner gdnative.Object
}

func (o *PhysicsBody) BaseClass() string {
	return "PhysicsBody"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *PhysicsBody) X_GetLayers() gdnative.Int {
	//log.Println("Calling PhysicsBody.X_GetLayers()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "_get_layers")

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
	Args: [{ false mask int}], Returns: void
*/
func (o *PhysicsBody) X_SetLayers(mask gdnative.Int) {
	//log.Println("Calling PhysicsBody.X_SetLayers()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "_set_layers")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a body to the list of bodies that this body can't collide with.
	Args: [{ false body Node}], Returns: void
*/
func (o *PhysicsBody) AddCollisionExceptionWith(body NodeImplementer) {
	//log.Println("Calling PhysicsBody.AddCollisionExceptionWith()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(body.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "add_collision_exception_with")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns an array of nodes that were added as collision exceptions for this body.
	Args: [], Returns: Array
*/
func (o *PhysicsBody) GetCollisionExceptions() gdnative.Array {
	//log.Println("Calling PhysicsBody.GetCollisionExceptions()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "get_collision_exceptions")

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
	Args: [], Returns: int
*/
func (o *PhysicsBody) GetCollisionLayer() gdnative.Int {
	//log.Println("Calling PhysicsBody.GetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "get_collision_layer")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns an individual bit on the [member collision_layer].
	Args: [{ false bit int}], Returns: bool
*/
func (o *PhysicsBody) GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool {
	//log.Println("Calling PhysicsBody.GetCollisionLayerBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "get_collision_layer_bit")

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
	Args: [], Returns: int
*/
func (o *PhysicsBody) GetCollisionMask() gdnative.Int {
	//log.Println("Calling PhysicsBody.GetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "get_collision_mask")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns an individual bit on the [member collision_mask].
	Args: [{ false bit int}], Returns: bool
*/
func (o *PhysicsBody) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	//log.Println("Calling PhysicsBody.GetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "get_collision_mask_bit")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes a body from the list of bodies that this body can't collide with.
	Args: [{ false body Node}], Returns: void
*/
func (o *PhysicsBody) RemoveCollisionExceptionWith(body NodeImplementer) {
	//log.Println("Calling PhysicsBody.RemoveCollisionExceptionWith()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(body.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "remove_collision_exception_with")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false layer int}], Returns: void
*/
func (o *PhysicsBody) SetCollisionLayer(layer gdnative.Int) {
	//log.Println("Calling PhysicsBody.SetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(layer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "set_collision_layer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets individual bits on the [member collision_layer] bitmask. Use this if you only need to change one layer's value.
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *PhysicsBody) SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling PhysicsBody.SetCollisionLayerBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "set_collision_layer_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *PhysicsBody) SetCollisionMask(mask gdnative.Int) {
	//log.Println("Calling PhysicsBody.SetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "set_collision_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets individual bits on the [member collision_mask] bitmask. Use this if you only need to change one layer's value.
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *PhysicsBody) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling PhysicsBody.SetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody", "set_collision_mask_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// PhysicsBodyImplementer is an interface that implements the methods
// of the PhysicsBody class.
type PhysicsBodyImplementer interface {
	CollisionObjectImplementer
	X_GetLayers() gdnative.Int
	X_SetLayers(mask gdnative.Int)
	AddCollisionExceptionWith(body NodeImplementer)
	GetCollisionExceptions() gdnative.Array
	GetCollisionLayer() gdnative.Int
	GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool
	GetCollisionMask() gdnative.Int
	GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool
	RemoveCollisionExceptionWith(body NodeImplementer)
	SetCollisionLayer(layer gdnative.Int)
	SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool)
	SetCollisionMask(mask gdnative.Int)
	SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool)
}
