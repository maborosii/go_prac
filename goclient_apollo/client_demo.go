package main

import (
	"flag"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/maborosii/goclient_apollo/appconfig"

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
	c := appconfig.NewAppConfig()("SampleApp", "sheng-pro", "application_properties.txt")

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
}
