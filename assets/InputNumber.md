# InputNumber 数字输入框

通过鼠标或键盘，输入范围内的数值，与 Input 不同的是它带有针对数字场景的步进器操作区，配合 Parser 使用可以展示更复杂的内容格式。

## API 参考

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| ---- | ---- | ---- | ------ | ---- |
| autofocus | 自动获取焦点 | boolean | false |  |
| className | 类名 | string | - |  |
| clearIcon | 可用于自定义清除按钮, showClear为true时有效 | ReactNode |  | 2.25.0 |
| currency | 货币种类，国际化模式下通过 currency={true} 开启，组件会自动根据 locale 展示对应货币种类, 也可以手动传入 localeCode 和 currency 指定展示的货币种类, currency 的可选值有 CNY,EUR,USD等 | boolean\|string | false | 2.77.0 |
| currencyDisplay | 货币展示方式，可选值：symbol、code、name | string | symbol | 2.77.0 |
| defaultValue | 默认值 | number |  |  |
| disabled | 禁用 | boolean | false |  |
| formatter | 指定输入框展示值的格式 | (value: number\|string) => string | - |  |
| hideButtons | 为 true 时隐藏 “上/下” 按钮 | boolean | false | 1.0.0 |
| innerButtons | 为 true 时 “上/下” 按钮显示在输入框内部 | boolean | false | 1.5.0 |
| keepFocus | 点击按钮时保持输入框聚焦 | boolean | false | 1.10.0 |
| localeCode | 货币模式下用于指定国家地区代码，可选值有 zh-CN, en-US, en-GB, ja-JP, ko-KR, ar, vi-VN, ru-RU, id-ID, ms-MY, th-TH, tr-TR, pt-BR, zh-TW, es, de, it, fr, ro, sv-SE, pl-PL, nl-NL等 | string | - | 2.77.0 |
| max | 限定最大值 | number | Infinity |  |
| min | 限定最小值 | number | -Infinity |  |
| parser | 指定从 formatter 里转换回数字串的方式，和 formatter 搭配使用 | (str: string) => string | - |  |
| precision | 数值精度 | number | - |  |
| prefix | 前缀内容 | string\|ReactNode |  |  |
| pressInterval | 长按按钮时，多久触发一次点击事件，单位毫秒 | number | 250 |  |
| pressTimeout | 长按按钮时，延迟多久后触发点击事件，单位毫秒 | number | 250 |  |
| preventScroll | 指示浏览器是否应滚动文档以显示新聚焦的元素，作用于组件内的 focus 方法 | boolean |  |  |
| shiftStep | 按住 shift 键每次改变步数，可以为小数，v2.13 默认值由 1 调整为 10 | number | 10 | 1.5.0 |
| showClear | 是否显示清除按钮 | boolean | false | 0.35.0 |
| showCurrencySymbol | 是否显示货币符号/代码/名称，仅货币模式下生效 | boolean | true | 2.77.0 |
| size | 输入框大小，可选值："default"\|"small"\|"large" | string | 'default' |  |
| step | 每次改变步数，可以为小数 | number | 1 |  |
| style | 样式 | CSSProperties | - |  |
| suffix | 自定义后缀 | ReactNode |  |  |
| value | 当前值 | number |  |  |
| onBlur | 失去焦点时的回调 | (e: domEvent) => void | () => {} | 1.0.0 |
| onChange | 变化回调 | (value: number\|string) => void | - |  |
| onFocus | 获得焦点时的回调 | (e: domEvent) => void | () => {} | 1.0.0 |
| onNumberChange | 数字变化回调 | (value: number) => void | - | 1.9.0 |

## Methods

| 名称 | 描述 |
| ---- | ---- |
| blur() | 移出焦点 |
| focus() | 获取焦点 |

## Accessibility

- 数字输入框具有 spinbutton role
- spinbutton 使用 aria-valuenow 表示当前值，aria-valuemax 表示可以接受的最大值，aria-valuemin 表示可以接受的最小值
- 当 InputNumber 在 Form 中使用时，输入框的 aria-labeledby 指向 Field label
- 可被获取焦点，键盘用户可以使用 Tab 及 Shift + Tab 切换焦点（增加/减少按钮不可以被键盘聚焦）
- 键盘用户可以按上键 ⬆️ 或下键 ⬇️ ，输入值将增加或减少 step（默认值为 1）
- 按住 Shift + 上键 ⬆️ 或下键 ⬇️ ，输入值将增加或减少 shiftStep（默认值为 10）

## 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ------ | ---- |
| $transition_duration-inputNumber-bg | var(--semi-transition_duration-none) | 数字输入框-背景色-动画持续时间 |
| $transition_function-inputNumber-bg | var(--semi-transition_function-easeIn) | 数字输入框-背景色-过渡曲线 |
| $transition_delay-inputNumber-bg | var(--semi-transition_delay-none) | 数字输入框-背景色-延迟时间 |
| $transition_duration-inputNumber-border | var(--semi-transition_duration-none) | 数字输入框-边框-动画持续时间 |
| $transition_function-inputNumber-border | var(--semi-transition_function-easeIn) | 数字输入框-边框-过渡曲线 |
| $transition_delay-inputNumber-border | var(--semi-transition_delay-none) | 数字输入框-边框-延迟时间 |
