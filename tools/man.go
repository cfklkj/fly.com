package main

import (
	"fmt"
	"strings"
	"syscall"

	"./smalltools"
)

func ColorPrint(v interface{}, i int) { //设置终端字体颜色
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Println(v)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

func main() {
	for {
		smalltools.WinColorPrintln(" smalltools | ", 1)
		fmt.Println("wait input...")
		choice := ""
		fmt.Scanln(&choice)
		switch strings.ToLower(choice) {
		case "smalltools":
			smalltools.Main()
		}
	}
}
