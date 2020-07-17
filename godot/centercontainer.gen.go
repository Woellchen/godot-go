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

//func NewCenterContainerFromPointer(ptr gdnative.Pointer) CenterContainer {
func newCenterContainerFromPointer(ptr gdnative.Pointer) CenterContainer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CenterContainer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
CenterContainer keeps children controls centered. This container keeps all children to their minimum size, in the center.
*/
type CenterContainer struct {
	Container
	owner gdnative.Object
}

func (o *CenterContainer) BaseClass() string {
	return "CenterContainer"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *CenterContainer) IsUsingTopLeft() gdnative.Bool {
	//log.Println("Calling CenterContainer.IsUsingTopLeft()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CenterContainer", "is_using_top_left")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *CenterContainer) SetUseTopLeft(enable gdnative.Bool) {
	//log.Println("Calling CenterContainer.SetUseTopLeft()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CenterContainer", "set_use_top_left")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// CenterContainerImplementer is an interface that implements the methods
// of the CenterContainer class.
type CenterContainerImplementer interface {
	ContainerImplementer
	IsUsingTopLeft() gdnative.Bool
	SetUseTopLeft(enable gdnative.Bool)
}
