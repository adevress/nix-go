package nix

/*
#cgo CFLAGS: -I/home/didi/workspace/cnix-install/include
#cgo LDFLAGS: -L/home/didi/workspace/cnix-install/lib -lcnix

#include <stdlib.h>
#include <cnix/cnix.h>


typedef void* cgo_nix_handle;

const int cnix_error_ok_code = CNIX_ERROR_OK;



*/
import "C"

import(
    "unsafe"
    "errors"    
)


type NixHandle struct {
    h C.cnix_handle_t
}


func nixCheckError() error {
    if(C.cnix_error_code() == C.cnix_error_ok_code ){
        return nil
    }
    
    s := C.cnix_error_string()
    res := C.GoString(s)
    
    C.cnix_error_reset()
    
    return errors.New(res)
}

func New() (NixHandle, error) {
    intern_handle := C.cnix_handle_new()
    handle := NixHandle{ h : intern_handle}
    err := nixCheckError()
    return handle, err
}


func (handle NixHandle) StorePath() string {
    s := C.cnix_store_path(handle.h)
    res := C.GoString(s)
    C.free(unsafe.Pointer(s))
    return res
}


func Delete(handle NixHandle)  {
    handle.h = nil
}
