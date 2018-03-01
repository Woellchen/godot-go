package circleshape2d

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
Circular shape for 2D collisions. This shape is useful for modeling balls or small characters and its collision detection with everything else is very fast.
*/
type CircleShape2D struct {
	Shape2D
}

func (o *CircleShape2D) BaseClass() string {
	return "CircleShape2D"
}

/*
   Undocumented
*/
func (o *CircleShape2D) GetRadius() gdnative.Float {
	log.Println("Calling CircleShape2D.GetRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_radius", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *CircleShape2D) SetRadius(radius gdnative.Float) {
	log.Println("Calling CircleShape2D.SetRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(radius)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_radius", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   CircleShape2DImplementer is an interface for CircleShape2D objects.
*/
type CircleShape2DImplementer interface {
	Class
}
