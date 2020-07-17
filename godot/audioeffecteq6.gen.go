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

//func NewAudioEffectEQ6FromPointer(ptr gdnative.Pointer) AudioEffectEQ6 {
func newAudioEffectEQ6FromPointer(ptr gdnative.Pointer) AudioEffectEQ6 {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectEQ6{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Frequency bands: Band 1: 32 Hz Band 2: 100 Hz Band 3: 320 Hz Band 4: 1000 Hz Band 5: 3200 Hz Band 6: 10000 Hz See also [AudioEffectEQ], [AudioEffectEQ10], [AudioEffectEQ21].
*/
type AudioEffectEQ6 struct {
	AudioEffectEQ
	owner gdnative.Object
}

func (o *AudioEffectEQ6) BaseClass() string {
	return "AudioEffectEQ6"
}

// AudioEffectEQ6Implementer is an interface that implements the methods
// of the AudioEffectEQ6 class.
type AudioEffectEQ6Implementer interface {
	AudioEffectEQImplementer
}
