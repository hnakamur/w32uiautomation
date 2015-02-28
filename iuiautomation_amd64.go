// +build amd64
package w32uiautomation

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/mattn/go-ole"
)

func createPropertyCondition(aut *IUIAutomation, propertyId PROPERTYID, value ole.VARIANT) (*IUIAutomationCondition, error) {
	v := VariantToUintptrArray(value)
	log.Printf("createPropertyCondition aut.VTable().CreatePropertyCondition=%x, aut=%x, propertyId=%d", aut.VTable().CreatePropertyCondition, aut, propertyId)
	log.Printf("createPropertyCondition v[0]=%x, v[1]=%d, v[2]=%x", v[0], v[1], v[2])
	condition := IUIAutomationCondition{}
	log.Printf("createPropertyCondition condition=%v", condition)
	var err error
	hr, _, _ := syscall.Syscall6(
		aut.VTable().CreatePropertyCondition,
		6,
		uintptr(unsafe.Pointer(aut)),
		uintptr(propertyId),
		v[0],
		v[1],
		v[2],
		uintptr(unsafe.Pointer(&condition)))
	// I got an Exception at the above line on IE11 - Win10 modern.IE VirtualBox.
	//
	// 2015/02/28 02:07:28 createPropertyCondition aut.VTable().CreatePropertyCondition=7ffd10a76f60, aut=&{{7ffd10b2ab90}}, propertyId=30005
	// 2015/02/28 02:07:28 createPropertyCondition v[0]=8, v[1]=3049688, v[2]=0
	// 2015/02/28 02:07:28 createPropertyCondition condition={{<nil>}}
	// Exception 0xc0000005 0x0 0x8 0x7ffd10a76f9f
	// PC=0x7ffd10a76f9f
	// signal arrived during cgo execution
	//
	// github.com/hnakamur/w32uiautomation.createPropertyCondition(0xbd6b10, 0x7535, 0x8, 0x2e88d8, 0x0, 0x8, 0x0, 0x0)
	//         C:/Users/IEUser/go/src/github.com/hnakamur/w32uiautomation/iuiautomation_amd64.go:27 +0x608
	// github.com/hnakamur/w32uiautomation.(*IUIAutomation).CreatePropertyCondition(0xbd6b10, 0x7535, 0x8, 0x2e88d8, 0x0, 0x447681, 0x0, 0x0)
	//         C:/Users/IEUser/go/src/github.com/hnakamur/w32uiautomation/iuiautomation.go:110 +0x61
	// github.com/hnakamur/moderniejapanizer.findChildElementByName(0xbd6b10, 0xbde1f0, 0x6f04b0, 0x8, 0x0, 0x0, 0x0)
	//         C:/Users/IEUser/go/src/github.com/hnakamur/moderniejapanizer/imeja.go:129 +0xb3
	// github.com/hnakamur/moderniejapanizer.switchInputMethodJaWin8(0x0, 0x0)
	//         C:/Users/IEUser/go/src/github.com/hnakamur/moderniejapanizer/imeja.go:39 +0x4ba
	// github.com/hnakamur/moderniejapanizer.SwitchInputMethodJa(0x406, 0x0, 0x0)
	//         C:/Users/IEUser/go/src/github.com/hnakamur/moderniejapanizer/imeja.go:15 +0x4a
	// main.main()
	//         C:/Users/IEUser/go/src/github.com/hnakamur/w32uiautomation/examples/switchimeja/main.go:21 +0xa5
	// rax     0x8fe00
	// rbx     0x8
	// rcx     0x7ffd10b662a8
	// rdx     0x7ffd10b2cb60
	// rdi     0x2e88d8
	// rsi     0x7535
	// rbp     0xbd6b10
	// rsp     0x8fdb0
	// r8      0x8fde0
	// r9      0x0
	// r10     0x7ffd10b2cb60
	// r11     0x0
	// r12     0x0
	// r13     0x0
	// r14     0x0
	// r15     0x10
	// rip     0x7ffd10a76f9f
	// rflags  0x7ffd00010246
	// cs      0x33
	// fs      0x53
	// gs      0x2b
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return &condition, err
}
