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

// JSONRPCErrorCode is an enum for ErrorCode values.
type JSONRPCErrorCode int

const (
	JSONRPCInternalError  JSONRPCErrorCode = -32603
	JSONRPCInvalidParams  JSONRPCErrorCode = -32602
	JSONRPCInvalidRequest JSONRPCErrorCode = -32600
	JSONRPCMethodNotFound JSONRPCErrorCode = -32601
	JSONRPCParseError     JSONRPCErrorCode = -32700
)

//func NewJSONRPCFromPointer(ptr gdnative.Pointer) JSONRPC {
func newJSONRPCFromPointer(ptr gdnative.Pointer) JSONRPC {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := JSONRPC{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type JSONRPC struct {
	Object
	owner gdnative.Object
}

func (o *JSONRPC) BaseClass() string {
	return "JSONRPC"
}

/*

	Args: [{ false method String} { false params Variant}], Returns: Dictionary
*/
func (o *JSONRPC) MakeNotification(method gdnative.String, params gdnative.Variant) gdnative.Dictionary {
	// log.Println("Calling JSONRPC.MakeNotification()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(method)
	ptrArguments[1] = gdnative.NewPointerFromVariant(params)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("JSONRPC", "make_notification")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false method String} { false params Variant} { false id Variant}], Returns: Dictionary
*/
func (o *JSONRPC) MakeRequest(method gdnative.String, params gdnative.Variant, id gdnative.Variant) gdnative.Dictionary {
	// log.Println("Calling JSONRPC.MakeRequest()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(method)
	ptrArguments[1] = gdnative.NewPointerFromVariant(params)
	ptrArguments[2] = gdnative.NewPointerFromVariant(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("JSONRPC", "make_request")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false result Variant} { false id Variant}], Returns: Dictionary
*/
func (o *JSONRPC) MakeResponse(result gdnative.Variant, id gdnative.Variant) gdnative.Dictionary {
	// log.Println("Calling JSONRPC.MakeResponse()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVariant(result)
	ptrArguments[1] = gdnative.NewPointerFromVariant(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("JSONRPC", "make_response")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false code int} { false message String} {Null true id Variant}], Returns: Dictionary
*/
func (o *JSONRPC) MakeResponseError(code gdnative.Int, message gdnative.String, id gdnative.Variant) gdnative.Dictionary {
	// log.Println("Calling JSONRPC.MakeResponseError()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(code)
	ptrArguments[1] = gdnative.NewPointerFromString(message)
	ptrArguments[2] = gdnative.NewPointerFromVariant(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("JSONRPC", "make_response_error")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false action Variant} {False true recurse bool}], Returns: Variant
*/
func (o *JSONRPC) ProcessAction(action gdnative.Variant, recurse gdnative.Bool) gdnative.Variant {
	// log.Println("Calling JSONRPC.ProcessAction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVariant(action)
	ptrArguments[1] = gdnative.NewPointerFromBool(recurse)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("JSONRPC", "process_action")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false action String}], Returns: String
*/
func (o *JSONRPC) ProcessString(action gdnative.String) gdnative.String {
	// log.Println("Calling JSONRPC.ProcessString()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("JSONRPC", "process_string")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false scope String} { false target Object}], Returns: void
*/
func (o *JSONRPC) SetScope(scope gdnative.String, target ObjectImplementer) {
	// log.Println("Calling JSONRPC.SetScope()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(scope)
	ptrArguments[1] = gdnative.NewPointerFromObject(target.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("JSONRPC", "set_scope")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// JSONRPCImplementer is an interface that implements the methods
// of the JSONRPC class.
type JSONRPCImplementer interface {
	ObjectImplementer
	MakeNotification(method gdnative.String, params gdnative.Variant) gdnative.Dictionary
	MakeRequest(method gdnative.String, params gdnative.Variant, id gdnative.Variant) gdnative.Dictionary
	MakeResponse(result gdnative.Variant, id gdnative.Variant) gdnative.Dictionary
	MakeResponseError(code gdnative.Int, message gdnative.String, id gdnative.Variant) gdnative.Dictionary
	ProcessAction(action gdnative.Variant, recurse gdnative.Bool) gdnative.Variant
	ProcessString(action gdnative.String) gdnative.String
	SetScope(scope gdnative.String, target ObjectImplementer)
}
