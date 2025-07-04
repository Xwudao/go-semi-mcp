# Input 输入框

## Input

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| aria-describedby | 设置 aria-describedby 属性 | string | - |
| aria-errormessage | 设置 aria-errormessage 属性 | string | - |
| aria-invalid | 设置 aria-invalid 属性 | string | - |
| aria-label | 设置 aria-label 属性 | string | - |
| aria-labelledby | 设置 aria-labelledby 属性 | string | - |
| aria-required | 设置 aria-required 属性 | string | - |
| addonAfter | 后置标签 | ReactNode |  |
| addonBefore | 前置标签 | ReactNode |  |
| borderless | 无边框模式 >=2.33.0 | boolean |  |
| className | 类名 | string |  |
| clearIcon | 自定义清除按钮, showClear为true时有效 >=2.25.0 | ReactNode |  |
| defaultValue | 输入框内容默认值 | ReactText |  |
| disabled | 是否禁用 | boolean | false |
| getValueLength | 自定义计算字符串长度 | (value: string) => number |  |
| hideSuffix | 清除按钮与后缀标签并存时隐藏后缀标签 | boolean | false |
| mode | 输入框的模式，可选值password >=v1.3.0 | string |  |
| prefix | 前缀标签 | ReactNode |  |
| preventScroll | 指示浏览器是否应滚动文档以显示新聚焦的元素 | boolean |  |
| showClear | 有内容且 hover 或 focus 时展示清除按钮 >=1.0.0 | boolean | false |
| size | 输入框大小，large、default、small | string | 'default' |
| style | 样式 | CSSProperties |  |
| suffix | 后缀标签 | ReactNode |  |
| type | 声明input类型，同原生input标签的type属性 | string | text |
| validateStatus | 校验状态，可选default、error、warning | string | 'default' |
| value | 输入框内容 | ReactText |  |
| onBlur | 失去焦点回调 | function(e:event) |  |
| onChange | 内容变化回调 | function(value:string, e:event) |  |
| onClear | 点击清除按钮回调 | function(e:event) |  |
| onEnterPress | 按回车时回调 | function(e:event) |  |
| onFocus | focus时回调 | function(e:event) |  |
| onKeyDown | keydown回调 | function(e:event) |  |
| onKeyPress | keypress回调 | function(e:event) |  |
| onKeyUp | keyup回调 | function(e:event) |  |

---

## TextArea

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| aria-describedby | 设置 aria-describedby 属性 | string | - |
| aria-errormessage | 设置 aria-errormessage 属性 | string | - |
| aria-invalid | 设置 aria-invalid 属性 | string | - |
| aria-label | 设置 aria-label 属性 | string | - |
| aria-labelledby | 设置 aria-labelledby 属性 | string | - |
| aria-required | 设置 aria-required 属性 | string | - |
| autosize | 自动适应内容高度，可为对象配置最小最大行数 | boolean \| object | false |
| borderless | 无边框模式 >=2.33.0 | boolean |  |
| className | 类名 | string | - |
| cols | 默认列数 | number | 无 |
| disabled | 禁用状态 | boolean | false |
| getValueLength | 自定义计算字符串长度 | (value: string) => number |  |
| maxCount | 字数限制并显示统计 | number | 无 |
| placeholder | 默认值 | string | 无 |
| readonly | 只读 | boolean | false |
| rows | 默认行数 | number | 4 |
| showClear | 支持清除 >=1.30.0 | boolean | false |
| style | 样式 | CSSProperties | - |
| onBlur | 失去焦点回调 | (e:event) => void | - |
| onChange | 内容变化回调 | (value:string, e:event) => void |  |
| onClear | 点击清除按钮回调 >=1.30.0 | (e:event) => void | - |
| onEnterPress | 按下回车回调 | (e:event) => void | 无 |
| onFocus | focus时回调 | (e:event) => void | - |
| onKeyDown | keydown回调 | (e:event) => void | - |
| onKeyPress | keypress回调 | (e:event) => void | - |
| onKeyUp | keyup回调 | (e:event) => void | - |
| onResize | 高度变化回调 >=v0.37.0 | ({ height:number }) => void | - |

---

## InputGroup

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| className | 组的类名 | string | - |
| disabled | 禁用 | boolean | - |
| label | InputGroup 的 label 属性 | LabelProps | - |
| labelPosition | label 位置，可选 top 或 left | string | - |
| size | 输入框大小，large、default、small | string | 'default' |
| style | 组的样式 | CSSProperties | - |
| onBlur | 失去焦点回调 | (e:event) => void | - |
| onFocus | focus时回调 | (e:event) => void | - |

---

## Methods

- `blur()`：移出焦点
- `focus()`：获取焦点

---

## Accessibility

- 当 validateStatus 为 error 时，输入框的 aria-invalid 为 true
- 在 Form 中使用时，field label 是 Input 的 aria-label
- Input 可被获取焦点，Tab/Shift+Tab 切换焦点
- 密码按钮可聚焦，Enter/空格激活

---

## 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ---- | ---- |
| $transition_duration-input-bg | var(--semi-transition_duration-none) | 输入框-背景色-动画持续时间 |
| $transition_function-input-bg | var(--semi-transition_function-easeIn) | 输入框-背景色-过渡曲线 |
| $transition_delay-input-bg | var(--semi-transition_delay-none) | 输入框-背景色-延迟时间 |
| $transition_duration-input-border | var(--semi-transition_duration-none) | 输入框-边框-动画持续时间 |
| $transition_function-input-border | var(--semi-transition_function-easeIn) | 输入框-边框-过渡曲线 |
| $transition_delay-input-border | var(--semi-transition_delay-none) | 输入框-边框-延迟时间 |
| $transition_duration-input-text | var(--semi-transition_duration-none) | 输入框-文字或图标-动画持续时间 |
| $transition_function-input-text | var(--semi-transition_function-easeIn) | 输入框-文字或图标-过渡曲线 |
| $transition_delay-input-text | var(--semi-transition_delay-none) | 输入框-文字或图标-延迟时间 |
