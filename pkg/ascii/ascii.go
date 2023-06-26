package ascii

import (
	"fmt"
	"os"
	"github.com/whyakari/v2ray_golang/pkg/clean"
)

func Ascii_Text(opt string) string {
    clean.Clear()

    if "sair" == opt || "3" == opt {
        fmt.Println(`
 ___
  | |_   _. ._  |    \_/ _
  | | | (_| | | |<    | (_) |_|
        `)
    }
    
    os.Exit(0)

    return opt

}
