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

// PacketPeerDTLSStatus is an enum for Status values.
type PacketPeerDTLSStatus int

const (
	PacketPeerDTLSStatusConnected             PacketPeerDTLSStatus = 2
	PacketPeerDTLSStatusDisconnected          PacketPeerDTLSStatus = 0
	PacketPeerDTLSStatusError                 PacketPeerDTLSStatus = 3
	PacketPeerDTLSStatusErrorHostnameMismatch PacketPeerDTLSStatus = 4
	PacketPeerDTLSStatusHandshaking           PacketPeerDTLSStatus = 1
)

//func NewPacketPeerDTLSFromPointer(ptr gdnative.Pointer) PacketPeerDTLS {
func newPacketPeerDTLSFromPointer(ptr gdnative.Pointer) PacketPeerDTLS {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PacketPeerDTLS{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This class represents a DTLS peer connection. It can be used to connect to a DTLS server, and is returned by [method DTLSServer.take_connection].
*/
type PacketPeerDTLS struct {
	PacketPeer
	owner gdnative.Object
}

func (o *PacketPeerDTLS) BaseClass() string {
	return "PacketPeerDTLS"
}

/*
        Connects a [code]peer[/code] beginning the DTLS handshake using the underlying [PacketPeerUDP] which must be connected (see [method PacketPeerUDP.connect_to_host]). If [code]validate_certs[/code] is [code]true[/code], [PacketPeerDTLS] will validate that the certificate presented by the remote peer and match it with the [code]for_hostname[/code] argument. You can specify a custom [X509Certificate] to use for validation via the [code]valid_certificate[/code] argument.
	Args: [{ false packet_peer PacketPeerUDP} {True true validate_certs bool} { true for_hostname String} {[Object:null] true valid_certificate X509Certificate}], Returns: enum.Error
*/
func (o *PacketPeerDTLS) ConnectToPeer(packetPeer PacketPeerUDPImplementer, validateCerts gdnative.Bool, forHostname gdnative.String, validCertificate X509CertificateImplementer) gdnative.Error {
	//log.Println("Calling PacketPeerDTLS.ConnectToPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(packetPeer.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromBool(validateCerts)
	ptrArguments[2] = gdnative.NewPointerFromString(forHostname)
	ptrArguments[3] = gdnative.NewPointerFromObject(validCertificate.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PacketPeerDTLS", "connect_to_peer")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Disconnects this peer, terminating the DTLS session.
	Args: [], Returns: void
*/
func (o *PacketPeerDTLS) DisconnectFromPeer() {
	//log.Println("Calling PacketPeerDTLS.DisconnectFromPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PacketPeerDTLS", "disconnect_from_peer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the status of the connection. See [enum Status] for values.
	Args: [], Returns: enum.PacketPeerDTLS::Status
*/
func (o *PacketPeerDTLS) GetStatus() PacketPeerDTLSStatus {
	//log.Println("Calling PacketPeerDTLS.GetStatus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PacketPeerDTLS", "get_status")

	// Call the parent method.
	// enum.PacketPeerDTLS::Status
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return PacketPeerDTLSStatus(ret)
}

/*
        Poll the connection to check for incoming packets. Call this frequently to update the status and keep the connection working.
	Args: [], Returns: void
*/
func (o *PacketPeerDTLS) Poll() {
	//log.Println("Calling PacketPeerDTLS.Poll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PacketPeerDTLS", "poll")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// PacketPeerDTLSImplementer is an interface that implements the methods
// of the PacketPeerDTLS class.
type PacketPeerDTLSImplementer interface {
	PacketPeerImplementer
	DisconnectFromPeer()
	Poll()
}
