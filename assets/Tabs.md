# Tabs 标签栏

## Tabs

| 属性                | 说明                                                                                   | 类型                                                                 | 默认值   |
|---------------------|----------------------------------------------------------------------------------------|----------------------------------------------------------------------|----------|
| activeKey           | 当前激活的 tab 页的 itemKey 值                                                         | string                                                               | 无       |
| arrowPosition       | 折叠模式下，左右切换箭头渲染位置 ("start" "end" "both")                                | string                                                               | 无       |
| className           | 类名                                                                                   | string                                                               | 无       |
| collapsible         | 折叠的 Tabs                                                                            | boolean                                                              | false    |
| dropdownProps       | 折叠模式下透传参数到下拉菜单的 Dropdown 组件                                           | { start: DropdownProps, end: DropdownProps }                         | 无       |
| visibleTabsStyle    | 整体滚动区域 Style                                                                     | style: CSSProperties                                                 | 无       |
| contentStyle        | 内容区域外层样式对象                                                                   | CSSProperties                                                        | 无       |
| defaultActiveKey    | 初始化选中的 tab 页的 key 值                                                           | string                                                               | '1'      |
| keepDOM             | 使用 TabPane 写法时是否渲染隐藏面板的 DOM 结构                                        | boolean                                                              | true     |
| lazyRender          | 懒渲染，仅当面板激活过才被渲染在 DOM 树中                                              | boolean                                                              | false    |
| more                | 一部分 Tab 渲染到下拉菜单中（number 或对象）                                           | number \| {count:number,render:()=>ReactNode,dropdownProps:DropDownProps}| -    |
| renderTabBar        | 二次封装标签栏                                                                         | (tabBarProps: object, defaultTabBar: React.ComponentType) => ReactNode| 无       |
| renderArrow         | 折叠滚动模式下自定义左右切换箭头渲染                                                   | (items, pos, handleArrowClick, defaultNode) => ReactNode              | 无       |
| preventScroll       | 指示浏览器是否应滚动文档以显示新聚焦的元素                                             | boolean                                                              |          |
| showRestInDropdown  | 收起的 Tab 是否展示在下拉菜单中（collapsible 时生效）                                  | boolean                                                              | true     |
| size                | 大小，large/medium/small，仅支持线性 Tabs                                              | string                                                               | large    |
| style               | 样式对象                                                                               | CSSProperties                                                        | 无       |
| tabBarExtraContent  | 标签栏右侧扩展内容                                                                     | ReactNode                                                            | 无       |
| tabList             | 标签页对象数组                                                                         | TabPane[]                                                            | 无       |
| tabPaneMotion       | 是否使用动画切换 tabs                                                                  | boolean                                                              | true     |
| tabPosition         | tab 的位置，top/left                                                                   | string                                                               | top      |
| type                | 标签栏样式，line/card/button/slash                                                     | string                                                               | line     |
| onChange            | 切换 tab 页时回调                                                                      | function(activeKey: string)                                           | 无       |
| onTabClick          | 单击事件                                                                               | function(key: string, e: Event)                                      | 无       |
| onTabClose          | 关闭 tab 页时回调                                                                      | function(tabKey: string)                                              | 无       |
| onVisibleTabsChange | 折叠滚动模式下，溢出项切换变化回调                                                     | (visibleState:Record<string,bool>)=>void                              | 无       |

## TabPane

| 属性      | 说明                 | 类型        | 默认值 |
|-----------|----------------------|-------------|--------|
| className | 类名                 | string      | 无     |
| disabled  | 是否禁用             | boolean     | 无     |
| icon      | 标签页栏 icon        | ReactNode   | 无     |
| itemKey   | 对应 activeKey       | string      | 无     |
| style     | 样式对象             | CSSProperties| 无    |
| tab       | 标签页栏显示文字     | ReactNode   | 无     |
| closable  | 允许关闭 tab         | boolean     | false  |
