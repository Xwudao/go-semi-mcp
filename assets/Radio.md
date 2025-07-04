# Radio 单选框

## Radio 属性

| 属性            | 说明                                                                 | 类型                        | 默认值    |
|-----------------|--------------------------------------------------------------------|-----------------------------|-----------|
| addonClassName  | 包裹内容容器的样式类名<br/>v1.16.0 后提供                         | string                      |           |
| addonId         | addon 节点 id，aria-labelledby 指向这个 id，若无设置会随机生成一个 id<br/>v2.11.0 后提供 | string                      |           |
| addonStyle      | 包裹内容容器的内联样式<br/>v1.16.0 后提供                         | CSSProperties               |           |
| aria-label      | Radio 的 label                                                     | string                      | -         |
| autoFocus       | 自动获取焦点                                                       | boolean                     | false     |
| checked         | 指定当前是否选中                                                   | boolean                     | false     |
| className       | 样式类名                                                           | string                      |           |
| defaultChecked  | 初始是否选中                                                       | boolean                     | false     |
| disabled        | 禁选单选框                                                         | boolean                     | false     |
| extra           | 副文本，只对type='default'生效                                     | ReactNode                   | -         |
| extraId         | 副文本的 id，aria-describedby 指向这个 id，若无设置会随机生成一个 id<br/>v2.11.0 后提供 | ReactNode                   | -         |
| mode            | 高级和普通模式，高级模式可以在 checked 时点击变成 unchecked，可选值 advanced | string                      | -         |
| name            | Radio组件中input[type="radio"]的name属性，具有相同name的Radio属于同一个RadioGroup | string                      | -         |
| preventScroll   | 指示浏览器是否应滚动文档以显示新聚焦的元素，作用于组件内的 focus 方法 | boolean                     |           |
| style           | 内联样式                                                           | CSSProperties               |           |
| type            | 设置 radio的样式类型，可选值为：default、button、card、pureCard<br/>v2.18.0 后提供 | string                      | default   |
| value           | 根据 value 进行比较，判断是否选中                                   | string \| number            | -         |
| onChange        | 选项变化时的回调函数                                               | function(e:Event)           | -         |
| onMouseEnter    | 鼠标移入选项时的回调函数                                           | function(e:Event)           | -         |
| onMouseLeave    | 鼠标移出选项时的回调函数                                           | function(e:Event)           | -         |

## RadioGroup 属性

| 属性         | 说明                                                                 | 类型                | 默认值    |
|--------------|--------------------------------------------------------------------|---------------------|-----------|
| aria-label   | RadioGroup 的 label                                                | string              | -         |
| buttonSize   | type='button'的radio的尺寸大小，可选值为：small、middle、large<br/>v1.26.0 后提供 | string              | middle    |
| className    | 样式类名                                                           | string              |           |
| defaultValue | 默认选中的值                                                       | string \| number    | -         |
| direction    | radio 排列方向, 只对type='default'生效，可选值horizontal、vertical<br/>v0.31.0 后提供 | string              | horizontal|
| disabled     | 禁选所有子单选器                                                   | boolean             | false     |
| mode         | 高级和普通模式，可以在 checked 时点击变成 unchecked，可选值 advanced<br/>v1.9.0 后提供 | string              | -         |
| name         | RadioGroup 下所有 input[type="radio"] 的 name 属性                 | string              | -         |
| options      | 以配置形式设置子元素                                               | Array               | -         |
| style        | 内联样式                                                           | CSSProperties       |           |
| value        | 用于设置当前选中的值                                               | string \| number    | -         |
| type         | 设置所有radio的样式类型，可选值为：default、button、card、pureCard<br/>v1.26.0 后提供，其中 card 和 pureCard 在 v1.30.0 后提供 | string              | default   |
| onChange     | 选项变化时的回调函数                                               | function(e:Event)   | -         |

## Methods

### Radio

| 名称    | 描述     |
|---------|----------|
| blur()  | 移除焦点 |
| focus() | 获取焦点 |

## Accessibility

- `aria-label`：用于解释 Radio 或 RadioGroup 的作用
- `aria-labelledby` 默认指向 addon 节点，用于解释 Radio 的内容
- `aria-describedby` 默认指向 extra 节点，用于补充解释 Radio 的内容

### 键盘和焦点

- RadioGroup 可以被获取焦点，初始焦点设置：
  - 没有被选择项时，初始焦点为第一个 Radio 项
  - 有选中项时，初始焦点为选中的 Radio 项
- 在同一个 radiogroup 内：
  - 右箭头/下箭头：焦点移动到下一个 Radio 项，并选中
  - 左箭头/上箭头：焦点移动到上一个 Radio 项，并选中
  - 没有选中项时，可用 Space 键选中初始焦点

## 设计变量

| 变量                              | 默认值                                 | 用法                   |
|-----------------------------------|----------------------------------------|------------------------|
| $transition_duration-radio-bg     | var(--semi-transition_duration-none)   | 单选框-背景色-动画持续时间 |
| $transition_function-radio-bg     | var(--semi-transition_function-easeIn) | 单选框-背景色-过渡曲线   |
| $transition_delay-radio-bg        | var(--semi-transition_delay-none)      | 单选框-背景色-延迟时间   |
| $transition_duration-radio-border | var(--semi-transition_duration-none)   | 单选框-边框-动画持续时间 |
| $transition_function-radio-border | var(--semi-transition_function-easeIn) | 单选框-边框-过渡曲线     |
| $transition_delay-radio-border    | var(--semi-transition_delay-none)      | 单选框-边框-延迟时间     |
