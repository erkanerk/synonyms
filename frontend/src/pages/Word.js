import React from "react"
import { useParams } from "react-router-dom"
import { Box, Typography } from "@mui/material"

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
      <Typography variant="h4">Word Page</Typography>
      <Typography variant="body1">You searched for: {word}</Typography>
    </Box>
  )
}

export default Word
