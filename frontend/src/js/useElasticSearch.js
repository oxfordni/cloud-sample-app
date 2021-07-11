import { useCallback, useEffect, useState } from 'react'

const useElasticSearch = (apiEndpoint, options) => {
  const [data, setData] = useState(null)
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState(null)

  const getData = useCallback(async (url, opts) => {
    try {
      setIsLoading(true)
      const res = await fetch(url, opts)
      const { ok, status, statusText } = res

      if (ok) {
        const resData = await res.json()
        setData(resData)
        setError(null)
      } else {
        setData(null)
        setError({ status, statusText })
      }
    } catch (err) {
      setData(null)
      setError(err)
    } finally {
      setIsLoading(false)
    }
  }, [setData, setError, setIsLoading])

  useEffect(() => {
    getData(apiEndpoint, options)
  }, [apiEndpoint, options])

  return [data, isLoading, error]
}

export default useElasticSearch
