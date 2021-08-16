package appconfig

import (
	"github.com/zouyx/agollo/v4/env/config"
)

type IConfigfile func(appid string, cluster string, namespace string) *config.AppConfig

//设置默认参数
func NewAppConfig() IConfigfile {

	return func(appid string, cluster string, namespace string) *config.AppConfig {
		return &config.AppConfig{
			AppID:          appid,
			Cluster:        cluster,
			IP:             "http://173.16.37.170:50080",
			NamespaceName:  namespace,
			IsBackupConfig: false,
			Secret:         "",
		}
	}

}
