# Rating 评分

展示评分的组件。

## API参考

| 属性           | 说明                                                                 | 类型                        | 默认值                                  |
|----------------|--------------------------------------------------------------------|-----------------------------|-----------------------------------------|
| allowClear     | 是否允许再次点击后清除                                              | boolean                     | true                                    |
| allowHalf      | 是否允许半选                                                        | boolean                     | false                                   |
| autoFocus      | 自动获取焦点                                                        | boolean                     | false                                   |
| character      | 自定义字符                                                          | ReactNode                   | `<IconStar size="extra-large"/>`        |
| className      | 自定义样式类名                                                      | string                      | -                                       |
| count          | star 总数                                                           | number                      | 5                                       |
| defaultValue   | 默认值                                                              | number                      | 0                                       |
| disabled       | 只读，无法进行交互                                                  | boolean                     | false                                   |
| preventScroll  | 指示浏览器是否应滚动文档以显示新聚焦的元素，作用于组件内的 focus 方法 | boolean                     | -                                       |
| size           | 尺寸，default，small，支持 number 类型自定义尺寸                    | string \| number            | default                                 |
| style          | 自定义样式对象                                                      | object                      | -                                       |
| tooltips       | 自定义每项的提示信息                                                | string[]                    | -                                       |
| value          | 当前受控值                                                          | number                      | -                                       |
| onBlur         | 失去焦点时的回调                                                    | function()                  | -                                       |
| onChange       | 选择时的回调                                                        | function(value: number)     | -                                       |
| onFocus        | 获取焦点时的回调                                                    | function()                  | -                                       |
| onHoverChange  | 鼠标经过时数值变化的回调                                            | function(value: number)     | -                                       |
| onKeyDown      | 按键回调                                                            | function(e: event)          | -                                       |

## Accessibility

- 支持 `aria-checked`、`aria-posinset`、`aria-setsize`。
- 可用 `aria-label` 定制语义化。
- 键盘支持：左右/上下箭头切换，支持半星，disabled 时不可聚焦。

## 设计变量

| 变量                                 | 默认值                                 | 用法                       |
|--------------------------------------|----------------------------------------|----------------------------|
| $transition_duration-rating-color    | var(--semi-transition_duration-none)   | 评分-背景色-动画持续时间    |
| $transition_function-rating-color    | var(--semi-transition_function-easeIn) | 评分-背景色-过渡曲线        |
| $transition_delay-rating-color       | var(--semi-transition_delay-none)      | 评分-背景色-延迟时间        |
