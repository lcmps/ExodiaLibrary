package app

import (
	"fmt"
	"sync"

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

func GetAllCardsLanguages() (english model.CardList, french model.CardList, portuguese model.CardList) {
	var wg sync.WaitGroup
	var languages = []string{
		"en",
		"fr",
		"pt",
	}
	wg.Add(len(languages))

	for _, lang := range languages {
		go func(lang string) {
			defer wg.Done()

			if lang == "en" {
				en, err := GetAllCards(nil)
				if err != nil {
					fmt.Printf("Error executing request on %s: %s\n", lang, err.Error())
				}
				fmt.Printf("Language: %s done. Results: %d\n", lang, len(en.Data))
				english = en
				return
			}
			if lang == "fr" {
				fr, err := GetAllCards(fmt.Sprintf("language=%s", lang))
				if err != nil {
					fmt.Printf("Error executing request on %s: %s\n", lang, err.Error())
				}
				fmt.Printf("Language: %s done. Results: %d\n", lang, len(fr.Data))
				french = fr
				return
			}
			if lang == "pt" {
				pt, err := GetAllCards(fmt.Sprintf("language=%s", lang))
				if err != nil {
					fmt.Printf("Error executing request on %s: %s\n", lang, err.Error())
				}
				fmt.Printf("Language: %s done. Results: %d\n", lang, len(pt.Data))
				portuguese = pt
				return
			}
		}(lang)
	}
	wg.Wait()

	return english, french, portuguese
}
