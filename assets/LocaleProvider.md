# LocaleProvider 多语言

为 Semi 组件提供多语言支持。

## 支持语言

| 语言         | 代码      | 版本      |
|--------------|-----------|-----------|
| 简体中文     | zh_CN     | v0.0.1    |
| 英语         | en_GB     | v0.7.0    |
| 日语         | ja_JP     | v0.7.0    |
| 韩语         | ko_KR     | v0.7.0    |
| ...          | ...       | ...       |

## 用法

```js
import zh_CN from '@douyinfe/semi-ui/lib/es/locale/source/zh_CN';
import en_GB from '@douyinfe/semi-ui/lib/es/locale/source/en_GB';
import { LocaleProvider } from '@douyinfe/semi-ui';

<LocaleProvider locale={en_GB}>
    <App />
</LocaleProvider>
```

切换语言时，直接切换 `locale` 参数。

## 支持多语言的组件

Calendar、Cascader、Chat、DatePicker、Form、Image、List、Modal、Navigation、Nav、Pagination、Popconfirm、Select、Table、TimePicker、Transfer、Tree、TreeSelect、Typography、Upload 等。

## 自定义国际化组件

可通过 `LocaleConsumer` 获取 context 中的 localeCode 或 localeData。

```js
import { LocaleProvider, LocaleConsumer } from '@douyinfe/semi-ui';

<LocaleConsumer componentName="TimePicker">
    {/* ... */}
</LocaleConsumer>
```
