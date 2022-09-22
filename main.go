package main

import (
	"fmt"

	"github.com/novanda1/image-uploader/api"
	"github.com/novanda1/image-uploader/conf"
	"github.com/sirupsen/logrus"
)

func main() {
	config, err := conf.LoadGlobal(".env")

	if err != nil {
		logrus.Fatalf("failed load .env: %v", err)
	}

	api := api.NewApi(config)

	l := fmt.Sprintf("%v:%v", config.API.Host, config.API.Port)
	logrus.Infof("Image Uploader API started on: %s", l)

	api.ListenAndServe(l)
}
