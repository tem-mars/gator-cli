package main

import (
	"fmt"
	"log"

	"github.com/tem-mars/gator-cli/internal/config"
)

func main() {
	// read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("cannot read config file: %v", err)
	}
	fmt.Println("default config:", cfg)

	// set user name and save
	err = cfg.SetUser("your name") // change to your name
	if err != nil {
		log.Fatalf("cannot save config file: %v", err)
	}
	fmt.Println("saved user name")

	// read config again
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("cannot read updated config file: %v", err)
	}
	fmt.Println("updated config:", updatedCfg)
}