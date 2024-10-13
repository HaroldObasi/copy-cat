package main

import (
	"github.com/HaroldObasi/copy-cat/bootstrap"
	"github.com/HaroldObasi/copy-cat/template"
)

func main() {
	const APP_NAME = "app"
	bootstrap.BootStrapApp(APP_NAME)
	template.FormatTemplate(APP_NAME)
}