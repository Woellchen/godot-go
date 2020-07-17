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

//func NewEditorExportPluginFromPointer(ptr gdnative.Pointer) EditorExportPlugin {
func newEditorExportPluginFromPointer(ptr gdnative.Pointer) EditorExportPlugin {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorExportPlugin{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type EditorExportPlugin struct {
	Reference
	owner gdnative.Object
}

func (o *EditorExportPlugin) BaseClass() string {
	return "EditorExportPlugin"
}

/*
        Virtual method to be overridden by the user. It is called when the export starts and provides all information about the export.
	Args: [{ false features PoolStringArray} { false is_debug bool} { false path String} { false flags int}], Returns: void
*/
func (o *EditorExportPlugin) X_ExportBegin(features gdnative.PoolStringArray, isDebug gdnative.Bool, path gdnative.String, flags gdnative.Int) {
	//log.Println("Calling EditorExportPlugin.X_ExportBegin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromPoolStringArray(features)
	ptrArguments[1] = gdnative.NewPointerFromBool(isDebug)
	ptrArguments[2] = gdnative.NewPointerFromString(path)
	ptrArguments[3] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "_export_begin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Virtual method to be overridden by the user. Called when the export is finished.
	Args: [], Returns: void
*/
func (o *EditorExportPlugin) X_ExportEnd() {
	//log.Println("Calling EditorExportPlugin.X_ExportEnd()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "_export_end")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false path String} { false type String} { false features PoolStringArray}], Returns: void
*/
func (o *EditorExportPlugin) X_ExportFile(path gdnative.String, aType gdnative.String, features gdnative.PoolStringArray) {
	//log.Println("Calling EditorExportPlugin.X_ExportFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)
	ptrArguments[2] = gdnative.NewPointerFromPoolStringArray(features)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "_export_file")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false path String} { false file PoolByteArray} { false remap bool}], Returns: void
*/
func (o *EditorExportPlugin) AddFile(path gdnative.String, file gdnative.PoolByteArray, remap gdnative.Bool) {
	//log.Println("Calling EditorExportPlugin.AddFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromPoolByteArray(file)
	ptrArguments[2] = gdnative.NewPointerFromBool(remap)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "add_file")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false path String}], Returns: void
*/
func (o *EditorExportPlugin) AddIosBundleFile(path gdnative.String) {
	//log.Println("Calling EditorExportPlugin.AddIosBundleFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "add_ios_bundle_file")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false code String}], Returns: void
*/
func (o *EditorExportPlugin) AddIosCppCode(code gdnative.String) {
	//log.Println("Calling EditorExportPlugin.AddIosCppCode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(code)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "add_ios_cpp_code")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false path String}], Returns: void
*/
func (o *EditorExportPlugin) AddIosFramework(path gdnative.String) {
	//log.Println("Calling EditorExportPlugin.AddIosFramework()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "add_ios_framework")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false flags String}], Returns: void
*/
func (o *EditorExportPlugin) AddIosLinkerFlags(flags gdnative.String) {
	//log.Println("Calling EditorExportPlugin.AddIosLinkerFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "add_ios_linker_flags")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false plist_content String}], Returns: void
*/
func (o *EditorExportPlugin) AddIosPlistContent(plistContent gdnative.String) {
	//log.Println("Calling EditorExportPlugin.AddIosPlistContent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(plistContent)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "add_ios_plist_content")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false path String}], Returns: void
*/
func (o *EditorExportPlugin) AddIosProjectStaticLib(path gdnative.String) {
	//log.Println("Calling EditorExportPlugin.AddIosProjectStaticLib()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "add_ios_project_static_lib")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false path String} { false tags PoolStringArray}], Returns: void
*/
func (o *EditorExportPlugin) AddSharedObject(path gdnative.String, tags gdnative.PoolStringArray) {
	//log.Println("Calling EditorExportPlugin.AddSharedObject()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromPoolStringArray(tags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "add_shared_object")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: void
*/
func (o *EditorExportPlugin) Skip() {
	//log.Println("Calling EditorExportPlugin.Skip()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorExportPlugin", "skip")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EditorExportPluginImplementer is an interface that implements the methods
// of the EditorExportPlugin class.
type EditorExportPluginImplementer interface {
	ReferenceImplementer
	X_ExportBegin(features gdnative.PoolStringArray, isDebug gdnative.Bool, path gdnative.String, flags gdnative.Int)
	X_ExportEnd()
	X_ExportFile(path gdnative.String, aType gdnative.String, features gdnative.PoolStringArray)
	AddFile(path gdnative.String, file gdnative.PoolByteArray, remap gdnative.Bool)
	AddIosBundleFile(path gdnative.String)
	AddIosCppCode(code gdnative.String)
	AddIosFramework(path gdnative.String)
	AddIosLinkerFlags(flags gdnative.String)
	AddIosPlistContent(plistContent gdnative.String)
	AddIosProjectStaticLib(path gdnative.String)
	AddSharedObject(path gdnative.String, tags gdnative.PoolStringArray)
	Skip()
}
