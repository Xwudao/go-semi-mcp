# Anchor 锚点

## Anchor 属性

| 属性           | 说明                                                                 | 类型                                                                 | 默认值      | 版本      |
| -------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | ----------- | --------- |
| autoCollapse   | 滚动时动态显示下一级锚点                                             | boolean                                                              | false       |           |
| className      | 类名                                                                 | string                                                               | -           |           |
| defaultAnchor  | 默认高亮锚点                                                         | string                                                               | -           | 1.20.0    |
| getContainer   | 指定滚动的容器                                                       | () => HTMLElement                                                    | window      |           |
| maxHeight      | 组件的 max-height，超出时显示滚动条                                  | string \| number                                                     | 750px       |           |
| maxWidth       | 组件的 max-width，超出时显示...                                      | string \| number                                                     | 200px       |           |
| offsetTop      | 滚动内容距离容器顶部达到指定偏移量时触发                             | number                                                               | 0           |           |
| onChange       | 改变锚点的回调函数                                                   | (currentLink, previousLink) => void                                  | -           |           |
| onClick        | 点击锚点回调函数                                                     | (event, currentLink) => void                                         | -           |           |
| position       | Tooltip 显示位置，可选值同 Tooltip 组件 position                     | string                                                               | -           |           |
| railTheme      | 滑轨主题，可选值：primary，tertiary，muted                           | string                                                               | primary     |           |
| scrollMotion   | 是否开启滚动动画                                                     | boolean                                                              | false       |           |
| showTooltip    | 文字缩略时是否显示 Tooltip 及相关配置                                | boolean \| {type: 'tooltip'\|'popover', opts: object}                | false       |           |
| size           | 锚点尺寸，可选值： small，default                                    | string                                                               | default     |           |
| style          | 样式对象                                                             | object                                                               | -           |           |
| targetOffset   | 锚点滚动时距离顶部偏移量                                             | number                                                               | 0           | 1.9.0     |

## Anchor.Link 属性

| 属性      | 说明             | 类型             | 默认值 | 版本    |
| --------- | ---------------- | ---------------- | ------ | ------- |
| className | 类名             | string           | -      |         |
| disabled  | 禁用，不响应点击 | boolean          | false  | 1.20.0  |
| href      | 跳转的链接       | string           | -      |         |
| style     | 样式对象         | object           | -      |         |
| title     | 文字内容         | string\|ReactNode| -      |         |

## FAQ

**为何我的 Link 没有高亮和滑动跟随？**
- 检查 href 是否正确，确保文档中存在该 id。
- 检查滚动容器设置是否正确，默认是 window，如有自定义容器需正确传递 getContainer。

## 设计变量

| 变量                                   | 默认值                                 | 用法                       |
| -------------------------------------- | -------------------------------------- | -------------------------- |
| $transition_duration-anchor_title-text | var(--semi-transition_duration-none)   | 锚点标题文字-动画持续时间  |
| $transition_function-anchor_title-text | var(--semi-transition_function-easeIn) | 锚点标题文字-过渡曲线      |
| $transition_delay-anchor_title-text    | var(--semi-transition_delay-none)      | 锚点标题文字-延迟时间      |
