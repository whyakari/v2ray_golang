package v2ray

import (
	"fmt"
	"os/exec"
	"github.com/AkariOficial/v2ray_golang/pkg/clean"
	"github.com/AkariOficial/v2ray_golang/pkg/sleep"
	"github.com/AkariOficial/v2ray_golang/pkg/types"
	"github.com/AkariOficial/v2ray_golang/pkg/v2ray"
)

func SubMenuV2rayFuncoes(opt string) string {
	if opt == "1" {
		exec.Command("su", "-c", "cat /data/v2ray/appid.list | grep -c '0' >> /dev/null").Run()
		exec.Command("rm", "-f", "appid.list").Run()
		exec.Command("echo", "-e", "'0' >> appid.list && su -c cp appid.list /data/v2ray/appid.list").Run()
		exec.Command("sudo", "bash", "/data/adb/modules/v2ray/scripts/v2ray.service", "start").Run()
		exec.Command("sudo", "/data/adb/modules/v2ray/scripts/v2ray.tproxy", "enable").Run()
		exec.Command("sleep", "6").Run()
		exec.Command("su", "-c", "settings put global airplane_mode_on 1").Run()
		exec.Command("su", "-c", "am broadcast -a android.intent.action.AIRPLANE_MODE > /dev/null 2>&1 && su -c settings put global airplane_mode_on 0").Run()
		exec.Command("su", "-c", "am broadcast -a android.intent.action.AIRPLANE_MODE > /dev/null 2>&1").Run()
		fmt.Println("serviço iniciado com sucesso!")
	} else if opt == "7" {
        clean.Clear()
		exec.Command("su", "-c", "cat /data/v2ray/config.json").Run()

		fmt.Scanln()
		fmt.Scanln("DE ENTER PARA VOLTAR")
	} else if opt == "4" {
		fmt.Println("instalando o tsu... aguarde...")
		exec.Command("apt", "install", "tsu", "-y").Run()
		fmt.Scanln()
    
	} else if opt == "2" {
		clean.Clear()
        
		exec.Command("sudo", "bash", "/data/adb/modules/v2ray/scripts/v2ray.service", "stop").Run()
		fmt.Println("serviço parado com sucesso, use enter para continuar.")
	} else if opt == "3" {
        clean.Clear()
		fmt.Println("1 - Configuração JSON")
		fmt.Println("2 - Configuração URL VMESS")
		fmt.Println("3 - Configuração URL VLESS")
		fmt.Println("4 - Configuração URL Trojan")
		fmt.Println("5 - Voltar")
		fmt.Println("")

		var opt string
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&opt)

		if opt == "1" {
            clean.Clear()
			fmt.Println("Configuração JSON em breve...")
            sleep.Sleep(2)
            clean.Clear()
			opt = v2ray.SubMenuV2rayFuncoes("3")
    
		} else if opt == "2" {
            clean.Clear()
            types.Vmess()
        
		} else if opt == "3" {
            clean.Clear()
			fmt.Println("VLESS em breve...")
            sleep.Sleep(2)
			clean.Clear()
		} else if opt == "4" {
			clean.Clear()
			fmt.Println("Trojan em breve...")
			sleep.Sleep(2)
            clean.Clear()
		} else if opt == "" || opt != "1" || opt != "2" || opt != "3" || opt != "4" || opt != "5" || opt != "7" {
			clean.Clear()
			fmt.Println("opção inválida! tente novamente.")
			sleep.Sleep(2)
            opt = SubMenuV2rayFuncoes("3")
		} else if opt == "5" {
			clean.Clear()
			opt = SubMenuV2ray()
		}
	}

	return opt
}

