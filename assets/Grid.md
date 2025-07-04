# Grid 栅格

## Row

| 属性      | 说明                                                         | 类型                                   | 默认值   |
|-----------|--------------------------------------------------------------|----------------------------------------|----------|
| align     | flex 布局下的垂直对齐方式：top middle bottom                 | string                                 |          |
| className | 类名                                                         | string                                 |          |
| gutter    | 栅格间隔，可以写成像素值或支持响应式的对象写法 { xs: 8, sm: 16, md: 24}<br>从1.11.0版本起支持垂直间隔 | number \| object \| Array<number\|object> |          |
| justify   | flex 布局下的水平排列方式：start end center space-around space-between | string                                 | start    |
| style     | 自定义样式                                                   | CSSProperties                          |          |
| type      | 布局模式，可选 flex，现代浏览器下有效                        | string                                 |          |

## Col

| 属性    | 说明                                         | 类型         | 默认值 |
|---------|----------------------------------------------|--------------|--------|
| lg      | ≥992px 响应式栅格，可为栅格数或对象配置      | number\|object | -      |
| md      | ≥768px 响应式栅格，可为栅格数或对象配置      | number\|object | -      |
| offset  | 栅格左侧的间隔格数，间隔内不可以有栅格       | number       | 0      |
| order   | 栅格顺序，flex 布局模式下有效                | number       | 0      |
| pull    | 栅格向左移动格数                             | number       | 0      |
| push    | 栅格向右移动格数                             | number       | 0      |
| sm      | ≥575px 响应式栅格，可为栅格数或对象配置      | number\|object | -      |
| span    | 栅格占位格数，为 0 时相当于 display: none    | number       | -      |
| xl      | ≥1200px 响应式栅格，可为栅格数或对象配置     | number\|object | -      |
| xs      | <576px 响应式栅格，可为栅格数或对象配置      | number\|object | -      |
| xxl     | ≥1600px 响应式栅格，可为栅格数或对象配置     | number\|object | -      |

## 设计变量

| 变量                        | 默认值  | 用法               |
|-----------------------------|---------|--------------------|
| $width-grid_screen-xs       | 480px   | 超小尺寸设备 - 手机 |
| $width-grid_screen-xs-min   | $width-grid_screen-xs | 暂无 |
| $width-grid_screen-sm       | 576px   | 小尺寸设备 - 平板   |
| $width-grid_screen-sm-min   | $width-grid_screen-sm | 暂无 |
| $width-grid_screen-md       | 768px   | 中尺寸设备 - 水平平板 |
| $width-grid_screen-md-min   | $width-grid_screen-md | 暂无 |
| $width-grid_screen-lg       | 992px   | 大尺寸设备 - 小尺寸桌面端 |
| $width-grid_screen-lg-min   | $width-grid_screen-lg | 暂无 |
| $width-grid_screen-xl       | 1200px  | 超大尺寸设备 - 桌面端 |
| $width-grid_screen-xl-min   | $width-grid_screen-xl | 暂无 |
