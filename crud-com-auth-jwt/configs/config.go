// NOTE: Pasta configs é muito utilizada como um ponto de boot do sitema, é onde definimos as configurações necessárias para o sistema iniciar
package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenJWTAuth  *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf

	viper.SetConfigName("app_config") // Nome da configuração
	viper.SetConfigType("env")        // Poderia ser yaml, toml, etc
	viper.AddConfigPath(path)         // caminho para o arquivo de .env, yaml, etc
	viper.SetConfigFile(".env")       // Nome do arquivo .env, .env.local, etc
	viper.AutomaticEnv()              // Carrega todas ao configurações automaticamente
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenJWTAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
