# Notification 通知

## API 参考

### 静态方法

- `Notification.open(options)`
- `Notification.info(options)`
- `Notification.error(options)`
- `Notification.warning(options)`
- `Notification.success(options)`
- `Notification.close(id)`
- `Notification.destroyAll()`
- `Notification.config(config)`
- `Notification.useNotification`

### options 属性

| 属性               | 说明                                                                 | 类型                       | 默认值     | 版本      |
|--------------------|----------------------------------------------------------------------|----------------------------|------------|-----------|
| content            | 通知内容                                                             | ReactNode                  | ''         |           |
| duration           | 自动关闭的延时，单位 s，设为 0 时不自动关闭                          | number                     | 3          |           |
| getPopupContainer  | 指定父级 DOM，弹层渲染至该 DOM，自定义需设置 position: relative      | () => HTMLElement          | document.body | 0.34.0 |
| icon               | 左上角 icon                                                          | ReactNode                  |            |           |
| position           | 弹出位置，可选 top、bottom、topLeft、topRight、bottomLeft、bottomRight | string                     | topRight   |           |
| showClose          | 是否展示关闭按钮                                                     | boolean                    | true       | 0.25.0    |
| theme              | 填充样式，支持 light, normal                                         | string                     | normal     | 1.0.0     |
| title              | 通知标题                                                             | ReactNode                  | ''         |           |
| zIndex             | 弹层 z-index 值，首次设置一次生效                                   | number                     | 1010       |           |
| onClick            | 点击通知的回调函数                                                   | (e: event) => void         |            | 0.27.0    |
| onClose            | 通知关闭的回调函数                                                   | () => void                 |            |           |
| onCloseClick       | 主动点击关闭按钮时的回调函数                                         | (id: string \| number) => void |        |           |

### Notification.config(config) 属性

| 属性     | 说明                   | 类型             | 默认值   | 版本    |
|----------|------------------------|------------------|----------|---------|
| bottom   | 弹出位置 bottom        | number \| string | -        | 0.25.0  |
| duration | 自动关闭延时           | number           | 3        | 0.25.0  |
| left     | 弹出位置 left          | number \| string | -        | 0.25.0  |
| position | 弹出位置               | string           | topRight | 0.25.0  |
| right    | 弹出位置 right         | number \| string | -        | 0.25.0  |
| top      | 弹出位置 top           | number \| string | -        | 0.25.0  |
| zIndex   | 弹层 z-index 值        | number           | 1010     | 0.25.0  |

### Accessibility

- 组件的 role 为 'alert'
- 通知的 aria-labelledby 标记为对应通知标题

### 设计变量

| 变量                                   | 默认值                                    | 用法                   |
|----------------------------------------|-------------------------------------------|------------------------|
| $animation_duration-notification-show  | 300ms                                     | 出现时-动画持续时间    |
| $animation_function-notification-show  | cubic-bezier(0.62, 0.63, 0, 1.13)         | 出现时-过渡曲线        |
| $animation_delay-notification-show     | 0ms                                       | 出现时-延迟时间        |
| $animation_duration-notification-hide  | 300ms                                     | 消失时-动画持续时间    |
| $animation_function-notification-hide  | cubic-bezier(0.62, 0.63, 0, 1.13)         | 消失时-过渡曲线        |
| $animation_delay-notification-hide     | 0ms                                       | 消失时-延迟时间        |
| $animation_opacity-notification-show   | 0                                         | 出现时-初始透明度      |
| $animation_opacity-notification-hide   | 0                                         | 消失时-结束透明度      |
