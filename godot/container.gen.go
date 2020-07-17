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

//func NewContainerFromPointer(ptr gdnative.Pointer) Container {
func newContainerFromPointer(ptr gdnative.Pointer) Container {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Container{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base node for containers. A [Container] contains other controls and automatically arranges them in a certain way. A Control can inherit this to create custom container classes.
*/
type Container struct {
	Control
	owner gdnative.Object
}

func (o *Container) BaseClass() string {
	return "Container"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Container) X_ChildMinsizeChanged() {
	//log.Println("Calling Container.X_ChildMinsizeChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Container", "_child_minsize_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Container) X_SortChildren() {
	//log.Println("Calling Container.X_SortChildren()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Container", "_sort_children")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Fit a child control in a given rect. This is mainly a helper for creating custom container classes.
	Args: [{ false child Control} { false rect Rect2}], Returns: void
*/
func (o *Container) FitChildInRect(child ControlImplementer, rect gdnative.Rect2) {
	//log.Println("Calling Container.FitChildInRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(child.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromRect2(rect)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Container", "fit_child_in_rect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Queue resort of the contained children. This is called automatically anyway, but can be called upon request.
	Args: [], Returns: void
*/
func (o *Container) QueueSort() {
	//log.Println("Calling Container.QueueSort()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Container", "queue_sort")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ContainerImplementer is an interface that implements the methods
// of the Container class.
type ContainerImplementer interface {
	ControlImplementer
	X_ChildMinsizeChanged()
	X_SortChildren()
	FitChildInRect(child ControlImplementer, rect gdnative.Rect2)
	QueueSort()
}
