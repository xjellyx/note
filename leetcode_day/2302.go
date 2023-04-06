package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// decodeMessage 02-06
func decodeMessage(key string, message string) string {
	var (
		arr      = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		keyMap   = make(map[string]string)
		isExists = make(map[string]bool)
		item     int
		res      string
	)
	for _, v := range key {
		if v == ' ' {
			continue
		}
		if _, ok := isExists[string(v)]; !ok {
			keyMap[string(v)] = arr[item]
			isExists[string(v)] = true
			item++
		} else {
			continue
		}
	}
	for _, v := range message {
		if v == ' ' {
			res += " "
		} else {
			res += keyMap[string(v)]
		}
	}
	return res
}

// alertNames 02-07
func alertNames(keyName []string, keyTime []string) []string {
	var (
		users = map[string][]string{}
		res   []string
	)
	for i, v := range keyName {
		users[v] = append(users[v], keyTime[i])
	}
	for k, v := range users {
		if len(v) < 3 {
			continue
		}
		for i := 0; i < len(v)-2; i++ {
			ts, _ := time.Parse("15:04", v[i+2])
			tn, _ := time.Parse("15:04", v[i])
			if ts.Unix()-tn.Unix() <= 3600 && (!strings.Contains(v[i+2], "00:") && !strings.Contains(v[i], "23:")) {
				res = append(res, k)
				break
			}
		}

	}
	sort.Strings(res)
	return res
}

// removeSubfolders 08
func removeSubfolders(folder []string) []string {
	var (
		res []string
	)
	if len(folder) == 0 {
		return res
	}
	first := folder[0]
	res = append(res, first)
	for i := 1; i < len(folder); i++ {
		if strings.HasPrefix(folder[i], first+"/") {
			continue
		} else {
			first = folder[i]
			res = append(res, first)
		}
	}
	return res
}

// 09
type AuthenticationManager struct {
	expireTime int
	tokenMap   map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{expireTime: timeToLive,
		tokenMap: map[string]int{},
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokenMap[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	v := this.tokenMap[tokenId]
	if v+this.expireTime <= currentTime {
		delete(this.tokenMap, tokenId)
	} else {
		this.tokenMap[tokenId] = currentTime
	}

}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var (
		res int
	)
	for _, v := range this.tokenMap {
		if v+this.expireTime > currentTime {
			res++
		}
	}
	return res
}

func main() {
	var (
		keyName = []string{"/ah/al/am", "/ah/al"}
	)
	sort.Strings(keyName)
	fmt.Println(keyName)
	fmt.Println(removeSubfolders(keyName))
}
