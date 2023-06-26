## V2ray Magisk Module

This is a v2ray module for Magisk, and includes binaries for arm64.

### Included
 - [v2ray-core](https://github.com/v2fly/v2ray-core)
 - V2Ray service script and Android transparent proxy iptables script

### Install
 - You can download the [release](https://github.com/whyakari/v2ray_golang/blob/main/modules/magisk/v2ray/v2ray-magisk-android64.zip) installer zip file and install it via the Magisk Manager App or in KernelSU.

### Changelog
V2Ray 5.4.1 | For architecture: Custom (go1.20.2 android/arm64)

Changelog
 - removed automatic download of v2ray/releases module (unnecessary).
 - zip smaller compared to the other with 4mb difference.
 - fix where iptables didn't work in version 4.x (updated to 5.4.1).
 - fix in v2ray/scripts/v2ray.service where the whole script didn't work.

### Config
 - V2ray config file is store in /data/v2ray/config.json.


## Usage
### Normal usage ( Default and Recommended )

Manage service start / stop
```
V2Ray service is auto-run
after system boot up by default.
You can start or stop v2ray 
service by simply turn on 
or turn off the module via Magisk Manager App. 
Start service may wait about 10 second and
stop service may take effect immediately.
```
