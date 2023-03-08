/**
* @Author: caoyongfei
* @Date: 2022/11/7
* @Description:
**/

package main

import (
	"fmt"
	"my_go/gotools/os/path"
	"os"
)

func main() {
	fmt.Println(path.GetBaseDirByCaller(1))
	fmt.Println(path.GetBaseDirByLookPath())
	fmt.Println(os.TempDir())
	fmt.Println(os.Executable())
	os.Getwd()
}
