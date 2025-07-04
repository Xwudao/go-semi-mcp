# Select 选择器

用户可以通过 Select 选择器从一个选项集合中选中一个或多个选项，并呈现最终选择结果。

## API 参考

### Select Props

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ---- |
| allowCreate | 是否允许用户创建新条目，需配合 filter 使用。该项为 true 时不再响应 optionList 的变更 | boolean | false |
| arrowIcon | 自定义右侧下拉箭头 Icon | ReactNode |  |
| autoAdjustOverflow | 浮层被遮挡时是否自动调整方向 | boolean | true |
| autoClearSearchValue | 选中选项后，是否自动清空搜索关键字（多选+filter时生效） | boolean | true |
| autoFocus | 初始渲染时是否自动 focus | boolean | false |
| borderless | 无边框模式 | boolean |  |
| className | 类名 | string |  |
| clearIcon | 自定义清除按钮 | ReactNode |  |
| clickToHide | 已展开时，点击选择框是否自动收起下拉列表 | boolean | false |
| defaultValue | 初始选中的值 | string\|number\|array |  |
| defaultOpen | 是否默认展开下拉列表 | boolean | false |
| disabled | 是否禁用 | boolean | false |
| defaultActiveFirstOption | 是否默认高亮第一个选项 | boolean | true |
| dropdownClassName | 弹出层的 className | string |  |
| dropdownMatchSelectWidth | 下拉菜单最小宽度是否等于 Select | boolean | true |
| dropdownStyle | 弹出层的样式 | object |  |
| dropdownMargin | 弹出层溢出冗余值 | object\|number |  |
| emptyContent | 无结果时展示的内容 | string\|ReactNode |  |
| ellipsisTrigger | 多选溢出 tag 自适应处理 | boolean | false |
| expandRestTagsOnClick | 多选下点击 +N 展示剩余 tag | boolean | false |
| filter | 是否可搜索/自定义过滤 | boolean\|function | false |
| getPopupContainer | 指定父级 DOM | function | () => document.body |
| inputProps | 搜索 input 的额外配置 | object |  |
| innerTopSlot | optionList 内顶部自定义 slot | ReactNode |  |
| innerBottomSlot | optionList 内底部自定义 slot | ReactNode |  |
| loading | 下拉列表加载动画 | boolean | false |
| maxTagCount | 多选已选项超出后 +N 展示 | number |  |
| max | 多选最大可选数量 | number |  |
| maxHeight | optionList 最大高度 | string\|number | 270 |
| multiple | 是否多选 | boolean | false |
| outerTopSlot | optionList 外顶部自定义 slot | ReactNode |  |
| outerBottomSlot | optionList 外底部自定义 slot | ReactNode |  |
| optionList | 通过数组传入 Option | array |  |
| placeholder | 选择框默认文字 | ReactNode |  |
| position | 菜单展开位置 | string | 'bottomLeft' |
| prefix | 选择框前缀 | ReactNode |  |
| preventScroll | 是否阻止浮层点击冒泡 | boolean |  |
| renderCreateItem | allowCreate 时自定义创建标签渲染 | function |  |
| renderSelectedItem | 自定义已选项标签渲染 | function |  |
| renderOptionItem | 完全自定义候选项渲染 | function |  |
| restTagsPopoverProps | Popover 配置 | PopoverProps | {} |
| remote | 是否开启远程搜索 | boolean | false |
| searchPosition | 搜索框位置 | string | 'trigger' |
| size | 大小，可选 default/small/large | string | 'default' |
| style | 样式 | object |  |
| stopPropagation | 是否阻止浮层点击冒泡 | boolean | true |
| suffix | 选择框后缀 | ReactNode |  |
| showClear | 是否展示清除按钮 | boolean | false |
| showArrow | 是否展示下拉箭头 | boolean | true |
| showRestTagsPopover | 超出 maxTagCount 时 hover +N 显示 Popover | boolean | false |
| spacing | 浮层与选择器距离 | number | 4 |
| triggerRender | 自定义触发器渲染 | function |  |
| value | 当前选中的值 | string\|number\|array |  |
| validateStatus | 校验结果 warning/error/default | string | 'default' |
| virtualize | 列表虚拟化配置 | object |  |
| zIndex | 弹层 zIndex | number | 1030 |
| onBlur | 失焦回调 | function |  |
| onChange | 变化时回调 | function |  |
| onCreate | allowCreate 时创建回调 | function |  |
| onClear | 清除按钮回调 | function |  |
| onChangeWithObject | 选中项 option 作为回调 | boolean | false |
| onDropdownVisibleChange | 下拉菜单展开/收起回调 | function |  |
| onListScroll | 候选项列表滚动回调 | function |  |
| onSearch | 搜索输入回调 | function |  |
| onSelect | 被选中时回调 | function |  |
| onDeselect | 取消选中时回调（多选） | function |  |
| onExceed | 超出 max 限制时回调 | function |  |
| onFocus | 获得焦点时回调 | function |  |

### Option Props

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ---- |
| className | 样式类名 | string |  |
| disabled | 是否禁用 | boolean | false |
| label | 展示文本 | string\|ReactNode |  |
| showTick | 选中时展示 √ | boolean | true |
| style | 样式 | object |  |
| value | 属性值 | string\|number |  |

### OptGroup Props

| 属性 | 说明 | 类型 |
| ---- | ---- | ---- |
| className | 样式类名 | string |
| label | 展示文本 | ReactNode |
| style | 样式 | object |

### Methods

- `close()` 手动关闭下拉列表
- `open()` 手动展开下拉列表
- `focus()` 手动聚焦
- `clearInput()` 清空搜索框
- `deselectAll()` 清空所有已选项
- `selectAll()` 选中所有 Option
- `search(value: string, event: event)` 进行搜索

### FAQ

- label 必须唯一，避免用户困惑
- allowCreate 开启后不再响应 optionList/children 更新
- remote={true} 时关闭本地筛选
- 受控组件需配合 value/onChange 使用

### 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ---- | ---- |
| $transition_duration-select-bg | var(--semi-transition_duration-none) | 选择器-背景色-动画持续时间 |
| $transition_function-select-bg | var(--semi-transition_function-easeIn) | 选择器-背景色-过渡曲线 |
| $transition_delay-select-bg | var(--semi-transition_delay-none) | 选择器-背景色-延迟时间 |
| $transition_duration-select-border | var(--semi-transition_duration-none) | 选择器-边框-动画持续时间 |
| $transition_function-select-border | var(--semi-transition_function-easeIn) | 选择器-边框-过渡曲线 |
| $transition_delay-select-border | var(--semi-transition_delay-none) | 选择器-边框-延迟时间 |
| $transition_duration-select_option-bg | var(--semi-transition_duration-none) | 选择器-选项-动画持续时间 |
| $transition_function-select_option-bg | var(--semi-transition_function-easeIn) | 选择器-选项-过渡曲线 |
| $transition_delay-select_option-bg | var(--semi-transition_delay-none) | 选择器-选项-延迟时间 |

更多用法详见 Semi Design 官网。
