# ConfigProvider 全局配置

为组件提供统一的全局化配置。

## 使用场景

- 覆盖多个组件公有 Props 配置（如 timezone、rtl），使用 ConfigProvider。
- 若需全局修改某组件某类 Props（如所有 Button 的 theme），可用 semiGlobal。

## 基本用法

```js
import { ConfigProvider } from '@douyinfe/semi-ui';
```

通过传入 `timeZone` 参数，为时间类组件配置时区。

## API 参考

| 属性                | 说明                                                                 | 类型                | 默认值 |
|---------------------|----------------------------------------------------------------------|---------------------|--------|
| direction           | 设置文本的方向                                                       | ltr \| rtl          | ltr    |
| getPopupContainer   | 指定父级 DOM，弹层渲染至该 DOM，自定义需设置 position: relative      | () => HTMLElement   | () => document.body |
| locale              | 多语言配置，同 LocaleProvider 中 locale 参数                         | object              |        |
| timeZone            | 时区标识，数字或字符串（如 GMT+08:00、Asia/Shanghai）                | string \| number    |        |

### 说明

- `direction`：rtl 表示从右到左，ltr 表示从左到右。
- `timeZone`：支持数字、GMT 字符串、IANA 标识。推荐使用 IANA 标识。
- `getPopupContainer`：自定义弹层父节点。

## semiGlobal

通过 `semiGlobal.config.overrideDefaultProps` 可全局覆盖组件默认 Props。

```js
import { semiGlobal } from "@douiyinfe/semi-ui";

semiGlobal.config.overrideDefaultProps = {
    Select: { zIndex: 2000 },
    Tooltip: { zIndex: 2001, trigger: 'click' },
};
```
