package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"synonyms-backend/data"
	"synonyms-backend/utils"

	"github.com/gorilla/mux"
)

func AddSynonym(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := strings.TrimSpace(strings.ToLower(vars["word"]))

	var input struct {
		Synonym string `json:"synonym"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid JSON format")
		return
	}

	synonym := strings.TrimSpace(strings.ToLower(input.Synonym))

	if err := data.SynonymDict.AddSynonym(word, synonym); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"synonym": synonym, "word": word})
}
