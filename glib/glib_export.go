package glib

// #cgo pkg-config: gio-2.0
// #include <gio/gio.h>
import "C"
import "unsafe"

//export goAsyncReadyCallbacks
func goAsyncReadyCallbacks(sourceObject *C.GObject, res *C.GAsyncResult, userData C.gpointer) {
	id := int(uintptr(userData))

	asyncReadyCallbackRegistry.Lock()
	r := asyncReadyCallbackRegistry.m[id]
	//delete(asyncReadyCallbackRegistry.m, id)
	asyncReadyCallbackRegistry.Unlock()

	var source *Object
	if sourceObject != nil {
		source = wrapObject(unsafe.Pointer(sourceObject))
	}

	r.fn(source, wrapAsyncResult(wrapObject(unsafe.Pointer(res))), r.userData)
}
