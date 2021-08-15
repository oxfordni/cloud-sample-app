import { Space } from 'antd'
import Template from './template'
import RandomQuote from './random-quote'
import 'antd/dist/antd.css'

const App = () => (
  <Template>
    <Space direction="vertical">
      <RandomQuote />
    </Space>
  </Template>
)

export default App
