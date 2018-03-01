package resource

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
Resource Preloader Node. This node is used to preload sub-resources inside a scene, so when the scene is loaded all the resources are ready to use and be retrieved from here.
*/
type ResourcePreloader struct {
	Node
}

func (o *ResourcePreloader) BaseClass() string {
	return "ResourcePreloader"
}

/*
   Undocumented
*/
func (o *ResourcePreloader) X_GetResources() *Array {
	log.Println("Calling ResourcePreloader.X_GetResources()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "_get_resources", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ResourcePreloader) X_SetResources(arg0 *Array) {
	log.Println("Calling ResourcePreloader.X_SetResources()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_set_resources", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ResourcePreloader) AddResource(name gdnative.String, resource *Resource) {
	log.Println("Calling ResourcePreloader.AddResource()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(name)
	goArguments[1] = reflect.ValueOf(resource)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_resource", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the resource given a text-id.
*/
func (o *ResourcePreloader) GetResource(name gdnative.String) *Resource {
	log.Println("Calling ResourcePreloader.GetResource()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_resource", goArguments, "*Resource")

	returnValue := goRet.Interface().(*Resource)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the list of resources inside the preloader.
*/
func (o *ResourcePreloader) GetResourceList() *PoolStringArray {
	log.Println("Calling ResourcePreloader.GetResourceList()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_resource_list", goArguments, "*PoolStringArray")

	returnValue := goRet.Interface().(*PoolStringArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return true if the preloader has a given resource.
*/
func (o *ResourcePreloader) HasResource(name gdnative.String) gdnative.Bool {
	log.Println("Calling ResourcePreloader.HasResource()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "has_resource", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Remove a resource from the preloader by text id.
*/
func (o *ResourcePreloader) RemoveResource(name gdnative.String) {
	log.Println("Calling ResourcePreloader.RemoveResource()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "remove_resource", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Rename a resource inside the preloader, from a text-id to a new text-id.
*/
func (o *ResourcePreloader) RenameResource(name gdnative.String, newname gdnative.String) {
	log.Println("Calling ResourcePreloader.RenameResource()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(name)
	goArguments[1] = reflect.ValueOf(newname)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "rename_resource", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ResourcePreloaderImplementer is an interface for ResourcePreloader objects.
*/
type ResourcePreloaderImplementer interface {
	Class
}
