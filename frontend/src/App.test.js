import { render, screen, fireEvent } from "@testing-library/react"
import App from "./App"
import {
  SEARCH_BAR_PLACEHOLDER,
  SEARCH_BUTTON_TEXT,
  SEARCH_RESULT_TEXT,
  HOME_PAGE_TITLE,
} from "./copy"

test("renders search bar and home page content", () => {
  render(<App />)

  const searchBar = screen.getByLabelText(
    new RegExp(SEARCH_BAR_PLACEHOLDER, "i")
  )
  expect(searchBar).toBeInTheDocument()

  const homeContent = screen.getByText(new RegExp(HOME_PAGE_TITLE, "i"))
  expect(homeContent).toBeInTheDocument()
})

test("navigates to word page on search", () => {
  render(<App />)

  const searchValue = "test"

  const searchBar = screen.getByLabelText(
    new RegExp(SEARCH_BAR_PLACEHOLDER, "i")
  )

  expect(searchBar).toBeInTheDocument()

  fireEvent.change(searchBar, { target: { value: searchValue } })

  const searchButton = screen.getByRole("button", {
    name: new RegExp(SEARCH_BUTTON_TEXT, "i"),
  })

  fireEvent.click(searchButton)

  const wordContent = screen.getByText(new RegExp(SEARCH_RESULT_TEXT(searchValue), "i"))
  expect(wordContent).toBeInTheDocument()
})
