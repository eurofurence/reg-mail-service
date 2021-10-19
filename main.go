package main

import (
	"fmt"
	"log"

	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/repository/logging/consolelogging/logformat"
	"github.com/eurofurence/reg-mail-service/web"
)

func main() {
	err := config.LoadConfiguration("config.yaml")
	if err != nil {
		log.Fatal(logformat.Logformat("ERROR", "00000000", fmt.Sprintf("Error while loading configuration: %v", err)))
	}
	server := web.Create()
	web.Serve(server)
}
