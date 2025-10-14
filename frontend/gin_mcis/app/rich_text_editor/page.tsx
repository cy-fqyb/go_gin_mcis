'use client'

import '@wangeditor/editor/dist/css/style.css'
import React, { useState, useEffect } from 'react'
import { Editor, Toolbar } from '@wangeditor/editor-for-react'
import type { IDomEditor, IEditorConfig, IToolbarConfig } from '@wangeditor/editor'

export default function RichTextEditor() {
  const [editor, setEditor] = useState<IDomEditor | null>(null)
  const [html, setHtml] = useState('<p>欢迎使用 wangEditor in Next.js + Bun 🚀</p>')

  const toolbarConfig: Partial<IToolbarConfig> = {}
  const editorConfig: Partial<IEditorConfig> = {
    placeholder: '开始编辑内容吧...',
  }

  // 销毁编辑器，防止内存泄露
  useEffect(() => {
    return () => {
      if (editor == null) return
      editor.destroy()
      setEditor(null)
    }
  }, [editor])

  return (
    <div className="border rounded-lg shadow-sm">
      <Toolbar editor={editor} defaultConfig={toolbarConfig} mode="default" />
      <Editor
        defaultConfig={editorConfig}
        value={html}
        onCreated={setEditor}
        onChange={editor => setHtml(editor.getHtml())}
        mode="default"
        style={{ height: '800px', overflowY: 'hidden' }}
      />
      <div className="p-2 border-t">
        <h3 className="font-semibold mb-1">实时预览：</h3>
        <div dangerouslySetInnerHTML={{ __html: html }} />
      </div>
    </div>
  )
}
