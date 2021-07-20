// package main

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/zouyx/agollo/v4"
// 	"github.com/zouyx/agollo/v4/env/config"
// )

// // convert apollo config  namespacename to lcoal config filename
// func ConfiFileName(namespace string) (filename string) {
// 	filename = strings.TrimSuffix(namespace, ".txt")
// 	p := strings.LastIndex(filename, "_")
// 	filename = filename[:p] + "." + filename[p+1:]
// 	return

// }

// func main() {
// 	c := &config.AppConfig{
// 		AppID:          "SampleApp",
// 		Cluster:        "default",
// 		IP:             "http://172.30.64.207:50080",
// 		NamespaceName:  "application_yml.txt",
// 		IsBackupConfig: false,
// 		Secret:         "",
// 	}

// 	agollo.SetLogger(&DefaultLogger{})

// 	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
// 		return c, nil
// 	})

// 	if err != nil {
// 		fmt.Println("err:", err)
// 		panic(err)
// 	}

// 	// checkKey(c.NamespaceName, client)
// 	conf := client.GetConfig(c.NamespaceName)
// 	fmt.Println(strings.TrimPrefix(conf.GetContent(), "content="))
// 	fmt.Println(ConfiFileName(c.NamespaceName))

// 	// c = &config.AppConfig{
// 	// 	AppID:          "hk109",
// 	// 	Cluster:        "dev",
// 	// 	IP:             "http://106.54.227.205:8080",
// 	// 	NamespaceName:  "dubbo",
// 	// 	IsBackupConfig: false,
// 	// 	Secret:         "6ce3ff7e96a24335a9634fe9abca6d51",
// 	// }

// 	// client, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
// 	// 	return c, nil
// 	// })

// 	// if err != nil {
// 	// 	fmt.Println("err:", err)
// 	// 	panic(err)
// 	// }

// 	// checkKey(c.NamespaceName, client)

// 	// time.Sleep(5 * time.Second)
// }

// // func checkKey(namespace string, client *agollo.Client) {
// // 	cache := client.GetConfigCache(namespace)
// // 	count := 0
// // 	cache.Range(func(key, value interface{}) bool {
// // 		fmt.Printf("%s=%s\n", key, value)
// // 		count++
// // 		return true
// // 	})
// // 	if count < 1 {
// // 		panic("config key can not be null")
// // 	}
// // 	// fmt.Println(count)
// // }

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
