package handlers

import (
	"net/http"
	"strings"
	"synonyms-backend/data"
	"synonyms-backend/utils"

	"github.com/gorilla/mux"
)

func GetSynonyms(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := strings.TrimSpace(strings.ToLower(vars["word"]))

	// Return 404 if word does not exist
	if !data.SynonymDict.WordExists(word) {
		utils.RespondWithError(w, http.StatusNotFound, "word does not exist in dictionary")
		return
	}

	synonyms, err := data.SynonymDict.GetSynonyms(word)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string][]string{"synonyms": synonyms})
}
