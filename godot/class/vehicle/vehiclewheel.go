package vehicle

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
This node needs to be used as a child node of [VehicleBody] and simulates the behaviour of one of its wheels. This node also acts as a collider to detect if the wheel is touching a surface.
*/
type VehicleWheel struct {
	Spatial
}

func (o *VehicleWheel) BaseClass() string {
	return "VehicleWheel"
}

/*
   Undocumented
*/
func (o *VehicleWheel) GetDampingCompression() gdnative.Float {
	log.Println("Calling VehicleWheel.GetDampingCompression()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_damping_compression", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetDampingRelaxation() gdnative.Float {
	log.Println("Calling VehicleWheel.GetDampingRelaxation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_damping_relaxation", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetFrictionSlip() gdnative.Float {
	log.Println("Calling VehicleWheel.GetFrictionSlip()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_friction_slip", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetRadius() gdnative.Float {
	log.Println("Calling VehicleWheel.GetRadius()")

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
func (o *VehicleWheel) GetRollInfluence() gdnative.Float {
	log.Println("Calling VehicleWheel.GetRollInfluence()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_roll_influence", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a value between 0.0 and 1.0 that indicates whether this wheel is skidding. 0.0 is not skidding, 1.0 means the wheel has lost grip.
*/
func (o *VehicleWheel) GetSkidinfo() gdnative.Float {
	log.Println("Calling VehicleWheel.GetSkidinfo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_skidinfo", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetSuspensionMaxForce() gdnative.Float {
	log.Println("Calling VehicleWheel.GetSuspensionMaxForce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_suspension_max_force", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetSuspensionRestLength() gdnative.Float {
	log.Println("Calling VehicleWheel.GetSuspensionRestLength()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_suspension_rest_length", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetSuspensionStiffness() gdnative.Float {
	log.Println("Calling VehicleWheel.GetSuspensionStiffness()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_suspension_stiffness", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetSuspensionTravel() gdnative.Float {
	log.Println("Calling VehicleWheel.GetSuspensionTravel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_suspension_travel", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns true if this wheel is in contact with a surface.
*/
func (o *VehicleWheel) IsInContact() gdnative.Bool {
	log.Println("Calling VehicleWheel.IsInContact()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_in_contact", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleWheel) IsUsedAsSteering() gdnative.Bool {
	log.Println("Calling VehicleWheel.IsUsedAsSteering()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_used_as_steering", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleWheel) IsUsedAsTraction() gdnative.Bool {
	log.Println("Calling VehicleWheel.IsUsedAsTraction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_used_as_traction", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetDampingCompression(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetDampingCompression()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(length)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_damping_compression", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetDampingRelaxation(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetDampingRelaxation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(length)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_damping_relaxation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetFrictionSlip(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetFrictionSlip()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(length)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_friction_slip", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetRadius(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(length)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_radius", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetRollInfluence(rollInfluence gdnative.Float) {
	log.Println("Calling VehicleWheel.SetRollInfluence()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rollInfluence)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_roll_influence", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetSuspensionMaxForce(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetSuspensionMaxForce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(length)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_suspension_max_force", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetSuspensionRestLength(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetSuspensionRestLength()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(length)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_suspension_rest_length", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetSuspensionStiffness(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetSuspensionStiffness()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(length)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_suspension_stiffness", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetSuspensionTravel(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetSuspensionTravel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(length)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_suspension_travel", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetUseAsSteering(enable gdnative.Bool) {
	log.Println("Calling VehicleWheel.SetUseAsSteering()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_use_as_steering", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetUseAsTraction(enable gdnative.Bool) {
	log.Println("Calling VehicleWheel.SetUseAsTraction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_use_as_traction", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VehicleWheelImplementer is an interface for VehicleWheel objects.
*/
type VehicleWheelImplementer interface {
	Class
}
