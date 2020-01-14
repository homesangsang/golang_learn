package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

// MessageBoxW函数的参数枚举值
const (
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_HELP              = 0x00004000
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_RETRYCANCEL       = 0x00000005
	MB_YESNO             = 0x00000004
	MB_YESNOCANCEL       = 0x00000003
)

func main() {
	// 调用user32.dll
	var mod = syscall.NewLazyDLL("user32.dll")
	// 定义要调用的方法user32.dll中的messageBoxW函数
	var proc = mod.NewProc("MessageBoxW")

	// 设置弹窗内容
	content, _ := syscall.UTF16PtrFromString("中文This test is Done.")
	// 设置弹窗标题
	title, _ := syscall.UTF16PtrFromString("Done Title")

	// 调用
	ret, _, _ := proc.Call(0,
		uintptr(unsafe.Pointer(content)),
		uintptr(unsafe.Pointer(title)),
		// MB_YESNO表示弹出框有一个yes和一个no选项
		uintptr(MB_YESNO))
	fmt.Printf("Return: %d\n", ret)
}
