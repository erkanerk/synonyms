import React, { useState } from "react"
import PropTypes from "prop-types"
import {
  Button,
  Typography,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
} from "@mui/material"
import { useAddSynonym } from "../hooks/api.js"
import {
  ADD_SYNONYM_BUTTON_TEXT,
  GENERIC_ERROR,
  CANCEL_BUTTON_TEXT,
  ADD_SYNONYM_DIALOG_TITLE,
  ADD_SYNONYM_DIALOG_BUTTON_TEXT,
} from "../copy"

const AddSynonymButton = ({ word, onSuccess }) => {
  const [openDialog, setOpenDialog] = useState(false)
  const [synonym, setSynonym] = useState("")
  const addSynonymMutation = useAddSynonym()

  const handleOpenDialog = () => {
    setOpenDialog(true)
  }

  const handleCloseDialog = () => {
    setOpenDialog(false)
    setSynonym("")
  }

  const handleInputChange = (event) => {
    setSynonym(event.target.value)
  }

  const handleKeyDown = (event) => {
    if (event.key === "Enter") {
      handleAddSynonym()
    }
  }

  const handleAddSynonym = () => {
    addSynonymMutation.mutate(
      { word, synonym },
      {
        onSuccess: () => {
          onSuccess()
          handleCloseDialog()
        },
      }
    )
  }

  return (
    <div>
      <Button
        disabled={addSynonymMutation.isLoading}
        variant="contained"
        color="primary"
        onClick={handleOpenDialog}
      >
        {ADD_SYNONYM_BUTTON_TEXT}
      </Button>
      <Dialog open={openDialog} onClose={handleCloseDialog}>
        <DialogTitle>{ADD_SYNONYM_DIALOG_TITLE(word)}</DialogTitle>
        <DialogContent>
          {addSynonymMutation.isError && (
            <Typography variant="body1" color="error">
              {GENERIC_ERROR}
            </Typography>
          )}
          <TextField
            autoFocus
            margin="dense"
            label="Synonym"
            fullWidth
            onKeyDown={handleKeyDown}
            value={synonym}
            onChange={handleInputChange}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseDialog} color="primary">
            {CANCEL_BUTTON_TEXT}
          </Button>
          <Button
            onClick={handleAddSynonym}
            color="primary"
            variant="contained"
          >
            {ADD_SYNONYM_DIALOG_BUTTON_TEXT}
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  )
}

AddSynonymButton.propTypes = {
  word: PropTypes.string.isRequired,
  onSuccess: PropTypes.func.isRequired,
}

export default AddSynonymButton
