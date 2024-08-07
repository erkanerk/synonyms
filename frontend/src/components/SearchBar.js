import React, { useState } from "react"
import { useNavigate } from "react-router-dom"
import { TextField, Button, Box } from "@mui/material"
import { SEARCH_BAR_PLACEHOLDER, SEARCH_BUTTON_TEXT } from "../copy"

function SearchBar() {
  const [searchValue, setSearchValue] = useState("")
  const navigate = useNavigate()

  const handleInputChange = (event) => {
    setSearchValue(event.target.value)
  }

  const handleSearch = () => {
    const sanitizedWord = searchValue.trim().toLowerCase()
    if (!sanitizedWord) {
      return
    }
    navigate(`/word/${sanitizedWord}`)
  }

  const handleKeyDown = (event) => {
    if (event.key === "Enter") {
      handleSearch()
    }
  }

  return (
    <Box
      sx={{
        position: "fixed",
        width: "100%",
        padding: "20px",
        display: "flex",
        gap: "10px",
        alignItems: "center",
        justifyContent: "center",
        boxShadow: "0px 2px 4px rgba(0, 0, 0, 0.1)",
      }}
    >
      <TextField
        label={SEARCH_BAR_PLACEHOLDER}
        variant="outlined"
        value={searchValue}
        onChange={handleInputChange}
        onKeyDown={handleKeyDown}
        size="medium"
        sx={{ width: "700px", maxWidth: "70%", backgroundColor: "white" }}
      />
      <Button
        variant="contained"
        color="primary"
        onClick={handleSearch}
        sx={{ height: "56px" }}
      >
        {SEARCH_BUTTON_TEXT}
      </Button>
    </Box>
  )
}

export default SearchBar
