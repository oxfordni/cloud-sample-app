import { Alert, Skeleton, Tag, Typography } from 'antd'
import useElasticSearch from './useElasticSearch'
import api from './api'

const { Text } = Typography

const minWidth = 512

const RandomQuote = () => {
  const [data, isLoading, error] = useElasticSearch(api.getRandomQuote)

  if (error) {
    return (
      <Alert
        message="ERROR"
        description={<Text code>{JSON.stringify(error, null, 2)}</Text>}
        type="error"
      />
    )
  }

  return (
    <Alert
      message="Random Quote"
      description={
        isLoading ? (
          <Skeleton active paragraph={{ minWidth }} />
        ) : (
          <>
            <p>
              <q>{data?.quote}</q>
            </p>
            <p>
              <cite>— {data?.role}</cite>
            </p>
            <div>
              <Tag color="volcano">{data?.show}</Tag>
              {data?.contain_adult_lang && <Tag color="magenta">Adult</Tag>}
            </div>
          </>
        )
      }
      type="warning"
      style={{ minWidth }}
    />
  )
}

RandomQuote.whyDidYouRender = true

export default RandomQuote
