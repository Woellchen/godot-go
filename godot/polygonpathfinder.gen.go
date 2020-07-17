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

//func NewPolygonPathFinderFromPointer(ptr gdnative.Pointer) PolygonPathFinder {
func newPolygonPathFinderFromPointer(ptr gdnative.Pointer) PolygonPathFinder {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PolygonPathFinder{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type PolygonPathFinder struct {
	Resource
	owner gdnative.Object
}

func (o *PolygonPathFinder) BaseClass() string {
	return "PolygonPathFinder"
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/
func (o *PolygonPathFinder) X_GetData() gdnative.Dictionary {
	//log.Println("Calling PolygonPathFinder.X_GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PolygonPathFinder", "_get_data")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false arg0 Dictionary}], Returns: void
*/
func (o *PolygonPathFinder) X_SetData(arg0 gdnative.Dictionary) {
	//log.Println("Calling PolygonPathFinder.X_SetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PolygonPathFinder", "_set_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false from Vector2} { false to Vector2}], Returns: PoolVector2Array
*/
func (o *PolygonPathFinder) FindPath(from gdnative.Vector2, to gdnative.Vector2) gdnative.PoolVector2Array {
	//log.Println("Calling PolygonPathFinder.FindPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector2(from)
	ptrArguments[1] = gdnative.NewPointerFromVector2(to)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PolygonPathFinder", "find_path")

	// Call the parent method.
	// PoolVector2Array
	retPtr := gdnative.NewEmptyPoolVector2Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector2ArrayFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: Rect2
*/
func (o *PolygonPathFinder) GetBounds() gdnative.Rect2 {
	//log.Println("Calling PolygonPathFinder.GetBounds()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PolygonPathFinder", "get_bounds")

	// Call the parent method.
	// Rect2
	retPtr := gdnative.NewEmptyRect2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRect2FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false point Vector2}], Returns: Vector2
*/
func (o *PolygonPathFinder) GetClosestPoint(point gdnative.Vector2) gdnative.Vector2 {
	//log.Println("Calling PolygonPathFinder.GetClosestPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PolygonPathFinder", "get_closest_point")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false from Vector2} { false to Vector2}], Returns: PoolVector2Array
*/
func (o *PolygonPathFinder) GetIntersections(from gdnative.Vector2, to gdnative.Vector2) gdnative.PoolVector2Array {
	//log.Println("Calling PolygonPathFinder.GetIntersections()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector2(from)
	ptrArguments[1] = gdnative.NewPointerFromVector2(to)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PolygonPathFinder", "get_intersections")

	// Call the parent method.
	// PoolVector2Array
	retPtr := gdnative.NewEmptyPoolVector2Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector2ArrayFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false idx int}], Returns: float
*/
func (o *PolygonPathFinder) GetPointPenalty(idx gdnative.Int) gdnative.Real {
	//log.Println("Calling PolygonPathFinder.GetPointPenalty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PolygonPathFinder", "get_point_penalty")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false point Vector2}], Returns: bool
*/
func (o *PolygonPathFinder) IsPointInside(point gdnative.Vector2) gdnative.Bool {
	//log.Println("Calling PolygonPathFinder.IsPointInside()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PolygonPathFinder", "is_point_inside")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false idx int} { false penalty float}], Returns: void
*/
func (o *PolygonPathFinder) SetPointPenalty(idx gdnative.Int, penalty gdnative.Real) {
	//log.Println("Calling PolygonPathFinder.SetPointPenalty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromReal(penalty)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PolygonPathFinder", "set_point_penalty")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false points PoolVector2Array} { false connections PoolIntArray}], Returns: void
*/
func (o *PolygonPathFinder) Setup(points gdnative.PoolVector2Array, connections gdnative.PoolIntArray) {
	//log.Println("Calling PolygonPathFinder.Setup()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector2Array(points)
	ptrArguments[1] = gdnative.NewPointerFromPoolIntArray(connections)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PolygonPathFinder", "setup")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// PolygonPathFinderImplementer is an interface that implements the methods
// of the PolygonPathFinder class.
type PolygonPathFinderImplementer interface {
	ResourceImplementer
	X_GetData() gdnative.Dictionary
	X_SetData(arg0 gdnative.Dictionary)
	FindPath(from gdnative.Vector2, to gdnative.Vector2) gdnative.PoolVector2Array
	GetBounds() gdnative.Rect2
	GetClosestPoint(point gdnative.Vector2) gdnative.Vector2
	GetIntersections(from gdnative.Vector2, to gdnative.Vector2) gdnative.PoolVector2Array
	GetPointPenalty(idx gdnative.Int) gdnative.Real
	IsPointInside(point gdnative.Vector2) gdnative.Bool
	SetPointPenalty(idx gdnative.Int, penalty gdnative.Real)
	Setup(points gdnative.PoolVector2Array, connections gdnative.PoolIntArray)
}
