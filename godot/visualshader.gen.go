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

// VisualShaderType is an enum for Type values.
type VisualShaderType int

const (
	VisualShaderTypeFragment VisualShaderType = 1
	VisualShaderTypeLight    VisualShaderType = 2
	VisualShaderTypeMax      VisualShaderType = 3
	VisualShaderTypeVertex   VisualShaderType = 0
)

//func NewVisualShaderFromPointer(ptr gdnative.Pointer) VisualShader {
func newVisualShaderFromPointer(ptr gdnative.Pointer) VisualShader {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShader{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This class allows you to define a custom shader program that can be used for various materials to render objects. The visual shader editor creates the shader.
*/
type VisualShader struct {
	Shader
	owner gdnative.Object
}

func (o *VisualShader) BaseClass() string {
	return "VisualShader"
}

/*
        Undocumented
	Args: [{ false arg0 int} { false arg1 int}], Returns: void
*/
func (o *VisualShader) X_InputTypeChanged(arg0 gdnative.Int, arg1 gdnative.Int) {
	//log.Println("Calling VisualShader.X_InputTypeChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)
	ptrArguments[1] = gdnative.NewPointerFromInt(arg1)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "_input_type_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *VisualShader) X_QueueUpdate() {
	//log.Println("Calling VisualShader.X_QueueUpdate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "_queue_update")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *VisualShader) X_UpdateShader() {
	//log.Println("Calling VisualShader.X_UpdateShader()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "_update_shader")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds the specified node to the shader.
	Args: [{ false type int} { false node VisualShaderNode} { false position Vector2} { false id int}], Returns: void
*/
func (o *VisualShader) AddNode(aType gdnative.Int, node VisualShaderNodeImplementer, position gdnative.Vector2, id gdnative.Int) {
	//log.Println("Calling VisualShader.AddNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)
	ptrArguments[1] = gdnative.NewPointerFromObject(node.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromVector2(position)
	ptrArguments[3] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "add_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns [code]true[/code] if the specified nodes and ports can be connected together.
	Args: [{ false type int} { false from_node int} { false from_port int} { false to_node int} { false to_port int}], Returns: bool
*/
func (o *VisualShader) CanConnectNodes(aType gdnative.Int, fromNode gdnative.Int, fromPort gdnative.Int, toNode gdnative.Int, toPort gdnative.Int) gdnative.Bool {
	//log.Println("Calling VisualShader.CanConnectNodes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)
	ptrArguments[1] = gdnative.NewPointerFromInt(fromNode)
	ptrArguments[2] = gdnative.NewPointerFromInt(fromPort)
	ptrArguments[3] = gdnative.NewPointerFromInt(toNode)
	ptrArguments[4] = gdnative.NewPointerFromInt(toPort)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "can_connect_nodes")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Connects the specified nodes and ports.
	Args: [{ false type int} { false from_node int} { false from_port int} { false to_node int} { false to_port int}], Returns: enum.Error
*/
func (o *VisualShader) ConnectNodes(aType gdnative.Int, fromNode gdnative.Int, fromPort gdnative.Int, toNode gdnative.Int, toPort gdnative.Int) gdnative.Error {
	//log.Println("Calling VisualShader.ConnectNodes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)
	ptrArguments[1] = gdnative.NewPointerFromInt(fromNode)
	ptrArguments[2] = gdnative.NewPointerFromInt(fromPort)
	ptrArguments[3] = gdnative.NewPointerFromInt(toNode)
	ptrArguments[4] = gdnative.NewPointerFromInt(toPort)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "connect_nodes")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Connects the specified nodes and ports, even if they can't be connected. Such connection is invalid and will not function properly.
	Args: [{ false type int} { false from_node int} { false from_port int} { false to_node int} { false to_port int}], Returns: void
*/
func (o *VisualShader) ConnectNodesForced(aType gdnative.Int, fromNode gdnative.Int, fromPort gdnative.Int, toNode gdnative.Int, toPort gdnative.Int) {
	//log.Println("Calling VisualShader.ConnectNodesForced()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)
	ptrArguments[1] = gdnative.NewPointerFromInt(fromNode)
	ptrArguments[2] = gdnative.NewPointerFromInt(fromPort)
	ptrArguments[3] = gdnative.NewPointerFromInt(toNode)
	ptrArguments[4] = gdnative.NewPointerFromInt(toPort)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "connect_nodes_forced")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Connects the specified nodes and ports.
	Args: [{ false type int} { false from_node int} { false from_port int} { false to_node int} { false to_port int}], Returns: void
*/
func (o *VisualShader) DisconnectNodes(aType gdnative.Int, fromNode gdnative.Int, fromPort gdnative.Int, toNode gdnative.Int, toPort gdnative.Int) {
	//log.Println("Calling VisualShader.DisconnectNodes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)
	ptrArguments[1] = gdnative.NewPointerFromInt(fromNode)
	ptrArguments[2] = gdnative.NewPointerFromInt(fromPort)
	ptrArguments[3] = gdnative.NewPointerFromInt(toNode)
	ptrArguments[4] = gdnative.NewPointerFromInt(toPort)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "disconnect_nodes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *VisualShader) GetGraphOffset() gdnative.Vector2 {
	//log.Println("Calling VisualShader.GetGraphOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "get_graph_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns the shader node instance with specified [code]type[/code] and [code]id[/code].
	Args: [{ false type int} { false id int}], Returns: VisualShaderNode
*/
func (o *VisualShader) GetNode(aType gdnative.Int, id gdnative.Int) VisualShaderNodeImplementer {
	//log.Println("Calling VisualShader.GetNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)
	ptrArguments[1] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "get_node")

	// Call the parent method.
	// VisualShaderNode
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newVisualShaderNodeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(VisualShaderNodeImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "VisualShaderNode" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(VisualShaderNodeImplementer)
	}

	return &ret
}

/*
        Returns the list of connected nodes with the specified type.
	Args: [{ false type int}], Returns: Array
*/
func (o *VisualShader) GetNodeConnections(aType gdnative.Int) gdnative.Array {
	//log.Println("Calling VisualShader.GetNodeConnections()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "get_node_connections")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the list of all nodes in the shader with the specified type.
	Args: [{ false type int}], Returns: PoolIntArray
*/
func (o *VisualShader) GetNodeList(aType gdnative.Int) gdnative.PoolIntArray {
	//log.Println("Calling VisualShader.GetNodeList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "get_node_list")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the position of the specified node within the shader graph.
	Args: [{ false type int} { false id int}], Returns: Vector2
*/
func (o *VisualShader) GetNodePosition(aType gdnative.Int, id gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling VisualShader.GetNodePosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)
	ptrArguments[1] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "get_node_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false type int}], Returns: int
*/
func (o *VisualShader) GetValidNodeId(aType gdnative.Int) gdnative.Int {
	//log.Println("Calling VisualShader.GetValidNodeId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "get_valid_node_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the specified node and port connection exist.
	Args: [{ false type int} { false from_node int} { false from_port int} { false to_node int} { false to_port int}], Returns: bool
*/
func (o *VisualShader) IsNodeConnection(aType gdnative.Int, fromNode gdnative.Int, fromPort gdnative.Int, toNode gdnative.Int, toPort gdnative.Int) gdnative.Bool {
	//log.Println("Calling VisualShader.IsNodeConnection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)
	ptrArguments[1] = gdnative.NewPointerFromInt(fromNode)
	ptrArguments[2] = gdnative.NewPointerFromInt(fromPort)
	ptrArguments[3] = gdnative.NewPointerFromInt(toNode)
	ptrArguments[4] = gdnative.NewPointerFromInt(toPort)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "is_node_connection")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes the specified node from the shader.
	Args: [{ false type int} { false id int}], Returns: void
*/
func (o *VisualShader) RemoveNode(aType gdnative.Int, id gdnative.Int) {
	//log.Println("Calling VisualShader.RemoveNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)
	ptrArguments[1] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "remove_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false offset Vector2}], Returns: void
*/
func (o *VisualShader) SetGraphOffset(offset gdnative.Vector2) {
	//log.Println("Calling VisualShader.SetGraphOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "set_graph_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the mode of this shader.
	Args: [{ false mode int}], Returns: void
*/
func (o *VisualShader) SetMode(mode gdnative.Int) {
	//log.Println("Calling VisualShader.SetMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "set_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the position of the specified node.
	Args: [{ false type int} { false id int} { false position Vector2}], Returns: void
*/
func (o *VisualShader) SetNodePosition(aType gdnative.Int, id gdnative.Int, position gdnative.Vector2) {
	//log.Println("Calling VisualShader.SetNodePosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)
	ptrArguments[1] = gdnative.NewPointerFromInt(id)
	ptrArguments[2] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShader", "set_node_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderImplementer is an interface that implements the methods
// of the VisualShader class.
type VisualShaderImplementer interface {
	ShaderImplementer
	X_InputTypeChanged(arg0 gdnative.Int, arg1 gdnative.Int)
	X_QueueUpdate()
	X_UpdateShader()
	AddNode(aType gdnative.Int, node VisualShaderNodeImplementer, position gdnative.Vector2, id gdnative.Int)
	CanConnectNodes(aType gdnative.Int, fromNode gdnative.Int, fromPort gdnative.Int, toNode gdnative.Int, toPort gdnative.Int) gdnative.Bool
	ConnectNodesForced(aType gdnative.Int, fromNode gdnative.Int, fromPort gdnative.Int, toNode gdnative.Int, toPort gdnative.Int)
	DisconnectNodes(aType gdnative.Int, fromNode gdnative.Int, fromPort gdnative.Int, toNode gdnative.Int, toPort gdnative.Int)
	GetGraphOffset() gdnative.Vector2
	GetNode(aType gdnative.Int, id gdnative.Int) VisualShaderNodeImplementer
	GetNodeConnections(aType gdnative.Int) gdnative.Array
	GetNodeList(aType gdnative.Int) gdnative.PoolIntArray
	GetNodePosition(aType gdnative.Int, id gdnative.Int) gdnative.Vector2
	GetValidNodeId(aType gdnative.Int) gdnative.Int
	IsNodeConnection(aType gdnative.Int, fromNode gdnative.Int, fromPort gdnative.Int, toNode gdnative.Int, toPort gdnative.Int) gdnative.Bool
	RemoveNode(aType gdnative.Int, id gdnative.Int)
	SetGraphOffset(offset gdnative.Vector2)
	SetMode(mode gdnative.Int)
	SetNodePosition(aType gdnative.Int, id gdnative.Int, position gdnative.Vector2)
}
