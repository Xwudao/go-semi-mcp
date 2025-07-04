# Cascader 级联选择

用于选择多级分类下的某个选项。

## API 参考

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| ---- | ---- | ---- | ------ | ---- |
| arrowIcon | 自定义右侧下拉箭头 Icon | ReactNode | - | 1.15.0 |
| autoAdjustOverflow | 是否自动调整下拉框展开方向 | boolean | true | - |
| autoMergeValue | 设置自动合并 value，选中父节点时 value 不包括子孙节点 | boolean | true | 1.28.0 |
| bottomSlot | 底部插槽 | ReactNode | - | 1.27.0 |
| borderless | 无边框模式 | boolean |  | 2.33.0 |
| changeOnSelect | 是否允许选择非叶子节点 | boolean | false | - |
| checkRelation | 多选时，节点之间选中状态的关系，可选：'related'、'unRelated' | string | 'related' | 2.71.0 |
| className | 选择框的 className 属性 | string | - | - |
| clearIcon | 可用于自定义清除按钮, showClear为true时有效 | ReactNode | - | 2.25.0 |
| defaultOpen | 设置是否默认打开下拉菜单 | boolean | false | - |
| defaultValue | 指定默认选中的条目 | string\|number\|CascaderData\|(string\|number\|CascaderData)[] | - | - |
| disabled | 是否禁用 | boolean | false | - |
| displayProp | 设置回填选项显示的属性值 | string | label | - |
| displayRender | 设置回填格式 | (selected: string[] \| Entity, idx?: number) => ReactNode | selected => selected.join('/') | - |
| dropdownMargin | 下拉菜单计算溢出时的增加的冗余值 | object\|number | - | 2.25.0 |
| dropdownClassName | 下拉菜单的 className 属性 | string | - | - |
| dropdownStyle | 下拉菜单的样式 | object | - | - |
| expandIcon | 自定义展开 icon | ReactNode | - | 2.68.0 |
| emptyContent | 当搜索无结果时展示的内容 | ReactNode | 暂无数据 | - |
| filterLeafOnly | 搜索结果是否只展示叶子结点路径 | boolean | true | 1.26.0 |
| filterRender | 自定义渲染筛选后的选项 | (props: FilterRenderProps) => ReactNode | - | 2.28.0 |
| filterSorter | 对筛选后的选项进行排序 | (first: CascaderData, second: CascaderData, inputValue: string) => number | - | 2.28.0 |
| filterTreeNode | 设置筛选，默认用 treeNodeFilterProp 的值作为要筛选的 TreeNode 的属性值 | ((inputValue: string, treeNodeString: string, data?: CascaderData) => boolean) \| boolean | false | - |
| getPopupContainer | 指定父级 DOM，下拉框将会渲染至该 DOM 中 | () => HTMLElement | () => document.body | - |
| leafOnly | 多选时设置 value 只包含叶子节点 | boolean | false | 2.2.0 |
| loadData | 异步加载数据，需要返回一个Promise | (selectOptions: CascaderData[]) => Promise<void> | - | 1.8.0 |
| max | 多选时，限制多选选中的数量，超出 max 后将触发 onExceed 回调 | number | - | 1.28.0 |
| maxTagCount | 多选时，标签的最大展示数量，超出后将以 +N 形式展示 | number | - | 1.28.0 |
| motion | 设置下拉框弹出的动画 | boolean | true | - |
| mouseEnterDelay | 鼠标移入后，延迟显示下拉框的时间，单位毫秒 | number | 50 | - |
| mouseLeaveDelay | 鼠标移出后，延迟消失下拉框的时间，单位毫秒 | number | 50 | - |
| multiple | 设置多选 | boolean | false | 1.28.0 |
| placeholder | 选择框默认文字 | string | - | - |
| position | 方向 | string | bottom | 2.16.0 |
| prefix | 前缀标签 | ReactNode | - | 0.28.0 |
| preventScroll | 指示浏览器是否应滚动文档以显示新聚焦的元素 | boolean | - | 2.15.0 |
| restTagsPopoverProps | Popover 的配置属性 | PopoverProps | {} | 1.28.0 |
| searchPlaceholder | 搜索框默认文字 | string | - | - |
| searchPosition | 设置搜索框的位置，可选: trigger、custom | string | trigger | 2.54.0 |
| separator | 自定义分隔符 | string | / | 2.2.0 |
| showClear | 是否展示清除按钮 | boolean | false | 0.35.0 |
| showNext | 设置展开 Dropdown 子菜单的方式，可选: click、hover | string | click | 1.29.0 |
| showRestTagsPopover | 超过 maxTagCount，hover 到 +N 时，是否通过 Popover 显示剩余内容 | boolean | false | 1.28.0 |
| size | 选择框大小，可选 large，small，default | string | default | - |
| stopPropagation | 是否阻止下拉框上的点击事件冒泡 | boolean | true | - |
| disableStrictly | 设置是否开启严格禁用 | boolean | false | 1.32.0 |
| style | 选择框的样式 | CSSProperties | - | - |
| suffix | 后缀标签 | ReactNode | - | 0.28.0 |
| topSlot | 顶部插槽 | ReactNode | - | 1.27.0 |
| treeData | 展示数据，具体属性参考 CascaderData | CascaderData[] | [] | - |
| treeNodeFilterProp | 搜索时输入项过滤对应的 CascaderData 属性 | string | label | - |
| triggerRender | 自定义触发器渲染方法 | (props: TriggerRenderProps) => ReactNode | - | 0.34.0 |
| validateStatus | trigger 的校验状态，仅影响展示样式 | string | default | - |
| value | （受控）选中的条目 | string\|number\|CascaderData\|(string\|number\|CascaderData)[] | - | - |
| virtualizeInSearch | 搜索列表虚拟化 | Object | - | - |
| zIndex | 下拉菜单的 zIndex | number | 1030 | - |
| enableLeafClick | 多选时，是否启动点击叶子节点选项触发勾选 | boolean | false | 2.2.0 |
| onBlur | 失焦 Cascader 的回调 | (e: MouseEvent) => void | - | - |
| onChange | 选中树节点时调用此函数 | (value: string\|number\|CascaderData\|(string\|number\|CascaderData)[]) => void | - | - |
| onChangeWithObject | 是否将选中项 option 的其他属性作为回调 | boolean | false | 1.16.0 |
| onClear | showClear 为 true 时，点击清空按钮触发的回调 | () => void | - | 1.29.0 |
| onDropdownVisibleChange | 下拉框切换时的回调 | (visible: boolean) => void | - | 0.35.0 |
| onExceed | 多选时，超出 max 后触发的回调 | (checkedItem: Entity[]) => void | - | 1.28.0 |
| onFocus | 聚焦 Cascader 的回调 | (e: MouseEvent) => void | - | - |
| onListScroll | 下拉面板滚动的回调 | (e: React.Event, panel: { panelIndex: number; activeNode: CascaderData; }) => void | - | 1.15.0 |
| onLoad | 节点加载完毕时触发的回调 | (newLoadedKeys: Set<string>, data: CascaderData) => void | - | 1.8.0 |
| onSearch | 文本框值变化时回调 | (value: string) => void | - | - |
| onSelect | 被选中时调用，返回选中项的 value | (value: string \| number \| (string \| number)[]) => void | - | - |

### CascaderData

| 属性 | 说明 | 类型 |
| ---- | ---- | ---- |
| children | 子节点 | CascaderData[] |
| disabled | 不可选状态 | boolean |
| isLeaf | 叶子节点 | boolean |
| label | 展示的文本（必填） | ReactNode |
| loading | 正在加载 | boolean |
| value | 属性值（必填） | string\|number |

### Methods

| 方法 | 说明 | 版本 |
| ---- | ---- | ---- |
| close | 手动关闭下拉列表 | v2.30.0 |
| open | 手动展开下拉列表 | v2.30.0 |
| focus | 手动聚焦 | v2.34.0 |
| blur | 手动失焦 | v2.34.0 |
| search(value: string) | 手动触发搜索，需设置 filterTreeNode 和 searchPosition | v2.54.0 |

## Accessibility

- 支持 ARIA 属性（aria-label、aria-describedby、aria-errormessage、aria-invalid、aria-labelledby、aria-required）。
- 支持键盘操作（Enter、清空、展开等）。

## 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ------ | ---- |
| $transition_duration-cascader_option-bg | var(--semi-transition_duration-none) | 级联选项-背景色-动画持续时间 |
| $transition_function-cascader_option-bg | var(--semi-transition_function-easeIn) | 级联选项-背景色-过渡曲线 |
| $transition_delay-cascader_option-bg | var(--semi-transition_delay-none) | 级联选项-背景色-延迟时间 |
