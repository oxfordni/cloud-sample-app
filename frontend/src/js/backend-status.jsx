import { useEffect, useState } from 'react'
import { Typography, Space } from 'antd'
import useAPI from './useAPI'
import api from './api'

const { Text, Link } = Typography

const BackendStatus = () => {
  const [goIsAlive, setGoIsAlive] = useState(false)
  const [pyIsAlive, setPyIsAlive] = useState(false)
  const [goData] = useAPI(api.backendGoStatus)
  const [pyData] = useAPI(api.backendPyStatus)

  useEffect(() => {
    if (goData) {
      setGoIsAlive(goData.alive === 'ok')
    }
  }, [goData])

  useEffect(() => {
    if (pyData) {
      setPyIsAlive(pyData.alive === 'ok')
    }
  }, [pyData])

  return (
    <Space direction="horizontal" size="large" style={{ float: 'right' }}>
      <Text type={goIsAlive ? 'success' : 'danger'}>go service</Text>
      <Text type={pyIsAlive ? 'success' : 'danger'}>python service</Text>
    </Space>
    // <span style={{ color: 'white', float: 'right' }}>{String(goIsAlive)}</span>
  )
}

export default BackendStatus
