package main

import (
	"fmt"
	"sync"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
)

func main() {
	c := &config.AppConfig{
		AppID:          "SampleApp",
		Cluster:        "default",
		IP:             "http://172.30.64.207:50080",
		NamespaceName:  "application_yml.txt",
		IsBackupConfig: false,
		Secret:         "",
	}
	c2 := &CustomChangeListener{}
	c2.wg.Add(5)

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	client.AddChangeListener(c2)

	fmt.Println("err:", err)

	c2.wg.Wait()
	writeConfig(c.NamespaceName, client)
}

func writeConfig(namespace string, client *agollo.Client) {
	cache := client.GetConfigCache(namespace)
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		return true
	})
}

type CustomChangeListener struct {
	wg sync.WaitGroup
}

func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	//write your code here
	fmt.Println(changeEvent.Changes)
	for key, value := range changeEvent.Changes {
		fmt.Println("change key : ", key, ", value :", value)
	}
	fmt.Println(changeEvent.Namespace)
	c.wg.Done()
}

func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	//write your code here
}
