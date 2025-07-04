# Table 表格

表格用于呈现结构化的数据内容，通常会伴随提供对数据进行操作（排序、搜索、分页等）的能力。

## 基本用法

- 通过 `columns` 配置列，通过 `dataSource` 配置数据。
- 每行数据需有唯一 `key`，或通过 `rowKey` 指定主键属性。

## 主要属性

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| bordered | 是否展示外边框和列边框 | boolean | false |
| columns | 列配置，详见 Column | Column[] | [] |
| dataSource | 数据源 | RecordType[] | [] |
| rowKey | 行主键字段名或函数 | string \| (record) => string | 'key' |
| rowSelection | 行选择配置 | object | - |
| pagination | 分页配置 | boolean \| TablePaginationProps | true |
| scroll | 滚动配置 | object | - |
| sticky | 固定表头 | boolean \| { top: number } | false |
| loading | 是否加载中 | boolean | false |
| size | 表格尺寸 | "default"\|"middle"\|"small" | "default" |
| virtualized | 虚拟化配置 | Virtualized | false |
| onChange | 分页、排序、筛选变化时触发 | function | - |

## Column 主要属性

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| title | 列标题 | ReactNode \| function | - |
| dataIndex | 数据字段 | string | - |
| key | React key | string | - |
| width | 列宽 | string \| number | - |
| align | 对齐方式 | 'left' \| 'right' \| 'center' | 'left' |
| fixed | 固定列 | boolean \| 'left' \| 'right' | false |
| sorter | 排序函数或 true | boolean \| function | false |
| filters | 筛选项 | Filter[] | - |
| onFilter | 本地筛选函数 | function | - |
| render | 自定义渲染 | function | - |
| ellipsis | 文本缩略 | boolean \| { showTitle: boolean } | false |

## rowSelection 主要属性

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| selectedRowKeys | 选中行 key 数组 | string[] | - |
| onChange | 选中项变化回调 | function | - |
| getCheckboxProps | 选择框属性 | function | - |
| fixed | 固定选择列 | boolean | false |

## scroll 主要属性

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| x | 横向滚动宽度 | string \| number | - |
| y | 纵向滚动高度 | number | - |

## pagination 主要属性

| 属性 | 说明 | 类型 | 默认值 |
| ---- | ---- | ---- | ------ |
| currentPage | 当前页码 | number | - |
| pageSize | 每页条数 | number | 10 |
| total | 数据总数 | number | 0 |

## 方法

- `getCurrentPageData()`：返回当前页数据对象。
- 通过 ref 可访问 Table 内部方法。

## 事件

- `onChange({ pagination, filters, sorter, extra })`：分页、排序、筛选变化时触发。
- `onExpand(expanded, record, DOMEvent)`：点击行展开图标时触发。
- `onRow(record, index)`：设置行属性。
- `onHeaderRow(columns, index)`：设置头部行属性。

## 说明

- 支持树形数据、分组、虚拟化、伸缩列、固定表头/列、完全自定义渲染等高级功能。
- 详细用法和示例请参考 Semi Design 官方文档。

