package remotetransform

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
RemoteTransform leads the [Transform] of another [Spatial] derived Node (called the remote node) in the scene. It can be set to track another Node's position, rotation and/or scale. It can update using either global or local coordinates.
*/
type RemoteTransform struct {
	Spatial
}

func (o *RemoteTransform) BaseClass() string {
	return "RemoteTransform"
}

/*
   Undocumented
*/
func (o *RemoteTransform) GetRemoteNode() *NodePath {
	log.Println("Calling RemoteTransform.GetRemoteNode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_remote_node", goArguments, "*NodePath")

	returnValue := goRet.Interface().(*NodePath)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RemoteTransform) GetUpdatePosition() gdnative.Bool {
	log.Println("Calling RemoteTransform.GetUpdatePosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_update_position", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RemoteTransform) GetUpdateRotation() gdnative.Bool {
	log.Println("Calling RemoteTransform.GetUpdateRotation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_update_rotation", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RemoteTransform) GetUpdateScale() gdnative.Bool {
	log.Println("Calling RemoteTransform.GetUpdateScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_update_scale", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RemoteTransform) GetUseGlobalCoordinates() gdnative.Bool {
	log.Println("Calling RemoteTransform.GetUseGlobalCoordinates()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_use_global_coordinates", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RemoteTransform) SetRemoteNode(path *NodePath) {
	log.Println("Calling RemoteTransform.SetRemoteNode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(path)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_remote_node", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RemoteTransform) SetUpdatePosition(updateRemotePosition gdnative.Bool) {
	log.Println("Calling RemoteTransform.SetUpdatePosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(updateRemotePosition)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_update_position", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RemoteTransform) SetUpdateRotation(updateRemoteRotation gdnative.Bool) {
	log.Println("Calling RemoteTransform.SetUpdateRotation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(updateRemoteRotation)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_update_rotation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RemoteTransform) SetUpdateScale(updateRemoteScale gdnative.Bool) {
	log.Println("Calling RemoteTransform.SetUpdateScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(updateRemoteScale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_update_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RemoteTransform) SetUseGlobalCoordinates(useGlobalCoordinates gdnative.Bool) {
	log.Println("Calling RemoteTransform.SetUseGlobalCoordinates()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(useGlobalCoordinates)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_use_global_coordinates", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   RemoteTransformImplementer is an interface for RemoteTransform objects.
*/
type RemoteTransformImplementer interface {
	Class
}
