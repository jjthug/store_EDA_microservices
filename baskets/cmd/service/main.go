package main

import (
	"EDA_GO/baskets"
	"EDA_GO/internal/config"
	"EDA_GO/internal/system"
	"EDA_GO/internal/web"
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("baskets exited abnormally: %s\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}
	s, err := system.NewSystem(cfg)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		if err = db.Close(); err != nil {
			return
		}
	}(s.DB())

	s.Mux().Mount("/", http.FileServer(http.FS(web.WebUI)))
	// call the model composition root
	if err = baskets.Root(s.Waiter().Context(), s); err != nil {
		return err
	}

	fmt.Println("Started baskets service")
	defer fmt.Println("Stopped baskets sevice")

	s.Waiter().Add(
		s.WaitForWeb,
		s.WaitForRPC,
		s.WaitForStream,
	)

	return s.Waiter().Wait()
}
