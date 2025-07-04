# Collapse 折叠面板

## Collapse 属性

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| --- | --- | --- | --- | --- |
| accordion | 手风琴模式 | boolean | false | - |
| activeKey | 受控属性, 当前展开的面板的 key | string \| string[] | 无 | - |
| className | 样式类名 | string | '' | - |
| clickHeaderToExpand | 点击 Header 展开收起，否则只响应点击箭头 | boolean | true | 2.32.0 |
| collapseIcon | 自定义折叠图标 | ReactNode | <IconChevronDown /> | - |
| defaultActiveKey | 初始化选中面板的 key | string \| string[] | 无 | - |
| expandIcon | 自定义展开图标 | ReactNode | <IconChevronUp /> | - |
| expandIconPosition | 展开图标位置 | left, right | right | 1.12.0 |
| keepDOM | 是否保留隐藏的面板 DOM 树 | bool | false | 0.25.0 |
| motion | 是否开启动画 | boolean | true | 1.4.0 |
| lazyRender | 配合 keepDOM 使用，为 true 时挂载时不会渲染组件 | boolean | false | 2.54.1 |
| style | 内联 CSS 样式 | CSSProperties | {} | - |
| onChange | 切换面板的回调 | function(activeKey: string \| string[], e: event) | 无 | - |

## Collapse.Panel 属性

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| --- | --- | --- | --- | --- |
| className | 样式类名 | string | 无 | - |
| disabled | 面板是否被禁用 | boolean | false | v2.17.0 |
| extra | 自定义渲染每个面板右上角的辅助内容 | ReactNode | 无 | - |
| header | 面板头内容 | ReactNode | 无 | - |
| itemKey | 必填且唯一，选中状态匹配 activeKey，defaultActiveKey | string | 无 | - |
| onMotionEnd | 动画结束的回调 | () => void | - | 2.47.0-beta.0 |
| reCalcKey | 重新计算子节点高度 | string \| number | 无 | 1.5.0 |
| showArrow | 是否展示箭头 | boolean | true | v2.17.0 |
| style | 内联 CSS 样式 | CSSProperties | 无 | - |

## Accessibility

- 面板 header 右侧按钮 设置了 aria-hidden=true
- 面板 header 可交互部分 设置了 aria-owns 值为对应面板内容
- 面板内容 设置了 aria-hidden 随面板内容展现隐藏其值在 true 和 false 之间自动切换
- 面板 aria-disabled 与 disabled 属性同步，表示面板禁用

## 设计变量

| 变量 | 默认值 | 用法 |
| --- | --- | --- |
| $color-collapse_item-border-default | var(--semi-color-border) | 分割线颜色 |
| $color-collapse_header-text-default | var(--semi-color-text-0) | 标题文字颜色 |
| $color-collapse_header-text-disabled | var(--semi-color-disabled-text) | 标题文字颜色 禁用 |
| $color-collapse_header-icon-default | var(--semi-color-text-2) | 展开箭头图标颜色 |
| $color-collapse_header-bg-hover | var(--semi-color-fill-0) | 菜单项背景颜色 - 悬浮 |
| $color-collapse_header-bg-active | var(--semi-color-fill-1) | 菜单项背景颜色 - 按下 |
| $color-collapse_content-text-default | var(--semi-color-text-1) | 内容文字颜色 |
