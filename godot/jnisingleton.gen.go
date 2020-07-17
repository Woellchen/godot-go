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

//func NewJNISingletonFromPointer(ptr gdnative.Pointer) JNISingleton {
func newJNISingletonFromPointer(ptr gdnative.Pointer) JNISingleton {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := JNISingleton{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type JNISingleton struct {
	Object
	owner gdnative.Object
}

func (o *JNISingleton) BaseClass() string {
	return "JNISingleton"
}

// JNISingletonImplementer is an interface that implements the methods
// of the JNISingleton class.
type JNISingletonImplementer interface {
	ObjectImplementer
}