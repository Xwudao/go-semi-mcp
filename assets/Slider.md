# Slider 滑动选择器

滑动选择器，使用拖动交互快速选择数值或数值范围，与 InputNumber 相比更直观。

## API 参考

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| ---- | ---- | ---- | ------ | ---- |
| aria-label | aria-label属性，用来给当前元素加上的标签描述, 提升可访问性 | string | - | - |
| aria-labelledby | aria-labelledby属性，表明某些元素的 id 是某一对象的标签。它被用来确定控件或控件组与它们标签之间的联系, 提升可访问性 | string | - | - |
| aria-valuetext | aria-valuetext属性，为滑块的当前值提供用户友好的名称。 | string | - | - |
| defaultValue | 设置初始取值 | number \| number[] | 0 | - |
| disabled | 滑块是否禁用 | boolean | false | - |
| handleDot | 滑块是否带有圆点 | { color: string, size: string} \| { color: string, size: string}[] | - | 2.52.0 |
| included | marks 不为空对象时有效，值为 true 时表示值为包含关系，false 表示并列 | boolean | true | - |
| marks | 刻度，key 的类型必须为 number 且取值在闭区间 [min, max] 内 | Record<number, string > | 无 | - |
| max | 最大值 | number | 100 | - |
| min | 最小值 | number | 0 | - |
| railStyle | 滑块轨道的样式 | CSSProperties | - | 0.31.0 |
| range | 是否支持两边同时可滑动 | boolean | false | - |
| showArrow | tooltip 是否带箭头 | boolean | true | 2.48.0 |
| showBoundary | 是否在 hover 时展示最大值最小值 | boolean | false | - |
| showMarkLabel | 是否隐藏标签 | boolean | true | 2.48.0 |
| step | 步长 | number | 1 | - |
| tipFormatter | 设置Tooltip的展示格式，默认显示当前选值 | (value: string \| number \| boolean \| (string \| number \| boolean)[]) => any | v => v | - |
| tooltipOnMark | 滑轨上的 mark 是否带有 tooltip | boolean | false | 2.48.0 |
| tooltipVisible | 是否始终显示Tooltip | boolean | 无 | - |
| value | 设置当前取值 | number \| number[] |  | - |
| vertical | 是否设置方向为垂直 | boolean | false | - |
| verticalReverse | 反转垂直方向，即上大下小 | boolean | false | >=1.29.0 |
| onAfterChange | 值变化后触发，把当前值作为参数传入 | (value: number \| number[]) => void | 无 | - |
| onChange | 当 Slider 的值发生改变时的回调 | (value: number \| number[]) => void | 无 | - |
| onMouseUp | 鼠标松开滑块时触发 | (e: React.MouseEvent<HTMLDivElement>) => void | 无 | 2.41.0 |
| getAriaValueText | 用于给滑块的当前值提供一个用户友好的名称，对屏幕阅读器用户很重要，参数value为当前滑块的值，index为当前滑块的顺序 | (value: number, index?: number) => string | - | - |

## Accessibility

- 可聚焦的控制元素 role 为 slider。
- aria-valuenow/aria-valuemin/aria-valuemax/aria-orientation/aria-valuetext 支持。
- 支持 aria-label 或 aria-labelledby。
- 键盘支持：Tab、Shift+Tab、上下左右箭头、Page Up/Down、Home/End。

## 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ------ | ---- |
| $transition_duration-slider_handle-bg | var(--semi-transition_duration-none) | 滑动条圆形按钮-背景色-动画持续时间 |
| $transition_function-slider_handle-bg | var(--semi-transition_function-easeIn) | 滑动条圆形按钮-背景色-过渡曲线 |
| $transition_delay-slider_handle-bg | var(--semi-transition_delay-none) | 滑动条圆形按钮-背景色-延迟时间 |
