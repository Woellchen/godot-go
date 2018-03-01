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

 */
type EditorFileDialog struct {
	ConfirmationDialog
}

func (o *EditorFileDialog) BaseClass() string {
	return "EditorFileDialog"
}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_ActionPressed() {
	log.Println("Calling EditorFileDialog.X_ActionPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_action_pressed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_CancelPressed() {
	log.Println("Calling EditorFileDialog.X_CancelPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_cancel_pressed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_DirEntered(arg0 gdnative.String) {
	log.Println("Calling EditorFileDialog.X_DirEntered()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_dir_entered", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_FavoriteMoveDown() {
	log.Println("Calling EditorFileDialog.X_FavoriteMoveDown()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_favorite_move_down", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_FavoriteMoveUp() {
	log.Println("Calling EditorFileDialog.X_FavoriteMoveUp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_favorite_move_up", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_FavoriteSelected(arg0 gdnative.Int) {
	log.Println("Calling EditorFileDialog.X_FavoriteSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_favorite_selected", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_FavoriteToggled(arg0 gdnative.Bool) {
	log.Println("Calling EditorFileDialog.X_FavoriteToggled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_favorite_toggled", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_FileEntered(arg0 gdnative.String) {
	log.Println("Calling EditorFileDialog.X_FileEntered()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_file_entered", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_FilterSelected(arg0 gdnative.Int) {
	log.Println("Calling EditorFileDialog.X_FilterSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_filter_selected", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_GoBack() {
	log.Println("Calling EditorFileDialog.X_GoBack()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_go_back", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_GoForward() {
	log.Println("Calling EditorFileDialog.X_GoForward()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_go_forward", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_GoUp() {
	log.Println("Calling EditorFileDialog.X_GoUp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_go_up", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_ItemDbSelected(arg0 gdnative.Int) {
	log.Println("Calling EditorFileDialog.X_ItemDbSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_item_db_selected", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_ItemListItemRmbSelected(arg0 gdnative.Int, arg1 *Vector2) {
	log.Println("Calling EditorFileDialog.X_ItemListItemRmbSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(arg0)
	goArguments[1] = reflect.ValueOf(arg1)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_item_list_item_rmb_selected", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_ItemListRmbClicked(arg0 *Vector2) {
	log.Println("Calling EditorFileDialog.X_ItemListRmbClicked()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_item_list_rmb_clicked", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_ItemMenuIdPressed(arg0 gdnative.Int) {
	log.Println("Calling EditorFileDialog.X_ItemMenuIdPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_item_menu_id_pressed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_ItemSelected(arg0 gdnative.Int) {
	log.Println("Calling EditorFileDialog.X_ItemSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_item_selected", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_ItemsClearSelection() {
	log.Println("Calling EditorFileDialog.X_ItemsClearSelection()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_items_clear_selection", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_MakeDir() {
	log.Println("Calling EditorFileDialog.X_MakeDir()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_make_dir", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_MakeDirConfirm() {
	log.Println("Calling EditorFileDialog.X_MakeDirConfirm()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_make_dir_confirm", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_MultiSelected(arg0 gdnative.Int, arg1 gdnative.Bool) {
	log.Println("Calling EditorFileDialog.X_MultiSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(arg0)
	goArguments[1] = reflect.ValueOf(arg1)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_multi_selected", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_RecentSelected(arg0 gdnative.Int) {
	log.Println("Calling EditorFileDialog.X_RecentSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_recent_selected", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_SaveConfirmPressed() {
	log.Println("Calling EditorFileDialog.X_SaveConfirmPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_save_confirm_pressed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_SelectDrive(arg0 gdnative.Int) {
	log.Println("Calling EditorFileDialog.X_SelectDrive()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_select_drive", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_ThumbnailDone(arg0 gdnative.String, arg1 *Texture, arg2 *Variant) {
	log.Println("Calling EditorFileDialog.X_ThumbnailDone()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(arg0)
	goArguments[1] = reflect.ValueOf(arg1)
	goArguments[2] = reflect.ValueOf(arg2)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_thumbnail_done", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_ThumbnailResult(arg0 gdnative.String, arg1 *Texture, arg2 *Variant) {
	log.Println("Calling EditorFileDialog.X_ThumbnailResult()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(arg0)
	goArguments[1] = reflect.ValueOf(arg1)
	goArguments[2] = reflect.ValueOf(arg2)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_thumbnail_result", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_UnhandledInput(arg0 *InputEvent) {
	log.Println("Calling EditorFileDialog.X_UnhandledInput()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_unhandled_input", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_UpdateDir() {
	log.Println("Calling EditorFileDialog.X_UpdateDir()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_update_dir", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) X_UpdateFileList() {
	log.Println("Calling EditorFileDialog.X_UpdateFileList()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_update_file_list", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a comma-delimited file extension filter option to the [code]EditorFileDialog[/code] with an optional semi-colon-delimited label. Example: "*.tscn, *.scn; Scenes", results in filter text "Scenes (*.tscn, *.scn)".
*/
func (o *EditorFileDialog) AddFilter(filter gdnative.String) {
	log.Println("Calling EditorFileDialog.AddFilter()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(filter)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_filter", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes all filters except for "All Files (*)".
*/
func (o *EditorFileDialog) ClearFilters() {
	log.Println("Calling EditorFileDialog.ClearFilters()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "clear_filters", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) GetAccess() gdnative.Int {
	log.Println("Calling EditorFileDialog.GetAccess()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_access", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *EditorFileDialog) GetCurrentDir() gdnative.String {
	log.Println("Calling EditorFileDialog.GetCurrentDir()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_current_dir", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *EditorFileDialog) GetCurrentFile() gdnative.String {
	log.Println("Calling EditorFileDialog.GetCurrentFile()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_current_file", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *EditorFileDialog) GetCurrentPath() gdnative.String {
	log.Println("Calling EditorFileDialog.GetCurrentPath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_current_path", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *EditorFileDialog) GetDisplayMode() gdnative.Int {
	log.Println("Calling EditorFileDialog.GetDisplayMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_display_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *EditorFileDialog) GetMode() gdnative.Int {
	log.Println("Calling EditorFileDialog.GetMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the [code]VBoxContainer[/code] used to display the file system.
*/
func (o *EditorFileDialog) GetVbox() *VBoxContainer {
	log.Println("Calling EditorFileDialog.GetVbox()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_vbox", goArguments, "*VBoxContainer")

	returnValue := goRet.Interface().(*VBoxContainer)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Notify the [code]EditorFileDialog[/code] that its view of the data is no longer accurate. Updates the view contents on next view update.
*/
func (o *EditorFileDialog) Invalidate() {
	log.Println("Calling EditorFileDialog.Invalidate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "invalidate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) IsOverwriteWarningDisabled() gdnative.Bool {
	log.Println("Calling EditorFileDialog.IsOverwriteWarningDisabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_overwrite_warning_disabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *EditorFileDialog) IsShowingHiddenFiles() gdnative.Bool {
	log.Println("Calling EditorFileDialog.IsShowingHiddenFiles()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_showing_hidden_files", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *EditorFileDialog) SetAccess(access gdnative.Int) {
	log.Println("Calling EditorFileDialog.SetAccess()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(access)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_access", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) SetCurrentDir(dir gdnative.String) {
	log.Println("Calling EditorFileDialog.SetCurrentDir()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(dir)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_current_dir", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) SetCurrentFile(file gdnative.String) {
	log.Println("Calling EditorFileDialog.SetCurrentFile()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(file)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_current_file", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) SetCurrentPath(path gdnative.String) {
	log.Println("Calling EditorFileDialog.SetCurrentPath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(path)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_current_path", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) SetDisableOverwriteWarning(disable gdnative.Bool) {
	log.Println("Calling EditorFileDialog.SetDisableOverwriteWarning()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(disable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_disable_overwrite_warning", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) SetDisplayMode(mode gdnative.Int) {
	log.Println("Calling EditorFileDialog.SetDisplayMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_display_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) SetMode(mode gdnative.Int) {
	log.Println("Calling EditorFileDialog.SetMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *EditorFileDialog) SetShowHiddenFiles(show gdnative.Bool) {
	log.Println("Calling EditorFileDialog.SetShowHiddenFiles()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(show)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_show_hidden_files", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   EditorFileDialogImplementer is an interface for EditorFileDialog objects.
*/
type EditorFileDialogImplementer interface {
	Class
}
