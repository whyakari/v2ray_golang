package v2ray

import (
	"fmt"
	"os"

	"github.com/AkariOficial/v2ray_golang/pkg/v2ray"
)

type Option struct {
    value string
}

func showMenu() {

    fmt.Println("--------------------------")
    fmt.Println("        Welcome!          ")
    fmt.Println("   Chooice your option    ")
    fmt.Println("--------------------------")
    fmt.Println("")
    fmt.Println("(1) - v2ray")
    fmt.Println("(2) - xray")
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
        opts := v2ray.SubMenuV2ray()
        opts = v2ray.SubMenuV2rayFuncoes(opts)

    } else if option.value == "2" {
		fmt.Println("nao disponivel")
        v2ray.Menu()
	} else if option.value == "0" {
		os.Exit(0)
	} else {
		fmt.Println("Opção inválida! Tente novamente.")
	}

    return option.value, nil
}

