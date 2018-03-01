package canvasitemmaterial

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
[code]CanvasItemMaterial[/code]s provide a means of modifying the textures associated with a CanvasItem. They specialize in describing blend and lighting behaviors for textures. Use a [ShaderMaterial] to more fully customize a material's interactions with a [CanvasItem].
*/
type CanvasItemMaterial struct {
	Material
}

func (o *CanvasItemMaterial) BaseClass() string {
	return "CanvasItemMaterial"
}

/*
   Undocumented
*/
func (o *CanvasItemMaterial) GetBlendMode() gdnative.Int {
	log.Println("Calling CanvasItemMaterial.GetBlendMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_blend_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *CanvasItemMaterial) GetLightMode() gdnative.Int {
	log.Println("Calling CanvasItemMaterial.GetLightMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_light_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *CanvasItemMaterial) SetBlendMode(blendMode gdnative.Int) {
	log.Println("Calling CanvasItemMaterial.SetBlendMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(blendMode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_blend_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *CanvasItemMaterial) SetLightMode(lightMode gdnative.Int) {
	log.Println("Calling CanvasItemMaterial.SetLightMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(lightMode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_light_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   CanvasItemMaterialImplementer is an interface for CanvasItemMaterial objects.
*/
type CanvasItemMaterialImplementer interface {
	Class
}
