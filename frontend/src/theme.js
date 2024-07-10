import { createTheme } from "@mui/material/styles"

const theme = createTheme({
  palette: {
    mode: "light",
    primary: {
      // A fitting primary color
      main: "#53c654",
    },
    background: {
      paper: "#f1f1f1",
      default: "#ffffff",
    },
  },
})

export default theme
