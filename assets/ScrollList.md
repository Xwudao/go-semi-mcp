# ScrollList 滚动列表

## ScrollList 属性

| 属性        | 说明         | 类型             | 默认值  |
| ----------- | ------------ | ---------------- | ------- |
| bodyHeight  | body高度     | string \| number |         |
| className   | 样式类名     | string           | ''      |
| footer      | 底部 addon   | ReactNode        | ''      |
| header      | 头部 addon   | ReactNode        | ''      |
| style       | 内联样式     | CSSProperties    | {}      |

## ScrollItem 属性

| 属性          | 说明                                         | 类型                                         | 默认值  |
| ------------- | -------------------------------------------- | -------------------------------------------- | ------- |
| className     | 样式类名                                     | string                                       | ''      |
| cycled        | 是否为无限循环，仅在 mode 为 "wheel" 时生效  | boolean                                      | false   |
| list          | 列表内容                                     | ItemData[]                                   | []      |
| mode          | 模式选择                                     | "normal" \| "wheel"                          | "wheel" |
| motion        | 是否开启滚动动画                             | Motion                                       | true    |
| onSelect      | 选中回调                                     | (data: ItemData) => void                     | NOOP    |
| selectedIndex | 选中项的索引                                 | number                                       | 0       |
| style         | 内联样式                                     | CSSProperties                                | {}      |
| transform     | 对选中项的变换，返回值会作为文案进行显示      | (value: any, text: string) => string         | v => v  |

## ItemData

| 属性      | 说明                                                         | 类型                                 | 默认值  |
| --------- | ------------------------------------------------------------ | ------------------------------------ | ------- |
| disabled  | 该项是否被禁止选择                                           | boolean                              |         |
| text      | 每一项的文案                                                 | string                               |         |
| transform | 该项处于选中状态时的变换，返回值会作为文案进行显示           | (value: any, text: string) => string | v => v  |
| value     | 每一项的值                                                   | any                                  |         |

## Accessibility

- ScrollItem 支持传入 aria-label, 指定该列标签
- ScrollItem 使用 aria-disabled 表示该项目是否被禁用
- ScrollItem 使用 aria-selected 表示该项目是否被选中

## 设计变量

| 变量                                         | 默认值                                 | 用法                                 |
| -------------------------------------------- | -------------------------------------- | ------------------------------------ |
| $transition_duration-scrollList_selected_item-bg | var(--semi-transition_duration-none)   | 滚动列表选中选项-背景颜色-动画持续时间 |
| $transition_function-scrollList_selected_item-bg | var(--semi-transition_function-easeOut)| 滚动列表选中选项-背景颜色-过渡曲线     |
| $transition_delay-scrollList_selected_item-bg    | 0ms                                   | 滚动列表选中选项-背景颜色-延迟时间     |
