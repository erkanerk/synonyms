package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"synonyms-backend/data"

	"github.com/gorilla/mux"
)

func TestAddWord(t *testing.T) {
	data.ResetSynonymDict()
	tt := []struct {
		name            string
		payload         []byte
		expectedStatus  int
		expectedBody    string
		expectedMessage string
	}{
		{
			name:            "Add New Word",
			payload:         []byte(`{"word": "frugal"}`),
			expectedStatus:  http.StatusCreated,
			expectedBody:    `{"word":"frugal"}`,
			expectedMessage: "",
		},
		{
			name:            "Add Empty Word",
			payload:         []byte(`{"word": ""}`),
			expectedStatus:  http.StatusBadRequest,
			expectedBody:    "",
			expectedMessage: `{"error":"word cannot be empty"}`,
		},
		{
			name:            "Add Existing Word",
			payload:         []byte(`{"word": "frugal"}`),
			expectedStatus:  http.StatusBadRequest,
			expectedBody:    "",
			expectedMessage: `{"error":"'frugal' already exists"}`,
		},
		{
			name:           "Invalid JSON Format",
			payload:        []byte(`{"word": frugal"}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid JSON format"}`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("POST", "/word", bytes.NewBuffer(tc.payload))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/word", AddWord)

			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code for test '%s': got %v want %v", tc.name, status, tc.expectedStatus)
			}

			if tc.expectedBody != "" && rr.Body.String() != tc.expectedBody {
				t.Errorf("handler returned unexpected body for test '%s': got %v want %v", tc.name, rr.Body.String(), tc.expectedBody)
			}

			if tc.expectedMessage != "" && rr.Body.String() != tc.expectedMessage {
				t.Errorf("handler returned unexpected message for test '%s': got %v want %v", tc.name, rr.Body.String(), tc.expectedMessage)
			}
		})
	}
}
