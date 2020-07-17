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

// BoxContainerAlignMode is an enum for AlignMode values.
type BoxContainerAlignMode int

const (
	BoxContainerAlignBegin  BoxContainerAlignMode = 0
	BoxContainerAlignCenter BoxContainerAlignMode = 1
	BoxContainerAlignEnd    BoxContainerAlignMode = 2
)

//func NewBoxContainerFromPointer(ptr gdnative.Pointer) BoxContainer {
func newBoxContainerFromPointer(ptr gdnative.Pointer) BoxContainer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := BoxContainer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Arranges child controls vertically or horizontally, and rearranges the controls automatically when their minimum size changes.
*/
type BoxContainer struct {
	Container
	owner gdnative.Object
}

func (o *BoxContainer) BaseClass() string {
	return "BoxContainer"
}

/*
        Adds a control to the box as a spacer. If [code]true[/code], [code]begin[/code] will insert the spacer control in front of other children.
	Args: [{ false begin bool}], Returns: void
*/
func (o *BoxContainer) AddSpacer(begin gdnative.Bool) {
	//log.Println("Calling BoxContainer.AddSpacer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(begin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BoxContainer", "add_spacer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: enum.BoxContainer::AlignMode
*/
func (o *BoxContainer) GetAlignment() BoxContainerAlignMode {
	//log.Println("Calling BoxContainer.GetAlignment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BoxContainer", "get_alignment")

	// Call the parent method.
	// enum.BoxContainer::AlignMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return BoxContainerAlignMode(ret)
}

/*
        Undocumented
	Args: [{ false alignment int}], Returns: void
*/
func (o *BoxContainer) SetAlignment(alignment gdnative.Int) {
	//log.Println("Calling BoxContainer.SetAlignment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(alignment)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BoxContainer", "set_alignment")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// BoxContainerImplementer is an interface that implements the methods
// of the BoxContainer class.
type BoxContainerImplementer interface {
	ContainerImplementer
	AddSpacer(begin gdnative.Bool)
	SetAlignment(alignment gdnative.Int)
}
