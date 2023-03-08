/**
* @Author: caoyongfei
* @Date: 2022/11/4
* @Description:
**/

package array

/**
* @FuncName: FindIndexString
* @Description: 寻找目标在切片中的索引下标。 string类型
* @Params <arr []string>:
* @Params <dest string>:
* @Return <index int>: -1 表示没找到
**/
func FindIndexString(arr []string, dest string) (index int) {
	for idx, v := range arr {
		if v == dest {
			return idx
		}
	}

	return -1
}

func FindIndexInt(arr []int, dest int) (index int) {
	for idx, v := range arr {
		if v == dest {
			return idx
		}
	}

	return -1
}

func FindIndexInt32(arr []int32, dest int32) (index int) {
	for idx, v := range arr {
		if v == dest {
			return idx
		}
	}

	return -1
}

func FindIndexInt64(arr []int64, dest int64) (index int) {
	for idx, v := range arr {
		if v == dest {
			return idx
		}
	}

	return -1
}

