# Tree 树形控件

树型结构列表。

## API 参考

### 属性

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| ---- | ---- | ---- | ------ | ---- |
| autoExpandParent | 是否自动展开父节点，默认为 false，当组件初次挂载时为 true | boolean | false | 0.34.0 |
| autoExpandWhenDragEnter | 拖拽到节点上时自动展开 | boolean | true | 1.8.0 |
| autoMergeValue | 自动合并 value（leafOnly=false 时生效） | boolean | true | 2.61.0 |
| blockNode | 行显示节点 | boolean | true | - |
| checkRelation | 多选时节点选中关系：'related'、'unRelated' | string | 'related' | 2.5.0 |
| className | 类名 | string | - | - |
| defaultExpandAll | 初始化时是否展开所有节点 | boolean | false | - |
| defaultExpandedKeys | 默认展开的节点 key | string[] | - | - |
| defaultValue | 默认选中的条目 | string \| number \| TreeNodeData \| (string \| number \| TreeNodeData)[] | - | - |
| directory | 目录树模式 | boolean | false | - |
| disableStrictly | 严格禁用 | boolean | false | 1.4.0 |
| disabled | 禁用整个树 | boolean | false | 0.32.0 |
| draggable | 允许拖拽 | boolean | false | 1.8.0 |
| emptyContent | 搜索无结果时展示内容 | ReactNode | 暂无数据 | 0.32.0 |
| expandAction | 展开逻辑：false, 'click', 'doubleClick' | boolean \| string | false | 0.35.0 |
| expandAll | 默认展开所有节点（数据变更时也生效） | boolean | false | 1.30.0 |
| expandedKeys | （受控）展开的节点 key | string[] | - | - |
| expandIcon | 自定义展开图标 | ReactNode \| (props) => ReactNode | - | 2.75.0 |
| keyMaps | 自定义 key、label、value 字段 | object | - | 2.47.0 |
| filterTreeNode | 搜索过滤 | boolean \| (inputValue, treeNodeString, data?) => boolean | false | - |
| hideDraggingNode | 隐藏拖拽节点的 dragImg | boolean | false | 1.8.0 |
| icon | 自定义图标 | ReactNode \| (props) => ReactNode | - | - |
| labelEllipsis | label 超出省略 | boolean | false/true(virtualized) | 1.8.0 |
| leafOnly | 多选模式下仅叶子节点可选 | boolean | false | 1.18.0 |
| loadData | 异步加载数据 | (treeNode?) => Promise<void> | - | 1.0.0 |
| loadedKeys | （受控）已加载节点 key | string[] | - | 1.0.0 |
| motion | 是否开启动画 | boolean | true | - |
| multiple | 支持多选 | boolean | false | - |
| preventScroll | 控制 focus 时是否滚动 | boolean |  |  |
| renderDraggingNode | 自定义拖拽节点 dragImg | (nodeInstance, node) => HTMLElement | - | 1.8.0 |
| renderFullLabel | 完全自定义 label 渲染 | (data) => ReactNode | - | 1.7.0 |
| renderLabel | 自定义 label 渲染 | (data, searchWord?) => ReactNode | - | 1.6.0 |
| searchClassName | 搜索框 className | string | - | - |
| searchPlaceholder | 搜索框默认文字 | string | - | - |
| searchRender | 自定义搜索框渲染/隐藏 | ((props) => ReactNode) \| false | - | 0.35.0 |
| searchStyle | 搜索框样式 | CSSProperties | - | - |
| showClear | 搜索框支持清除 | boolean | true | 0.35.0 |
| showFilteredOnly | 搜索时只展示过滤结果 | boolean | false | 0.32.0 |
| showLine | 显示连接线 | boolean | false | 2.50.0 |
| style | 样式 | CSSProperties | - | - |
| treeData | 树节点数据 | TreeNodeData[] | [] | - |
| treeDataSimpleJson | 简单 JSON 形式数据 | TreeDataSimpleJson | {} | - |
| treeNodeFilterProp | 搜索过滤属性 | string | label | - |
| value | 当前选中节点 value | string \| number \| TreeNodeData \| (string \| number \| TreeNodeData)[] | - | - |
| virtualize | 列表虚拟化 | VirtualizeObj | - | 0.32.0 |
| onChange | 选中节点时回调 | (value) => void | - | - |
| onChangeWithObject | onChange 回调是否返回对象 | boolean | false | - |
| onDoubleClick | 双击事件回调 | (e, node) => void | - | 0.35.0 |
| onDragEnd | 拖拽结束回调 | (props) => void | - | 1.8.0 |
| onDragEnter | 拖拽进入回调 | (props) => void | - | 1.8.0 |
| onDragLeave | 拖拽离开回调 | (props) => void | - | 1.8.0 |
| onDragOver | 拖拽经过回调 | (props) => void | - | 1.8.0 |
| onDragStart | 拖拽开始回调 | (props) => void | - | 1.8.0 |
| onDrop | 拖拽释放回调 | (props) => void | - | 1.8.0 |
| onExpand | 展开节点时回调 | (expandedKeys, {expanded, node}) => void | - | - |
| onLoad | 节点加载完毕回调 | (loadedKeys, treeNode) => void | - | 1.0.0 |
| onContextMenu | 右键点击回调 | (e, node) => void | - | 0.35.0 |
| onSearch | 搜索时回调 | (input, filteredExpandedKeys) => void | - | - |
| onSelect | 选中时回调 | (selectedKey, selected, selectedNode) => void | - | - |

### TreeNodeData

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| value | 属性值 | string/number | - |
| label | 展示文本 | string/ReactNode | - |
| icon | 自定义图标 | ReactNode | - |
| disabled | 是否禁用 | boolean | false |
| key | required且唯一 | string | - |
| isLeaf | 是否叶子节点（异步加载时有效） | boolean | - |

### Virtualize Object

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| height | 高度 | number/string | '100%' |
| itemSize | 每行高度 | number | - |
| width | 宽度 | number/string | '100%' |

### Methods

- `search(value: string): void` 手动触发搜索
- `scrollTo({key, align}): void` 虚拟化树滚动到指定节点（2.18.0+）

### Accessibility

- 支持 aria-label
- 子节点自动设置 aria-disabled、aria-checked、aria-selected、aria-level
- role 分别为 tree、treeitem

