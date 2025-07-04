# Calendar 日历

日历组件，允许以日/周/月视图展示对应事件。

## API 参考

### Calendar

| 属性                | 说明                                         | 类型                | 默认值    |
|---------------------|----------------------------------------------|---------------------|-----------|
| dateGridRender      | 自定义单元格/列渲染，需要绝对定位            | function(dateString: string, date: Date) | -         |
| allDayEventsRender  | 自定义日/多日/周视图下的顶部事件渲染         | function(events: EventObject[]): ReactNode | -      |
| displayValue        | 展示日期                                     | Date                | 当前日期  |
| events              | 渲染事件，具体格式请参考 event object         | EventObject[]       | -         |
| header              | 自定义头部内容                               | ReactNode           | -         |
| height              | 日历高度                                     | string\|number      | 600       |
| markWeekend         | 区分周末列和工作日，以灰色显示               | boolean             | false     |
| minEventHeight      | 日/多日/周视图下事件的最小高度               | number              | Number.MIN_SAFE_INTEGER |
| mode                | 初始模式，day, week, month, range            | "day"\|"week"\|"month"\|"range" | week |
| onClick             | 单击日期格的回调                             | function(e: Event, date: Date） | -    |
| onClose             | 月视图下，展示所有 event 的卡片关闭时的回调  | function(e: Event） | -         |
| range               | 多日视图模式下展示的日期范围，左闭右开        | Date[]              | -         |
| renderTimeDisplay   | 自定义日/周视图下的时间文案                  | function(time: number): ReactNode | -   |
| renderDateDisplay   | 自定义日期文案                               | function(date: Date): ReactNode | -   |
| scrollTop           | 日/周视图模式下，设置展示内容默认的滚动高度   | number              | 400       |
| showCurrTime        | 显示当前时间                                 | boolean             | true      |
| width               | 日历宽度                                     | string\|number      | -         |
| weekStartsOn        | 以周几作为每周第一天，0 代表周日，1 代表周一  | number              | 0         |

### Event Object

| 属性      | 说明             | 类型      | 默认值 |
|-----------|------------------|-----------|--------|
| allDay    | 全天事件         | boolean   | false  |
| children  | 展示日期         | ReactNode | -      |
| end       | 事情结束的时间   | Date      | -      |
| key       | required 且唯一  | string    | -      |
| start     | 事情起始的时间   | Date      | -      |

### 设计变量

| 变量                              | 默认值                        | 用法                   |
|-----------------------------------|-------------------------------|------------------------|
| $color-calendar-bg-active         | var(--semi-color-primary)     | 今日标识颜色           |
| $color-calendar-bg                | var(--semi-color-bg-2)        | 日历背景颜色           |
| $color-calendar-text-active       | var(--semi-color-bg-1)        | 日历文字颜色 - 今日    |
| $color-calendar_curr-bg           | var(--semi-color-danger)      | 当前时间标识线颜色     |
| $color-calendar_curr-border       | var(--semi-color-danger)      | 当前时间标识线颜色     |
| $color-calendar_currCircle-bg-default | var(--semi-color-danger)  | 当前时间标识原点颜色   |
| $color-calendar_date-text-default | var(--semi-color-text-0)      | 日历文字颜色 - 当月    |
| $color-calendar_day-border        | var(--semi-color-border)      | 日历描边颜色           |
| $color-calendar_day-text-default  | var(--semi-color-text-2)      | 日历文字颜色 - 非当月  |
| $color-calendar_sticky-bg         | var(--semi-color-bg-2)        | 日历背景颜色 - 吸顶部分|

### 文案规范

- 12 小时制和 24 小时制都可用，12 小时制需搭配 AM/PM。
- 月份、星期、时间缩写规则可参考官方规范。
