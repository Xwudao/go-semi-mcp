# Highlight 高亮文本

## 组件说明
用于高亮显示文本中的指定内容。

## API

| 属性                  | 说明                                         | 类型                | 默认值   |
|----------------------|--------------------------------------------|---------------------|----------|
| searchWords          | 期望高亮显示的文本（对象数组在v2.71后支持）      | string[] \| object[]| []       |
| sourceString         | 源文本                                       | string              | -        |
| component            | 高亮标签                                     | string              | mark     |
| highlightClassName   | 高亮标签的样式类名                             | string              | -        |
| highlightStyle       | 高亮标签的内联样式                             | CSSProperties       | -        |
| caseSensitive        | 是否大小写敏感                                 | boolean             | false    |
| autoEscape           | 是否自动转义                                   | boolean             | true     |

## 说明
- `searchWords` 支持字符串数组和对象数组（对象数组需 v2.71.0 及以上）。
- 可通过 `component` 指定高亮包裹标签，默认为 `<mark>`。
- 支持自定义高亮样式和类名。
