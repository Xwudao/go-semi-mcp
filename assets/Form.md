# Form 表单

Semi Design 的 Form 组件提供高性能、结构简洁、无障碍支持的表单解决方案，支持多种表单控件、布局、校验和扩展机制。

## 主要特性

- 按需重绘，避免不必要的全量渲染
- 结构极简，避免多余嵌套
- 完善的无障碍支持
- 支持外部获取 formState / fieldState
- 提供 formApi / fieldApi 操作表单
- 支持自定义组件接入
- 支持 Form/Field 级别赋值与校验（同步/异步）

## 基本用法

```jsx
import { Form } from '@douyinfe/semi-ui';

<Form layout='horizontal' onValueChange={values=>console.log(values)}>
    <Form.Input field='UserName' label='用户名' style={{ width: 80 }}/>
    <Form.Input field='Password' label='密码' style={{ width: 80 }}/>
</Form>
```

## 支持的表单控件

- Input、InputNumber、TextArea、Select、Checkbox、Radio、RadioGroup、Switch、DatePicker、TimePicker、Slider、InputGroup、TreeSelect、Cascader、Rating、AutoComplete、Upload、Label、ErrorMessage、Section、TagInput

## 表单布局

- `layout`: 'vertical'（默认）或 'horizontal'
- `labelPosition`: 'top' | 'left' | 'inset'
- `labelAlign`: 'left' | 'right'
- 支持 Grid 布局、分组（Form.Section）、InputGroup 组合

## 校验与初始值

- 支持 rules、trigger、initValue、initValues
- 支持同步/异步校验
- 支持自定义校验函数（Form/Field 级别）

## API 参考

### Form Props

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| autoScrollToError | 校验失败时自动滚动到出错字段 | boolean/object | false |
| allowEmpty | 是否保留空值字段 | boolean | false |
| className | 表单 class | string |  |
| component | 直接声明表单控件 | ReactNode |  |
| disabled | 统一禁用 | boolean | false |
| extraTextPosition | extraText 显示位置 | string | 'bottom' |
| getFormApi | 获取 formApi 回调 | function |  |
| initValues | 初始值 | object |  |
| layout | 布局 | string | 'vertical' |
| labelAlign | label 对齐 | string | 'left' |
| labelCol | label 布局 | object |  |
| labelPosition | label 位置 | string | 'top' |
| labelWidth | label 宽度 | string/number |  |
| onChange | 表单更新回调 | function |  |
| onValueChange | 值变更回调 | function |  |
| onErrorChange | 校验状态变更回调 | function |  |
| onReset | 重置回调 | function |  |
| onSubmit | 提交回调 | function |  |
| onSubmitFail | 提交失败回调 | function |  |
| render | render props | function |  |
| showValidateIcon | 校验信息区块显示 icon | boolean | true |
| style | 内联样式 | object |  |
| stopValidateWithError | 校验遇错即停 | boolean | false |
| stopPropagation | 阻止事件冒泡 | object |  |
| trigger | 校验触发时机 | string/array | 'change' |
| validateFields | 自定义校验函数 | function |  |
| wrapperCol | Field 布局 | object |  |

### FormState

| Name | 说明 | 初始值 | 示例 |
| ---- | ---- | ------ | ---- |
| values | 表单的值 | {} | { fieldA: 'str' } |
| errors | 错误信息 | {} | { fieldA: 'error' } |
| touched | 点击过的字段 | {} | { fieldA: true } |

### FormApi

- `getFormProps()`
- `getFormState()`
- `submitForm()`
- `reset(fields?)`
- `validate(fields?)`
- `setValues(newValues, { isOverride })`
- `setValue(field, value)`
- `getValue(field)`
- `getValues()`
- `setTouched(field, isTouched)`
- `getTouched(field)`
- `setError(field, error)`
- `getError(field)`
- `getFieldExist(field)`
- `scrollToField(field, opts)`
- `scrollToError(fieldOrIndex)`

### Field Props

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| field | 字段名（必填） | string |  |
| label | 标签 | string/object |  |
| labelPosition | label 位置 | string |  |
| labelAlign | label 对齐 | string |  |
| labelWidth | label 宽度 | string/number |  |
| noLabel | 不自动插入 label | boolean |  |
| noErrorMessage | 不自动插入 ErrorMessage | boolean |  |
| name | 控件名称 | string |  |
| fieldClassName | fieldWrapper class | string |  |
| fieldStyle | fieldWrapper 样式 | object |  |
| initValue | 初始值 | any |  |
| validate | 自定义校验函数 | function |  |
| rules | 校验规则 | array |  |
| validateStatus | 校验状态 | string | 'default' |
| trigger | 校验触发时机 | string/array | 'change' |
| onChange | 值变化回调 | function |  |
| onBlur | 失焦回调 | function |  |
| transform | 校验前转换 | function |  |
| allowEmptyString | 允许空字符串 | boolean | false |
| stopValidateWithError | 校验遇错即停 | boolean | false |
| helpText | 自定义提示 | ReactNode |  |
| extraText | 额外提示 | ReactNode |  |
| pure | 仅接管数据流 | boolean | false |
| extraTextPosition | extraText 显示位置 | string | 'bottom' |

### 其他

- 支持 ArrayField、Form.Section、Form.Label、Form.InputGroup、Form.Slot、Form.ErrorMessage
- 支持 hooks: useFormApi, useFormState, useFieldApi, useFieldState
- 支持 HOC: withFormApi, withFormState, withField

## FAQ

- field 属性必填，未传不会自动映射
- defaultValue、defaultChecked 不生效，需用 initValue/initValues
- initValue/initValues 只在挂载时消费
- setValues 只会更新已存在的 field
- async-validator 的自定义 validator 返回值需为 boolean

更多细节请参考 Semi Design 官方文档。
