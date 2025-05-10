package registry

type Config struct {
	Port          int    `yaml:"port"`
	ValidationKey string `yaml:"validation_key"`
}
