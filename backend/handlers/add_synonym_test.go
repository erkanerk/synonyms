package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"synonyms-backend/data"
	"testing"

	"github.com/gorilla/mux"
)

func TestAddSynonym(t *testing.T) {
	// Reset the dictionary before each test
	data.ResetSynonymDict()

	// Add a word 'frugal' to the dictionary
	data.SynonymDict.AddWord("frugal")

	tt := []struct {
		name           string
		payload        []byte
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Synonym",
			payload:        []byte(`{"synonym": "thrifty"}`),
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"synonym":"thrifty","word":"frugal"}`,
		},
		{
			name:           "Empty Synonym",
			payload:        []byte(`{"synonym": ""}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"word cannot be empty"}`,
		},
		{
			name:           "Synonym Already Exists",
			payload:        []byte(`{"synonym": "thrifty"}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"'thrifty' is already a synonym of 'frugal'"}`,
		},
		{
			name:           "Invalid JSON Format",
			payload:        []byte(`{"synonym": "economical"`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid JSON format"}`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/synonym/frugal", bytes.NewBuffer(tc.payload))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/synonym/{word}", AddSynonym)

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
