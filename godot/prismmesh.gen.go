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

//func NewPrismMeshFromPointer(ptr gdnative.Pointer) PrismMesh {
func newPrismMeshFromPointer(ptr gdnative.Pointer) PrismMesh {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PrismMesh{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Class representing a prism-shaped [PrimitiveMesh].
*/
type PrismMesh struct {
	PrimitiveMesh
	owner gdnative.Object
}

func (o *PrismMesh) BaseClass() string {
	return "PrismMesh"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *PrismMesh) GetLeftToRight() gdnative.Real {
	// log.Println("Calling PrismMesh.GetLeftToRight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrismMesh", "get_left_to_right")

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
	Args: [], Returns: Vector3
*/
func (o *PrismMesh) GetSize() gdnative.Vector3 {
	// log.Println("Calling PrismMesh.GetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrismMesh", "get_size")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *PrismMesh) GetSubdivideDepth() gdnative.Int {
	// log.Println("Calling PrismMesh.GetSubdivideDepth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrismMesh", "get_subdivide_depth")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *PrismMesh) GetSubdivideHeight() gdnative.Int {
	// log.Println("Calling PrismMesh.GetSubdivideHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrismMesh", "get_subdivide_height")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *PrismMesh) GetSubdivideWidth() gdnative.Int {
	// log.Println("Calling PrismMesh.GetSubdivideWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrismMesh", "get_subdivide_width")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false left_to_right float}], Returns: void
*/
func (o *PrismMesh) SetLeftToRight(leftToRight gdnative.Real) {
	// log.Println("Calling PrismMesh.SetLeftToRight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(leftToRight)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrismMesh", "set_left_to_right")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false size Vector3}], Returns: void
*/
func (o *PrismMesh) SetSize(size gdnative.Vector3) {
	// log.Println("Calling PrismMesh.SetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrismMesh", "set_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false segments int}], Returns: void
*/
func (o *PrismMesh) SetSubdivideDepth(segments gdnative.Int) {
	// log.Println("Calling PrismMesh.SetSubdivideDepth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(segments)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrismMesh", "set_subdivide_depth")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false segments int}], Returns: void
*/
func (o *PrismMesh) SetSubdivideHeight(segments gdnative.Int) {
	// log.Println("Calling PrismMesh.SetSubdivideHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(segments)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrismMesh", "set_subdivide_height")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false segments int}], Returns: void
*/
func (o *PrismMesh) SetSubdivideWidth(segments gdnative.Int) {
	// log.Println("Calling PrismMesh.SetSubdivideWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(segments)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrismMesh", "set_subdivide_width")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// PrismMeshImplementer is an interface that implements the methods
// of the PrismMesh class.
type PrismMeshImplementer interface {
	PrimitiveMeshImplementer
	GetLeftToRight() gdnative.Real
	GetSize() gdnative.Vector3
	GetSubdivideDepth() gdnative.Int
	GetSubdivideHeight() gdnative.Int
	GetSubdivideWidth() gdnative.Int
	SetLeftToRight(leftToRight gdnative.Real)
	SetSize(size gdnative.Vector3)
	SetSubdivideDepth(segments gdnative.Int)
	SetSubdivideHeight(segments gdnative.Int)
	SetSubdivideWidth(segments gdnative.Int)
}
