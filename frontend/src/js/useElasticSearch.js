import { useCallback, useEffect, useState } from 'react'

const initialStatus = { data: null, error: null, loading: false }

const useElasticSearch = (apiEndpoint, options) => {
  const [status, setStatus] = useState(initialStatus)
  const { data, error, isLoading } = status

  const getData = useCallback(
    async (url, opts) => {
      try {
        setStatus({ ...initialStatus, isLoading: true })

        const res = await fetch(url, opts)
        const { ok, status, statusText } = res

        if (ok) {
          const resData = await res.json()
          setStatus({ ...initialStatus, data: resData })
        } else {
          setStatus({ ...initialStatus, error: { status, statusText } })
        }
      } catch (err) {
        setStatus({ ...initialStatus, error: err })
      }
    },
    [setStatus],
  )

  useEffect(() => {
    getData(apiEndpoint, options)
  }, [apiEndpoint, options])

  return [data, isLoading, error]
}

export default useElasticSearch
