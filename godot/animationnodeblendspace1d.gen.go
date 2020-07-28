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

//func NewAnimationNodeBlendSpace1DFromPointer(ptr gdnative.Pointer) AnimationNodeBlendSpace1D {
func newAnimationNodeBlendSpace1DFromPointer(ptr gdnative.Pointer) AnimationNodeBlendSpace1D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNodeBlendSpace1D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A resource to add to an [AnimationNodeBlendTree]. This is a virtual axis on which you can add any type of [AnimationNode] using [method add_blend_point]. Outputs the linear blend of the two [AnimationNode]s closest to the node's current value. You can set the extents of the axis using the [member min_space] and [member max_space].
*/
type AnimationNodeBlendSpace1D struct {
	AnimationRootNode
	owner gdnative.Object
}

func (o *AnimationNodeBlendSpace1D) BaseClass() string {
	return "AnimationNodeBlendSpace1D"
}

/*
        Undocumented
	Args: [{ false index int} { false node AnimationRootNode}], Returns: void
*/
func (o *AnimationNodeBlendSpace1D) X_AddBlendPoint(index gdnative.Int, node AnimationRootNodeImplementer) {
	// log.Println("Calling AnimationNodeBlendSpace1D.X_AddBlendPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "_add_blend_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AnimationNodeBlendSpace1D) X_TreeChanged() {
	// log.Println("Calling AnimationNodeBlendSpace1D.X_TreeChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "_tree_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a new point that represents a [code]node[/code] on the virtual axis at a given position set by [code]pos[/code]. You can insert it at a specific index using the [code]at_index[/code] argument. If you use the default value for [code]at_index[/code], the point is inserted at the end of the blend points array.
	Args: [{ false node AnimationRootNode} { false pos float} {-1 true at_index int}], Returns: void
*/
func (o *AnimationNodeBlendSpace1D) AddBlendPoint(node AnimationRootNodeImplementer, pos gdnative.Real, atIndex gdnative.Int) {
	// log.Println("Calling AnimationNodeBlendSpace1D.AddBlendPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromReal(pos)
	ptrArguments[2] = gdnative.NewPointerFromInt(atIndex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "add_blend_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the number of points on the blend axis.
	Args: [], Returns: int
*/
func (o *AnimationNodeBlendSpace1D) GetBlendPointCount() gdnative.Int {
	// log.Println("Calling AnimationNodeBlendSpace1D.GetBlendPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "get_blend_point_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the [AnimationNode] referenced by the point at index [code]point[/code].
	Args: [{ false point int}], Returns: AnimationRootNode
*/
func (o *AnimationNodeBlendSpace1D) GetBlendPointNode(point gdnative.Int) AnimationRootNodeImplementer {
	// log.Println("Calling AnimationNodeBlendSpace1D.GetBlendPointNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "get_blend_point_node")

	// Call the parent method.
	// AnimationRootNode
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAnimationRootNodeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AnimationRootNodeImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AnimationRootNode" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AnimationRootNodeImplementer)
	}

	return &ret
}

/*
        Returns the position of the point at index [code]point[/code].
	Args: [{ false point int}], Returns: float
*/
func (o *AnimationNodeBlendSpace1D) GetBlendPointPosition(point gdnative.Int) gdnative.Real {
	// log.Println("Calling AnimationNodeBlendSpace1D.GetBlendPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "get_blend_point_position")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AnimationNodeBlendSpace1D) GetMaxSpace() gdnative.Real {
	// log.Println("Calling AnimationNodeBlendSpace1D.GetMaxSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "get_max_space")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AnimationNodeBlendSpace1D) GetMinSpace() gdnative.Real {
	// log.Println("Calling AnimationNodeBlendSpace1D.GetMinSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "get_min_space")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AnimationNodeBlendSpace1D) GetSnap() gdnative.Real {
	// log.Println("Calling AnimationNodeBlendSpace1D.GetSnap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "get_snap")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AnimationNodeBlendSpace1D) GetValueLabel() gdnative.String {
	// log.Println("Calling AnimationNodeBlendSpace1D.GetValueLabel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "get_value_label")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Removes the point at index [code]point[/code] from the blend axis.
	Args: [{ false point int}], Returns: void
*/
func (o *AnimationNodeBlendSpace1D) RemoveBlendPoint(point gdnative.Int) {
	// log.Println("Calling AnimationNodeBlendSpace1D.RemoveBlendPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "remove_blend_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Changes the [AnimationNode] referenced by the point at index [code]point[/code].
	Args: [{ false point int} { false node AnimationRootNode}], Returns: void
*/
func (o *AnimationNodeBlendSpace1D) SetBlendPointNode(point gdnative.Int, node AnimationRootNodeImplementer) {
	// log.Println("Calling AnimationNodeBlendSpace1D.SetBlendPointNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(point)
	ptrArguments[1] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "set_blend_point_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Updates the position of the point at index [code]point[/code] on the blend axis.
	Args: [{ false point int} { false pos float}], Returns: void
*/
func (o *AnimationNodeBlendSpace1D) SetBlendPointPosition(point gdnative.Int, pos gdnative.Real) {
	// log.Println("Calling AnimationNodeBlendSpace1D.SetBlendPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(point)
	ptrArguments[1] = gdnative.NewPointerFromReal(pos)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "set_blend_point_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false max_space float}], Returns: void
*/
func (o *AnimationNodeBlendSpace1D) SetMaxSpace(maxSpace gdnative.Real) {
	// log.Println("Calling AnimationNodeBlendSpace1D.SetMaxSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(maxSpace)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "set_max_space")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false min_space float}], Returns: void
*/
func (o *AnimationNodeBlendSpace1D) SetMinSpace(minSpace gdnative.Real) {
	// log.Println("Calling AnimationNodeBlendSpace1D.SetMinSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(minSpace)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "set_min_space")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false snap float}], Returns: void
*/
func (o *AnimationNodeBlendSpace1D) SetSnap(snap gdnative.Real) {
	// log.Println("Calling AnimationNodeBlendSpace1D.SetSnap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(snap)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "set_snap")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false text String}], Returns: void
*/
func (o *AnimationNodeBlendSpace1D) SetValueLabel(text gdnative.String) {
	// log.Println("Calling AnimationNodeBlendSpace1D.SetValueLabel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace1D", "set_value_label")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimationNodeBlendSpace1DImplementer is an interface that implements the methods
// of the AnimationNodeBlendSpace1D class.
type AnimationNodeBlendSpace1DImplementer interface {
	AnimationRootNodeImplementer
	X_AddBlendPoint(index gdnative.Int, node AnimationRootNodeImplementer)
	X_TreeChanged()
	AddBlendPoint(node AnimationRootNodeImplementer, pos gdnative.Real, atIndex gdnative.Int)
	GetBlendPointCount() gdnative.Int
	GetBlendPointNode(point gdnative.Int) AnimationRootNodeImplementer
	GetBlendPointPosition(point gdnative.Int) gdnative.Real
	GetMaxSpace() gdnative.Real
	GetMinSpace() gdnative.Real
	GetSnap() gdnative.Real
	GetValueLabel() gdnative.String
	RemoveBlendPoint(point gdnative.Int)
	SetBlendPointNode(point gdnative.Int, node AnimationRootNodeImplementer)
	SetBlendPointPosition(point gdnative.Int, pos gdnative.Real)
	SetMaxSpace(maxSpace gdnative.Real)
	SetMinSpace(minSpace gdnative.Real)
	SetSnap(snap gdnative.Real)
	SetValueLabel(text gdnative.String)
}
