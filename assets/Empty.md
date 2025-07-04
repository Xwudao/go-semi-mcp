# Empty 空状态

## API 参考

| 属性         | 说明                                                         | 类型                                             | 默认值     |
|--------------|--------------------------------------------------------------|--------------------------------------------------|------------|
| className    | 类名                                                         | string                                           | -          |
| darkModeImage| 暗色模式下的占位图，响应 theme-mode 属性变化                 | ReactNode                                        | -          |
| description  | 内容描述                                                     | ReactNode                                        | -          |
| image        | 占位图                                                       | ReactNode \| { id?: string; viewBox?: string; url?: string;} | -          |
| imageStyle   | 占位图样式                                                   | CSSProperties                                    | -          |
| layout       | 布局方式，支持 vertical, horizontal                          | string                                           | vertical   |
| style        | 样式名                                                       | CSSProperties                                    | -          |
| title        | 标题                                                         | ReactNode                                        | -          |

## Accessibility

- Empty 插图的 aria-hidden 为 true

## 文案规范

- 标题应简洁易懂
- 正文可展示空状态原因或后续操作，避免重复标题内容，建议 1-2 句话
- 按钮文案需清晰易懂，推荐“动词+名词”格式

## 设计变量

| 变量                              | 默认值                        | 用法         |
|-----------------------------------|-------------------------------|--------------|
| $color-empty_description-text-default | var(--semi-color-text-1)     | 描述文字颜色 |
