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

//func NewmarshallsFromPointer(ptr gdnative.Pointer) marshalls {
func new_MarshallsFromPointer(ptr gdnative.Pointer) marshalls {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := marshalls{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonMarshalls() *marshalls {
	return &marshalls{}
}

/*
   Provides data transformation and encoding utility functions.
*/
var Marshalls = newSingletonMarshalls()

/*
Provides data transformation and encoding utility functions.
*/
type marshalls struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *marshalls) ensureSingleton() {
	if o.initialized == true {
		return
	}
	//log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("Marshalls")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *marshalls) BaseClass() string {
	return "_Marshalls"
}

/*
        Undocumented
	Args: [{ false base64_str String}], Returns: PoolByteArray
*/
func (o *marshalls) Base64ToRaw(base64Str gdnative.String) gdnative.PoolByteArray {
	o.ensureSingleton()
	//log.Println("Calling _Marshalls.Base64ToRaw()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(base64Str)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Marshalls", "base64_to_raw")

	// Call the parent method.
	// PoolByteArray
	retPtr := gdnative.NewEmptyPoolByteArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolByteArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false base64_str String}], Returns: String
*/
func (o *marshalls) Base64ToUtf8(base64Str gdnative.String) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling _Marshalls.Base64ToUtf8()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(base64Str)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Marshalls", "base64_to_utf8")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false base64_str String} {False true allow_objects bool}], Returns: Variant
*/
func (o *marshalls) Base64ToVariant(base64Str gdnative.String, allowObjects gdnative.Bool) gdnative.Variant {
	o.ensureSingleton()
	//log.Println("Calling _Marshalls.Base64ToVariant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(base64Str)
	ptrArguments[1] = gdnative.NewPointerFromBool(allowObjects)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Marshalls", "base64_to_variant")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false array PoolByteArray}], Returns: String
*/
func (o *marshalls) RawToBase64(array gdnative.PoolByteArray) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling _Marshalls.RawToBase64()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolByteArray(array)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Marshalls", "raw_to_base64")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false utf8_str String}], Returns: String
*/
func (o *marshalls) Utf8ToBase64(utf8Str gdnative.String) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling _Marshalls.Utf8ToBase64()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(utf8Str)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Marshalls", "utf8_to_base64")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false variant Variant} {False true full_objects bool}], Returns: String
*/
func (o *marshalls) VariantToBase64(variant gdnative.Variant, fullObjects gdnative.Bool) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling _Marshalls.VariantToBase64()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVariant(variant)
	ptrArguments[1] = gdnative.NewPointerFromBool(fullObjects)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Marshalls", "variant_to_base64")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

// MarshallsImplementer is an interface that implements the methods
// of the Marshalls class.
type MarshallsImplementer interface {
	ObjectImplementer
	Base64ToRaw(base64Str gdnative.String) gdnative.PoolByteArray
	Base64ToUtf8(base64Str gdnative.String) gdnative.String
	Base64ToVariant(base64Str gdnative.String, allowObjects gdnative.Bool) gdnative.Variant
	RawToBase64(array gdnative.PoolByteArray) gdnative.String
	Utf8ToBase64(utf8Str gdnative.String) gdnative.String
	VariantToBase64(variant gdnative.Variant, fullObjects gdnative.Bool) gdnative.String
}
