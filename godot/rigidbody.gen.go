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

// RigidBodyMode is an enum for Mode values.
type RigidBodyMode int

const (
	RigidBodyModeCharacter RigidBodyMode = 2
	RigidBodyModeKinematic RigidBodyMode = 3
	RigidBodyModeRigid     RigidBodyMode = 0
	RigidBodyModeStatic    RigidBodyMode = 1
)

//func NewRigidBodyFromPointer(ptr gdnative.Pointer) RigidBody {
func newRigidBodyFromPointer(ptr gdnative.Pointer) RigidBody {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := RigidBody{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This is the node that implements full 3D physics. This means that you do not control a RigidBody directly. Instead, you can apply forces to it (gravity, impulses, etc.), and the physics simulation will calculate the resulting movement, collision, bouncing, rotating, etc. A RigidBody has 4 behavior [member mode]s: Rigid, Static, Character, and Kinematic. [b]Note:[/b] Don't change a RigidBody's position every frame or very often. Sporadic changes work fine, but physics runs at a different granularity (fixed Hz) than usual rendering (process callback) and maybe even in a separate thread, so changing this from a process loop may result in strange behavior. If you need to directly affect the body's state, use [method _integrate_forces], which allows you to directly access the physics state. If you need to override the default physics behavior, you can write a custom force integration function. See [member custom_integrator].
*/
type RigidBody struct {
	PhysicsBody
	owner gdnative.Object
}

func (o *RigidBody) BaseClass() string {
	return "RigidBody"
}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *RigidBody) X_BodyEnterTree(arg0 gdnative.Int) {
	//log.Println("Calling RigidBody.X_BodyEnterTree()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "_body_enter_tree")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *RigidBody) X_BodyExitTree(arg0 gdnative.Int) {
	//log.Println("Calling RigidBody.X_BodyExitTree()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "_body_exit_tree")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 Object}], Returns: void
*/
func (o *RigidBody) X_DirectStateChanged(arg0 ObjectImplementer) {
	//log.Println("Calling RigidBody.X_DirectStateChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "_direct_state_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it works in addition to the usual physics behavior, but the [member custom_integrator] property allows you to disable the default behavior and do fully custom force integration for a body.
	Args: [{ false state PhysicsDirectBodyState}], Returns: void
*/
func (o *RigidBody) X_IntegrateForces(state PhysicsDirectBodyStateImplementer) {
	//log.Println("Calling RigidBody.X_IntegrateForces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(state.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "_integrate_forces")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *RigidBody) X_ReloadPhysicsCharacteristics() {
	//log.Println("Calling RigidBody.X_ReloadPhysicsCharacteristics()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "_reload_physics_characteristics")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a constant directional force (i.e. acceleration) without affecting rotation. This is equivalent to [code]add_force(force, Vector3(0,0,0))[/code].
	Args: [{ false force Vector3}], Returns: void
*/
func (o *RigidBody) AddCentralForce(force gdnative.Vector3) {
	//log.Println("Calling RigidBody.AddCentralForce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(force)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "add_central_force")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a constant directional force (i.e. acceleration). The position uses the rotation of the global coordinate system, but is centered at the object's origin.
	Args: [{ false force Vector3} { false position Vector3}], Returns: void
*/
func (o *RigidBody) AddForce(force gdnative.Vector3, position gdnative.Vector3) {
	//log.Println("Calling RigidBody.AddForce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector3(force)
	ptrArguments[1] = gdnative.NewPointerFromVector3(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "add_force")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a constant rotational force (i.e. a motor) without affecting position.
	Args: [{ false torque Vector3}], Returns: void
*/
func (o *RigidBody) AddTorque(torque gdnative.Vector3) {
	//log.Println("Calling RigidBody.AddTorque()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(torque)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "add_torque")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Applies a directional impulse without affecting rotation. This is equivalent to [code]apply_impulse(Vector3(0,0,0), impulse)[/code].
	Args: [{ false impulse Vector3}], Returns: void
*/
func (o *RigidBody) ApplyCentralImpulse(impulse gdnative.Vector3) {
	//log.Println("Calling RigidBody.ApplyCentralImpulse()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(impulse)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "apply_central_impulse")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Applies a positioned impulse to the body. An impulse is time independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason it should only be used when simulating one-time impacts. The position uses the rotation of the global coordinate system, but is centered at the object's origin.
	Args: [{ false position Vector3} { false impulse Vector3}], Returns: void
*/
func (o *RigidBody) ApplyImpulse(position gdnative.Vector3, impulse gdnative.Vector3) {
	//log.Println("Calling RigidBody.ApplyImpulse()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector3(position)
	ptrArguments[1] = gdnative.NewPointerFromVector3(impulse)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "apply_impulse")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Applies a torque impulse which will be affected by the body mass and shape. This will rotate the body around the [code]impulse[/code] vector passed.
	Args: [{ false impulse Vector3}], Returns: void
*/
func (o *RigidBody) ApplyTorqueImpulse(impulse gdnative.Vector3) {
	//log.Println("Calling RigidBody.ApplyTorqueImpulse()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(impulse)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "apply_torque_impulse")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *RigidBody) GetAngularDamp() gdnative.Real {
	//log.Println("Calling RigidBody.GetAngularDamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_angular_damp")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *RigidBody) GetAngularVelocity() gdnative.Vector3 {
	//log.Println("Calling RigidBody.GetAngularVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_angular_velocity")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the specified linear or rotational axis is locked.
	Args: [{ false axis int}], Returns: bool
*/
func (o *RigidBody) GetAxisLock(axis gdnative.Int) gdnative.Bool {
	//log.Println("Calling RigidBody.GetAxisLock()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(axis)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_axis_lock")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *RigidBody) GetBounce() gdnative.Real {
	//log.Println("Calling RigidBody.GetBounce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_bounce")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns a list of the bodies colliding with this one. By default, number of max contacts reported is at 0, see the [member contacts_reported] property to increase it. [b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of collisions is updated once per frame and before the physics step. Consider using signals instead.
	Args: [], Returns: Array
*/
func (o *RigidBody) GetCollidingBodies() gdnative.Array {
	//log.Println("Calling RigidBody.GetCollidingBodies()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_colliding_bodies")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *RigidBody) GetFriction() gdnative.Real {
	//log.Println("Calling RigidBody.GetFriction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_friction")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *RigidBody) GetGravityScale() gdnative.Real {
	//log.Println("Calling RigidBody.GetGravityScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_gravity_scale")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *RigidBody) GetLinearDamp() gdnative.Real {
	//log.Println("Calling RigidBody.GetLinearDamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_linear_damp")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *RigidBody) GetLinearVelocity() gdnative.Vector3 {
	//log.Println("Calling RigidBody.GetLinearVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_linear_velocity")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *RigidBody) GetMass() gdnative.Real {
	//log.Println("Calling RigidBody.GetMass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_mass")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *RigidBody) GetMaxContactsReported() gdnative.Int {
	//log.Println("Calling RigidBody.GetMaxContactsReported()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_max_contacts_reported")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.RigidBody::Mode
*/
func (o *RigidBody) GetMode() RigidBodyMode {
	//log.Println("Calling RigidBody.GetMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_mode")

	// Call the parent method.
	// enum.RigidBody::Mode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return RigidBodyMode(ret)
}

/*
        Undocumented
	Args: [], Returns: PhysicsMaterial
*/
func (o *RigidBody) GetPhysicsMaterialOverride() PhysicsMaterialImplementer {
	//log.Println("Calling RigidBody.GetPhysicsMaterialOverride()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_physics_material_override")

	// Call the parent method.
	// PhysicsMaterial
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newPhysicsMaterialFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(PhysicsMaterialImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "PhysicsMaterial" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(PhysicsMaterialImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *RigidBody) GetWeight() gdnative.Real {
	//log.Println("Calling RigidBody.GetWeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "get_weight")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RigidBody) IsAbleToSleep() gdnative.Bool {
	//log.Println("Calling RigidBody.IsAbleToSleep()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "is_able_to_sleep")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RigidBody) IsContactMonitorEnabled() gdnative.Bool {
	//log.Println("Calling RigidBody.IsContactMonitorEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "is_contact_monitor_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RigidBody) IsSleeping() gdnative.Bool {
	//log.Println("Calling RigidBody.IsSleeping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "is_sleeping")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RigidBody) IsUsingContinuousCollisionDetection() gdnative.Bool {
	//log.Println("Calling RigidBody.IsUsingContinuousCollisionDetection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "is_using_continuous_collision_detection")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RigidBody) IsUsingCustomIntegrator() gdnative.Bool {
	//log.Println("Calling RigidBody.IsUsingCustomIntegrator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "is_using_custom_integrator")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false angular_damp float}], Returns: void
*/
func (o *RigidBody) SetAngularDamp(angularDamp gdnative.Real) {
	//log.Println("Calling RigidBody.SetAngularDamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(angularDamp)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_angular_damp")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false angular_velocity Vector3}], Returns: void
*/
func (o *RigidBody) SetAngularVelocity(angularVelocity gdnative.Vector3) {
	//log.Println("Calling RigidBody.SetAngularVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(angularVelocity)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_angular_velocity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Locks the specified linear or rotational axis.
	Args: [{ false axis int} { false lock bool}], Returns: void
*/
func (o *RigidBody) SetAxisLock(axis gdnative.Int, lock gdnative.Bool) {
	//log.Println("Calling RigidBody.SetAxisLock()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(axis)
	ptrArguments[1] = gdnative.NewPointerFromBool(lock)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_axis_lock")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets an axis velocity. The velocity in the given vector axis will be set as the given vector length. This is useful for jumping behavior.
	Args: [{ false axis_velocity Vector3}], Returns: void
*/
func (o *RigidBody) SetAxisVelocity(axisVelocity gdnative.Vector3) {
	//log.Println("Calling RigidBody.SetAxisVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(axisVelocity)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_axis_velocity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bounce float}], Returns: void
*/
func (o *RigidBody) SetBounce(bounce gdnative.Real) {
	//log.Println("Calling RigidBody.SetBounce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(bounce)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_bounce")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false able_to_sleep bool}], Returns: void
*/
func (o *RigidBody) SetCanSleep(ableToSleep gdnative.Bool) {
	//log.Println("Calling RigidBody.SetCanSleep()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(ableToSleep)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_can_sleep")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *RigidBody) SetContactMonitor(enabled gdnative.Bool) {
	//log.Println("Calling RigidBody.SetContactMonitor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_contact_monitor")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false friction float}], Returns: void
*/
func (o *RigidBody) SetFriction(friction gdnative.Real) {
	//log.Println("Calling RigidBody.SetFriction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(friction)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_friction")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false gravity_scale float}], Returns: void
*/
func (o *RigidBody) SetGravityScale(gravityScale gdnative.Real) {
	//log.Println("Calling RigidBody.SetGravityScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(gravityScale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_gravity_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false linear_damp float}], Returns: void
*/
func (o *RigidBody) SetLinearDamp(linearDamp gdnative.Real) {
	//log.Println("Calling RigidBody.SetLinearDamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(linearDamp)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_linear_damp")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false linear_velocity Vector3}], Returns: void
*/
func (o *RigidBody) SetLinearVelocity(linearVelocity gdnative.Vector3) {
	//log.Println("Calling RigidBody.SetLinearVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(linearVelocity)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_linear_velocity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mass float}], Returns: void
*/
func (o *RigidBody) SetMass(mass gdnative.Real) {
	//log.Println("Calling RigidBody.SetMass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(mass)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_mass")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount int}], Returns: void
*/
func (o *RigidBody) SetMaxContactsReported(amount gdnative.Int) {
	//log.Println("Calling RigidBody.SetMaxContactsReported()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_max_contacts_reported")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *RigidBody) SetMode(mode gdnative.Int) {
	//log.Println("Calling RigidBody.SetMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false physics_material_override PhysicsMaterial}], Returns: void
*/
func (o *RigidBody) SetPhysicsMaterialOverride(physicsMaterialOverride PhysicsMaterialImplementer) {
	//log.Println("Calling RigidBody.SetPhysicsMaterialOverride()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(physicsMaterialOverride.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_physics_material_override")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false sleeping bool}], Returns: void
*/
func (o *RigidBody) SetSleeping(sleeping gdnative.Bool) {
	//log.Println("Calling RigidBody.SetSleeping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(sleeping)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_sleeping")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *RigidBody) SetUseContinuousCollisionDetection(enable gdnative.Bool) {
	//log.Println("Calling RigidBody.SetUseContinuousCollisionDetection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_use_continuous_collision_detection")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *RigidBody) SetUseCustomIntegrator(enable gdnative.Bool) {
	//log.Println("Calling RigidBody.SetUseCustomIntegrator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_use_custom_integrator")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false weight float}], Returns: void
*/
func (o *RigidBody) SetWeight(weight gdnative.Real) {
	//log.Println("Calling RigidBody.SetWeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(weight)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RigidBody", "set_weight")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// RigidBodyImplementer is an interface that implements the methods
// of the RigidBody class.
type RigidBodyImplementer interface {
	PhysicsBodyImplementer
	X_BodyEnterTree(arg0 gdnative.Int)
	X_BodyExitTree(arg0 gdnative.Int)
	X_DirectStateChanged(arg0 ObjectImplementer)
	X_IntegrateForces(state PhysicsDirectBodyStateImplementer)
	X_ReloadPhysicsCharacteristics()
	AddCentralForce(force gdnative.Vector3)
	AddForce(force gdnative.Vector3, position gdnative.Vector3)
	AddTorque(torque gdnative.Vector3)
	ApplyCentralImpulse(impulse gdnative.Vector3)
	ApplyImpulse(position gdnative.Vector3, impulse gdnative.Vector3)
	ApplyTorqueImpulse(impulse gdnative.Vector3)
	GetAngularDamp() gdnative.Real
	GetAngularVelocity() gdnative.Vector3
	GetAxisLock(axis gdnative.Int) gdnative.Bool
	GetBounce() gdnative.Real
	GetCollidingBodies() gdnative.Array
	GetFriction() gdnative.Real
	GetGravityScale() gdnative.Real
	GetLinearDamp() gdnative.Real
	GetLinearVelocity() gdnative.Vector3
	GetMass() gdnative.Real
	GetMaxContactsReported() gdnative.Int
	GetPhysicsMaterialOverride() PhysicsMaterialImplementer
	GetWeight() gdnative.Real
	IsAbleToSleep() gdnative.Bool
	IsContactMonitorEnabled() gdnative.Bool
	IsSleeping() gdnative.Bool
	IsUsingContinuousCollisionDetection() gdnative.Bool
	IsUsingCustomIntegrator() gdnative.Bool
	SetAngularDamp(angularDamp gdnative.Real)
	SetAngularVelocity(angularVelocity gdnative.Vector3)
	SetAxisLock(axis gdnative.Int, lock gdnative.Bool)
	SetAxisVelocity(axisVelocity gdnative.Vector3)
	SetBounce(bounce gdnative.Real)
	SetCanSleep(ableToSleep gdnative.Bool)
	SetContactMonitor(enabled gdnative.Bool)
	SetFriction(friction gdnative.Real)
	SetGravityScale(gravityScale gdnative.Real)
	SetLinearDamp(linearDamp gdnative.Real)
	SetLinearVelocity(linearVelocity gdnative.Vector3)
	SetMass(mass gdnative.Real)
	SetMaxContactsReported(amount gdnative.Int)
	SetMode(mode gdnative.Int)
	SetPhysicsMaterialOverride(physicsMaterialOverride PhysicsMaterialImplementer)
	SetSleeping(sleeping gdnative.Bool)
	SetUseContinuousCollisionDetection(enable gdnative.Bool)
	SetUseCustomIntegrator(enable gdnative.Bool)
	SetWeight(weight gdnative.Real)
}
