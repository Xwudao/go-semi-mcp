# Banner 通知横幅

## API 参考

| 属性       | 说明                                         | 类型                | 默认值   | 版本   |
|------------|----------------------------------------------|---------------------|----------|--------|
| bordered   | 是否展示边框，仅在非全屏模式下有效           | boolean             | false    | 1.0    |
| className  | 类名                                         | string              | -        | -      |
| closeIcon  | 自定义关闭icon，为 null 时不显示关闭按钮     | string\|ReactNode   | -        | 1.0    |
| description| 描述内容                                     | ReactNode           | -        | 1.0    |
| fullMode   | 是否为全屏模式                               | boolean             | true     | 1.0    |
| icon       | 自定义 icon，为 null 时不显示 icon           | string\|ReactNode   | -        | 1.0    |
| onClose    | 关闭时的回调函数                             | function            | -        | -      |
| style      | 样式名                                       | object              | -        | -      |
| title      | 标题                                         | ReactNode           | -        | 1.0    |
| type       | 类型，支持 info, success, danger, warning    | string              | info     | -      |
