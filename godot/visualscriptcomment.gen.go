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

//func NewVisualScriptCommentFromPointer(ptr gdnative.Pointer) VisualScriptComment {
func newVisualScriptCommentFromPointer(ptr gdnative.Pointer) VisualScriptComment {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptComment{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptComment struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptComment) BaseClass() string {
	return "VisualScriptComment"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptComment) GetDescription() gdnative.String {
	// log.Println("Calling VisualScriptComment.GetDescription()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptComment", "get_description")

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
	Args: [], Returns: Vector2
*/
func (o *VisualScriptComment) GetSize() gdnative.Vector2 {
	// log.Println("Calling VisualScriptComment.GetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptComment", "get_size")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptComment) GetTitle() gdnative.String {
	// log.Println("Calling VisualScriptComment.GetTitle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptComment", "get_title")

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
	Args: [{ false description String}], Returns: void
*/
func (o *VisualScriptComment) SetDescription(description gdnative.String) {
	// log.Println("Calling VisualScriptComment.SetDescription()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(description)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptComment", "set_description")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false size Vector2}], Returns: void
*/
func (o *VisualScriptComment) SetSize(size gdnative.Vector2) {
	// log.Println("Calling VisualScriptComment.SetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptComment", "set_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false title String}], Returns: void
*/
func (o *VisualScriptComment) SetTitle(title gdnative.String) {
	// log.Println("Calling VisualScriptComment.SetTitle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(title)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptComment", "set_title")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptCommentImplementer is an interface that implements the methods
// of the VisualScriptComment class.
type VisualScriptCommentImplementer interface {
	VisualScriptNodeImplementer
	GetDescription() gdnative.String
	GetSize() gdnative.Vector2
	GetTitle() gdnative.String
	SetDescription(description gdnative.String)
	SetSize(size gdnative.Vector2)
	SetTitle(title gdnative.String)
}
