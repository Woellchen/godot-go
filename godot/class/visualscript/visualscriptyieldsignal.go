package visualscript

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
Undocumented
*/
type VisualScriptYieldSignal struct {
	VisualScriptNode
}

func (o *VisualScriptYieldSignal) BaseClass() string {
	return "VisualScriptYieldSignal"
}

/*
   Undocumented
*/
func (o *VisualScriptYieldSignal) GetBasePath() *NodePath {
	log.Println("Calling VisualScriptYieldSignal.GetBasePath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_base_path", goArguments, "*NodePath")

	returnValue := goRet.Interface().(*NodePath)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptYieldSignal) GetBaseType() gdnative.String {
	log.Println("Calling VisualScriptYieldSignal.GetBaseType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_base_type", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptYieldSignal) GetCallMode() gdnative.Int {
	log.Println("Calling VisualScriptYieldSignal.GetCallMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_call_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptYieldSignal) GetSignal() gdnative.String {
	log.Println("Calling VisualScriptYieldSignal.GetSignal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_signal", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptYieldSignal) SetBasePath(basePath *NodePath) {
	log.Println("Calling VisualScriptYieldSignal.SetBasePath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(basePath)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_base_path", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VisualScriptYieldSignal) SetBaseType(baseType gdnative.String) {
	log.Println("Calling VisualScriptYieldSignal.SetBaseType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(baseType)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_base_type", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VisualScriptYieldSignal) SetCallMode(mode gdnative.Int) {
	log.Println("Calling VisualScriptYieldSignal.SetCallMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_call_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VisualScriptYieldSignal) SetSignal(signal gdnative.String) {
	log.Println("Calling VisualScriptYieldSignal.SetSignal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(signal)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_signal", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptYieldSignalImplementer is an interface for VisualScriptYieldSignal objects.
*/
type VisualScriptYieldSignalImplementer interface {
	Class
}
