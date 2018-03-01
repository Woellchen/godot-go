package world2d

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
Class that has everything pertaining to a 2D world. A physics space, a visual scenario and a sound space. 2D nodes register their resources into the current 2D world.
*/
type World2D struct {
	Resource
}

func (o *World2D) BaseClass() string {
	return "World2D"
}

/*
   Undocumented
*/
func (o *World2D) GetCanvas() *RID {
	log.Println("Calling World2D.GetCanvas()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_canvas", goArguments, "*RID")

	returnValue := goRet.Interface().(*RID)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *World2D) GetDirectSpaceState() *Physics2DDirectSpaceState {
	log.Println("Calling World2D.GetDirectSpaceState()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_direct_space_state", goArguments, "*Physics2DDirectSpaceState")

	returnValue := goRet.Interface().(*Physics2DDirectSpaceState)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *World2D) GetSpace() *RID {
	log.Println("Calling World2D.GetSpace()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_space", goArguments, "*RID")

	returnValue := goRet.Interface().(*RID)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   World2DImplementer is an interface for World2D objects.
*/
type World2DImplementer interface {
	Class
}
