# OverflowList 折叠列表

## 简介
OverflowList 是一个行为组件，用于展示列表，并支持自适应来展示尽可能多的项目。因过长而溢出项目将折叠为一个元素。当检测到调整大小时，可见项将被重新计算。

## API 参考

### 通用属性

| 属性         | 说明                 | 类型                  | 默认值     | 版本    |
| ------------ | -------------------- | --------------------- | ---------- | ------- |
| className    | 类名                 | string                | -          | 1.1.0   |
| renderMode   | 渲染模式             | 'collapse' \| 'scroll'| 'collapse' | 1.1.0   |
| style        | OverflowList的样式   | React.CSSProperties   | -          | 1.1.0   |

### renderMode='collapse' 专属属性

| 属性               | 说明                     | 类型                                   | 默认值 | 版本  |
| ------------------ | ------------------------ | -------------------------------------- | ------ | ----- |
| collapseFrom       | 折叠方向                 | 'start' \| 'end'                       | 'end'  | 1.1.0 |
| items              | 渲染项目                 | Record<string, any>[]                  | -      | 1.1.0 |
| minVisibleItems    | 最小展示的可见项数目     | number                                 | 0      | 1.1.0 |
| onOverflow         | 溢出回调                 | (overflowItems: Record<string, any>[]) => void | - | 1.1.0 |
| overflowRenderer   | 溢出项的自定义渲染函数   | (overflowItems: Record<string, any>[]) => React.ReactNode | - | 1.1.0 |
| visibleItemRenderer| 展示项的自定义渲染函数   | (item: Record<string, any>, index: number) => React.ReactElement | - | 1.1.0 |

### renderMode='scroll' 专属属性

| 属性                  | 说明                     | 类型                                   | 默认值 | 版本    |
| --------------------- | ------------------------ | -------------------------------------- | ------ | ------- |
| items                 | 渲染项目，要求必含 key 项 | Record<string, any>[]                  | -      | 1.1.0   |
| onIntersect           | 溢出回调                 | ({[key: string]: IntersectionObserverEntry}) => void | - | 1.1.0 |
| onVisibleStateChange  | 隐藏显示状态变化回调     | (visibleState: Map<string, boolean>) => void | - | 2.61.0 |
| overflowRenderer      | 溢出项的自定义渲染函数   | (overflowItems: Record<string, any>[]) => React.ReactNode[] | - | 1.1.0 |
| threshold             | 触发溢出回调的阈值       | number                                 | 0.75   | 1.1.0   |
| visibleItemRenderer   | 展示项的自定义渲染函数   | (item: Record<string, any>, index: number) => React.ReactElement | - | 1.1.0 |
| wrapperClassName      | 滚动 wrapper 的类名      | string                                 | -      | 1.1.0   |
| wrapperStyle          | 滚动 wrapper 的样式      | React.CSSProperties                    | -      | 1.1.0   |
