package audioeffect

import (
	"log"
	"reflect"
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
Allows frequencies other than the [member cutoff_hz] to pass.
*/
type AudioEffectFilter struct {
	AudioEffect
}

func (o *AudioEffectFilter) BaseClass() string {
	return "AudioEffectFilter"
}

/*
   Undocumented
*/
func (o *AudioEffectFilter) GetCutoff() gdnative.Float {
	log.Println("Calling AudioEffectFilter.GetCutoff()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_cutoff", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioEffectFilter) GetDb() gdnative.Int {
	log.Println("Calling AudioEffectFilter.GetDb()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_db", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioEffectFilter) GetGain() gdnative.Float {
	log.Println("Calling AudioEffectFilter.GetGain()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_gain", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioEffectFilter) GetResonance() gdnative.Float {
	log.Println("Calling AudioEffectFilter.GetResonance()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_resonance", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AudioEffectFilter) SetCutoff(freq gdnative.Float) {
	log.Println("Calling AudioEffectFilter.SetCutoff()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(freq)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_cutoff", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioEffectFilter) SetDb(amount gdnative.Int) {
	log.Println("Calling AudioEffectFilter.SetDb()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_db", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioEffectFilter) SetGain(amount gdnative.Float) {
	log.Println("Calling AudioEffectFilter.SetGain()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_gain", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AudioEffectFilter) SetResonance(amount gdnative.Float) {
	log.Println("Calling AudioEffectFilter.SetResonance()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_resonance", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AudioEffectFilterImplementer is an interface for AudioEffectFilter objects.
*/
type AudioEffectFilterImplementer interface {
	Class
}
