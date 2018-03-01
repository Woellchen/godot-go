package colorrect

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
An object that is represented on the canvas as a rect with color. [Color] is used to set or get color info for the rect.
*/
type ColorRect struct {
	Control
}

func (o *ColorRect) BaseClass() string {
	return "ColorRect"
}

/*
   Undocumented
*/
func (o *ColorRect) GetFrameColor() *Color {
	log.Println("Calling ColorRect.GetFrameColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_frame_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ColorRect) SetFrameColor(color *Color) {
	log.Println("Calling ColorRect.SetFrameColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_frame_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ColorRectImplementer is an interface for ColorRect objects.
*/
type ColorRectImplementer interface {
	Class
}
