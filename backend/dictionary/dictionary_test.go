package dictionary

import (
	"testing"
)

func TestAddWord(t *testing.T) {
	dict := NewSynonymDict()

	// Test adding a new word
	if err := dict.AddWord("frugal"); err != nil {
		t.Errorf("unexpected error adding word: %v", err)
	}

	// Test adding an existing word
	err := dict.AddWord("frugal")
	if err == nil {
		t.Error("expected error adding existing word")
	} else if err.Error() != "'frugal' already exists" {
		t.Errorf("unexpected error message: %v", err)
	}

	// Check if 'frugal' is in the dictionary
	if _, exists := dict.words["frugal"]; !exists {
		t.Errorf("expected 'frugal' to be in the dictionary")
	}
}

func TestAddSynonym(t *testing.T) {
	dict := NewSynonymDict()

	// Add 'frugal' and 'economical'
	if err := dict.AddWord("frugal"); err != nil {
		t.Errorf("unexpected error adding word: %v", err)
	}
	if err := dict.AddWord("economical"); err != nil {
		t.Errorf("unexpected error adding word: %v", err)
	}

	// Test adding valid synonym
	if err := dict.AddSynonym("frugal", "economical"); err != nil {
		t.Errorf("unexpected error adding synonym: %v", err)
	}

	// Test adding synonym that is the same as the word
	err := dict.AddSynonym("frugal", "frugal")
	if err == nil {
		t.Error("expected error adding synonym equal to word")
	} else if err.Error() != "synonym cannot equal word" {
		t.Errorf("unexpected error message: %v", err)
	}

	// Test adding synonym to non-existing word
	err = dict.AddSynonym("nonexistent", "test")
	if err == nil {
		t.Error("expected error adding synonym to non-existing word")
	} else if err.Error() != "word 'nonexistent' does not exist in dictionary" {
		t.Errorf("unexpected error message: %v", err)
	}

	// Check synonyms in the dictionary
	synonyms := dict.words["frugal"]
	if _, exists := synonyms["economical"]; !exists {
		t.Errorf("expected 'economical' to be a synonym of 'frugal'")
	}
}

func TestGetSynonyms(t *testing.T) {
	dict := NewSynonymDict()

	// Add 'frugal', 'economical', 'thrifty'
	dict.AddWord("frugal")
	dict.AddSynonym("frugal", "economical")
	dict.AddSynonym("frugal", "thrifty")

	// Test getting synonyms for an existing word
	synonyms, err := dict.GetSynonyms("frugal")
	if err != nil {
		t.Errorf("unexpected error getting synonyms: %v", err)
	}

	expected := []string{"economical", "thrifty"}
	foundSynonyms := make(map[string]bool)

	// Mark all synonyms found in the result
	for _, syn := range synonyms {
		foundSynonyms[syn] = true
	}

	// Check if all expected synonyms are found
	for _, syn := range expected {
		if !foundSynonyms[syn] {
			t.Errorf("expected synonym %s not found in result", syn)
		}
	}

	// Test getting synonyms for a non-existing word
	_, err = dict.GetSynonyms("nonexistent")
	if err == nil {
		t.Error("expected error getting synonyms for non-existing word")
	} else if err.Error() != "word 'nonexistent' does not exist in dictionary" {
		t.Errorf("unexpected error message: %v", err)
	}
}
