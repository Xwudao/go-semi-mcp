# Checkbox 复选框

## Checkbox API

| 参数           | 说明                                                                 | 类型        | 默认值   |
| -------------- | -------------------------------------------------------------------- | ----------- | -------- |
| addonId        | addon 节点 id，aria-labelledby 指向这个 id，若无设置会随机生成一个 id | string      | -        |
| aria-label     | 定义 Checkbox 的作用                                                 | string      | -        |
| checked        | 指定当前Checkbox是否选中（在Group中使用时无效）                      | boolean     | false    |
| type           | 设置checkbox 的样式类型，可选: default、card、pureCard               | string      | default  |
| defaultChecked | 初始是否选中（在Group中使用时无效）                                  | boolean     | false    |
| disabled       | 失效状态                                                             | boolean     | false    |
| extra          | 副文本                                                               | ReactNode   | -        |
| extraId        | 副文本的 id，aria-describedby 指向这个 id                            | ReactNode   | -        |
| value          | 该checkbox在CheckboxGroup中代表的value                               | any         | -        |
| indeterminate  | 设置 indeterminate 状态，只负责样式控制                              | boolean     | false    |
| preventScroll  | 指示浏览器是否应滚动文档以显示新聚焦的元素                           | boolean     | -        |
| onChange       | 变化时回调函数                                                       | function(e) | -        |

### 方法

| 名称    | 描述         |
| ------- | ------------ |
| blur()  | 移除焦点     |
| focus() | 获取焦点     |

---

## CheckboxGroup API

| 参数        | 说明                                               | 类型      | 默认值   |
| ----------- | -------------------------------------------------- | --------- | -------- |
| defaultValue| 组内默认选中的选项，会与Checkbox的value值做匹配    | any[]     | []       |
| direction   | 组内checkbox布局，可选 horizontal 或 vertical      | string    | vertical |
| disabled    | 整组失效                                           | boolean   | false    |
| name        | CheckboxGroup 下所有 input 的 name 属性            | string    | -        |
| options     | 指定可选项                                         | any[]     | []       |
| type        | 设置所有 checkbox 的样式类型                       | string    | default  |
| value       | 指定选中的选项                                     | any[]     | []       |
| onChange    | 变化时回调函数                                     | function  | -        |

---

## Accessibility

- Checkbox 的 role 为 `checkbox`，CheckboxGroup 的 role 为 `list`，其子元素为 `listitem`
- `aria-label`：无文本时建议传入，屏幕阅读器可读
- `aria-labelledby` 指向 addon 节点
- `aria-describedby` 指向 extra 节点
- `aria-disabled` 与 `disabled` prop 保持一致
- `aria-checked` 表示当前选中状态

---

## 设计变量

| 变量                                 | 默认值                                 | 用法                   |
| ------------------------------------ | -------------------------------------- | ---------------------- |
| $transition_duration-checkbox-bg     | var(--semi-transition_duration-none)   | 背景色动画持续时间     |
| $transition_function-checkbox-bg     | var(--semi-transition_function-easeIn) | 背景色过渡曲线         |
| $transition_delay-checkbox-bg        | var(--semi-transition_delay-none)      | 背景色延迟时间         |
| $transition_duration-checkbox-border | var(--semi-transition_duration-none)   | 边框动画持续时间       |
| $transition_function-checkbox-border | var(--semi-transition_function-easeIn) | 边框过渡曲线           |
| $transition_delay-checkbox-border    | var(--semi-transition_delay-none)      | 边框延迟时间           |
