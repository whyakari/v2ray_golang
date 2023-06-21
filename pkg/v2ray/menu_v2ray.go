package v2ray

import (
	"fmt"
	"os"

	"github.com/whyakari/v2ray_golang/pkg/clean"
	"github.com/whyakari/v2ray_golang/pkg/input"
	menufunctions "github.com/whyakari/v2ray_golang/pkg/v2ray/menu_functions"
	submenu "github.com/whyakari/v2ray_golang/pkg/v2ray/sub_menu"
)

type Option struct {
    value string
}


func showMenu() {

    clean.Clear()

    fmt.Println("--------------------------")
    fmt.Println("        Welcome!          ")
    fmt.Println("   Chooice your option    ")
    fmt.Println("--------------------------")
    fmt.Println("")
    fmt.Println("(1) - v2ray")
    fmt.Println("(2) - xray")
    fmt.Println("(3) - exit")
    fmt.Println("")
}

func readOption() (Option, error) {
    var opt string
    fmt.Print("Your option: ")
    _, err := fmt.Scanln(&opt)
    if err != nil {
        return Option{}, err
    }
    return Option{value: opt}, nil
}

func validateOption(option Option) error {
    if option.value == "" {
        return fmt.Errorf("error: option is empty")
    }
    return nil
}

/*
Menu of the v2ray
*/
func Menu() (string, error) {
    showMenu()
    option, err := readOption()
    if err != nil {
        return "", err
    }
    if err := validateOption(option); err != nil {
        return "", err
    }

    if option.value == "1" {
        opts := submenu.SubMenuV2ray()
        opts = menufunctions.SubMenuV2rayFuncoes(opts)

    } else if option.value == "2" {
		fmt.Println("nao disponivel")
        input.Input()
        clean.Clear()
        Menu()
	} else if option.value == "3"{
		fmt.Println("Bye!")
        os.Exit(0)
    } else {
        fmt.Println("Option invalid! try again.")
    }

    return option.value, nil
}

