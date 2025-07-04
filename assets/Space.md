# Space 间距

设置组件之间的间距。

## API参考

| 属性       | 说明                         | 类型                   | 默认值   | 版本       |
|------------|------------------------------|------------------------|----------|------------|
| align      | 对齐方式, 支持 start、end、center、baseline | string                | center   | >=1.17.0   |
| className  | 样式类名                     | string                 | -        | >=1.17.0   |
| spacing    | 间距尺寸, 支持 loose、medium、tight 或 number、array | string\|number\|array | tight    | >=1.17.0   |
| style      | 内联样式                     | CSSProperties          | -        | >=1.17.0   |
| vertical   | 是否为垂直间距               | boolean                | false    | >=1.17.0   |
| wrap       | 是否自动换行                 | boolean                | false    | >=1.17.0   |

## 设计变量

| 变量                   | 默认值         | 用法         |
|------------------------|---------------|--------------|
| $spacing-space-tight   | $spacing-tight| 默认间距尺寸 |
| $spacing-space-medium  | $spacing-base | 中等间距尺寸 |
| $spacing-space-loose   | $spacing-loose| 宽松间距尺寸 |
