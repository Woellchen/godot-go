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

// AudioEffectDistortionMode is an enum for Mode values.
type AudioEffectDistortionMode int

const (
	AudioEffectDistortionModeAtan      AudioEffectDistortionMode = 1
	AudioEffectDistortionModeClip      AudioEffectDistortionMode = 0
	AudioEffectDistortionModeLofi      AudioEffectDistortionMode = 2
	AudioEffectDistortionModeOverdrive AudioEffectDistortionMode = 3
	AudioEffectDistortionModeWaveshape AudioEffectDistortionMode = 4
)

//func NewAudioEffectDistortionFromPointer(ptr gdnative.Pointer) AudioEffectDistortion {
func newAudioEffectDistortionFromPointer(ptr gdnative.Pointer) AudioEffectDistortion {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectDistortion{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Modify the sound and make it dirty. Different types are available: clip, tan, lo-fi (bit crushing), overdrive, or waveshape. By distorting the waveform the frequency content change, which will often make the sound "crunchy" or "abrasive". For games, it can simulate sound coming from some saturated device or speaker very efficiently.
*/
type AudioEffectDistortion struct {
	AudioEffect
	owner gdnative.Object
}

func (o *AudioEffectDistortion) BaseClass() string {
	return "AudioEffectDistortion"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectDistortion) GetDrive() gdnative.Real {
	// log.Println("Calling AudioEffectDistortion.GetDrive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "get_drive")

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
func (o *AudioEffectDistortion) GetKeepHfHz() gdnative.Real {
	// log.Println("Calling AudioEffectDistortion.GetKeepHfHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "get_keep_hf_hz")

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
	Args: [], Returns: enum.AudioEffectDistortion::Mode
*/
func (o *AudioEffectDistortion) GetMode() AudioEffectDistortionMode {
	// log.Println("Calling AudioEffectDistortion.GetMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "get_mode")

	// Call the parent method.
	// enum.AudioEffectDistortion::Mode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return AudioEffectDistortionMode(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectDistortion) GetPostGain() gdnative.Real {
	// log.Println("Calling AudioEffectDistortion.GetPostGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "get_post_gain")

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
func (o *AudioEffectDistortion) GetPreGain() gdnative.Real {
	// log.Println("Calling AudioEffectDistortion.GetPreGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "get_pre_gain")

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
	Args: [{ false drive float}], Returns: void
*/
func (o *AudioEffectDistortion) SetDrive(drive gdnative.Real) {
	// log.Println("Calling AudioEffectDistortion.SetDrive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(drive)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "set_drive")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false keep_hf_hz float}], Returns: void
*/
func (o *AudioEffectDistortion) SetKeepHfHz(keepHfHz gdnative.Real) {
	// log.Println("Calling AudioEffectDistortion.SetKeepHfHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(keepHfHz)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "set_keep_hf_hz")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *AudioEffectDistortion) SetMode(mode gdnative.Int) {
	// log.Println("Calling AudioEffectDistortion.SetMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "set_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false post_gain float}], Returns: void
*/
func (o *AudioEffectDistortion) SetPostGain(postGain gdnative.Real) {
	// log.Println("Calling AudioEffectDistortion.SetPostGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(postGain)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "set_post_gain")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pre_gain float}], Returns: void
*/
func (o *AudioEffectDistortion) SetPreGain(preGain gdnative.Real) {
	// log.Println("Calling AudioEffectDistortion.SetPreGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(preGain)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "set_pre_gain")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioEffectDistortionImplementer is an interface that implements the methods
// of the AudioEffectDistortion class.
type AudioEffectDistortionImplementer interface {
	AudioEffectImplementer
	GetDrive() gdnative.Real
	GetKeepHfHz() gdnative.Real
	GetPostGain() gdnative.Real
	GetPreGain() gdnative.Real
	SetDrive(drive gdnative.Real)
	SetKeepHfHz(keepHfHz gdnative.Real)
	SetMode(mode gdnative.Int)
	SetPostGain(postGain gdnative.Real)
	SetPreGain(preGain gdnative.Real)
}
