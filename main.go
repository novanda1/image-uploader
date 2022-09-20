/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/novanda1/image-uploader/api"
	"github.com/novanda1/image-uploader/conf"
	"github.com/sirupsen/logrus"
)

func main() {
	config, _ := conf.LoadGlobal(".env")
	api := api.NewApi(config)

	l := fmt.Sprintf("%v:%v", config.API.Host, config.API.Port)
	logrus.Infof("Image Uploader API started on: %s", l)

	http.ListenAndServe(l, api)
}
