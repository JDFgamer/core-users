package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DataBaseMongo DataBase `yaml:"database_mongo"`
}

type DataBase struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// Cargar configuraciones desde el archivo YAML correspondiente
func LoadConfig() (*Config, error) {
	// Determina el entorno (puedes usar una variable de entorno o un valor hardcodeado)
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local" // Por defecto, se carga el entorno local
	}

	// Configuración de Viper
	viper.SetConfigName(fmt.Sprintf("%s", env)) // El nombre del archivo de configuración cambia según el entorno
	viper.AddConfigPath("./config")             // Ruta donde buscar el archivo
	viper.SetConfigType("yaml")                 // Tipo de archivo

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error loading config file for environment '%s': %v", env, err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Error unmarshalling config: %v", err)
	}

	return &cfg, nil
}
