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

//func NewSemaphoreFromPointer(ptr gdnative.Pointer) Semaphore {
func new_SemaphoreFromPointer(ptr gdnative.Pointer) Semaphore {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Semaphore{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A synchronization semaphore which can be used to synchronize multiple [Thread]s. Initialized to zero on creation. Be careful to avoid deadlocks. For a binary version, see [Mutex].
*/
type Semaphore struct {
	Reference
	owner gdnative.Object
}

func (o *Semaphore) BaseClass() string {
	return "_Semaphore"
}

/*
        Undocumented
	Args: [], Returns: enum.Error
*/
func (o *Semaphore) Post() gdnative.Error {
	// log.Println("Calling _Semaphore.Post()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Semaphore", "post")

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
	Args: [], Returns: enum.Error
*/
func (o *Semaphore) Wait() gdnative.Error {
	// log.Println("Calling _Semaphore.Wait()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Semaphore", "wait")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

// SemaphoreImplementer is an interface that implements the methods
// of the Semaphore class.
type SemaphoreImplementer interface {
	ReferenceImplementer
}
