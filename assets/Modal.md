# Modal 模态对话框

## API 参考

### Modal Props

| 属性                | 说明                                                                                      | 类型                                      | 默认值         |
|---------------------|-------------------------------------------------------------------------------------------|-------------------------------------------|----------------|
| afterClose          | 对话框完全关闭后的回调函数                                                                | () => void                                | 无             |
| bodyStyle           | 对话框内容的样式                                                                          | CSSProperties                             | 无             |
| cancelButtonProps   | 取消按钮的 props                                                                          | ButtonProps                               | 无             |
| cancelText          | 取消按钮的文字                                                                            | string                                    | 无             |
| centered            | 是否居中显示                                                                              | boolean                                   | false          |
| className           | 可用于设置样式类名                                                                        | string                                    | 无             |
| closable            | 是否显示右上角的关闭按钮                                                                  | boolean                                   | true           |
| closeIcon           | 关闭按钮的 icon                                                                           | ReactNode                                 | <IconClose />  |
| closeOnEsc          | 允许通过键盘事件 Esc 触发关闭                                                             | boolean                                   | true           |
| confirmLoading      | 确认按钮 loading                                                                          | boolean                                   | false          |
| content             | 对话框内容                                                                                | ReactNode                                 | 无             |
| footer              | 对话框底部                                                                                | ReactNode                                 | 无             |
| fullScreen          | 对话是否是全屏（会覆盖 width height）                                                     | boolean                                   | false          |
| getPopupContainer   | 指定父级 DOM，弹层将会渲染至该 DOM 中                                                     | () => HTMLElement                         | () => document.body |
| hasCancel           | 是否显示取消按钮                                                                          | boolean                                   | true           |
| header              | 对话框头部                                                                                | ReactNode                                 | 无             |
| height              | 高度                                                                                      | number                                    | 无             |
| icon                | 自定义 icon                                                                              | ReactNode                                 | -              |
| keepDOM             | 关闭对话框时是否保留内部组件不销毁                                                        | boolean                                   | false          |
| lazyRender          | 配合 keepDOM 使用，为 true 时挂载时不会渲染对话框组件                                     | boolean                                   | true           |
| mask                | 是否显示遮罩                                                                              | boolean                                   | true           |
| maskClosable        | 是否允许通过点击遮罩来关闭对话框                                                          | boolean                                   | true           |
| maskStyle           | 遮罩的样式                                                                                | CSSProperties                             | 无             |
| modalContentClass   | 可用于设置对话框内容的样式类名                                                            | string                                    | 无             |
| modalRender         | 自定义渲染 Modal                                                                          | (modal: ReactNode) => ReactNode           | -              |
| motion              | 动画效果开关                                                                              | boolean                                   | true           |
| okButtonProps       | 确认按钮的 props                                                                          | ButtonProps                               | 无             |
| okText              | 确认按钮的文字                                                                            | string                                    | 无             |
| okType              | 确认按钮的类型, 可选: 'primary'、'secondary'、'tertiary'、'warning'、'danger'             | string                                    | primary        |
| preventScroll       | 指示浏览器是否应滚动文档以显示新聚焦的元素，作用于组件内的 focus 方法，不包含用户传入的组件 | boolean                                   |                |
| size                | 对话框宽度尺寸，支持 small(448px)， medium(684px), large(920px)，full-width(100vw - 64px) | string                                    | small          |
| style               | 可用于设置样式                                                                            | CSSProperties                             | 无             |
| title               | 对话框的标题                                                                              | ReactNode                                 | 无             |
| visible             | 对话框是否可见                                                                            | boolean                                   | false          |
| width               | 宽度                                                                                      | number                                    | 448            |
| zIndex              | 遮罩的 z-index 值                                                                         | number                                    | 1000           |
| onCancel            | 取消对话框时的回调函数，返回 Promise 时，取消按钮会出现 loading 态                        | (e: any) => void \| Promise<any>          | 无             |
| onOk                | 点击确认按钮时的回调函数，返回 Promise 时，确认按钮会出现 loading 态                      | (e: any) => void \| Promise<any>          | 无             |

---

### Modal.method()

Modal.info / Modal.success / Modal.error / Modal.warning / Modal.confirm

| 属性              | 说明                                         | 类型                        | 默认值   |
|-------------------|----------------------------------------------|-----------------------------|----------|
| bodyStyle         | 对话框内容的样式                             | CSSProperties               | 无       |
| cancelButtonProps | 取消按钮的 props                             | ButtonProps                 | 无       |
| cancelText        | 取消按钮的文字                               | string                      | 无       |
| centered          | 是否居中显示                                 | boolean                     | false    |
| className         | 可用于设置样式类名                           | string                      | 无       |
| closable          | 是否显示右上角的关闭按钮                     | boolean                     | true     |
| confirmLoading    | 确认按钮 loading                             | boolean                     | false    |
| content           | 对话框内容                                   | ReactNode                   | 无       |
| footer            | 对话框底部                                   | ReactNode                   | 无       |
| footerFill        | 底部按钮是否撑满 (>= 2.xx.0 )                | boolean                     | false    |
| header            | 对话框头部                                   | ReactNode                   | 无       |
| height            | 高度                                         | number                      | 无       |
| icon              | 自定义 icon                                  | ReactNode                   | -        |
| mask              | 是否显示遮罩                                 | boolean                     | true     |
| maskClosable      | 是否允许通过点击遮罩来关闭对话框             | boolean                     | true     |
| maskStyle         | 遮罩的样式                                   | CSSProperties               | 无       |
| modalContentClass | 可用于设置对话框内容的样式类名               | string                      | 无       |
| modalRender       | 自定义渲染 Modal                             | (modal: ReactNode) => ReactNode | -    |
| okButtonProps     | 确认按钮的 props                             | ButtonProps                 | 无       |
| okText            | 确认按钮的文字                               | string                      | 无       |
| okType            | 确认按钮的类型                               | string                      | primary  |
| style             | 可用于设置样式                               | CSSProperties               | 无       |
| title             | 对话框的标题                                 | ReactNode                   | 无       |
| width             | 宽度                                         | number                      | 520      |
| zIndex            | 遮罩的 z-index 值                            | number                      | 1000     |
| onCancel          | 取消回调，参数为关闭函数，返回 promise 时 resolve 后自动关闭 | (e: any) => void \| Promise<any> | 无 |
| onOk              | 点击确定回调，参数为关闭函数，返回 promise 时 resolve 后自动关闭 | (e: any) => void \| Promise<any> | 无 |

---

### Modal.destroyAll

使用 `Modal.destroyAll()` 可以销毁命令式及以上.info()等创建的弹窗。

---

### Modal.useModal

当你需要使用 Context 时，可以通过 `Modal.useModal` 创建一个 contextHolder 插入相应的节点中。此时通过 hooks 创建的 Modal 将会得到 contextHolder 所在位置的所有上下文。创建的 modal 对象拥有与 Modal.method 相同的创建通知方法。

---

### Accessibility

- role 设置为 dialog
- aria-modal 设置为 true
- aria-labelledby 对应 Modal header
- aria-describedby 对应 Modal body
- Modal 在弹出时自动获得焦点，关闭时焦点自动回归到打开前元素
- 支持键盘 Tab/Shift+Tab 在 Modal 内部移动焦点

