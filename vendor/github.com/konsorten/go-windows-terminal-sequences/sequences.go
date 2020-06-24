// +build windows

package sequences

import (
	"syscall"
<<<<<<< HEAD
=======
	"unsafe"
>>>>>>> dev
)

var (
	kernel32Dll    *syscall.LazyDLL  = syscall.NewLazyDLL("Kernel32.dll")
	setConsoleMode *syscall.LazyProc = kernel32Dll.NewProc("SetConsoleMode")
)

func EnableVirtualTerminalProcessing(stream syscall.Handle, enable bool) error {
	const ENABLE_VIRTUAL_TERMINAL_PROCESSING uint32 = 0x4

	var mode uint32
	err := syscall.GetConsoleMode(syscall.Stdout, &mode)
	if err != nil {
		return err
	}

	if enable {
		mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING
	} else {
		mode &^= ENABLE_VIRTUAL_TERMINAL_PROCESSING
	}

<<<<<<< HEAD
	ret, _, err := setConsoleMode.Call(uintptr(stream), uintptr(mode))
=======
	ret, _, err := setConsoleMode.Call(uintptr(unsafe.Pointer(stream)), uintptr(mode))
>>>>>>> dev
	if ret == 0 {
		return err
	}

	return nil
}
