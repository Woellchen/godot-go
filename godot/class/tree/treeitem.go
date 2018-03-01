package tree

import (
	"log"
	"reflect"
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
Control for a single item inside a [Tree]. May have child [code]TreeItem[/code]s and be styled as well as contain buttons.
*/
type TreeItem struct {
	Object
}

func (o *TreeItem) BaseClass() string {
	return "TreeItem"
}

/*
   Adds a button with [Texture] [code]button[/code] at column [code]column[/code]. The [code]button_idx[/code] index is used to identify the button when calling other methods. If not specified, the next available index is used, which may be retrieved by calling [code]get_buton_count()[/code] immediately after this method. Optionally, the button can be [code]disabled[/code] and have a [code]tooltip[/code].
*/
func (o *TreeItem) AddButton(column gdnative.Int, button *Texture, buttonIdx gdnative.Int, disabled gdnative.Bool, tooltip gdnative.String) {
	log.Println("Calling TreeItem.AddButton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(button)
	goArguments[2] = reflect.ValueOf(buttonIdx)
	goArguments[3] = reflect.ValueOf(disabled)
	goArguments[4] = reflect.ValueOf(tooltip)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_button", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Resets the background color for the given column to default.
*/
func (o *TreeItem) ClearCustomBgColor(column gdnative.Int) {
	log.Println("Calling TreeItem.ClearCustomBgColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "clear_custom_bg_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Resets the color for the given column to default.
*/
func (o *TreeItem) ClearCustomColor(column gdnative.Int) {
	log.Println("Calling TreeItem.ClearCustomColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "clear_custom_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Deselects the given column.
*/
func (o *TreeItem) Deselect(column gdnative.Int) {
	log.Println("Calling TreeItem.Deselect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "deselect", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes the button at index [code]button_idx[/code] in column [code]column[/code].
*/
func (o *TreeItem) EraseButton(column gdnative.Int, buttonIdx gdnative.Int) {
	log.Println("Calling TreeItem.EraseButton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(buttonIdx)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "erase_button", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns the [Texture] of the button at index [code]button_idx[/code] in column [code]column[/code].
*/
func (o *TreeItem) GetButton(column gdnative.Int, buttonIdx gdnative.Int) *Texture {
	log.Println("Calling TreeItem.GetButton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(buttonIdx)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_button", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the number of buttons in column [code]column[/code]. May be used to get the most recently added button's index, if no index was specified.
*/
func (o *TreeItem) GetButtonCount(column gdnative.Int) gdnative.Int {
	log.Println("Calling TreeItem.GetButtonCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_button_count", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the column's cell mode. See [code]CELL_MODE_*[/code] constants.
*/
func (o *TreeItem) GetCellMode(column gdnative.Int) gdnative.Int {
	log.Println("Calling TreeItem.GetCellMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_cell_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the TreeItem's child items.
*/
func (o *TreeItem) GetChildren() *TreeItem {
	log.Println("Calling TreeItem.GetChildren()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_children", goArguments, "*TreeItem")

	returnValue := goRet.Interface().(*TreeItem)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the custom background color of column [code]column[/code].
*/
func (o *TreeItem) GetCustomBgColor(column gdnative.Int) *Color {
	log.Println("Calling TreeItem.GetCustomBgColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_custom_bg_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *TreeItem) GetCustomMinimumHeight() gdnative.Int {
	log.Println("Calling TreeItem.GetCustomMinimumHeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_custom_minimum_height", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if [code]expand_right[/code] is set.
*/
func (o *TreeItem) GetExpandRight(column gdnative.Int) gdnative.Bool {
	log.Println("Calling TreeItem.GetExpandRight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_expand_right", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the given column's icon [Texture]. Error if no icon is set.
*/
func (o *TreeItem) GetIcon(column gdnative.Int) *Texture {
	log.Println("Calling TreeItem.GetIcon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_icon", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the column's icon's maximum width.
*/
func (o *TreeItem) GetIconMaxWidth(column gdnative.Int) gdnative.Int {
	log.Println("Calling TreeItem.GetIconMaxWidth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_icon_max_width", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the icon [Texture] region as [Rect2].
*/
func (o *TreeItem) GetIconRegion(column gdnative.Int) *Rect2 {
	log.Println("Calling TreeItem.GetIconRegion()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_icon_region", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *TreeItem) GetMetadata(column gdnative.Int) *Variant {
	log.Println("Calling TreeItem.GetMetadata()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_metadata", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the next TreeItem in the tree.
*/
func (o *TreeItem) GetNext() *TreeItem {
	log.Println("Calling TreeItem.GetNext()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_next", goArguments, "*TreeItem")

	returnValue := goRet.Interface().(*TreeItem)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the next visible TreeItem in the tree.
*/
func (o *TreeItem) GetNextVisible() *TreeItem {
	log.Println("Calling TreeItem.GetNextVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_next_visible", goArguments, "*TreeItem")

	returnValue := goRet.Interface().(*TreeItem)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the parent TreeItem.
*/
func (o *TreeItem) GetParent() *TreeItem {
	log.Println("Calling TreeItem.GetParent()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_parent", goArguments, "*TreeItem")

	returnValue := goRet.Interface().(*TreeItem)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the previous TreeItem in the tree.
*/
func (o *TreeItem) GetPrev() *TreeItem {
	log.Println("Calling TreeItem.GetPrev()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_prev", goArguments, "*TreeItem")

	returnValue := goRet.Interface().(*TreeItem)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the previous visible TreeItem in the tree.
*/
func (o *TreeItem) GetPrevVisible() *TreeItem {
	log.Println("Calling TreeItem.GetPrevVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_prev_visible", goArguments, "*TreeItem")

	returnValue := goRet.Interface().(*TreeItem)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *TreeItem) GetRange(column gdnative.Int) gdnative.Float {
	log.Println("Calling TreeItem.GetRange()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_range", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *TreeItem) GetRangeConfig(column gdnative.Int) *Dictionary {
	log.Println("Calling TreeItem.GetRangeConfig()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_range_config", goArguments, "*Dictionary")

	returnValue := goRet.Interface().(*Dictionary)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the given column's text.
*/
func (o *TreeItem) GetText(column gdnative.Int) gdnative.String {
	log.Println("Calling TreeItem.GetText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_text", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the given column's text alignment.
*/
func (o *TreeItem) GetTextAlign(column gdnative.Int) gdnative.Int {
	log.Println("Calling TreeItem.GetTextAlign()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_text_align", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the given column's tooltip.
*/
func (o *TreeItem) GetTooltip(column gdnative.Int) gdnative.String {
	log.Println("Calling TreeItem.GetTooltip()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_tooltip", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the button at index [code]button_idx[/code] for the given column is disabled.
*/
func (o *TreeItem) IsButtonDisabled(column gdnative.Int, buttonIdx gdnative.Int) gdnative.Bool {
	log.Println("Calling TreeItem.IsButtonDisabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(buttonIdx)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_button_disabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the given column is checked.
*/
func (o *TreeItem) IsChecked(column gdnative.Int) gdnative.Bool {
	log.Println("Calling TreeItem.IsChecked()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_checked", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *TreeItem) IsCollapsed() gdnative.Bool {
	log.Println("Calling TreeItem.IsCollapsed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_collapsed", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *TreeItem) IsCustomSetAsButton(column gdnative.Int) gdnative.Bool {
	log.Println("Calling TreeItem.IsCustomSetAsButton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_custom_set_as_button", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if column [code]column[/code] is editable.
*/
func (o *TreeItem) IsEditable(column gdnative.Int) gdnative.Bool {
	log.Println("Calling TreeItem.IsEditable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_editable", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *TreeItem) IsFoldingDisabled() gdnative.Bool {
	log.Println("Calling TreeItem.IsFoldingDisabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_folding_disabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if column [code]column[/code] is selectable.
*/
func (o *TreeItem) IsSelectable(column gdnative.Int) gdnative.Bool {
	log.Println("Calling TreeItem.IsSelectable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_selectable", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if column [code]column[/code] is selected.
*/
func (o *TreeItem) IsSelected(column gdnative.Int) gdnative.Bool {
	log.Println("Calling TreeItem.IsSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_selected", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Moves this TreeItem to the bottom in the [Tree] hierarchy.
*/
func (o *TreeItem) MoveToBottom() {
	log.Println("Calling TreeItem.MoveToBottom()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "move_to_bottom", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Moves this TreeItem to the top in the [Tree] hierarchy.
*/
func (o *TreeItem) MoveToTop() {
	log.Println("Calling TreeItem.MoveToTop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "move_to_top", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes the child TreeItem at index [code]index[/code].
*/
func (o *TreeItem) RemoveChild(child *Object) {
	log.Println("Calling TreeItem.RemoveChild()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(child)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "remove_child", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Selects the column [code]column[/code].
*/
func (o *TreeItem) Select(column gdnative.Int) {
	log.Println("Calling TreeItem.Select()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "select", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the given column's button [Texture] at index [code]button_idx[/code] to [code]button[/code].
*/
func (o *TreeItem) SetButton(column gdnative.Int, buttonIdx gdnative.Int, button *Texture) {
	log.Println("Calling TreeItem.SetButton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(buttonIdx)
	goArguments[2] = reflect.ValueOf(button)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_button", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the given column's cell mode to [code]mode[/code]. See [code]CELL_MODE_*[/code] constants.
*/
func (o *TreeItem) SetCellMode(column gdnative.Int, mode gdnative.Int) {
	log.Println("Calling TreeItem.SetCellMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_cell_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If [code]true[/code] the column [code]column[/code] is checked.
*/
func (o *TreeItem) SetChecked(column gdnative.Int, checked gdnative.Bool) {
	log.Println("Calling TreeItem.SetChecked()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(checked)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_checked", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *TreeItem) SetCollapsed(enable gdnative.Bool) {
	log.Println("Calling TreeItem.SetCollapsed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collapsed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *TreeItem) SetCustomAsButton(column gdnative.Int, enable gdnative.Bool) {
	log.Println("Calling TreeItem.SetCustomAsButton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_custom_as_button", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the given column's custom background color and whether to just use it as an outline.
*/
func (o *TreeItem) SetCustomBgColor(column gdnative.Int, color *Color, justOutline gdnative.Bool) {
	log.Println("Calling TreeItem.SetCustomBgColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(color)
	goArguments[2] = reflect.ValueOf(justOutline)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_custom_bg_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the given column's custom color.
*/
func (o *TreeItem) SetCustomColor(column gdnative.Int, color *Color) {
	log.Println("Calling TreeItem.SetCustomColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_custom_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the given column's custom draw callback to [code]callback[/code] method on [code]object[/code].
*/
func (o *TreeItem) SetCustomDraw(column gdnative.Int, object *Object, callback gdnative.String) {
	log.Println("Calling TreeItem.SetCustomDraw()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(object)
	goArguments[2] = reflect.ValueOf(callback)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_custom_draw", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *TreeItem) SetCustomMinimumHeight(height gdnative.Int) {
	log.Println("Calling TreeItem.SetCustomMinimumHeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(height)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_custom_minimum_height", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *TreeItem) SetDisableFolding(disable gdnative.Bool) {
	log.Println("Calling TreeItem.SetDisableFolding()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(disable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_disable_folding", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If [code]true[/code] column [code]column[/code] is editable.
*/
func (o *TreeItem) SetEditable(column gdnative.Int, enabled gdnative.Bool) {
	log.Println("Calling TreeItem.SetEditable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_editable", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If [code]true[/code] column [code]column[/code] is expanded to the right.
*/
func (o *TreeItem) SetExpandRight(column gdnative.Int, enable gdnative.Bool) {
	log.Println("Calling TreeItem.SetExpandRight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_expand_right", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the given column's icon [Texture].
*/
func (o *TreeItem) SetIcon(column gdnative.Int, texture *Texture) {
	log.Println("Calling TreeItem.SetIcon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_icon", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the given column's icon's maximum width.
*/
func (o *TreeItem) SetIconMaxWidth(column gdnative.Int, width gdnative.Int) {
	log.Println("Calling TreeItem.SetIconMaxWidth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(width)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_icon_max_width", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the given column's icon's texture region.
*/
func (o *TreeItem) SetIconRegion(column gdnative.Int, region *Rect2) {
	log.Println("Calling TreeItem.SetIconRegion()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(region)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_icon_region", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *TreeItem) SetMetadata(column gdnative.Int, meta *Variant) {
	log.Println("Calling TreeItem.SetMetadata()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(meta)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_metadata", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *TreeItem) SetRange(column gdnative.Int, value gdnative.Float) {
	log.Println("Calling TreeItem.SetRange()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_range", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *TreeItem) SetRangeConfig(column gdnative.Int, min gdnative.Float, max gdnative.Float, step gdnative.Float, expr gdnative.Bool) {
	log.Println("Calling TreeItem.SetRangeConfig()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(min)
	goArguments[2] = reflect.ValueOf(max)
	goArguments[3] = reflect.ValueOf(step)
	goArguments[4] = reflect.ValueOf(expr)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_range_config", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If [code]true[/code] the given column is selectable.
*/
func (o *TreeItem) SetSelectable(column gdnative.Int, selectable gdnative.Bool) {
	log.Println("Calling TreeItem.SetSelectable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(selectable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_selectable", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *TreeItem) SetText(column gdnative.Int, text gdnative.String) {
	log.Println("Calling TreeItem.SetText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(text)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_text", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the given column's text alignment. See [code]ALIGN_*[/code] constants.
*/
func (o *TreeItem) SetTextAlign(column gdnative.Int, textAlign gdnative.Int) {
	log.Println("Calling TreeItem.SetTextAlign()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(textAlign)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_text_align", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the given column's tooltip text.
*/
func (o *TreeItem) SetTooltip(column gdnative.Int, tooltip gdnative.String) {
	log.Println("Calling TreeItem.SetTooltip()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(tooltip)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_tooltip", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   TreeItemImplementer is an interface for TreeItem objects.
*/
type TreeItemImplementer interface {
	Class
}
