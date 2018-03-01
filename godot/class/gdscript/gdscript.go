package gdscript

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
type GDScript struct {
	Script
}

func (o *GDScript) BaseClass() string {
	return "GDScript"
}

/*
   Undocumented
*/
func (o *GDScript) GetAsByteCode() *PoolByteArray {
	log.Println("Calling GDScript.GetAsByteCode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_as_byte_code", goArguments, "*PoolByteArray")

	returnValue := goRet.Interface().(*PoolByteArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *GDScript) New() *Object {
	log.Println("Calling GDScript.New()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "new", goArguments, "*Object")

	returnValue := goRet.Interface().(*Object)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   GDScriptImplementer is an interface for GDScript objects.
*/
type GDScriptImplementer interface {
	Class
}
