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

//func NewEditorNavigationMeshGeneratorFromPointer(ptr gdnative.Pointer) EditorNavigationMeshGenerator {
func newEditorNavigationMeshGeneratorFromPointer(ptr gdnative.Pointer) EditorNavigationMeshGenerator {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorNavigationMeshGenerator{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type EditorNavigationMeshGenerator struct {
	Object
	owner gdnative.Object
}

func (o *EditorNavigationMeshGenerator) BaseClass() string {
	return "EditorNavigationMeshGenerator"
}

/*

	Args: [{ false nav_mesh NavigationMesh} { false root_node Node}], Returns: void
*/
func (o *EditorNavigationMeshGenerator) Bake(navMesh NavigationMeshImplementer, rootNode NodeImplementer) {
	//log.Println("Calling EditorNavigationMeshGenerator.Bake()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(navMesh.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromObject(rootNode.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorNavigationMeshGenerator", "bake")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false nav_mesh NavigationMesh}], Returns: void
*/
func (o *EditorNavigationMeshGenerator) Clear(navMesh NavigationMeshImplementer) {
	//log.Println("Calling EditorNavigationMeshGenerator.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(navMesh.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorNavigationMeshGenerator", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EditorNavigationMeshGeneratorImplementer is an interface that implements the methods
// of the EditorNavigationMeshGenerator class.
type EditorNavigationMeshGeneratorImplementer interface {
	ObjectImplementer
	Bake(navMesh NavigationMeshImplementer, rootNode NodeImplementer)
	Clear(navMesh NavigationMeshImplementer)
}
