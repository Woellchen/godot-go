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

// UPNPUPNPResult is an enum for UPNPResult values.
type UPNPUPNPResult int

const (
	UPNPUpnpResultActionFailed                UPNPUPNPResult = 5
	UPNPUpnpResultConflictWithOtherMapping    UPNPUPNPResult = 13
	UPNPUpnpResultConflictWithOtherMechanism  UPNPUPNPResult = 12
	UPNPUpnpResultExtPortMustBeWildcard       UPNPUPNPResult = 10
	UPNPUpnpResultExtPortWildcardNotPermitted UPNPUPNPResult = 7
	UPNPUpnpResultHttpError                   UPNPUPNPResult = 23
	UPNPUpnpResultInconsistentParameters      UPNPUPNPResult = 3
	UPNPUpnpResultIntPortWildcardNotPermitted UPNPUPNPResult = 8
	UPNPUpnpResultInvalidArgs                 UPNPUPNPResult = 20
	UPNPUpnpResultInvalidDuration             UPNPUPNPResult = 19
	UPNPUpnpResultInvalidGateway              UPNPUPNPResult = 16
	UPNPUpnpResultInvalidParam                UPNPUPNPResult = 22
	UPNPUpnpResultInvalidPort                 UPNPUPNPResult = 17
	UPNPUpnpResultInvalidProtocol             UPNPUPNPResult = 18
	UPNPUpnpResultInvalidResponse             UPNPUPNPResult = 21
	UPNPUpnpResultMemAllocError               UPNPUPNPResult = 25
	UPNPUpnpResultNotAuthorized               UPNPUPNPResult = 1
	UPNPUpnpResultNoDevices                   UPNPUPNPResult = 27
	UPNPUpnpResultNoGateway                   UPNPUPNPResult = 26
	UPNPUpnpResultNoPortMapsAvailable         UPNPUPNPResult = 11
	UPNPUpnpResultNoSuchEntryInArray          UPNPUPNPResult = 4
	UPNPUpnpResultOnlyPermanentLeaseSupported UPNPUPNPResult = 15
	UPNPUpnpResultPortMappingNotFound         UPNPUPNPResult = 2
	UPNPUpnpResultRemoteHostMustBeWildcard    UPNPUPNPResult = 9
	UPNPUpnpResultSamePortValuesRequired      UPNPUPNPResult = 14
	UPNPUpnpResultSocketError                 UPNPUPNPResult = 24
	UPNPUpnpResultSrcIpWildcardNotPermitted   UPNPUPNPResult = 6
	UPNPUpnpResultSuccess                     UPNPUPNPResult = 0
	UPNPUpnpResultUnknownError                UPNPUPNPResult = 28
)

//func NewUPNPFromPointer(ptr gdnative.Pointer) UPNP {
func newUPNPFromPointer(ptr gdnative.Pointer) UPNP {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := UPNP{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type UPNP struct {
	Reference
	owner gdnative.Object
}

func (o *UPNP) BaseClass() string {
	return "UPNP"
}

/*
        Undocumented
	Args: [{ false device UPNPDevice}], Returns: void
*/
func (o *UPNP) AddDevice(device UPNPDeviceImplementer) {
	// log.Println("Calling UPNP.AddDevice()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(device.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "add_device")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false port int} {0 true port_internal int} { true desc String} {UDP true proto String} {0 true duration int}], Returns: int
*/
func (o *UPNP) AddPortMapping(port gdnative.Int, portInternal gdnative.Int, desc gdnative.String, proto gdnative.String, duration gdnative.Int) gdnative.Int {
	// log.Println("Calling UPNP.AddPortMapping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(port)
	ptrArguments[1] = gdnative.NewPointerFromInt(portInternal)
	ptrArguments[2] = gdnative.NewPointerFromString(desc)
	ptrArguments[3] = gdnative.NewPointerFromString(proto)
	ptrArguments[4] = gdnative.NewPointerFromInt(duration)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "add_port_mapping")

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
	Args: [], Returns: void
*/
func (o *UPNP) ClearDevices() {
	// log.Println("Calling UPNP.ClearDevices()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "clear_devices")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false port int} {UDP true proto String}], Returns: int
*/
func (o *UPNP) DeletePortMapping(port gdnative.Int, proto gdnative.String) gdnative.Int {
	// log.Println("Calling UPNP.DeletePortMapping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(port)
	ptrArguments[1] = gdnative.NewPointerFromString(proto)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "delete_port_mapping")

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
	Args: [{2000 true timeout int} {2 true ttl int} {InternetGatewayDevice true device_filter String}], Returns: int
*/
func (o *UPNP) Discover(timeout gdnative.Int, ttl gdnative.Int, deviceFilter gdnative.String) gdnative.Int {
	// log.Println("Calling UPNP.Discover()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(timeout)
	ptrArguments[1] = gdnative.NewPointerFromInt(ttl)
	ptrArguments[2] = gdnative.NewPointerFromString(deviceFilter)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "discover")

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
	Args: [{ false index int}], Returns: UPNPDevice
*/
func (o *UPNP) GetDevice(index gdnative.Int) UPNPDeviceImplementer {
	// log.Println("Calling UPNP.GetDevice()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "get_device")

	// Call the parent method.
	// UPNPDevice
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newUPNPDeviceFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(UPNPDeviceImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "UPNPDevice" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(UPNPDeviceImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *UPNP) GetDeviceCount() gdnative.Int {
	// log.Println("Calling UPNP.GetDeviceCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "get_device_count")

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
	Args: [], Returns: int
*/
func (o *UPNP) GetDiscoverLocalPort() gdnative.Int {
	// log.Println("Calling UPNP.GetDiscoverLocalPort()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "get_discover_local_port")

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
func (o *UPNP) GetDiscoverMulticastIf() gdnative.String {
	// log.Println("Calling UPNP.GetDiscoverMulticastIf()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "get_discover_multicast_if")

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
	Args: [], Returns: UPNPDevice
*/
func (o *UPNP) GetGateway() UPNPDeviceImplementer {
	// log.Println("Calling UPNP.GetGateway()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "get_gateway")

	// Call the parent method.
	// UPNPDevice
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newUPNPDeviceFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(UPNPDeviceImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "UPNPDevice" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(UPNPDeviceImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *UPNP) IsDiscoverIpv6() gdnative.Bool {
	// log.Println("Calling UPNP.IsDiscoverIpv6()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "is_discover_ipv6")

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
func (o *UPNP) QueryExternalAddress() gdnative.String {
	// log.Println("Calling UPNP.QueryExternalAddress()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "query_external_address")

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
	Args: [{ false index int}], Returns: void
*/
func (o *UPNP) RemoveDevice(index gdnative.Int) {
	// log.Println("Calling UPNP.RemoveDevice()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "remove_device")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false index int} { false device UPNPDevice}], Returns: void
*/
func (o *UPNP) SetDevice(index gdnative.Int, device UPNPDeviceImplementer) {
	// log.Println("Calling UPNP.SetDevice()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromObject(device.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "set_device")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ipv6 bool}], Returns: void
*/
func (o *UPNP) SetDiscoverIpv6(ipv6 gdnative.Bool) {
	// log.Println("Calling UPNP.SetDiscoverIpv6()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(ipv6)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "set_discover_ipv6")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false port int}], Returns: void
*/
func (o *UPNP) SetDiscoverLocalPort(port gdnative.Int) {
	// log.Println("Calling UPNP.SetDiscoverLocalPort()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(port)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "set_discover_local_port")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false m_if String}], Returns: void
*/
func (o *UPNP) SetDiscoverMulticastIf(mIf gdnative.String) {
	// log.Println("Calling UPNP.SetDiscoverMulticastIf()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(mIf)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("UPNP", "set_discover_multicast_if")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// UPNPImplementer is an interface that implements the methods
// of the UPNP class.
type UPNPImplementer interface {
	ReferenceImplementer
	AddDevice(device UPNPDeviceImplementer)
	AddPortMapping(port gdnative.Int, portInternal gdnative.Int, desc gdnative.String, proto gdnative.String, duration gdnative.Int) gdnative.Int
	ClearDevices()
	DeletePortMapping(port gdnative.Int, proto gdnative.String) gdnative.Int
	Discover(timeout gdnative.Int, ttl gdnative.Int, deviceFilter gdnative.String) gdnative.Int
	GetDevice(index gdnative.Int) UPNPDeviceImplementer
	GetDeviceCount() gdnative.Int
	GetDiscoverLocalPort() gdnative.Int
	GetDiscoverMulticastIf() gdnative.String
	GetGateway() UPNPDeviceImplementer
	IsDiscoverIpv6() gdnative.Bool
	QueryExternalAddress() gdnative.String
	RemoveDevice(index gdnative.Int)
	SetDevice(index gdnative.Int, device UPNPDeviceImplementer)
	SetDiscoverIpv6(ipv6 gdnative.Bool)
	SetDiscoverLocalPort(port gdnative.Int)
	SetDiscoverMulticastIf(mIf gdnative.String)
}
