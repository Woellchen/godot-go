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

//func NewAnimationNodeTimeSeekFromPointer(ptr gdnative.Pointer) AnimationNodeTimeSeek {
func newAnimationNodeTimeSeekFromPointer(ptr gdnative.Pointer) AnimationNodeTimeSeek {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNodeTimeSeek{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This node can be used to cause a seek command to happen to any sub-children of the graph. After setting the time, this value returns to -1.
*/
type AnimationNodeTimeSeek struct {
	AnimationNode
	owner gdnative.Object
}

func (o *AnimationNodeTimeSeek) BaseClass() string {
	return "AnimationNodeTimeSeek"
}

// AnimationNodeTimeSeekImplementer is an interface that implements the methods
// of the AnimationNodeTimeSeek class.
type AnimationNodeTimeSeekImplementer interface {
	AnimationNodeImplementer
}
