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

//func NewScrollContainerFromPointer(ptr gdnative.Pointer) ScrollContainer {
func newScrollContainerFromPointer(ptr gdnative.Pointer) ScrollContainer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ScrollContainer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A ScrollContainer node meant to contain a [Control] child. ScrollContainers will automatically create a scrollbar child ([HScrollBar], [VScrollBar], or both) when needed and will only draw the Control within the ScrollContainer area. Scrollbars will automatically be drawn at the right (for vertical) or bottom (for horizontal) and will enable dragging to move the viewable Control (and its children) within the ScrollContainer. Scrollbars will also automatically resize the grabber based on the [member Control.rect_min_size] of the Control relative to the ScrollContainer. Works great with a [Panel] control. You can set [code]EXPAND[/code] on the children's size flags, so they will upscale to the ScrollContainer's size if it's larger (scroll is invisible for the chosen dimension).
*/
type ScrollContainer struct {
	Container
	owner gdnative.Object
}

func (o *ScrollContainer) BaseClass() string {
	return "ScrollContainer"
}

/*
        Undocumented
	Args: [{ false arg0 Control}], Returns: void
*/
func (o *ScrollContainer) X_EnsureFocusedVisible(arg0 ControlImplementer) {
	//log.Println("Calling ScrollContainer.X_EnsureFocusedVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "_ensure_focused_visible")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *ScrollContainer) X_GuiInput(arg0 InputEventImplementer) {
	//log.Println("Calling ScrollContainer.X_GuiInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "_gui_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 float}], Returns: void
*/
func (o *ScrollContainer) X_ScrollMoved(arg0 gdnative.Real) {
	//log.Println("Calling ScrollContainer.X_ScrollMoved()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "_scroll_moved")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ScrollContainer) X_UpdateScrollbarPosition() {
	//log.Println("Calling ScrollContainer.X_UpdateScrollbarPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "_update_scrollbar_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *ScrollContainer) GetDeadzone() gdnative.Int {
	//log.Println("Calling ScrollContainer.GetDeadzone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "get_deadzone")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *ScrollContainer) GetHScroll() gdnative.Int {
	//log.Println("Calling ScrollContainer.GetHScroll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "get_h_scroll")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the horizontal scrollbar [HScrollBar] of this [ScrollContainer].
	Args: [], Returns: HScrollBar
*/
func (o *ScrollContainer) GetHScrollbar() HScrollBarImplementer {
	//log.Println("Calling ScrollContainer.GetHScrollbar()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "get_h_scrollbar")

	// Call the parent method.
	// HScrollBar
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newHScrollBarFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(HScrollBarImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "HScrollBar" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(HScrollBarImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *ScrollContainer) GetVScroll() gdnative.Int {
	//log.Println("Calling ScrollContainer.GetVScroll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "get_v_scroll")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the vertical scrollbar [VScrollBar] of this [ScrollContainer].
	Args: [], Returns: VScrollBar
*/
func (o *ScrollContainer) GetVScrollbar() VScrollBarImplementer {
	//log.Println("Calling ScrollContainer.GetVScrollbar()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "get_v_scrollbar")

	// Call the parent method.
	// VScrollBar
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newVScrollBarFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(VScrollBarImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "VScrollBar" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(VScrollBarImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *ScrollContainer) IsFollowingFocus() gdnative.Bool {
	//log.Println("Calling ScrollContainer.IsFollowingFocus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "is_following_focus")

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
func (o *ScrollContainer) IsHScrollEnabled() gdnative.Bool {
	//log.Println("Calling ScrollContainer.IsHScrollEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "is_h_scroll_enabled")

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
func (o *ScrollContainer) IsVScrollEnabled() gdnative.Bool {
	//log.Println("Calling ScrollContainer.IsVScrollEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "is_v_scroll_enabled")

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
	Args: [{ false deadzone int}], Returns: void
*/
func (o *ScrollContainer) SetDeadzone(deadzone gdnative.Int) {
	//log.Println("Calling ScrollContainer.SetDeadzone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(deadzone)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "set_deadzone")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *ScrollContainer) SetEnableHScroll(enable gdnative.Bool) {
	//log.Println("Calling ScrollContainer.SetEnableHScroll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "set_enable_h_scroll")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *ScrollContainer) SetEnableVScroll(enable gdnative.Bool) {
	//log.Println("Calling ScrollContainer.SetEnableVScroll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "set_enable_v_scroll")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *ScrollContainer) SetFollowFocus(enabled gdnative.Bool) {
	//log.Println("Calling ScrollContainer.SetFollowFocus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "set_follow_focus")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false value int}], Returns: void
*/
func (o *ScrollContainer) SetHScroll(value gdnative.Int) {
	//log.Println("Calling ScrollContainer.SetHScroll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "set_h_scroll")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false value int}], Returns: void
*/
func (o *ScrollContainer) SetVScroll(value gdnative.Int) {
	//log.Println("Calling ScrollContainer.SetVScroll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollContainer", "set_v_scroll")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ScrollContainerImplementer is an interface that implements the methods
// of the ScrollContainer class.
type ScrollContainerImplementer interface {
	ContainerImplementer
	X_EnsureFocusedVisible(arg0 ControlImplementer)
	X_ScrollMoved(arg0 gdnative.Real)
	X_UpdateScrollbarPosition()
	GetDeadzone() gdnative.Int
	GetHScroll() gdnative.Int
	GetHScrollbar() HScrollBarImplementer
	GetVScroll() gdnative.Int
	GetVScrollbar() VScrollBarImplementer
	IsFollowingFocus() gdnative.Bool
	IsHScrollEnabled() gdnative.Bool
	IsVScrollEnabled() gdnative.Bool
	SetDeadzone(deadzone gdnative.Int)
	SetEnableHScroll(enable gdnative.Bool)
	SetEnableVScroll(enable gdnative.Bool)
	SetFollowFocus(enabled gdnative.Bool)
	SetHScroll(value gdnative.Int)
	SetVScroll(value gdnative.Int)
}
