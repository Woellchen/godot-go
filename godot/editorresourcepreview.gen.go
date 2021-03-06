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

//func NewEditorResourcePreviewFromPointer(ptr gdnative.Pointer) EditorResourcePreview {
func newEditorResourcePreviewFromPointer(ptr gdnative.Pointer) EditorResourcePreview {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorResourcePreview{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This object is used to generate previews for resources of files. [b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_resource_previewer].
*/
type EditorResourcePreview struct {
	Node
	owner gdnative.Object
}

func (o *EditorResourcePreview) BaseClass() string {
	return "EditorResourcePreview"
}

/*
        Undocumented
	Args: [{ false arg0 String} { false arg1 Texture} { false arg2 Texture} { false arg3 int} { false arg4 String} { false arg5 Variant}], Returns: void
*/
func (o *EditorResourcePreview) X_PreviewReady(arg0 gdnative.String, arg1 TextureImplementer, arg2 TextureImplementer, arg3 gdnative.Int, arg4 gdnative.String, arg5 gdnative.Variant) {
	// log.Println("Calling EditorResourcePreview.X_PreviewReady()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 6, 6)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)
	ptrArguments[1] = gdnative.NewPointerFromObject(arg1.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromObject(arg2.GetBaseObject())
	ptrArguments[3] = gdnative.NewPointerFromInt(arg3)
	ptrArguments[4] = gdnative.NewPointerFromString(arg4)
	ptrArguments[5] = gdnative.NewPointerFromVariant(arg5)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorResourcePreview", "_preview_ready")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Create an own, custom preview generator.
	Args: [{ false generator EditorResourcePreviewGenerator}], Returns: void
*/
func (o *EditorResourcePreview) AddPreviewGenerator(generator EditorResourcePreviewGeneratorImplementer) {
	// log.Println("Calling EditorResourcePreview.AddPreviewGenerator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(generator.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorResourcePreview", "add_preview_generator")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Check if the resource changed, if so, it will be invalidated and the corresponding signal emitted.
	Args: [{ false path String}], Returns: void
*/
func (o *EditorResourcePreview) CheckForInvalidation(path gdnative.String) {
	// log.Println("Calling EditorResourcePreview.CheckForInvalidation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorResourcePreview", "check_for_invalidation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Queue a resource being edited for preview (using an instance). Once the preview is ready, your receiver.receiver_func will be called either containing the preview texture or an empty texture (if no preview was possible). Callback must have the format: (path,texture,userdata). Userdata can be anything.
	Args: [{ false resource Resource} { false receiver Object} { false receiver_func String} { false userdata Variant}], Returns: void
*/
func (o *EditorResourcePreview) QueueEditedResourcePreview(resource ResourceImplementer, receiver ObjectImplementer, receiverFunc gdnative.String, userdata gdnative.Variant) {
	// log.Println("Calling EditorResourcePreview.QueueEditedResourcePreview()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(resource.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromObject(receiver.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromString(receiverFunc)
	ptrArguments[3] = gdnative.NewPointerFromVariant(userdata)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorResourcePreview", "queue_edited_resource_preview")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Queue a resource file for preview (using a path). Once the preview is ready, your receiver.receiver_func will be called either containing the preview texture or an empty texture (if no preview was possible). Callback must have the format: (path,texture,userdata). Userdata can be anything.
	Args: [{ false path String} { false receiver Object} { false receiver_func String} { false userdata Variant}], Returns: void
*/
func (o *EditorResourcePreview) QueueResourcePreview(path gdnative.String, receiver ObjectImplementer, receiverFunc gdnative.String, userdata gdnative.Variant) {
	// log.Println("Calling EditorResourcePreview.QueueResourcePreview()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromObject(receiver.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromString(receiverFunc)
	ptrArguments[3] = gdnative.NewPointerFromVariant(userdata)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorResourcePreview", "queue_resource_preview")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes a custom preview generator.
	Args: [{ false generator EditorResourcePreviewGenerator}], Returns: void
*/
func (o *EditorResourcePreview) RemovePreviewGenerator(generator EditorResourcePreviewGeneratorImplementer) {
	// log.Println("Calling EditorResourcePreview.RemovePreviewGenerator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(generator.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorResourcePreview", "remove_preview_generator")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EditorResourcePreviewImplementer is an interface that implements the methods
// of the EditorResourcePreview class.
type EditorResourcePreviewImplementer interface {
	NodeImplementer
	X_PreviewReady(arg0 gdnative.String, arg1 TextureImplementer, arg2 TextureImplementer, arg3 gdnative.Int, arg4 gdnative.String, arg5 gdnative.Variant)
	AddPreviewGenerator(generator EditorResourcePreviewGeneratorImplementer)
	CheckForInvalidation(path gdnative.String)
	QueueEditedResourcePreview(resource ResourceImplementer, receiver ObjectImplementer, receiverFunc gdnative.String, userdata gdnative.Variant)
	QueueResourcePreview(path gdnative.String, receiver ObjectImplementer, receiverFunc gdnative.String, userdata gdnative.Variant)
	RemovePreviewGenerator(generator EditorResourcePreviewGeneratorImplementer)
}
