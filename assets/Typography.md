# Typography API 参考

## Typography.Text

| 属性        | 说明 | 类型 | 默认值 | 版本 |
|-------------|------|------|--------|------|
| component   | 自定义渲染元素 | html element | span | - |
| code        | 是否被 code 元素包裹 | boolean | - | - |
| copyable    | 是否可拷贝 | boolean \| object:Copyable Config | false | - |
| delete      | 添加删除线样式 | boolean | false | - |
| disabled    | 禁用文本 | boolean | false | - |
| ellipsis    | 设置自动溢出省略 | boolean\|object:Ellipsis Config | false | - |
| icon        | 前缀图标 | ReactNode | - | - |
| link        | 是否为链接，传object时，属性将透传给a标签 | boolean\|object | false | - |
| mark        | 添加标记样式 | boolean | false | - |
| size        | 文本大小，可选normal，small，inherit | string | normal | - |
| strong      | 是否加粗 | boolean | false | - |
| type        | 文本类型，可选 primary, secondary, warning, danger, tertiary(v>=1.2.0), quaternary(v>=1.2.0), success(v>=1.7.0) | string | primary | - |
| underline   | 添加下划线样式 | boolean | false | - |
| weight      | 设置字重 | number | - | 2.34.0 |

---

## Typography.Title

| 属性        | 说明 | 类型 | 默认值 | 版本 |
|-------------|------|------|--------|------|
| component   | 自定义渲染元素，默认由 heading 决定 | html element | h1~h6 | - |
| copyable    | 是否可拷贝 | boolean \| object:Copyable Config | false | - |
| delete      | 添加删除线样式 | boolean | false | - |
| disabled    | 禁用文本 | boolean | false | - |
| ellipsis    | 设置自动溢出省略 | boolean\|object:Ellipsis Config | false | - |
| heading     | 标题级别，可选1，2，3，4，5，6 | number | 1 | - |
| link        | 是否为链接，传object时，属性将透传给a标签 | boolean\|object | false | - |
| mark        | 添加标记样式 | boolean | false | - |
| type        | 文本类型，同 Text | string | primary | - |
| underline   | 添加下划线样式 | boolean | false | - |
| weight      | 设置字重, 可选 light, regular, medium, semibold, bold, default | string, number | - | 2.34.0 |

---

## Typography.Paragraph

| 属性        | 说明 | 类型 | 默认值 | 版本 |
|-------------|------|------|--------|------|
| component   | 自定义渲染元素 | html element | p | - |
| copyable    | 是否可拷贝 | boolean \| object:Copyable Config | false | - |
| delete      | 添加删除线样式 | boolean | false | - |
| disabled    | 禁用文本 | boolean | false | - |
| ellipsis    | 设置自动溢出省略 | boolean\|object:Ellipsis Config | false | - |
| link        | 是否为链接，传object时，属性将透传给a标签 | boolean\|object | false | - |
| mark        | 添加标记样式 | boolean | false | - |
| size        | 文本大小，可选normal，small | string | normal | - |
| spacing     | 行距大小，可选normal，extended | string | normal | - |
| strong      | 是否加粗 | boolean | false | - |
| type        | 文本类型，同 Text | string | primary | - |
| underline   | 添加下划线样式 | boolean | false | - |

---

## Typography.Numeral

| 属性        | 说明 | 类型 | 默认值 | 版本 |
|-------------|------|------|--------|------|
| rule        | 解析规则，可选 text, numbers, bytes-decimal, bytes-binary, percentages, exponential | string | text | 2.22.0 |
| precision   | 小数点后保留位数 | number | 0 | 2.22.0 |
| truncate    | 小数点后保留位截段取整方式，可选 ceil, floor, round | string | round | 2.22.0 |
| parser      | 自定义数值解析函数 | (str: string) => string | - | 2.22.0 |
| component   | 自定义渲染元素 | html element | span | 2.22.0 |
| code        | 是否被 code 元素包裹 | boolean | - | 2.22.0 |
| copyable    | 是否可拷贝 | boolean \| object:Copyable Config | false | 2.22.0 |
| delete      | 添加删除线样式 | boolean | false | 2.22.0 |
| disabled    | 禁用文本 | boolean | false | 2.22.0 |
| ellipsis    | 设置自动溢出省略 | boolean\|object:Ellipsis Config | false | 2.22.0 |
| icon        | 前缀图标 | ReactNode | - | 2.22.0 |
| link        | 是否为链接，传object时，属性将透传给a标签 | boolean\|object | false | 2.22.0 |
| mark        | 添加标记样式 | boolean | false | 2.22.0 |
| size        | 文本大小，可选normal，small | string | normal | 2.22.0 |
| strong      | 是否加粗 | boolean | false | 2.22.0 |
| type        | 文本类型，同 Text | string | primary | 2.22.0 |
| underline   | 添加下划线样式 | boolean | false | 2.22.0 |

---

## Ellipsis Config

| 属性         | 说明 | 类型 | 默认值 |
|--------------|------|------|--------|
| collapseText | 折叠的展示文本 | string | 收起 |
| collapsible  | 是否支持折叠 | boolean | false |
| expandText   | 展开的展示文本 | string | 展开 |
| expandable   | 是否支持展开 | boolean | false |
| pos          | 省略截断的位置，支持末尾和中间截断：end, middle | string | end |
| rows         | 省略溢出行数 | number | 1 |
| showTooltip  | 是否展示 tooltip 及相关配置 | boolean\|object | false |
| suffix       | 始终展示的后缀 | string | - |
| onExpand     | 展开/收起的回调 | function(expanded: bool, Event: e) | - |

---

## Copyable Config

| 属性      | 说明 | 类型 | 默认值 | 版本 |
|-----------|------|------|--------|------|
| content   | 复制出的文本 | string | - | - |
| copyTip   | 复制图标的 tooltip 展示内容 | React.node | - | - |
| icon      | 自定义渲染复制节点 | React.node | - | 2.31.0 |
| onCopy    | 复制回调 | Function(e:Event, content:string, res:boolean) | - | - |
| render    | 自定义渲染复制节点 | - | - | 2.65.0 |
| successTip| 复制成功的展示内容 | React.node | - | - |
