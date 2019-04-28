# hotkey

适用于window的热键

## example

```go
package main

import (
	"fmt"
	"time"

	"github.com/jinzhongmin/hotkey"
)

func main() {

	hk := hotkey.New()
	hk.Bind(hotkey.ModCtrl+hotkey.ModAlt, 'S', func() {
		fmt.Println("hello")
	})
	hk.Bind(hotkey.ModNone, hotkey.VK_F8, func() {
		fmt.Println("stop")
		hk.Stop()
	})
	hk.BindContinue(hotkey.ModNone, 'C', func() {
		fmt.Println("continue")
	})

	go hk.Listen()

	fmt.Println("Ctrl+Alt+S Say hello")
	fmt.Println("F8 Stop")
	fmt.Println("C Continue")
	for {
		time.Sleep(time.Microsecond * 10)
	}
}

```

