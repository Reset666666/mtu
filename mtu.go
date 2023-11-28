package mtu

/*
#cgo LDFLAGS: -L. -lnetwork
#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>
#include <string.h>
extern int SetAdapterMtu(const char* adapterName, uint32_t mtu);
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func SetAdapterMtu(adapterName string, mtu uint32) error {
	cAdapterName := C.CString(adapterName)
	defer C.free(unsafe.Pointer(cAdapterName))
	errno := C.SetAdapterMtu(cAdapterName, C.uint32_t(mtu))
	// 调用 strerror 函数获取错误描述
	if errno == 0 {
		return nil
	}
	errorStr := C.GoString(C.strerror(C.int(errno)))
	return fmt.Errorf(errorStr)
}
