```go
package main

import (
	"fmt"

	"github.com/erning/mobiledb"
)

func main() {
	info, err := mobiledb.GetMobileInfo("13916020000")
	fmt.Println(info, err)
	info, err = mobiledb.GetMobileInfo("13906900000")
	fmt.Println(info, err)
	info, err = mobiledb.GetMobileInfo("13906010000")
	fmt.Println(info, err)
	info, err = mobiledb.GetMobileInfo("14599990000")
	fmt.Println(info, err)
	info, err = mobiledb.GetMobileInfo("11199990000")
	fmt.Println(info, err)
}
```

```
&{上海 上海 移动 21 200000 中国移动 GSM} <nil>
&{福建 福州 移动 591 350000 中国移动 GSM} <nil>
&{福建 厦门 移动 592 361000 中国移动 GSM} <nil>
<nil> Unknown number
<nil> Unknown prefix
```
