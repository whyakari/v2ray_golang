package submenu

import (
	"fmt"

	"github.com/AkariOficial/v2ray_golang/pkg/clean"
)

func SubMenuV2ray() string {
    clean.Clear()
	fmt.Println("V2RAY Core Magisk")
	fmt.Println("")
	fmt.Println("Necessario ter instalado o modulo do magisk:")
	fmt.Println("V2RAY For Android")
	fmt.Println("")
	fmt.Println("1 - Iniciar serviços V2RAY")
	fmt.Println("2 - Parar serviços V2RAY")
	fmt.Println("3 - Mudar configuração V2RAY")
	fmt.Println("4 - Instalar sudo")
	fmt.Println("5 - Corrigir instalação do Modulo")
	fmt.Println("6 - Roteamento")
	fmt.Println("7 - Visualizar arquivo")
	fmt.Println("0 - Voltar")
	fmt.Println("")

	var opts string
	fmt.Print("Escolha sua opção: ")
	fmt.Scanln(&opts)
	return opts
}
