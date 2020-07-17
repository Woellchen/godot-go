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

//func NewPacketPeerGDNativeFromPointer(ptr gdnative.Pointer) PacketPeerGDNative {
func newPacketPeerGDNativeFromPointer(ptr gdnative.Pointer) PacketPeerGDNative {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PacketPeerGDNative{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type PacketPeerGDNative struct {
	PacketPeer
	owner gdnative.Object
}

func (o *PacketPeerGDNative) BaseClass() string {
	return "PacketPeerGDNative"
}

// PacketPeerGDNativeImplementer is an interface that implements the methods
// of the PacketPeerGDNative class.
type PacketPeerGDNativeImplementer interface {
	PacketPeerImplementer
}
