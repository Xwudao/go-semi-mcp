# Navigation 导航

为页面和功能提供导航的菜单列表。

## API 参考

### Nav

| 属性                | 描述                                                                 | 类型                | 默认值   |
|---------------------|----------------------------------------------------------------------|---------------------|----------|
| bodyStyle           | 导航项列表的自定义样式                                               | CSSProperties       |          |
| className           | 最外层元素的样式名                                                   | string              |          |
| defaultIsCollapsed  | 默认是否处于收起状态，仅 mode = "vertical" 时有效                    | boolean             | false    |
| defaultOpenKeys     | 初始打开的子导航 itemKey 数组，仅 mode = "vertical" 且展开时有效      | string[]            | []       |
| defaultSelectedKeys | 初始选中的导航项 key 数组                                            | string[]            | []       |
| subDropdownProps    | 控制 nav.sub 中的 dropdown 参数(v >= 2.69)                           | DropdownProps       |          |
| expandIcon          | 默认下拉箭头Icon, v>=2.36                                            | ReactNode           |          |
| footer              | 底部区域配置对象或元素，详见 Nav.Footer                              | object\|ReactNode   |          |
| getPopupContainer   | 指定弹出层容器，改变浮层 DOM 树位置                                  | Function            |          |
| header              | 头部区域配置对象或元素，详见 Nav.Header                              | object\|ReactNode   |          |
| isCollapsed         | 是否处于收起状态的受控属性，仅 mode = "vertical" 时有效              | boolean             |          |
| items               | 导航项目列表，支持嵌套                                               | object\|string[]    |          |
| limitIndent         | 解除缩进限制，水平模式只能为true                                     | boolean             | true     |
| mode                | 导航类型，可选 vertical 或 horizontal                                | string              | vertical |
| openKeys            | 受控的打开的子导航 itemKey 数组                                      | string[]            |          |
| prefixCls           | 类名前缀                                                             | string              | semi     |
| renderWrapper       | 自定义导航项外层组件，v>=2.24.0                                      |                     |          |
| selectedKeys        | 受控的导航项 key 数组                                                | string[]            |          |
| style               | 最外层元素的自定义样式                                               | CSSProperties       |          |
| subNavCloseDelay    | 子导航浮层关闭的延迟，单位 ms                                        | number              | 100      |
| subNavMotion        | 子导航折叠动画                                                       | boolean             | true     |
| subNavOpenDelay     | 子导航浮层显示的延迟，单位 ms                                        | number              | 0        |
| toggleIconPosition  | 子导航项箭头位置，可选 left 或 right                                 | string              | right    |
| tooltipHideDelay    | tooltip 隐藏的延迟，collapse 为 true 时有效，单位 ms                 | number              | 100      |
| tooltipShowDelay    | tooltip 显示的延迟，collapse 为 true 时有效，单位 ms                 | number              | 0        |
| onClick             | 点击任意导航项时触发                                                 | () => {}            |          |
| onCollapseChange    | 收起状态变化时的回调                                                 | () => {}            |          |
| onOpenChange        | 切换某个子导航项目显隐状态时触发                                     | () => {}            |          |
| onSelect            | 第一次选中某个可选中导航项目时触发                                   | () => {}            |          |

---

### Nav.Item

| 属性        | 描述                                 | 类型              | 默认值 |
|-------------|--------------------------------------|-------------------|--------|
| disabled    | 是否禁用                             | boolean           | false  |
| icon        | 导航项目图标                         | ReactNode         |        |
| indent      | icon 为空时是否保留占位，仅一级导航   | boolean           | false  |
| itemKey     | 导航项目唯一 key                     | string            | ""     |
| level       | 当前项所在嵌套层级                   | number            |        |
| link        | 导航项 href 链接                     | string            | -      |
| linkOptions | 透传给 a 标签的参数                  | object            | -      |
| text        | 导航项目文案或元素                   | string\|ReactNode | ""     |
| onClick     | 点击导航项时触发                     | () => {}          |        |
| onMouseEnter| mouse enter 时触发                   | function(e) => {} |        |
| onMouseLeave| mouse leave 时触发                   | function(e) => {} |        |

---

### Nav.Sub

| 属性         | 描述                                 | 类型              | 默认值 |
|--------------|--------------------------------------|-------------------|--------|
| disabled     | 是否禁用                             | boolean           | false  |
| dropdownProps| 弹出层 dropdown 参数配置             | DropdownProps     |        |
| dropdownStyle| 弹出层的 style                       | CSSProperties     |        |
| icon         | 导航项目图标                         | ReactNode         |        |
| indent       | icon 为空时是否保留占位，仅一级导航   | boolean           | false  |
| isCollapsed  | 是否处于收起状态，仅 mode = "vertical"| boolean           | false  |
| isOpen       | 是否打开                             | boolean           | false  |
| itemKey      | 导航项目唯一 key                     | string            | -      |
| level        | 当前项所在嵌套层级                   | number            | 0      |
| maxHeight    | 最大高度                             | number            | 999    |
| text         | 导航项目文案或组件                   | string\|ReactNode | ""     |
| onMouseEnter | mouse enter 时触发                   | function(e) => {} |        |
| onMouseLeave | mouse leave 时触发                   | function(e) => {} |        |

---

### Nav.Header

| 属性      | 描述           | 类型        | 默认值 |
|-----------|----------------|-------------|--------|
| children  | 子元素         | ReactNode   |        |
| className | 最外层样式名   | string      |        |
| link      | href 链接      | string      | -      |
| linkOptions| a 标签参数    | object      | -      |
| logo      | Logo           | ReactNode   |        |
| style     | 最外层样式     | CSSProperties|       |
| text      | Logo 文案      | ReactNode   |        |

---

### Nav.Footer

| 属性          | 描述                                         | 类型                | 默认值 |
|---------------|----------------------------------------------|---------------------|--------|
| children      | 子元素                                       | ReactNode           |        |
| className     | 最外层样式名                                 | string              |        |
| collapseButton| 是否展示底部“收起侧边栏”按钮                 | boolean\|ReactNode  | false  |
| collapseText  | “收起”按钮的文案                             | (collapsed:boolean)=>ReactNode | |
| style         | 最外层样式                                   | CSSProperties       |        |
| onClick       | 点击事件回调                                 | (event) => void     |        |
