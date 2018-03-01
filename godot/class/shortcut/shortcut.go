package shortcut

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
A shortcut for binding input. Shortcuts are commonly used for interacting with a [Control] element from a [InputEvent].
*/
type ShortCut struct {
	Resource
}

func (o *ShortCut) BaseClass() string {
	return "ShortCut"
}

/*
   Returns the Shortcut's [InputEvent] as a [String].
*/
func (o *ShortCut) GetAsText() gdnative.String {
	log.Println("Calling ShortCut.GetAsText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_as_text", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ShortCut) GetShortcut() *InputEvent {
	log.Println("Calling ShortCut.GetShortcut()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_shortcut", goArguments, "*InputEvent")

	returnValue := goRet.Interface().(*InputEvent)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the Shortcut's [InputEvent] equals [code]event[/code].
*/
func (o *ShortCut) IsShortcut(event *InputEvent) gdnative.Bool {
	log.Println("Calling ShortCut.IsShortcut()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(event)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_shortcut", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   If [code]true[/code] this Shortcut is valid.
*/
func (o *ShortCut) IsValid() gdnative.Bool {
	log.Println("Calling ShortCut.IsValid()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_valid", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ShortCut) SetShortcut(event *InputEvent) {
	log.Println("Calling ShortCut.SetShortcut()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(event)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_shortcut", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ShortCutImplementer is an interface for ShortCut objects.
*/
type ShortCutImplementer interface {
	Class
}
