package types

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Config struct {
	Add   string `json:"add"`
	Aid   string `json:"aid"`
	Host  string `json:"host"`
	ID    string `json:"id"`
	Net   string `json:"net"`
	Path  string `json:"path"`
	Port  string `json:"port"`
	TLS   string `json:"tls"`
	SNI   string `json:"sni"`
}

func Vmess() string {
	var config Config

	fmt.Print("Cole a URL VMESS: ")
	var ws_vmess string
	fmt.Scanln(&ws_vmess)
	ws_vmess = strings.Replace(ws_vmess, "vmess://", "", 1)
	data, _ := base64.StdEncoding.DecodeString(ws_vmess)
	json.Unmarshal(data, &config)


	ws_ip := config.Add
	ws_alter_id := config.Aid
	ws_host := config.Host
	ws_uuid := config.ID
	ws_net := config.Net
	ws_path := config.Path
	ws_port := config.Port
	ws_tls := config.TLS
    // ws_sni := config.SNI

	var json_vmess string

	if strings.Contains(ws_port, "443") {
		json_vmess = fmt.Sprintf(`
{
  "log": {
    "access": "none",
    "error": "/data/v2ray/run/error.log",
    "loglevel": "warning"
  },
  "inbounds": [
    {
      "port": 65535,
      "listen": "0.0.0.0",
      "tag": "proxy-inbound",
      "protocol": "dokodemo-door",
      "settings": {
        "udp": true,
        "userLevel": 8,
        "network": "tcp,udp",
        "followRedirect": true
      },
      "sniffing": {
        "enabled": false
      }
    }
  ],
  "outbounds": [
    {
      "mux": {
        "concurrency": -1,
        "enabled": false
      },
      "protocol": "vmess",
      "settings": {
        "vnext": [
          {
            "address": "%s",
            "port": %s,
            "users": [
              {
                "alterId": %s,
                "encryption": "",
                "flow": "",
                "id": "%s",
                "level": 8,
                "security": "auto"
              }
            ]
          }
        ]
      },
      "streamSettings": {
        "network": "%s",
        "security": "%s",
        "tlsSettings": {
          "allowInsecure": true,
          "serverName": "%s"
        },
        "wsSettings": {
          "headers": {
            "Host": "%s"
          },
          "path": "%s"
        }
      },
      "tag": "proxy"
    },
    {
      "protocol": "freedom",
      "settings": {},
      "tag": "direct"
    },
    {
      "protocol": "blackhole",
      "settings": {
        "response": {
          "type": "http"
        }
      },
      "tag": "block"
    }
  ],
  "policy": {
    "levels": {
      "8": {
        "connIdle": 300,
        "downlinkOnly": 1,
        "handshake": 4,
        "uplinkOnly": 1
      }
    },
    "system": {
      "statsOutboundUplink": false,
      "statsOutboundDownlink": false
    }
  },
  "routing": {
    "domainStrategy": "AsIs",
    "rules": []
  },
  "other": {}
}
`, ws_ip, ws_port, ws_alter_id, ws_uuid, ws_net, ws_tls, ws_host, ws_host, ws_path)

	} else if strings.Contains(ws_port, "8080") {
		json_vmess = fmt.Sprintf(`
{
  "log": {
    "access": "none",
    "error": "/data/v2ray/run/error.log",
    "loglevel": "none"
  },
  "inbounds": [
    {
      "port": 65535,
      "listen": "0.0.0.0",
      "tag": "proxy-inbound",
      "protocol": "dokodemo-door",
      "settings": {
        "udp": true,
        "userLevel": 8,
        "network": "tcp,udp",
        "followRedirect": true
      },
      "sniffing": {
        "enabled": false
      }
    }
  ],
  "outbounds": [
    {
      "mux": {
        "concurrency": 8,
        "enabled": false
      },
      "protocol": "vmess",
      "settings": {
        "vnext": [
          {
            "address": "%s",
            "port": %s,
            "users": [
              {
                "alterId": %s,
                "encryption": "auto",
                "flow": "",
                "id": "%s",
                "level": 8,
                "security": "auto"
              }
            ]
          }
        ]
      },
      "streamSettings": {
        "network": "%s",
        "security": "",
        "wsSettings": {
          "headers": {
            "Host": "%s"
          },
          "path": "%s"
        }
      },
      "tag": "proxy"
    },
    {
      "protocol": "freedom",
      "settings": {},
      "tag": "direct"
    },
    {
      "protocol": "blackhole",
      "settings": {
        "response": {
          "type": "http"
        }
      },
      "tag": "block"
    }
  ],
  "policy": {
    "levels": {
      "8": {
        "connIdle": 300,
        "downlinkOnly": 1,
        "handshake": 4,
        "uplinkOnly": 1
      }
    },
    "system": {
      "statsOutboundUplink": false,
      "statsOutboundDownlink": false
    }
  },
  "routing": {
    "domainStrategy": "AsIs",
    "rules": []
  },
  "other": {}
}
`, ws_ip, ws_port, ws_alter_id, ws_uuid, ws_net, ws_host, ws_path)
	}

	err := ioutil.WriteFile("/data/v2ray/config.json", []byte(json_vmess), 0644)
	if err != nil {
		panic(err)
	}


	fmt.Println("Pressione ENTER para continuar")
	fmt.Scanln()

	return json_vmess
}

