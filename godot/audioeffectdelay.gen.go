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

//func NewAudioEffectDelayFromPointer(ptr gdnative.Pointer) AudioEffectDelay {
func newAudioEffectDelayFromPointer(ptr gdnative.Pointer) AudioEffectDelay {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectDelay{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Plays input signal back after a period of time. The delayed signal may be played back multiple times to create the sound of a repeating, decaying echo. Delay effects range from a subtle echo effect to a pronounced blending of previous sounds with new sounds.
*/
type AudioEffectDelay struct {
	AudioEffect
	owner gdnative.Object
}

func (o *AudioEffectDelay) BaseClass() string {
	return "AudioEffectDelay"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectDelay) GetDry() gdnative.Real {
	// log.Println("Calling AudioEffectDelay.GetDry()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "get_dry")

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
	Args: [], Returns: float
*/
func (o *AudioEffectDelay) GetFeedbackDelayMs() gdnative.Real {
	// log.Println("Calling AudioEffectDelay.GetFeedbackDelayMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "get_feedback_delay_ms")

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
	Args: [], Returns: float
*/
func (o *AudioEffectDelay) GetFeedbackLevelDb() gdnative.Real {
	// log.Println("Calling AudioEffectDelay.GetFeedbackLevelDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "get_feedback_level_db")

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
	Args: [], Returns: float
*/
func (o *AudioEffectDelay) GetFeedbackLowpass() gdnative.Real {
	// log.Println("Calling AudioEffectDelay.GetFeedbackLowpass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "get_feedback_lowpass")

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
	Args: [], Returns: float
*/
func (o *AudioEffectDelay) GetTap1DelayMs() gdnative.Real {
	// log.Println("Calling AudioEffectDelay.GetTap1DelayMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "get_tap1_delay_ms")

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
	Args: [], Returns: float
*/
func (o *AudioEffectDelay) GetTap1LevelDb() gdnative.Real {
	// log.Println("Calling AudioEffectDelay.GetTap1LevelDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "get_tap1_level_db")

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
	Args: [], Returns: float
*/
func (o *AudioEffectDelay) GetTap1Pan() gdnative.Real {
	// log.Println("Calling AudioEffectDelay.GetTap1Pan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "get_tap1_pan")

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
	Args: [], Returns: float
*/
func (o *AudioEffectDelay) GetTap2DelayMs() gdnative.Real {
	// log.Println("Calling AudioEffectDelay.GetTap2DelayMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "get_tap2_delay_ms")

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
	Args: [], Returns: float
*/
func (o *AudioEffectDelay) GetTap2LevelDb() gdnative.Real {
	// log.Println("Calling AudioEffectDelay.GetTap2LevelDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "get_tap2_level_db")

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
	Args: [], Returns: float
*/
func (o *AudioEffectDelay) GetTap2Pan() gdnative.Real {
	// log.Println("Calling AudioEffectDelay.GetTap2Pan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "get_tap2_pan")

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
	Args: [], Returns: bool
*/
func (o *AudioEffectDelay) IsFeedbackActive() gdnative.Bool {
	// log.Println("Calling AudioEffectDelay.IsFeedbackActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "is_feedback_active")

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
func (o *AudioEffectDelay) IsTap1Active() gdnative.Bool {
	// log.Println("Calling AudioEffectDelay.IsTap1Active()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "is_tap1_active")

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
func (o *AudioEffectDelay) IsTap2Active() gdnative.Bool {
	// log.Println("Calling AudioEffectDelay.IsTap2Active()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "is_tap2_active")

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
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectDelay) SetDry(amount gdnative.Real) {
	// log.Println("Calling AudioEffectDelay.SetDry()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_dry")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount bool}], Returns: void
*/
func (o *AudioEffectDelay) SetFeedbackActive(amount gdnative.Bool) {
	// log.Println("Calling AudioEffectDelay.SetFeedbackActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_feedback_active")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectDelay) SetFeedbackDelayMs(amount gdnative.Real) {
	// log.Println("Calling AudioEffectDelay.SetFeedbackDelayMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_feedback_delay_ms")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectDelay) SetFeedbackLevelDb(amount gdnative.Real) {
	// log.Println("Calling AudioEffectDelay.SetFeedbackLevelDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_feedback_level_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectDelay) SetFeedbackLowpass(amount gdnative.Real) {
	// log.Println("Calling AudioEffectDelay.SetFeedbackLowpass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_feedback_lowpass")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount bool}], Returns: void
*/
func (o *AudioEffectDelay) SetTap1Active(amount gdnative.Bool) {
	// log.Println("Calling AudioEffectDelay.SetTap1Active()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_tap1_active")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectDelay) SetTap1DelayMs(amount gdnative.Real) {
	// log.Println("Calling AudioEffectDelay.SetTap1DelayMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_tap1_delay_ms")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectDelay) SetTap1LevelDb(amount gdnative.Real) {
	// log.Println("Calling AudioEffectDelay.SetTap1LevelDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_tap1_level_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectDelay) SetTap1Pan(amount gdnative.Real) {
	// log.Println("Calling AudioEffectDelay.SetTap1Pan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_tap1_pan")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount bool}], Returns: void
*/
func (o *AudioEffectDelay) SetTap2Active(amount gdnative.Bool) {
	// log.Println("Calling AudioEffectDelay.SetTap2Active()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_tap2_active")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectDelay) SetTap2DelayMs(amount gdnative.Real) {
	// log.Println("Calling AudioEffectDelay.SetTap2DelayMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_tap2_delay_ms")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectDelay) SetTap2LevelDb(amount gdnative.Real) {
	// log.Println("Calling AudioEffectDelay.SetTap2LevelDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_tap2_level_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectDelay) SetTap2Pan(amount gdnative.Real) {
	// log.Println("Calling AudioEffectDelay.SetTap2Pan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDelay", "set_tap2_pan")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioEffectDelayImplementer is an interface that implements the methods
// of the AudioEffectDelay class.
type AudioEffectDelayImplementer interface {
	AudioEffectImplementer
	GetDry() gdnative.Real
	GetFeedbackDelayMs() gdnative.Real
	GetFeedbackLevelDb() gdnative.Real
	GetFeedbackLowpass() gdnative.Real
	GetTap1DelayMs() gdnative.Real
	GetTap1LevelDb() gdnative.Real
	GetTap1Pan() gdnative.Real
	GetTap2DelayMs() gdnative.Real
	GetTap2LevelDb() gdnative.Real
	GetTap2Pan() gdnative.Real
	IsFeedbackActive() gdnative.Bool
	IsTap1Active() gdnative.Bool
	IsTap2Active() gdnative.Bool
	SetDry(amount gdnative.Real)
	SetFeedbackActive(amount gdnative.Bool)
	SetFeedbackDelayMs(amount gdnative.Real)
	SetFeedbackLevelDb(amount gdnative.Real)
	SetFeedbackLowpass(amount gdnative.Real)
	SetTap1Active(amount gdnative.Bool)
	SetTap1DelayMs(amount gdnative.Real)
	SetTap1LevelDb(amount gdnative.Real)
	SetTap1Pan(amount gdnative.Real)
	SetTap2Active(amount gdnative.Bool)
	SetTap2DelayMs(amount gdnative.Real)
	SetTap2LevelDb(amount gdnative.Real)
	SetTap2Pan(amount gdnative.Real)
}
