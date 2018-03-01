package pathfollow

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
This node takes its parent [Path], and returns the coordinates of a point within it, given a distance from the first vertex. It is useful for making other nodes follow a path, without coding the movement pattern. For that, the nodes must be descendants of this node. Then, when setting an offset in this node, the descendant nodes will move accordingly.
*/
type PathFollow struct {
	Spatial
}

func (o *PathFollow) BaseClass() string {
	return "PathFollow"
}

/*
   Undocumented
*/
func (o *PathFollow) GetCubicInterpolation() gdnative.Bool {
	log.Println("Calling PathFollow.GetCubicInterpolation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_cubic_interpolation", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PathFollow) GetHOffset() gdnative.Float {
	log.Println("Calling PathFollow.GetHOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_h_offset", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PathFollow) GetOffset() gdnative.Float {
	log.Println("Calling PathFollow.GetOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_offset", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PathFollow) GetRotationMode() gdnative.Int {
	log.Println("Calling PathFollow.GetRotationMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_rotation_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PathFollow) GetUnitOffset() gdnative.Float {
	log.Println("Calling PathFollow.GetUnitOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_unit_offset", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PathFollow) GetVOffset() gdnative.Float {
	log.Println("Calling PathFollow.GetVOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_v_offset", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PathFollow) HasLoop() gdnative.Bool {
	log.Println("Calling PathFollow.HasLoop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "has_loop", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PathFollow) SetCubicInterpolation(enable gdnative.Bool) {
	log.Println("Calling PathFollow.SetCubicInterpolation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_cubic_interpolation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PathFollow) SetHOffset(hOffset gdnative.Float) {
	log.Println("Calling PathFollow.SetHOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(hOffset)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_h_offset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PathFollow) SetLoop(loop gdnative.Bool) {
	log.Println("Calling PathFollow.SetLoop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(loop)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_loop", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PathFollow) SetOffset(offset gdnative.Float) {
	log.Println("Calling PathFollow.SetOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(offset)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_offset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PathFollow) SetRotationMode(rotationMode gdnative.Int) {
	log.Println("Calling PathFollow.SetRotationMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rotationMode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_rotation_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PathFollow) SetUnitOffset(unitOffset gdnative.Float) {
	log.Println("Calling PathFollow.SetUnitOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(unitOffset)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_unit_offset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PathFollow) SetVOffset(vOffset gdnative.Float) {
	log.Println("Calling PathFollow.SetVOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(vOffset)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_v_offset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   PathFollowImplementer is an interface for PathFollow objects.
*/
type PathFollowImplementer interface {
	Class
}
