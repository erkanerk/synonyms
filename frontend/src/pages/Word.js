import React from "react"
import { useParams, Link } from "react-router-dom"
import { Box, Typography, CircularProgress } from "@mui/material"
import {
  SEARCH_RESULT_TEXT,
  FAILED_TO_FETCH_SYNONYMS,
  WORD_MISSING,
  SYNONYMS_LIST_PREFIX,
  NO_SYNONYMS,
} from "../copy"
import { useGetSynonyms } from "../hooks/api"
import AddWordButton from "../components/AddWordButton"
import AddSynonymButton from "../components/AddSynonymButton"

function Word() {
  const { word } = useParams()
  const { data, isLoading, isError, error, isSuccess, refetch } =
    useGetSynonyms(word)

  const onAddWord = () => {
    refetch()
  }

  const renderContent = () => {
    if (isLoading) {
      return <CircularProgress />
    }

    if (isError) {
      if (error?.response?.status == 404) {
        return (
          <Box sx={{ display: "flex", flexDirection: "column", gap: "50px" }}>
            <Typography variant="h4">{WORD_MISSING(word)}</Typography>
            <AddWordButton word={word} onSuccess={onAddWord} />
          </Box>
        )
      } else {
        return (
          <Typography variant="h4">{FAILED_TO_FETCH_SYNONYMS(word)}</Typography>
        )
      }
    }

    if (isSuccess) {
      return (
        <>
          <Typography variant="h4">{SEARCH_RESULT_TEXT(word)}</Typography>
          <Box
            sx={{ display: "flex", flexDirection: "column", gap: "20px" }}
            mt={2}
          >
            {data?.synonyms?.length > 0 ? (
              <>
                <Typography variant="h5">{SYNONYMS_LIST_PREFIX}</Typography>
                <ul>
                  {data?.synonyms.map((synonym, index) => (
                    <li key={index}>
                      <Link to={`/word/${synonym}`}>{synonym}</Link>
                    </li>
                  ))}
                </ul>
              </>
            ) : (
              <Typography variant="h6">{NO_SYNONYMS}</Typography>
            )}
            <AddSynonymButton word={word} onSuccess={refetch} />
          </Box>
        </>
      )
    }
  }

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
      {renderContent()}
    </Box>
  )
}

export default Word
