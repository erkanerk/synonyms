import React from "react"
import PropTypes from "prop-types"
import { Button, Typography } from "@mui/material"
import { useAddWord } from "../hooks/api.js"
import { ADD_WORD_BUTTON_TEXT, GENERIC_ERROR } from "../copy"

const AddWordButton = ({ word, onSuccess }) => {
  const addWordMutation = useAddWord()

  const handleAddWord = () => {
    addWordMutation.mutate(word, {
      onSuccess: () => {
        onSuccess()
      },
    })
  }

  return (
    <div>
      {addWordMutation.isError && (
        <Typography variant="body1" color="error">
          {GENERIC_ERROR}
        </Typography>
      )}
      {!addWordMutation.isError && (
        <Button
          disabled={addWordMutation.isLoading}
          variant="contained"
          color="primary"
          onClick={handleAddWord}
        >
          {ADD_WORD_BUTTON_TEXT(word)}
        </Button>
      )}
    </div>
  )
}

AddWordButton.propTypes = {
  word: PropTypes.string.isRequired,
  onSuccess: PropTypes.func.isRequired,
}

export default AddWordButton
