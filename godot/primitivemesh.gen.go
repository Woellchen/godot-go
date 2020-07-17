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

//func NewPrimitiveMeshFromPointer(ptr gdnative.Pointer) PrimitiveMesh {
func newPrimitiveMeshFromPointer(ptr gdnative.Pointer) PrimitiveMesh {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PrimitiveMesh{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base class for all primitive meshes. Handles applying a [Material] to a primitive mesh. Examples include [CapsuleMesh], [CubeMesh], [CylinderMesh], [PlaneMesh], [PrismMesh], [QuadMesh], and [SphereMesh].
*/
type PrimitiveMesh struct {
	Mesh
	owner gdnative.Object
}

func (o *PrimitiveMesh) BaseClass() string {
	return "PrimitiveMesh"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *PrimitiveMesh) X_Update() {
	//log.Println("Calling PrimitiveMesh.X_Update()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrimitiveMesh", "_update")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: AABB
*/
func (o *PrimitiveMesh) GetCustomAabb() gdnative.Aabb {
	//log.Println("Calling PrimitiveMesh.GetCustomAabb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrimitiveMesh", "get_custom_aabb")

	// Call the parent method.
	// AABB
	retPtr := gdnative.NewEmptyAabb()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewAabbFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *PrimitiveMesh) GetFlipFaces() gdnative.Bool {
	//log.Println("Calling PrimitiveMesh.GetFlipFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrimitiveMesh", "get_flip_faces")

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
	Args: [], Returns: Material
*/
func (o *PrimitiveMesh) GetMaterial() MaterialImplementer {
	//log.Println("Calling PrimitiveMesh.GetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrimitiveMesh", "get_material")

	// Call the parent method.
	// Material
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newMaterialFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(MaterialImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Material" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(MaterialImplementer)
	}

	return &ret
}

/*
        Returns mesh arrays used to constitute surface of [Mesh]. Mesh arrays can be used with [ArrayMesh] to create new surfaces.
	Args: [], Returns: Array
*/
func (o *PrimitiveMesh) GetMeshArrays() gdnative.Array {
	//log.Println("Calling PrimitiveMesh.GetMeshArrays()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrimitiveMesh", "get_mesh_arrays")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false aabb AABB}], Returns: void
*/
func (o *PrimitiveMesh) SetCustomAabb(aabb gdnative.Aabb) {
	//log.Println("Calling PrimitiveMesh.SetCustomAabb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromAabb(aabb)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrimitiveMesh", "set_custom_aabb")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flip_faces bool}], Returns: void
*/
func (o *PrimitiveMesh) SetFlipFaces(flipFaces gdnative.Bool) {
	//log.Println("Calling PrimitiveMesh.SetFlipFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(flipFaces)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrimitiveMesh", "set_flip_faces")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false material Material}], Returns: void
*/
func (o *PrimitiveMesh) SetMaterial(material MaterialImplementer) {
	//log.Println("Calling PrimitiveMesh.SetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PrimitiveMesh", "set_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// PrimitiveMeshImplementer is an interface that implements the methods
// of the PrimitiveMesh class.
type PrimitiveMeshImplementer interface {
	MeshImplementer
	X_Update()
	GetCustomAabb() gdnative.Aabb
	GetFlipFaces() gdnative.Bool
	GetMaterial() MaterialImplementer
	GetMeshArrays() gdnative.Array
	SetCustomAabb(aabb gdnative.Aabb)
	SetFlipFaces(flipFaces gdnative.Bool)
	SetMaterial(material MaterialImplementer)
}
