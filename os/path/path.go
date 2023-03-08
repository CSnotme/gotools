/**
* @Author: caoyongfei
* @Date: 2022/11/21
* @Description: path的一些应用
**/

package path

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)


/**
* @FuncName: GetAppBaseDir
* @Description: 获取项目根目录(兼容go run方式)
*           目前可以获取项目目录的方式有3种:  这里设二进制文件路径   /Users/test/demo/app
*			1.os.Getwd(), 获取的是执行go程序的目录路径，除非在项目目录下执行，否则获取的路径不是项目目录
*				在/Users/test/demo 目录下执行 ./app      => /Users/test/demo
*				在/Users/test      目录下执行 .demo/app  => /Users/test
*			2.os.Args[0], 通过filePath, _ := exec.LookPath(os.Args[0])获取到执行的文件路径, 但如果go run运行情况下获取的是生成的临时文件
*				在/Users/test/demo 目录下执行 ./app           => /Users/test/demo/app
*				在/Users/test      目录下执行 .demo/app       => /Users/test/demo/app
*				在/Users/test/demo 目录下执行 go run main.go  => /var/folders/3g/f2sh8sgs5ls_z62npf80v69w0000gn/T/go-build1053443992/b001/exe
*			3.runtime.Caller(), 获取程序执行的堆栈信息，如果清楚调用层级通过设置skip=x，就能获取到项目的执行路径
*    		目前来看,方式3更为通用，当前函数通过方式3实现, 可以兼容go run形式。1、2的方式也可用，只是需要做一些约定，或者在提前配置中设定好根路径。
* @Params <skip int>:
* @Return <string>:
**/
func GetBaseDirByCaller(skip int) string {
	// 这个skip看调用层级设定， 0表示当前文件
	_, file, _, _ := runtime.Caller(skip)
	baseDir := path.Dir(file)
	return baseDir
}

/**
* @FuncName: GetBaseDirByLookPath
* @Description: 通过filePath, _ := exec.LookPath(os.Args[0])获取到执行的文件路径, 但如果go run运行情况下获取的是生成的临时文件
* @Return <string>:
**/
func GetBaseDirByLookPath() (string, error) {
	// 获取执行的二进制文件路径
	filePath, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}

	// 获取二进制文件绝对路径
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	// 获取baseDir
	baseDir := path.Dir(absFilePath)
	return baseDir, nil
}

