# Transfer 穿梭框

一个更直观高效的多选选择器，支持搜索、分组、树形、拖拽等功能。

## API 参考

### Transfer Props

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| ---- | ---- | ---- | ------ | ---- |
| className | 样式类名 | string |  |  |
| dataSource | 数据源 | Array&lt;Item&gt;\|Array&lt;GroupItem&gt;\|Array&lt;TreeItem&gt; | [] |  |
| defaultValue | 默认已选中值 | Array&lt;string\|number&gt; |  |  |
| disabled | 是否禁用 | boolean | false |  |
| draggable | 是否开启拖拽排序 | boolean | false |  |
| emptyContent | 自定义空状态的提示文本 | {left: ReactNode; right: ReactNode; search: ReactNode;} |  |  |
| filter | 自定义筛选逻辑，或关闭搜索框 | boolean \| (input:string, item: Item) => boolean | true |  |
| inputProps | 自定义搜索框 Input 属性 | InputProps |  |  |
| loading | 是否正在加载左侧选项 | boolean | - |  |
| onChange | 选中值变化时触发 | (values: Array&lt;string\|number&gt;, items: Array&lt;Item&gt;) => void |  |  |
| onDeselect | 取消勾选时的回调 | (item: Item) => void |  |  |
| onSearch | 搜索框输入内容变化时调用 | (inputValue: string) => void |  |  |
| onSelect | 勾选时的回调 | (item: Item) => void |  |  |
| renderSelectedHeader | 自定义右侧面板头部渲染 | (props: SelectedHeaderProps) => ReactNode |  | 2.29.0 |
| renderSelectedItem | 自定义右侧单个已选项渲染 | (item: { onRemove, sortableHandle } & Item) => ReactNode |  |  |
| renderSelectedPanel | 自定义右侧已选面板渲染 | (selectedPanelProps) => ReactNode |  | 1.11.0 |
| renderSourceHeader | 自定义左侧面板头部渲染 | (props: SourceHeaderProps) => ReactNode |  | 2.29.0 |
| renderSourceItem | 自定义左侧单个候选项渲染 | (item: { onChange, checked } & Item) => ReactNode |  |  |
| renderSourcePanel | 自定义左侧候选面板渲染 | (sourcePanelProps) => ReactNode |  | 1.11.0 |
| showPath | 树形模式下，右侧选中项是否显示选择路径 | boolean | false | 1.20.0 |
| style | 内联样式 | CSSProperties |  |  |
| treeProps | 树形模式下，传递给 Tree 的属性 | TreeProps |  | 1.20.0 |
| type | Transfer 类型，可选 list、groupList、treeList | string | 'list' | 1.20.0 |
| value | 已选中值，受控用法 | Array&lt;string\|number&gt; |  |  |

### Item Interface

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| className | 样式类名 | string |  |
| disabled | 是否禁用 | boolean | false |
| key | 唯一标识 | string \| number |  |
| label | 选项展示内容 | ReactNode |  |
| style | 内联样式 | CSSProperties |  |
| value | 选项代表的值 | string \| number |  |

### GroupItem Interface

继承 Item 的所有属性

| 属性 | 说明 | 类型 |
| ---- | ---- | ---- |
| children | 分组元素 | Array&lt;Item&gt; |
| title | 分组名称 | string |

### TreeItem Interface

继承 Item 的所有属性

| 属性 | 说明 | 类型 |
| ---- | ---- | ---- |
| children | 子元素 | Array&lt;TreeItem&gt; |

### Methods

| 名称 | 描述 |
| ---- | ---- |
| search(value: string) | 可通过 ref 调用该方法进行搜索 |

### 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ------ | ---- |
| $transition_duration_transfer_item-bg | var(--semi-transition_duration-none) | 穿梭框条目-背景色-动画持续时间 |
| $transition_function_transfer_item-bg | var(--semi-transition_function-easeIn) | 穿梭框条目-背景色-过渡曲线 |
| $transition_delay_transfer_item-bg | var(--semi-transition_delay-none) | 穿梭框条目-背景色-延迟时间 |
