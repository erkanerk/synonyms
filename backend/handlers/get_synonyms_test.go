package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"synonyms-backend/data"

	"github.com/gorilla/mux"
)

func TestGetSynonyms(t *testing.T) {

	data.ResetSynonymDict()

	data.SynonymDict.AddWord("careful")
	data.SynonymDict.AddSynonym("careful", "frugal")

	tt := []struct {
		name           string
		word           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Case",
			word:           "careful",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"synonyms":["frugal"]}`,
		},
		{
			name:           "Word Not Found",
			word:           "unknown",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"word 'unknown' does not exist in dictionary"}`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/synonyms/"+tc.word, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/synonyms/{word}", GetSynonyms)

			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}

			if rr.Body.String() != tc.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expectedBody)
			}
		})
	}
}
