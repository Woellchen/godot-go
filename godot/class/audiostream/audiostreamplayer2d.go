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
Plays audio that dampens with distance from screen center.
*/
type AudioStreamPlayer2D struct {
	Node2D
}

func (o *AudioStreamPlayer2D) BaseClass() string {
	return "AudioStreamPlayer2D"
}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) X_BusLayoutChanged() {
	log.Println("Calling AudioStreamPlayer2D.X_BusLayoutChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_bus_layout_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) X_IsActive() gdnative.Bool {
	log.Println("Calling AudioStreamPlayer2D.X_IsActive()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "_is_active", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) X_SetPlaying(enable gdnative.Bool) {
	log.Println("Calling AudioStreamPlayer2D.X_SetPlaying()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_set_playing", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) GetAreaMask() gdnative.Int {
	log.Println("Calling AudioStreamPlayer2D.GetAreaMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_area_mask", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) GetAttenuation() gdnative.Float {
	log.Println("Calling AudioStreamPlayer2D.GetAttenuation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_attenuation", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) GetBus() gdnative.String {
	log.Println("Calling AudioStreamPlayer2D.GetBus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_bus", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) GetMaxDistance() gdnative.Float {
	log.Println("Calling AudioStreamPlayer2D.GetMaxDistance()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_max_distance", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) GetPitchScale() gdnative.Float {
	log.Println("Calling AudioStreamPlayer2D.GetPitchScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_pitch_scale", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the position in the [AudioStream].
*/
func (o *AudioStreamPlayer2D) GetPlaybackPosition() gdnative.Float {
	log.Println("Calling AudioStreamPlayer2D.GetPlaybackPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_playback_position", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) GetStream() *AudioStream {
	log.Println("Calling AudioStreamPlayer2D.GetStream()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_stream", goArguments, "*AudioStream")

	returnValue := goRet.Interface().(*AudioStream)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) GetVolumeDb() gdnative.Float {
	log.Println("Calling AudioStreamPlayer2D.GetVolumeDb()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_volume_db", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) IsAutoplayEnabled() gdnative.Bool {
	log.Println("Calling AudioStreamPlayer2D.IsAutoplayEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_autoplay_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) IsPlaying() gdnative.Bool {
	log.Println("Calling AudioStreamPlayer2D.IsPlaying()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_playing", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Plays the audio from the given position 'from_position', in seconds.
*/
func (o *AudioStreamPlayer2D) Play(fromPosition gdnative.Float) {
	log.Println("Calling AudioStreamPlayer2D.Play()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(fromPosition)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "play", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the position from which audio will be played, in seconds.
*/
func (o *AudioStreamPlayer2D) Seek(toPosition gdnative.Float) {
	log.Println("Calling AudioStreamPlayer2D.Seek()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(toPosition)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "seek", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) SetAreaMask(mask gdnative.Int) {
	log.Println("Calling AudioStreamPlayer2D.SetAreaMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_area_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) SetAttenuation(curve gdnative.Float) {
	log.Println("Calling AudioStreamPlayer2D.SetAttenuation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(curve)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_attenuation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) SetAutoplay(enable gdnative.Bool) {
	log.Println("Calling AudioStreamPlayer2D.SetAutoplay()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_autoplay", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) SetBus(bus gdnative.String) {
	log.Println("Calling AudioStreamPlayer2D.SetBus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bus)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_bus", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) SetMaxDistance(pixels gdnative.Float) {
	log.Println("Calling AudioStreamPlayer2D.SetMaxDistance()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(pixels)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_max_distance", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) SetPitchScale(pitchScale gdnative.Float) {
	log.Println("Calling AudioStreamPlayer2D.SetPitchScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(pitchScale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_pitch_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) SetStream(stream *AudioStream) {
	log.Println("Calling AudioStreamPlayer2D.SetStream()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(stream)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_stream", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioStreamPlayer2D) SetVolumeDb(volumeDb gdnative.Float) {
	log.Println("Calling AudioStreamPlayer2D.SetVolumeDb()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(volumeDb)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_volume_db", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Stops the audio.
*/
func (o *AudioStreamPlayer2D) Stop() {
	log.Println("Calling AudioStreamPlayer2D.Stop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "stop", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AudioStreamPlayer2DImplementer is an interface for AudioStreamPlayer2D objects.
*/
type AudioStreamPlayer2DImplementer interface {
	Class
}
