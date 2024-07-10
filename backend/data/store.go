package data

import "synonyms-backend/dictionary"

// SynonymDict is a global variable that holds the dictionary of synonyms acting as an in memory database
var SynonymDict = dictionary.NewSynonymDict()

func ResetSynonymDict() {
	SynonymDict = dictionary.NewSynonymDict()
}