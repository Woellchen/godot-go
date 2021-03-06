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

//func NewMultiMeshInstance2DFromPointer(ptr gdnative.Pointer) MultiMeshInstance2D {
func newMultiMeshInstance2DFromPointer(ptr gdnative.Pointer) MultiMeshInstance2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := MultiMeshInstance2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
[MultiMeshInstance2D] is a specialized node to instance a [MultiMesh] resource in 2D. Usage is the same as [MultiMeshInstance].
*/
type MultiMeshInstance2D struct {
	Node2D
	owner gdnative.Object
}

func (o *MultiMeshInstance2D) BaseClass() string {
	return "MultiMeshInstance2D"
}

/*
        Undocumented
	Args: [], Returns: MultiMesh
*/
func (o *MultiMeshInstance2D) GetMultimesh() MultiMeshImplementer {
	// log.Println("Calling MultiMeshInstance2D.GetMultimesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMeshInstance2D", "get_multimesh")

	// Call the parent method.
	// MultiMesh
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newMultiMeshFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(MultiMeshImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "MultiMesh" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(MultiMeshImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *MultiMeshInstance2D) GetNormalMap() TextureImplementer {
	// log.Println("Calling MultiMeshInstance2D.GetNormalMap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMeshInstance2D", "get_normal_map")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *MultiMeshInstance2D) GetTexture() TextureImplementer {
	// log.Println("Calling MultiMeshInstance2D.GetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMeshInstance2D", "get_texture")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false multimesh MultiMesh}], Returns: void
*/
func (o *MultiMeshInstance2D) SetMultimesh(multimesh MultiMeshImplementer) {
	// log.Println("Calling MultiMeshInstance2D.SetMultimesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(multimesh.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMeshInstance2D", "set_multimesh")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false normal_map Texture}], Returns: void
*/
func (o *MultiMeshInstance2D) SetNormalMap(normalMap TextureImplementer) {
	// log.Println("Calling MultiMeshInstance2D.SetNormalMap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(normalMap.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMeshInstance2D", "set_normal_map")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *MultiMeshInstance2D) SetTexture(texture TextureImplementer) {
	// log.Println("Calling MultiMeshInstance2D.SetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMeshInstance2D", "set_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// MultiMeshInstance2DImplementer is an interface that implements the methods
// of the MultiMeshInstance2D class.
type MultiMeshInstance2DImplementer interface {
	Node2DImplementer
	GetMultimesh() MultiMeshImplementer
	GetNormalMap() TextureImplementer
	GetTexture() TextureImplementer
	SetMultimesh(multimesh MultiMeshImplementer)
	SetNormalMap(normalMap TextureImplementer)
	SetTexture(texture TextureImplementer)
}
