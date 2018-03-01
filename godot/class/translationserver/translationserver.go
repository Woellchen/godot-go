package translationserver

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

func newSingletonTranslationServer() *translationServer {
	obj := &translationServer{}
	ptr := C.godot_global_get_singleton(C.CString("TranslationServer"))
	obj.owner = (*C.godot_object)(ptr)
	return obj
}

/*

 */
var TranslationServer = newSingletonTranslationServer()

/*

 */
type translationServer struct {
	Object
}

func (o *translationServer) BaseClass() string {
	return "TranslationServer"
}

/*

 */
func (o *translationServer) AddTranslation(translation *Translation) {
	log.Println("Calling TranslationServer.AddTranslation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(translation)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_translation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *translationServer) Clear() {
	log.Println("Calling TranslationServer.Clear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "clear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *translationServer) GetLocale() gdnative.String {
	log.Println("Calling TranslationServer.GetLocale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_locale", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *translationServer) GetLocaleName(locale gdnative.String) gdnative.String {
	log.Println("Calling TranslationServer.GetLocaleName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(locale)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_locale_name", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *translationServer) RemoveTranslation(translation *Translation) {
	log.Println("Calling TranslationServer.RemoveTranslation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(translation)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "remove_translation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *translationServer) SetLocale(locale gdnative.String) {
	log.Println("Calling TranslationServer.SetLocale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(locale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_locale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *translationServer) Translate(message gdnative.String) gdnative.String {
	log.Println("Calling TranslationServer.Translate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(message)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "translate", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}
