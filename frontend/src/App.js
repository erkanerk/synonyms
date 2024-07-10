import React from "react"
import { BrowserRouter as Router, Route, Routes } from "react-router-dom"
import Home from "./pages/Home"
import Word from "./pages/Word"
import SearchBar from "./components/SearchBar"

function App() {
  return (
    <Router>
      <SearchBar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/word/:word" element={<Word />} />
      </Routes>
    </Router>
  )
}

export default App
