package confirmationdialog

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
Dialog for confirmation of actions. This dialog inherits from [AcceptDialog], but has by default an OK and Cancel button (in host OS order).
*/
type ConfirmationDialog struct {
	AcceptDialog
}

func (o *ConfirmationDialog) BaseClass() string {
	return "ConfirmationDialog"
}

/*
   Return the cancel button.
*/
func (o *ConfirmationDialog) GetCancel() *Button {
	log.Println("Calling ConfirmationDialog.GetCancel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_cancel", goArguments, "*Button")

	returnValue := goRet.Interface().(*Button)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   ConfirmationDialogImplementer is an interface for ConfirmationDialog objects.
*/
type ConfirmationDialogImplementer interface {
	Class
}
