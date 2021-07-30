package main

import (
	"flag"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/zouyx/agollo/v4"
	"github.com/zouyx/agollo/v4/env/config"
)

func ConfiFileName(namespace string) (filename string) {
	/*convert apollo config  namespacename to lcoal config filename
	 */
	filename = strings.TrimSuffix(namespace, ".txt")
	p := strings.LastIndex(filename, "_")
	filename = filename[:p] + "." + filename[p+1:]
	return

}

func SaveConfigFile(parentpath string, filename string, content string) {
	/* 默认path为当前执行路径
	 */
	f, err := os.Create(path.Join(parentpath, filename))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(content)
}

func GetPath() string {
	allargs := os.Args
	if len(allargs) == 1 {
		default_path, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
		return default_path
	}
	customize_path := flag.String("path", os.Args[0], "configfile's location")
	flag.Parse()
	return *customize_path
}

func main() {
	c := &config.AppConfig{
		AppID:          "SampleApp",
		Cluster:        "sheng-pro",
		IP:             "http://173.16.37.170:50080",
		NamespaceName:  "application_properties.txt",
		IsBackupConfig: false,
		Secret:         "",
	}

	// agollo.SetLogger(&DefaultLogger{})

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err != nil {
		panic(err)
	}

	conf := client.GetConfig(c.NamespaceName)
	content := strings.TrimPrefix(conf.GetContent(), "content=")
	content = strings.Replace(content, "\r", "", -1)
	// fmt.Println(content)
	SaveConfigFile(GetPath(), ConfiFileName(c.NamespaceName), content)
	// fmt.Println(ConfiFileName(c.NamespaceName))
}

// type DefaultLogger struct {
// }

// func (logger *DefaultLogger) Debugf(format string, params ...interface{}) {
// 	logger.Debug(format, params)
// }

// func (logger *DefaultLogger) Infof(format string, params ...interface{}) {
// 	logger.Debug(format, params)
// }

// func (logger *DefaultLogger) Warnf(format string, params ...interface{}) {
// 	logger.Debug(format, params)
// }

// func (logger *DefaultLogger) Errorf(format string, params ...interface{}) {
// 	logger.Debug(format, params)
// }

// func (logger *DefaultLogger) Debug(v ...interface{}) {
// 	fmt.Printf("%v\n", v)
// }
// func (logger *DefaultLogger) Info(v ...interface{}) {
// 	logger.Debug(v)
// }

// func (logger *DefaultLogger) Warn(v ...interface{}) {
// 	logger.Debug(v)
// }

// func (logger *DefaultLogger) Error(v ...interface{}) {
// 	logger.Debug(v)
// }
