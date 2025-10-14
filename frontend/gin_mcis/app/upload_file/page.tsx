'use client';

import { useState } from 'react';

export default function UploadPage() {
  const [file, setFile] = useState<File | null>(null);
  const [message, setMessage] = useState('');
  const [uploading, setUploading] = useState(false);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFile(e.target.files?.[0] || null);
    setMessage('');
  };

  const handleUpload = async () => {
    if (!file) {
      setMessage('请先选择一个文件');
      return;
    }

    const formData = new FormData();
    formData.append('file', file);

    setUploading(true);
    setMessage('');

    try {
      const res = await fetch('http://localhost:8081/api/public/upload', {
        method: 'POST',
        body: formData,
        credentials: 'include', // 支持携带 cookie
      });

      // 尝试解析 JSON，如果失败就用文本
      let data: any;
      try {
        data = await res.json();
      } catch {
        data = await res.text();
      }

      if (!res.ok) {
        setMessage(`❌ 上传失败：${data?.error || data || res.statusText}`);
      } else {
        setMessage(`✅ 上传成功：${data?.filename || JSON.stringify(data)}`);
      }
    } catch (err: any) {
      setMessage(`🚨 网络错误：${err.message}`);
    } finally {
      setUploading(false);
    }
  };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center gap-4 p-6">
      <h1 className="text-2xl font-bold mb-4">文件上传测试</h1>

      <input
        type="file"
        onChange={handleFileChange}
        className="border p-2 rounded"
      />

      <button
        onClick={handleUpload}
        disabled={!file || uploading}
        className={`px-4 py-2 rounded text-white ${
          uploading ? 'bg-gray-400' : 'bg-blue-600 hover:bg-blue-700'
        }`}
      >
        {uploading ? '上传中...' : '开始上传'}
      </button>

      {message && (
        <p className="mt-4 text-sm text-center whitespace-pre-wrap">
          {message}
        </p>
      )}
    </div>
  );
}
