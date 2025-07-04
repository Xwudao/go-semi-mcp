# AutoComplete 自动完成

## 说明
输入框自动填充。用于对输入框提供输入建议，进行自动补全的操作。

## API 参考

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| ---- | ---- | ---- | ------ | ---- |
| autoFocus | 是否自动聚焦 | bool | false | 1.16.0 |
| autoAdjustOverflow | 浮层被遮挡时是否自动调整方向 | bool | true |  |
| className | 样式类名 | string |  |  |
| clearIcon | 可用于自定义清除按钮, showClear为true时有效 | ReactNode |  | 2.25.0 |
| data | 候选项的数据源，可以为字符串数组或对象数组 | array | [] |  |
| defaultActiveFirstOption | 是否默认高亮第一个选项（按回车可直接选中） | bool | false |  |
| defaultOpen | 是否默认展开下拉菜单 | boolean | false |  |
| defaultValue | 默认值 | string |  |  |
| disabled | 是否禁用 | boolean | false |  |
| dropdownClassName | 下拉列表的 CSS 类名 | string |  |  |
| dropdownStyle | 下拉列表的内联样式 | object |  |  |
| emptyContent | data 为空时自定义下拉内容 | ReactNode | null | 1.16.0 |
| getPopupContainer | 指定下拉列表浮层的父级容器，浮层将会渲染至该 DOM 中。自定义该项时需给容器设置 position: relative 这会改变浮层 DOM 树位置，但不会改变视图渲染位置 | () => HTMLElement | () => document.body |  |
| loading | 下拉列表是否展示加载动画 | boolean | false |  |
| maxHeight | 下拉列表的最大高度 | number\|string | 300 |  |
| motion | 下拉列表出现/隐藏时，是否有动画 | boolean | true |  |
| onSelectWithObject | 点击候选项时，是否将选中项 option 的其他属性也作为回调入参。设为 true 时，onSelect 的入参类型会从 string 变为 object: { value, label, ...rest } | boolean | false | 1.23.0 |
| placeholder | 输入框默认提示文案 | string |  |  |
| position | 下拉菜单的显示位置，可选值同 tooltip 组件 | string | 'bottomLeft' |  |
| prefix | 选择框的前缀标签 | ReactNode |  | 0.23.0 |
| renderItem | 控制下拉列表候选项的渲染 | (option: string\|Item)=> React.Node |  |  |
| renderSelectedItem | 通过 renderSelectedItem 自定义下拉列表候选项被点击选中后，在选择框中的渲染内容，仅支持 String 类型的返回值 | (option: string\|Item) => string |  |  |
| showClear | 是否展示清除按钮 | boolean | false |  |
| size | 尺寸，可选small, default, large | string | default |  |
| style | 样式 | object |  |  |
| suffix | 选择框的前缀标签 | ReactNode |  |  |
| validateStatus | 校验状态，可选值default、error、warning，默认 default。仅影响展示样式 | string | 'default' | 1.14.0 |
| value | 当前值 | string\|number | 无 |  |
| zIndex | 下拉菜单的 zIndex | number |  |  |
| onBlur | 失去焦点时的回调 | Function(event) |  |  |
| onChange | 输入框变化/候选项选中时变化 | Function(value:string\|number) |  | 1.23.0 |
| onFocus | 获得焦点时的回调 | Function(event) |  |  |
| onKeyDown | keydown 回调 | (e: React.KeyboardEvent) => void |  | 2.21.0 |
| onSearch | 输入变化时的回调 | Function(value: string) |  |  |
| onSelect | 下拉菜单候选项被选中时的回调 | Function(item: string\|number\|Item) |  |  |

## Accessibility

- 支持键盘和焦点操作，详见官网文档。

## 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ------ | ---- |
| $transition_duration-autoComplete_option-bg | var(--semi-transition_duration-none) | 选择器-选项-动画持续时间 |
| $transition_function-autoComplete_option-bg | var(--semi-transition_function-easeIn) | 选择器-选项-过渡曲线 |
| $transition_delay-autoComplete_option-bg | var(--semi-transition_delay-none) | 选择器-选项-延迟时间 |
