package main

import (
	"log"

	"github.com/5Asuka/order/common/config"
	"github.com/spf13/viper"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	log.Printf("%v", viper.Get("order"))

}
