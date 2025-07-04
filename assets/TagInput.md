# TagInput 标签输入框

标签输入框能够将输入的内容生成标签。

## API 参考

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| ---- | ---- | ---- | ------ | ---- |
| addOnBlur | 是否在 blur 事件触发时，将当前 input 的值自动创建成 tag | boolean | false | 1.20.0 |
| allowDuplicates | 是否允许添加相同 tag | boolean | true | 1.20.0 |
| autoFocus | 初始渲染时是否自动 focus | boolean | false | 1.29.0 |
| className | 样式类名 | string | - | 1.19.0 |
| defaultValue | 初始标签 | string[] | - | 1.19.0 |
| disabled | 是否禁用 | boolean | false | 1.19.0 |
| inputValue | 当前输入框，配合 onInputChange 实现受控 | string | - | 1.19.0 |
| maxLength | 单个标签的最大长度 | number | - | 1.19.0 |
| max | 允许标签的最大数量 | number | - | 1.21.0 |
| maxTagCount | 标签的最大展示数量，超出后将以 +N 形式展示 | number | - | 1.21.0 |
| showRestTagsPopover | 超过 maxTagCount 时，hover 到 +N 是否通过 Popover 显示剩余内容 | boolean | true | 1.21.0 |
| restTagsPopoverProps | Popover 的配置属性，具体参考 Popover | PopoverProps | {} | 1.21.0 |
| showContentTooltip | 标签长度过长截断时，hover 是否通过 Tooltip/Popover 显示全部内容 | boolean \| {type: 'tooltip'\|'popover', opts: object} | true | 1.21.0 |
| placeholder | 占位默认值 | string | - | 1.19.0 |
| prefix | 前缀标签 | ReactNode | - | 1.19.0 |
| preventScroll | 指示浏览器是否应滚动文档以显示新聚焦的元素，作用于组件内的 focus 方法 | boolean |  |  |
| renderTagItem | 自定义标签渲染, 参数 onClose 于 2.23.0 版本提供 |  |  | 1.19.0 |
| separator | 设置批量输入时的分隔符 | string\|string[] | , | 1.19.0, string[]是从1.29.0开始支持 |
| showClear | 是否支持一键删除所有标签和输入内容 | boolean | false | 1.19.0 |
| size | 设置输入框尺寸, 可选: small、large、default | string | default | 1.19.0 |
| style | 内联样式 | React.CSSProperties | - | 1.19.0 |
| suffix | 后缀标签 | ReactNode | - | 1.19.0 |
| validateStatus | 设置校验状态样式, 可选: default、warning、error | string | default | 1.19.0 |
| value | 当前标签，配合 onChange 实现受控 | string[] \| undefined | - | 1.19.0 |
| draggable | 设置是否可拖拽 | boolean | false | 2.17.0 |
| expandRestTagsOnClick | 在不可拖拽的情况下，TagInput 被点击后是否展开多余的 Tag | boolean | true | 2.17.0 |
| onAdd | 添加标签时的回调 | (addedValue: string[]) => void | - | 1.19.0 |
| onBlur | 输入框失去焦点时的回调 | (e:React.MouseEvent) => void | - | 1.19.0 |
| onChange | 标签变化时的回调 | (value:string[]) => void | - | 1.19.0 |
| onExceed | 超过 max 时的回调 | (value:string[]) => void | - | 1.19.0 |
| onFocus | 输入框获取焦点时的回调 | (e:React.MouseEvent) => void | - | 1.19.0 |
| onInputChange | 输入框内容变化时的回调 | (value:string,e: React.KeyboardEvent) => void | - | 1.19.0 |
| onInputExceed | 超过 maxLength 时的回调 | (value:string) => void | - | 1.19.0 |
| onKeyDown | keydown 回调 | (e: React.KeyboardEvent) => void | - | 2.1.0 |
| onRemove | 移除标签时的回调 | (removedValue: string, idx: number) => void | - | 1.19.0 |

### Methods

| 名称 | 描述 | 版本 |
| ---- | ---- | ---- |
| blur() | 移出焦点 | 1.19.0 |
| focus() | 获取焦点 | 1.19.0 |

### Accessibility

- 支持传入 aria-label。
- 会依据 disabled 及 validateStatus props 设置 aria-disabled、aria-invalid。
- 输入框和清空按钮均具有 aria-label。

### 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ------ | ---- |
| $transition_duration-tagInput-bg | var(--semi-transition_duration-none) | 标签输入框-背景色-动画持续时间 |
| $transition_function-tagInput-bg | var(--semi-transition_function-easeIn) | 标签输入框-背景色-过渡曲线 |
| $transition_delay-tagInput-bg | var(--semi-transition_delay-none) | 标签输入框-背景色-延迟时间 |
| $transition_duration-tagInput-border | var(--semi-transition_duration-none) | 标签输入框-边框-动画持续时间 |
| $transition_function-tagInput-border | var(--semi-transition_function-easeIn) | 标签输入框-边框-过渡曲线 |
| $transition_delay-tagInput-border | var(--semi-transition_delay-none) | 标签输入框-边框-延迟时间 |
