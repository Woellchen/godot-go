package audiostream

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
Randomly varies pitch on each start.
*/
type AudioStreamRandomPitch struct {
	AudioStream
}

func (o *AudioStreamRandomPitch) BaseClass() string {
	return "AudioStreamRandomPitch"
}

/*
   Undocumented
*/
func (o *AudioStreamRandomPitch) GetAudioStream() *AudioStream {
	log.Println("Calling AudioStreamRandomPitch.GetAudioStream()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_audio_stream", goArguments, "*AudioStream")

	returnValue := goRet.Interface().(*AudioStream)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioStreamRandomPitch) GetRandomPitch() gdnative.Float {
	log.Println("Calling AudioStreamRandomPitch.GetRandomPitch()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_random_pitch", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioStreamRandomPitch) SetAudioStream(stream *AudioStream) {
	log.Println("Calling AudioStreamRandomPitch.SetAudioStream()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(stream)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_audio_stream", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioStreamRandomPitch) SetRandomPitch(scale gdnative.Float) {
	log.Println("Calling AudioStreamRandomPitch.SetRandomPitch()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(scale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_random_pitch", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AudioStreamRandomPitchImplementer is an interface for AudioStreamRandomPitch objects.
*/
type AudioStreamRandomPitchImplementer interface {
	Class
}
