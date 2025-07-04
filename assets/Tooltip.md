# Tooltip 工具提示

## Tooltip Props

| 属性                | 说明                                                         | 类型                       | 默认值      | 版本      |
| ------------------- | ------------------------------------------------------------ | -------------------------- | ----------- | --------- |
| autoAdjustOverflow  | 弹出层被遮挡时是否自动调整方向                               | boolean                    | true        | -         |
| arrowPointAtCenter  | “小三角”是否指向元素中心，需要同时传入"showArrow=true"       | boolean                    | true        | -         |
| content             | 弹出层内容                                                   | string\|ReactNode          | -           | -         |
| className           | 弹出层的样式名                                               | string                     | -           | -         |
| clickToHide         | 点击弹出层及内部任一元素时是否自动关闭弹层                   | boolean                    | false       | -         |
| disableFocusListener| trigger为hover时，不响应键盘聚焦弹出浮层事件                 | boolean                    | false       | 2.17.0    |
| getPopupContainer   | 指定父级 DOM，弹层将会渲染至该 DOM 中                        | function():HTMLElement     | () => document.body | -   |
| keepDOM             | 关闭时是否保留内部组件不销毁                                 | boolean                    | false       | 2.31.0    |
| margin              | 计算溢出时的增加的冗余值                                     | number                     | 0           | 2.23.0    |
| mouseEnterDelay     | 鼠标移入后，延迟显示的时间，单位毫秒                         | number                     | 50          | -         |
| mouseLeaveDelay     | 鼠标移出后，延迟消失的时间，单位毫秒                         | number                     | 50          | -         |
| motion              | 是否展示弹出层动画                                           | boolean                    | true        | -         |
| position            | 弹出层展示位置                                               | string                     | 'top'       | -         |
| prefixCls           | 弹出层 wrapper div 的 className 前缀                        | string                     | 'semi-tooltip' | -     |
| preventScroll       | 指示浏览器是否应滚动文档以显示新聚焦的元素                   | boolean                    | -           | -         |
| rePosKey            | 可以更新该项值手动触发弹出层的重新定位                       | string\|number             | -           | -         |
| style               | 弹出层的内联样式                                             | object                     | -           | -         |
| spacing             | 弹出层与 children 元素的距离，单位 px                        | number                     | 8           | -         |
| showArrow           | 是否显示箭头三角形                                           | boolean                    | true        | -         |
| stopPropagation     | 是否阻止弹层上的点击事件冒泡                                 | boolean                    | false       | -         |
| transformFromCenter | 是否从包裹的元素中心变换，影响动效变换的 transform-origin     | boolean                    | true        | -         |
| trigger             | 触发展示的时机，可选 hover / focus / click / custom / contextMenu | string              | 'hover'     | -         |
| visible             | 是否展示弹出层, 需配合 trigger='custom' 使用                 | boolean                    | -           | -         |
| wrapperClassName    | 外层 span 的样式类名                                         | string                     | -           | -         |
| wrapperId           | 弹出层 wrapper 节点的 id                                     | string                     | -           | 2.11.0    |
| zIndex              | 弹层层级                                                     | number                     | 1060        | -         |
| onVisibleChange     | 弹出层展示/隐藏时触发的回调                                  | function(isVisible:boolean)| -           | -         |
| onClickOutSide      | 点击非Children、非浮层内部区域时的回调                       | function(e:event)          | -           | 2.1.0     |

## Accessibility

- Tooltip 具有 tooltip role，遵循 WAI-ARIA 规范
- Tooltip 的 content wrapper 会自动添加 id 属性，用于与 children 的 aria-describedby 匹配

## 设计变量

| 变量                                   | 默认值                           | 用法                       |
| -------------------------------------- | -------------------------------- | -------------------------- |
| $animation_duration-tooltip_in         | 100ms                            | tooltip-弹出动画-动画持续时间 |
| $animation_duration-tooltip_out        | 100ms                            | tooltip-收起动画-动画持续时间 |
| $animation_function-tooltip_in         | cubic-bezier(0.215, 0.61, 0.355, 1) | tooltip-弹出动画-动画插值函数 |
| $animation_function-tooltip_out        | cubic-bezier(0.215, 0.61, 0.355, 1) | tooltip-收起动画-动画插值函数 |
