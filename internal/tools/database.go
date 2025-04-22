package tools

import (

	log "github.com/sirupsen/logrus"
)

//Database collections
type LoginDetails struct{
	AuthToken string
	Username string
}

type CoinDetails struct {
	Coins int64
	UserName string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoinDetails(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}