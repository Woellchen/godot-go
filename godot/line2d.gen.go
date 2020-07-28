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

// Line2DLineCapMode is an enum for LineCapMode values.
type Line2DLineCapMode int

const (
	Line2DLineCapBox   Line2DLineCapMode = 1
	Line2DLineCapNone  Line2DLineCapMode = 0
	Line2DLineCapRound Line2DLineCapMode = 2
)

// Line2DLineJointMode is an enum for LineJointMode values.
type Line2DLineJointMode int

const (
	Line2DLineJointBevel Line2DLineJointMode = 1
	Line2DLineJointRound Line2DLineJointMode = 2
	Line2DLineJointSharp Line2DLineJointMode = 0
)

// Line2DLineTextureMode is an enum for LineTextureMode values.
type Line2DLineTextureMode int

const (
	Line2DLineTextureNone    Line2DLineTextureMode = 0
	Line2DLineTextureStretch Line2DLineTextureMode = 2
	Line2DLineTextureTile    Line2DLineTextureMode = 1
)

//func NewLine2DFromPointer(ptr gdnative.Pointer) Line2D {
func newLine2DFromPointer(ptr gdnative.Pointer) Line2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Line2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A line through several points in 2D space. [b]Note:[/b] By default, Godot can only draw up to 4,096 polygon points at a time. To increase this limit, open the Project Settings and increase [member ProjectSettings.rendering/limits/buffers/canvas_polygon_buffer_size_kb] and [member ProjectSettings.rendering/limits/buffers/canvas_polygon_index_buffer_size_kb].
*/
type Line2D struct {
	Node2D
	owner gdnative.Object
}

func (o *Line2D) BaseClass() string {
	return "Line2D"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Line2D) X_CurveChanged() {
	// log.Println("Calling Line2D.X_CurveChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "_curve_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Line2D) X_GradientChanged() {
	// log.Println("Calling Line2D.X_GradientChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "_gradient_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a point at the [code]position[/code]. Appends the point at the end of the line. If [code]at_position[/code] is given, the point is inserted before the point number [code]at_position[/code], moving that point (and every point after) after the inserted point. If [code]at_position[/code] is not given, or is an illegal value ([code]at_position < 0[/code] or [code]at_position >= [method get_point_count][/code]), the point will be appended at the end of the point list.
	Args: [{ false position Vector2} {-1 true at_position int}], Returns: void
*/
func (o *Line2D) AddPoint(position gdnative.Vector2, atPosition gdnative.Int) {
	// log.Println("Calling Line2D.AddPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector2(position)
	ptrArguments[1] = gdnative.NewPointerFromInt(atPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "add_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all points from the line.
	Args: [], Returns: void
*/
func (o *Line2D) ClearPoints() {
	// log.Println("Calling Line2D.ClearPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "clear_points")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Line2D) GetAntialiased() gdnative.Bool {
	// log.Println("Calling Line2D.GetAntialiased()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_antialiased")

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
	Args: [], Returns: enum.Line2D::LineCapMode
*/
func (o *Line2D) GetBeginCapMode() Line2DLineCapMode {
	// log.Println("Calling Line2D.GetBeginCapMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_begin_cap_mode")

	// Call the parent method.
	// enum.Line2D::LineCapMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return Line2DLineCapMode(ret)
}

/*
        Undocumented
	Args: [], Returns: Curve
*/
func (o *Line2D) GetCurve() CurveImplementer {
	// log.Println("Calling Line2D.GetCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_curve")

	// Call the parent method.
	// Curve
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newCurveFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(CurveImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Curve" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(CurveImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *Line2D) GetDefaultColor() gdnative.Color {
	// log.Println("Calling Line2D.GetDefaultColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_default_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.Line2D::LineCapMode
*/
func (o *Line2D) GetEndCapMode() Line2DLineCapMode {
	// log.Println("Calling Line2D.GetEndCapMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_end_cap_mode")

	// Call the parent method.
	// enum.Line2D::LineCapMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return Line2DLineCapMode(ret)
}

/*
        Undocumented
	Args: [], Returns: Gradient
*/
func (o *Line2D) GetGradient() GradientImplementer {
	// log.Println("Calling Line2D.GetGradient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_gradient")

	// Call the parent method.
	// Gradient
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newGradientFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(GradientImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Gradient" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(GradientImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: enum.Line2D::LineJointMode
*/
func (o *Line2D) GetJointMode() Line2DLineJointMode {
	// log.Println("Calling Line2D.GetJointMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_joint_mode")

	// Call the parent method.
	// enum.Line2D::LineJointMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return Line2DLineJointMode(ret)
}

/*
        Returns the Line2D's amount of points.
	Args: [], Returns: int
*/
func (o *Line2D) GetPointCount() gdnative.Int {
	// log.Println("Calling Line2D.GetPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_point_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns point [code]i[/code]'s position.
	Args: [{ false i int}], Returns: Vector2
*/
func (o *Line2D) GetPointPosition(i gdnative.Int) gdnative.Vector2 {
	// log.Println("Calling Line2D.GetPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(i)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_point_position")

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
	Args: [], Returns: PoolVector2Array
*/
func (o *Line2D) GetPoints() gdnative.PoolVector2Array {
	// log.Println("Calling Line2D.GetPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_points")

	// Call the parent method.
	// PoolVector2Array
	retPtr := gdnative.NewEmptyPoolVector2Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector2ArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Line2D) GetRoundPrecision() gdnative.Int {
	// log.Println("Calling Line2D.GetRoundPrecision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_round_precision")

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
	Args: [], Returns: float
*/
func (o *Line2D) GetSharpLimit() gdnative.Real {
	// log.Println("Calling Line2D.GetSharpLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_sharp_limit")

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
	Args: [], Returns: Texture
*/
func (o *Line2D) GetTexture() TextureImplementer {
	// log.Println("Calling Line2D.GetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_texture")

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
	Args: [], Returns: enum.Line2D::LineTextureMode
*/
func (o *Line2D) GetTextureMode() Line2DLineTextureMode {
	// log.Println("Calling Line2D.GetTextureMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_texture_mode")

	// Call the parent method.
	// enum.Line2D::LineTextureMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return Line2DLineTextureMode(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Line2D) GetWidth() gdnative.Real {
	// log.Println("Calling Line2D.GetWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "get_width")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Removes the point at index [code]i[/code] from the line.
	Args: [{ false i int}], Returns: void
*/
func (o *Line2D) RemovePoint(i gdnative.Int) {
	// log.Println("Calling Line2D.RemovePoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(i)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "remove_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false antialiased bool}], Returns: void
*/
func (o *Line2D) SetAntialiased(antialiased gdnative.Bool) {
	// log.Println("Calling Line2D.SetAntialiased()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(antialiased)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_antialiased")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *Line2D) SetBeginCapMode(mode gdnative.Int) {
	// log.Println("Calling Line2D.SetBeginCapMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_begin_cap_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false curve Curve}], Returns: void
*/
func (o *Line2D) SetCurve(curve CurveImplementer) {
	// log.Println("Calling Line2D.SetCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(curve.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_curve")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *Line2D) SetDefaultColor(color gdnative.Color) {
	// log.Println("Calling Line2D.SetDefaultColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_default_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *Line2D) SetEndCapMode(mode gdnative.Int) {
	// log.Println("Calling Line2D.SetEndCapMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_end_cap_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Gradient}], Returns: void
*/
func (o *Line2D) SetGradient(color GradientImplementer) {
	// log.Println("Calling Line2D.SetGradient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(color.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_gradient")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *Line2D) SetJointMode(mode gdnative.Int) {
	// log.Println("Calling Line2D.SetJointMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_joint_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Overwrites the position in point [code]i[/code] with the supplied [code]position[/code].
	Args: [{ false i int} { false position Vector2}], Returns: void
*/
func (o *Line2D) SetPointPosition(i gdnative.Int, position gdnative.Vector2) {
	// log.Println("Calling Line2D.SetPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(i)
	ptrArguments[1] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_point_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false points PoolVector2Array}], Returns: void
*/
func (o *Line2D) SetPoints(points gdnative.PoolVector2Array) {
	// log.Println("Calling Line2D.SetPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector2Array(points)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_points")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false precision int}], Returns: void
*/
func (o *Line2D) SetRoundPrecision(precision gdnative.Int) {
	// log.Println("Calling Line2D.SetRoundPrecision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(precision)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_round_precision")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false limit float}], Returns: void
*/
func (o *Line2D) SetSharpLimit(limit gdnative.Real) {
	// log.Println("Calling Line2D.SetSharpLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(limit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_sharp_limit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *Line2D) SetTexture(texture TextureImplementer) {
	// log.Println("Calling Line2D.SetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *Line2D) SetTextureMode(mode gdnative.Int) {
	// log.Println("Calling Line2D.SetTextureMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_texture_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false width float}], Returns: void
*/
func (o *Line2D) SetWidth(width gdnative.Real) {
	// log.Println("Calling Line2D.SetWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(width)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Line2D", "set_width")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// Line2DImplementer is an interface that implements the methods
// of the Line2D class.
type Line2DImplementer interface {
	Node2DImplementer
	X_CurveChanged()
	X_GradientChanged()
	AddPoint(position gdnative.Vector2, atPosition gdnative.Int)
	ClearPoints()
	GetAntialiased() gdnative.Bool
	GetCurve() CurveImplementer
	GetDefaultColor() gdnative.Color
	GetGradient() GradientImplementer
	GetPointCount() gdnative.Int
	GetPointPosition(i gdnative.Int) gdnative.Vector2
	GetPoints() gdnative.PoolVector2Array
	GetRoundPrecision() gdnative.Int
	GetSharpLimit() gdnative.Real
	GetTexture() TextureImplementer
	GetWidth() gdnative.Real
	RemovePoint(i gdnative.Int)
	SetAntialiased(antialiased gdnative.Bool)
	SetBeginCapMode(mode gdnative.Int)
	SetCurve(curve CurveImplementer)
	SetDefaultColor(color gdnative.Color)
	SetEndCapMode(mode gdnative.Int)
	SetGradient(color GradientImplementer)
	SetJointMode(mode gdnative.Int)
	SetPointPosition(i gdnative.Int, position gdnative.Vector2)
	SetPoints(points gdnative.PoolVector2Array)
	SetRoundPrecision(precision gdnative.Int)
	SetSharpLimit(limit gdnative.Real)
	SetTexture(texture TextureImplementer)
	SetTextureMode(mode gdnative.Int)
	SetWidth(width gdnative.Real)
}
