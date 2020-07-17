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

// BaseButtonActionMode is an enum for ActionMode values.
type BaseButtonActionMode int

const (
	BaseButtonActionModeButtonPress   BaseButtonActionMode = 0
	BaseButtonActionModeButtonRelease BaseButtonActionMode = 1
)

// BaseButtonDrawMode is an enum for DrawMode values.
type BaseButtonDrawMode int

const (
	BaseButtonDrawDisabled     BaseButtonDrawMode = 3
	BaseButtonDrawHover        BaseButtonDrawMode = 2
	BaseButtonDrawHoverPressed BaseButtonDrawMode = 4
	BaseButtonDrawNormal       BaseButtonDrawMode = 0
	BaseButtonDrawPressed      BaseButtonDrawMode = 1
)

//func NewBaseButtonFromPointer(ptr gdnative.Pointer) BaseButton {
func newBaseButtonFromPointer(ptr gdnative.Pointer) BaseButton {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := BaseButton{}
	obj.SetBaseObject(owner)

	return obj
}

/*
BaseButton is the abstract base class for buttons, so it shouldn't be used directly (it doesn't display anything). Other types of buttons inherit from it.
*/
type BaseButton struct {
	Control
	owner gdnative.Object
}

func (o *BaseButton) BaseClass() string {
	return "BaseButton"
}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *BaseButton) X_GuiInput(arg0 InputEventImplementer) {
	//log.Println("Calling BaseButton.X_GuiInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "_gui_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Called when the button is pressed.
	Args: [], Returns: void
*/
func (o *BaseButton) X_Pressed() {
	//log.Println("Calling BaseButton.X_Pressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "_pressed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Called when the button is toggled (only if [member toggle_mode] is active).
	Args: [{ false button_pressed bool}], Returns: void
*/
func (o *BaseButton) X_Toggled(buttonPressed gdnative.Bool) {
	//log.Println("Calling BaseButton.X_Toggled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(buttonPressed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "_toggled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *BaseButton) X_UnhandledInput(arg0 InputEventImplementer) {
	//log.Println("Calling BaseButton.X_UnhandledInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "_unhandled_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: enum.BaseButton::ActionMode
*/
func (o *BaseButton) GetActionMode() BaseButtonActionMode {
	//log.Println("Calling BaseButton.GetActionMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "get_action_mode")

	// Call the parent method.
	// enum.BaseButton::ActionMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return BaseButtonActionMode(ret)
}

/*
        Undocumented
	Args: [], Returns: ButtonGroup
*/
func (o *BaseButton) GetButtonGroup() ButtonGroupImplementer {
	//log.Println("Calling BaseButton.GetButtonGroup()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "get_button_group")

	// Call the parent method.
	// ButtonGroup
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newButtonGroupFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ButtonGroupImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "ButtonGroup" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ButtonGroupImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *BaseButton) GetButtonMask() gdnative.Int {
	//log.Println("Calling BaseButton.GetButtonMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "get_button_mask")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the visual state used to draw the button. This is useful mainly when implementing your own draw code by either overriding _draw() or connecting to "draw" signal. The visual state of the button is defined by the [enum DrawMode] enum.
	Args: [], Returns: enum.BaseButton::DrawMode
*/
func (o *BaseButton) GetDrawMode() BaseButtonDrawMode {
	//log.Println("Calling BaseButton.GetDrawMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "get_draw_mode")

	// Call the parent method.
	// enum.BaseButton::DrawMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return BaseButtonDrawMode(ret)
}

/*
        Undocumented
	Args: [], Returns: enum.Control::FocusMode
*/
func (o *BaseButton) GetEnabledFocusMode() ControlFocusMode {
	//log.Println("Calling BaseButton.GetEnabledFocusMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "get_enabled_focus_mode")

	// Call the parent method.
	// enum.Control::FocusMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ControlFocusMode(ret)
}

/*
        Undocumented
	Args: [], Returns: ShortCut
*/
func (o *BaseButton) GetShortcut() ShortCutImplementer {
	//log.Println("Calling BaseButton.GetShortcut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "get_shortcut")

	// Call the parent method.
	// ShortCut
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newShortCutFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ShortCutImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "ShortCut" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ShortCutImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *BaseButton) IsDisabled() gdnative.Bool {
	//log.Println("Calling BaseButton.IsDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "is_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the mouse has entered the button and has not left it yet.
	Args: [], Returns: bool
*/
func (o *BaseButton) IsHovered() gdnative.Bool {
	//log.Println("Calling BaseButton.IsHovered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "is_hovered")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *BaseButton) IsKeepPressedOutside() gdnative.Bool {
	//log.Println("Calling BaseButton.IsKeepPressedOutside()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "is_keep_pressed_outside")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *BaseButton) IsPressed() gdnative.Bool {
	//log.Println("Calling BaseButton.IsPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "is_pressed")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *BaseButton) IsShortcutInTooltipEnabled() gdnative.Bool {
	//log.Println("Calling BaseButton.IsShortcutInTooltipEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "is_shortcut_in_tooltip_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *BaseButton) IsToggleMode() gdnative.Bool {
	//log.Println("Calling BaseButton.IsToggleMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "is_toggle_mode")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *BaseButton) SetActionMode(mode gdnative.Int) {
	//log.Println("Calling BaseButton.SetActionMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "set_action_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false button_group ButtonGroup}], Returns: void
*/
func (o *BaseButton) SetButtonGroup(buttonGroup ButtonGroupImplementer) {
	//log.Println("Calling BaseButton.SetButtonGroup()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(buttonGroup.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "set_button_group")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *BaseButton) SetButtonMask(mask gdnative.Int) {
	//log.Println("Calling BaseButton.SetButtonMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "set_button_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false disabled bool}], Returns: void
*/
func (o *BaseButton) SetDisabled(disabled gdnative.Bool) {
	//log.Println("Calling BaseButton.SetDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(disabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "set_disabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *BaseButton) SetEnabledFocusMode(mode gdnative.Int) {
	//log.Println("Calling BaseButton.SetEnabledFocusMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "set_enabled_focus_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *BaseButton) SetKeepPressedOutside(enabled gdnative.Bool) {
	//log.Println("Calling BaseButton.SetKeepPressedOutside()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "set_keep_pressed_outside")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pressed bool}], Returns: void
*/
func (o *BaseButton) SetPressed(pressed gdnative.Bool) {
	//log.Println("Calling BaseButton.SetPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(pressed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "set_pressed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false shortcut ShortCut}], Returns: void
*/
func (o *BaseButton) SetShortcut(shortcut ShortCutImplementer) {
	//log.Println("Calling BaseButton.SetShortcut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(shortcut.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "set_shortcut")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *BaseButton) SetShortcutInTooltip(enabled gdnative.Bool) {
	//log.Println("Calling BaseButton.SetShortcutInTooltip()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "set_shortcut_in_tooltip")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *BaseButton) SetToggleMode(enabled gdnative.Bool) {
	//log.Println("Calling BaseButton.SetToggleMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BaseButton", "set_toggle_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// BaseButtonImplementer is an interface that implements the methods
// of the BaseButton class.
type BaseButtonImplementer interface {
	ControlImplementer
	X_Pressed()
	X_Toggled(buttonPressed gdnative.Bool)
	GetButtonGroup() ButtonGroupImplementer
	GetButtonMask() gdnative.Int
	GetShortcut() ShortCutImplementer
	IsDisabled() gdnative.Bool
	IsHovered() gdnative.Bool
	IsKeepPressedOutside() gdnative.Bool
	IsPressed() gdnative.Bool
	IsShortcutInTooltipEnabled() gdnative.Bool
	IsToggleMode() gdnative.Bool
	SetActionMode(mode gdnative.Int)
	SetButtonGroup(buttonGroup ButtonGroupImplementer)
	SetButtonMask(mask gdnative.Int)
	SetDisabled(disabled gdnative.Bool)
	SetEnabledFocusMode(mode gdnative.Int)
	SetKeepPressedOutside(enabled gdnative.Bool)
	SetPressed(pressed gdnative.Bool)
	SetShortcut(shortcut ShortCutImplementer)
	SetShortcutInTooltip(enabled gdnative.Bool)
	SetToggleMode(enabled gdnative.Bool)
}
