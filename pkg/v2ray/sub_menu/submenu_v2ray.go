package submenu

import (
	"fmt"
	"github.com/whyakari/v2ray_golang/pkg/clean"
)

func SubMenuV2ray() string {
    clean.Clear()
	fmt.Println("V2RAY Core Magisk")
	fmt.Println("")
	fmt.Println("You need to have installed the magisk module:")
	fmt.Println("V2RAY For Android")
	fmt.Println("")
	fmt.Println("1 - Start services V2RAY")
	fmt.Println("2 - Stop services V2RAY")
	fmt.Println("3 - Change Configuration V2RAY")
	fmt.Println("4 - Install sudo")
	fmt.Println("5 - Fix install of the Module")
	fmt.Println("6 - Routing")
	fmt.Println("7 - View conf file v2ray")
	fmt.Println("0 - Back")
	fmt.Println("")

	var opts string
	fmt.Print("console ")
	fmt.Scanln(&opts)
	return opts
}
