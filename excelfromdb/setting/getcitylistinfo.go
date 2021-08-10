package setting

import (
	"embed"
	. "excelfromdb/locallog"
	"reflect"

	"github.com/spf13/viper"
)

//go:embed citylist.yaml
var cityfile embed.FS

func GetCity() ([]string, []string) {

	citieslist := []string{}
	codelist := []string{}

	config := viper.New()
	// config.AddConfigPath(".")
	// config.SetConfigName("citylist")
	config.SetConfigType("yaml")

	iofile, err := cityfile.ReadFile("citylist.yaml")

	if err != nil {
		Log.Fatal(err)
	}
	if err = config.ReadIOInConfig(iofile); err != nil {
		// 这里修改了viper的源码，增加了读取fs.file的选项
		Log.Fatal(err)
	}

	c := reflect.ValueOf(config.Get("cities.name"))
	for i := 0; i < c.Len(); i++ {
		elec := c.Index(i)
		citieslist = append(citieslist, elec.Interface().(string))
	}

	d := reflect.ValueOf(config.Get("cities.code"))
	for i := 0; i < d.Len(); i++ {
		eled := d.Index(i)
		codelist = append(codelist, eled.Interface().(string))
	}
	return citieslist, codelist
}
