# Pagination 翻页器

分页器帮助用户在多个页之间进行导航。

## API 参考

| 属性                  | 说明                                                                 | 类型                                      | 默认值                | 版本      |
|-----------------------|----------------------------------------------------------------------|-------------------------------------------|-----------------------|-----------|
| className             | 类名                                                                 | string                                    |                       |           |
| currentPage           | 当前页码                                                             | number                                    |                       |           |
| defaultCurrentPage    | 默认的当前页码                                                       | number                                    |                       |           |
| disabled              | 禁用                                                                 | boolean                                   | false                 | 2.37.0    |
| hideOnSinglePage      | 总页数小于 2 时，是否自动隐藏分页器，当 showSizeChanger 为true时，此开关不再生效 | boolean                                   | false                 |           |
| hoverShowPageSelect   | hover 页码时是否展示切换页数的Select控件，仅当 size = 'small'时生效   | boolean                                   | false                 | 1.27.0    |
| nextText              | 下一页文本                                                           | string\|ReactNode                         |                       |           |
| pageSize              | 每页条数                                                             | number                                    | 10                    |           |
| pageSizeOpts          | 指定每页显示多少条                                                   | array                                     | [10, 20, 40, 100]     |           |
| popoverPosition       | 浮层方向，具体可见 Popover·API 参考·position                         | string                                    | "bottomLeft"          |           |
| popoverZIndex         | 浮层 z-index 值                                                      | number                                    | 1030                  |           |
| prevText              | 上一页文本                                                           | string\|ReactNode                         |                       |           |
| style                 | 样式                                                                 | object                                    |                       |           |
| size                  | 尺寸                                                                 | string                                    |                       |           |
| showTotal             | 是否显示总页数                                                       | boolean                                   |                       |           |
| showSizeChanger       | 是否显示切换页容量的 Select，size为small时不生效                     | boolean                                   | false                 |           |
| showQuickJumper       | 是否显示切换页码的 Input                                             | boolean                                   | false                 | 1.31.0    |
| total                 | 总条数                                                               | number                                    | 1                     |           |
| onChange              | 页码、每页容量变化时的回调函数                                       | function(currentPage: number, pageSize: number) |                |           |
| onPageChange          | 页码变化的回调函数                                                   | function(currentPage: number)              |                       |           |
| onPageSizeChange      | 每页容量变化时的回调函数                                             | function(pageSize: number)                 |                       |           |

## Accessibility

- `aria-label`: 描述组件内页码、前一页、后一页等元素的标签
- `aria-current`: 指向当前页的页码元素

## 设计变量

| 变量                                      | 默认值                                 | 用法                         |
|-------------------------------------------|----------------------------------------|------------------------------|
| $transition_duration-pagination_item-bg   | var(--semi-transition_duration-none)   | 翻页器页码-背景色-动画持续时间 |
| $transition_function-pagination_item-bg   | var(--semi-transition_function-easeIn) | 翻页器页码-背景色-过渡曲线    |
| $transition_delay-pagination_item-bg      | var(--semi-transition_delay-none)      | 翻页器页码-背景色-延迟时间    |
| $transition_duration-pagination_item-text | var(--semi-transition_duration-none)   | 翻页器页码文本-背景色-动画持续时间 |
| $transition_function-pagination_item-text | var(--semi-transition_function-easeIn) | 翻页器页码文本-背景色-过渡曲线    |
| $transition_delay-pagination_item-text    | var(--semi-transition_delay-none)      | 翻页器页码文本-背景色-延迟时间    |

## FAQ

**为什么页数下拉选择器最多只有1,000,000条？**  
因为创建列表时, 浏览器对Array.from()创建数组的大小存在限制; 同时为了兼顾Array.from()的开销，我们设定了1,000,000这个阈值。
