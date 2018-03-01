package rigidbody

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
This is the node that implements full 3D physics. This means that you do not control a RigidBody directly. Instead you can apply forces to it (gravity, impulses, etc.), and the physics simulation will calculate the resulting movement, collision, bouncing, rotating, etc. A RigidBody has 4 behavior [member mode]s: Rigid, Static, Character, and Kinematic. [b]Note:[/b] Don't change a RigidBody's position every frame or very often. Sporadic changes work fine, but physics runs at a different granularity (fixed hz) than usual rendering (process callback) and maybe even in a separate thread, so changing this from a process loop will yield strange behavior. If you need to directly affect the body's state, use [method _integrate_forces], which allows you to directly access the physics state. If you need to override the default physics behavior, you can write a custom force integration. See [member custom_integrator].
*/
type RigidBody struct {
	PhysicsBody
}

func (o *RigidBody) BaseClass() string {
	return "RigidBody"
}

/*
   Undocumented
*/
func (o *RigidBody) X_BodyEnterTree(arg0 gdnative.Int) {
	log.Println("Calling RigidBody.X_BodyEnterTree()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_body_enter_tree", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) X_BodyExitTree(arg0 gdnative.Int) {
	log.Println("Calling RigidBody.X_BodyExitTree()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_body_exit_tree", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) X_DirectStateChanged(arg0 *Object) {
	log.Println("Calling RigidBody.X_DirectStateChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_direct_state_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default it works in addition to the usual physics behavior, but [method set_use_custom_integrator] allows you to disable the default behavior and do fully custom force integration for a body.
*/
func (o *RigidBody) X_IntegrateForces(state *PhysicsDirectBodyState) {
	log.Println("Calling RigidBody.X_IntegrateForces()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(state)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_integrate_forces", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Apply a positioned impulse (which will be affected by the body mass and shape). This is the equivalent of hitting a billiard ball with a cue: a force that is applied once, and only once. Both the impulse and the position are in global coordinates, and the position is relative to the object's origin.
*/
func (o *RigidBody) ApplyImpulse(position *Vector3, impulse *Vector3) {
	log.Println("Calling RigidBody.ApplyImpulse()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(position)
	goArguments[1] = reflect.ValueOf(impulse)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "apply_impulse", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) GetAngularDamp() gdnative.Float {
	log.Println("Calling RigidBody.GetAngularDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_angular_damp", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) GetAngularVelocity() *Vector3 {
	log.Println("Calling RigidBody.GetAngularVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_angular_velocity", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) GetAxisLock(axis gdnative.Int) gdnative.Bool {
	log.Println("Calling RigidBody.GetAxisLock()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(axis)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_axis_lock", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) GetBounce() gdnative.Float {
	log.Println("Calling RigidBody.GetBounce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_bounce", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return a list of the bodies colliding with this one. By default, number of max contacts reported is at 0 , see [method set_max_contacts_reported] to increase it. Note that the result of this test is not immediate after moving objects. For performance, list of collisions is updated once per frame and before the physics step. Consider using signals instead.
*/
func (o *RigidBody) GetCollidingBodies() *Array {
	log.Println("Calling RigidBody.GetCollidingBodies()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_colliding_bodies", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) GetFriction() gdnative.Float {
	log.Println("Calling RigidBody.GetFriction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_friction", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) GetGravityScale() gdnative.Float {
	log.Println("Calling RigidBody.GetGravityScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_gravity_scale", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) GetLinearDamp() gdnative.Float {
	log.Println("Calling RigidBody.GetLinearDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_linear_damp", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) GetLinearVelocity() *Vector3 {
	log.Println("Calling RigidBody.GetLinearVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_linear_velocity", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) GetMass() gdnative.Float {
	log.Println("Calling RigidBody.GetMass()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mass", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) GetMaxContactsReported() gdnative.Int {
	log.Println("Calling RigidBody.GetMaxContactsReported()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_max_contacts_reported", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) GetMode() gdnative.Int {
	log.Println("Calling RigidBody.GetMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) GetWeight() gdnative.Float {
	log.Println("Calling RigidBody.GetWeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_weight", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) IsAbleToSleep() gdnative.Bool {
	log.Println("Calling RigidBody.IsAbleToSleep()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_able_to_sleep", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) IsContactMonitorEnabled() gdnative.Bool {
	log.Println("Calling RigidBody.IsContactMonitorEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_contact_monitor_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) IsSleeping() gdnative.Bool {
	log.Println("Calling RigidBody.IsSleeping()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_sleeping", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) IsUsingContinuousCollisionDetection() gdnative.Bool {
	log.Println("Calling RigidBody.IsUsingContinuousCollisionDetection()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_using_continuous_collision_detection", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) IsUsingCustomIntegrator() gdnative.Bool {
	log.Println("Calling RigidBody.IsUsingCustomIntegrator()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_using_custom_integrator", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody) SetAngularDamp(angularDamp gdnative.Float) {
	log.Println("Calling RigidBody.SetAngularDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(angularDamp)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_angular_damp", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetAngularVelocity(angularVelocity *Vector3) {
	log.Println("Calling RigidBody.SetAngularVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(angularVelocity)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_angular_velocity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetAxisLock(axis gdnative.Int, lock gdnative.Bool) {
	log.Println("Calling RigidBody.SetAxisLock()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(axis)
	goArguments[1] = reflect.ValueOf(lock)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_axis_lock", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set an axis velocity. The velocity in the given vector axis will be set as the given vector length. This is useful for jumping behavior.
*/
func (o *RigidBody) SetAxisVelocity(axisVelocity *Vector3) {
	log.Println("Calling RigidBody.SetAxisVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(axisVelocity)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_axis_velocity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetBounce(bounce gdnative.Float) {
	log.Println("Calling RigidBody.SetBounce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bounce)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_bounce", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetCanSleep(ableToSleep gdnative.Bool) {
	log.Println("Calling RigidBody.SetCanSleep()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(ableToSleep)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_can_sleep", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetContactMonitor(enabled gdnative.Bool) {
	log.Println("Calling RigidBody.SetContactMonitor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_contact_monitor", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetFriction(friction gdnative.Float) {
	log.Println("Calling RigidBody.SetFriction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(friction)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_friction", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetGravityScale(gravityScale gdnative.Float) {
	log.Println("Calling RigidBody.SetGravityScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(gravityScale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_gravity_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetLinearDamp(linearDamp gdnative.Float) {
	log.Println("Calling RigidBody.SetLinearDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(linearDamp)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_linear_damp", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetLinearVelocity(linearVelocity *Vector3) {
	log.Println("Calling RigidBody.SetLinearVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(linearVelocity)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_linear_velocity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetMass(mass gdnative.Float) {
	log.Println("Calling RigidBody.SetMass()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mass)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_mass", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetMaxContactsReported(amount gdnative.Int) {
	log.Println("Calling RigidBody.SetMaxContactsReported()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_max_contacts_reported", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetMode(mode gdnative.Int) {
	log.Println("Calling RigidBody.SetMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetSleeping(sleeping gdnative.Bool) {
	log.Println("Calling RigidBody.SetSleeping()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(sleeping)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_sleeping", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetUseContinuousCollisionDetection(enable gdnative.Bool) {
	log.Println("Calling RigidBody.SetUseContinuousCollisionDetection()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_use_continuous_collision_detection", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetUseCustomIntegrator(enable gdnative.Bool) {
	log.Println("Calling RigidBody.SetUseCustomIntegrator()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_use_custom_integrator", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody) SetWeight(weight gdnative.Float) {
	log.Println("Calling RigidBody.SetWeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(weight)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_weight", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   RigidBodyImplementer is an interface for RigidBody objects.
*/
type RigidBodyImplementer interface {
	Class
}
