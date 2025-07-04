# Timeline 时间轴

## Timeline Props

| 属性        | 说明                                 | 类型                                   | 默认值  |
| ----------- | ------------------------------------ | -------------------------------------- | ------- |
| className   | 类名                                 | string                                 | -       |
| mode        | 时间轴和内容的相对位置               | left \| right \| center \| alternate   | left    |
| style       | 样式                                 | CSSProperties                          | -       |
| dataSource  | 时间轴数据源，支持 content 属性及 TimeLine.Item 的所有属性 | array | - |

## Timeline.Item Props

| 属性      | 说明                   | 类型                                   | 默认值    | 版本    |
| --------- | ---------------------- | -------------------------------------- | --------- | ------- |
| className | 类名                   | string                                 | -         | -       |
| color     | 自定义的圆圈色值       | string                                 | -         | -       |
| dot       | 自定义时间轴点         | ReactNode                              | -         | -       |
| extra     | 自定义辅助内容         | ReactNode                              | -         | -       |
| position  | 自定义节点位置         | left \| right                          | -         | -       |
| style     | 样式                   | CSSProperties                          | -         | -       |
| time      | 时间文本               | string                                 | -         | -       |
| type      | 当前圆圈的模式         | default \| ongoing \| success \| warning \| error | default | -       |
| onClick   | 鼠标点击事件的回调     | (e: MouseEvent) => void                | -         | 2.2.0   |

## Accessibility

- 组件中时间点的连线以及时间点本身被设置了 aria-hidden，不会响应 Accessibility API
- 可以通过传入 aria-label 设置 TimeLine 组件的标签

## 设计变量

| 变量                                   | 默认值                           | 用法                       |
| -------------------------------------- | -------------------------------- | -------------------------- |
| $color-timeline_dot_default-bg-default | var(--semi-color-primary)        | 时间轴节点圆点背景色-进行中 |
| $color-timeline_dot_success-bg-default | var(--semi-color-success)        | 时间轴节点圆点背景色-成功   |
| $color-timeline_dot_warning-bg-default | var(--semi-color-warning)        | 时间轴节点圆点背景色-警告   |
| $color-timeline_dot_error-bg-default   | var(--semi-color-danger)         | 时间轴节点圆点背景色-错误   |
| $color-timeline_dot_info-bg-default    | var(--semi-color-tertiary-light-active) | 时间轴节点圆点背景色-默认   |
| $color-timeline_tail-border            | var(--semi-color-text-3)         | 时间轴连线颜色             |
| $color-timeline_time_default-text-default | var(--semi-color-text-2)        | 时间轴文字颜色-默认        |
| $color-timeline_item_head-bg           | transparent                      | 时间轴头部背景颜色         |
| $color-timeline_item_content-text-default | var(--semi-color-text-0)        | 时间轴标题文字颜色         |
