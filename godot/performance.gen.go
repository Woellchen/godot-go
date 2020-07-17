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

// PerformanceMonitor is an enum for Monitor values.
type PerformanceMonitor int

const (
	PerformanceAudioOutputLatency           PerformanceMonitor = 30
	PerformanceMemoryDynamic                PerformanceMonitor = 4
	PerformanceMemoryDynamicMax             PerformanceMonitor = 6
	PerformanceMemoryMessageBufferMax       PerformanceMonitor = 7
	PerformanceMemoryStatic                 PerformanceMonitor = 3
	PerformanceMemoryStaticMax              PerformanceMonitor = 5
	PerformanceMonitorMax                   PerformanceMonitor = 31
	PerformanceObjectCount                  PerformanceMonitor = 8
	PerformanceObjectNodeCount              PerformanceMonitor = 10
	PerformanceObjectOrphanNodeCount        PerformanceMonitor = 11
	PerformanceObjectResourceCount          PerformanceMonitor = 9
	PerformancePhysics2DActiveObjects       PerformanceMonitor = 24
	PerformancePhysics2DCollisionPairs      PerformanceMonitor = 25
	PerformancePhysics2DIslandCount         PerformanceMonitor = 26
	PerformancePhysics3DActiveObjects       PerformanceMonitor = 27
	PerformancePhysics3DCollisionPairs      PerformanceMonitor = 28
	PerformancePhysics3DIslandCount         PerformanceMonitor = 29
	PerformanceRender2DDrawCallsInFrame     PerformanceMonitor = 19
	PerformanceRender2DItemsInFrame         PerformanceMonitor = 18
	PerformanceRenderDrawCallsInFrame       PerformanceMonitor = 17
	PerformanceRenderMaterialChangesInFrame PerformanceMonitor = 14
	PerformanceRenderObjectsInFrame         PerformanceMonitor = 12
	PerformanceRenderShaderChangesInFrame   PerformanceMonitor = 15
	PerformanceRenderSurfaceChangesInFrame  PerformanceMonitor = 16
	PerformanceRenderTextureMemUsed         PerformanceMonitor = 21
	PerformanceRenderUsageVideoMemTotal     PerformanceMonitor = 23
	PerformanceRenderVertexMemUsed          PerformanceMonitor = 22
	PerformanceRenderVerticesInFrame        PerformanceMonitor = 13
	PerformanceRenderVideoMemUsed           PerformanceMonitor = 20
	PerformanceTimeFps                      PerformanceMonitor = 0
	PerformanceTimePhysicsProcess           PerformanceMonitor = 2
	PerformanceTimeProcess                  PerformanceMonitor = 1
)

//func NewperformanceFromPointer(ptr gdnative.Pointer) performance {
func newPerformanceFromPointer(ptr gdnative.Pointer) performance {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := performance{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonPerformance() *performance {
	return &performance{}
}

/*
   This class provides access to a number of different monitors related to performance, such as memory usage, draw calls, and FPS. These are the same as the values displayed in the [b]Monitor[/b] tab in the editor's [b]Debugger[/b] panel. By using the [method get_monitor] method of this class, you can access this data from your code. [b]Note:[/b] A few of these monitors are only available in debug mode and will always return 0 when used in a release build. [b]Note:[/b] Many of these monitors are not updated in real-time, so there may be a short delay between changes.
*/
var Performance = newSingletonPerformance()

/*
This class provides access to a number of different monitors related to performance, such as memory usage, draw calls, and FPS. These are the same as the values displayed in the [b]Monitor[/b] tab in the editor's [b]Debugger[/b] panel. By using the [method get_monitor] method of this class, you can access this data from your code. [b]Note:[/b] A few of these monitors are only available in debug mode and will always return 0 when used in a release build. [b]Note:[/b] Many of these monitors are not updated in real-time, so there may be a short delay between changes.
*/
type performance struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *performance) ensureSingleton() {
	if o.initialized == true {
		return
	}
	//log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("Performance")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *performance) BaseClass() string {
	return "Performance"
}

/*
        Returns the value of one of the available monitors. You should provide one of the [enum Monitor] constants as the argument, like this: [codeblock] print(Performance.get_monitor(Performance.TIME_FPS)) # Prints the FPS to the console [/codeblock]
	Args: [{ false monitor int}], Returns: float
*/
func (o *performance) GetMonitor(monitor gdnative.Int) gdnative.Real {
	o.ensureSingleton()
	//log.Println("Calling Performance.GetMonitor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(monitor)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Performance", "get_monitor")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

// PerformanceImplementer is an interface that implements the methods
// of the Performance class.
type PerformanceImplementer interface {
	ObjectImplementer
	GetMonitor(monitor gdnative.Int) gdnative.Real
}
