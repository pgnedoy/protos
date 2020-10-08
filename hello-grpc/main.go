package main

import (
	"github.com/spf13/viper"

	"github.com/9count/go-services/boilerplates/hello-grpc/cmd"
)

func init() {
	viper.AutomaticEnv()
}

func main() {
	cmd.Execute()
}
