# Upload 上传

## 组件简介
用于文件选择和上传，支持多种上传场景和自定义扩展。

## API 参考

### Upload Props

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| ---- | ---- | ---- | ---- | ---- |
| accept | 接受上传的文件类型 | string |  |  |
| action | 文件上传地址，必填 | string |  |  |
| addOnPasting | 粘贴时自动添加文件 | boolean | false | 2.43.0 |
| afterUpload | 上传后钩子 | function |  |  |
| beforeClear | 清空前回调 | function |  |  |
| beforeRemove | 移除前回调 | function |  |  |
| beforeUpload | 上传前钩子 | function |  |  |
| capture | 控件媒体拍摄方式 | boolean/string/undefined |  |  |
| className | 类名 | string |  |  |
| customRequest | 自定义上传请求 | function |  |  |
| data | 上传额外参数 | object/function | {} |  |
| defaultFileList | 默认文件列表 | Array<FileItem> | [] |  |
| directory | 支持文件夹上传 | boolean | false |  |
| disabled | 是否禁用 | boolean | false |  |
| dragIcon | 拖拽区左侧 Icon | ReactNode | `<IconUpload />` |  |
| dragMainText | 拖拽区主文本 | ReactNode | '点击上传文件或拖拽文件到这里' |  |
| dragSubText | 拖拽区帮助文本 | ReactNode | '' |  |
| draggable | 支持拖拽上传 | boolean | false |  |
| fileList | 受控文件列表 | Array<FileItem> |  |  |
| fileName | 上传文件名 | string |  |  |
| headers | 上传 headers | object/function | {} |  |
| hotSpotLocation | 照片墙点击热区位置 | string | 'end' | 2.5.0 |
| itemStyle | fileCard 样式 | CSSProperties |  |  |
| limit | 最大上传文件数 | number |  |  |
| listType | 文件列表展示类型 | string | 'list' |  |
| maxSize | 文件体积最大限制(KB) | number |  |  |
| minSize | 文件体积最小限制(KB) | number |  |  |
| multiple | 支持多选 | boolean | false |  |
| name | 上传字段名 | string | '' |  |
| onAcceptInvalid | 文件类型不符时回调 | function |  |  |
| onChange | 文件状态变化时回调 | function |  |  |
| onClear | 清空时回调 | function |  |  |
| onDrop | 拖拽释放时回调 | function |  |  |
| onError | 上传错误回调 | function |  |  |
| onExceed | 超出 limit 回调 | function |  |  |
| onFileChange | 选中文件后回调 | function |  |  |
| onOpenFileDialog | 打开文件选择弹窗时回调 | function |  |  |
| onPreviewClick | 点击文件卡片回调 | function |  |  |
| onProgress | 上传进度回调 | function |  |  |
| onPastingError | 粘贴读取失败回调 | function |  | 2.43.0 |
| onRemove | 移除文件回调 | function |  |  |
| onRetry | 上传重试回调 | function |  |  |
| onSizeError | 文件尺寸非法回调 | function |  |  |
| onSuccess | 上传成功回调 | function |  |  |
| picHeight | 图片墙图片高度 | string/number |  | 2.42.0 |
| picWidth | 图片墙图片宽度 | string/number |  | 2.42.0 |
| previewFile | 自定义预览逻辑 | function |  |  |
| prompt | 自定义提示文本 | ReactNode |  |  |
| promptPosition | 提示文本位置 | string | 'right' |  |
| renderFileItem | 自定义 fileCard 渲染 | function |  |  |
| renderFileOperation | 自定义操作区 | function |  | 2.5.0 |
| renderPicClose | 自定义照片墙 close 按钮 | function |  | 2.75.0 |
| renderPicInfo | 自定义照片墙信息 | function |  | 2.2.0 |
| renderPicPreviewIcon | 自定义预览图标 | function |  | 2.5.0 |
| renderThumbnail | 自定义缩略图 | function |  | 2.2.0 |
| showClear | 是否展示清空按钮 | boolean | true |  |
| showPicInfo | 是否显示图片信息 | boolean | false | 2.2.0 |
| showReplace | 上传成功后展示替换按钮 | boolean | false |  |
| showRetry | 上传失败后展示重试按钮 | boolean | true |  |
| showTooltip | 文件名超长时展示 tooltip | boolean/object | true |  |
| showUploadList | 是否显示文件列表 | boolean | true |  |
| style | 样式 | CSSProperties |  |  |
| transformFile | 上传前自定义转换 | function |  |  |
| uploadTrigger | 触发上传时机 | string | 'auto' |  |
| validateMessage | 整体错误信息 | ReactNode |  |  |
| withCredentials | 是否带上 Cookie | boolean | false |  |

### FileItem Interface

| 属性 | 说明 | 类型 |
| ---- | ---- | ---- |
| event | xhr event | event |
| fileInstance | 原始文件对象 | File |
| name | 文件名 | string |
| percent | 上传进度百分比 | number |
| preview | 是否预览 | boolean |
| response | 上传响应 | any |
| shouldUpload | 是否继续上传 | boolean |
| showReplace | 是否展示替换按钮 | boolean |
| showRetry | 是否展示重试按钮 | boolean |
| size | 文件大小(KB) | string |
| status | 文件状态 | string |
| uid | 文件唯一标识符 | string |
| url | 文件 url | string |
| validateMessage | 校验信息 | ReactNode/string |

### Methods

| 名称 | 描述 | 类型 | 版本 |
| ---- | ---- | ---- | ---- |
| insert | 上传文件，支持插入指定位置 | (files: Array<File>, index?: number) => void | 2.2.0 |
| upload | 手动开始上传 | () => void |  |
| openFileDialog | 打开文件选择窗口 | () => void | 2.21.0 |

### Accessibility

- 可交互控件，支持点击/拖拽选择文件。
- 可点击元素添加 `role="button"`。
- 文件列表添加 `role="list"` 并用 `aria-label` 描述。

### 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ---- | ---- |
| $color-upload-text | var(--semi-color-text-0) | 默认文本颜色 |
| $color-upload_assist-text | var(--semi-color-text-2) | 辅助文本颜色 |
| $color-upload-border | var(--semi-color-border) | 描边颜色 |
| $color-upload_card-bg-hover | var(--semi-color-fill-1) | 卡片背景色-悬浮 |
| $color-upload_card-bg | var(--semi-color-fill-0) | 卡片背景色-默认 |
| $color-upload_card_fail-bg-hover | var(--semi-color-danger-light-hover) | 失败卡片背景色-悬浮 |
| $color-upload_card_fail-bg | var(--semi-color-danger-light-default) | 失败卡片背景色-默认 |
| $color-upload_clear-text | var(--semi-color-primary) | 清空按钮文本色 |
| $color-upload_drag_area-bg-hover | var(--semi-color-primary-light-default) | 拖拽区背景色-悬浮 |
| $color-upload_drag_area-border-hover | var(--semi-color-primary) | 拖拽区描边色-悬浮 |

