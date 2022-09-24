package helpers

import (
	"CYAMod/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func MapJSONToStories() models.Story {
	var(
		 story models.Story
		 err error
		 jsonFile *os.File
		 byteValue []byte
	)

	jsonFile ,err = os.Open("stories.json")
	if err != nil {
		log.Fatal("Open: ",err)
	}
	defer jsonFile.Close()
	byteValue ,err = ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("ReadAll: ",err)
	}
	json.Unmarshal(byteValue,&story)
	return story
}