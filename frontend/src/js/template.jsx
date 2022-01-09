import { Layout, Typography } from 'antd'

const { Header, Footer, Sider, Content } = Layout
const { Text, Title, Link } = Typography

const Template = ({ children }) => (
  <Layout style={{ minHeight: '100vh' }}>
    <Header>
      <Title type="warning">go+es</Title>
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
        <Link
          href="https://joaocarmo.com"
          target="_blank"
          rel="noopener noreferrer"
        >
          João Carmo
        </Link>
      </Text>
    </Footer>
  </Layout>
)

export default Template
