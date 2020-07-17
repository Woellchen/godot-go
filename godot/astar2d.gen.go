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

//func NewAStar2DFromPointer(ptr gdnative.Pointer) AStar2D {
func newAStar2DFromPointer(ptr gdnative.Pointer) AStar2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AStar2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This is a wrapper for the [AStar] class which uses 2D vectors instead of 3D vectors.
*/
type AStar2D struct {
	Reference
	owner gdnative.Object
}

func (o *AStar2D) BaseClass() string {
	return "AStar2D"
}

/*
        Called when computing the cost between two connected points. Note that this function is hidden in the default [code]AStar2D[/code] class.
	Args: [{ false from_id int} { false to_id int}], Returns: float
*/
func (o *AStar2D) X_ComputeCost(fromId gdnative.Int, toId gdnative.Int) gdnative.Real {
	//log.Println("Calling AStar2D.X_ComputeCost()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(fromId)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "_compute_cost")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Called when estimating the cost between a point and the path's ending point. Note that this function is hidden in the default [code]AStar2D[/code] class.
	Args: [{ false from_id int} { false to_id int}], Returns: float
*/
func (o *AStar2D) X_EstimateCost(fromId gdnative.Int, toId gdnative.Int) gdnative.Real {
	//log.Println("Calling AStar2D.X_EstimateCost()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(fromId)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "_estimate_cost")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Adds a new point at the given position with the given identifier. The algorithm prefers points with lower [code]weight_scale[/code] to form a path. The [code]id[/code] must be 0 or larger, and the [code]weight_scale[/code] must be 1 or larger. [codeblock] var astar = AStar2D.new() astar.add_point(1, Vector2(1, 0), 4) # Adds the point (1, 0) with weight_scale 4 and id 1 [/codeblock] If there already exists a point for the given [code]id[/code], its position and weight scale are updated to the given values.
	Args: [{ false id int} { false position Vector2} {1 true weight_scale float}], Returns: void
*/
func (o *AStar2D) AddPoint(id gdnative.Int, position gdnative.Vector2, weightScale gdnative.Real) {
	//log.Println("Calling AStar2D.AddPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromVector2(position)
	ptrArguments[2] = gdnative.NewPointerFromReal(weightScale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "add_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns whether there is a connection/segment between the given points.
	Args: [{ false id int} { false to_id int}], Returns: bool
*/
func (o *AStar2D) ArePointsConnected(id gdnative.Int, toId gdnative.Int) gdnative.Bool {
	//log.Println("Calling AStar2D.ArePointsConnected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "are_points_connected")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Clears all the points and segments.
	Args: [], Returns: void
*/
func (o *AStar2D) Clear() {
	//log.Println("Calling AStar2D.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Creates a segment between the given points. If [code]bidirectional[/code] is [code]false[/code], only movement from [code]id[/code] to [code]to_id[/code] is allowed, not the reverse direction. [codeblock] var astar = AStar2D.new() astar.add_point(1, Vector2(1, 1)) astar.add_point(2, Vector2(0, 5)) astar.connect_points(1, 2, false) [/codeblock]
	Args: [{ false id int} { false to_id int} {True true bidirectional bool}], Returns: void
*/
func (o *AStar2D) ConnectPoints(id gdnative.Int, toId gdnative.Int, bidirectional gdnative.Bool) {
	//log.Println("Calling AStar2D.ConnectPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)
	ptrArguments[2] = gdnative.NewPointerFromBool(bidirectional)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "connect_points")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Deletes the segment between the given points.
	Args: [{ false id int} { false to_id int}], Returns: void
*/
func (o *AStar2D) DisconnectPoints(id gdnative.Int, toId gdnative.Int) {
	//log.Println("Calling AStar2D.DisconnectPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "disconnect_points")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the next available point ID with no point associated to it.
	Args: [], Returns: int
*/
func (o *AStar2D) GetAvailablePointId() gdnative.Int {
	//log.Println("Calling AStar2D.GetAvailablePointId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "get_available_point_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the ID of the closest point to [code]to_position[/code], optionally taking disabled points into account. Returns [code]-1[/code] if there are no points in the points pool. [b]Note:[/b] If several points are the closest to [code]to_position[/code], the one with the smallest ID will be returned, ensuring a deterministic result.
	Args: [{ false to_position Vector2} {False true include_disabled bool}], Returns: int
*/
func (o *AStar2D) GetClosestPoint(toPosition gdnative.Vector2, includeDisabled gdnative.Bool) gdnative.Int {
	//log.Println("Calling AStar2D.GetClosestPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector2(toPosition)
	ptrArguments[1] = gdnative.NewPointerFromBool(includeDisabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "get_closest_point")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the closest position to [code]to_position[/code] that resides inside a segment between two connected points. [codeblock] var astar = AStar2D.new() astar.add_point(1, Vector2(0, 0)) astar.add_point(2, Vector2(0, 5)) astar.connect_points(1, 2) var res = astar.get_closest_position_in_segment(Vector2(3, 3)) # Returns (0, 3) [/codeblock] The result is in the segment that goes from [code]y = 0[/code] to [code]y = 5[/code]. It's the closest position in the segment to the given point.
	Args: [{ false to_position Vector2}], Returns: Vector2
*/
func (o *AStar2D) GetClosestPositionInSegment(toPosition gdnative.Vector2) gdnative.Vector2 {
	//log.Println("Calling AStar2D.GetClosestPositionInSegment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(toPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "get_closest_position_in_segment")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns an array with the IDs of the points that form the path found by AStar2D between the given points. The array is ordered from the starting point to the ending point of the path. [codeblock] var astar = AStar2D.new() astar.add_point(1, Vector2(0, 0)) astar.add_point(2, Vector2(0, 1), 1) # Default weight is 1 astar.add_point(3, Vector2(1, 1)) astar.add_point(4, Vector2(2, 0)) astar.connect_points(1, 2, false) astar.connect_points(2, 3, false) astar.connect_points(4, 3, false) astar.connect_points(1, 4, false) var res = astar.get_id_path(1, 3) # Returns [1, 2, 3] [/codeblock] If you change the 2nd point's weight to 3, then the result will be [code][1, 4, 3][/code] instead, because now even though the distance is longer, it's "easier" to get through point 4 than through point 2.
	Args: [{ false from_id int} { false to_id int}], Returns: PoolIntArray
*/
func (o *AStar2D) GetIdPath(fromId gdnative.Int, toId gdnative.Int) gdnative.PoolIntArray {
	//log.Println("Calling AStar2D.GetIdPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(fromId)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "get_id_path")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the capacity of the structure backing the points, useful in conjunction with [code]reserve_space[/code].
	Args: [], Returns: int
*/
func (o *AStar2D) GetPointCapacity() gdnative.Int {
	//log.Println("Calling AStar2D.GetPointCapacity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "get_point_capacity")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns an array with the IDs of the points that form the connection with the given point. [codeblock] var astar = AStar2D.new() astar.add_point(1, Vector2(0, 0)) astar.add_point(2, Vector2(0, 1)) astar.add_point(3, Vector2(1, 1)) astar.add_point(4, Vector2(2, 0)) astar.connect_points(1, 2, true) astar.connect_points(1, 3, true) var neighbors = astar.get_point_connections(1) # Returns [2, 3] [/codeblock]
	Args: [{ false id int}], Returns: PoolIntArray
*/
func (o *AStar2D) GetPointConnections(id gdnative.Int) gdnative.PoolIntArray {
	//log.Println("Calling AStar2D.GetPointConnections()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "get_point_connections")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the number of points currently in the points pool.
	Args: [], Returns: int
*/
func (o *AStar2D) GetPointCount() gdnative.Int {
	//log.Println("Calling AStar2D.GetPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "get_point_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns an array with the points that are in the path found by AStar2D between the given points. The array is ordered from the starting point to the ending point of the path.
	Args: [{ false from_id int} { false to_id int}], Returns: PoolVector2Array
*/
func (o *AStar2D) GetPointPath(fromId gdnative.Int, toId gdnative.Int) gdnative.PoolVector2Array {
	//log.Println("Calling AStar2D.GetPointPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(fromId)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "get_point_path")

	// Call the parent method.
	// PoolVector2Array
	retPtr := gdnative.NewEmptyPoolVector2Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector2ArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the position of the point associated with the given [code]id[/code].
	Args: [{ false id int}], Returns: Vector2
*/
func (o *AStar2D) GetPointPosition(id gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling AStar2D.GetPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "get_point_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns the weight scale of the point associated with the given [code]id[/code].
	Args: [{ false id int}], Returns: float
*/
func (o *AStar2D) GetPointWeightScale(id gdnative.Int) gdnative.Real {
	//log.Println("Calling AStar2D.GetPointWeightScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "get_point_weight_scale")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns an array of all points.
	Args: [], Returns: Array
*/
func (o *AStar2D) GetPoints() gdnative.Array {
	//log.Println("Calling AStar2D.GetPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "get_points")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns whether a point associated with the given [code]id[/code] exists.
	Args: [{ false id int}], Returns: bool
*/
func (o *AStar2D) HasPoint(id gdnative.Int) gdnative.Bool {
	//log.Println("Calling AStar2D.HasPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "has_point")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns whether a point is disabled or not for pathfinding. By default, all points are enabled.
	Args: [{ false id int}], Returns: bool
*/
func (o *AStar2D) IsPointDisabled(id gdnative.Int) gdnative.Bool {
	//log.Println("Calling AStar2D.IsPointDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "is_point_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes the point associated with the given [code]id[/code] from the points pool.
	Args: [{ false id int}], Returns: void
*/
func (o *AStar2D) RemovePoint(id gdnative.Int) {
	//log.Println("Calling AStar2D.RemovePoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "remove_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Reserves space internally for [code]num_nodes[/code] points, useful if you're adding a known large number of points at once, for a grid for instance. New capacity must be greater or equals to old capacity.
	Args: [{ false num_nodes int}], Returns: void
*/
func (o *AStar2D) ReserveSpace(numNodes gdnative.Int) {
	//log.Println("Calling AStar2D.ReserveSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(numNodes)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "reserve_space")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Disables or enables the specified point for pathfinding. Useful for making a temporary obstacle.
	Args: [{ false id int} {True true disabled bool}], Returns: void
*/
func (o *AStar2D) SetPointDisabled(id gdnative.Int, disabled gdnative.Bool) {
	//log.Println("Calling AStar2D.SetPointDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromBool(disabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "set_point_disabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the [code]position[/code] for the point with the given [code]id[/code].
	Args: [{ false id int} { false position Vector2}], Returns: void
*/
func (o *AStar2D) SetPointPosition(id gdnative.Int, position gdnative.Vector2) {
	//log.Println("Calling AStar2D.SetPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "set_point_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the [code]weight_scale[/code] for the point with the given [code]id[/code].
	Args: [{ false id int} { false weight_scale float}], Returns: void
*/
func (o *AStar2D) SetPointWeightScale(id gdnative.Int, weightScale gdnative.Real) {
	//log.Println("Calling AStar2D.SetPointWeightScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromReal(weightScale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar2D", "set_point_weight_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AStar2DImplementer is an interface that implements the methods
// of the AStar2D class.
type AStar2DImplementer interface {
	ReferenceImplementer
	X_ComputeCost(fromId gdnative.Int, toId gdnative.Int) gdnative.Real
	X_EstimateCost(fromId gdnative.Int, toId gdnative.Int) gdnative.Real
	AddPoint(id gdnative.Int, position gdnative.Vector2, weightScale gdnative.Real)
	ArePointsConnected(id gdnative.Int, toId gdnative.Int) gdnative.Bool
	Clear()
	ConnectPoints(id gdnative.Int, toId gdnative.Int, bidirectional gdnative.Bool)
	DisconnectPoints(id gdnative.Int, toId gdnative.Int)
	GetAvailablePointId() gdnative.Int
	GetClosestPoint(toPosition gdnative.Vector2, includeDisabled gdnative.Bool) gdnative.Int
	GetClosestPositionInSegment(toPosition gdnative.Vector2) gdnative.Vector2
	GetIdPath(fromId gdnative.Int, toId gdnative.Int) gdnative.PoolIntArray
	GetPointCapacity() gdnative.Int
	GetPointConnections(id gdnative.Int) gdnative.PoolIntArray
	GetPointCount() gdnative.Int
	GetPointPath(fromId gdnative.Int, toId gdnative.Int) gdnative.PoolVector2Array
	GetPointPosition(id gdnative.Int) gdnative.Vector2
	GetPointWeightScale(id gdnative.Int) gdnative.Real
	GetPoints() gdnative.Array
	HasPoint(id gdnative.Int) gdnative.Bool
	IsPointDisabled(id gdnative.Int) gdnative.Bool
	RemovePoint(id gdnative.Int)
	ReserveSpace(numNodes gdnative.Int)
	SetPointDisabled(id gdnative.Int, disabled gdnative.Bool)
	SetPointPosition(id gdnative.Int, position gdnative.Vector2)
	SetPointWeightScale(id gdnative.Int, weightScale gdnative.Real)
}
