package lightoccluder2d

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
Occludes light cast by a Light2D, casting shadows. The LightOccluder2D must be provided with an [OccluderPolygon2D] in order for the shadow to be computed.
*/
type LightOccluder2D struct {
	Node2D
}

func (o *LightOccluder2D) BaseClass() string {
	return "LightOccluder2D"
}

/*
   Undocumented
*/
func (o *LightOccluder2D) X_PolyChanged() {
	log.Println("Calling LightOccluder2D.X_PolyChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_poly_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *LightOccluder2D) GetOccluderLightMask() gdnative.Int {
	log.Println("Calling LightOccluder2D.GetOccluderLightMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_occluder_light_mask", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *LightOccluder2D) GetOccluderPolygon() *OccluderPolygon2D {
	log.Println("Calling LightOccluder2D.GetOccluderPolygon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_occluder_polygon", goArguments, "*OccluderPolygon2D")

	returnValue := goRet.Interface().(*OccluderPolygon2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *LightOccluder2D) SetOccluderLightMask(mask gdnative.Int) {
	log.Println("Calling LightOccluder2D.SetOccluderLightMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_occluder_light_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *LightOccluder2D) SetOccluderPolygon(polygon *OccluderPolygon2D) {
	log.Println("Calling LightOccluder2D.SetOccluderPolygon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(polygon)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_occluder_polygon", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   LightOccluder2DImplementer is an interface for LightOccluder2D objects.
*/
type LightOccluder2DImplementer interface {
	Class
}
