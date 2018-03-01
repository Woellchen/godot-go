package spatial

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
Most basic 3D game object, with a 3D [Transform] and visibility settings. All other 3D game objects inherit from Spatial. Use Spatial as a parent node to move, scale, rotate and show/hide children in a 3D project. Affine operations (rotate, scale, translate) happen in parent's local coordinate system, unless the Spatial object is set as top level. Affine operations in this coordinate system correspond to direct affine operations on the Spatial's transform. The word local below refers to this coordinate system. The coordinate system that is attached to the Spatial object itself is referred to as object-local coordinate system.
*/
type Spatial struct {
	Node
}

func (o *Spatial) BaseClass() string {
	return "Spatial"
}

/*
   Undocumented
*/
func (o *Spatial) X_UpdateGizmo() {
	log.Println("Calling Spatial.X_UpdateGizmo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_update_gizmo", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Spatial) GetGizmo() *SpatialGizmo {
	log.Println("Calling Spatial.GetGizmo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_gizmo", goArguments, "*SpatialGizmo")

	returnValue := goRet.Interface().(*SpatialGizmo)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Spatial) GetGlobalTransform() *Transform {
	log.Println("Calling Spatial.GetGlobalTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_global_transform", goArguments, "*Transform")

	returnValue := goRet.Interface().(*Transform)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the parent [code]Spatial[/code], or an empty [Object] if no parent exists or parent is not of type [code]Spatial[/code].
*/
func (o *Spatial) GetParentSpatial() *Spatial {
	log.Println("Calling Spatial.GetParentSpatial()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_parent_spatial", goArguments, "*Spatial")

	returnValue := goRet.Interface().(*Spatial)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Spatial) GetRotation() *Vector3 {
	log.Println("Calling Spatial.GetRotation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_rotation", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Spatial) GetRotationDegrees() *Vector3 {
	log.Println("Calling Spatial.GetRotationDegrees()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_rotation_degrees", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Spatial) GetScale() *Vector3 {
	log.Println("Calling Spatial.GetScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_scale", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Spatial) GetTransform() *Transform {
	log.Println("Calling Spatial.GetTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_transform", goArguments, "*Transform")

	returnValue := goRet.Interface().(*Transform)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Spatial) GetTranslation() *Vector3 {
	log.Println("Calling Spatial.GetTranslation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_translation", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the current [World] resource this Spatial node is registered to.
*/
func (o *Spatial) GetWorld() *World {
	log.Println("Calling Spatial.GetWorld()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_world", goArguments, "*World")

	returnValue := goRet.Interface().(*World)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Rotates the global (world) transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in global coordinate system.
*/
func (o *Spatial) GlobalRotate(axis *Vector3, angle gdnative.Float) {
	log.Println("Calling Spatial.GlobalRotate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(axis)
	goArguments[1] = reflect.ValueOf(angle)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "global_rotate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *Spatial) GlobalScale(scale *Vector3) {
	log.Println("Calling Spatial.GlobalScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(scale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "global_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Moves the global (world) transformation by [Vector3] offset. The offset is in global coordinate system.
*/
func (o *Spatial) GlobalTranslate(offset *Vector3) {
	log.Println("Calling Spatial.GlobalTranslate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(offset)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "global_translate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Disables rendering of this node. Change Spatial Visible property to false.
*/
func (o *Spatial) Hide() {
	log.Println("Calling Spatial.Hide()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "hide", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns whether node notifies about its local transformation changes. Spatial will not propagate this by default.
*/
func (o *Spatial) IsLocalTransformNotificationEnabled() gdnative.Bool {
	log.Println("Calling Spatial.IsLocalTransformNotificationEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_local_transform_notification_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns whether this node is set as Toplevel, that is whether it ignores its parent nodes transformations.
*/
func (o *Spatial) IsSetAsToplevel() gdnative.Bool {
	log.Println("Calling Spatial.IsSetAsToplevel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_set_as_toplevel", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns whether the node notifies about its global and local transformation changes. Spatial will not propagate this by default.
*/
func (o *Spatial) IsTransformNotificationEnabled() gdnative.Bool {
	log.Println("Calling Spatial.IsTransformNotificationEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_transform_notification_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Spatial) IsVisible() gdnative.Bool {
	log.Println("Calling Spatial.IsVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_visible", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns whether the node is visible, taking into consideration that its parents visibility.
*/
func (o *Spatial) IsVisibleInTree() gdnative.Bool {
	log.Println("Calling Spatial.IsVisibleInTree()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_visible_in_tree", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Rotates itself to point into direction of target position. Operations take place in global space.
*/
func (o *Spatial) LookAt(target *Vector3, up *Vector3) {
	log.Println("Calling Spatial.LookAt()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(target)
	goArguments[1] = reflect.ValueOf(up)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "look_at", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Moves the node to specified position and then rotates itself to point into direction of target position. Operations take place in global space.
*/
func (o *Spatial) LookAtFromPosition(position *Vector3, target *Vector3, up *Vector3) {
	log.Println("Calling Spatial.LookAtFromPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(position)
	goArguments[1] = reflect.ValueOf(target)
	goArguments[2] = reflect.ValueOf(up)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "look_at_from_position", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Resets this node's transformations (like scale, skew and taper) preserving its rotation and translation by performing Gram-Schmidt orthonormalization on this node's [Transform3D].
*/
func (o *Spatial) Orthonormalize() {
	log.Println("Calling Spatial.Orthonormalize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "orthonormalize", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians.
*/
func (o *Spatial) Rotate(axis *Vector3, angle gdnative.Float) {
	log.Println("Calling Spatial.Rotate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(axis)
	goArguments[1] = reflect.ValueOf(angle)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "rotate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in object-local coordinate system.
*/
func (o *Spatial) RotateObjectLocal(axis *Vector3, angle gdnative.Float) {
	log.Println("Calling Spatial.RotateObjectLocal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(axis)
	goArguments[1] = reflect.ValueOf(angle)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "rotate_object_local", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Rotates the local transformation around the X axis by angle in radians
*/
func (o *Spatial) RotateX(angle gdnative.Float) {
	log.Println("Calling Spatial.RotateX()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(angle)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "rotate_x", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Rotates the local transformation around the Y axis by angle in radians.
*/
func (o *Spatial) RotateY(angle gdnative.Float) {
	log.Println("Calling Spatial.RotateY()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(angle)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "rotate_y", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Rotates the local transformation around the Z axis by angle in radians.
*/
func (o *Spatial) RotateZ(angle gdnative.Float) {
	log.Println("Calling Spatial.RotateZ()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(angle)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "rotate_z", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Scales the local transformation by given 3D scale factors in object-local coordinate system.
*/
func (o *Spatial) ScaleObjectLocal(scale *Vector3) {
	log.Println("Calling Spatial.ScaleObjectLocal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(scale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "scale_object_local", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Makes the node ignore its parents transformations. Node transformations are only in global space.
*/
func (o *Spatial) SetAsToplevel(enable gdnative.Bool) {
	log.Println("Calling Spatial.SetAsToplevel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_as_toplevel", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Spatial) SetGizmo(gizmo *SpatialGizmo) {
	log.Println("Calling Spatial.SetGizmo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(gizmo)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_gizmo", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Spatial) SetGlobalTransform(global *Transform) {
	log.Println("Calling Spatial.SetGlobalTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(global)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_global_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Reset all transformations for this node. Set its [Transform3D] to identity matrix.
*/
func (o *Spatial) SetIdentity() {
	log.Println("Calling Spatial.SetIdentity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_identity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether the node ignores notification that its transformation (global or local) changed.
*/
func (o *Spatial) SetIgnoreTransformNotification(enabled gdnative.Bool) {
	log.Println("Calling Spatial.SetIgnoreTransformNotification()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_ignore_transform_notification", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether the node notifies about its local transformation changes. Spatial will not propagate this by default.
*/
func (o *Spatial) SetNotifyLocalTransform(enable gdnative.Bool) {
	log.Println("Calling Spatial.SetNotifyLocalTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_notify_local_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether the node notifies about its global and local transformation changes. Spatial will not propagate this by default.
*/
func (o *Spatial) SetNotifyTransform(enable gdnative.Bool) {
	log.Println("Calling Spatial.SetNotifyTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_notify_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Spatial) SetRotation(euler *Vector3) {
	log.Println("Calling Spatial.SetRotation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(euler)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_rotation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Spatial) SetRotationDegrees(eulerDegrees *Vector3) {
	log.Println("Calling Spatial.SetRotationDegrees()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(eulerDegrees)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_rotation_degrees", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Spatial) SetScale(scale *Vector3) {
	log.Println("Calling Spatial.SetScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(scale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Spatial) SetTransform(local *Transform) {
	log.Println("Calling Spatial.SetTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(local)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Spatial) SetTranslation(translation *Vector3) {
	log.Println("Calling Spatial.SetTranslation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(translation)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_translation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Spatial) SetVisible(visible gdnative.Bool) {
	log.Println("Calling Spatial.SetVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(visible)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_visible", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Enables rendering of this node. Change Spatial Visible property to "True".
*/
func (o *Spatial) Show() {
	log.Println("Calling Spatial.Show()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "show", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Transforms [Vector3] "local_point" from this node's local space to world space.
*/
func (o *Spatial) ToGlobal(localPoint *Vector3) *Vector3 {
	log.Println("Calling Spatial.ToGlobal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(localPoint)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "to_global", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Transforms [Vector3] "global_point" from world space to this node's local space.
*/
func (o *Spatial) ToLocal(globalPoint *Vector3) *Vector3 {
	log.Println("Calling Spatial.ToLocal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(globalPoint)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "to_local", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Changes the node's position by given offset [Vector3].
*/
func (o *Spatial) Translate(offset *Vector3) {
	log.Println("Calling Spatial.Translate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(offset)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "translate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *Spatial) TranslateObjectLocal(offset *Vector3) {
	log.Println("Calling Spatial.TranslateObjectLocal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(offset)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "translate_object_local", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Updates the [SpatialGizmo] of this node.
*/
func (o *Spatial) UpdateGizmo() {
	log.Println("Calling Spatial.UpdateGizmo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "update_gizmo", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   SpatialImplementer is an interface for Spatial objects.
*/
type SpatialImplementer interface {
	Class
}
