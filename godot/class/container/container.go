package container

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
Base node for containers. A [code]Container[/code] contains other controls and automatically arranges them in a certain way. A Control can inherit this to create custom container classes.
*/
type Container struct {
	Control
}

func (o *Container) BaseClass() string {
	return "Container"
}

/*
   Undocumented
*/
func (o *Container) X_ChildMinsizeChanged() {
	log.Println("Calling Container.X_ChildMinsizeChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_child_minsize_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Container) X_SortChildren() {
	log.Println("Calling Container.X_SortChildren()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_sort_children", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Fit a child control in a given rect. This is mainly a helper for creating custom container classes.
*/
func (o *Container) FitChildInRect(child *Object, rect *Rect2) {
	log.Println("Calling Container.FitChildInRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(child)
	goArguments[1] = reflect.ValueOf(rect)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "fit_child_in_rect", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Queue resort of the contained children. This is called automatically anyway, but can be called upon request.
*/
func (o *Container) QueueSort() {
	log.Println("Calling Container.QueueSort()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "queue_sort", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ContainerImplementer is an interface for Container objects.
*/
type ContainerImplementer interface {
	Class
}
