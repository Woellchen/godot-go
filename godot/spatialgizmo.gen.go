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

//func NewSpatialGizmoFromPointer(ptr gdnative.Pointer) SpatialGizmo {
func newSpatialGizmoFromPointer(ptr gdnative.Pointer) SpatialGizmo {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SpatialGizmo{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type SpatialGizmo struct {
	Reference
	owner gdnative.Object
}

func (o *SpatialGizmo) BaseClass() string {
	return "SpatialGizmo"
}

// SpatialGizmoImplementer is an interface that implements the methods
// of the SpatialGizmo class.
type SpatialGizmoImplementer interface {
	ReferenceImplementer
}
