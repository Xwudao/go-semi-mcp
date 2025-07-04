# Card 卡片

常规的卡片容器，可以承载标题、段落、图片、列表等内容。

## API 参考

### Card

| 属性                | 说明                                         | 类型                | 默认值    | 版本     |
|---------------------|----------------------------------------------|---------------------|-----------|----------|
| actions             | 卡片操作组，位于卡片内容区的底部             | Array<ReactNode>    | -         | 1.21.0   |
| bodyStyle           | 卡片内容区内联样式                           | CSSProperties       | -         | 1.21.0   |
| bordered            | 是否设置卡片的外边框                         | boolean             | true      | 1.21.0   |
| className           | 卡片的样式类名                               | string              | -         | 1.21.0   |
| cover               | 卡片封面                                     | ReactNode           | -         | 1.21.0   |
| headerExtraContent  | 卡片标题右侧的额外内容                       | ReactNode           | -         | 1.21.0   |
| footer              | 自定义卡片页脚                               | ReactNode           | -         | 1.21.0   |
| footerLine          | 卡片页脚区与内容区是否有边线                 | boolean             | false     | 1.21.0   |
| footerStyle         | 卡片页脚区内联样式                           | CSSProperties       | -         | 1.21.0   |
| header              | 自定义卡片头部，若传入将覆盖 title 和 headerExtraContent | ReactNode | - | 1.21.0   |
| headerLine          | 卡片标题区与内容区是否有边线                 | boolean             | true      | 1.21.0   |
| headerStyle         | 卡片标题区内联样式                           | CSSProperties       | -         | 1.21.0   |
| loading             | 是否设置加载时的占位                         | boolean             | false     | 1.21.0   |
| shadows             | 设置显示阴影的时机，可选 hover、always       | string              | -         | 1.21.0   |
| style               | 卡片内联样式                                 | CSSProperties       | -         | 1.21.0   |
| title               | 卡片标题                                     | ReactNode           | -         | 1.21.0   |

### CardGroup

| 属性      | 说明                                         | 类型                | 默认值    | 版本     |
|-----------|----------------------------------------------|---------------------|-----------|----------|
| className | 卡片组的样式类名                             | string              | -         | 1.21.0   |
| spacing   | 间距尺寸，支持数值或数组                     | number \| number[]  | 12px      | 1.21.0   |
| style     | 卡片组的内联样式                             | CSSProperties       | -         | 1.21.0   |
| type      | 网格型卡片组，可选值：grid                   | string              | -         | 1.21.0   |

### Card.Meta

| 属性        | 说明         | 类型          | 默认值 | 版本   |
|-------------|--------------|---------------|--------|--------|
| avatar      | 头像         | ReactNode     | -      | 1.21.0 |
| className   | 类名         | string        | -      | 1.21.0 |
| description | 描述         | ReactNode     | -      | 1.21.0 |
| style       | 内联样式     | CSSProperties | -      | 1.21.0 |
| title       | 标题         | ReactNode     | -      | 1.21.0 |

### 设计变量

| 变量                        | 默认值                    | 用法               |
|-----------------------------|---------------------------|--------------------|
| $color-card-bg-default      | var(--semi-color-bg-0)    | 卡片背景颜色       |
| $color-card-border          | var(--semi-color-border)  | 卡片描边颜色       |
| $color-card_title-text      | var(--semi-color-text-0)  | 卡片标题文字颜色   |
| $color-card_description-text| var(--semi-color-text-2)  | 卡片描述文字颜色   |
| $color-card_extra-text      | var(--semi-color-text-0)  | 卡片附加文字颜色   |
| $color-card_body-text       | var(--semi-color-text-1)  | 卡片正文文字颜色   |

### 文案规范

- 卡片标题应具有信息描述性，聚焦最重要的信息，限制在 1 个短语或句段中。
- 正文可操作的：使用祈使句而不是“你可以”来描述正文。
