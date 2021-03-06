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

// ProceduralSkyTextureSize is an enum for TextureSize values.
type ProceduralSkyTextureSize int

const (
	ProceduralSkyTextureSize1024 ProceduralSkyTextureSize = 2
	ProceduralSkyTextureSize2048 ProceduralSkyTextureSize = 3
	ProceduralSkyTextureSize256  ProceduralSkyTextureSize = 0
	ProceduralSkyTextureSize4096 ProceduralSkyTextureSize = 4
	ProceduralSkyTextureSize512  ProceduralSkyTextureSize = 1
	ProceduralSkyTextureSizeMax  ProceduralSkyTextureSize = 5
)

//func NewProceduralSkyFromPointer(ptr gdnative.Pointer) ProceduralSky {
func newProceduralSkyFromPointer(ptr gdnative.Pointer) ProceduralSky {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ProceduralSky{}
	obj.SetBaseObject(owner)

	return obj
}

/*
ProceduralSky provides a way to create an effective background quickly by defining procedural parameters for the sun, the sky and the ground. The sky and ground are very similar, they are defined by a color at the horizon, another color, and finally an easing curve to interpolate between these two colors. Similarly, the sun is described by a position in the sky, a color, and an easing curve. However, the sun also defines a minimum and maximum angle, these two values define at what distance the easing curve begins and ends from the sun, and thus end up defining the size of the sun in the sky. The ProceduralSky is updated on the CPU after the parameters change. It is stored in a texture and then displayed as a background in the scene. This makes it relatively unsuitable for real-time updates during gameplay. However, with a small enough texture size, it can still be updated relatively frequently, as it is updated on a background thread when multi-threading is available.
*/
type ProceduralSky struct {
	Sky
	owner gdnative.Object
}

func (o *ProceduralSky) BaseClass() string {
	return "ProceduralSky"
}

/*
        Undocumented
	Args: [{ false image Image}], Returns: void
*/
func (o *ProceduralSky) X_ThreadDone(image ImageImplementer) {
	// log.Println("Calling ProceduralSky.X_ThreadDone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(image.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "_thread_done")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ProceduralSky) X_UpdateSky() {
	// log.Println("Calling ProceduralSky.X_UpdateSky()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "_update_sky")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *ProceduralSky) GetGroundBottomColor() gdnative.Color {
	// log.Println("Calling ProceduralSky.GetGroundBottomColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_ground_bottom_color")

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
	Args: [], Returns: float
*/
func (o *ProceduralSky) GetGroundCurve() gdnative.Real {
	// log.Println("Calling ProceduralSky.GetGroundCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_ground_curve")

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
	Args: [], Returns: float
*/
func (o *ProceduralSky) GetGroundEnergy() gdnative.Real {
	// log.Println("Calling ProceduralSky.GetGroundEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_ground_energy")

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
	Args: [], Returns: Color
*/
func (o *ProceduralSky) GetGroundHorizonColor() gdnative.Color {
	// log.Println("Calling ProceduralSky.GetGroundHorizonColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_ground_horizon_color")

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
	Args: [], Returns: float
*/
func (o *ProceduralSky) GetSkyCurve() gdnative.Real {
	// log.Println("Calling ProceduralSky.GetSkyCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sky_curve")

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
	Args: [], Returns: float
*/
func (o *ProceduralSky) GetSkyEnergy() gdnative.Real {
	// log.Println("Calling ProceduralSky.GetSkyEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sky_energy")

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
	Args: [], Returns: Color
*/
func (o *ProceduralSky) GetSkyHorizonColor() gdnative.Color {
	// log.Println("Calling ProceduralSky.GetSkyHorizonColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sky_horizon_color")

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
	Args: [], Returns: Color
*/
func (o *ProceduralSky) GetSkyTopColor() gdnative.Color {
	// log.Println("Calling ProceduralSky.GetSkyTopColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sky_top_color")

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
	Args: [], Returns: float
*/
func (o *ProceduralSky) GetSunAngleMax() gdnative.Real {
	// log.Println("Calling ProceduralSky.GetSunAngleMax()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_angle_max")

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
	Args: [], Returns: float
*/
func (o *ProceduralSky) GetSunAngleMin() gdnative.Real {
	// log.Println("Calling ProceduralSky.GetSunAngleMin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_angle_min")

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
	Args: [], Returns: Color
*/
func (o *ProceduralSky) GetSunColor() gdnative.Color {
	// log.Println("Calling ProceduralSky.GetSunColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_color")

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
	Args: [], Returns: float
*/
func (o *ProceduralSky) GetSunCurve() gdnative.Real {
	// log.Println("Calling ProceduralSky.GetSunCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_curve")

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
	Args: [], Returns: float
*/
func (o *ProceduralSky) GetSunEnergy() gdnative.Real {
	// log.Println("Calling ProceduralSky.GetSunEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_energy")

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
	Args: [], Returns: float
*/
func (o *ProceduralSky) GetSunLatitude() gdnative.Real {
	// log.Println("Calling ProceduralSky.GetSunLatitude()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_latitude")

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
	Args: [], Returns: float
*/
func (o *ProceduralSky) GetSunLongitude() gdnative.Real {
	// log.Println("Calling ProceduralSky.GetSunLongitude()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_longitude")

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
	Args: [], Returns: enum.ProceduralSky::TextureSize
*/
func (o *ProceduralSky) GetTextureSize() ProceduralSkyTextureSize {
	// log.Println("Calling ProceduralSky.GetTextureSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_texture_size")

	// Call the parent method.
	// enum.ProceduralSky::TextureSize
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ProceduralSkyTextureSize(ret)
}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ProceduralSky) SetGroundBottomColor(color gdnative.Color) {
	// log.Println("Calling ProceduralSky.SetGroundBottomColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_ground_bottom_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false curve float}], Returns: void
*/
func (o *ProceduralSky) SetGroundCurve(curve gdnative.Real) {
	// log.Println("Calling ProceduralSky.SetGroundCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(curve)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_ground_curve")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false energy float}], Returns: void
*/
func (o *ProceduralSky) SetGroundEnergy(energy gdnative.Real) {
	// log.Println("Calling ProceduralSky.SetGroundEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(energy)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_ground_energy")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ProceduralSky) SetGroundHorizonColor(color gdnative.Color) {
	// log.Println("Calling ProceduralSky.SetGroundHorizonColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_ground_horizon_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false curve float}], Returns: void
*/
func (o *ProceduralSky) SetSkyCurve(curve gdnative.Real) {
	// log.Println("Calling ProceduralSky.SetSkyCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(curve)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sky_curve")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false energy float}], Returns: void
*/
func (o *ProceduralSky) SetSkyEnergy(energy gdnative.Real) {
	// log.Println("Calling ProceduralSky.SetSkyEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(energy)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sky_energy")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ProceduralSky) SetSkyHorizonColor(color gdnative.Color) {
	// log.Println("Calling ProceduralSky.SetSkyHorizonColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sky_horizon_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ProceduralSky) SetSkyTopColor(color gdnative.Color) {
	// log.Println("Calling ProceduralSky.SetSkyTopColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sky_top_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *ProceduralSky) SetSunAngleMax(degrees gdnative.Real) {
	// log.Println("Calling ProceduralSky.SetSunAngleMax()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_angle_max")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *ProceduralSky) SetSunAngleMin(degrees gdnative.Real) {
	// log.Println("Calling ProceduralSky.SetSunAngleMin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_angle_min")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ProceduralSky) SetSunColor(color gdnative.Color) {
	// log.Println("Calling ProceduralSky.SetSunColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false curve float}], Returns: void
*/
func (o *ProceduralSky) SetSunCurve(curve gdnative.Real) {
	// log.Println("Calling ProceduralSky.SetSunCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(curve)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_curve")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false energy float}], Returns: void
*/
func (o *ProceduralSky) SetSunEnergy(energy gdnative.Real) {
	// log.Println("Calling ProceduralSky.SetSunEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(energy)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_energy")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *ProceduralSky) SetSunLatitude(degrees gdnative.Real) {
	// log.Println("Calling ProceduralSky.SetSunLatitude()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_latitude")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *ProceduralSky) SetSunLongitude(degrees gdnative.Real) {
	// log.Println("Calling ProceduralSky.SetSunLongitude()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_longitude")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false size int}], Returns: void
*/
func (o *ProceduralSky) SetTextureSize(size gdnative.Int) {
	// log.Println("Calling ProceduralSky.SetTextureSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_texture_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ProceduralSkyImplementer is an interface that implements the methods
// of the ProceduralSky class.
type ProceduralSkyImplementer interface {
	SkyImplementer
	X_ThreadDone(image ImageImplementer)
	X_UpdateSky()
	GetGroundBottomColor() gdnative.Color
	GetGroundCurve() gdnative.Real
	GetGroundEnergy() gdnative.Real
	GetGroundHorizonColor() gdnative.Color
	GetSkyCurve() gdnative.Real
	GetSkyEnergy() gdnative.Real
	GetSkyHorizonColor() gdnative.Color
	GetSkyTopColor() gdnative.Color
	GetSunAngleMax() gdnative.Real
	GetSunAngleMin() gdnative.Real
	GetSunColor() gdnative.Color
	GetSunCurve() gdnative.Real
	GetSunEnergy() gdnative.Real
	GetSunLatitude() gdnative.Real
	GetSunLongitude() gdnative.Real
	SetGroundBottomColor(color gdnative.Color)
	SetGroundCurve(curve gdnative.Real)
	SetGroundEnergy(energy gdnative.Real)
	SetGroundHorizonColor(color gdnative.Color)
	SetSkyCurve(curve gdnative.Real)
	SetSkyEnergy(energy gdnative.Real)
	SetSkyHorizonColor(color gdnative.Color)
	SetSkyTopColor(color gdnative.Color)
	SetSunAngleMax(degrees gdnative.Real)
	SetSunAngleMin(degrees gdnative.Real)
	SetSunColor(color gdnative.Color)
	SetSunCurve(curve gdnative.Real)
	SetSunEnergy(energy gdnative.Real)
	SetSunLatitude(degrees gdnative.Real)
	SetSunLongitude(degrees gdnative.Real)
	SetTextureSize(size gdnative.Int)
}
