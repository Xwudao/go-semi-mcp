# Popover 气泡卡片

## 简介
Popover 气泡卡片是由用户自主打开的临时性浮层卡片，能够承载一些额外内容和交互行为而不影响原页面。和 Tooltip 的区别是，它可以承载更复杂的内容，而不仅仅是提示文本。

## API 参考

| 属性                  | 说明                                                         | 类型                                               | 默认值      | 版本     |
| --------------------- | ------------------------------------------------------------ | -------------------------------------------------- | ----------- | -------- |
| autoAdjustOverflow    | 是否自动调整弹出层展开方向                                   | boolean                                            | true        |          |
| arrowPointAtCenter    | “小三角”是否指向元素中心（需 showArrow=true）                | boolean                                            | true        | 0.34.0   |
| closeOnEsc            | 在 trigger 或弹出层按 Esc 是否关闭面板（受控时无效）         | boolean                                            | true        | 2.8.0    |
| content               | 显示的内容（支持函数类型）                                   | ReactNode \| ({ initialFocusRef }) => ReactNode    |             | 2.8.0    |
| clickToHide           | 点击弹出层及内部任一元素时是否自动关闭弹层                   | boolean                                            | false       | 0.24.0   |
| disableFocusListener  | trigger为hover时，不响应键盘聚焦弹出浮层事件                 | boolean                                            | true        | 2.17.0   |
| getPopupContainer     | 指定父级 DOM，弹层渲染至该 DOM                              | function():HTMLElement                             | () => document.body |         |
| guardFocus            | 焦点处于弹出层内时，Tab 是否循环                             | boolean                                            | true        | 2.8.0    |
| keepDOM               | 关闭时是否保留内部组件不销毁                                 | boolean                                            | false       | 2.31.0   |
| margin                | 弹出层计算溢出时的冗余值                                    | number\|object                                     |             | 2.25.0   |
| mouseEnterDelay       | 鼠标移入后延迟显示时间（ms，仅 hover/focus 时生效）          | number                                             | 50          |          |
| mouseLeaveDelay       | 鼠标移出后延迟消失时间（ms，仅 hover/focus 时生效）          | number                                             | 50          |          |
| rePosKey              | 更新该项值手动触发弹出层重新定位                            | string\|number                                     |             |          |
| returnFocusOnClose    | 按 Esc 后焦点是否回到 trigger 上（hover, focus, click 时生效）| boolean                                            | true        | 2.8.0    |
| position              | 方向（top,topLeft,topRight,left,leftTop,leftBottom,right,rightTop,rightBottom,bottom,bottomLeft,bottomRight） | string | "bottom" |          |
| spacing               | 弹出层与 children 元素的距离（px）                           | number\|object                                     | 4/10        |          |
| showArrow             | 是否显示“小三角”                                            | boolean                                            |             |          |
| stopPropagation       | 是否阻止弹出层点击事件冒泡                                  | boolean                                            | false       | 0.34.0   |
| trigger               | 触发方式（hover, focus, click, custom, contextMenu）         | string                                             | 'hover'     |          |
| visible               | 是否显示，配合trigger='custom'可实现完全受控                 | boolean                                            |             |          |
| zIndex                | 弹出层 z-index 值                                            | number                                             | 1030        |          |
| onClickOutSide        | 展示状态下，点击非Children、非浮层内部区域时回调             | function(e:event)                                  |             | 2.1.0    |
| onEscKeyDown          | 在 trigger 或弹出层按 Esc 时调用                             | function(e:event)                                  |             | 2.8.0    |
| onVisibleChange       | 弹出层展示/隐藏时触发的回调                                 | function(isVisible:boolean)                        |             |          |

### Accessibility

- trigger 为 click、custom 时，content 具有 dialog role
- trigger 为 hover 时，content 具有 tooltip role
- children 自动添加 aria-expanded、aria-haspopup、aria-controls 属性
- 支持键盘操作和焦点管理
