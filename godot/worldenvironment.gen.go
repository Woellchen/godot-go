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

//func NewWorldEnvironmentFromPointer(ptr gdnative.Pointer) WorldEnvironment {
func newWorldEnvironmentFromPointer(ptr gdnative.Pointer) WorldEnvironment {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := WorldEnvironment{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The [WorldEnvironment] node is used to configure the default [Environment] for the scene. The parameters defined in the [WorldEnvironment] can be overridden by an [Environment] node set on the current [Camera]. Additionally, only one [WorldEnvironment] may be instanced in a given scene at a time. The [WorldEnvironment] allows the user to specify default lighting parameters (e.g. ambient lighting), various post-processing effects (e.g. SSAO, DOF, Tonemapping), and how to draw the background (e.g. solid color, skybox). Usually, these are added in order to improve the realism/color balance of the scene.
*/
type WorldEnvironment struct {
	Node
	owner gdnative.Object
}

func (o *WorldEnvironment) BaseClass() string {
	return "WorldEnvironment"
}

/*
        Undocumented
	Args: [], Returns: Environment
*/
func (o *WorldEnvironment) GetEnvironment() EnvironmentImplementer {
	//log.Println("Calling WorldEnvironment.GetEnvironment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WorldEnvironment", "get_environment")

	// Call the parent method.
	// Environment
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newEnvironmentFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(EnvironmentImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Environment" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(EnvironmentImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false env Environment}], Returns: void
*/
func (o *WorldEnvironment) SetEnvironment(env EnvironmentImplementer) {
	//log.Println("Calling WorldEnvironment.SetEnvironment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(env.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WorldEnvironment", "set_environment")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// WorldEnvironmentImplementer is an interface that implements the methods
// of the WorldEnvironment class.
type WorldEnvironmentImplementer interface {
	NodeImplementer
	GetEnvironment() EnvironmentImplementer
	SetEnvironment(env EnvironmentImplementer)
}
