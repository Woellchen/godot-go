package worldenvironment

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
The [code]WorldEnvironment[/code] node is used to configure the default [Environment] for the scene. The parameters defined in the [code]WorldEnvironment[/code] can be overridden by an [Environment] node set on the current [Camera]. Additionally, only one [code]WorldEnvironment[/code] may be instanced in a given scene at a time. The [code]WorldEnvironment[/code] allows the user to specify default lighting parameters (e.g. ambient lighting), various post-processing effects (e.g. SSAO, DOF, Tonemapping), and how to draw the background (e.g. solid color, skybox). Usually, these are added in order to improve the realism/color balance of the scene.
*/
type WorldEnvironment struct {
	Node
}

func (o *WorldEnvironment) BaseClass() string {
	return "WorldEnvironment"
}

/*
   Undocumented
*/
func (o *WorldEnvironment) GetEnvironment() *Environment {
	log.Println("Calling WorldEnvironment.GetEnvironment()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_environment", goArguments, "*Environment")

	returnValue := goRet.Interface().(*Environment)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *WorldEnvironment) SetEnvironment(env *Environment) {
	log.Println("Calling WorldEnvironment.SetEnvironment()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(env)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_environment", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   WorldEnvironmentImplementer is an interface for WorldEnvironment objects.
*/
type WorldEnvironmentImplementer interface {
	Class
}
