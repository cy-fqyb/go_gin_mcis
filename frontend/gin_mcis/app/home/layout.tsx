'use client'

import { Layout, Menu, Avatar, Space } from 'antd'
import {
  UserOutlined,
  HomeOutlined,
  AppstoreOutlined,
  SettingOutlined,
  InfoCircleOutlined
} from '@ant-design/icons'
import { useRouter, usePathname } from 'next/navigation'

const { Header, Content } = Layout

export default function HomeLayout({ children }: { children: React.ReactNode }) {
  const router = useRouter()
  const pathname = usePathname()

  // 顶部导航菜单
  const menuItems = [
    { key: '/home', label: '首页', icon: <HomeOutlined /> },
    { key: '/home/projects', label: '项目', icon: <AppstoreOutlined /> },
    { key: '/home/settings', label: '设置', icon: <SettingOutlined /> },
    { key: '/home/about', label: '关于', icon: <InfoCircleOutlined /> },
  ]

  const handleMenuClick = (e: any) => {
    router.push(e.key)
  }

  const handleLogout = () => {
    localStorage.removeItem('token')
    router.push('/login')
  }

  return (
    <Layout className="min-h-screen flex flex-col !bg-white h-full">
      {/* 顶部导航栏 */}
      <Header className="flex justify-between items-center px-10 !bg-white shadow-md flex-shrink-0">
        <div className="flex items-center gap-8">
          <h1
            className="text-xl font-bold text-blue-600 cursor-pointer"
            onClick={() => router.push('/home')}
          >
            🌟 go_mcis
          </h1>
          <Menu
            mode="horizontal"
            selectedKeys={[pathname]}
            items={menuItems}
            onClick={handleMenuClick}
            className="border-0 bg-transparent"
          />
        </div>

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

      {/* 子页面内容区（自动填满剩余空间） */}
      <Content className="flex-1 bg-gray-50 p-10 overflow-auto h-full">
        {children}
      </Content>
    </Layout>
  )
}
