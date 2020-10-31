package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus")

var Properties Config

func Read() {
	err := cleanenv.ReadConfig("configs/dev.yaml", &Properties)
	if err != nil {
		logrus.Panic(err)
	}
}
func ReadByPath(path string) {
	err := cleanenv.ReadConfig(path+"/acom-mp-pok.yaml", &Properties)
	if err != nil {
		logrus.Panic(err)
	}
}

type Config struct {
	Brand  string `yaml:"brand"`
	Server server `yaml:"server"`
	Db     db     `yaml:"mongo"`
}

type server struct {
	Port string `yaml:"port" env:"PORT" env-default:"9000"`
}
type db struct {
	Uri       string `env:"DB_URI" env-default:"mongodb://localhost:27017"`
	DockerUri string `yaml:"docker.uri" env:"DOCKER_URI"`
}

