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

//func NewRemoteTransformFromPointer(ptr gdnative.Pointer) RemoteTransform {
func newRemoteTransformFromPointer(ptr gdnative.Pointer) RemoteTransform {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := RemoteTransform{}
	obj.SetBaseObject(owner)

	return obj
}

/*
RemoteTransform pushes its own [Transform] to another [Spatial] derived Node (called the remote node) in the scene. It can be set to update another Node's position, rotation and/or scale. It can use either global or local coordinates.
*/
type RemoteTransform struct {
	Spatial
	owner gdnative.Object
}

func (o *RemoteTransform) BaseClass() string {
	return "RemoteTransform"
}

/*
        [RemoteTransform] caches the remote node. It may not notice if the remote node disappears; [method force_update_cache] forces it to update the cache again.
	Args: [], Returns: void
*/
func (o *RemoteTransform) ForceUpdateCache() {
	// log.Println("Calling RemoteTransform.ForceUpdateCache()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform", "force_update_cache")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *RemoteTransform) GetRemoteNode() gdnative.NodePath {
	// log.Println("Calling RemoteTransform.GetRemoteNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform", "get_remote_node")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RemoteTransform) GetUpdatePosition() gdnative.Bool {
	// log.Println("Calling RemoteTransform.GetUpdatePosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform", "get_update_position")

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
func (o *RemoteTransform) GetUpdateRotation() gdnative.Bool {
	// log.Println("Calling RemoteTransform.GetUpdateRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform", "get_update_rotation")

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
func (o *RemoteTransform) GetUpdateScale() gdnative.Bool {
	// log.Println("Calling RemoteTransform.GetUpdateScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform", "get_update_scale")

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
func (o *RemoteTransform) GetUseGlobalCoordinates() gdnative.Bool {
	// log.Println("Calling RemoteTransform.GetUseGlobalCoordinates()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform", "get_use_global_coordinates")

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
	Args: [{ false path NodePath}], Returns: void
*/
func (o *RemoteTransform) SetRemoteNode(path gdnative.NodePath) {
	// log.Println("Calling RemoteTransform.SetRemoteNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform", "set_remote_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false update_remote_position bool}], Returns: void
*/
func (o *RemoteTransform) SetUpdatePosition(updateRemotePosition gdnative.Bool) {
	// log.Println("Calling RemoteTransform.SetUpdatePosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(updateRemotePosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform", "set_update_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false update_remote_rotation bool}], Returns: void
*/
func (o *RemoteTransform) SetUpdateRotation(updateRemoteRotation gdnative.Bool) {
	// log.Println("Calling RemoteTransform.SetUpdateRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(updateRemoteRotation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform", "set_update_rotation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false update_remote_scale bool}], Returns: void
*/
func (o *RemoteTransform) SetUpdateScale(updateRemoteScale gdnative.Bool) {
	// log.Println("Calling RemoteTransform.SetUpdateScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(updateRemoteScale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform", "set_update_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false use_global_coordinates bool}], Returns: void
*/
func (o *RemoteTransform) SetUseGlobalCoordinates(useGlobalCoordinates gdnative.Bool) {
	// log.Println("Calling RemoteTransform.SetUseGlobalCoordinates()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(useGlobalCoordinates)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform", "set_use_global_coordinates")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// RemoteTransformImplementer is an interface that implements the methods
// of the RemoteTransform class.
type RemoteTransformImplementer interface {
	SpatialImplementer
	ForceUpdateCache()
	GetRemoteNode() gdnative.NodePath
	GetUpdatePosition() gdnative.Bool
	GetUpdateRotation() gdnative.Bool
	GetUpdateScale() gdnative.Bool
	GetUseGlobalCoordinates() gdnative.Bool
	SetRemoteNode(path gdnative.NodePath)
	SetUpdatePosition(updateRemotePosition gdnative.Bool)
	SetUpdateRotation(updateRemoteRotation gdnative.Bool)
	SetUpdateScale(updateRemoteScale gdnative.Bool)
	SetUseGlobalCoordinates(useGlobalCoordinates gdnative.Bool)
}
