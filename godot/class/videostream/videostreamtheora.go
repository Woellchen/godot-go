package videostream

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
type VideoStreamTheora struct {
	VideoStream
}

func (o *VideoStreamTheora) BaseClass() string {
	return "VideoStreamTheora"
}

/*
   Undocumented
*/
func (o *VideoStreamTheora) GetFile() gdnative.String {
	log.Println("Calling VideoStreamTheora.GetFile()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_file", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VideoStreamTheora) SetFile(file gdnative.String) {
	log.Println("Calling VideoStreamTheora.SetFile()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(file)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_file", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VideoStreamTheoraImplementer is an interface for VideoStreamTheora objects.
*/
type VideoStreamTheoraImplementer interface {
	Class
}
