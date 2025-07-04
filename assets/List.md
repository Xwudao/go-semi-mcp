# List 列表

基础列表组件，支持多种布局、栅格、分页、加载等功能。

## API 参考

### List

| 属性           | 说明                         | 类型           | 默认值   |
| -------------- | ---------------------------- | -------------- | -------- |
| bordered       | 是否显示边框                 | boolean        | false    |
| className      | 自定义样式类名               | string         | -        |
| dataSource     | 列表数据源                   | any[]          | -        |
| emptyContent   | 空列表的展示内容             | ReactNode      | -        |
| footer         | 列表底部                     | ReactNode      | -        |
| grid           | 列表栅格配置                 | Grid           | -        |
| header         | 列表头部                     | ReactNode      | -        |
| layout         | 列表布局 vertical/horizontal | string         | vertical |
| loadMore       | 加载更多的按钮               | ReactNode      | -        |
| loading        | 是否处于加载中               | boolean        | false    |
| renderItem     | 自定义渲染列表项             | (item, ind) => ReactNode | - |
| size           | 列表尺寸 small/default/large | string         | default  |
| split          | 是否展示分割线               | boolean        | true     |
| style          | 自定义样式对象               | CSSProperties  | -        |
| onClick        | 点击回调事件                 | (e: event) => void | -    |
| onRightClick   | 右键点击回调事件             | (e: event) => void | -    |

#### List grid props

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| span | 栅格占位格数 | number | - |
| gutter | 栅格间隔 | number | 0 |
| xs/sm/md/lg/xl/xxl | 响应式栅格 | number/object | - |

### List.Item

| 属性      | 说明                                   | 类型      | 默认值     |
| --------- | -------------------------------------- | --------- | ---------- |
| align     | 头内容和主体内容的垂直对齐方式         | string    | flex-start |
| className | 自定义样式类名                         | string    | -          |
| extra     | 列表项附加内容                         | ReactNode | -          |
| header    | 列表项头内容                           | ReactNode | -          |
| main      | 列表项主体内容                         | ReactNode | -          |
| style     | 自定义样式对象                         | CSSProperties | -      |
| onClick   | 点击回调事件                           | (e: event) => void | -   |
| onRightClick | 右键点击回调事件                    | (e: event) => void | -   |

---

## 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ------ | ---- |
| $color-list_default-border-default | var(--semi-color-border) | 列表描边颜色 |
| $color-list_empty-text-default | var(--semi-color-text-2) | 空状态下列表文字颜色 |
