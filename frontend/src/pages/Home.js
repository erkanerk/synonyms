import React from "react"
import { HOME_PAGE_TITLE } from "../copy"
import { Typography, Box } from "@mui/material"

function Home() {
  return (
    <Box
      sx={{
        width: "100vw",
        height: "100vh",
        padding: "10%",
        paddingBottom: "30%",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        textAlign: "center",
      }}
    >
      <Typography variant="h3">{HOME_PAGE_TITLE}</Typography>
    </Box>
  )
}

export default Home
