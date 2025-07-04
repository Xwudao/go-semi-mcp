# Dropdown 下拉框

## API 参考

### Dropdown

| 属性                | 说明                                                                 | 类型                                      | 默认值      | 版本      |
|---------------------|----------------------------------------------------------------------|-------------------------------------------|-------------|-----------|
| autoAdjustOverflow  | 弹出层被遮挡时是否自动调整方向                                       | boolean                                   | true        |           |
| closeOnEsc          | 在 trigger 或弹出层按 Esc 键是否关闭面板，受控时不生效               | boolean                                   | true        | 2.13.0    |
| className           | 下拉弹层外层样式类名                                                 | string                                    |             |           |
| children            | 触发弹出层的 Trigger 元素                                            | ReactNode                                 |             |           |
| clickToHide         | 在弹出层内点击时是否自动关闭弹出层                                   | boolean                                   |             |           |
| contentClassName    | 下拉菜单根元素类名                                                   | string                                    |             |           |
| disableFocusListener| trigger为hover时，不响应键盘聚焦弹出浮层事件                         | boolean                                   | false       | 2.17.0    |
| getPopupContainer   | 指定父级 DOM，弹层渲染至该 DOM                                       | function():HTMLElement                    | () => document.body |      |
| keepDOM             | 关闭时是否保留内部组件 DOM 不销毁                                    | boolean                                   | false       | 2.31.0    |
| margin              | 弹出层溢出时的冗余值，作用同 Tooltip margin                         | number\|object                            |             | 2.25.0    |
| mouseEnterDelay     | 鼠标移入 Trigger 后延迟显示时间（ms，仅 hover/focus 时生效）         | number                                    | 50          |           |
| mouseLeaveDelay     | 鼠标移出弹出层后延迟消失时间（ms，仅 hover/focus 时生效）            | number                                    | 50          |           |
| menu                | 通过 JSON Array 快速配置 Dropdown 内容                              | Array<DropdownMenuItem>                   | []          |           |
| position            | 弹出菜单位置，常用："bottom", "bottomLeft", "bottomRight"            | string                                    | "bottom"    |           |
| render              | 弹出层内容，由 Dropdown.Menu 及 Dropdown.Item、Dropdown.Title 构成   | ReactNode                                 |             |           |
| rePosKey            | 更新该项值手动触发弹出层重新定位                                     | string \| number                          |             |           |
| spacing             | 弹出层与 Trigger 元素的距离（px）                                   | number                                    | 4           |           |
| style               | 弹出层内联样式                                                       | object                                    |             |           |
| showTick            | 是否自动在 active 的 Dropdown.Item 左侧展示勾                       | boolean                                   | false       |           |
| stopPropagation     | 是否阻止弹出层点击事件冒泡                                           | boolean                                   | false       |           |
| trigger             | 触发下拉的行为："hover", "focus", "click", "custom", "contextMenu"   | string                                    | "hover"     |           |
| visible             | 是否显示菜单，需配合 trigger custom 使用                             | boolean                                   |             |           |
| zIndex              | 弹出层 z-index 值                                                    | number                                    | 1050        |           |
| onClickOutSide      | 弹出层展示时，点击非 Children/弹出层内部区域时回调                  | function(e:event)                         |             | 2.1.0     |
| onEscKeyDown        | 在 trigger 或弹出层按 Esc 键时调用                                   | function(e:event)                         |             | 2.13.0    |
| onVisibleChange     | 弹出层显示状态改变时的回调                                           | function(visible: boolean)                |             |           |

### Dropdown.Menu

| 属性      | 说明                           | 类型      | 默认值 |
|-----------|--------------------------------|-----------|--------|
| className | 下拉弹层菜单样式类名           | string    |        |
| children  | 下拉弹层菜单包裹的子元素       | ReactNode |        |
| style     | 下拉弹层菜单样式               | object    |        |

### Dropdown.Item

| 属性         | 说明                                                                 | 类型      | 默认值   |
|--------------|----------------------------------------------------------------------|-----------|----------|
| active       | 当前项是否激活，激活时左侧有 √，字体加粗，颜色加深                   | bool      | false    |
| className    | 样式类名                                                             | string    |          |
| disabled     | 是否禁用菜单                                                         | boolean   | false    |
| icon         | 图标                                                                 | ReactNode |          |
| style        | 内联样式                                                             | object    |          |
| type         | 类型："primary"、"secondary"、"tertiary"、"warning"、"danger"         | string    | "tertiary"|
| onClick      | 单击触发的回调事件                                                   | function  |          |
| onMouseEnter | MouseEnter 触发的回调事件                                            | function  |          |
| onMouseLeave | MouseLeave 触发的回调事件                                            | function  |          |
| onContextMenu| 点击鼠标右键触发的回调事件                                           | function  |          |

### Dropdown.Title

| 属性      | 说明         | 类型      | 默认值 |
|-----------|--------------|-----------|--------|
| className | 样式类名     | string    | ""     |
| style     | 内联样式     | object    | {}     |

### DropdownMenuItem

| 属性 | 说明                         | 类型   | 默认值 |
|------|------------------------------|--------|--------|
| node | 按钮类型：title/item/divider | string |        |
| name | 菜单文本                     | string |        |
| ...  | 其他属性与 Title、Item、Divider 属性对应 |        |        |

## Accessibility

- Dropdown.Menu `role=menu`，`aria-orientation=vertical`
- Dropdown.Item `role=menuitem`
- 支持键盘和焦点操作

## 设计变量

| 变量                                 | 默认值                                 | 用法                   |
|--------------------------------------|----------------------------------------|------------------------|
| $transition_duration-dropdown_item-bg| var(--semi-transition_duration-none)   | 下拉菜单项背景动画时长 |
| $transition_function-dropdown_item-bg| var(--semi-transition_function-easeOut)| 下拉菜单项背景过渡曲线 |
| $transition_delay-dropdown_item-bg   | 0ms                                   | 下拉菜单项背景延迟     |
