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

//func NewImmediateGeometryFromPointer(ptr gdnative.Pointer) ImmediateGeometry {
func newImmediateGeometryFromPointer(ptr gdnative.Pointer) ImmediateGeometry {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ImmediateGeometry{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Draws simple geometry from code. Uses a drawing mode similar to OpenGL 1.x. See also [ArrayMesh], [MeshDataTool] and [SurfaceTool] for procedural geometry generation. [b]Note:[/b] ImmediateGeometry3D is best suited to small amounts of mesh data that change every frame. It will be slow when handling large amounts of mesh data. If mesh data doesn't change often, use [ArrayMesh], [MeshDataTool] or [SurfaceTool] instead. [b]Note:[/b] Godot uses clockwise [url=https://learnopengl.com/Advanced-OpenGL/Face-culling]winding order[/url] for front faces of triangle primitive modes.
*/
type ImmediateGeometry struct {
	GeometryInstance
	owner gdnative.Object
}

func (o *ImmediateGeometry) BaseClass() string {
	return "ImmediateGeometry"
}

/*
        Simple helper to draw an UV sphere with given latitude, longitude and radius.
	Args: [{ false lats int} { false lons int} { false radius float} {True true add_uv bool}], Returns: void
*/
func (o *ImmediateGeometry) AddSphere(lats gdnative.Int, lons gdnative.Int, radius gdnative.Real, addUv gdnative.Bool) {
	//log.Println("Calling ImmediateGeometry.AddSphere()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(lats)
	ptrArguments[1] = gdnative.NewPointerFromInt(lons)
	ptrArguments[2] = gdnative.NewPointerFromReal(radius)
	ptrArguments[3] = gdnative.NewPointerFromBool(addUv)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImmediateGeometry", "add_sphere")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a vertex in local coordinate space with the currently set color/uv/etc.
	Args: [{ false position Vector3}], Returns: void
*/
func (o *ImmediateGeometry) AddVertex(position gdnative.Vector3) {
	//log.Println("Calling ImmediateGeometry.AddVertex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImmediateGeometry", "add_vertex")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Begin drawing (and optionally pass a texture override). When done call [method end]. For more information on how this works, search for [code]glBegin()[/code] and [code]glEnd()[/code] references. For the type of primitive, see the [enum Mesh.PrimitiveType] enum.
	Args: [{ false primitive int} {[Object:null] true texture Texture}], Returns: void
*/
func (o *ImmediateGeometry) Begin(primitive gdnative.Int, texture TextureImplementer) {
	//log.Println("Calling ImmediateGeometry.Begin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(primitive)
	ptrArguments[1] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImmediateGeometry", "begin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clears everything that was drawn using begin/end.
	Args: [], Returns: void
*/
func (o *ImmediateGeometry) Clear() {
	//log.Println("Calling ImmediateGeometry.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImmediateGeometry", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Ends a drawing context and displays the results.
	Args: [], Returns: void
*/
func (o *ImmediateGeometry) End() {
	//log.Println("Calling ImmediateGeometry.End()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImmediateGeometry", "end")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        The current drawing color.
	Args: [{ false color Color}], Returns: void
*/
func (o *ImmediateGeometry) SetColor(color gdnative.Color) {
	//log.Println("Calling ImmediateGeometry.SetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImmediateGeometry", "set_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        The next vertex's normal.
	Args: [{ false normal Vector3}], Returns: void
*/
func (o *ImmediateGeometry) SetNormal(normal gdnative.Vector3) {
	//log.Println("Calling ImmediateGeometry.SetNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(normal)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImmediateGeometry", "set_normal")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        The next vertex's tangent (and binormal facing).
	Args: [{ false tangent Plane}], Returns: void
*/
func (o *ImmediateGeometry) SetTangent(tangent gdnative.Plane) {
	//log.Println("Calling ImmediateGeometry.SetTangent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPlane(tangent)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImmediateGeometry", "set_tangent")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        The next vertex's UV.
	Args: [{ false uv Vector2}], Returns: void
*/
func (o *ImmediateGeometry) SetUv(uv gdnative.Vector2) {
	//log.Println("Calling ImmediateGeometry.SetUv()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(uv)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImmediateGeometry", "set_uv")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        The next vertex's second layer UV.
	Args: [{ false uv Vector2}], Returns: void
*/
func (o *ImmediateGeometry) SetUv2(uv gdnative.Vector2) {
	//log.Println("Calling ImmediateGeometry.SetUv2()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(uv)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImmediateGeometry", "set_uv2")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ImmediateGeometryImplementer is an interface that implements the methods
// of the ImmediateGeometry class.
type ImmediateGeometryImplementer interface {
	GeometryInstanceImplementer
	AddSphere(lats gdnative.Int, lons gdnative.Int, radius gdnative.Real, addUv gdnative.Bool)
	AddVertex(position gdnative.Vector3)
	Begin(primitive gdnative.Int, texture TextureImplementer)
	Clear()
	End()
	SetColor(color gdnative.Color)
	SetNormal(normal gdnative.Vector3)
	SetTangent(tangent gdnative.Plane)
	SetUv(uv gdnative.Vector2)
	SetUv2(uv gdnative.Vector2)
}
