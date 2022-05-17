package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
	"strings"
)

func main() {

	localIp := getLocalIpV4()
	if len(localIp) == 0 {
		localIp = GetInterfaceIpv4Addr("ens33")
	}
	f, _ := ioutil.ReadFile("./config/production.yaml")
	var (
		m = map[string]interface{}{}
	)
	yaml.Unmarshal(f, &m)
	m["IP"] = localIp
	yaml.
		f, _ = ioutil.ReadFile("./")
}

// getLocalIpV4 获取 IPV4 IP，没有则返回空
func getLocalIpV4() string {
	inters, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, inter := range inters {
		// 判断网卡是否开启，过滤本地环回接口
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			// 获取网卡下所有的地址
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//判断是否存在IPV4 IP 如果没有过滤
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

func GetInterfaceIpv4Addr(interfaceName string) string {
	var (
		ief      *net.Interface
		addrs    []net.Addr
		ipv4Addr net.IP
		err      error
	)
	if ief, err = net.InterfaceByName(interfaceName); err != nil { // get interface
		return ""
	}
	if addrs, err = ief.Addrs(); err != nil { // get addresses
		return ""
	}
	for _, addr := range addrs { // get ipv4 address
		if ipv4Addr = addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			break
		}
	}
	if ipv4Addr == nil {
		return ""
	}
	return ipv4Addr.String()
}
