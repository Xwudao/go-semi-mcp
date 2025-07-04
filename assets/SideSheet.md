# SideSheet 滑动侧边栏

## 属性

| 属性                | 说明                                                         | 类型                        | 默认值         | 版本    |
| ------------------- | ------------------------------------------------------------ | --------------------------- | -------------- | ------- |
| afterVisibleChange  | 面板展示/隐藏时动画结束触发的回调                            | (isVisible: boolean) => void| -              | 1.0.0   |
| bodyStyle           | 面板内容的样式                                               | CSSProperties               | -              | -       |
| className           | 类名                                                         | string                      | -              | -       |
| closable            | 是否允许通过右上角的关闭按钮关闭                             | boolean                     | true           | -       |
| closeIcon           | 关闭按钮的 icon                                              | ReactNode                   | <IconClose />  | -       |
| closeOnEsc          | 允许通过键盘事件 Esc 触发关闭                                | boolean                     | false          | 1.0.0   |
| disableScroll       | 默认渲染在 document.body 层时是否禁止 body 的滚动            | boolean                     | true           | -       |
| footer              | 侧边栏底部                                                   | ReactNode                   | null           | 1.3.0   |
| getPopupContainer   | 指定父级 DOM，弹层将会渲染至该 DOM 中                        | () => HTMLElement           | -              | 0.29.0  |
| headerStyle         | 面板头部的样式                                               | CSSProperties               | -              | 1.0.0   |
| height              | 高度，位置为 top 或 bottom 时生效                            | number \| string            | 400            | -       |
| keepDOM             | 关闭 SideSheet 时是否保留内部组件不销毁                      | boolean                     | false          | 1.18.0  |
| mask                | 是否显示遮罩                                                 | boolean                     | true           | -       |
| maskClosable        | 是否允许通过点击遮罩来关闭面板                               | boolean                     | true           | -       |
| maskStyle           | 遮罩的样式                                                   | CSSProperties               | -              | -       |
| motion              | 是否允许动画                                                 | boolean                     | true           | -       |
| placement           | 侧边栏滑出位置，支持top, bottom, left, right                | string                      | right          | -       |
| size                | 尺寸，支持 small(448px)， medium(684px), large(920px)        | string                      | small          | 0.29.0  |
| style               | 可用于设置样式                                               | CSSProperties               | -              | -       |
| title               | 面板的标题                                                   | ReactNode                   | -              | -       |
| visible             | 面板是否可见                                                 | boolean                     | false          | -       |
| width               | 宽度，位置为 left 或 right 时生效                            | number \| string            | 448            | -       |
| zIndex              | 弹层 z-index 值                                              | number                      | 1000           | 0.29.0  |
| onCancel            | 取消面板时的回调函数                                         | (e: MouseEvent) => void     | -              | -       |

## Accessibility

- SideSheet 具有 dialog role 来表示它是一个弹窗组件， 内部 header 具有 heading role 表明是 header。

## 设计变量

| 变量                                    | 默认值                                      | 用法                       |
| --------------------------------------- | ------------------------------------------- | -------------------------- |
| $animation_duration-sideSheet_mask-show | 180ms                                      | 侧边栏打开时-蒙层颜色-动画持续时间 |
| $animation_function-sideSheet_mask-show | cubic-bezier(0.25, 0.46, 0.45, 0.94)        | 侧边栏打开时-蒙层颜色-过渡曲线   |
| $animation_delay-sideSheet_mask-show    | 0ms                                        | 侧边栏打开时-蒙层颜色-延迟时间   |
| $animation_duration-sideSheet_mask-hide | 180ms                                      | 侧边栏关闭时-蒙层颜色-动画持续时间 |
| $animation_function-sideSheet_mask-hide | cubic-bezier(0.25, 0.46, 0.45, 0.94)        | 侧边栏关闭时-蒙层颜色-过渡曲线   |
| $animation_delay-sideSheet_mask-hide    | 0ms                                        | 侧边栏关闭时-蒙层颜色-延迟时间   |
| $animation_duration-sideSheet_inner-show| 180ms                                      | 侧边栏打开-动画持续时间          |
| $animation_function-sideSheet_inner-show| cubic-bezier(0.25, 0.46, 0.45, 0.94)        | 侧边栏打开-过渡曲线              |
| $animation_delay-sideSheet_inner-show   | 0ms                                        | 侧边栏打开-延迟时间              |
| $animation_duration-sideSheet_inner-hide| 180ms                                      | 侧边栏关闭-动画持续时间          |
