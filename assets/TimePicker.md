# TimePicker 时间选择器

时间选择器用于选择某一符合要求的、格式化的时间点。

## API 参考

| 参数 | 说明 | 类型 | 默认值 | 版本 |
| ---- | ---- | ---- | ------ | ---- |
| autoAdjustOverflow | 浮层被遮挡时是否自动调整方向 | boolean | true |  |
| autoFocus | 自动获取焦点 | boolean | false |  |
| borderless | 无边框模式 | boolean |  | >=2.33.0 |
| className | 外层样式名 | string |  |  |
| clearIcon | 自定义清除按钮, showClear为true时有效 | ReactNode |  | 2.25.0 |
| defaultOpen | 面板是否默认打开 | boolean |  |  |
| defaultValue | 默认时间 | Date\|timeStamp\|String（type="timeRange"时为数组） |  |  |
| disabled | 禁用全部操作 | boolean | false |  |
| disabledHours | 禁止选择部分小时选项 | Function(): number[] |  |  |
| disabledMinutes | 禁止选择部分分钟选项 | Function(selectedHour: number): number[] |  |  |
| disabledSeconds | 禁止选择部分秒选项 | Function(selectedHour: number, selectedMinute: number): number[] |  |  |
| dropdownMargin | 浮层算溢出时的增加的冗余值 | object\|number |  | 2.25.0 |
| focusOnOpen | 挂载时是否打开面板并focus输入框 | boolean | false |  |
| format | 展示的时间格式 | string | "HH:mm:ss" |  |
| getPopupContainer | 指定容器，浮层将会渲染至该元素内 | Function(): HTMLElement | () => document.body |  |
| hideDisabledOptions | 隐藏禁止选择的选项 | boolean | false |  |
| hourStep | 小时选项间隔 | number | 1 |  |
| inputReadOnly | 设置输入框为只读 | boolean | false |  |
| minuteStep | 分钟选项间隔 | number | 1 |  |
| motion | 是否展示弹出层动画 | boolean | true |  |
| open | 面板是否打开的受控属性 | boolean |  |  |
| panelFooter | 面板底部 addon | ReactNode\|ReactNode[]\|string | 无 |  |
| panelHeader | 面板头部 addon | ReactNode\|ReactNode[]\|string | 无 |  |
| placeholder | 没有值的时候显示的内容 | string | "请选择时间" |  |
| popupClassName | 弹出层类名 | string | '' |  |
| popupStyle | 弹出层样式对象 | object | - |  |
| position | 浮层位置 | string | type="timeRange"时默认为"bottom"，type="time"时默认为"bottomLeft" |  |
| prefix | 前缀内容 | string\|ReactNode |  |  |
| preventScroll | 指示浏览器是否应滚动文档以显示新聚焦的元素，作用于组件内的 focus 方法 | boolean |  |  |
| rangeSeparator | 时间范围分隔符 | string | " ~ " |  |
| scrollItemProps | 透传给 scrollItem 的属性 | object |  |  |
| secondStep | 秒选项间隔 | number | 1 |  |
| showClear | 是否展示清除按钮 | boolean | true |  |
| stopPropagation | 是否阻止弹出层上的点击事件冒泡 | boolean | true | 2.49.0 |
| size | 输入框的大小，可选 'default'，'small'，'large' | string | 'default' |  |
| triggerRender | 自定义触发器渲染方法 | ({ placeholder: string }) => ReactNode | - |  |
| type | 类型 | "time"\|"timeRange" | "time" |  |
| use12Hours | 使用 12 小时制 | boolean | false |  |
| value | 当前时间 | Date\|timeStamp\|String（type="timeRange"时为数组） |  |  |
| onBlur | 失去焦点时的回调 | (e: domEvent) => void | () => {} |  |
| onChange | 时间发生变化的回调 | Function(time: Date, timeString: string): void （type="timeRange"时入参皆为数组） | 无 |  |
| onChangeWithDateFirst | onChange 的入参顺序为 (Date, string) | boolean | true | 2.4.0 |
| onFocus | 获得焦点时的回调 | (e: domEvent) => void | () => {} |  |
| onOpenChange | 面板打开/关闭时的回调 | Function(isOpen: boolean): void | 无 |  |

### Methods

| 名称 | 描述 |
| ---- | ---- |
| blur() | 移除焦点 |
| focus() | 获取焦点 |

### 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ------ | ---- |
| $color-timePicker_range_picker_panel_split-border | var(--semi-color-border) | 时间选择器菜单分割线颜色 |
| $color-timePicker_range_panel-border | rgba(0, 0, 0, .1) | 时间选择器描边颜色 |
