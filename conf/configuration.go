package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type APIConfiguration struct {
	Host        string
	Port        int    `envconfig:"PORT" default:"8081"`
	ExternalURL string `json:"external_url" envconfig:"API_EXTERNAL_URL"`
}

type ImageKitConfiguration struct {
	PubKey  string `json:"imagekit_pubkey" envconfig:"IMAGEKIT_PUBKEY"`
	PrivKey string `json:"imagekit_privkey" envconfig:"IMAGEKIT_PRIVKEY"`
	Id      string `json:"imagekit_id" envconfig:"IMAGEKIT_ID"`
}

type GlobalConfiguration struct {
	API APIConfiguration
	IK  ImageKitConfiguration
}

func LoadGlobal(filename string) (*GlobalConfiguration, error) {
	if err := loadEnvironment(filename); err != nil {
		return nil, err
	}

	config := new(GlobalConfiguration)
	if err := envconfig.Process("app", config); err != nil {
		return nil, err
	}

	return config, nil
}

func loadEnvironment(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Overload(filename)
	} else {
		err = godotenv.Load()
		// handle if .env file does not exist, this is OK
		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}
