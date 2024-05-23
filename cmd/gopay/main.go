package main

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gopay/internal"
)

func initDB() {
	// temp fake db for accounts
	internal.Accounts["0001"] = &internal.Account{
		AccountId: "0001",
		Name:      "Shankar",
		LastName:  "Nakai",
	}

	internal.Accounts["0002"] = &internal.Account{
		AccountId: "0002",
		Name:      "Jessica",
		LastName:  "Lourenco",
	}

	internal.Accounts["0003"] = &internal.Account{
		AccountId: "0003",
		Name:      "Caio",
		LastName:  "Henrique",
	}

	internal.Accounts["0004"] = &internal.Account{
		AccountId: "0004",
		Name:      "Karina",
		LastName:  "Domingues",
	}
}

func main() {
	router := internal.Router(internal.Routes())
	initDB()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Server started at port :8080")

	log.
		Fatal().
		Err(http.ListenAndServe(":8080", router)).
		Msg("Server closed")
}