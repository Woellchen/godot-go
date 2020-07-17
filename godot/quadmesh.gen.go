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

//func NewQuadMeshFromPointer(ptr gdnative.Pointer) QuadMesh {
func newQuadMeshFromPointer(ptr gdnative.Pointer) QuadMesh {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := QuadMesh{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Class representing a square [PrimitiveMesh]. This flat mesh does not have a thickness. By default, this mesh is aligned on the X and Y axes; this default rotation is more suited for use with billboarded materials. Unlike [PlaneMesh], this mesh doesn't provide subdivision options.
*/
type QuadMesh struct {
	PrimitiveMesh
	owner gdnative.Object
}

func (o *QuadMesh) BaseClass() string {
	return "QuadMesh"
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *QuadMesh) GetSize() gdnative.Vector2 {
	//log.Println("Calling QuadMesh.GetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("QuadMesh", "get_size")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false size Vector2}], Returns: void
*/
func (o *QuadMesh) SetSize(size gdnative.Vector2) {
	//log.Println("Calling QuadMesh.SetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("QuadMesh", "set_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// QuadMeshImplementer is an interface that implements the methods
// of the QuadMesh class.
type QuadMeshImplementer interface {
	PrimitiveMeshImplementer
	GetSize() gdnative.Vector2
	SetSize(size gdnative.Vector2)
}
