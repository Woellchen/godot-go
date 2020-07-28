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

//func NewWebRTCMultiplayerFromPointer(ptr gdnative.Pointer) WebRTCMultiplayer {
func newWebRTCMultiplayerFromPointer(ptr gdnative.Pointer) WebRTCMultiplayer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := WebRTCMultiplayer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type WebRTCMultiplayer struct {
	NetworkedMultiplayerPeer
	owner gdnative.Object
}

func (o *WebRTCMultiplayer) BaseClass() string {
	return "WebRTCMultiplayer"
}

/*
        Undocumented
	Args: [{ false peer WebRTCPeerConnection} { false peer_id int} {1 true unreliable_lifetime int}], Returns: enum.Error
*/
func (o *WebRTCMultiplayer) AddPeer(peer WebRTCPeerConnectionImplementer, peerId gdnative.Int, unreliableLifetime gdnative.Int) gdnative.Error {
	// log.Println("Calling WebRTCMultiplayer.AddPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(peer.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(peerId)
	ptrArguments[2] = gdnative.NewPointerFromInt(unreliableLifetime)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WebRTCMultiplayer", "add_peer")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *WebRTCMultiplayer) Close() {
	// log.Println("Calling WebRTCMultiplayer.Close()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WebRTCMultiplayer", "close")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false peer_id int}], Returns: Dictionary
*/
func (o *WebRTCMultiplayer) GetPeer(peerId gdnative.Int) gdnative.Dictionary {
	// log.Println("Calling WebRTCMultiplayer.GetPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(peerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WebRTCMultiplayer", "get_peer")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/
func (o *WebRTCMultiplayer) GetPeers() gdnative.Dictionary {
	// log.Println("Calling WebRTCMultiplayer.GetPeers()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WebRTCMultiplayer", "get_peers")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false peer_id int}], Returns: bool
*/
func (o *WebRTCMultiplayer) HasPeer(peerId gdnative.Int) gdnative.Bool {
	// log.Println("Calling WebRTCMultiplayer.HasPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(peerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WebRTCMultiplayer", "has_peer")

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
	Args: [{ false peer_id int} {False true server_compatibility bool}], Returns: enum.Error
*/
func (o *WebRTCMultiplayer) Initialize(peerId gdnative.Int, serverCompatibility gdnative.Bool) gdnative.Error {
	// log.Println("Calling WebRTCMultiplayer.Initialize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(peerId)
	ptrArguments[1] = gdnative.NewPointerFromBool(serverCompatibility)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WebRTCMultiplayer", "initialize")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false peer_id int}], Returns: void
*/
func (o *WebRTCMultiplayer) RemovePeer(peerId gdnative.Int) {
	// log.Println("Calling WebRTCMultiplayer.RemovePeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(peerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WebRTCMultiplayer", "remove_peer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// WebRTCMultiplayerImplementer is an interface that implements the methods
// of the WebRTCMultiplayer class.
type WebRTCMultiplayerImplementer interface {
	NetworkedMultiplayerPeerImplementer
	Close()
	GetPeer(peerId gdnative.Int) gdnative.Dictionary
	GetPeers() gdnative.Dictionary
	HasPeer(peerId gdnative.Int) gdnative.Bool
	RemovePeer(peerId gdnative.Int)
}
