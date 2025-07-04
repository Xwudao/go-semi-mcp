# Descriptions 描述列表

## API 参考

### Descriptions

| 属性      | 说明                                               | 类型              | 默认值    |
|-----------|----------------------------------------------------|-------------------|-----------|
| align     | 描述列表的对齐方式，可选 center、justify、left、plain | string            | center    |
| className | 类名                                               | string            | -         |
| data      | 列表显示的内容                                     | DataItem[]        | -         |
| row       | 是否双行显示                                       | boolean           | false     |
| size      | 设置双行显示时的列表的大小，small、medium、large    | string            | medium    |
| style     | 列表的样式                                         | CSSProperties     | -         |
| layout    | 列表布局模式，vertical、horizontal (v>=2.54.0)      | string            | vertical  |
| column    | 横向布局下每行的总列数 (v>=2.54.0)                  | number            | 3         |

#### DataItem

| 属性   | 说明                         | 类型                        | 默认值 |
|--------|------------------------------|-----------------------------|--------|
| key    | 键值                         | ReactNode                   | -      |
| value  | 属性值                       | ReactNode \| () => ReactNode| -      |
| hidden | 该数据是否需要展示 (v>=1.12.0)| boolean                     | -      |
| span   | 单元格应跨越的列数 (v>=2.54.0)| number                      | 1      |

#### DescriptionItem

| 属性      | 说明                         | 类型          | 默认值 |
|-----------|------------------------------|---------------|--------|
| itemKey   | 键值                         | ReactNode     | -      |
| hidden    | 该数据是否需要展示           | boolean       | -      |
| className | Item 外部 wrapper: tr 的类名 | string        | -      |
| style     | Item 外部 wrapper: tr 的样式 | CSSProperties | -      |
| span      | 单元格应跨越的列数 (v>=2.54.0)| number        | 1      |

### 设计变量

| 变量                                   | 默认值                        | 用法         |
|----------------------------------------|-------------------------------|--------------|
| $color-descriptions_key-text-default   | var(--semi-color-text-2)      | key 文字颜色 |
| $color-descriptions_value-text-default | var(--semi-color-text-0)      | value 文字颜色 |

