package smalltools

/* https://www.cnblogs.com/askpad/p/8361492.html
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     tofuutils
 * @date        2018-01-25 19:19
 */

// //使用方法，直接调用即可输出带颜色的文本  --https://www.qiansw.com/golang-windows-setconsoletextattribute.html
// ColorPrint("[OK];", 2|8) //亮绿色

import (
	"fmt"
	"strings"
	"syscall"
)

//绿色字体，modifier里，第一个控制闪烁，第二个控制下划线
func Green(str string, modifier ...interface{}) string {
	return cliColorRender(str, 32, 0, modifier...)
}

//淡绿
func LightGreen(str string, modifier ...interface{}) string {
	return cliColorRender(str, 32, 1, modifier...)
}

//青色/蓝绿色
func Cyan(str string, modifier ...interface{}) string {
	return cliColorRender(str, 36, 0, modifier...)
}

//淡青色
func LightCyan(str string, modifier ...interface{}) string {
	return cliColorRender(str, 36, 1, modifier...)
}

//红字体
func Red(str string, modifier ...interface{}) string {
	return cliColorRender(str, 31, 0, modifier...)
}

//淡红色
func LightRed(str string, modifier ...interface{}) string {
	return cliColorRender(str, 31, 1, modifier...)
}

//黄色字体
func Yellow(str string, modifier ...interface{}) string {
	return cliColorRender(str, 33, 0, modifier...)
}

//黑色
func Black(str string, modifier ...interface{}) string {
	return cliColorRender(str, 30, 0, modifier...)
}

//深灰色
func DarkGray(str string, modifier ...interface{}) string {
	return cliColorRender(str, 30, 1, modifier...)
}

//浅灰色
func LightGray(str string, modifier ...interface{}) string {
	return cliColorRender(str, 37, 0, modifier...)
}

//白色
func White(str string, modifier ...interface{}) string {
	return cliColorRender(str, 37, 1, modifier...)
}

//蓝色
func Blue(str string, modifier ...interface{}) string {
	return cliColorRender(str, 34, 0, modifier...)
}

//淡蓝
func LightBlue(str string, modifier ...interface{}) string {
	return cliColorRender(str, 34, 1, modifier...)
}

//紫色
func Purple(str string, modifier ...interface{}) string {
	return cliColorRender(str, 35, 0, modifier...)
}

//淡紫色
func LightPurple(v interface{}, modifier ...interface{}) string {
	//return cliColorRender(v, 35, 1, modifier...)
	return ""
}

//棕色
func Brown(str string, modifier ...interface{}) string {
	return cliColorRender(str, 33, 0, modifier...)
}

func cliColorRender(v string, color int, weight int, extraArgs ...interface{}) string {
	// //闪烁效果
	isBlink := 0
	// if len(extraArgs) > 0 {
	// 	isBlink, _ = MakeItemElem(extraArgs[0]).ToInt()
	// }
	// //下划线效果
	isUnderLine := 0
	// if len(extraArgs) > 1 {
	// 	isUnderLine, _ = MakeItemElem(extraArgs[1]).ToInt()
	// }
	var mo []string
	if isBlink > 0 {
		mo = append(mo, "05")
	}
	if isUnderLine > 0 {
		mo = append(mo, "04")
	}
	if weight > 0 {
		mo = append(mo, fmt.Sprintf("%d", weight))
	}
	if len(mo) <= 0 {
		mo = append(mo, "0")
	}
	return fmt.Sprintf("\033[%s;%dm"+v+"\033[0m", strings.Join(mo, ";"), color)
}

func WinColorPrintln(v interface{}, i int) { //设置终端字体颜色
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Println(v)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}
