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

//func NewInputEventFromPointer(ptr gdnative.Pointer) InputEvent {
func newInputEventFromPointer(ptr gdnative.Pointer) InputEvent {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputEvent{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base class of all sort of input event. See [method Node._input].
*/
type InputEvent struct {
	Resource
	owner gdnative.Object
}

func (o *InputEvent) BaseClass() string {
	return "InputEvent"
}

/*
        Returns [code]true[/code] if the given input event and this input event can be added together (only for events of type [InputEventMouseMotion]). The given input event's position, global position and speed will be copied. The resulting [code]relative[/code] is a sum of both events. Both events' modifiers have to be identical.
	Args: [{ false with_event InputEvent}], Returns: bool
*/
func (o *InputEvent) Accumulate(withEvent InputEventImplementer) gdnative.Bool {
	// log.Println("Calling InputEvent.Accumulate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(withEvent.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "accumulate")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns a [String] representation of the event.
	Args: [], Returns: String
*/
func (o *InputEvent) AsText() gdnative.String {
	// log.Println("Calling InputEvent.AsText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "as_text")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns a value between 0.0 and 1.0 depending on the given actions' state. Useful for getting the value of events of type [InputEventJoypadMotion].
	Args: [{ false action String}], Returns: float
*/
func (o *InputEvent) GetActionStrength(action gdnative.String) gdnative.Real {
	// log.Println("Calling InputEvent.GetActionStrength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "get_action_strength")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *InputEvent) GetDevice() gdnative.Int {
	// log.Println("Calling InputEvent.GetDevice()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "get_device")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if this input event matches a pre-defined action of any type.
	Args: [{ false action String}], Returns: bool
*/
func (o *InputEvent) IsAction(action gdnative.String) gdnative.Bool {
	// log.Println("Calling InputEvent.IsAction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_action")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the given action is being pressed (and is not an echo event for [InputEventKey] events, unless [code]allow_echo[/code] is [code]true[/code]). Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
	Args: [{ false action String} {False true allow_echo bool}], Returns: bool
*/
func (o *InputEvent) IsActionPressed(action gdnative.String, allowEcho gdnative.Bool) gdnative.Bool {
	// log.Println("Calling InputEvent.IsActionPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(action)
	ptrArguments[1] = gdnative.NewPointerFromBool(allowEcho)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_action_pressed")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the given action is released (i.e. not pressed). Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
	Args: [{ false action String}], Returns: bool
*/
func (o *InputEvent) IsActionReleased(action gdnative.String) gdnative.Bool {
	// log.Println("Calling InputEvent.IsActionReleased()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_action_released")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if this input event's type is one that can be assigned to an input action.
	Args: [], Returns: bool
*/
func (o *InputEvent) IsActionType() gdnative.Bool {
	// log.Println("Calling InputEvent.IsActionType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_action_type")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if this input event is an echo event (only for events of type [InputEventKey]).
	Args: [], Returns: bool
*/
func (o *InputEvent) IsEcho() gdnative.Bool {
	// log.Println("Calling InputEvent.IsEcho()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_echo")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if this input event is pressed. Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
	Args: [], Returns: bool
*/
func (o *InputEvent) IsPressed() gdnative.Bool {
	// log.Println("Calling InputEvent.IsPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_pressed")

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
	Args: [{ false device int}], Returns: void
*/
func (o *InputEvent) SetDevice(device gdnative.Int) {
	// log.Println("Calling InputEvent.SetDevice()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "set_device")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns [code]true[/code] if the given input event is checking for the same key ([InputEventKey]), button ([InputEventJoypadButton]) or action ([InputEventAction]).
	Args: [{ false event InputEvent}], Returns: bool
*/
func (o *InputEvent) ShortcutMatch(event InputEventImplementer) gdnative.Bool {
	// log.Println("Calling InputEvent.ShortcutMatch()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(event.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "shortcut_match")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns a copy of the given input event which has been offset by [code]local_ofs[/code] and transformed by [code]xform[/code]. Relevant for events of type [InputEventMouseButton], [InputEventMouseMotion], [InputEventScreenTouch], [InputEventScreenDrag], [InputEventMagnifyGesture] and [InputEventPanGesture].
	Args: [{ false xform Transform2D} {(0, 0) true local_ofs Vector2}], Returns: InputEvent
*/
func (o *InputEvent) XformedBy(xform gdnative.Transform2D, localOfs gdnative.Vector2) InputEventImplementer {
	// log.Println("Calling InputEvent.XformedBy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromTransform2D(xform)
	ptrArguments[1] = gdnative.NewPointerFromVector2(localOfs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "xformed_by")

	// Call the parent method.
	// InputEvent
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newInputEventFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(InputEventImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "InputEvent" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(InputEventImplementer)
	}

	return &ret
}

// InputEventImplementer is an interface that implements the methods
// of the InputEvent class.
type InputEventImplementer interface {
	ResourceImplementer
	Accumulate(withEvent InputEventImplementer) gdnative.Bool
	AsText() gdnative.String
	GetActionStrength(action gdnative.String) gdnative.Real
	GetDevice() gdnative.Int
	IsAction(action gdnative.String) gdnative.Bool
	IsActionPressed(action gdnative.String, allowEcho gdnative.Bool) gdnative.Bool
	IsActionReleased(action gdnative.String) gdnative.Bool
	IsActionType() gdnative.Bool
	IsEcho() gdnative.Bool
	IsPressed() gdnative.Bool
	SetDevice(device gdnative.Int)
	ShortcutMatch(event InputEventImplementer) gdnative.Bool
	XformedBy(xform gdnative.Transform2D, localOfs gdnative.Vector2) InputEventImplementer
}
