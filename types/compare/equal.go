/**
* @Author: caoyongfei
* @Date: 2022/11/4
* @Description: 值的比较
**/

package compare

//func IsEqualIfcE(v1, v2 interface{}) bool {
//	v1Type := reflect.TypeOf(v1)
//
//	reflect.DeepEqual(v1, v2)
//
//
//}

func IsEqualString(v1, v2 string) bool {
	return v1 == v2
}

func IsEqualInt(v1, v2 int) bool {
	return v1 == v2
}

func IsEqualInt32(v1, v2 int32) bool {
	return v1 == v2
}

func IsEqualInt64(v1, v2 int64) bool {
	return v1 == v2
}

func IsEqualFloat32(v1, v2 float32) bool {
	return v1 == v2
}

func IsEqualFloat64(v1, v2 float64) bool {
	return v1 == v2
}

func IsEqualSlice(v1, v2 float64) bool {
	return v1 == v2
}
