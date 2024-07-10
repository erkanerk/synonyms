package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"synonyms-backend/data"
	"synonyms-backend/utils"
)

func AddWord(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Word string `json:"word"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid JSON format")
		return
	}

	word := strings.TrimSpace(strings.ToLower(input.Word))

	if err := data.SynonymDict.AddWord(word); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"word": word})
}
