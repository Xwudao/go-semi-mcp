# Toast 提示

Toast 提示是对用户的操作做出及时反馈，由用户的操作触发，反馈信息可以是操作的结果状态，如成功、失败、出错、警告等。

## API 参考

### 组件静态方法

- Toast.config(config)
- Toast.info(options \|\| string)
- Toast.error(options \|\| string)
- Toast.warning(options \|\| string)
- Toast.success(options \|\| string)
- Toast.close(toastId)
- Toast.destroyAll()
- Toast.useToast()

### Options

| 属性          | 说明                         | 类型                | 默认值   | 版本      |
|---------------|------------------------------|---------------------|----------|-----------|
| content       | 提示内容                     | ReactNode           | ''       |           |
| icon          | 自定义图标                   | ReactNode           |          | 0.25.0    |
| showClose     | 是否展示关闭按钮             | boolean             | true     | 0.25.0    |
| textMaxWidth  | 内容的最大宽度               | number \| string    | 450      | 0.25.0    |
| onClose       | toast 关闭的回调函数         | () => void          |          |           |
| stack         | 是否堆叠 Toast               | boolean             | false    | 2.42.0    |
| id            | 自定义 ToastId               | number              |          |           |

### Config

| 属性              | 说明                         | 类型                | 默认值   | 版本      |
|-------------------|------------------------------|---------------------|----------|-----------|
| bottom            | 弹出位置 bottom              | number \| string    | -        | 0.25.0    |
| left              | 弹出位置 left                | number \| string    | -        | 0.25.0    |
| right             | 弹出位置 right               | number \| string    | -        | 0.25.0    |
| top               | 弹出位置 top                 | number \| string    | -        | 0.25.0    |
| zIndex            | 弹层 z-index 值              | number              | 1010     |           |
| theme             | 填充样式，支持 light, normal | string              | normal   | 2.54.0    |
| duration          | 自动关闭的延时，单位 s       | number              | 3        |           |
| getPopupContainer | 指定父级 DOM                 | () => HTMLElement   | document.body | 0.34.0 |

### Accessibility

- Toast 的 role 为 alert

## 文案规范

- 保持简洁，句尾不使用句号
- 使用 名词 + 动词 的格式进行说明
- 只提供一个动作，不使用类似于「已读」类的动作

## 设计变量

| 变量                                 | 默认值   | 用法 |
|--------------------------------------|----------|------|
| $animation_duration-toast-show       | 300ms    | 暂无 |
| $animation_duration-toast-hide       | 300ms    | 暂无 |
| $animation_function-toast-show       | cubic-bezier(.22,.57,.02,1.2) | 暂无 |
| $animation_function-toast-hide       | cubic-bezier(.22,.57,.02,1.2) | 暂无 |
| $animation_delay-toast-show          | 0s       | 暂无 |
| $animation_delay-toast-hide          | 0s       | 暂无 |
| $animation_opacity-toast-show        | 0        | 暂无 |
| $animation_opacity-toast-hide        | 0        | 暂无 |
| $animation_transform_translateY-toast-show | -100% | 暂无 |
| $animation_transform_translateY-toast-hide | -100% | 暂无 |
