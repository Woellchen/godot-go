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

//func NewOptionButtonFromPointer(ptr gdnative.Pointer) OptionButton {
func newOptionButtonFromPointer(ptr gdnative.Pointer) OptionButton {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := OptionButton{}
	obj.SetBaseObject(owner)

	return obj
}

/*
OptionButton is a type button that provides a selectable list of items when pressed. The item selected becomes the "current" item and is displayed as the button text.
*/
type OptionButton struct {
	Button
	owner gdnative.Object
}

func (o *OptionButton) BaseClass() string {
	return "OptionButton"
}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *OptionButton) X_Focused(arg0 gdnative.Int) {
	//log.Println("Calling OptionButton.X_Focused()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "_focused")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *OptionButton) X_GetItems() gdnative.Array {
	//log.Println("Calling OptionButton.X_GetItems()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "_get_items")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *OptionButton) X_SelectInt(arg0 gdnative.Int) {
	//log.Println("Calling OptionButton.X_SelectInt()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "_select_int")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *OptionButton) X_Selected(arg0 gdnative.Int) {
	//log.Println("Calling OptionButton.X_Selected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 Array}], Returns: void
*/
func (o *OptionButton) X_SetItems(arg0 gdnative.Array) {
	//log.Println("Calling OptionButton.X_SetItems()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "_set_items")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds an item, with a [code]texture[/code] icon, text [code]label[/code] and (optionally) [code]id[/code]. If no [code]id[/code] is passed, the item index will be used as the item's ID. New items are appended at the end.
	Args: [{ false texture Texture} { false label String} {-1 true id int}], Returns: void
*/
func (o *OptionButton) AddIconItem(texture TextureImplementer, label gdnative.String, id gdnative.Int) {
	//log.Println("Calling OptionButton.AddIconItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromString(label)
	ptrArguments[2] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "add_icon_item")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds an item, with text [code]label[/code] and (optionally) [code]id[/code]. If no [code]id[/code] is passed, the item index will be used as the item's ID. New items are appended at the end.
	Args: [{ false label String} {-1 true id int}], Returns: void
*/
func (o *OptionButton) AddItem(label gdnative.String, id gdnative.Int) {
	//log.Println("Calling OptionButton.AddItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(label)
	ptrArguments[1] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "add_item")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a separator to the list of items. Separators help to group items. Separator also takes up an index and is appended at the end.
	Args: [], Returns: void
*/
func (o *OptionButton) AddSeparator() {
	//log.Println("Calling OptionButton.AddSeparator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "add_separator")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clears all the items in the [OptionButton].
	Args: [], Returns: void
*/
func (o *OptionButton) Clear() {
	//log.Println("Calling OptionButton.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the amount of items in the OptionButton, including separators.
	Args: [], Returns: int
*/
func (o *OptionButton) GetItemCount() gdnative.Int {
	//log.Println("Calling OptionButton.GetItemCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "get_item_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the icon of the item at index [code]idx[/code].
	Args: [{ false idx int}], Returns: Texture
*/
func (o *OptionButton) GetItemIcon(idx gdnative.Int) TextureImplementer {
	//log.Println("Calling OptionButton.GetItemIcon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "get_item_icon")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Returns the ID of the item at index [code]idx[/code].
	Args: [{ false idx int}], Returns: int
*/
func (o *OptionButton) GetItemId(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling OptionButton.GetItemId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "get_item_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the index of the item with the given [code]id[/code].
	Args: [{ false id int}], Returns: int
*/
func (o *OptionButton) GetItemIndex(id gdnative.Int) gdnative.Int {
	//log.Println("Calling OptionButton.GetItemIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "get_item_index")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Retrieves the metadata of an item. Metadata may be any type and can be used to store extra information about an item, such as an external string ID.
	Args: [{ false idx int}], Returns: Variant
*/
func (o *OptionButton) GetItemMetadata(idx gdnative.Int) gdnative.Variant {
	//log.Println("Calling OptionButton.GetItemMetadata()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "get_item_metadata")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Returns the text of the item at index [code]idx[/code].
	Args: [{ false idx int}], Returns: String
*/
func (o *OptionButton) GetItemText(idx gdnative.Int) gdnative.String {
	//log.Println("Calling OptionButton.GetItemText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "get_item_text")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the [PopupMenu] contained in this button.
	Args: [], Returns: PopupMenu
*/
func (o *OptionButton) GetPopup() PopupMenuImplementer {
	//log.Println("Calling OptionButton.GetPopup()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "get_popup")

	// Call the parent method.
	// PopupMenu
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newPopupMenuFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(PopupMenuImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "PopupMenu" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(PopupMenuImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *OptionButton) GetSelected() gdnative.Int {
	//log.Println("Calling OptionButton.GetSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "get_selected")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the ID of the selected item, or [code]0[/code] if no item is selected.
	Args: [], Returns: int
*/
func (o *OptionButton) GetSelectedId() gdnative.Int {
	//log.Println("Calling OptionButton.GetSelectedId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "get_selected_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Gets the metadata of the selected item. Metadata for items can be set using [method set_item_metadata].
	Args: [], Returns: Variant
*/
func (o *OptionButton) GetSelectedMetadata() gdnative.Variant {
	//log.Println("Calling OptionButton.GetSelectedMetadata()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "get_selected_metadata")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the item at index [code]idx[/code] is disabled.
	Args: [{ false idx int}], Returns: bool
*/
func (o *OptionButton) IsItemDisabled(idx gdnative.Int) gdnative.Bool {
	//log.Println("Calling OptionButton.IsItemDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "is_item_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes the item at index [code]idx[/code].
	Args: [{ false idx int}], Returns: void
*/
func (o *OptionButton) RemoveItem(idx gdnative.Int) {
	//log.Println("Calling OptionButton.RemoveItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "remove_item")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Selects an item by index and makes it the current item. This will work even if the item is disabled.
	Args: [{ false idx int}], Returns: void
*/
func (o *OptionButton) Select(idx gdnative.Int) {
	//log.Println("Calling OptionButton.Select()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "select")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets whether the item at index [code]idx[/code] is disabled. Disabled items are drawn differently in the dropdown and are not selectable by the user. If the current selected item is set as disabled, it will remain selected.
	Args: [{ false idx int} { false disabled bool}], Returns: void
*/
func (o *OptionButton) SetItemDisabled(idx gdnative.Int, disabled gdnative.Bool) {
	//log.Println("Calling OptionButton.SetItemDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromBool(disabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "set_item_disabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the icon of the item at index [code]idx[/code].
	Args: [{ false idx int} { false texture Texture}], Returns: void
*/
func (o *OptionButton) SetItemIcon(idx gdnative.Int, texture TextureImplementer) {
	//log.Println("Calling OptionButton.SetItemIcon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "set_item_icon")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the ID of the item at index [code]idx[/code].
	Args: [{ false idx int} { false id int}], Returns: void
*/
func (o *OptionButton) SetItemId(idx gdnative.Int, id gdnative.Int) {
	//log.Println("Calling OptionButton.SetItemId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "set_item_id")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the metadata of an item. Metadata may be of any type and can be used to store extra information about an item, such as an external string ID.
	Args: [{ false idx int} { false metadata Variant}], Returns: void
*/
func (o *OptionButton) SetItemMetadata(idx gdnative.Int, metadata gdnative.Variant) {
	//log.Println("Calling OptionButton.SetItemMetadata()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVariant(metadata)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "set_item_metadata")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the text of the item at index [code]idx[/code].
	Args: [{ false idx int} { false text String}], Returns: void
*/
func (o *OptionButton) SetItemText(idx gdnative.Int, text gdnative.String) {
	//log.Println("Calling OptionButton.SetItemText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OptionButton", "set_item_text")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// OptionButtonImplementer is an interface that implements the methods
// of the OptionButton class.
type OptionButtonImplementer interface {
	ButtonImplementer
	X_Focused(arg0 gdnative.Int)
	X_GetItems() gdnative.Array
	X_SelectInt(arg0 gdnative.Int)
	X_Selected(arg0 gdnative.Int)
	X_SetItems(arg0 gdnative.Array)
	AddIconItem(texture TextureImplementer, label gdnative.String, id gdnative.Int)
	AddItem(label gdnative.String, id gdnative.Int)
	AddSeparator()
	Clear()
	GetItemCount() gdnative.Int
	GetItemIcon(idx gdnative.Int) TextureImplementer
	GetItemId(idx gdnative.Int) gdnative.Int
	GetItemIndex(id gdnative.Int) gdnative.Int
	GetItemMetadata(idx gdnative.Int) gdnative.Variant
	GetItemText(idx gdnative.Int) gdnative.String
	GetPopup() PopupMenuImplementer
	GetSelected() gdnative.Int
	GetSelectedId() gdnative.Int
	GetSelectedMetadata() gdnative.Variant
	IsItemDisabled(idx gdnative.Int) gdnative.Bool
	RemoveItem(idx gdnative.Int)
	Select(idx gdnative.Int)
	SetItemDisabled(idx gdnative.Int, disabled gdnative.Bool)
	SetItemIcon(idx gdnative.Int, texture TextureImplementer)
	SetItemId(idx gdnative.Int, id gdnative.Int)
	SetItemMetadata(idx gdnative.Int, metadata gdnative.Variant)
	SetItemText(idx gdnative.Int, text gdnative.String)
}
