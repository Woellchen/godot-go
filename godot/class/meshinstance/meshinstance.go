package meshinstance

import (
	"log"
	"reflect"

	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
MeshInstance is a node that takes a [Mesh] resource and adds it to the current scenario by creating an instance of it. This is the class most often used to get 3D geometry rendered and can be used to instance a single [Mesh] in many places. This allows to reuse geometry and save on resources. When a [Mesh] has to be instanced more than thousands of times at close proximity, consider using a [MultiMesh] in a [MultiMeshInstance] instead.
*/
type MeshInstance struct {
	GeometryInstance
}

func (o *MeshInstance) BaseClass() string {
	return "MeshInstance"
}

/*
   Undocumented
*/
func (o *MeshInstance) X_MeshChanged() {
	log.Println("Calling MeshInstance.X_MeshChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_mesh_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   This helper creates a [StaticBody] child node with a [ConvexPolygonShape] collision shape calculated from the mesh geometry. It's mainly used for testing.
*/
func (o *MeshInstance) CreateConvexCollision() {
	log.Println("Calling MeshInstance.CreateConvexCollision()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "create_convex_collision", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   This helper creates a [MeshInstance] child node with gizmos at every vertex calculated from the mesh geometry. It's mainly used for testing.
*/
func (o *MeshInstance) CreateDebugTangents() {
	log.Println("Calling MeshInstance.CreateDebugTangents()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "create_debug_tangents", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   This helper creates a [StaticBody] child node with a [ConcavePolygonShape] collision shape calculated from the mesh geometry. It's mainly used for testing.
*/
func (o *MeshInstance) CreateTrimeshCollision() {
	log.Println("Calling MeshInstance.CreateTrimeshCollision()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "create_trimesh_collision", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *MeshInstance) GetMesh() *Mesh {
	log.Println("Calling MeshInstance.GetMesh()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mesh", goArguments, "*Mesh")

	returnValue := goRet.Interface().(*Mesh)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *MeshInstance) GetSkeletonPath() *NodePath {
	log.Println("Calling MeshInstance.GetSkeletonPath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_skeleton_path", goArguments, "*NodePath")

	returnValue := goRet.Interface().(*NodePath)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the [Material] for a surface of the [Mesh] resource.
*/
func (o *MeshInstance) GetSurfaceMaterial(surface gdnative.Int) *Material {
	log.Println("Calling MeshInstance.GetSurfaceMaterial()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(surface)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_surface_material", goArguments, "*Material")

	returnValue := goRet.Interface().(*Material)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *MeshInstance) SetMesh(mesh *Mesh) {
	log.Println("Calling MeshInstance.SetMesh()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mesh)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_mesh", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *MeshInstance) SetSkeletonPath(skeletonPath *NodePath) {
	log.Println("Calling MeshInstance.SetSkeletonPath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(skeletonPath)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_skeleton_path", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the [Material] for a surface of the [Mesh] resource.
*/
func (o *MeshInstance) SetSurfaceMaterial(surface gdnative.Int, material *Material) {
	log.Println("Calling MeshInstance.SetSurfaceMaterial()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(surface)
	goArguments[1] = reflect.ValueOf(material)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_surface_material", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   MeshInstanceImplementer is an interface for MeshInstance objects.
*/
type MeshInstanceImplementer interface {
	Class
}
