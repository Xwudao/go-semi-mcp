# Button 按钮

用户使用按钮来触发一个操作或者进行跳转。

## API 参考

### Button

| 属性                  | 说明                                                         | 类型                              | 默认值      |
|----------------------|------------------------------------------------------------|-----------------------------------|-------------|
| aria-label           | 按钮的标签                                                   | string                            | -           |
| block                | 块级按钮                                                     | bool                              | false       |
| className            | 类名                                                         | string                            | -           |
| contentClassName     | 内容区域 className                                           | string                            | -           |
| disabled             | 禁用状态                                                     | boolean                           | false       |
| htmlType             | button 原生 type，可选：button、reset、submit                | string                            | "button"    |
| icon                 | 图标                                                         | ReactNode                         | -           |
| iconPosition         | 图标位置，可选：left、right                                  | string                            | "left"      |
| loading              | 加载状态                                                     | boolean                           | false       |
| noHorizontalPadding  | 水平去除内边距，仅对有 icon 的 Button 有效                   | boolean \| string \| Array<string>| false       |
| size                 | 按钮大小，可选：large、default、small                        | string                            | "default"   |
| style                | 自定义样式                                                   | CSSProperties                     | -           |
| theme                | 按钮主题，可选：solid、borderless、light、outline            | string                            | "light"     |
| type                 | 类型，可选：primary、secondary、tertiary、warning、danger     | string                            | "primary"   |
| onClick              | 单击事件                                                     | function(MouseEvent)              | -           |
| onMouseDown          | 鼠标按下事件                                                 | function(MouseEvent)              | -           |
| onMouseEnter         | 鼠标移入事件                                                 | function(MouseEvent)              | -           |
| onMouseLeave         | 鼠标移出事件                                                 | function(MouseEvent)              | -           |

### ButtonGroup

| 属性        | 说明                   | 类型        | 默认值    | 版本    |
|-------------|----------------------|------------|----------|---------|
| aria-label  | 按钮组的标签           | string     | -        |         |
| className   | 自定义类名             | string     | -        |         |
| disabled    | 禁用状态               | boolean    | false    |         |
| size        | 按钮大小               | string     | "default"|         |
| style       | 自定义样式             | CSSProperties | -     | 2.20.0  |
| theme       | 按钮主题               | string     | "light"  |         |
| type        | 类型                   | string     | "primary"|         |

### SplitButtonGroup

| 属性        | 说明                   | 类型        | 默认值    |
|-------------|----------------------|------------|----------|
| aria-label  | 分裂按钮组的标签        | string     | -        |
| className   | 自定义类名             | string     | -        |
| style       | 自定义样式             | CSSProperties | -     |

---

## 其他说明

- `aria-label` 用于无文本按钮的可访问性。
- `disabled` 状态优先级高于 `loading`。
- `ButtonGroup` 可统一设置组合按钮的尺寸、禁用和类型。
- `SplitButtonGroup` 用于分裂按钮场景。

