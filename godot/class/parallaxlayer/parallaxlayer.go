package parallaxlayer

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
A ParallaxLayer must be the child of a [ParallaxBackground] node. Each ParallaxLayer can be set to move at different speeds relative to the camera movement or the [member ParallaxBackground.scroll_offset] value. This node's children will be affected by its scroll offset.
*/
type ParallaxLayer struct {
	Node2D
}

func (o *ParallaxLayer) BaseClass() string {
	return "ParallaxLayer"
}

/*
   Undocumented
*/
func (o *ParallaxLayer) GetMirroring() *Vector2 {
	log.Println("Calling ParallaxLayer.GetMirroring()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mirroring", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ParallaxLayer) GetMotionOffset() *Vector2 {
	log.Println("Calling ParallaxLayer.GetMotionOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_motion_offset", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ParallaxLayer) GetMotionScale() *Vector2 {
	log.Println("Calling ParallaxLayer.GetMotionScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_motion_scale", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ParallaxLayer) SetMirroring(mirror *Vector2) {
	log.Println("Calling ParallaxLayer.SetMirroring()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mirror)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_mirroring", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ParallaxLayer) SetMotionOffset(offset *Vector2) {
	log.Println("Calling ParallaxLayer.SetMotionOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(offset)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_motion_offset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ParallaxLayer) SetMotionScale(scale *Vector2) {
	log.Println("Calling ParallaxLayer.SetMotionScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(scale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_motion_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ParallaxLayerImplementer is an interface for ParallaxLayer objects.
*/
type ParallaxLayerImplementer interface {
	Class
}
