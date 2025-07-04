# TreeSelect 树选择器

## 组件简介
树选择器用于多层级树形数据的结构化展示与选取，例如文件夹、组织架构等。

## API 参考

### TreeSelect Props

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ---- |
| arrowIcon | 自定义右侧下拉箭头Icon | ReactNode | - |
| autoAdjustOverflow | 浮层被遮挡时是否自动调整方向 | boolean | true |
| autoExpandParent | 是否自动展开父节点 | boolean | false |
| autoMergeValue | 自动合并 value，父节点选中时不包括子孙节点 | boolean | true |
| borderless | 无边框模式 | boolean | false |
| checkRelation | 多选时节点选中关系：'related'、'unRelated' | string | 'related' |
| className | 选择框 className | string | - |
| clearIcon | 自定义清除按钮 | ReactNode | - |
| clickToHide | 选择后是否自动关闭下拉，仅单选有效 | boolean | true |
| clickTriggerToHide | 面板打开时点击 Trigger 是否关闭面板 | boolean | true |
| defaultExpandAll | 初始化时是否展开所有节点 | boolean | false |
| defaultExpandedKeys | 默认展开的节点 key | string[] | - |
| defaultOpen | 默认展开下拉菜单 | boolean | false |
| defaultValue | 默认选中的条目 | - | - |
| disabled | 是否禁用 | boolean | false |
| disableStrictly | 是否严格禁用 | boolean | false |
| dropdownClassName | 下拉菜单 className | string | - |
| dropdownMatchSelectWidth | 下拉菜单宽度是否等于Select | boolean | true |
| dropdownMargin | 下拉菜单溢出冗余值 | object/number | - |
| dropdownStyle | 下拉菜单样式 | CSSProperties | - |
| emptyContent | 搜索无结果时展示内容 | ReactNode | 暂无数据 |
| expandAction | 展开逻辑：false/'click'/'doubleClick' | boolean/string | false |
| expandAll | 是否默认展开所有节点 | boolean | false |
| expandedKeys | 受控展开节点 key | string[] | - |
| expandIcon | 自定义展开图标 | ReactNode/Function | - |
| keyMaps | 自定义节点字段映射 | object | - |
| filterTreeNode | 是否支持搜索 | boolean | false |
| getPopupContainer | 指定弹层父级 DOM | function | - |
| labelEllipsis | label 超出省略 | boolean | false/true(虚拟化) |
| leafOnly | 多选时仅叶子节点生效 | boolean | false |
| loadData | 异步加载数据 | Function | - |
| loadedKeys | 受控已加载节点 key | Set<string> | - |
| maxTagCount | 最多显示多少个 tag | number | - |
| motionExpand | 是否开启节点动画 | boolean | true |
| multiple | 是否多选 | boolean | false |
| optionListStyle | optionList 样式 | CSSProperties | - |
| outerBottomSlot | 弹层底部自定义 slot | ReactNode | - |
| outerTopSlot | 弹层顶部自定义 slot | ReactNode | - |
| placeholder | 选择框默认文字 | string | - |
| position | 下拉菜单位置 | string | bottomLeft |
| prefix | 前缀标签 | ReactNode | - |
| preventScroll | 是否阻止滚动 | boolean | - |
| renderFullLabel | 完全自定义 label 渲染 | Function | - |
| renderLabel | 自定义 label 渲染 | Function | - |
| renderSelectedItem | 自定义已选项渲染 | Function | - |
| restTagsPopoverProps | Popover 配置属性 | PopoverProps | {} |
| searchAutoFocus | 搜索框自动聚焦 | boolean | false |
| searchPlaceholder | 搜索框默认文字 | string | - |
| searchPosition | 搜索框位置：dropdown/trigger | string | dropdown |
| showClear | 是否展示清除按钮 | boolean | false |
| showFilteredOnly | 搜索时只展示过滤结果 | boolean | false |
| showLine | 选项面板显示连接线 | boolean | false |
| showRestTagsPopover | 超出 maxTagCount 时 Popover 展示 | boolean | false |
| showSearchClear | 是否显示搜索框清除按钮 | boolean | true |
| size | 选择框大小：large/small/default | string | default |
| style | 选择框样式 | CSSProperties | - |
| suffix | 后缀标签 | ReactNode | - |
| treeData | 树节点数据 | TreeNodeData[] | [] |
| treeNodeFilterProp | 搜索过滤属性 | string | label |
| treeNodeLabelProp | 显示的 prop | string | label |
| triggerRender | 自定义触发器渲染 | Function | - |
| validateStatus | 校验结果：warning/error/default | string | - |
| value | 当前选中节点 value | - | - |
| virtualize | 列表虚拟化配置 | object | - |
| zIndex | 下拉菜单 zIndex | number | 1030 |
| onBlur | 失焦回调 | Function | - |
| onChange | 选中节点回调 | Function | - |
| onChangeWithObject | onChange 回调是否带对象 | boolean | false |
| onClear | 清除按钮回调 | Function | - |
| onExpand | 展开节点回调 | Function | - |
| onFocus | 聚焦回调 | Function | - |
| onLoad | 节点加载完毕回调 | Function | - |
| onSearch | 搜索回调 | Function | - |
| onSelect | 被选中时回调 | Function | - |
| onVisibleChange | 弹层展示/隐藏回调 | Function | - |

### TreeNodeData

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ---- |
| value | 属性值 | string/number | - |
| label | 展示文本 | string/ReactNode | - |
| icon | 自定义图标 | ReactNode | - |
| disabled | 是否禁用 | boolean | false |
| key | 唯一 key，必填 | string | - |
| isLeaf | 是否叶子节点 | boolean | - |

### Methods

| 名称 | 说明 |
| ---- | ---- |
| search(sugInput: string) | 外部自定义搜索框时调用，改变筛选结果 |

### Accessibility

- 自动设置 `aria-label='TreeSelect'`，支持自定义。
- 支持 `aria-describedby`、`aria-errormessage`、`aria-invalid`、`aria-labelledby`、`aria-required`。
- 子节点自动设置 `aria-disabled`、`aria-checked`、`aria-selected`、`aria-level`。

### 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ---- | ---- |
| $color-treeSelect_default-bg-default | var(--semi-color-fill-0) | 选择框背景色-默认 |
| $color-treeSelect_default-bg-hover | var(--semi-color-fill-1) | 选择框背景色-悬停 |
| $color-treeSelect_default-bg-active | var(--semi-color-fill-2) | 选择框背景色-按下 |
| $color-treeSelect_default-bg-focus | var(--semi-color-fill-0) | 选择框背景色-选中 |
| $color-treeSelect_default-border-default | transparent | 选择框描边色-默认 |
| $color-treeSelect_selection-text-default | var(--semi-color-text-0) | 回填内容文本色-默认 |
| $color-treeSelect_default-icon-default | var(--semi-color-text-2) | 图标色-默认 |
| $color-treeSelect_default-icon-hover | var(--semi-color-primary-hover) | 清空按钮色-悬停 |
| $color-treeSelect_default-icon-active | var(--semi-color-primary-active) | 清空按钮色-按下 |
| $color-treeSelect_default-border-hover | transparent | 选择框描边色-悬浮 |

