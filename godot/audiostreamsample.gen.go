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

// AudioStreamSampleFormat is an enum for Format values.
type AudioStreamSampleFormat int

const (
	AudioStreamSampleFormat16Bits   AudioStreamSampleFormat = 1
	AudioStreamSampleFormat8Bits    AudioStreamSampleFormat = 0
	AudioStreamSampleFormatImaAdpcm AudioStreamSampleFormat = 2
)

// AudioStreamSampleLoopMode is an enum for LoopMode values.
type AudioStreamSampleLoopMode int

const (
	AudioStreamSampleLoopBackward AudioStreamSampleLoopMode = 3
	AudioStreamSampleLoopDisabled AudioStreamSampleLoopMode = 0
	AudioStreamSampleLoopForward  AudioStreamSampleLoopMode = 1
	AudioStreamSampleLoopPingPong AudioStreamSampleLoopMode = 2
)

//func NewAudioStreamSampleFromPointer(ptr gdnative.Pointer) AudioStreamSample {
func newAudioStreamSampleFromPointer(ptr gdnative.Pointer) AudioStreamSample {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioStreamSample{}
	obj.SetBaseObject(owner)

	return obj
}

/*
AudioStreamSample stores sound samples loaded from WAV files. To play the stored sound, use an [AudioStreamPlayer] (for non-positional audio) or [AudioStreamPlayer2D]/[AudioStreamPlayer3D] (for positional audio). The sound can be looped. This class can also be used to store dynamically-generated PCM audio data.
*/
type AudioStreamSample struct {
	AudioStream
	owner gdnative.Object
}

func (o *AudioStreamSample) BaseClass() string {
	return "AudioStreamSample"
}

/*
        Undocumented
	Args: [], Returns: PoolByteArray
*/
func (o *AudioStreamSample) GetData() gdnative.PoolByteArray {
	// log.Println("Calling AudioStreamSample.GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "get_data")

	// Call the parent method.
	// PoolByteArray
	retPtr := gdnative.NewEmptyPoolByteArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolByteArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.AudioStreamSample::Format
*/
func (o *AudioStreamSample) GetFormat() AudioStreamSampleFormat {
	// log.Println("Calling AudioStreamSample.GetFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "get_format")

	// Call the parent method.
	// enum.AudioStreamSample::Format
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return AudioStreamSampleFormat(ret)
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *AudioStreamSample) GetLoopBegin() gdnative.Int {
	// log.Println("Calling AudioStreamSample.GetLoopBegin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "get_loop_begin")

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
func (o *AudioStreamSample) GetLoopEnd() gdnative.Int {
	// log.Println("Calling AudioStreamSample.GetLoopEnd()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "get_loop_end")

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
	Args: [], Returns: enum.AudioStreamSample::LoopMode
*/
func (o *AudioStreamSample) GetLoopMode() AudioStreamSampleLoopMode {
	// log.Println("Calling AudioStreamSample.GetLoopMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "get_loop_mode")

	// Call the parent method.
	// enum.AudioStreamSample::LoopMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return AudioStreamSampleLoopMode(ret)
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *AudioStreamSample) GetMixRate() gdnative.Int {
	// log.Println("Calling AudioStreamSample.GetMixRate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "get_mix_rate")

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
	Args: [], Returns: bool
*/
func (o *AudioStreamSample) IsStereo() gdnative.Bool {
	// log.Println("Calling AudioStreamSample.IsStereo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "is_stereo")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Saves the AudioStreamSample as a WAV file to [code]path[/code]. Samples with IMA ADPCM format can't be saved. [b]Note:[/b] A [code].wav[/code] extension is automatically appended to [code]path[/code] if it is missing.
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *AudioStreamSample) SaveToWav(path gdnative.String) gdnative.Error {
	// log.Println("Calling AudioStreamSample.SaveToWav()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "save_to_wav")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false data PoolByteArray}], Returns: void
*/
func (o *AudioStreamSample) SetData(data gdnative.PoolByteArray) {
	// log.Println("Calling AudioStreamSample.SetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolByteArray(data)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "set_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false format int}], Returns: void
*/
func (o *AudioStreamSample) SetFormat(format gdnative.Int) {
	// log.Println("Calling AudioStreamSample.SetFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(format)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "set_format")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false loop_begin int}], Returns: void
*/
func (o *AudioStreamSample) SetLoopBegin(loopBegin gdnative.Int) {
	// log.Println("Calling AudioStreamSample.SetLoopBegin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(loopBegin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "set_loop_begin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false loop_end int}], Returns: void
*/
func (o *AudioStreamSample) SetLoopEnd(loopEnd gdnative.Int) {
	// log.Println("Calling AudioStreamSample.SetLoopEnd()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(loopEnd)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "set_loop_end")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false loop_mode int}], Returns: void
*/
func (o *AudioStreamSample) SetLoopMode(loopMode gdnative.Int) {
	// log.Println("Calling AudioStreamSample.SetLoopMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(loopMode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "set_loop_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mix_rate int}], Returns: void
*/
func (o *AudioStreamSample) SetMixRate(mixRate gdnative.Int) {
	// log.Println("Calling AudioStreamSample.SetMixRate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mixRate)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "set_mix_rate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false stereo bool}], Returns: void
*/
func (o *AudioStreamSample) SetStereo(stereo gdnative.Bool) {
	// log.Println("Calling AudioStreamSample.SetStereo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(stereo)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamSample", "set_stereo")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioStreamSampleImplementer is an interface that implements the methods
// of the AudioStreamSample class.
type AudioStreamSampleImplementer interface {
	AudioStreamImplementer
	GetData() gdnative.PoolByteArray
	GetLoopBegin() gdnative.Int
	GetLoopEnd() gdnative.Int
	GetMixRate() gdnative.Int
	IsStereo() gdnative.Bool
	SetData(data gdnative.PoolByteArray)
	SetFormat(format gdnative.Int)
	SetLoopBegin(loopBegin gdnative.Int)
	SetLoopEnd(loopEnd gdnative.Int)
	SetLoopMode(loopMode gdnative.Int)
	SetMixRate(mixRate gdnative.Int)
	SetStereo(stereo gdnative.Bool)
}
