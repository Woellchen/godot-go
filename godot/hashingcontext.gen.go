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

// HashingContextHashType is an enum for HashType values.
type HashingContextHashType int

const (
	HashingContextHashMd5    HashingContextHashType = 0
	HashingContextHashSha1   HashingContextHashType = 1
	HashingContextHashSha256 HashingContextHashType = 2
)

//func NewHashingContextFromPointer(ptr gdnative.Pointer) HashingContext {
func newHashingContextFromPointer(ptr gdnative.Pointer) HashingContext {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HashingContext{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The HashingContext class provides an interface for computing cryptographic hashes over multiple iterations. This is useful for example when computing hashes of big files (so you don't have to load them all in memory), network streams, and data streams in general (so you don't have to hold buffers). The [enum HashType] enum shows the supported hashing algorithms. [codeblock] const CHUNK_SIZE = 1024 func hash_file(path): var ctx = HashingContext.new() var file = File.new() # Start a SHA-256 context. ctx.start(HashingContext.HASH_SHA256) # Check that file exists. if not file.file_exists(path): return # Open the file to hash. file.open(path, File.READ) # Update the context after reading each chunk. while not file.eof_reached(): ctx.update(file.get_buffer(CHUNK_SIZE)) # Get the computed hash. var res = ctx.finish() # Print the result as hex string and array. printt(res.hex_encode(), Array(res)) [/codeblock] [b]Note:[/b] Not available in HTML5 exports.
*/
type HashingContext struct {
	Reference
	owner gdnative.Object
}

func (o *HashingContext) BaseClass() string {
	return "HashingContext"
}

/*
        Closes the current context, and return the computed hash.
	Args: [], Returns: PoolByteArray
*/
func (o *HashingContext) Finish() gdnative.PoolByteArray {
	//log.Println("Calling HashingContext.Finish()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HashingContext", "finish")

	// Call the parent method.
	// PoolByteArray
	retPtr := gdnative.NewEmptyPoolByteArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolByteArrayFromPointer(retPtr)
	return ret
}

/*
        Starts a new hash computation of the given [code]type[/code] (e.g. [constant HASH_SHA256] to start computation of a SHA-256).
	Args: [{ false type int}], Returns: enum.Error
*/
func (o *HashingContext) Start(aType gdnative.Int) gdnative.Error {
	//log.Println("Calling HashingContext.Start()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HashingContext", "start")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Updates the computation with the given [code]chunk[/code] of data.
	Args: [{ false chunk PoolByteArray}], Returns: enum.Error
*/
func (o *HashingContext) Update(chunk gdnative.PoolByteArray) gdnative.Error {
	//log.Println("Calling HashingContext.Update()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolByteArray(chunk)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HashingContext", "update")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

// HashingContextImplementer is an interface that implements the methods
// of the HashingContext class.
type HashingContextImplementer interface {
	ReferenceImplementer
	Finish() gdnative.PoolByteArray
}
