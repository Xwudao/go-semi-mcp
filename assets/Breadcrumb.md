# Breadcrumb 面包屑

## API 参考

### Breadcrumb

| 属性           | 说明                                                                 | 类型                                              | 默认值                                                                                  | 版本      |
| -------------- | -------------------------------------------------------------------- | ------------------------------------------------- | --------------------------------------------------------------------------------------- | --------- |
| activeIndex    | 受控使用，当前选择的导航序号                                         | -                                                 | -                                                                                       | 2.61.0    |
| autoCollapse   | 是否超出maxItemCount后自动折叠                                       | boolean                                           | true                                                                                    | 1.9.0     |
| className      | 类名                                                                 | string                                            | -                                                                                       |            |
| compact        | 显示尺寸，是否紧凑                                                   | boolean                                           | true                                                                                    |            |
| maxItemCount   | 超出多少个进行自动折叠                                               | number                                            | 4                                                                                       | 1.9.0     |
| moreType       | 内置的...区域的渲染类型，可选值为 'default'、'popover'               | string                                            | 'default'                                                                               | 1.27.0    |
| renderItem     | 自定义链接函数，配合 routes 使用                                      | (Route: Route) => ReactNode                       | -                                                                                       | 0.27.0    |
| renderMore     | 自定义...区域的渲染                                                  | (restItem: ReactNode[]) => ReactNode              | -                                                                                       | 1.27.0    |
| routes         | router 的路由信息，由路由对象或字符串组成的数组，路由对象格式参考: Route | Array<Route \| string>                            | -                                                                                       |            |
| separator      | 自定义的分隔符                                                       | ReactNode                                         | '/'                                                                                     |            |
| showTooltip    | 是否展示 Tooltip 及相关配置: width，溢出宽度； ellipsisPos，截断方式，从中间/末尾截断； opts，透传给Tooltip的属性 | boolean \| showToolTipProps                       | {width: 150, ellipsisPos: 'end', opts: { autoAdjustOverflow: true, position: "bottomLeft" }} | 0.34.0    |
| style          | 内联样式                                                             | CSSProperties                                     | -                                                                                       |            |
| onClick        | 单击事件                                                             | (item: Route , e: Event) => void                  | -                                                                                       | 0.27.0    |

### Breadcrumb.Item

| 属性      | 说明                   | 类型                       | 默认值 | 版本    |
| --------- | ---------------------- | -------------------------- | ------ | ------- |
| href      | 链接的目的地           | string                     | -      |         |
| icon      | 标签的显示图标         | ReactNode                  | -      |         |
| onClick   | 单击事件               | function(item: Route, e: Event) | -      | 0.27.0 |
| separator | 分隔符，可以覆盖父级的分隔符 | ReactNode                  | -      | 1.16.0 |
| noLink    | 移除 hover 和 active 的样式 | boolean                    | false  | 1.16.0 |

### Route

| 属性  | 说明         | 类型      | 默认值 | 版本    |
| ----- | ------------ | --------- | ------ | ------- |
| href  | 链接目的地   | string    | -      | 0.27.0  |
| icon  | 标签的显示图标 | ReactNode | -      |         |
| name  | 路由名       | string    | -      |         |
| path  | 路由路径     | string    | -      |         |

> v>=1.16.0 之后 Route 支持 Breadcrumb.Item 上的相应属性。

### Accessibility

- Breadcrumb 支持传入 aria-label 来表示该 Breadcrumb 作用
- Breadcrumb 会对当前项设置 aria-current='page'

### 文案规范

- 每个页面链接都应该很简短，并且清楚地反映它链接到的位置或链接的实体
- 按句子大小写书写

### 设计变量

| 变量                                      | 默认值                                   | 用法                       |
| ----------------------------------------- | ---------------------------------------- | -------------------------- |
| $transition_duration-breadcrumb_link-text | var(--semi-transition_duration-none)     | 面包屑文字-文字-动画持续时间 |
| $transition_function-breadcrumb_link-text | var(--semi-transition_function-easeIn)   | 面包屑文字-文字-过渡曲线   |
| $transition_delay-breadcrumb_link-text    | var(--semi-transition_delay-none)        | 面包屑文字-文字-延迟时间   |
