package ysort

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
Sort all child nodes based on their Y positions. The child node must inherit from [CanvasItem] for it to be sorted. Nodes that have a higher Y position will be drawn later, so they will appear on top of nodes that have a lower Y position.
*/
type YSort struct {
	Node2D
}

func (o *YSort) BaseClass() string {
	return "YSort"
}

/*
   Undocumented
*/
func (o *YSort) IsSortEnabled() gdnative.Bool {
	log.Println("Calling YSort.IsSortEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_sort_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *YSort) SetSortEnabled(enabled gdnative.Bool) {
	log.Println("Calling YSort.SetSortEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_sort_enabled", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   YSortImplementer is an interface for YSort objects.
*/
type YSortImplementer interface {
	Class
}
