package animatedsprite3d

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
type AnimatedSprite3D struct {
	SpriteBase3D
}

func (o *AnimatedSprite3D) BaseClass() string {
	return "AnimatedSprite3D"
}

/*
   Undocumented
*/
func (o *AnimatedSprite3D) X_IsPlaying() gdnative.Bool {
	log.Println("Calling AnimatedSprite3D.X_IsPlaying()")

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
func (o *AnimatedSprite3D) X_ResChanged() {
	log.Println("Calling AnimatedSprite3D.X_ResChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_res_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AnimatedSprite3D) X_SetPlaying(playing gdnative.Bool) {
	log.Println("Calling AnimatedSprite3D.X_SetPlaying()")

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
func (o *AnimatedSprite3D) GetAnimation() gdnative.String {
	log.Println("Calling AnimatedSprite3D.GetAnimation()")

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
func (o *AnimatedSprite3D) GetFrame() gdnative.Int {
	log.Println("Calling AnimatedSprite3D.GetFrame()")

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
func (o *AnimatedSprite3D) GetSpriteFrames() *SpriteFrames {
	log.Println("Calling AnimatedSprite3D.GetSpriteFrames()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_sprite_frames", goArguments, "*SpriteFrames")

	returnValue := goRet.Interface().(*SpriteFrames)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return true if an animation if currently being played.
*/
func (o *AnimatedSprite3D) IsPlaying() gdnative.Bool {
	log.Println("Calling AnimatedSprite3D.IsPlaying()")

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
func (o *AnimatedSprite3D) Play(anim gdnative.String) {
	log.Println("Calling AnimatedSprite3D.Play()")

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
func (o *AnimatedSprite3D) SetAnimation(animation gdnative.String) {
	log.Println("Calling AnimatedSprite3D.SetAnimation()")

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
func (o *AnimatedSprite3D) SetFrame(frame gdnative.Int) {
	log.Println("Calling AnimatedSprite3D.SetFrame()")

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
func (o *AnimatedSprite3D) SetSpriteFrames(spriteFrames *SpriteFrames) {
	log.Println("Calling AnimatedSprite3D.SetSpriteFrames()")

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
func (o *AnimatedSprite3D) Stop() {
	log.Println("Calling AnimatedSprite3D.Stop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "stop", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AnimatedSprite3DImplementer is an interface for AnimatedSprite3D objects.
*/
type AnimatedSprite3DImplementer interface {
	Class
}
