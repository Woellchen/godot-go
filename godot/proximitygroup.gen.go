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

// ProximityGroupDispatchMode is an enum for DispatchMode values.
type ProximityGroupDispatchMode int

const (
	ProximityGroupModeProxy  ProximityGroupDispatchMode = 0
	ProximityGroupModeSignal ProximityGroupDispatchMode = 1
)

//func NewProximityGroupFromPointer(ptr gdnative.Pointer) ProximityGroup {
func newProximityGroupFromPointer(ptr gdnative.Pointer) ProximityGroup {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ProximityGroup{}
	obj.SetBaseObject(owner)

	return obj
}

/*
General-purpose proximity detection node.
*/
type ProximityGroup struct {
	Spatial
	owner gdnative.Object
}

func (o *ProximityGroup) BaseClass() string {
	return "ProximityGroup"
}

/*
        Undocumented
	Args: [{ false name String} { false params Variant}], Returns: void
*/
func (o *ProximityGroup) X_ProximityGroupBroadcast(name gdnative.String, params gdnative.Variant) {
	// log.Println("Calling ProximityGroup.X_ProximityGroupBroadcast()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromVariant(params)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProximityGroup", "_proximity_group_broadcast")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false name String} { false parameters Variant}], Returns: void
*/
func (o *ProximityGroup) Broadcast(name gdnative.String, parameters gdnative.Variant) {
	// log.Println("Calling ProximityGroup.Broadcast()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromVariant(parameters)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProximityGroup", "broadcast")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: enum.ProximityGroup::DispatchMode
*/
func (o *ProximityGroup) GetDispatchMode() ProximityGroupDispatchMode {
	// log.Println("Calling ProximityGroup.GetDispatchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProximityGroup", "get_dispatch_mode")

	// Call the parent method.
	// enum.ProximityGroup::DispatchMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ProximityGroupDispatchMode(ret)
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *ProximityGroup) GetGridRadius() gdnative.Vector3 {
	// log.Println("Calling ProximityGroup.GetGridRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProximityGroup", "get_grid_radius")

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
	Args: [], Returns: String
*/
func (o *ProximityGroup) GetGroupName() gdnative.String {
	// log.Println("Calling ProximityGroup.GetGroupName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProximityGroup", "get_group_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *ProximityGroup) SetDispatchMode(mode gdnative.Int) {
	// log.Println("Calling ProximityGroup.SetDispatchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProximityGroup", "set_dispatch_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false radius Vector3}], Returns: void
*/
func (o *ProximityGroup) SetGridRadius(radius gdnative.Vector3) {
	// log.Println("Calling ProximityGroup.SetGridRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(radius)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProximityGroup", "set_grid_radius")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false name String}], Returns: void
*/
func (o *ProximityGroup) SetGroupName(name gdnative.String) {
	// log.Println("Calling ProximityGroup.SetGroupName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProximityGroup", "set_group_name")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ProximityGroupImplementer is an interface that implements the methods
// of the ProximityGroup class.
type ProximityGroupImplementer interface {
	SpatialImplementer
	X_ProximityGroupBroadcast(name gdnative.String, params gdnative.Variant)
	Broadcast(name gdnative.String, parameters gdnative.Variant)
	GetGridRadius() gdnative.Vector3
	GetGroupName() gdnative.String
	SetDispatchMode(mode gdnative.Int)
	SetGridRadius(radius gdnative.Vector3)
	SetGroupName(name gdnative.String)
}
