package xicon

import (
	"fmt"
	"github.com/gogf/gf/v2/util/grand"
)

// GenIcon  随机生成头像
// 国内网络不能用
func GenIcon() string {
	switch grand.N(1, 100) % 5 {
	case 1:
		return fmt.Sprintf("https://www.gravatar.com/avatar/%d?d=wavatar&f=y", grand.Intn(10000000000))
	case 2:
		return fmt.Sprintf("https://www.gravatar.com/avatar/%d?d=identicon&f=y", grand.Intn(10000000000))
	case 3:
		return fmt.Sprintf("https://www.gravatar.com/avatar/%d?d=monsterid&f=y", grand.Intn(10000000000))
	case 4:
		return fmt.Sprintf("https://www.gravatar.com/avatar/%d?d=retro&f=y", grand.Intn(10000000000))
	default:
		return fmt.Sprintf("https://www.gravatar.com/avatar/%d?d=robohash&f=y", grand.Intn(10000000000))
	}
}
