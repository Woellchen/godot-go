package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#cgo CFLAGS: -I../godot_headers
#include "gdnative.gen.h"
// #include <godot_headers/net/godot_net.h>
// Include all headers for now. TODO: Look up all the required
// headers we need to import based on the method arguments and return types.
#include <gdnative/aabb.h>
#include <gdnative/array.h>
#include <gdnative/basis.h>
#include <gdnative/color.h>
#include <gdnative/dictionary.h>
#include <gdnative/gdnative.h>
#include <gdnative/node_path.h>
#include <gdnative/plane.h>
#include <gdnative/pool_arrays.h>
#include <gdnative/quat.h>
#include <gdnative/rect2.h>
#include <gdnative/rid.h>
#include <gdnative/string.h>
#include <gdnative/string_name.h>
#include <gdnative/transform.h>
#include <gdnative/transform2d.h>
#include <gdnative/variant.h>
#include <gdnative/vector2.h>
#include <gdnative/vector3.h>
#include <gdnative_api_struct.gen.h>
*/
import "C"
import "unsafe"

// NewEmptyNetStreamPeer will return a pointer to an empty
// initialized NetStreamPeer. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyNetStreamPeer() Pointer {
	var obj C.godot_net_stream_peer
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromNetStreamPeer will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromNetStreamPeer(obj NetStreamPeer) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewNetStreamPeerFromPointer will return a NetStreamPeer from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewNetStreamPeerFromPointer(ptr Pointer) NetStreamPeer {

	return NetStreamPeer{base: (*C.godot_net_stream_peer)(ptr.getBase())}
}

type NetStreamPeer struct {
	base *C.godot_net_stream_peer

	Version GdnativeApiVersion
}

func (gdt NetStreamPeer) getBase() *C.godot_net_stream_peer {
	return gdt.base
}

// NewEmptyNetPacketPeer will return a pointer to an empty
// initialized NetPacketPeer. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyNetPacketPeer() Pointer {
	var obj C.godot_net_packet_peer
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromNetPacketPeer will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromNetPacketPeer(obj NetPacketPeer) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewNetPacketPeerFromPointer will return a NetPacketPeer from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewNetPacketPeerFromPointer(ptr Pointer) NetPacketPeer {

	return NetPacketPeer{base: (*C.godot_net_packet_peer)(ptr.getBase())}
}

type NetPacketPeer struct {
	base *C.godot_net_packet_peer

	Version GdnativeApiVersion
}

func (gdt NetPacketPeer) getBase() *C.godot_net_packet_peer {
	return gdt.base
}

// NewEmptyNetMultiplayerPeer will return a pointer to an empty
// initialized NetMultiplayerPeer. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyNetMultiplayerPeer() Pointer {
	var obj C.godot_net_multiplayer_peer
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromNetMultiplayerPeer will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromNetMultiplayerPeer(obj NetMultiplayerPeer) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewNetMultiplayerPeerFromPointer will return a NetMultiplayerPeer from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewNetMultiplayerPeerFromPointer(ptr Pointer) NetMultiplayerPeer {

	return NetMultiplayerPeer{base: (*C.godot_net_multiplayer_peer)(ptr.getBase())}
}

type NetMultiplayerPeer struct {
	base *C.godot_net_multiplayer_peer

	Version GdnativeApiVersion
}

func (gdt NetMultiplayerPeer) getBase() *C.godot_net_multiplayer_peer {
	return gdt.base
}
