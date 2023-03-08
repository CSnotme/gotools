/**
* @Author: caoyongfei
* @Date: 2022/11/4
* @Description:
**/

package compare

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsEqualFloat32(t *testing.T) {
	a := []string{"1", "2"}
	b := []string{"2", "1"}

	ret := reflect.DeepEqual(a, b)
	fmt.Println(ret)
}