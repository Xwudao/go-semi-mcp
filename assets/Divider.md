# Divider 分割线 API 参考

## 属性

| 属性       | 说明                                   | 类型                              | 默认值      | 版本    |
| ---------- | -------------------------------------- | --------------------------------- | ----------- | ------- |
| align      | 带内容时，内容对齐方式                 | 'left' \| 'center' \| 'right'     | 'center'    | 2.9.0   |
| children   | 内容                                   | ReactNode                         | 无          | 2.9.0   |
| className  | 类名                                   | string                            | 无          | 2.9.0   |
| dashed     | 是否为虚线                             | boolean                           | false       | 2.9.0   |
| layout     | 分割线方向                             | 'horizontal' \| 'vertical'        | 'horizontal'| 2.9.0   |
| margin     | 分割线上下 margin（垂直时为左右 margin）| number \| string                  | 无          | 2.9.0   |
| style      | 自定义样式                             | CSSProperties                     | 无          | 2.9.0   |

## 设计变量

| 变量                        | 默认值                        | 用法         |
| --------------------------- | ----------------------------- | ------------ |
| $color-divider_border-color | var(--semi-color-border)      | 分割线颜色   |
| $color-divider_text-default | var(--semi-color-text-0)      | 标题颜色     |
