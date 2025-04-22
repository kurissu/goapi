package tools
import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"kurissu": {
		AuthToken: "123ABC",
		Username: "kurissu",
	},
	"poppy":{
		AuthToken: "456DEF",
		Username: "poppy",
	},
	"tobi":{
		AuthToken: "789GHI",
		Username: "tobi",
	}	,
}
	
var mockCoinDetails = map[string]CoinDetails{
	"kurissu": {
		Coins: 100,
		UserName: "kurissu",
	},
	"poppy":{
		Coins: 200,
		UserName: "poppy",
	},
	"tobi":{
		Coins: 300,
		UserName: "tobi",
	}	,
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}
