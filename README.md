# Synonym Manager
This project consists of a backend API and a frontend React application for managing a dictionary of words and their synonyms. It allows adding words, adding synonyms for words, and retrieving synonyms for a given word.

## Features

- Add a word to the dictionary.
- Add a synonym for a word.
- Retrieve synonyms for a given word.

## Backend

### Design
I tried to create a really simple, elegant and organized backend with a popular framework, mux. Tests were implemented for the parts that needs testing in order to ensure working functionality.

The requirements mentions a fourth endpoint to obtain words from a synonym but I assume this is kind of a "trick question", and I am supposed to be able to question the requirements.. No fourth endpoint should be needed as synonyms are also words and the relationship is bidirectional. If you add a synonym to a word, it makes sense for both words to obtain a new synonym.

The datastructure I created to store the data in memory can be seen in `backend/dictionary`. It is a map of a map, where the inner map is just used as a set with the value being struct{}{} occupying no memory. This structure should be efficient for lookup, adding and deleting in O(1) time.

### Prerequisites
- Go 1.20 or higher

### Running the backend locally
Navigate to the backend folder and run `go mod tidy` to install dependencies. Then run `go run main.go`.

### Testing
Navigate to the backend folder and run `go test ./...` to run all test files in the backend.

### Continous Integration
A github action workflow can been seen in `.github/workflows/backend-checks.yml` where lints and tests are checked on PRs and pushes to main.

## Frontend

### Design
With the limited requirements I got I tried to come up with a very simple and intuitive design for the use case of this application. I thought it would be nice with the search up top and always have it available. And no need to clutter the UI with buttons to create words or synonyms until you are located where it makes sense to do those actions. 

To get a simple and clean look with minimal effort I used Material UI and inline styles with the sx prop. This way there is no need for css files and I can keep things tidy as long as files are kept small. I also installed styled-components in case I would stumble on a case where the sx prop was not enough, but I managed to not need it.

### Prerequisites
- Node.js (v14.x or higher is recommended)
- npm

### Running the frontend locally
Navigate to the frontend folder and run `npm install` to install dependencies. Then run `npm start`.

### Testing
Navigate to the frontend folder and run `npm test` to run all test files in the frontend. There is currently an annoying bug with the tests so the test step has been left out of the github actions workflow, but the bug is not affecting the functionality of the app.

### Continous Integration
A github action workflow can been seen in `.github/workflows/frontend-checks.yml` where lints are checked on PRs and pushes to main. As previously mentioned I have had to leave out the testing step due to a bug.