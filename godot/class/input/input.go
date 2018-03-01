package input

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

func newSingletonInput() *input {
	obj := &input{}
	ptr := C.godot_global_get_singleton(C.CString("Input"))
	obj.owner = (*C.godot_object)(ptr)
	return obj
}

/*
   A Singleton that deals with inputs. This includes key presses, mouse buttons and movement, joypads, and input actions. Actions and their events can be set in the Project Settings / Input Map tab. Or be set with [InputMap].
*/
var Input = newSingletonInput()

/*
A Singleton that deals with inputs. This includes key presses, mouse buttons and movement, joypads, and input actions. Actions and their events can be set in the Project Settings / Input Map tab. Or be set with [InputMap].
*/
type input struct {
	Object
}

func (o *input) BaseClass() string {
	return "Input"
}

/*
   This will simulate pressing the specified action.
*/
func (o *input) ActionPress(action gdnative.String) {
	log.Println("Calling Input.ActionPress()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(action)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "action_press", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If the specified action is already pressed, this will release it.
*/
func (o *input) ActionRelease(action gdnative.String) {
	log.Println("Calling Input.ActionRelease()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(action)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "action_release", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Add a new mapping entry (in SDL2 format) to the mapping database. Optionally update already connected devices.
*/
func (o *input) AddJoyMapping(mapping gdnative.String, updateExisting gdnative.Bool) {
	log.Println("Calling Input.AddJoyMapping()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(mapping)
	goArguments[1] = reflect.ValueOf(updateExisting)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_joy_mapping", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If the device has an accelerometer, this will return the movement.
*/
func (o *input) GetAccelerometer() *Vector3 {
	log.Println("Calling Input.GetAccelerometer()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_accelerometer", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns an [Array] containing the device IDs of all currently connected joypads.
*/
func (o *input) GetConnectedJoypads() *Array {
	log.Println("Calling Input.GetConnectedJoypads()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_connected_joypads", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *input) GetGravity() *Vector3 {
	log.Println("Calling Input.GetGravity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_gravity", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   If the device has a gyroscope, this will return the rate of rotation in rad/s around a device's x, y, and z axis.
*/
func (o *input) GetGyroscope() *Vector3 {
	log.Println("Calling Input.GetGyroscope()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_gyroscope", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the current value of the joypad axis at given index (see [code]JOY_*[/code] constants in [@GlobalScope])
*/
func (o *input) GetJoyAxis(device gdnative.Int, axis gdnative.Int) gdnative.Float {
	log.Println("Calling Input.GetJoyAxis()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(device)
	goArguments[1] = reflect.ValueOf(axis)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_joy_axis", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *input) GetJoyAxisIndexFromString(axis gdnative.String) gdnative.Int {
	log.Println("Calling Input.GetJoyAxisIndexFromString()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(axis)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_joy_axis_index_from_string", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *input) GetJoyAxisString(axisIndex gdnative.Int) gdnative.String {
	log.Println("Calling Input.GetJoyAxisString()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(axisIndex)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_joy_axis_string", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *input) GetJoyButtonIndexFromString(button gdnative.String) gdnative.Int {
	log.Println("Calling Input.GetJoyButtonIndexFromString()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(button)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_joy_button_index_from_string", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *input) GetJoyButtonString(buttonIndex gdnative.Int) gdnative.String {
	log.Println("Calling Input.GetJoyButtonString()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(buttonIndex)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_joy_button_string", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a SDL2 compatible device guid on platforms that use gamepad remapping. Returns "Default Gamepad" otherwise.
*/
func (o *input) GetJoyGuid(device gdnative.Int) gdnative.String {
	log.Println("Calling Input.GetJoyGuid()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(device)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_joy_guid", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the name of the joypad at the specified device index
*/
func (o *input) GetJoyName(device gdnative.Int) gdnative.String {
	log.Println("Calling Input.GetJoyName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(device)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_joy_name", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the duration of the current vibration effect in seconds.
*/
func (o *input) GetJoyVibrationDuration(device gdnative.Int) gdnative.Float {
	log.Println("Calling Input.GetJoyVibrationDuration()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(device)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_joy_vibration_duration", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the strength of the joypad vibration: x is the strength of the weak motor, and y is the strength of the strong motor.
*/
func (o *input) GetJoyVibrationStrength(device gdnative.Int) *Vector2 {
	log.Println("Calling Input.GetJoyVibrationStrength()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(device)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_joy_vibration_strength", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the mouse speed for the last time the cursor was moved, and this until the next frame where the mouse moves. This means that even if the mouse is not moving, this function will still return the value of the last motion.
*/
func (o *input) GetLastMouseSpeed() *Vector2 {
	log.Println("Calling Input.GetLastMouseSpeed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_last_mouse_speed", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   If the device has a magnetometer, this will return the magnetic field strength in micro-Tesla for all axes.
*/
func (o *input) GetMagnetometer() *Vector3 {
	log.Println("Calling Input.GetMagnetometer()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_magnetometer", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns mouse buttons as a bitmask. If multiple mouse buttons are pressed at the same time the bits are added together.
*/
func (o *input) GetMouseButtonMask() gdnative.Int {
	log.Println("Calling Input.GetMouseButtonMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mouse_button_mask", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the mouse mode. See the constants for more information.
*/
func (o *input) GetMouseMode() gdnative.Int {
	log.Println("Calling Input.GetMouseMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mouse_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] when you start pressing the action event.
*/
func (o *input) IsActionJustPressed(action gdnative.String) gdnative.Bool {
	log.Println("Calling Input.IsActionJustPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(action)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_action_just_pressed", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] when you stop pressing the action event.
*/
func (o *input) IsActionJustReleased(action gdnative.String) gdnative.Bool {
	log.Println("Calling Input.IsActionJustReleased()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(action)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_action_just_released", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if you are pressing the action event.
*/
func (o *input) IsActionPressed(action gdnative.String) gdnative.Bool {
	log.Println("Calling Input.IsActionPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(action)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_action_pressed", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if you are pressing the joypad button. (see [code]JOY_*[/code] constants in [@GlobalScope])
*/
func (o *input) IsJoyButtonPressed(device gdnative.Int, button gdnative.Int) gdnative.Bool {
	log.Println("Calling Input.IsJoyButtonPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(device)
	goArguments[1] = reflect.ValueOf(button)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_joy_button_pressed", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the system knows the specified device. This means that it sets all button and axis indices exactly as defined in the [code]JOY_*[/code] constants (see [@GlobalScope]). Unknown joypads are not expected to match these constants, but you can still retrieve events from them.
*/
func (o *input) IsJoyKnown(device gdnative.Int) gdnative.Bool {
	log.Println("Calling Input.IsJoyKnown()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(device)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_joy_known", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if you are pressing the key. You can pass [code]KEY_*[/code], which are pre-defined constants listed in [@GlobalScope].
*/
func (o *input) IsKeyPressed(scancode gdnative.Int) gdnative.Bool {
	log.Println("Calling Input.IsKeyPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(scancode)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_key_pressed", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if you are pressing the mouse button. You can pass [code]BUTTON_*[/code], which are pre-defined constants listed in [@GlobalScope].
*/
func (o *input) IsMouseButtonPressed(button gdnative.Int) gdnative.Bool {
	log.Println("Calling Input.IsMouseButtonPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(button)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_mouse_button_pressed", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *input) JoyConnectionChanged(device gdnative.Int, connected gdnative.Bool, name gdnative.String, guid gdnative.String) {
	log.Println("Calling Input.JoyConnectionChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(device)
	goArguments[1] = reflect.ValueOf(connected)
	goArguments[2] = reflect.ValueOf(name)
	goArguments[3] = reflect.ValueOf(guid)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "joy_connection_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *input) ParseInputEvent(event *InputEvent) {
	log.Println("Calling Input.ParseInputEvent()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(event)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "parse_input_event", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes all mappings from the internal db that match the given uid.
*/
func (o *input) RemoveJoyMapping(guid gdnative.String) {
	log.Println("Calling Input.RemoveJoyMapping()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(guid)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "remove_joy_mapping", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set a custom mouse cursor image, which is only visible inside the game window. The hotspot can also be specified. See enum [code]CURSOR_*[/code] for the list of shapes.
*/
func (o *input) SetCustomMouseCursor(image *Resource, shape gdnative.Int, hotspot *Vector2) {
	log.Println("Calling Input.SetCustomMouseCursor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(image)
	goArguments[1] = reflect.ValueOf(shape)
	goArguments[2] = reflect.ValueOf(hotspot)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_custom_mouse_cursor", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the mouse mode. See the constants for more information.
*/
func (o *input) SetMouseMode(mode gdnative.Int) {
	log.Println("Calling Input.SetMouseMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_mouse_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Starts to vibrate the joypad. Joypads usually come with two rumble motors, a strong and a weak one. weak_magnitude is the strength of the weak motor (between 0 and 1) and strong_magnitude is the strength of the strong motor (between 0 and 1). duration is the duration of the effect in seconds (a duration of 0 will try to play the vibration indefinitely). Note that not every hardware is compatible with long effect durations, it is recommended to restart an effect if in need to play it for more than a few seconds.
*/
func (o *input) StartJoyVibration(device gdnative.Int, weakMagnitude gdnative.Float, strongMagnitude gdnative.Float, duration gdnative.Float) {
	log.Println("Calling Input.StartJoyVibration()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(device)
	goArguments[1] = reflect.ValueOf(weakMagnitude)
	goArguments[2] = reflect.ValueOf(strongMagnitude)
	goArguments[3] = reflect.ValueOf(duration)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "start_joy_vibration", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Stops the vibration of the joypad.
*/
func (o *input) StopJoyVibration(device gdnative.Int) {
	log.Println("Calling Input.StopJoyVibration()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(device)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "stop_joy_vibration", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the mouse position to the specified vector.
*/
func (o *input) WarpMousePosition(to *Vector2) {
	log.Println("Calling Input.WarpMousePosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(to)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "warp_mouse_position", goArguments, "")

	log.Println("  Function successfully completed.")

}
