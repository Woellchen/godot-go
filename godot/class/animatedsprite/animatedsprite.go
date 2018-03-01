package animatedsprite

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
Animations are created using a [SpriteFrames] resource, which can be configured in the editor via the SpriteFrames panel.
*/
type AnimatedSprite struct {
	Node2D
}

func (o *AnimatedSprite) BaseClass() string {
	return "AnimatedSprite"
}

/*
   Undocumented
*/
func (o *AnimatedSprite) X_IsPlaying() gdnative.Bool {
	log.Println("Calling AnimatedSprite.X_IsPlaying()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "_is_playing", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AnimatedSprite) X_ResChanged() {
	log.Println("Calling AnimatedSprite.X_ResChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_res_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AnimatedSprite) X_SetPlaying(playing gdnative.Bool) {
	log.Println("Calling AnimatedSprite.X_SetPlaying()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(playing)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_set_playing", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AnimatedSprite) GetAnimation() gdnative.String {
	log.Println("Calling AnimatedSprite.GetAnimation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_animation", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AnimatedSprite) GetFrame() gdnative.Int {
	log.Println("Calling AnimatedSprite.GetFrame()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_frame", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AnimatedSprite) GetOffset() *Vector2 {
	log.Println("Calling AnimatedSprite.GetOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_offset", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AnimatedSprite) GetSpriteFrames() *SpriteFrames {
	log.Println("Calling AnimatedSprite.GetSpriteFrames()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_sprite_frames", goArguments, "*SpriteFrames")

	returnValue := goRet.Interface().(*SpriteFrames)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AnimatedSprite) IsCentered() gdnative.Bool {
	log.Println("Calling AnimatedSprite.IsCentered()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_centered", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AnimatedSprite) IsFlippedH() gdnative.Bool {
	log.Println("Calling AnimatedSprite.IsFlippedH()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_flipped_h", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *AnimatedSprite) IsFlippedV() gdnative.Bool {
	log.Println("Calling AnimatedSprite.IsFlippedV()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_flipped_v", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return true if an animation if currently being played.
*/
func (o *AnimatedSprite) IsPlaying() gdnative.Bool {
	log.Println("Calling AnimatedSprite.IsPlaying()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_playing", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Play the animation set in parameter. If no parameter is provided, the current animation is played.
*/
func (o *AnimatedSprite) Play(anim gdnative.String) {
	log.Println("Calling AnimatedSprite.Play()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(anim)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "play", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AnimatedSprite) SetAnimation(animation gdnative.String) {
	log.Println("Calling AnimatedSprite.SetAnimation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(animation)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_animation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AnimatedSprite) SetCentered(centered gdnative.Bool) {
	log.Println("Calling AnimatedSprite.SetCentered()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(centered)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_centered", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AnimatedSprite) SetFlipH(flipH gdnative.Bool) {
	log.Println("Calling AnimatedSprite.SetFlipH()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(flipH)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_flip_h", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AnimatedSprite) SetFlipV(flipV gdnative.Bool) {
	log.Println("Calling AnimatedSprite.SetFlipV()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(flipV)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_flip_v", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AnimatedSprite) SetFrame(frame gdnative.Int) {
	log.Println("Calling AnimatedSprite.SetFrame()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(frame)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_frame", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AnimatedSprite) SetOffset(offset *Vector2) {
	log.Println("Calling AnimatedSprite.SetOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(offset)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_offset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AnimatedSprite) SetSpriteFrames(spriteFrames *SpriteFrames) {
	log.Println("Calling AnimatedSprite.SetSpriteFrames()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(spriteFrames)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_sprite_frames", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Stop the current animation (does not reset the frame counter).
*/
func (o *AnimatedSprite) Stop() {
	log.Println("Calling AnimatedSprite.Stop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "stop", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AnimatedSpriteImplementer is an interface for AnimatedSprite objects.
*/
type AnimatedSpriteImplementer interface {
	Class
}
