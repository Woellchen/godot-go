package light2d

import (
	"log"
	"reflect"
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
Casts light in a 2D environment. Light is defined by a (usually grayscale) texture, a color, an energy value, a mode (see constants), and various other parameters (range and shadows-related). Note that Light2D can be used as a mask.
*/
type Light2D struct {
	Node2D
}

func (o *Light2D) BaseClass() string {
	return "Light2D"
}

/*
   Undocumented
*/
func (o *Light2D) GetColor() *Color {
	log.Println("Calling Light2D.GetColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetEnergy() gdnative.Float {
	log.Println("Calling Light2D.GetEnergy()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_energy", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetHeight() gdnative.Float {
	log.Println("Calling Light2D.GetHeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_height", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetItemCullMask() gdnative.Int {
	log.Println("Calling Light2D.GetItemCullMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_item_cull_mask", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetItemShadowCullMask() gdnative.Int {
	log.Println("Calling Light2D.GetItemShadowCullMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_item_shadow_cull_mask", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetLayerRangeMax() gdnative.Int {
	log.Println("Calling Light2D.GetLayerRangeMax()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_layer_range_max", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetLayerRangeMin() gdnative.Int {
	log.Println("Calling Light2D.GetLayerRangeMin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_layer_range_min", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetMode() gdnative.Int {
	log.Println("Calling Light2D.GetMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetShadowBufferSize() gdnative.Int {
	log.Println("Calling Light2D.GetShadowBufferSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_shadow_buffer_size", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetShadowColor() *Color {
	log.Println("Calling Light2D.GetShadowColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_shadow_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetShadowFilter() gdnative.Int {
	log.Println("Calling Light2D.GetShadowFilter()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_shadow_filter", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetShadowGradientLength() gdnative.Float {
	log.Println("Calling Light2D.GetShadowGradientLength()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_shadow_gradient_length", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetShadowSmooth() gdnative.Float {
	log.Println("Calling Light2D.GetShadowSmooth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_shadow_smooth", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetTexture() *Texture {
	log.Println("Calling Light2D.GetTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_texture", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetTextureOffset() *Vector2 {
	log.Println("Calling Light2D.GetTextureOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_texture_offset", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetTextureScale() gdnative.Float {
	log.Println("Calling Light2D.GetTextureScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_texture_scale", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetZRangeMax() gdnative.Int {
	log.Println("Calling Light2D.GetZRangeMax()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_z_range_max", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) GetZRangeMin() gdnative.Int {
	log.Println("Calling Light2D.GetZRangeMin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_z_range_min", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) IsEditorOnly() gdnative.Bool {
	log.Println("Calling Light2D.IsEditorOnly()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_editor_only", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) IsEnabled() gdnative.Bool {
	log.Println("Calling Light2D.IsEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) IsShadowEnabled() gdnative.Bool {
	log.Println("Calling Light2D.IsShadowEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_shadow_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Light2D) SetColor(color *Color) {
	log.Println("Calling Light2D.SetColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetEditorOnly(editorOnly gdnative.Bool) {
	log.Println("Calling Light2D.SetEditorOnly()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(editorOnly)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_editor_only", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetEnabled(enabled gdnative.Bool) {
	log.Println("Calling Light2D.SetEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_enabled", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetEnergy(energy gdnative.Float) {
	log.Println("Calling Light2D.SetEnergy()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(energy)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_energy", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetHeight(height gdnative.Float) {
	log.Println("Calling Light2D.SetHeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(height)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_height", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetItemCullMask(itemCullMask gdnative.Int) {
	log.Println("Calling Light2D.SetItemCullMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(itemCullMask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_item_cull_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetItemShadowCullMask(itemShadowCullMask gdnative.Int) {
	log.Println("Calling Light2D.SetItemShadowCullMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(itemShadowCullMask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_item_shadow_cull_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetLayerRangeMax(layer gdnative.Int) {
	log.Println("Calling Light2D.SetLayerRangeMax()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(layer)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_layer_range_max", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetLayerRangeMin(layer gdnative.Int) {
	log.Println("Calling Light2D.SetLayerRangeMin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(layer)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_layer_range_min", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetMode(mode gdnative.Int) {
	log.Println("Calling Light2D.SetMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetShadowBufferSize(size gdnative.Int) {
	log.Println("Calling Light2D.SetShadowBufferSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_shadow_buffer_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetShadowColor(shadowColor *Color) {
	log.Println("Calling Light2D.SetShadowColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(shadowColor)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_shadow_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetShadowEnabled(enabled gdnative.Bool) {
	log.Println("Calling Light2D.SetShadowEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_shadow_enabled", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetShadowFilter(filter gdnative.Int) {
	log.Println("Calling Light2D.SetShadowFilter()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(filter)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_shadow_filter", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetShadowGradientLength(multiplier gdnative.Float) {
	log.Println("Calling Light2D.SetShadowGradientLength()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(multiplier)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_shadow_gradient_length", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetShadowSmooth(smooth gdnative.Float) {
	log.Println("Calling Light2D.SetShadowSmooth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(smooth)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_shadow_smooth", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetTexture(texture *Texture) {
	log.Println("Calling Light2D.SetTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_texture", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetTextureOffset(textureOffset *Vector2) {
	log.Println("Calling Light2D.SetTextureOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(textureOffset)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_texture_offset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetTextureScale(textureScale gdnative.Float) {
	log.Println("Calling Light2D.SetTextureScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(textureScale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_texture_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetZRangeMax(z gdnative.Int) {
	log.Println("Calling Light2D.SetZRangeMax()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(z)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_z_range_max", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Light2D) SetZRangeMin(z gdnative.Int) {
	log.Println("Calling Light2D.SetZRangeMin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(z)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_z_range_min", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Light2DImplementer is an interface for Light2D objects.
*/
type Light2DImplementer interface {
	Class
}
