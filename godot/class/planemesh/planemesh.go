package planemesh

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
Class representing a planar [PrimitiveMesh]. This flat mesh does not have a thickness.
*/
type PlaneMesh struct {
	PrimitiveMesh
}

func (o *PlaneMesh) BaseClass() string {
	return "PlaneMesh"
}

/*
   Undocumented
*/
func (o *PlaneMesh) GetSize() *Vector2 {
	log.Println("Calling PlaneMesh.GetSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_size", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PlaneMesh) GetSubdivideDepth() gdnative.Int {
	log.Println("Calling PlaneMesh.GetSubdivideDepth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_subdivide_depth", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PlaneMesh) GetSubdivideWidth() gdnative.Int {
	log.Println("Calling PlaneMesh.GetSubdivideWidth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_subdivide_width", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PlaneMesh) SetSize(size *Vector2) {
	log.Println("Calling PlaneMesh.SetSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PlaneMesh) SetSubdivideDepth(subdivide gdnative.Int) {
	log.Println("Calling PlaneMesh.SetSubdivideDepth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(subdivide)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_subdivide_depth", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PlaneMesh) SetSubdivideWidth(subdivide gdnative.Int) {
	log.Println("Calling PlaneMesh.SetSubdivideWidth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(subdivide)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_subdivide_width", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   PlaneMeshImplementer is an interface for PlaneMesh objects.
*/
type PlaneMeshImplementer interface {
	Class
}
