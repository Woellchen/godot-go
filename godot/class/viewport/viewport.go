package viewport

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
A Viewport creates a different view into the screen, or a sub-view inside another viewport. Children 2D Nodes will display on it, and children Camera 3D nodes will render on it too. Optionally, a viewport can have its own 2D or 3D world, so they don't share what they draw with other viewports. If a viewport is a child of a [Control], it will automatically take up its same rect and position, otherwise they must be set manually. Viewports can also choose to be audio listeners, so they generate positional audio depending on a 2D or 3D camera child of it. Also, viewports can be assigned to different screens in case the devices have multiple screens. Finally, viewports can also behave as render targets, in which case they will not be visible unless the associated texture is used to draw.
*/
type Viewport struct {
	Node
}

func (o *Viewport) BaseClass() string {
	return "Viewport"
}

/*
   Undocumented
*/
func (o *Viewport) X_GuiRemoveFocus() {
	log.Println("Calling Viewport.X_GuiRemoveFocus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_gui_remove_focus", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) X_GuiShowTooltip() {
	log.Println("Calling Viewport.X_GuiShowTooltip()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_gui_show_tooltip", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) X_VpInput(arg0 *InputEvent) {
	log.Println("Calling Viewport.X_VpInput()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_vp_input", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) X_VpInputText(text gdnative.String) {
	log.Println("Calling Viewport.X_VpInputText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(text)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_vp_input_text", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) X_VpUnhandledInput(arg0 *InputEvent) {
	log.Println("Calling Viewport.X_VpUnhandledInput()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_vp_unhandled_input", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the 3D world of the viewport, or if no such present, the one of the parent viewport.
*/
func (o *Viewport) FindWorld() *World {
	log.Println("Calling Viewport.FindWorld()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "find_world", goArguments, "*World")

	returnValue := goRet.Interface().(*World)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the 2D world of the viewport.
*/
func (o *Viewport) FindWorld2D() *World2D {
	log.Println("Calling Viewport.FindWorld2D()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "find_world_2d", goArguments, "*World2D")

	returnValue := goRet.Interface().(*World2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the active 3D camera.
*/
func (o *Viewport) GetCamera() *Camera {
	log.Println("Calling Viewport.GetCamera()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_camera", goArguments, "*Camera")

	returnValue := goRet.Interface().(*Camera)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetCanvasTransform() *Transform2D {
	log.Println("Calling Viewport.GetCanvasTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_canvas_transform", goArguments, "*Transform2D")

	returnValue := goRet.Interface().(*Transform2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetClearMode() gdnative.Int {
	log.Println("Calling Viewport.GetClearMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_clear_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetDebugDraw() gdnative.Int {
	log.Println("Calling Viewport.GetDebugDraw()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_debug_draw", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the total transform of the viewport.
*/
func (o *Viewport) GetFinalTransform() *Transform2D {
	log.Println("Calling Viewport.GetFinalTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_final_transform", goArguments, "*Transform2D")

	returnValue := goRet.Interface().(*Transform2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetGlobalCanvasTransform() *Transform2D {
	log.Println("Calling Viewport.GetGlobalCanvasTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_global_canvas_transform", goArguments, "*Transform2D")

	returnValue := goRet.Interface().(*Transform2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetHdr() gdnative.Bool {
	log.Println("Calling Viewport.GetHdr()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_hdr", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the mouse position, relative to the viewport.
*/
func (o *Viewport) GetMousePosition() *Vector2 {
	log.Println("Calling Viewport.GetMousePosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mouse_position", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetMsaa() gdnative.Int {
	log.Println("Calling Viewport.GetMsaa()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_msaa", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetPhysicsObjectPicking() gdnative.Bool {
	log.Println("Calling Viewport.GetPhysicsObjectPicking()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_physics_object_picking", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the specific information about the viewport from rendering pipeline.
*/
func (o *Viewport) GetRenderInfo(info gdnative.Int) gdnative.Int {
	log.Println("Calling Viewport.GetRenderInfo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(info)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_render_info", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetShadowAtlasQuadrantSubdiv(quadrant gdnative.Int) gdnative.Int {
	log.Println("Calling Viewport.GetShadowAtlasQuadrantSubdiv()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(quadrant)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_shadow_atlas_quadrant_subdiv", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetShadowAtlasSize() gdnative.Int {
	log.Println("Calling Viewport.GetShadowAtlasSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_shadow_atlas_size", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetSize() *Vector2 {
	log.Println("Calling Viewport.GetSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_size", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the size override set with [method set_size_override].
*/
func (o *Viewport) GetSizeOverride() *Vector2 {
	log.Println("Calling Viewport.GetSizeOverride()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_size_override", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the viewport's texture, for use with various objects that you want to texture with the viewport.
*/
func (o *Viewport) GetTexture() *ViewportTexture {
	log.Println("Calling Viewport.GetTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_texture", goArguments, "*ViewportTexture")

	returnValue := goRet.Interface().(*ViewportTexture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetUpdateMode() gdnative.Int {
	log.Println("Calling Viewport.GetUpdateMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_update_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetUsage() gdnative.Int {
	log.Println("Calling Viewport.GetUsage()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_usage", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetVflip() gdnative.Bool {
	log.Println("Calling Viewport.GetVflip()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_vflip", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the viewport RID from the [VisualServer].
*/
func (o *Viewport) GetViewportRid() *RID {
	log.Println("Calling Viewport.GetViewportRid()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_viewport_rid", goArguments, "*RID")

	returnValue := goRet.Interface().(*RID)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the final, visible rect in global screen coordinates.
*/
func (o *Viewport) GetVisibleRect() *Rect2 {
	log.Println("Calling Viewport.GetVisibleRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_visible_rect", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetWorld() *World {
	log.Println("Calling Viewport.GetWorld()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_world", goArguments, "*World")

	returnValue := goRet.Interface().(*World)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) GetWorld2D() *World2D {
	log.Println("Calling Viewport.GetWorld2D()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_world_2d", goArguments, "*World2D")

	returnValue := goRet.Interface().(*World2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the drag data from the GUI, that was previously returned by [method Control.get_drag_data].
*/
func (o *Viewport) GuiGetDragData() *Variant {
	log.Println("Calling Viewport.GuiGetDragData()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "gui_get_drag_data", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns whether there are shown modals on-screen.
*/
func (o *Viewport) GuiHasModalStack() gdnative.Bool {
	log.Println("Calling Viewport.GuiHasModalStack()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "gui_has_modal_stack", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) HasTransparentBackground() gdnative.Bool {
	log.Println("Calling Viewport.HasTransparentBackground()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "has_transparent_background", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *Viewport) Input(localEvent *InputEvent) {
	log.Println("Calling Viewport.Input()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(localEvent)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "input", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) Is3DDisabled() gdnative.Bool {
	log.Println("Calling Viewport.Is3DDisabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_3d_disabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) IsAudioListener() gdnative.Bool {
	log.Println("Calling Viewport.IsAudioListener()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_audio_listener", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) IsAudioListener2D() gdnative.Bool {
	log.Println("Calling Viewport.IsAudioListener2D()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_audio_listener_2d", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) IsInputDisabled() gdnative.Bool {
	log.Println("Calling Viewport.IsInputDisabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_input_disabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the enabled status of the size override set with [method set_size_override].
*/
func (o *Viewport) IsSizeOverrideEnabled() gdnative.Bool {
	log.Println("Calling Viewport.IsSizeOverrideEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_size_override_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the enabled status of the size stretch override set with [method set_size_override_stretch].
*/
func (o *Viewport) IsSizeOverrideStretchEnabled() gdnative.Bool {
	log.Println("Calling Viewport.IsSizeOverrideStretchEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_size_override_stretch_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) IsSnapControlsToPixelsEnabled() gdnative.Bool {
	log.Println("Calling Viewport.IsSnapControlsToPixelsEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_snap_controls_to_pixels_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) IsUsingOwnWorld() gdnative.Bool {
	log.Println("Calling Viewport.IsUsingOwnWorld()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_using_own_world", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Viewport) SetAsAudioListener(enable gdnative.Bool) {
	log.Println("Calling Viewport.SetAsAudioListener()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_as_audio_listener", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetAsAudioListener2D(enable gdnative.Bool) {
	log.Println("Calling Viewport.SetAsAudioListener2D()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_as_audio_listener_2d", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *Viewport) SetAttachToScreenRect(rect *Rect2) {
	log.Println("Calling Viewport.SetAttachToScreenRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rect)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_attach_to_screen_rect", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetCanvasTransform(xform *Transform2D) {
	log.Println("Calling Viewport.SetCanvasTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(xform)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_canvas_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetClearMode(mode gdnative.Int) {
	log.Println("Calling Viewport.SetClearMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_clear_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetDebugDraw(debugDraw gdnative.Int) {
	log.Println("Calling Viewport.SetDebugDraw()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(debugDraw)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_debug_draw", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetDisable3D(disable gdnative.Bool) {
	log.Println("Calling Viewport.SetDisable3D()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(disable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_disable_3d", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetDisableInput(disable gdnative.Bool) {
	log.Println("Calling Viewport.SetDisableInput()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(disable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_disable_input", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetGlobalCanvasTransform(xform *Transform2D) {
	log.Println("Calling Viewport.SetGlobalCanvasTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(xform)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_global_canvas_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetHdr(enable gdnative.Bool) {
	log.Println("Calling Viewport.SetHdr()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_hdr", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetMsaa(msaa gdnative.Int) {
	log.Println("Calling Viewport.SetMsaa()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(msaa)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_msaa", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetPhysicsObjectPicking(enable gdnative.Bool) {
	log.Println("Calling Viewport.SetPhysicsObjectPicking()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_physics_object_picking", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetShadowAtlasQuadrantSubdiv(quadrant gdnative.Int, subdiv gdnative.Int) {
	log.Println("Calling Viewport.SetShadowAtlasQuadrantSubdiv()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(quadrant)
	goArguments[1] = reflect.ValueOf(subdiv)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_shadow_atlas_quadrant_subdiv", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetShadowAtlasSize(size gdnative.Int) {
	log.Println("Calling Viewport.SetShadowAtlasSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_shadow_atlas_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetSize(size *Vector2) {
	log.Println("Calling Viewport.SetSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the size override of the viewport. If the enable parameter is true, it would use the override, otherwise it would use the default size. If the size parameter is equal to [code](-1, -1)[/code], it won't update the size.
*/
func (o *Viewport) SetSizeOverride(enable gdnative.Bool, size *Vector2, margin *Vector2) {
	log.Println("Calling Viewport.SetSizeOverride()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(enable)
	goArguments[1] = reflect.ValueOf(size)
	goArguments[2] = reflect.ValueOf(margin)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_size_override", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether the size override affects stretch as well.
*/
func (o *Viewport) SetSizeOverrideStretch(enabled gdnative.Bool) {
	log.Println("Calling Viewport.SetSizeOverrideStretch()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_size_override_stretch", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetSnapControlsToPixels(enabled gdnative.Bool) {
	log.Println("Calling Viewport.SetSnapControlsToPixels()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_snap_controls_to_pixels", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetTransparentBackground(enable gdnative.Bool) {
	log.Println("Calling Viewport.SetTransparentBackground()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_transparent_background", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetUpdateMode(mode gdnative.Int) {
	log.Println("Calling Viewport.SetUpdateMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_update_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetUsage(usage gdnative.Int) {
	log.Println("Calling Viewport.SetUsage()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(usage)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_usage", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetUseArvr(use gdnative.Bool) {
	log.Println("Calling Viewport.SetUseArvr()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(use)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_use_arvr", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetUseOwnWorld(enable gdnative.Bool) {
	log.Println("Calling Viewport.SetUseOwnWorld()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_use_own_world", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetVflip(enable gdnative.Bool) {
	log.Println("Calling Viewport.SetVflip()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_vflip", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetWorld(world *World) {
	log.Println("Calling Viewport.SetWorld()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(world)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_world", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) SetWorld2D(world2D *World2D) {
	log.Println("Calling Viewport.SetWorld2D()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(world2D)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_world_2d", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *Viewport) UnhandledInput(localEvent *InputEvent) {
	log.Println("Calling Viewport.UnhandledInput()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(localEvent)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "unhandled_input", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Force update of the 2D and 3D worlds.
*/
func (o *Viewport) UpdateWorlds() {
	log.Println("Calling Viewport.UpdateWorlds()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "update_worlds", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Viewport) UseArvr() gdnative.Bool {
	log.Println("Calling Viewport.UseArvr()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "use_arvr", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Warp the mouse to a position, relative to the viewport.
*/
func (o *Viewport) WarpMouse(toPosition *Vector2) {
	log.Println("Calling Viewport.WarpMouse()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(toPosition)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "warp_mouse", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ViewportImplementer is an interface for Viewport objects.
*/
type ViewportImplementer interface {
	Class
}
