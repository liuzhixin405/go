package config
import (
	"fmt"
	"net/url"
	"strings"

	"microproject/internal/pkg/config"
	"microproject/internal/pkg/config/environment"
)

type Config struct {
	AppOptions      AppOptions      `mapstructure:"app_options"       env:"APP_OPTIONS"`
	EchohttpOptions EchoHttpOptions `mapstructure:"echoHttpOptions"`
}
func NewConfig(env environment.Environment) (*Config, error) {
	cfg,err:=config.BindConfig[*Config](env)
	if err!=nil{
		return nil,err
	}
	return cfg,nil
}

type AppOptions struct {
	Name string `mapstructure:"name" env:"Name"`
}
func(c *AppOptions) GetMicroserviceNameUpper() string{
	return strings.ToUpper(c.Name)
}

func(c *AppOptions) GetMicroserviceName() string{
	return c.Name
}

type EchoHttpOptions struct {
	Port                string   `mapstructure:"port"                validate:"required" env:"Port"`
	Development         bool     `mapstructure:"development"                             env:"Development"`
	BasePath            string   `mapstructure:"basePath"            validate:"required" env:"BasePath"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"                     env:"DebugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
	Timeout             int      `mapstructure:"timeout"                                 env:"Timeout"`
	Host                string   `mapstructure:"host"                                    env:"Host"`
	Name                string   `mapstructure:"name"                                    env:"Name"`
}

func(c *EchoHttpOptions) Address() string{
	return fmt.Sprintf("%s%s",c.Host,c.Port)
}

func (c *EchoHttpOptions) BasePathAddress() string {
	path, errr:=url.JoinPath(c.Address(),c.BasePath)
	if errr!=nil{
		return ""
	}
	return path
}