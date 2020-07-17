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

// UPNPDeviceIGDStatus is an enum for IGDStatus values.
type UPNPDeviceIGDStatus int

const (
	UPNPDeviceIgdStatusDisconnected   UPNPDeviceIGDStatus = 5
	UPNPDeviceIgdStatusHttpEmpty      UPNPDeviceIGDStatus = 2
	UPNPDeviceIgdStatusHttpError      UPNPDeviceIGDStatus = 1
	UPNPDeviceIgdStatusInvalidControl UPNPDeviceIGDStatus = 7
	UPNPDeviceIgdStatusMallocError    UPNPDeviceIGDStatus = 8
	UPNPDeviceIgdStatusNoIgd          UPNPDeviceIGDStatus = 4
	UPNPDeviceIgdStatusNoUrls         UPNPDeviceIGDStatus = 3
	UPNPDeviceIgdStatusOk             UPNPDeviceIGDStatus = 0
	UPNPDeviceIgdStatusUnknownDevice  UPNPDeviceIGDStatus = 6
	UPNPDeviceIgdStatusUnknownError   UPNPDeviceIGDStatus = 9
)

//func NewUPNPDeviceFromPointer(ptr gdnative.Pointer) UPNPDevice {
func newUPNPDeviceFromPointer(ptr gdnative.Pointer) UPNPDevice {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := UPNPDevice{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type UPNPDevice struct {
	Reference
	owner gdnative.Object
}

func (o *UPNPDevice) BaseClass() string {
	return "UPNPDevice"
}

/*
        Undocumented
	Args: [{ false port int} {0 true port_internal int} { true desc String} {UDP true proto String} {0 true duration int}], Returns: int
*/
func (o *UPNPDevice) AddPortMapping(port gdnative.Int, portInternal gdnative.Int, desc gdnative.String, proto gdnative.String, duration gdnative.Int) gdnative.Int {
	//log.Println("Calling UPNPDevice.AddPortMapping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(port)
	ptrArguments[1] = gdnative.NewPointerFromInt(portInternal)
	ptrArguments[2] = gdnative.NewPointerFromString(desc)
	ptrArguments[3] = gdnative.NewPointerFromString(proto)
	ptrArguments[4] = gdnative.NewPointerFromInt(duration)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "add_port_mapping")

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
	Args: [{ false port int} {UDP true proto String}], Returns: int
*/
func (o *UPNPDevice) DeletePortMapping(port gdnative.Int, proto gdnative.String) gdnative.Int {
	//log.Println("Calling UPNPDevice.DeletePortMapping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(port)
	ptrArguments[1] = gdnative.NewPointerFromString(proto)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "delete_port_mapping")

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
	Args: [], Returns: String
*/
func (o *UPNPDevice) GetDescriptionUrl() gdnative.String {
	//log.Println("Calling UPNPDevice.GetDescriptionUrl()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "get_description_url")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *UPNPDevice) GetIgdControlUrl() gdnative.String {
	//log.Println("Calling UPNPDevice.GetIgdControlUrl()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "get_igd_control_url")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *UPNPDevice) GetIgdOurAddr() gdnative.String {
	//log.Println("Calling UPNPDevice.GetIgdOurAddr()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "get_igd_our_addr")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *UPNPDevice) GetIgdServiceType() gdnative.String {
	//log.Println("Calling UPNPDevice.GetIgdServiceType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "get_igd_service_type")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.UPNPDevice::IGDStatus
*/
func (o *UPNPDevice) GetIgdStatus() UPNPDeviceIGDStatus {
	//log.Println("Calling UPNPDevice.GetIgdStatus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "get_igd_status")

	// Call the parent method.
	// enum.UPNPDevice::IGDStatus
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return UPNPDeviceIGDStatus(ret)
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *UPNPDevice) GetServiceType() gdnative.String {
	//log.Println("Calling UPNPDevice.GetServiceType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "get_service_type")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *UPNPDevice) IsValidGateway() gdnative.Bool {
	//log.Println("Calling UPNPDevice.IsValidGateway()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "is_valid_gateway")

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
	Args: [], Returns: String
*/
func (o *UPNPDevice) QueryExternalAddress() gdnative.String {
	//log.Println("Calling UPNPDevice.QueryExternalAddress()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "query_external_address")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false url String}], Returns: void
*/
func (o *UPNPDevice) SetDescriptionUrl(url gdnative.String) {
	//log.Println("Calling UPNPDevice.SetDescriptionUrl()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(url)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "set_description_url")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false url String}], Returns: void
*/
func (o *UPNPDevice) SetIgdControlUrl(url gdnative.String) {
	//log.Println("Calling UPNPDevice.SetIgdControlUrl()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(url)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "set_igd_control_url")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false addr String}], Returns: void
*/
func (o *UPNPDevice) SetIgdOurAddr(addr gdnative.String) {
	//log.Println("Calling UPNPDevice.SetIgdOurAddr()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(addr)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "set_igd_our_addr")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false type String}], Returns: void
*/
func (o *UPNPDevice) SetIgdServiceType(aType gdnative.String) {
	//log.Println("Calling UPNPDevice.SetIgdServiceType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "set_igd_service_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false status int}], Returns: void
*/
func (o *UPNPDevice) SetIgdStatus(status gdnative.Int) {
	//log.Println("Calling UPNPDevice.SetIgdStatus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(status)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "set_igd_status")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false type String}], Returns: void
*/
func (o *UPNPDevice) SetServiceType(aType gdnative.String) {
	//log.Println("Calling UPNPDevice.SetServiceType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNPDevice", "set_service_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// UPNPDeviceImplementer is an interface that implements the methods
// of the UPNPDevice class.
type UPNPDeviceImplementer interface {
	ReferenceImplementer
	AddPortMapping(port gdnative.Int, portInternal gdnative.Int, desc gdnative.String, proto gdnative.String, duration gdnative.Int) gdnative.Int
	DeletePortMapping(port gdnative.Int, proto gdnative.String) gdnative.Int
	GetDescriptionUrl() gdnative.String
	GetIgdControlUrl() gdnative.String
	GetIgdOurAddr() gdnative.String
	GetIgdServiceType() gdnative.String
	GetServiceType() gdnative.String
	IsValidGateway() gdnative.Bool
	QueryExternalAddress() gdnative.String
	SetDescriptionUrl(url gdnative.String)
	SetIgdControlUrl(url gdnative.String)
	SetIgdOurAddr(addr gdnative.String)
	SetIgdServiceType(aType gdnative.String)
	SetIgdStatus(status gdnative.Int)
	SetServiceType(aType gdnative.String)
}
