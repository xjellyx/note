package config

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/olongfen/note/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

// 设置保存地址对对象指针
type InterfaceConfig interface {
	SetSavePath(savePath string) (err error)
	SetSavePoint(saveTarget interface{}) (err error)
	Save(newConf interface{}) error
	MonitorChange()
}

// Config
type Config struct {
	sync.RWMutex
	MonitorTime time.Duration
	lastModify  time.Time
	pathYaml    *string           // yaml配置文件保存地址
	savePoint   interface{}       //
	comments    map[string]string // 配置文件的备注信息
	silent      bool
}

// LoadConfigAndSave
func LoadConfigAndSave(configPath string, targetConfig InterfaceConfig, defaultConfig InterfaceConfig) (err error) {
	var (
		data     []byte
		fileInfo os.FileInfo
	)
	//
	if fileInfo, err = os.Stat(configPath); err != nil {
		if os.IsNotExist(err) {
			if targetConfig == nil {
				err = fmt.Errorf(`[Config] dedaultConfig undefined, "%s" error: %v`, configPath, err)
				return
			}
			// 自动创建配置文件
			if d, _err := yaml.Marshal(targetConfig); _err != nil {
				err = _err
				return
			} else if err = ioutil.WriteFile(configPath, d, 0666); err != nil {
				return
			}
			err = fmt.Errorf(`[Config] please modify "%s" and run again`, configPath)
			return err
		}
		if fileInfo.IsDir() {
			return errors.New("config path is dir")
		}

	}

	if data, err = ioutil.ReadFile(configPath); err != nil {
		return
	}

	if err = yaml.Unmarshal(data, targetConfig); err != nil {
		return
	}

	if _c, _ok := targetConfig.(InterfaceConfig); _ok == true {
		if err = _c.SetSavePath(configPath); err != nil {
			return
		}
		if err = _c.SetSavePoint(targetConfig); err != nil {
			return
		}
		if err = _c.Save(defaultConfig); err != nil {
			return
		}
		go func() {
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				_c.MonitorChange()
			}()
			wg.Wait()
		}()
	}

	return
}

// Save save config
func (c *Config) Save(newConfig interface{}) (err error) {
	c.Lock()
	defer c.Unlock()
	var (
		savePath    string
		readContent []byte
		saveContent []byte
	)
	if savePath, err = c.GetSavePath(); err != nil {
		return
	}
	// 读旧记录
	readContent, _ = ioutil.ReadFile(savePath)

	if newConfig == nil {
		newConfig = c.savePoint
	}
	if saveContent, err = yaml.Marshal(newConfig); err != nil {
		return
	} else if bytes.Equal(readContent, saveContent) == true {
		// 不重复保存
		return
	}

	// 写入记录
	if err = ioutil.WriteFile(savePath, saveContent, 0666); err != nil {
		return
	}
	if c.silent == false {
		log.Println(fmt.Sprintf("[Config] save Config to %s bytes:%d->%d",
			savePath, len(readContent), len(saveContent)))
	}

	return
}

func (c *Config) change() (err error) {
	c.Lock()
	defer c.Unlock()
	var (
		savePath    string
		readContent []byte
	)
	if savePath, err = c.GetSavePath(); err != nil {
		return
	}
	// 读旧记录
	readContent, _ = ioutil.ReadFile(savePath)

	if err = yaml.Unmarshal(readContent, c.savePoint); err != nil {
		return
	}
	if c.silent == false {
		log.Println(fmt.Sprintf("[Config] save Config to %s bytes:%d",
			savePath, len(readContent)))
	}
	return
}
func (c *Config) MonitorChange() {
	if c.MonitorTime == 0 {
		c.MonitorTime = time.Millisecond * 500
	}

	ticker := time.NewTicker(c.MonitorTime)
	for range ticker.C {
		func() {
			fileInfo, err := os.Stat(*c.pathYaml)
			if err != nil {
				if os.IsNotExist(err) {
					log.Println(err)
				}

				if fileInfo.IsDir() {
					log.Println(err)
				}
				log.Println("get file stat error: ", err)
				return
			}

			if fileInfo.ModTime().Equal(c.lastModify) {
				return
			}

			if err = c.change(); err == nil {
				c.lastModify = fileInfo.ModTime()
			} else {
				log.Errorln("[MonitorChange]", err)
			}
		}()
	}

}

// GetSavePath 取保存地址
func (c *Config) GetSavePath() (ret string, err error) {
	if c.pathYaml == nil {
		err = errors.New("param invalid")
		return
	} else {
		ret = *c.pathYaml
	}
	return
}

// SetSavePath 设置保存地址
func (c *Config) SetSavePath(savePath string) (err error) {
	if len(savePath) == 0 {
		c.pathYaml = nil
		err = errors.New("param invalid")
		return
	} else {
		c.pathYaml = &savePath
	}
	return
}

// SetSavePoint 设置要保存的对象
func (c *Config) SetSavePoint(saveTarget interface{}) (err error) {
	c.savePoint = saveTarget
	return
}
