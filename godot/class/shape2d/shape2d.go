package shape2d

import (
	"log"
	"reflect"
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
Base class for all 2D Shapes. All 2D shape types inherit from this.
*/
type Shape2D struct {
	Resource
}

func (o *Shape2D) BaseClass() string {
	return "Shape2D"
}

/*
   Return whether this shape is colliding with another. This method needs the transformation matrix for this shape ([code]local_xform[/code]), the shape to check collisions with ([code]with_shape[/code]), and the transformation matrix of that shape ([code]shape_xform[/code]).
*/
func (o *Shape2D) Collide(localXform *Transform2D, withShape *Shape2D, shapeXform *Transform2D) gdnative.Bool {
	log.Println("Calling Shape2D.Collide()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(localXform)
	goArguments[1] = reflect.ValueOf(withShape)
	goArguments[2] = reflect.ValueOf(shapeXform)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "collide", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return a list of the points where this shape touches another. If there are no collisions, the list is empty. This method needs the transformation matrix for this shape ([code]local_xform[/code]), the shape to check collisions with ([code]with_shape[/code]), and the transformation matrix of that shape ([code]shape_xform[/code]).
*/
func (o *Shape2D) CollideAndGetContacts(localXform *Transform2D, withShape *Shape2D, shapeXform *Transform2D) *Variant {
	log.Println("Calling Shape2D.CollideAndGetContacts()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(localXform)
	goArguments[1] = reflect.ValueOf(withShape)
	goArguments[2] = reflect.ValueOf(shapeXform)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "collide_and_get_contacts", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether this shape would collide with another, if a given movement was applied. This method needs the transformation matrix for this shape ([code]local_xform[/code]), the movement to test on this shape ([code]local_motion[/code]), the shape to check collisions with ([code]with_shape[/code]), the transformation matrix of that shape ([code]shape_xform[/code]), and the movement to test onto the other object ([code]shape_motion[/code]).
*/
func (o *Shape2D) CollideWithMotion(localXform *Transform2D, localMotion *Vector2, withShape *Shape2D, shapeXform *Transform2D, shapeMotion *Vector2) gdnative.Bool {
	log.Println("Calling Shape2D.CollideWithMotion()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(localXform)
	goArguments[1] = reflect.ValueOf(localMotion)
	goArguments[2] = reflect.ValueOf(withShape)
	goArguments[3] = reflect.ValueOf(shapeXform)
	goArguments[4] = reflect.ValueOf(shapeMotion)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "collide_with_motion", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return a list of the points where this shape would touch another, if a given movement was applied. If there are no collisions, the list is empty. This method needs the transformation matrix for this shape ([code]local_xform[/code]), the movement to test on this shape ([code]local_motion[/code]), the shape to check collisions with ([code]with_shape[/code]), the transformation matrix of that shape ([code]shape_xform[/code]), and the movement to test onto the other object ([code]shape_motion[/code]).
*/
func (o *Shape2D) CollideWithMotionAndGetContacts(localXform *Transform2D, localMotion *Vector2, withShape *Shape2D, shapeXform *Transform2D, shapeMotion *Vector2) *Variant {
	log.Println("Calling Shape2D.CollideWithMotionAndGetContacts()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(localXform)
	goArguments[1] = reflect.ValueOf(localMotion)
	goArguments[2] = reflect.ValueOf(withShape)
	goArguments[3] = reflect.ValueOf(shapeXform)
	goArguments[4] = reflect.ValueOf(shapeMotion)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "collide_with_motion_and_get_contacts", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Shape2D) GetCustomSolverBias() gdnative.Float {
	log.Println("Calling Shape2D.GetCustomSolverBias()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_custom_solver_bias", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Shape2D) SetCustomSolverBias(bias gdnative.Float) {
	log.Println("Calling Shape2D.SetCustomSolverBias()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bias)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_custom_solver_bias", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Shape2DImplementer is an interface for Shape2D objects.
*/
type Shape2DImplementer interface {
	Class
}
