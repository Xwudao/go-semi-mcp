# Tag 标签

## Tag Props

| 属性         | 说明                                                         | 类型         | 默认值   | 版本    |
| ------------ | ------------------------------------------------------------ | ------------ | -------- | ------- |
| avatarShape  | 头像 Tag 形状，可选 square 和 circle                         | string       | square   | 1.6.0   |
| avatarSrc    | 头像的资源地址                                               | string       | -        | 1.6.0   |
| className    | 类名                                                         | string       | -        | -       |
| closable     | 标签是否可以关闭                                             | boolean      | false    | -       |
| color        | 标签的颜色，可选 amber、blue、cyan、green、grey、indigo、light-blue、light-green、lime、orange、pink、purple、red、teal、violet、yellow、white | string | grey | - |
| prefixIcon   | 前缀图标                                                     | ReactNode    | -        | 2.44.0  |
| suffixIcon   | 后缀图标                                                     | ReactNode    | -        | 2.44.0  |
| shape        | 标签的形状，可选 square、circle                              | string       | square   | 2.20.0  |
| size         | 标签的尺寸，可选 small、large                                | string       | small    | -       |
| style        | 样式                                                         | CSSProperties| -        | -       |
| type         | 标签的样式类型，可选 ghost、solid、light                     | string       | light    | -       |
| visible      | 标签是否可见                                                 | boolean      | true     | -       |
| tagKey       | React 需要的 key，作为每个标签的唯一标识，不允许重复         | string/number| -        | -       |
| onClick      | 单击标签时的回调函数                                         | (e: MouseEvent) => void | 无 | - |
| onClose      | 关闭标签时的回调函数                                         | (tagChildren: ReactNode, e: MouseEvent, tagKey: string \| number ) => void | 无 | e 于 v1.18 提供, tagKey 于 v2.18.0 提供 |

## TagGroup Props

| 属性         | 说明                                                         | 类型         | 默认值   | 版本    |
| ------------ | ------------------------------------------------------------ | ------------ | -------- | ------- |
| avatarShape  | 头像 Tag 形状，可选 square 和 circle                         | string       | square   | 1.6.0   |
| className    | 类名                                                         | string       | -        | -       |
| maxTagCount  | 最大数量限制，超出后显示为 +N                                | number       | -        | -       |
| popoverProps | popover 的配置属性，可控制 direction, zIndex, trigger 等     | PopoverProps | {}       | -       |
| showPopover  | hover 到 +N 时，是否通过 Popover 显示剩余内容                | boolean      | false    | -       |
| size         | 标签的尺寸，可选 small、large                                | string       | small    | -       |
| style        | 样式                                                         | CSSProperties| -        | -       |
| tagList      | 标签组                                                       | (TagProps)[] | -        | -       |
| onTagClose   | 删除TagGroup中的Tag时候的回调函数                            | (tagChildren: ReactNode, e: MouseEvent, tagKey: string \| number ) => void | - | 2.18.0 |

## Accessibility

- aria-label 用于表示 Tag 的作用，对于可删除或者可点击的 Tag ，推荐使用此属性
- 可交互的 Tag 可被聚焦，支持键盘操作

## 设计变量

| 变量                                   | 默认值                           | 用法                       |
| -------------------------------------- | -------------------------------- | -------------------------- |
| $color-tag_avatar-bg-default           | var(--semi-color-bg-4)           | 头像标签背景颜色 - 默认    |
| $color-tag_avatar-border-default       | var(--semi-color-border)         | 头像标签描边颜色 - 默认    |
| $color-tag_avatar-text-default         | var(--semi-color-text-0)         | 头像标签文字颜色 - 默认    |
| $color-tag_white-bg-default            | var(--semi-color-bg-4)           | 白色标签背景颜色 - 默认    |
| $color-tag_white-border-default        | rgba(var(--semi-grey-2), 0.7)    | 白色标签描边颜色 - 默认    |
| $color-tag_white-text-default          | var(--semi-color-text-0)         | 白色标签文字颜色 - 默认    |
| $color-tag_white-icon-default          | var(--semi-color-text-2)         | 白色标签图标颜色 - 默认    |
| $color-tag-outline-focus               | var(--semi-color-primary-light-active) | 标签轮廓 - 聚焦      |
| $color-tag_close-icon-default          | var(--semi-color-text-2)         | 可删除的标签删除按钮颜色   |
| $color-tag_close-icon-hover            | var(--semi-color-text-1)         | 可删除的标签删除按钮颜色-悬浮 |
