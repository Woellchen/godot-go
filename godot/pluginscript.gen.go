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

//func NewPluginScriptFromPointer(ptr gdnative.Pointer) PluginScript {
func newPluginScriptFromPointer(ptr gdnative.Pointer) PluginScript {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PluginScript{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type PluginScript struct {
	Script
	owner gdnative.Object
}

func (o *PluginScript) BaseClass() string {
	return "PluginScript"
}

/*
        Undocumented
	Args: [], Returns: Variant
*/
func (o *PluginScript) New(args ...gdnative.Variant) gdnative.Variant {
	//log.Println("Calling PluginScript.New()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0+len(args), 0+len(args))

	for i, arg := range args {
		ptrArguments[i+0] = gdnative.NewPointerFromVariant(arg)
	}

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PluginScript", "new")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

// PluginScriptImplementer is an interface that implements the methods
// of the PluginScript class.
type PluginScriptImplementer interface {
	ScriptImplementer
	New(args ...gdnative.Variant) gdnative.Variant
}
