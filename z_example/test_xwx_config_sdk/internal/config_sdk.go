package internal

import (
	"codeup.aliyun.com/61e54b0e0bb300d827e1ae27/xue/bigclass_jichujiagou_common/config-sdk-go"
	"fmt"
	"time"
)

func UseCase() {
	config_sdk.SetClient("44c388c3ae65d930252fa0da756c1785", "ea2a85634612372b82e6f92da138b4af", 10*time.Second, 10*time.Second)
	config_sdk.SetAppId("224")
	config_sdk.SetCluster("1")
	data, err := config_sdk.GetConfig("test-111")
	fmt.Println(data, err)
}
