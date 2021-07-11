import React from 'react'
import { Alert, Skeleton, Tag } from 'antd'
import useElasticSearch from './useElasticSearch'
import api from './api'

const RandomQuote = () => {
  const [data, isLoading, error] = useElasticSearch(api.getRandomQuote)

  return (
      <Alert
        message="Random Quote"
        description={
          isLoading
            ? <Skeleton active />
            : (
              <>
                <p>
                  <q>{data?.quote}</q>
                </p>
                <p>
                  <cite>— {data?.role}</cite>
                </p>
                <div>
                  <Tag color="volcano">{data?.show}</Tag>
                  {data?.contain_adult_lang && (
                    <Tag color="magenta">{data?.contain_adult_lang}</Tag>
                  )}
                </div>
              </>
            )
          }
        type="warning"
      />
  )
}

export default RandomQuote
