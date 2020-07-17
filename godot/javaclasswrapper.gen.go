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

//func NewjavaClassWrapperFromPointer(ptr gdnative.Pointer) javaClassWrapper {
func newJavaClassWrapperFromPointer(ptr gdnative.Pointer) javaClassWrapper {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := javaClassWrapper{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonJavaClassWrapper() *javaClassWrapper {
	return &javaClassWrapper{}
}

/*

 */
var JavaClassWrapper = newSingletonJavaClassWrapper()

/*

 */
type javaClassWrapper struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *javaClassWrapper) ensureSingleton() {
	if o.initialized == true {
		return
	}
	//log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("JavaClassWrapper")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *javaClassWrapper) BaseClass() string {
	return "JavaClassWrapper"
}

/*

	Args: [{ false name String}], Returns: JavaClass
*/
func (o *javaClassWrapper) Wrap(name gdnative.String) JavaClassImplementer {
	o.ensureSingleton()
	//log.Println("Calling JavaClassWrapper.Wrap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("JavaClassWrapper", "wrap")

	// Call the parent method.
	// JavaClass
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newJavaClassFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(JavaClassImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "JavaClass" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(JavaClassImplementer)
	}

	return &ret
}

// JavaClassWrapperImplementer is an interface that implements the methods
// of the JavaClassWrapper class.
type JavaClassWrapperImplementer interface {
	ObjectImplementer
	Wrap(name gdnative.String) JavaClassImplementer
}
