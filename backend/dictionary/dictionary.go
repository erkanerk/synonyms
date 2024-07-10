package dictionary

import "fmt"

type SynonymSet map[string]struct{}

// Using a map of maps to store synonyms allows for O(1) lookup, addition and deletion
type SynonymDict struct {
	words map[string]SynonymSet
}

// Constructor
func NewSynonymDict() *SynonymDict {
	return &SynonymDict{
		words: make(map[string]SynonymSet),
	}
}

func (d *SynonymDict) AddWord(word string) error {
	if err := d.wordValidation(word); err != nil {
		return err
	}

	if d.WordExists(word) {
		return fmt.Errorf("'%s' already exists", word)
	}

	d.words[word] = make(SynonymSet)

	return nil
}

// Adds a synonym to a word bidirectionally
func (d *SynonymDict) AddSynonym(word, synonym string) error {
	if err := d.wordValidation(word); err != nil {
		return err
	}

	if err := d.wordValidation(synonym); err != nil {
		return err
	}

	if synonym == word {
		return fmt.Errorf("synonym cannot equal word")
	}

	if !d.WordExists(word) {
		return fmt.Errorf("word '%s' does not exist in dictionary", word)
	}

	if _, exists := d.words[word][synonym]; exists {
		return fmt.Errorf("'%s' is already a synonym of '%s'", synonym, word)
	}

	if !d.WordExists(synonym) {
		if err := d.AddWord(synonym); err != nil {
			return err
		}
	}

	// Add synonym bidirectionally
	d.words[word][synonym] = struct{}{}
	d.words[synonym][word] = struct{}{}

	return nil
}

func (d *SynonymDict) GetSynonyms(word string) ([]string, error) {
	if err := d.wordValidation(word); err != nil {
		return nil, err
	}

	if !d.WordExists(word) {
		return nil, fmt.Errorf("word '%s' does not exist in dictionary", word)
	}

	synonyms := []string{}
	if synSet, exists := d.words[word]; exists {
		for syn := range synSet {
			synonyms = append(synonyms, syn)
		}
	}
	return synonyms, nil
}

func (d *SynonymDict) WordExists(word string) bool {
	_, exists := d.words[word]
	return exists
}

func (d *SynonymDict) wordValidation(word string) error {
	if word == "" {
		return fmt.Errorf("word cannot be empty")
	}
	return nil
}
