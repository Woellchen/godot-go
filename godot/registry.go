package godot

import (
	"github.com/Woellchen/godot-go/gdnative"
	"reflect"
)

// AutoRegister will register the given object(s) as a Godot class, so it will be available
// inside Godot. This method uses reflection at runtime to automatically discover the given
// Go structs' methods and properties. It will automatically map Godot method names in
// CamelCase to Godot's conventional snake_case. It will also automatically use the
// struct name for the class name to register as.
func AutoRegister(constructor ...ClassConstructor) {
	for _, construct := range constructor {
		godotConstructorsToAutoRegister = append(godotConstructorsToAutoRegister, construct)
	}
}

// Register class will register the given Go struct as a Godot class, so it will be
// available inside Godot. If you use this method, you will also need to manually
// register this method's properties and methods. If you want to automatically
// discover this class's methods and properties, use the 'AutoRegister' function.
//func RegisterClass(className string, constructor ClassConstructor) {
//	godotConstructorsToRegister[className] = constructor
//}

// Class is an interface for any objects that can have Godot
// inheritance.
type Class interface {
	BaseClass() string
	SetBaseObject(object gdnative.Object)
	GetBaseObject() gdnative.Object
}

// ClassConstructor is any function that will build and return a class to be registered
// with Godot.
type ClassConstructor func() Class

// registeredClass is a structure for holding on to the reflected details of a Go
// struct that has been registered as a Godot class. It has the struct's name,
// methods, the arguments that method takes, and the value types it returns.
type registeredClass struct {
	name       string
	structType reflect.Type
	methods    map[string]*registeredMethod
}

// newRegisteredClass will return a structure for
func newRegisteredClass(classType reflect.Type) *registeredClass {
	class := &registeredClass{
		name:       classType.String(),
		structType: classType,
		methods:    map[string]*registeredMethod{},
	}

	return class
}

// addMethod will add the given registered method to the class.
func (r *registeredClass) addMethod(name string, method *registeredMethod) {
	r.methods[name] = method
}

// registeredMethod is a structure for holding on to the reflected details of a Go
// method that has been registered as a Godot method. It contains the method's
// argument types and return types.
type registeredMethod struct {
	method    reflect.Method
	arguments []reflect.Type
	returns   []reflect.Type
}

// newRegisteredMethod takes in a struct type and uses reflection to discover all
// of the methods it has attached to it, and returns a registered method structure.
func newRegisteredMethod(classMethod reflect.Method) *registeredMethod {
	method := &registeredMethod{}
	method.method = classMethod

	// Get the type of the method.
	methodType := method.method.Type

	// Check the number of arguments the method takes and returns.
	numArgs := methodType.NumIn()
	numReturns := methodType.NumOut()
	method.arguments = make([]reflect.Type, numArgs)
	method.returns = make([]reflect.Type, numReturns)

	// Loop through the arguments and check what types of arguments they
	// take.
	for j := 0; j < numArgs; j++ {
		argType := methodType.In(j)
		method.arguments[j] = argType
	}

	// Loop through the returns and check what types it returns.
	for j := 0; j < numReturns; j++ {
		retType := methodType.Out(j)
		method.returns[j] = retType
	}

	return method
}

// godotClassesToAutoRegister is a slice of objects that will be registered as a Godot class
// upon library initialization. It will automatically inspect the object for exported
// properties and methods.
var godotConstructorsToAutoRegister = []ClassConstructor{}

// godotConstructorsToRegister is a slice of objects that will be registered as a Godot class
// upon library initialization. It will not attempt to discover methods or properties.
var godotConstructorsToRegister = map[string]ClassConstructor{}

// classRegistry is a mapping of all classes that have been registered in Godot.
var classRegistry = map[string]*registeredClass{}

// InstanceRegistry is a mapping of all instances that have currently been created. This map
// allows instance methods to find which instance they belong to.
var InstanceRegistry = &classInstanceRegistry{registry: map[string]Class{}}

// classInstanceRegistry is a structure for holding on to Class instances that have
// been constructed.
type classInstanceRegistry struct {
	registry map[string]Class
}

// Add will add the instance to the registry.
func (i *classInstanceRegistry) Add(instanceID string, class Class) {
	i.registry[instanceID] = class
}

// Get will return the instance with the given instance ID.
func (i *classInstanceRegistry) Get(instanceID string) (Class, bool) {
	if instance, ok := i.registry[instanceID]; ok {
		return instance, true
	}
	return nil, false
}

// Delete will delete the given instance from the registry, so it can be
// garbage collected.
func (i *classInstanceRegistry) Delete(instanceID string) {
	delete(i.registry, instanceID)
}
