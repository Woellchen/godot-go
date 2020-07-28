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

//func NewtranslationServerFromPointer(ptr gdnative.Pointer) translationServer {
func newTranslationServerFromPointer(ptr gdnative.Pointer) translationServer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := translationServer{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonTranslationServer() *translationServer {
	return &translationServer{}
}

/*
   Server that manages all translations. Translations can be set to it and removed from it.
*/
var TranslationServer = newSingletonTranslationServer()

/*
Server that manages all translations. Translations can be set to it and removed from it.
*/
type translationServer struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *translationServer) ensureSingleton() {
	if o.initialized == true {
		return
	}
	//log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("TranslationServer")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *translationServer) BaseClass() string {
	return "TranslationServer"
}

/*
        Adds a [Translation] resource.
	Args: [{ false translation Translation}], Returns: void
*/
func (o *translationServer) AddTranslation(translation TranslationImplementer) {
	o.ensureSingleton()
	// log.Println("Calling TranslationServer.AddTranslation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(translation.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TranslationServer", "add_translation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clears the server from all translations.
	Args: [], Returns: void
*/
func (o *translationServer) Clear() {
	o.ensureSingleton()
	// log.Println("Calling TranslationServer.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TranslationServer", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns an Array of all loaded locales of the game.
	Args: [], Returns: Array
*/
func (o *translationServer) GetLoadedLocales() gdnative.Array {
	o.ensureSingleton()
	// log.Println("Calling TranslationServer.GetLoadedLocales()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TranslationServer", "get_loaded_locales")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the current locale of the game.
	Args: [], Returns: String
*/
func (o *translationServer) GetLocale() gdnative.String {
	o.ensureSingleton()
	// log.Println("Calling TranslationServer.GetLocale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TranslationServer", "get_locale")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns a locale's language and its variant (e.g. [code]"en_US"[/code] would return [code]"English (United States)"[/code]).
	Args: [{ false locale String}], Returns: String
*/
func (o *translationServer) GetLocaleName(locale gdnative.String) gdnative.String {
	o.ensureSingleton()
	// log.Println("Calling TranslationServer.GetLocaleName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(locale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TranslationServer", "get_locale_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Removes the given translation from the server.
	Args: [{ false translation Translation}], Returns: void
*/
func (o *translationServer) RemoveTranslation(translation TranslationImplementer) {
	o.ensureSingleton()
	// log.Println("Calling TranslationServer.RemoveTranslation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(translation.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TranslationServer", "remove_translation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the locale of the game.
	Args: [{ false locale String}], Returns: void
*/
func (o *translationServer) SetLocale(locale gdnative.String) {
	o.ensureSingleton()
	// log.Println("Calling TranslationServer.SetLocale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(locale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TranslationServer", "set_locale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the current locale's translation for the given message (key).
	Args: [{ false message String}], Returns: String
*/
func (o *translationServer) Translate(message gdnative.String) gdnative.String {
	o.ensureSingleton()
	// log.Println("Calling TranslationServer.Translate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(message)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TranslationServer", "translate")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

// TranslationServerImplementer is an interface that implements the methods
// of the TranslationServer class.
type TranslationServerImplementer interface {
	ObjectImplementer
	AddTranslation(translation TranslationImplementer)
	Clear()
	GetLoadedLocales() gdnative.Array
	GetLocale() gdnative.String
	GetLocaleName(locale gdnative.String) gdnative.String
	RemoveTranslation(translation TranslationImplementer)
	SetLocale(locale gdnative.String)
	Translate(message gdnative.String) gdnative.String
}
