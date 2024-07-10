import React from "react"
import { useParams } from "react-router-dom"
import { Box, Typography } from "@mui/material"
import { SEARCH_RESULT_TEXT } from "../copy"

function Word() {
  const { word } = useParams()

  return (
    <Box
      sx={{
        width: "100vw",
        height: "100vh",
        padding: "10%",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        textAlign: "center",
      }}
    >
      <Typography variant="h4">{SEARCH_RESULT_TEXT(word)}</Typography>
    </Box>
  )
}

export default Word
