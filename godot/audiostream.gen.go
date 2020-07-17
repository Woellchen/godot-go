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

//func NewAudioStreamFromPointer(ptr gdnative.Pointer) AudioStream {
func newAudioStreamFromPointer(ptr gdnative.Pointer) AudioStream {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioStream{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base class for audio streams. Audio streams are used for sound effects and music playback, and support WAV (via [AudioStreamSample]) and OGG (via [AudioStreamOGGVorbis]) file formats.
*/
type AudioStream struct {
	Resource
	owner gdnative.Object
}

func (o *AudioStream) BaseClass() string {
	return "AudioStream"
}

/*
        Returns the length of the audio stream in seconds.
	Args: [], Returns: float
*/
func (o *AudioStream) GetLength() gdnative.Real {
	//log.Println("Calling AudioStream.GetLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStream", "get_length")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

// AudioStreamImplementer is an interface that implements the methods
// of the AudioStream class.
type AudioStreamImplementer interface {
	ResourceImplementer
	GetLength() gdnative.Real
}
