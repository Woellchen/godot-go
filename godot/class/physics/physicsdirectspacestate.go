package physics

import (
	"log"
	"reflect"

	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
Direct access object to a space in the [PhysicsServer]. It's used mainly to do queries against objects and areas residing in a given space.
*/
type PhysicsDirectSpaceState struct {
	Object
}

func (o *PhysicsDirectSpaceState) BaseClass() string {
	return "PhysicsDirectSpaceState"
}

/*
   Checks whether the shape can travel to a point. The method will return an array with two floats between 0 and 1, both representing a fraction of [code]motion[/code]. The first is how far the shape can move without triggering a collision, and the second is the point at which a collision will occur. If no collision is detected, the returned array will be [1, 1]. If the shape can not move, the array will be empty ([code]dir.empty()==true[/code]).
*/
func (o *PhysicsDirectSpaceState) CastMotion(shape *PhysicsShapeQueryParameters, motion *Vector3) *Array {
	log.Println("Calling PhysicsDirectSpaceState.CastMotion()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(shape)
	goArguments[1] = reflect.ValueOf(motion)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "cast_motion", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters] object, against the space. The resulting array contains a list of points where the shape intersects another. Like with [method intersect_shape], the number of returned results can be limited to save processing time.
*/
func (o *PhysicsDirectSpaceState) CollideShape(shape *PhysicsShapeQueryParameters, maxResults gdnative.Int) *Array {
	log.Println("Calling PhysicsDirectSpaceState.CollideShape()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(shape)
	goArguments[1] = reflect.ValueOf(maxResults)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "collide_shape", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters] object, against the space. If it collides with more than a shape, the nearest one is selected. The returned object is a dictionary containing the following fields: [code]collider_id[/code]: The colliding object's ID. [code]linear_velocity[/code]: The colliding object's velocity [Vector3]. If the object is an [Area], the result is [code](0, 0, 0)[/code]. [code]normal[/code]: The object's surface normal at the intersection point. [code]point[/code]: The intersection point. [code]rid[/code]: The intersecting object's [RID]. [code]shape[/code]: The shape index of the colliding shape. If the shape did not intersect anything, then an empty dictionary ([code]dir.empty()==true[/code]) is returned instead.
*/
func (o *PhysicsDirectSpaceState) GetRestInfo(shape *PhysicsShapeQueryParameters) *Dictionary {
	log.Println("Calling PhysicsDirectSpaceState.GetRestInfo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(shape)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_rest_info", goArguments, "*Dictionary")

	returnValue := goRet.Interface().(*Dictionary)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Intersects a ray in a given space. The returned object is a dictionary with the following fields: [code]collider[/code]: The colliding object. [code]collider_id[/code]: The colliding object's ID. [code]normal[/code]: The object's surface normal at the intersection point. [code]position[/code]: The intersection point. [code]rid[/code]: The intersecting object's [RID]. [code]shape[/code]: The shape index of the colliding shape. If the ray did not intersect anything, then an empty dictionary ([code]dir.empty()==true[/code]) is returned instead. Additionally, the method can take an array of objects or [RID]s that are to be excluded from collisions, or a bitmask representing the physics layers to check in.
*/
func (o *PhysicsDirectSpaceState) IntersectRay(from *Vector3, to *Vector3, exclude *Array, collisionLayer gdnative.Int) *Dictionary {
	log.Println("Calling PhysicsDirectSpaceState.IntersectRay()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(from)
	goArguments[1] = reflect.ValueOf(to)
	goArguments[2] = reflect.ValueOf(exclude)
	goArguments[3] = reflect.ValueOf(collisionLayer)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "intersect_ray", goArguments, "*Dictionary")

	returnValue := goRet.Interface().(*Dictionary)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters] object, against the space. The intersected shapes are returned in an array containing dictionaries with the following fields: [code]collider[/code]: The colliding object. [code]collider_id[/code]: The colliding object's ID. [code]rid[/code]: The intersecting object's [RID]. [code]shape[/code]: The shape index of the colliding shape. The number of intersections can be limited with the second parameter, to reduce the processing time.
*/
func (o *PhysicsDirectSpaceState) IntersectShape(shape *PhysicsShapeQueryParameters, maxResults gdnative.Int) *Array {
	log.Println("Calling PhysicsDirectSpaceState.IntersectShape()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(shape)
	goArguments[1] = reflect.ValueOf(maxResults)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "intersect_shape", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   PhysicsDirectSpaceStateImplementer is an interface for PhysicsDirectSpaceState objects.
*/
type PhysicsDirectSpaceStateImplementer interface {
	Class
}
