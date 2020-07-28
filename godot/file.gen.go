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

// _FileCompressionMode is an enum for CompressionMode values.
type _FileCompressionMode int

const (
	_FileCompressionDeflate _FileCompressionMode = 1
	_FileCompressionFastlz  _FileCompressionMode = 0
	_FileCompressionGzip    _FileCompressionMode = 3
	_FileCompressionZstd    _FileCompressionMode = 2
)

// _FileModeFlags is an enum for ModeFlags values.
type _FileModeFlags int

const (
	_FileRead      _FileModeFlags = 1
	_FileReadWrite _FileModeFlags = 3
	_FileWrite     _FileModeFlags = 2
	_FileWriteRead _FileModeFlags = 7
)

//func NewFileFromPointer(ptr gdnative.Pointer) File {
func new_FileFromPointer(ptr gdnative.Pointer) File {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := File{}
	obj.SetBaseObject(owner)

	return obj
}

/*
File type. This is used to permanently store data into the user device's file system and to read from it. This can be used to store game save data or player configuration files, for example. Here's a sample on how to write and read from a file: [codeblock] func save(content): var file = File.new() file.open("user://save_game.dat", File.WRITE) file.store_string(content) file.close() func load(): var file = File.new() file.open("user://save_game.dat", File.READ) var content = file.get_as_text() file.close() return content [/codeblock]
*/
type File struct {
	Reference
	owner gdnative.Object
}

func (o *File) BaseClass() string {
	return "_File"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *File) Close() {
	// log.Println("Calling _File.Close()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "close")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *File) EofReached() gdnative.Bool {
	// log.Println("Calling _File.EofReached()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "eof_reached")

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
	Args: [{ false path String}], Returns: bool
*/
func (o *File) FileExists(path gdnative.String) gdnative.Bool {
	// log.Println("Calling _File.FileExists()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "file_exists")

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
	Args: [], Returns: int
*/
func (o *File) Get16() gdnative.Int {
	// log.Println("Calling _File.Get16()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_16")

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
func (o *File) Get32() gdnative.Int {
	// log.Println("Calling _File.Get32()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_32")

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
func (o *File) Get64() gdnative.Int {
	// log.Println("Calling _File.Get64()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_64")

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
func (o *File) Get8() gdnative.Int {
	// log.Println("Calling _File.Get8()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_8")

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
func (o *File) GetAsText() gdnative.String {
	// log.Println("Calling _File.GetAsText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_as_text")

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
	Args: [{ false len int}], Returns: PoolByteArray
*/
func (o *File) GetBuffer(len gdnative.Int) gdnative.PoolByteArray {
	// log.Println("Calling _File.GetBuffer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(len)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_buffer")

	// Call the parent method.
	// PoolByteArray
	retPtr := gdnative.NewEmptyPoolByteArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolByteArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{, true delim String}], Returns: PoolStringArray
*/
func (o *File) GetCsvLine(delim gdnative.String) gdnative.PoolStringArray {
	// log.Println("Calling _File.GetCsvLine()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(delim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_csv_line")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *File) GetDouble() gdnative.Real {
	// log.Println("Calling _File.GetDouble()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_double")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *File) GetEndianSwap() gdnative.Bool {
	// log.Println("Calling _File.GetEndianSwap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_endian_swap")

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
	Args: [], Returns: enum.Error
*/
func (o *File) GetError() gdnative.Error {
	// log.Println("Calling _File.GetError()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_error")

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
	Args: [], Returns: float
*/
func (o *File) GetFloat() gdnative.Real {
	// log.Println("Calling _File.GetFloat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_float")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *File) GetLen() gdnative.Int {
	// log.Println("Calling _File.GetLen()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_len")

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
func (o *File) GetLine() gdnative.String {
	// log.Println("Calling _File.GetLine()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_line")

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
	Args: [{ false path String}], Returns: String
*/
func (o *File) GetMd5(path gdnative.String) gdnative.String {
	// log.Println("Calling _File.GetMd5()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_md5")

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
	Args: [{ false file String}], Returns: int
*/
func (o *File) GetModifiedTime(file gdnative.String) gdnative.Int {
	// log.Println("Calling _File.GetModifiedTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(file)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_modified_time")

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
func (o *File) GetPascalString() gdnative.String {
	// log.Println("Calling _File.GetPascalString()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_pascal_string")

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
func (o *File) GetPath() gdnative.String {
	// log.Println("Calling _File.GetPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_path")

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
func (o *File) GetPathAbsolute() gdnative.String {
	// log.Println("Calling _File.GetPathAbsolute()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_path_absolute")

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
	Args: [], Returns: int
*/
func (o *File) GetPosition() gdnative.Int {
	// log.Println("Calling _File.GetPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_position")

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
	Args: [], Returns: float
*/
func (o *File) GetReal() gdnative.Real {
	// log.Println("Calling _File.GetReal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_real")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false path String}], Returns: String
*/
func (o *File) GetSha256(path gdnative.String) gdnative.String {
	// log.Println("Calling _File.GetSha256()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_sha256")

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
	Args: [{False true allow_objects bool}], Returns: Variant
*/
func (o *File) GetVar(allowObjects gdnative.Bool) gdnative.Variant {
	// log.Println("Calling _File.GetVar()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(allowObjects)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "get_var")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *File) IsOpen() gdnative.Bool {
	// log.Println("Calling _File.IsOpen()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "is_open")

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
	Args: [{ false path String} { false flags int}], Returns: enum.Error
*/
func (o *File) Open(path gdnative.String, flags gdnative.Int) gdnative.Error {
	// log.Println("Calling _File.Open()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "open")

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
	Args: [{ false path String} { false mode_flags int} {0 true compression_mode int}], Returns: enum.Error
*/
func (o *File) OpenCompressed(path gdnative.String, modeFlags gdnative.Int, compressionMode gdnative.Int) gdnative.Error {
	// log.Println("Calling _File.OpenCompressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromInt(modeFlags)
	ptrArguments[2] = gdnative.NewPointerFromInt(compressionMode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "open_compressed")

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
	Args: [{ false path String} { false mode_flags int} { false key PoolByteArray}], Returns: enum.Error
*/
func (o *File) OpenEncrypted(path gdnative.String, modeFlags gdnative.Int, key gdnative.PoolByteArray) gdnative.Error {
	// log.Println("Calling _File.OpenEncrypted()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromInt(modeFlags)
	ptrArguments[2] = gdnative.NewPointerFromPoolByteArray(key)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "open_encrypted")

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
	Args: [{ false path String} { false mode_flags int} { false pass String}], Returns: enum.Error
*/
func (o *File) OpenEncryptedWithPass(path gdnative.String, modeFlags gdnative.Int, pass gdnative.String) gdnative.Error {
	// log.Println("Calling _File.OpenEncryptedWithPass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromInt(modeFlags)
	ptrArguments[2] = gdnative.NewPointerFromString(pass)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "open_encrypted_with_pass")

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
	Args: [{ false position int}], Returns: void
*/
func (o *File) Seek(position gdnative.Int) {
	// log.Println("Calling _File.Seek()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "seek")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{0 true position int}], Returns: void
*/
func (o *File) SeekEnd(position gdnative.Int) {
	// log.Println("Calling _File.SeekEnd()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "seek_end")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *File) SetEndianSwap(enable gdnative.Bool) {
	// log.Println("Calling _File.SetEndianSwap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "set_endian_swap")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false value int}], Returns: void
*/
func (o *File) Store16(value gdnative.Int) {
	// log.Println("Calling _File.Store16()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_16")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false value int}], Returns: void
*/
func (o *File) Store32(value gdnative.Int) {
	// log.Println("Calling _File.Store32()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_32")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false value int}], Returns: void
*/
func (o *File) Store64(value gdnative.Int) {
	// log.Println("Calling _File.Store64()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_64")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false value int}], Returns: void
*/
func (o *File) Store8(value gdnative.Int) {
	// log.Println("Calling _File.Store8()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_8")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false buffer PoolByteArray}], Returns: void
*/
func (o *File) StoreBuffer(buffer gdnative.PoolByteArray) {
	// log.Println("Calling _File.StoreBuffer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolByteArray(buffer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_buffer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false values PoolStringArray} {, true delim String}], Returns: void
*/
func (o *File) StoreCsvLine(values gdnative.PoolStringArray, delim gdnative.String) {
	// log.Println("Calling _File.StoreCsvLine()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromPoolStringArray(values)
	ptrArguments[1] = gdnative.NewPointerFromString(delim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_csv_line")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false value float}], Returns: void
*/
func (o *File) StoreDouble(value gdnative.Real) {
	// log.Println("Calling _File.StoreDouble()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_double")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false value float}], Returns: void
*/
func (o *File) StoreFloat(value gdnative.Real) {
	// log.Println("Calling _File.StoreFloat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_float")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false line String}], Returns: void
*/
func (o *File) StoreLine(line gdnative.String) {
	// log.Println("Calling _File.StoreLine()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(line)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_line")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false string String}], Returns: void
*/
func (o *File) StorePascalString(string gdnative.String) {
	// log.Println("Calling _File.StorePascalString()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(string)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_pascal_string")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false value float}], Returns: void
*/
func (o *File) StoreReal(value gdnative.Real) {
	// log.Println("Calling _File.StoreReal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_real")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false string String}], Returns: void
*/
func (o *File) StoreString(string gdnative.String) {
	// log.Println("Calling _File.StoreString()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(string)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_string")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false value Variant} {False true full_objects bool}], Returns: void
*/
func (o *File) StoreVar(value gdnative.Variant, fullObjects gdnative.Bool) {
	// log.Println("Calling _File.StoreVar()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVariant(value)
	ptrArguments[1] = gdnative.NewPointerFromBool(fullObjects)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_File", "store_var")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// FileImplementer is an interface that implements the methods
// of the File class.
type FileImplementer interface {
	ReferenceImplementer
	Close()
	EofReached() gdnative.Bool
	FileExists(path gdnative.String) gdnative.Bool
	Get16() gdnative.Int
	Get32() gdnative.Int
	Get64() gdnative.Int
	Get8() gdnative.Int
	GetAsText() gdnative.String
	GetBuffer(len gdnative.Int) gdnative.PoolByteArray
	GetCsvLine(delim gdnative.String) gdnative.PoolStringArray
	GetDouble() gdnative.Real
	GetEndianSwap() gdnative.Bool
	GetFloat() gdnative.Real
	GetLen() gdnative.Int
	GetLine() gdnative.String
	GetMd5(path gdnative.String) gdnative.String
	GetModifiedTime(file gdnative.String) gdnative.Int
	GetPascalString() gdnative.String
	GetPath() gdnative.String
	GetPathAbsolute() gdnative.String
	GetPosition() gdnative.Int
	GetReal() gdnative.Real
	GetSha256(path gdnative.String) gdnative.String
	GetVar(allowObjects gdnative.Bool) gdnative.Variant
	IsOpen() gdnative.Bool
	Seek(position gdnative.Int)
	SeekEnd(position gdnative.Int)
	SetEndianSwap(enable gdnative.Bool)
	Store16(value gdnative.Int)
	Store32(value gdnative.Int)
	Store64(value gdnative.Int)
	Store8(value gdnative.Int)
	StoreBuffer(buffer gdnative.PoolByteArray)
	StoreCsvLine(values gdnative.PoolStringArray, delim gdnative.String)
	StoreDouble(value gdnative.Real)
	StoreFloat(value gdnative.Real)
	StoreLine(line gdnative.String)
	StorePascalString(string gdnative.String)
	StoreReal(value gdnative.Real)
	StoreString(string gdnative.String)
	StoreVar(value gdnative.Variant, fullObjects gdnative.Bool)
}
