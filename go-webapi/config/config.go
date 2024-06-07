package config

type Config struct {
	AppOptions      AppOptions      `mapstructure:"app_options"       env:"APP_OPTIONS"`
	EchohttpOptions EchoHttpOptions `mapstructure:"echoHttpOptions"`
}

type AppOptions struct {
	Name string `mapstructure:"name" env:"Name"`
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
