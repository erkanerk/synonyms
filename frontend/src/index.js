import React from "react"
import { createRoot } from "react-dom/client"
import { ThemeProvider } from "@emotion/react"
import CssBaseline from "@mui/material/CssBaseline"
import App from "./App"
import theme from "./theme"

const container = document.getElementById("root")
const root = createRoot(container)

root.render(
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <App />
  </ThemeProvider>
)
