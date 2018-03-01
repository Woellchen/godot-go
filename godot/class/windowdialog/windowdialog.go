package windowdialog

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
Windowdialog is the base class for all window-based dialogs. It's a by-default toplevel [Control] that draws a window decoration and allows motion and resizing.
*/
type WindowDialog struct {
	Popup
}

func (o *WindowDialog) BaseClass() string {
	return "WindowDialog"
}

/*
   Undocumented
*/
func (o *WindowDialog) X_Closed() {
	log.Println("Calling WindowDialog.X_Closed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_closed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *WindowDialog) X_GuiInput(arg0 *InputEvent) {
	log.Println("Calling WindowDialog.X_GuiInput()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_gui_input", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the close [TextureButton].
*/
func (o *WindowDialog) GetCloseButton() *TextureButton {
	log.Println("Calling WindowDialog.GetCloseButton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_close_button", goArguments, "*TextureButton")

	returnValue := goRet.Interface().(*TextureButton)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *WindowDialog) GetResizable() gdnative.Bool {
	log.Println("Calling WindowDialog.GetResizable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_resizable", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *WindowDialog) GetTitle() gdnative.String {
	log.Println("Calling WindowDialog.GetTitle()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_title", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *WindowDialog) SetResizable(resizable gdnative.Bool) {
	log.Println("Calling WindowDialog.SetResizable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(resizable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_resizable", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *WindowDialog) SetTitle(title gdnative.String) {
	log.Println("Calling WindowDialog.SetTitle()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(title)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_title", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   WindowDialogImplementer is an interface for WindowDialog objects.
*/
type WindowDialogImplementer interface {
	Class
}
