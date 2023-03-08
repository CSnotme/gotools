/**
* @Author: caoyongfei
* @Date: 2022/11/21
* @Description: golang  runtime.Caller使用
**/

package runt

import (
	"fmt"
	"path"
	"runtime"
)

/**
* @FuncName: GetCallerInfo
* @Description: 测试一下runtime.Caller的使用
*		main函数中调用当前函数，不同的skip得到不同的结果, 可以看到skip每次+1,打印更外层调用者的堆栈信息， runtime.goexit为最外层了
*		fmt.Println(runt.GetCallerInfo(0))   // 0层skip  stdout: FuncName:my_go/gotools/os/runt.GetCallerInfo, file:caller.go, line:23
*		fmt.Println(runt.GetCallerInfo(1))   // 1层skip  stdout: FuncName:main.main, file:main.go, line:16
*		fmt.Println(runt.GetCallerInfo(2))   // 2层skip  stdout: FuncName:runtime.main, file:proc.go, line:225
*		fmt.Println(runt.GetCallerInfo(3))   // 3层skip  stdout: FuncName:runtime.goexit, file:asm_arm64.s, line:1130
*		fmt.Println(runt.GetCallerInfo(4))   // 4层skip  stdout: runtime.Caller() failed
* @Params <skip int>:
* @Return <info string>:
**/
func GetCallerInfo(skip int) (info string) {
	// pc:函数栈帧、 file:文件名、 lineNo:行号、 ok:是否获取到信息
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		info = "runtime.Caller() failed"
		return
	}

	// 传入函数栈帧，获取函数名称
	funcName := runtime.FuncForPC(pc).Name()

	// Base函数返回路径的最后一个元素
	fileName := path.Base(file)
	return fmt.Sprintf("FuncName:%s, file:%s, line:%d ", funcName, fileName, lineNo)
}
