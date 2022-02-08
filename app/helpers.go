package app

import (
	"fmt"

	"github.com/lcmps/ExodiaLibrary/model"
)

func GetAllCards(opt interface{}) (model.CardList, error) {
	var url string
	var jsonResp model.CardList

	if opt != nil {
		url = fmt.Sprintf("https://db.ygoprodeck.com/api/v7/cardinfo.php?%v", opt)
	} else {
		url = "https://db.ygoprodeck.com/api/v7/cardinfo.php"
	}

	_, err := MakeRequestFastHTTP(url, nil, &jsonResp)
	if err != nil {
		fmt.Printf("Error executing request: %s\n", err.Error())
		return jsonResp, err
	}
	return jsonResp, nil
}
