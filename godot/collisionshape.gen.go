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

//func NewCollisionShapeFromPointer(ptr gdnative.Pointer) CollisionShape {
func newCollisionShapeFromPointer(ptr gdnative.Pointer) CollisionShape {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CollisionShape{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Editor facility for creating and editing collision shapes in 3D space. You can use this node to represent all sorts of collision shapes, for example, add this to an [Area] to give it a detection shape, or add it to a [PhysicsBody] to create a solid object. [b]IMPORTANT[/b]: this is an Editor-only helper to create shapes, use [method CollisionObject.shape_owner_get_shape] to get the actual shape.
*/
type CollisionShape struct {
	Spatial
	owner gdnative.Object
}

func (o *CollisionShape) BaseClass() string {
	return "CollisionShape"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *CollisionShape) X_ShapeChanged() {
	// log.Println("Calling CollisionShape.X_ShapeChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape", "_shape_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *CollisionShape) X_UpdateDebugShape() {
	// log.Println("Calling CollisionShape.X_UpdateDebugShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape", "_update_debug_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Shape
*/
func (o *CollisionShape) GetShape() ShapeImplementer {
	// log.Println("Calling CollisionShape.GetShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape", "get_shape")

	// Call the parent method.
	// Shape
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newShapeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ShapeImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Shape" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ShapeImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *CollisionShape) IsDisabled() gdnative.Bool {
	// log.Println("Calling CollisionShape.IsDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape", "is_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Sets the collision shape's shape to the addition of all its convexed [MeshInstance] siblings geometry.
	Args: [], Returns: void
*/
func (o *CollisionShape) MakeConvexFromBrothers() {
	// log.Println("Calling CollisionShape.MakeConvexFromBrothers()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape", "make_convex_from_brothers")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If this method exists within a script it will be called whenever the shape resource has been modified.
	Args: [{ false resource Resource}], Returns: void
*/
func (o *CollisionShape) ResourceChanged(resource ResourceImplementer) {
	// log.Println("Calling CollisionShape.ResourceChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(resource.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape", "resource_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *CollisionShape) SetDisabled(enable gdnative.Bool) {
	// log.Println("Calling CollisionShape.SetDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape", "set_disabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false shape Shape}], Returns: void
*/
func (o *CollisionShape) SetShape(shape ShapeImplementer) {
	// log.Println("Calling CollisionShape.SetShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape", "set_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// CollisionShapeImplementer is an interface that implements the methods
// of the CollisionShape class.
type CollisionShapeImplementer interface {
	SpatialImplementer
	X_ShapeChanged()
	X_UpdateDebugShape()
	GetShape() ShapeImplementer
	IsDisabled() gdnative.Bool
	MakeConvexFromBrothers()
	ResourceChanged(resource ResourceImplementer)
	SetDisabled(enable gdnative.Bool)
	SetShape(shape ShapeImplementer)
}
