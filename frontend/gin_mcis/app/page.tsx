'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { Form, Input, Button, message, Card } from 'antd'
import { LockOutlined, UserOutlined } from '@ant-design/icons'

export default function LoginPage() {
  const router = useRouter()
  const [loading, setLoading] = useState(false)
  const [msg, contextHolder] = message.useMessage()

  const handleLogin = async (values: any) => {
    setLoading(true)

    // 模拟登录请求
    setTimeout(() => {
      setLoading(false)
      if (values.username === 'admin' && values.password === '123456') {
        msg.success('🎉 登录成功！正在跳转...')
        // 模拟延时跳转
        setTimeout(() => {
          router.push('/home')
        }, 500)
      } else {
        msg.error('用户名或密码错误')
      }
    }, 1200)
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-gradient-to-br from-blue-100 via-blue-50 to-white">
      {contextHolder}
      <Card className="shadow-2xl w-[360px] rounded-2xl border border-gray-100 bg-white/90 backdrop-blur-md">
        <h1 className="text-2xl font-semibold text-center mb-6 text-blue-600">
          欢迎登录 👋
        </h1>

        <Form
          name="login"
          layout="vertical"
          onFinish={handleLogin}
          autoComplete="off"
        >
          <Form.Item
            name="username"
            label="用户名"
            rules={[{ required: true, message: '请输入用户名！' }]}
          >
            <Input
              prefix={<UserOutlined className="text-gray-400" />}
              placeholder="admin"
              size="large"
            />
          </Form.Item>

          <Form.Item
            name="password"
            label="密码"
            rules={[{ required: true, message: '请输入密码！' }]}
          >
            <Input.Password
              prefix={<LockOutlined className="text-gray-400" />}
              placeholder="123456"
              size="large"
            />
          </Form.Item>

          <Form.Item className="mt-6">
            <Button
              type="primary"
              htmlType="submit"
              size="large"
              block
              loading={loading}
              className="rounded-lg bg-blue-600 hover:bg-blue-700"
            >
              登录
            </Button>
          </Form.Item>
        </Form>

        <p className="text-center text-sm text-gray-500 mt-4">
          没有账号？<a href="#" className="text-blue-500 hover:underline">注册一个</a>
        </p>
      </Card>
    </div>
  )
}
