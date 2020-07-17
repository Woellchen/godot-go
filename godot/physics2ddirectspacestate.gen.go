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

//func NewPhysics2DDirectSpaceStateFromPointer(ptr gdnative.Pointer) Physics2DDirectSpaceState {
func newPhysics2DDirectSpaceStateFromPointer(ptr gdnative.Pointer) Physics2DDirectSpaceState {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Physics2DDirectSpaceState{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Direct access object to a space in the [Physics2DServer]. It's used mainly to do queries against objects and areas residing in a given space.
*/
type Physics2DDirectSpaceState struct {
	Object
	owner gdnative.Object
}

func (o *Physics2DDirectSpaceState) BaseClass() string {
	return "Physics2DDirectSpaceState"
}

/*
        Checks how far the shape can travel toward a point. If the shape can not move, the array will be empty. [b]Note:[/b] Both the shape and the motion are supplied through a [Physics2DShapeQueryParameters] object. The method will return an array with two floats between 0 and 1, both representing a fraction of [code]motion[/code]. The first is how far the shape can move without triggering a collision, and the second is the point at which a collision will occur. If no collision is detected, the returned array will be [code][1, 1][/code].
	Args: [{ false shape Physics2DShapeQueryParameters}], Returns: Array
*/
func (o *Physics2DDirectSpaceState) CastMotion(shape Physics2DShapeQueryParametersImplementer) gdnative.Array {
	//log.Println("Calling Physics2DDirectSpaceState.CastMotion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectSpaceState", "cast_motion")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Checks the intersections of a shape, given through a [Physics2DShapeQueryParameters] object, against the space. The resulting array contains a list of points where the shape intersects another. Like with [method intersect_shape], the number of returned results can be limited to save processing time.
	Args: [{ false shape Physics2DShapeQueryParameters} {32 true max_results int}], Returns: Array
*/
func (o *Physics2DDirectSpaceState) CollideShape(shape Physics2DShapeQueryParametersImplementer, maxResults gdnative.Int) gdnative.Array {
	//log.Println("Calling Physics2DDirectSpaceState.CollideShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(maxResults)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectSpaceState", "collide_shape")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Checks the intersections of a shape, given through a [Physics2DShapeQueryParameters] object, against the space. If it collides with more than one shape, the nearest one is selected. If the shape did not intersect anything, then an empty dictionary is returned instead. [b]Note:[/b] This method does not take into account the [code]motion[/code] property of the object. The returned object is a dictionary containing the following fields: [code]collider_id[/code]: The colliding object's ID. [code]linear_velocity[/code]: The colliding object's velocity [Vector2]. If the object is an [Area2D], the result is [code](0, 0)[/code]. [code]metadata[/code]: The intersecting shape's metadata. This metadata is different from [method Object.get_meta], and is set with [method Physics2DServer.shape_set_data]. [code]normal[/code]: The object's surface normal at the intersection point. [code]point[/code]: The intersection point. [code]rid[/code]: The intersecting object's [RID]. [code]shape[/code]: The shape index of the colliding shape.
	Args: [{ false shape Physics2DShapeQueryParameters}], Returns: Dictionary
*/
func (o *Physics2DDirectSpaceState) GetRestInfo(shape Physics2DShapeQueryParametersImplementer) gdnative.Dictionary {
	//log.Println("Calling Physics2DDirectSpaceState.GetRestInfo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectSpaceState", "get_rest_info")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Checks whether a point is inside any shape. The shapes the point is inside of are returned in an array containing dictionaries with the following fields: [code]collider[/code]: The colliding object. [code]collider_id[/code]: The colliding object's ID. [code]metadata[/code]: The intersecting shape's metadata. This metadata is different from [method Object.get_meta], and is set with [method Physics2DServer.shape_set_data]. [code]rid[/code]: The intersecting object's [RID]. [code]shape[/code]: The shape index of the colliding shape. Additionally, the method can take an [code]exclude[/code] array of objects or [RID]s that are to be excluded from collisions, a [code]collision_mask[/code] bitmask representing the physics layers to check in, or booleans to determine if the ray should collide with [PhysicsBody]s or [Area]s, respectively.
	Args: [{ false point Vector2} {32 true max_results int} {[] true exclude Array} {2147483647 true collision_layer int} {True true collide_with_bodies bool} {False true collide_with_areas bool}], Returns: Array
*/
func (o *Physics2DDirectSpaceState) IntersectPoint(point gdnative.Vector2, maxResults gdnative.Int, exclude gdnative.Array, collisionLayer gdnative.Int, collideWithBodies gdnative.Bool, collideWithAreas gdnative.Bool) gdnative.Array {
	//log.Println("Calling Physics2DDirectSpaceState.IntersectPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 6, 6)
	ptrArguments[0] = gdnative.NewPointerFromVector2(point)
	ptrArguments[1] = gdnative.NewPointerFromInt(maxResults)
	ptrArguments[2] = gdnative.NewPointerFromArray(exclude)
	ptrArguments[3] = gdnative.NewPointerFromInt(collisionLayer)
	ptrArguments[4] = gdnative.NewPointerFromBool(collideWithBodies)
	ptrArguments[5] = gdnative.NewPointerFromBool(collideWithAreas)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectSpaceState", "intersect_point")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false point Vector2} { false canvas_instance_id int} {32 true max_results int} {[] true exclude Array} {2147483647 true collision_layer int} {True true collide_with_bodies bool} {False true collide_with_areas bool}], Returns: Array
*/
func (o *Physics2DDirectSpaceState) IntersectPointOnCanvas(point gdnative.Vector2, canvasInstanceId gdnative.Int, maxResults gdnative.Int, exclude gdnative.Array, collisionLayer gdnative.Int, collideWithBodies gdnative.Bool, collideWithAreas gdnative.Bool) gdnative.Array {
	//log.Println("Calling Physics2DDirectSpaceState.IntersectPointOnCanvas()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 7, 7)
	ptrArguments[0] = gdnative.NewPointerFromVector2(point)
	ptrArguments[1] = gdnative.NewPointerFromInt(canvasInstanceId)
	ptrArguments[2] = gdnative.NewPointerFromInt(maxResults)
	ptrArguments[3] = gdnative.NewPointerFromArray(exclude)
	ptrArguments[4] = gdnative.NewPointerFromInt(collisionLayer)
	ptrArguments[5] = gdnative.NewPointerFromBool(collideWithBodies)
	ptrArguments[6] = gdnative.NewPointerFromBool(collideWithAreas)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectSpaceState", "intersect_point_on_canvas")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Intersects a ray in a given space. The returned object is a dictionary with the following fields: [code]collider[/code]: The colliding object. [code]collider_id[/code]: The colliding object's ID. [code]metadata[/code]: The intersecting shape's metadata. This metadata is different from [method Object.get_meta], and is set with [method Physics2DServer.shape_set_data]. [code]normal[/code]: The object's surface normal at the intersection point. [code]position[/code]: The intersection point. [code]rid[/code]: The intersecting object's [RID]. [code]shape[/code]: The shape index of the colliding shape. If the ray did not intersect anything, then an empty dictionary is returned instead. Additionally, the method can take an [code]exclude[/code] array of objects or [RID]s that are to be excluded from collisions, a [code]collision_mask[/code] bitmask representing the physics layers to check in, or booleans to determine if the ray should collide with [PhysicsBody]s or [Area]s, respectively.
	Args: [{ false from Vector2} { false to Vector2} {[] true exclude Array} {2147483647 true collision_layer int} {True true collide_with_bodies bool} {False true collide_with_areas bool}], Returns: Dictionary
*/
func (o *Physics2DDirectSpaceState) IntersectRay(from gdnative.Vector2, to gdnative.Vector2, exclude gdnative.Array, collisionLayer gdnative.Int, collideWithBodies gdnative.Bool, collideWithAreas gdnative.Bool) gdnative.Dictionary {
	//log.Println("Calling Physics2DDirectSpaceState.IntersectRay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 6, 6)
	ptrArguments[0] = gdnative.NewPointerFromVector2(from)
	ptrArguments[1] = gdnative.NewPointerFromVector2(to)
	ptrArguments[2] = gdnative.NewPointerFromArray(exclude)
	ptrArguments[3] = gdnative.NewPointerFromInt(collisionLayer)
	ptrArguments[4] = gdnative.NewPointerFromBool(collideWithBodies)
	ptrArguments[5] = gdnative.NewPointerFromBool(collideWithAreas)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectSpaceState", "intersect_ray")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Checks the intersections of a shape, given through a [Physics2DShapeQueryParameters] object, against the space. [b]Note:[/b] This method does not take into account the [code]motion[/code] property of the object. The intersected shapes are returned in an array containing dictionaries with the following fields: [code]collider[/code]: The colliding object. [code]collider_id[/code]: The colliding object's ID. [code]metadata[/code]: The intersecting shape's metadata. This metadata is different from [method Object.get_meta], and is set with [method Physics2DServer.shape_set_data]. [code]rid[/code]: The intersecting object's [RID]. [code]shape[/code]: The shape index of the colliding shape. The number of intersections can be limited with the [code]max_results[/code] parameter, to reduce the processing time.
	Args: [{ false shape Physics2DShapeQueryParameters} {32 true max_results int}], Returns: Array
*/
func (o *Physics2DDirectSpaceState) IntersectShape(shape Physics2DShapeQueryParametersImplementer, maxResults gdnative.Int) gdnative.Array {
	//log.Println("Calling Physics2DDirectSpaceState.IntersectShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(maxResults)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectSpaceState", "intersect_shape")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

// Physics2DDirectSpaceStateImplementer is an interface that implements the methods
// of the Physics2DDirectSpaceState class.
type Physics2DDirectSpaceStateImplementer interface {
	ObjectImplementer
	CastMotion(shape Physics2DShapeQueryParametersImplementer) gdnative.Array
	CollideShape(shape Physics2DShapeQueryParametersImplementer, maxResults gdnative.Int) gdnative.Array
	GetRestInfo(shape Physics2DShapeQueryParametersImplementer) gdnative.Dictionary
	IntersectPoint(point gdnative.Vector2, maxResults gdnative.Int, exclude gdnative.Array, collisionLayer gdnative.Int, collideWithBodies gdnative.Bool, collideWithAreas gdnative.Bool) gdnative.Array
	IntersectPointOnCanvas(point gdnative.Vector2, canvasInstanceId gdnative.Int, maxResults gdnative.Int, exclude gdnative.Array, collisionLayer gdnative.Int, collideWithBodies gdnative.Bool, collideWithAreas gdnative.Bool) gdnative.Array
	IntersectRay(from gdnative.Vector2, to gdnative.Vector2, exclude gdnative.Array, collisionLayer gdnative.Int, collideWithBodies gdnative.Bool, collideWithAreas gdnative.Bool) gdnative.Dictionary
	IntersectShape(shape Physics2DShapeQueryParametersImplementer, maxResults gdnative.Int) gdnative.Array
}
