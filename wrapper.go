package gonix

/*
#cgo CFLAGS: -I/home/didi/workspace/cnix-install/include
#cgo LDFLAGS: -L/home/didi/workspace/cnix-install/lib -lcnix

#include <stdlib.h>
#include <cnix/cnix.h>


typedef void* cgo_nix_handle;





*/
import "C"

import(
    "unsafe"
)



func NixHandleCreate() C.cnix_handle_t {
    h := C.cnix_handle_new()
    return h
}


func NixStorePath(h C.cnix_handle_t) string {
    s := C.cnix_store_path(h)
    res := C.GoString(s)
    C.free(unsafe.Pointer(s))
    return res
}


func NixHandleDestroy(h C.cnix_handle_t)  {
    C.cnix_handle_delete(h)
}
