# DatePicker 日期选择器

## 组件简介
日期选择器用于帮助用户选择一个符合要求的、格式化的日期（时间）或日期（时间）范围。

---

## API 参考

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| ---- | ---- | ---- | ------ | ---- |
| autoAdjustOverflow | 浮层被遮挡时是否自动调整方向 | boolean | true |  |
| autoFocus | 自动获取焦点 | boolean | false |  |
| autoSwitchDate | 通过面板上方左右按钮、下拉菜单更改年月时，自动切换日期。仅对 date type 生效。 | boolean | true |  |
| borderless | 无边框模式 | boolean |  | 2.33.0 |
| bottomSlot | 渲染底部额外区域 | ReactNode |  |  |
| className | 类名 | string | - |  |
| clearIcon | 可用于自定义清除按钮, showClear为true时有效 | ReactNode |  | 2.25.0 |
| defaultOpen | 面板默认显示或隐藏 | boolean | false |  |
| defaultPickerValue | 默认面板日期 | ValueType |  |  |
| defaultValue | 默认值 | ValueType |  |  |
| density | 面板的尺寸，可选值：default, compact | string | default |  |
| disabled | 是否禁用 | boolean | false |  |
| disabledDate | 日期禁止判断方法，返回为 true 时禁止该日期，options 参数 1.9.0 后支持，其中 rangeEnd 1.29 后支持，rangeInputFocus 2.22 后支持 | (date, options) => boolean | () => false |  |
| disabledTime | 时间禁止配置，返回值将会作为参数透传给 TimePicker | (date) => object | () => false |  |
| disabledTimePicker | 是否禁止时间选择 | boolean |  |  |
| dropdownClassName | 下拉列表的 CSS 类名 | string |  |  |
| dropdownStyle | 下拉列表的内联样式 | object |  |  |
| dropdownMargin | 下拉列表算溢出时的增加的冗余值 | object\|number |  | 2.25.0 |
| endDateOffset | type 为 dateRange 时，设置单击选择范围的结束日期 | (selectedDate?: Date) => Date | - |  |
| format | 在输入框内展现的日期串格式 | string | 与 type 对应 |  |
| getPopupContainer | 指定父级 DOM，弹层将会渲染至该 DOM 中 | function():HTMLElement | () => document.body |  |
| hideDisabledOptions | 隐藏禁止选择的时间 | boolean | false |  |
| insetInput | 面板中是否嵌入输入框，InsetInputProps 类型 v2.29 支持 | boolean | false | 2.7.0 |
| inputReadOnly | 文本框是否 readonly | boolean | false |  |
| inputStyle | 输入框样式 | object |  |  |
| leftSlot | 渲染左侧额外区域 | ReactNode |  | 2.65.0 |
| max | multiple 为 true 时，多选的数目 | number | - |  |
| motion | 是否开启面板展开的动画 | boolean | true |  |
| multiple | 是否可以选择多个，仅支持 type="date" | boolean | false |  |
| needConfirm | 是否需要“确认选择”，仅 type="dateTime"\|"dateTimeRange" 时有效 | boolean |  |  |
| open | 面板显示或隐藏的受控属性 | boolean |  |  |
| placeholder | 输入框提示文字 | string\|string[] | 'Select date' |  |
| position | 浮层位置 | string | 'bottomLeft' |  |
| prefix | 前缀内容 | string\|ReactNode |  |  |
| presets | 日期时间快捷方式, start 和 end 在 v2.52 版本支持函数类型 | [] |  |  |
| preventScroll | 指示浏览器是否应滚动文档以显示新聚焦的元素 | boolean |  |  |
| presetPosition | 日期时间快捷方式面板位置 | string | 'bottom' | 2.18.0 |
| rangeSeparator | 自定义范围类型输入框的日期分隔符 | string | '~' |  |
| renderDate | 自定义日期显示内容 | (dayNumber, fullDate) => ReactNode | - |  |
| renderFullDate | 自定义显示日期格子内容 | (dayNumber, fullDate, dayStatus) => ReactNode | - |  |
| rightSlot | 渲染右侧额外区域 | ReactNode |  | 2.65.0 |
| showClear | 是否显示清除按钮 | boolean | true |  |
| size | 尺寸，可选值："small", "default", "large" | string | 'default' |  |
| spacing | 浮层与 trigger 的距离 | number | 4 |  |
| startDateOffset | type 为 dateRange 时，设置单击选择范围的开始日期 | (selectedDate?: Date) => Date | - |  |
| startYear | 滚轮的开始年 | number | 当前年前 100 年 | 2.36.0 |
| endYear | 滚轮的结束年，结束年需要大于开始年 | number | 当前年后 100 年 | 2.36.0 |
| stopPropagation | 是否阻止弹出层上的点击事件冒泡 | boolean | true |  |
| style | 自定义样式 | CSSProperties |  |  |
| syncSwitchMonth | 在范围选择的场景中，支持同步切换双面板的月份 | boolean | false |  |
| timePickerOpts | 其他可以透传给时间选择器的参数 | object |  |  |
| topSlot | 渲染顶部额外区域 | ReactNode |  |  |
| triggerRender | 自定义触发器渲染方法 | (props) => ReactNode |  |  |
| type | 类型，可选值："date", "dateRange", "dateTime", "dateTimeRange", "month", "monthRange" | string | 'date' |  |
| validateStatus | 校验状态，可选值 default、error、warning，默认 default。仅影响展示样式 | string |  |  |
| value | 受控的值 | ValueType |  |  |
| weekStartsOn | 以周几作为每周第一天，0 代表周日，1 代表周一，以此类推 | number | 0 |  |
| zIndex | 弹出面板的 zIndex | number | 1030 |  |
| onBlur | 失去焦点时的回调 | (e: event) => void | () => {} |  |
| onCancel | 取消选择时的回调，仅 type="dateTime"\|"dateTimeRange" 且 needConfirm=true 时有效 | (value) => void |  |  |
| onChange | 值变化时的回调 | (value, valueStr) => void |  |  |
| onChangeWithDateFirst | 0.x 中 onChange(string, Date), 1.0 后(Date, string)。此开关设为 false 时，入参顺序将与 0.x 版本保持一致 | boolean | true |  |
| onClear | 点击 clear 按钮时触发 | (e: event) => void | () => {} |  |
| onClickOutSide | 弹出层展示时，点击非弹出层、触发器的回调 | (event: React.mouseEvent) => void | () => {} | 2.31.0 |
| onConfirm | 确认选择时的回调，仅 type="dateTime"\|"dateTimeRange" 且 needConfirm=true 时有效 | (value) => void |  |  |
| onFocus | 获得焦点时的回调 | (e: event) => void | () => {} |  |
| onOpenChange | 面板显示或隐藏状态切换的回调 | (open) => void |  |  |
| onPanelChange | 切换面板的年份或者日期时的回调 | (date, dateStr) => void |  |  |
| onPresetClick | 点击快捷选择按钮的回调 | () => {} |  |  |
| yearAndMonthOpts | 其他可以透传给年月选择器的参数 | object |  | 2.20.0 |

---

## Methods

| 方法 | 说明 | 类型 | 版本 |
| ---- | ---- | ---- | ---- |
| open | 手动展开下拉列表 | () => void | 2.31.0 |
| close | 手动关闭下拉列表 | () => void | 2.31.0 |
| focus | 手动聚焦输入框 | (focusType?: 'rangeStart' \| 'rangeEnd') => void | 2.31.0 |
| blur | 手动失焦输入框 | () => void | 2.31.0 |

---

## 类型定义

```typescript
type BaseValueType = string | number | Date;
type ValueType = BaseValueType | BaseValueType[];
type DateType = Date | Date[];
type StringType = string | string[];
type TriggerRenderProps = {
    value?: ValueType;
    inputValue?: string;
    placeholder?: string | string[];
    autoFocus?: boolean;
    size?: InputSize;
    disabled?: boolean;
    inputReadOnly?: boolean;
    componentProps?: DatePickerProps;
    [x: string]: any;
};
```

---

## Accessibility

- 未选中日期时，触发器的 aria-label 为 Choose date，选中日期时为 Change date
- 日期面板中月的 role 为 grid，周的 role 设置为 row，日期格子设置为 gridcell
- 禁用时 aria-disabled 为 true
- 多选时，月的 aria-multiselectable 为 true，选中时日期格子的 aria-selected 为 true
- 装饰 icon 的 aria-hidden 为 true

---

## 日期时间格式

Semi UI 组件库中采用 date-fns(v2.9.0) 作为日期时间引擎，格式化 token 含义如下：

- "y" ：年
- "M" ：月
- "d" ：日
- "H" ：小时
- "m" ：分钟
- "s" ：秒

示例：

| 类型 | format | 展示值 |
| ---- | ------ | ------ |
| date | yyyy-MM-dd | 2023-12-09 |
| dateTime | yyyy-MM-dd HH:mm:ss | 2023-12-09 08:08:00 |
| month | yyyy-MM | 2023-12 |
| dateRange | yyyy-MM-dd | 2023-12-09 ～ 2023-12-10 |
| dateTimeRange | yyyy-MM-dd HH:mm:ss | 2023-12-09 08:08 ～ 2023-12-10 10:08 |

---

## FAQ

- 如何设置面板打开时默认显示的时间？可通过 defaultPickerValue 属性。
- 日期时间选择、范围日期选择，输入部分日期后，面板没有回显日期？输入框需要输入完整后才会回显到面板上。
- 日期时间选择面板底部的展示的时间是什么？未选择时间时，为 defaultPickerValue 的值，否则为已选择的时间。
