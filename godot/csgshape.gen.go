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

// CSGShapeOperation is an enum for Operation values.
type CSGShapeOperation int

const (
	CSGShapeOperationIntersection CSGShapeOperation = 1
	CSGShapeOperationSubtraction  CSGShapeOperation = 2
	CSGShapeOperationUnion        CSGShapeOperation = 0
)

//func NewCSGShapeFromPointer(ptr gdnative.Pointer) CSGShape {
func newCSGShapeFromPointer(ptr gdnative.Pointer) CSGShape {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CSGShape{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type CSGShape struct {
	GeometryInstance
	owner gdnative.Object
}

func (o *CSGShape) BaseClass() string {
	return "CSGShape"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *CSGShape) X_UpdateShape() {
	// log.Println("Calling CSGShape.X_UpdateShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "_update_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *CSGShape) GetCollisionLayer() gdnative.Int {
	// log.Println("Calling CSGShape.GetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "get_collision_layer")

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
	Args: [{ false bit int}], Returns: bool
*/
func (o *CSGShape) GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool {
	// log.Println("Calling CSGShape.GetCollisionLayerBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "get_collision_layer_bit")

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
func (o *CSGShape) GetCollisionMask() gdnative.Int {
	// log.Println("Calling CSGShape.GetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "get_collision_mask")

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
	Args: [{ false bit int}], Returns: bool
*/
func (o *CSGShape) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	// log.Println("Calling CSGShape.GetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "get_collision_mask_bit")

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
	Args: [], Returns: Array
*/
func (o *CSGShape) GetMeshes() gdnative.Array {
	// log.Println("Calling CSGShape.GetMeshes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "get_meshes")

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
	Args: [], Returns: enum.CSGShape::Operation
*/
func (o *CSGShape) GetOperation() CSGShapeOperation {
	// log.Println("Calling CSGShape.GetOperation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "get_operation")

	// Call the parent method.
	// enum.CSGShape::Operation
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return CSGShapeOperation(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *CSGShape) GetSnap() gdnative.Real {
	// log.Println("Calling CSGShape.GetSnap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "get_snap")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *CSGShape) IsCalculatingTangents() gdnative.Bool {
	// log.Println("Calling CSGShape.IsCalculatingTangents()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "is_calculating_tangents")

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
	Args: [], Returns: bool
*/
func (o *CSGShape) IsRootShape() gdnative.Bool {
	// log.Println("Calling CSGShape.IsRootShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "is_root_shape")

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
	Args: [], Returns: bool
*/
func (o *CSGShape) IsUsingCollision() gdnative.Bool {
	// log.Println("Calling CSGShape.IsUsingCollision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "is_using_collision")

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
	Args: [{ false enabled bool}], Returns: void
*/
func (o *CSGShape) SetCalculateTangents(enabled gdnative.Bool) {
	// log.Println("Calling CSGShape.SetCalculateTangents()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "set_calculate_tangents")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false layer int}], Returns: void
*/
func (o *CSGShape) SetCollisionLayer(layer gdnative.Int) {
	// log.Println("Calling CSGShape.SetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(layer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "set_collision_layer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *CSGShape) SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool) {
	// log.Println("Calling CSGShape.SetCollisionLayerBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "set_collision_layer_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *CSGShape) SetCollisionMask(mask gdnative.Int) {
	// log.Println("Calling CSGShape.SetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "set_collision_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *CSGShape) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	// log.Println("Calling CSGShape.SetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "set_collision_mask_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false operation int}], Returns: void
*/
func (o *CSGShape) SetOperation(operation gdnative.Int) {
	// log.Println("Calling CSGShape.SetOperation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(operation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "set_operation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false snap float}], Returns: void
*/
func (o *CSGShape) SetSnap(snap gdnative.Real) {
	// log.Println("Calling CSGShape.SetSnap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(snap)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "set_snap")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false operation bool}], Returns: void
*/
func (o *CSGShape) SetUseCollision(operation gdnative.Bool) {
	// log.Println("Calling CSGShape.SetUseCollision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(operation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGShape", "set_use_collision")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// CSGShapeImplementer is an interface that implements the methods
// of the CSGShape class.
type CSGShapeImplementer interface {
	GeometryInstanceImplementer
	X_UpdateShape()
	GetCollisionLayer() gdnative.Int
	GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool
	GetCollisionMask() gdnative.Int
	GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool
	GetMeshes() gdnative.Array
	GetSnap() gdnative.Real
	IsCalculatingTangents() gdnative.Bool
	IsRootShape() gdnative.Bool
	IsUsingCollision() gdnative.Bool
	SetCalculateTangents(enabled gdnative.Bool)
	SetCollisionLayer(layer gdnative.Int)
	SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool)
	SetCollisionMask(mask gdnative.Int)
	SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool)
	SetOperation(operation gdnative.Int)
	SetSnap(snap gdnative.Real)
	SetUseCollision(operation gdnative.Bool)
}
