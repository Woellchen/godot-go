package editor

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
This object holds information of all resources in the filesystem, their types, etc.
*/
type EditorFileSystem struct {
	Node
}

func (o *EditorFileSystem) BaseClass() string {
	return "EditorFileSystem"
}

/*
   Get the type of the file, given the full path.
*/
func (o *EditorFileSystem) GetFileType(path gdnative.String) gdnative.String {
	log.Println("Calling EditorFileSystem.GetFileType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(path)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_file_type", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the root directory object.
*/
func (o *EditorFileSystem) GetFilesystem() *EditorFileSystemDirectory {
	log.Println("Calling EditorFileSystem.GetFilesystem()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_filesystem", goArguments, "*EditorFileSystemDirectory")

	returnValue := goRet.Interface().(*EditorFileSystemDirectory)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a view into the filesystem at [code]path[/code].
*/
func (o *EditorFileSystem) GetFilesystemPath(path gdnative.String) *EditorFileSystemDirectory {
	log.Println("Calling EditorFileSystem.GetFilesystemPath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(path)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_filesystem_path", goArguments, "*EditorFileSystemDirectory")

	returnValue := goRet.Interface().(*EditorFileSystemDirectory)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the scan progress for 0 to 1 if the FS is being scanned.
*/
func (o *EditorFileSystem) GetScanningProgress() gdnative.Float {
	log.Println("Calling EditorFileSystem.GetScanningProgress()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_scanning_progress", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return true of the filesystem is being scanned.
*/
func (o *EditorFileSystem) IsScanning() gdnative.Bool {
	log.Println("Calling EditorFileSystem.IsScanning()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_scanning", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Scan the filesystem for changes.
*/
func (o *EditorFileSystem) Scan() {
	log.Println("Calling EditorFileSystem.Scan()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "scan", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Check if the source of any imported resource changed.
*/
func (o *EditorFileSystem) ScanSources() {
	log.Println("Calling EditorFileSystem.ScanSources()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "scan_sources", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Update a file information. Call this if an external program (not Godot) modified the file.
*/
func (o *EditorFileSystem) UpdateFile(path gdnative.String) {
	log.Println("Calling EditorFileSystem.UpdateFile()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(path)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "update_file", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   EditorFileSystemImplementer is an interface for EditorFileSystem objects.
*/
type EditorFileSystemImplementer interface {
	Class
}
