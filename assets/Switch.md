# Switch 开关

开关是用于切换两种互斥状态的交互形式。

## API 参考

| 属性             | 说明                                                                 | 类型                         | 默认值    | 版本     |
|------------------|----------------------------------------------------------------------|------------------------------|-----------|----------|
| aria-label       | aria-label属性，用来给当前元素加上的标签描述, 提升可访问性           | string                       |           | 2.2.0    |
| aria-labelledby  | aria-labelledby属性，表明某些元素的 id 是某一对象的标签。提升可访问性 | string                       |           | 2.2.0    |
| className        | 类名                                                                 | string                       |           |          |
| checked          | 指示当前是否选中,配合 onChange 使用                                  | boolean                      | false     |          |
| checkedText      | 打开时展示的内容, size 为 small 时无效                               | ReactNode                    |           |          |
| defaultChecked   | 初始是否选中                                                        | boolean                      | false     |          |
| disabled         | 是否禁用                                                            | boolean                      | false     |          |
| loading          | 设置加载状态                                                        | boolean                      | false     | 1.29.0   |
| onChange         | 变化时回调函数                                                      | function(checked:boolean)    |           |          |
| onMouseEnter     | 鼠标移入时回调                                                      | function()                   |           |          |
| onMouseLeave     | 鼠标移出时回调                                                      | function()                   |           |          |
| size             | 尺寸,可选值 large,default,small                                     | string                       | 'default' |          |
| style            | 内联样式                                                            | object                       |           |          |
| uncheckedText    | 关闭时展示的内容, size 为 small 时无效                              | ReactNode                    |           |          |

## Accessibility

- Switch 具有 switch role，当 checked 为 true 时，aria-checked 将被自动设置为 true，反之亦然。
- 作为表单控件应该带有 Label，当你使用 Form.Switch 时会自动被带上。
- 如果你单独使用 Switch，建议使用 aria-label 描述当前标签作用。

## 键盘和焦点

- 键盘用户可以使用 Tab 及 Shift + Tab 切换焦点。
- 聚焦时可以通过 Space 键切换开启或关闭状态。

## 设计变量

| 变量                           | 默认值                                 | 用法                       |
|--------------------------------|----------------------------------------|----------------------------|
| $transition_duration-switch-bg | 200ms                                  | 开关-背景色-动画持续时间    |
| $transition_function-switch-bg | var(--semi-transition_function-easeIn) | 开关-背景色-过渡曲线        |
| $transition_delay-switch-bg    | var(--semi-transition_delay-none)      | 开关-背景色-延迟时间        |
