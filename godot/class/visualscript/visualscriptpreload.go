package visualscript

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
Undocumented
*/
type VisualScriptPreload struct {
	VisualScriptNode
}

func (o *VisualScriptPreload) BaseClass() string {
	return "VisualScriptPreload"
}

/*
   Undocumented
*/
func (o *VisualScriptPreload) GetPreload() *Resource {
	log.Println("Calling VisualScriptPreload.GetPreload()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_preload", goArguments, "*Resource")

	returnValue := goRet.Interface().(*Resource)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptPreload) SetPreload(resource *Resource) {
	log.Println("Calling VisualScriptPreload.SetPreload()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(resource)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_preload", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptPreloadImplementer is an interface for VisualScriptPreload objects.
*/
type VisualScriptPreloadImplementer interface {
	Class
}
