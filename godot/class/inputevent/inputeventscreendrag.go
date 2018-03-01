package inputevent

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
Contains screen drag information. See [method Node._input].
*/
type InputEventScreenDrag struct {
	InputEvent
}

func (o *InputEventScreenDrag) BaseClass() string {
	return "InputEventScreenDrag"
}

/*
   Undocumented
*/
func (o *InputEventScreenDrag) GetIndex() gdnative.Int {
	log.Println("Calling InputEventScreenDrag.GetIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_index", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *InputEventScreenDrag) GetPosition() *Vector2 {
	log.Println("Calling InputEventScreenDrag.GetPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_position", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *InputEventScreenDrag) GetRelative() *Vector2 {
	log.Println("Calling InputEventScreenDrag.GetRelative()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_relative", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *InputEventScreenDrag) GetSpeed() *Vector2 {
	log.Println("Calling InputEventScreenDrag.GetSpeed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_speed", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *InputEventScreenDrag) SetIndex(index gdnative.Int) {
	log.Println("Calling InputEventScreenDrag.SetIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(index)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_index", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *InputEventScreenDrag) SetPosition(position *Vector2) {
	log.Println("Calling InputEventScreenDrag.SetPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(position)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_position", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *InputEventScreenDrag) SetRelative(relative *Vector2) {
	log.Println("Calling InputEventScreenDrag.SetRelative()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(relative)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_relative", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *InputEventScreenDrag) SetSpeed(speed *Vector2) {
	log.Println("Calling InputEventScreenDrag.SetSpeed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(speed)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_speed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   InputEventScreenDragImplementer is an interface for InputEventScreenDrag objects.
*/
type InputEventScreenDragImplementer interface {
	Class
}
