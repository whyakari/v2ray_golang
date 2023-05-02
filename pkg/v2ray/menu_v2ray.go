package v2ray

import (
    "fmt"
)

type Option struct {
    value string
}

func showMenu() {

    fmt.Println("--------------------------")
    fmt.Println("       Bem-Vindo          ")
    fmt.Println("   Escolha sua opção      ")
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
Menu do v2ray
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

    fmt.Println(option.value)
    return option.value, nil
}

