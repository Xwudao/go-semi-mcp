# Layout 布局

用于快捷划分页面整体布局。

## 组件结构

- `Layout`：布局容器，可嵌套 `Header`、`Sider`、`Content`、`Footer` 或 `Layout` 本身。
- `Header`：顶部布局，只能放在 `Layout` 中。
- `Sider`：侧边栏，只能放在 `Layout` 中。
- `Content`：内容部分，只能放在 `Layout` 中。
- `Footer`：底部布局，只能放在 `Layout` 中。

## 注意事项

1. 布局组件采用 Flex 布局实现，无法在非现代浏览器中工作。
2. 仅实现布局，不包含背景色、文本色、宽高等样式，可通过 `style` 或 `className` 自定义。

## API 参考

### Layout

| 属性        | 说明                                         | 类型           | 默认值 |
| ----------- | -------------------------------------------- | -------------- | ------ |
| className   | 类名                                         | string         | -      |
| hasSider    | 表示子元素里有 Sider，服务端渲染时可用       | boolean        | -      |
| style       | 样式                                         | CSSProperties  | -      |
| aria-label  | aria-label 属性，提升可访问性 (>=2.3.0)      | string         | -      |
| role        | role 属性，提升可访问性 (>=2.3.0)            | string         | -      |

### Layout.Sider

| 属性         | 说明                                         | 类型                       | 默认值 |
| ------------ | -------------------------------------------- | -------------------------- | ------ |
| breakpoint   | 触发响应式布局的断点，可选值见下表           | string[]                   | -      |
| className    | 类名                                         | string                     | -      |
| style        | 样式                                         | CSSProperties              | -      |
| onBreakpoint | 响应式断点回调 `(screen: string, broken: bool) => void` | function                   | -      |
| aria-label   | aria-label 属性，提升可访问性 (>=2.3.0)      | string                     | -      |
| role         | role 属性，提升可访问性 (>=2.3.0)            | string                     | -      |

#### 响应式断点

| 尺寸 | 媒体查询                  |
| ---- | ------------------------- |
| xs   | (max-width: 575px)        |
| sm   | (min-width: 576px)        |
| md   | (min-width: 768px)        |
| lg   | (min-width: 992px)        |
| xl   | (min-width: 1200px)       |
| xxl  | (min-width: 1600px)       |

## Accessibility

- `Sider` 可传入 `aria-label` props，描述该 Sider 作用。
- `Header`、`Content`、`Footer` 可传入 `role`、`aria-label` 描述对应元素作用。
