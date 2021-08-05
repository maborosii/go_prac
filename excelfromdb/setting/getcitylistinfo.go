package setting

import (
	"reflect"

	"github.com/spf13/viper"
)

func GetCity(path string) ([]string, []string) {

	citieslist := []string{}
	codelist := []string{}

	config := viper.New()
	config.AddConfigPath(path)
	config.SetConfigName("citylist")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
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
