package flashcard

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/Tenjin0/go_network/chap10/dictionary"
)

const FLASHCARDSET_PATH = "chap10/flashcardSets"

type FlashCard struct {
	Simplified string
	English    string
	Dictionary *dictionary.Dictionary
}

type FlashCards struct {
	Name      string
	CardOrder string
	ShowHalf  string
	Cards     []*FlashCard
}

func Gets() []string {

	flashcardsDir, err := os.Open(FLASHCARDSET_PATH)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	files, err := flashcardsDir.Readdir(-1)

	fileNames := make([]string, len(files))

	for n, f := range files {
		fileNames[n] = f.Name()
	}
	sort.Strings(fileNames)
	return fileNames
}

func LoadJSON(cardname string, key interface{}) error {

	jsonFile, err := os.Open(filepath.Join(FLASHCARDSET_PATH, cardname))
	if err != nil {
		return err
	}

	if jsonFile != nil {
		defer jsonFile.Close()
	}

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(key)
	if err != nil {
		return err
	}
	return nil
}
