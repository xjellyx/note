package main

import (
	"bytes"
	"fmt"
	"github.com/olongfen/note/config"
	"github.com/olongfen/note/log"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode"
)

type Compose struct {
	config.Config `yaml:"-"`
	Version       string                 `json:"version" yaml:"version"`
	Services      map[string]interface{} `json:"services" yaml:"services"`
}

func main() {
	var (
		f   string
		err error
	)
	f1 := "./demo.sh"
	d, _ := ioutil.ReadFile(f1)

	if f, err = RunContainer("dsadasdasdas", "chart1.py", d, []string{}); err != nil {
		log.Warnln(err)
	}
	println(f)
	time.Sleep(time.Second)
	if err = Output(f); err != nil {
		log.Warnln(err)
	}
	//if err = Down(f);err!=nil{
	//	log.Warnln(err)
	//}

}

func RunContainer(accessToken string, pyFilename string, pyCode []byte, pkgs []string) (ret string, err error) {
	dir := fmt.Sprintf("./public/user/%s", accessToken)
	ymlFileDir := fmt.Sprintf("%s/%s", dir, "docker-compose.yml")
	pyFileDir := fmt.Sprintf("%s/%s", dir, pyFilename)
	pipCmds := ""
	defer func() {
		ret = ymlFileDir
	}()
	for i, v := range pkgs {
		if i == 0 {
			pipCmds += fmt.Sprintf("pip install %s", v)
		} else {
			pipCmds += fmt.Sprintf(" && pip install %s", v)
		}
	}
	dockerCmd := ""
	if len(pipCmds) != 0 {
		dockerCmd = fmt.Sprintf(`/bin/bash -c " %s && python  %s"`, pipCmds, pyFilename)
	} else {
		dockerCmd = fmt.Sprintf(`/bin/bash -c " python  %s"`, pyFilename)
	}
	println(dockerCmd)
	c := new(Compose)
	c.Version = "3.8"
	c.Services = map[string]interface{}{
		"demo": map[string]interface{}{
			"image":          "registry.cn-hangzhou.aliyuncs.com/olongfen/python:3",
			"network_mode":   "host",
			"container_name": "py_" + accessToken,
			"working_dir":    "/app",
			"volumes": []string{
				fmt.Sprintf(`./%s:/app/%s`, pyFilename, pyFilename),
				fmt.Sprintf(`./%s:/app/%s`, "config.ini", "config.ini"),
			},
			"command": dockerCmd,
			"deploy": map[string]interface{}{
				"resources": map[string]interface{}{
					"limits": map[string]interface{}{
						"cpus":   "0.2",
						"memory": "1000M",
					},
				},
			},
		},
	}
	if _, err = os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(dir, os.ModePerm); err != nil {
				return "", err
			}
		}
	}

	var (
		file *os.File
	)
	if file, err = os.Create(pyFileDir); err != nil {

		return
	}
	if _, err = file.Write(pyCode); err != nil {
		return
	}
	file.Close()

	config.LoadConfigAndSave(ymlFileDir, c, c)

	cmd := exec.Command("docker-compose", "--compatibility", "-f", ymlFileDir, "up", "-d")
	if err = cmd.Run(); err != nil {
		return
	}
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	if err = cmd.Run(); err != nil {
		return
	}
	cmd = exec.Command("docker-compose", "--compatibility", "-f", ymlFileDir, "ps")
	println(stdout.String())
	cmd.Wait()

	return
}

func Output(ymlFileDir string) (err error) {
	cmd := exec.Command("docker-compose", "--compatibility", "-f", ymlFileDir, "logs")
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	if err = cmd.Run(); err != nil {
		return
	}
	println(stdout.String())
	return
}

func Down(ymlFileDir string) (err error) {
	cmd := exec.Command("docker-compose", "--compatibility", "-f", ymlFileDir, "down")
	if err = cmd.Run(); err != nil {
		return
	}
	return os.Remove(ymlFileDir)
}

func get(spec string) (ret []string) {
	var (
		data  = strings.Split(spec, " ")
		title = ""
	)
	for i, v := range data {
		switch i {
		case 0:
			title = "minute"
		case 1:
			title = "hour"
		case 2:
			title = "day"
		case 3:
			title = "month"
		case 4:
			title = "week"
		default:
			return
		}
		var (
			mean string
			and  string
		)
		if strings.Contains(v, ",") {
			if !strings.Contains(v, "-") && !strings.Contains(v, "/") {
				mean = " every " + v + " " + title + " run "
			} else {
				_data := strings.Split(v, ",")
				if len(_data) >= 2 {
					and = "and"
				}
				for _i, _v := range _data {
					if strings.Contains(_v, "-") && !strings.Contains(_v, "/") {
						if _i == 0 {
							mean += " every " + _v + " run interval 1 " + title + " "
						} else {
							mean += and + " every " + _v + " run interval 1 " + title + " "
						}
					} else if strings.Contains(_v, "-") && strings.Contains(_v, "/") {
						arr := strings.Split(_v, "/")
						if _i == 0 {
							mean += " every " + arr[0] + " " + title + " run interval " + arr[1] + " " + title + " "
						} else {
							mean += and + " every " + arr[0] + " " + title + " run interval " + arr[1] + " " + title + " "
						}
					}
				}
			}
		} else {
			if strings.Contains(v, "-") && !strings.Contains(v, "/") {
				mean = " every " + v + " run interval 1 " + title + " "
			} else if strings.Contains(v, "/") {
				arr := strings.Split(v, "/")
				if strings.Contains(v, "-") {
					mean = " every " + arr[0] + " " + title + " run interval " + arr[1] + " " + title + " "
				} else {
					mean = " every " + arr[0] + " begin " + title + " run interval " + arr[1] + " " + title + " "
				}

			} else {
				if v == "*" {
					v = ""
				} else if v == "?" {
					continue
				}
				mean += " every " + v + " " + title + " run "
			}
		}
		ret = append(ret, mean)
	}
	return
}

// SQLColumnToHumpStyle sql转换成驼峰模式
func SQLColumnToHumpStyle(in string) (ret string) {
	for i := 0; i < len(in); i++ {
		if i > 0 && in[i-1] == '_' && in[i] != '_' {
			s := strings.ToUpper(string(in[i]))
			ret += s
		} else if in[i] == '_' {
			continue
		} else {
			ret += string(in[i])
		}
	}
	return
}

// HumpToSQLColumnStyle 驼峰转sql
func HumpToSQLColumnStyle(in string) (ret string) {
	for i := 0; i < len(in); i++ {
		if unicode.IsUpper(rune(in[i])) {
			ret += "_" + strings.ToLower(string(in[i]))
		} else {
			ret += string(in[i])
		}
	}
	return
}
