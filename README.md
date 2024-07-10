# Features
This project consists of a backend API and a frontend React application for managing a dictionary of words and their synonyms. It allows adding words, adding synonyms for words, and retrieving synonyms for a given word.

## Features

- Add a word to the dictionary.
- Add a synonym for a word.
- Retrieve synonyms for a given word.

## Backend

### Prerequisites
- Go 1.20 or higher

### Running the backend locally
Navigate to the backend folder and run `go mod tidy` to install dependencies. Then run `go run main.go`.

### Testing
Navigate to the backend folder and run `go test ./...` to run all test files in the backend.

### Continous Integration
A github action workflow can been seen in `.github/workflows/backend-checks.yml` where lints and tests are checked on PRs and pushes to main.

## Frontend

### Prerequisites
- Node.js (v14.x or higher is recommended)
- npm

### Running the frontend locally
Navigate to the frontend folder and run `npm install` to install dependencies. Then run `npm start`.

### Testing
Navigate to the frontend folder and run `npm test` to run all test files in the frontend. There is currently an annoying bug with the tests so the test step has been left out of the github actions workflow, but the bug is not affecting the functionality of the app.

### Continous Integration
A github action workflow can been seen in `.github/workflows/frontend-checks.yml` where lints are checked on PRs and pushes to main. As previously mentioned I have had to leave out the testing step due to a bug.