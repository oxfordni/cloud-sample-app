import { Row, Col, Layout, Typography } from 'antd'
import BackendStatus from './backend-status'
import pkg from '../../package.json'

const { Header, Footer, Sider, Content } = Layout
const { Text, Title, Link } = Typography

const Template = ({ children }) => (
  <Layout style={{ minHeight: '100vh' }}>
    <Header>
      <Row align="middle" justify="space-between">
        <Col span={12}>
          <Title type="warning">{pkg.title}</Title>
        </Col>
        <Col span={12}>
          <BackendStatus />
        </Col>
      </Row>
    </Header>
    <Content
      style={{
        alignItems: 'center',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        overflow: 'auto',
      }}
    >
      {children}
    </Content>
    <Footer>
      <Text>
        {'Created by '}
        <Link href={pkg.author.url} target="_blank" rel="noopener noreferrer">
          {pkg.author.name}
        </Link>
      </Text>
    </Footer>
  </Layout>
)

export default Template
