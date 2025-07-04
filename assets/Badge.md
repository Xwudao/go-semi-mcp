# Badge 徽章

用徽章来给用户提示。

## API 参考

| 属性           | 说明 | 类型 | 默认值 |
| -------------- | ---- | ---- | ------ |
| children       | 徽章的 base | ReactNode | 无 |
| className      | 外侧 className | string | - |
| count          | 展示的内容 | ReactNode | 无 |
| countClassName | 内容区域 className | string | 无 |
| dot            | 不展示数字，显示小圆点 | boolean | false |
| overflowCount  | 最大的展示数字值 | number | 无 |
| position       | 徽章位置，可选 leftTop、leftBottom、rightTop、rightBottom | string | rightTop |
| countStyle     | 徽章内容的样式，v2.59.1后生效 | CSSProperties | 无 |
| theme          | 徽章主题，可选 solid、light、inverted | string | solid |
| type           | 徽章类型，可选 primary、secondary、tertiary、danger、warning、success | string | primary |

## 文案规范

- Badge 内容若为英文时，首字母应大写

## 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ------ | ---- |
| $color-badge_default-border-default | var(--semi-color-bg-1) | 描边颜色 - 默认 |
| $color-badge_default-bg-default | var(--semi-color-bg-1) | 背景颜色 - 默认 |
| $color-badge_default_light-bg-default | var(--semi-color-bg-2) | 背景颜色 - 浅版 |
| $color-badge_default-text-default | var(--semi-color-bg-2) | 文字颜色 - 默认 |
| $color-badge_primary_solid-bg-default | var(--semi-color-primary) | 背景颜色 - 主色 |
| $color-badge_primary_light-bg-default | var(--semi-color-primary-light-default) | 背景颜色 - 浅版主色 |
| $color-badge_primary_light-text-default | var(--semi-color-primary) | 文字颜色 - 浅版主色 |
| $color-badge_primary_inverted-text-default | var(--semi-color-primary) | 文字颜色 - 白底 |
| $color-badge_secondary_solid-bg-default | var(--semi-color-secondary) | 背景颜色 - 次要 |
| $color-badge_secondary_light-bg-default | var(--semi-color-secondary-light-default) | 背景颜色 - 浅版次要 |
