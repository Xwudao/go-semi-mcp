# Popconfirm 气泡确认框

## API 参考

| 属性                | 说明                                                                 | 类型                                 | 默认值         | 版本      |
|---------------------|----------------------------------------------------------------------|--------------------------------------|----------------|-----------|
| arrowPointAtCenter  | “小三角”是否指向元素中心，需 showArrow=true                          | boolean                              | false          |           |
| cancelText          | 取消按钮文字                                                         | string                               | "取消"         |           |
| cancelButtonProps   | 取消按钮的 props                                                     | object                               |                | 0.29.0    |
| cancelType          | 取消按钮类型                                                         | string                               | "tertiary"     |           |
| closeOnEsc          | 按 Esc 是否关闭面板，受控时不生效                                    | boolean                              | true           | 2.8.0     |
| content             | 显示的内容（函数类型 2.10.0 支持）                                   | ReactNode \| ({ initialFocusRef })=>ReactNode |         |           |
| defaultVisible      | 气泡框默认是否展示                                                   | boolean                              |                |           |
| disabled            | 点击子元素是否弹出气泡确认框                                         | boolean                              | false          |           |
| getPopupContainer   | 指定父级 DOM，弹层渲染至该 DOM，自定义需设置 position: relative      | Function():HTMLElement               | document.body  |           |
| guardFocus          | 焦点在弹出层内时，Tab 是否循环                                       | boolean                              | true           | 2.8.0     |
| icon                | 自定义弹出气泡 Icon 图标                                             | ReactNode                            |                |           |
| motion              | 下拉列表出现/隐藏时，是否有动画                                      | boolean                              | true           |           |
| okText              | 确认按钮文字                                                         | string                               | "确认"         |           |
| okType              | 确认按钮类型                                                         | string                               | "primary"      |           |
| okButtonProps       | 确认按钮的 props                                                     | object                               |                |           |
| position            | 方向，可选 top,topLeft,topRight,left,leftTop,leftBottom,right,rightTop,rightBottom,bottom,bottomLeft,bottomRight | string | "bottomLeft" |           |
| returnFocusOnClose  | 按 Esc 后焦点是否回到 trigger 上，仅 trigger 为 click 时生效         | boolean                              | true           | 2.8.0     |
| showArrow           | 是否显示箭头三角形                                                   | boolean                              | false          |           |
| stopPropagation     | 是否阻止弹层点击事件冒泡                                             | boolean                              | true           |           |
| title               | 显示的标题                                                           | ReactNode                            |                |           |
| trigger             | 触发展示的时机，可选 hover / focus / click / custom                  | string                               | 'click'        |           |
| visible             | 气泡框是否展示的受控属性                                             | boolean                              |                |           |
| zIndex              | 浮层 z-index 值                                                      | number                               | 1030           |           |
| onConfirm           | 点击确认按钮回调                                                     | Function(e)                          |                |           |
| onCancel            | 点击取消按钮回调                                                     | Function(e)                          |                |           |
| onClickOutSide      | 展示状态下，点击非 Children、非浮层内部区域时回调                    | Function(e)                          |                | 2.1.0     |
| onEscKeyDown        | 在 trigger 或弹出层按 Esc 键时调用                                    | Function(e:event)                    |                | 2.8.0     |
| onVisibleChange     | 气泡框切换显示隐藏的回调                                             | Function(visible: boolean): void     | () => {}       |           |

### Accessibility

- 语义化请参考 Popover
- 必须带有触发器，触发器可被聚焦，使用 Enter 键打开
- 激活后，方向键 ⬇️ 将焦点移动到 Popconfirm 上
- 初始焦点建议设置在最常用或最安全的交互元素上（okButtonProps/cancelButtonProps 传 autoFocus）
- 支持 Esc 关闭并返回焦点到触发器

### 设计变量

| 变量                             | 默认值                        | 用法           |
|----------------------------------|-------------------------------|----------------|
| $color-popconfirm_header-text    | var(--semi-color-text-0)      | 标题文字颜色   |
| $color-popconfirm_body-text      | var(--semi-color-text-2)      | 正文字颜色     |
| $color-popconfirm_header_alert-icon | var(--semi-color-warning)   | 警示图标颜色   |
