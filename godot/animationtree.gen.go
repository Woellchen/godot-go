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

// AnimationTreeAnimationProcessMode is an enum for AnimationProcessMode values.
type AnimationTreeAnimationProcessMode int

const (
	AnimationTreeAnimationProcessIdle    AnimationTreeAnimationProcessMode = 1
	AnimationTreeAnimationProcessManual  AnimationTreeAnimationProcessMode = 2
	AnimationTreeAnimationProcessPhysics AnimationTreeAnimationProcessMode = 0
)

//func NewAnimationTreeFromPointer(ptr gdnative.Pointer) AnimationTree {
func newAnimationTreeFromPointer(ptr gdnative.Pointer) AnimationTree {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationTree{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Note: When linked with an [AnimationPlayer], several properties and methods of the corresponding [AnimationPlayer] will not function as expected. Playback and transitions should be handled using only the [AnimationTree] and its constituent [AnimationNode](s). The [AnimationPlayer] node should be used solely for adding, deleting, and editing animations.
*/
type AnimationTree struct {
	Node
	owner gdnative.Object
}

func (o *AnimationTree) BaseClass() string {
	return "AnimationTree"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AnimationTree) X_ClearCaches() {
	// log.Println("Calling AnimationTree.X_ClearCaches()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "_clear_caches")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 Node}], Returns: void
*/
func (o *AnimationTree) X_NodeRemoved(arg0 NodeImplementer) {
	// log.Println("Calling AnimationTree.X_NodeRemoved()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "_node_removed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AnimationTree) X_TreeChanged() {
	// log.Println("Calling AnimationTree.X_TreeChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "_tree_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AnimationTree) X_UpdateProperties() {
	// log.Println("Calling AnimationTree.X_UpdateProperties()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "_update_properties")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Manually advance the animations by the specified time (in seconds).
	Args: [{ false delta float}], Returns: void
*/
func (o *AnimationTree) Advance(delta gdnative.Real) {
	// log.Println("Calling AnimationTree.Advance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(delta)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "advance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *AnimationTree) GetAnimationPlayer() gdnative.NodePath {
	// log.Println("Calling AnimationTree.GetAnimationPlayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "get_animation_player")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.AnimationTree::AnimationProcessMode
*/
func (o *AnimationTree) GetProcessMode() AnimationTreeAnimationProcessMode {
	// log.Println("Calling AnimationTree.GetProcessMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "get_process_mode")

	// Call the parent method.
	// enum.AnimationTree::AnimationProcessMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return AnimationTreeAnimationProcessMode(ret)
}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *AnimationTree) GetRootMotionTrack() gdnative.NodePath {
	// log.Println("Calling AnimationTree.GetRootMotionTrack()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "get_root_motion_track")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Retrieve the motion of the [member root_motion_track] as a [Transform] that can be used elsewhere. If [member root_motion_track] is not a path to a track of type [constant Animation.TYPE_TRANSFORM], returns an identity transformation.
	Args: [], Returns: Transform
*/
func (o *AnimationTree) GetRootMotionTransform() gdnative.Transform {
	// log.Println("Calling AnimationTree.GetRootMotionTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "get_root_motion_transform")

	// Call the parent method.
	// Transform
	retPtr := gdnative.NewEmptyTransform()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransformFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: AnimationNode
*/
func (o *AnimationTree) GetTreeRoot() AnimationNodeImplementer {
	// log.Println("Calling AnimationTree.GetTreeRoot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "get_tree_root")

	// Call the parent method.
	// AnimationNode
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAnimationNodeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AnimationNodeImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AnimationNode" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AnimationNodeImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AnimationTree) IsActive() gdnative.Bool {
	// log.Println("Calling AnimationTree.IsActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "is_active")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false old_name String} { false new_name String}], Returns: void
*/
func (o *AnimationTree) RenameParameter(oldName gdnative.String, newName gdnative.String) {
	// log.Println("Calling AnimationTree.RenameParameter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(oldName)
	ptrArguments[1] = gdnative.NewPointerFromString(newName)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "rename_parameter")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false active bool}], Returns: void
*/
func (o *AnimationTree) SetActive(active gdnative.Bool) {
	// log.Println("Calling AnimationTree.SetActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(active)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "set_active")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false root NodePath}], Returns: void
*/
func (o *AnimationTree) SetAnimationPlayer(root gdnative.NodePath) {
	// log.Println("Calling AnimationTree.SetAnimationPlayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(root)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "set_animation_player")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *AnimationTree) SetProcessMode(mode gdnative.Int) {
	// log.Println("Calling AnimationTree.SetProcessMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "set_process_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false path NodePath}], Returns: void
*/
func (o *AnimationTree) SetRootMotionTrack(path gdnative.NodePath) {
	// log.Println("Calling AnimationTree.SetRootMotionTrack()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "set_root_motion_track")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false root AnimationNode}], Returns: void
*/
func (o *AnimationTree) SetTreeRoot(root AnimationNodeImplementer) {
	// log.Println("Calling AnimationTree.SetTreeRoot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(root.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationTree", "set_tree_root")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimationTreeImplementer is an interface that implements the methods
// of the AnimationTree class.
type AnimationTreeImplementer interface {
	NodeImplementer
	X_ClearCaches()
	X_NodeRemoved(arg0 NodeImplementer)
	X_TreeChanged()
	X_UpdateProperties()
	Advance(delta gdnative.Real)
	GetAnimationPlayer() gdnative.NodePath
	GetRootMotionTrack() gdnative.NodePath
	GetRootMotionTransform() gdnative.Transform
	GetTreeRoot() AnimationNodeImplementer
	IsActive() gdnative.Bool
	RenameParameter(oldName gdnative.String, newName gdnative.String)
	SetActive(active gdnative.Bool)
	SetAnimationPlayer(root gdnative.NodePath)
	SetProcessMode(mode gdnative.Int)
	SetRootMotionTrack(path gdnative.NodePath)
	SetTreeRoot(root AnimationNodeImplementer)
}
