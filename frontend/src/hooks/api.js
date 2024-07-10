import { useQuery, useMutation } from "react-query"
import axios from "axios"

const instance = axios.create({
  baseURL: "http://localhost:8080",
})

const useAddWord = () => {
  return useMutation((word) =>
    instance.post("/word", { word }).then((res) => res.data)
  )
}

const useAddSynonym = () => {
  return useMutation(({ word, synonym }) =>
    instance.post(`/synonym/${word}`, { synonym }).then((res) => res.data)
  )
}

const useGetSynonyms = (word) => {
  return useQuery(
    ["synonyms", word],
    () => instance.get(`/synonyms/${word}`).then((res) => res.data),
    {
      retry: false,
    }
  )
}

export { useAddWord, useAddSynonym, useGetSynonyms }
