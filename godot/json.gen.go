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

//func NewjsonFromPointer(ptr gdnative.Pointer) json {
func new_JSONFromPointer(ptr gdnative.Pointer) json {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := json{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonJSON() *json {
	return &json{}
}

/*
   Helper class for parsing JSON data. For usage example and other important hints, see [JSONParseResult].
*/
var JSON = newSingletonJSON()

/*
Helper class for parsing JSON data. For usage example and other important hints, see [JSONParseResult].
*/
type json struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *json) ensureSingleton() {
	if o.initialized == true {
		return
	}
	//log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("JSON")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *json) BaseClass() string {
	return "_JSON"
}

/*
        Undocumented
	Args: [{ false json String}], Returns: JSONParseResult
*/
func (o *json) Parse(json gdnative.String) JSONParseResultImplementer {
	o.ensureSingleton()
	// log.Println("Calling _JSON.Parse()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(json)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_JSON", "parse")

	// Call the parent method.
	// JSONParseResult
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newJSONParseResultFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(JSONParseResultImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "JSONParseResult" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(JSONParseResultImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false value Variant} { true indent String} {False true sort_keys bool}], Returns: String
*/
func (o *json) Print(value gdnative.Variant, indent gdnative.String, sortKeys gdnative.Bool) gdnative.String {
	o.ensureSingleton()
	// log.Println("Calling _JSON.Print()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromVariant(value)
	ptrArguments[1] = gdnative.NewPointerFromString(indent)
	ptrArguments[2] = gdnative.NewPointerFromBool(sortKeys)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_JSON", "print")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

// JSONImplementer is an interface that implements the methods
// of the JSON class.
type JSONImplementer interface {
	ObjectImplementer
	Parse(json gdnative.String) JSONParseResultImplementer
	Print(value gdnative.Variant, indent gdnative.String, sortKeys gdnative.Bool) gdnative.String
}
