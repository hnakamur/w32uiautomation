// +build amd64
package w32uiautomation

import "github.com/mattn/go-ole"

// type VARIANT struct {
// 	VT         uint16  //  2
// 	wReserved1 uint16  //  4
// 	wReserved2 uint16  //  6
// 	wReserved3 uint16  //  8
// 	Val        int64   // 16
// 	_          [8]byte // 24
// }

func VariantToUintptrArray(v ole.VARIANT) []uintptr {
	// Size of uintptr on 64bit system is 8
	return []uintptr{
		uintptr(v.VT), // uintptr(v.wReserved3)<<48 | uintptr(v.wReserved2)<<32 | uintptr(v.wReserved1)<<16 | uintptr(v.VT)
		uintptr(v.Val),
		uintptr(0),
	}
}
