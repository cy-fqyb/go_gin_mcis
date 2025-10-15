'use client'

import { Layout, Menu, Avatar, Space } from 'antd'
import { UserOutlined, HomeOutlined, AppstoreOutlined, SettingOutlined, InfoCircleOutlined } from '@ant-design/icons'
import { useRouter } from 'next/navigation'

const { Header, Content } = Layout

export default function HomePage() {
  const router = useRouter()

  // 顶部导航菜单
  const menuItems = [
    { key: 'home', label: '首页', icon: <HomeOutlined /> },
    { key: 'projects', label: '项目', icon: <AppstoreOutlined /> },
    { key: 'settings', label: '设置', icon: <SettingOutlined /> },
    { key: 'about', label: '关于', icon: <InfoCircleOutlined /> },
  ]

  const onMenuClick = (e: any) => {
    switch (e.key) {
      case 'home':
        router.push('/home')
        break
      case 'projects':
        router.push('/projects')
        break
      case 'settings':
        router.push('/settings')
        break
      case 'about':
        router.push('/about')
        break
      default:
        break
    }
  }

  const handleLogout = () => {
    localStorage.removeItem('token')
    router.push('/login')
  }

  return (
    <Layout className="min-h-screen bg-white">
      {/* 顶部导航栏 */}
      <Header className="flex justify-between items-center px-10 !bg-white shadow-md">
        {/* 左侧 Logo + 菜单 */}
        <div className="flex items-center gap-8">
          <h1
            className="text-xl font-bold text-blue-600 cursor-pointer"
            onClick={() => router.push('/home')}
          >
            🌟 MyDashboard
          </h1>
          <Menu
            mode="horizontal"
            items={menuItems}
            onClick={onMenuClick}
            defaultSelectedKeys={['home']}
            className="border-0 bg-transparent"
          />
        </div>

        {/* 右侧用户信息 */}
        <div
          className="flex items-center gap-3 cursor-pointer hover:bg-gray-100 px-3 py-1 rounded-lg transition"
          onClick={handleLogout}
        >
          <Space>
            <Avatar icon={<UserOutlined />} />
            <span className="text-gray-700 font-medium">退出登录</span>
          </Space>
        </div>
      </Header>

      {/* 页面主体 */}
      <Content className="p-10">
        <div className="bg-white rounded-xl shadow-sm p-8">
          <h2 className="text-2xl font-semibold mb-4 text-gray-800">欢迎回来 👋</h2>
          <p className="text-gray-600 leading-relaxed">
            枫叶疏林映晚霞，
            桥横烟水接天涯。
            古渡寒鸦归渔火，
            道旁流水绕村家。
          </p>
        </div>
      </Content>
    </Layout>
  )
}
