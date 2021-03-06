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

//func NewNavigationPolygonInstanceFromPointer(ptr gdnative.Pointer) NavigationPolygonInstance {
func newNavigationPolygonInstanceFromPointer(ptr gdnative.Pointer) NavigationPolygonInstance {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := NavigationPolygonInstance{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type NavigationPolygonInstance struct {
	Node2D
	owner gdnative.Object
}

func (o *NavigationPolygonInstance) BaseClass() string {
	return "NavigationPolygonInstance"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *NavigationPolygonInstance) X_NavpolyChanged() {
	// log.Println("Calling NavigationPolygonInstance.X_NavpolyChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NavigationPolygonInstance", "_navpoly_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: NavigationPolygon
*/
func (o *NavigationPolygonInstance) GetNavigationPolygon() NavigationPolygonImplementer {
	// log.Println("Calling NavigationPolygonInstance.GetNavigationPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NavigationPolygonInstance", "get_navigation_polygon")

	// Call the parent method.
	// NavigationPolygon
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newNavigationPolygonFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(NavigationPolygonImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "NavigationPolygon" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(NavigationPolygonImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *NavigationPolygonInstance) IsEnabled() gdnative.Bool {
	// log.Println("Calling NavigationPolygonInstance.IsEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NavigationPolygonInstance", "is_enabled")

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
	Args: [{ false enabled bool}], Returns: void
*/
func (o *NavigationPolygonInstance) SetEnabled(enabled gdnative.Bool) {
	// log.Println("Calling NavigationPolygonInstance.SetEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NavigationPolygonInstance", "set_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false navpoly NavigationPolygon}], Returns: void
*/
func (o *NavigationPolygonInstance) SetNavigationPolygon(navpoly NavigationPolygonImplementer) {
	// log.Println("Calling NavigationPolygonInstance.SetNavigationPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(navpoly.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NavigationPolygonInstance", "set_navigation_polygon")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// NavigationPolygonInstanceImplementer is an interface that implements the methods
// of the NavigationPolygonInstance class.
type NavigationPolygonInstanceImplementer interface {
	Node2DImplementer
	X_NavpolyChanged()
	GetNavigationPolygon() NavigationPolygonImplementer
	IsEnabled() gdnative.Bool
	SetEnabled(enabled gdnative.Bool)
	SetNavigationPolygon(navpoly NavigationPolygonImplementer)
}
