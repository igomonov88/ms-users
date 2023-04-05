package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Source struct {
	ds *viper.Viper
}

func NewSource() (*Source, error) {
	vp := viper.New()
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vp.SetConfigType("yml")
	vp.SetConfigName("config")
	vp.AddConfigPath("etc")
	vp.AutomaticEnv()
	vp.SetEnvPrefix("DAILYBURN")

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Source{
		ds: vp,
	}, nil
}

func (src *Source) Read(key string, dst interface{}) error {
	return src.ds.UnmarshalKey(key, dst)
}

func (src *Source) ReadConfig(dst interface{}) error {
	return src.ds.Unmarshal(dst)
}

func (src *Source) GetString(key string) string {
	return src.ds.GetString(key)
}

func (src *Source) GetInt(key string) int {
	return src.ds.GetInt(key)
}

func (src *Source) GetBool(key string) bool {
	return src.ds.GetBool(key)
}
