# Skeleton 骨架屏

在需要等待加载内容的位置提供的占位组件。

## Skeleton

| 属性        | 说明                   | 类型          | 默认值   |
|-------------|------------------------|---------------|----------|
| active      | 是否展示动画效果       | boolean       | false    |
| className   | 类名                   | string        | -        |
| loading     | 是否显示占位元素       | boolean       | true     |
| placeholder | 加载等待时的占位元素   | ReactNode     | -        |
| style       | 样式                   | CSSProperties | -        |

## Skeleton.Avatar / Skeleton.Image / Skeleton.Title / Skeleton.Button

> Skeleton.Image、Skeleton.Title、Skeleton.Button 大部分API 与 Skeleton.Avatar 相同。其中 shape 仅 Skeleton.Avatar支持

| 属性      | 说明           | 类型    | 默认值   |
|-----------|----------------|---------|----------|
| className | 类名           | string  | -        |
| size      | 头像大小       | string  | medium   |
| style     | 样式           | CSSProperties | -    |
| shape     | 头像形状（仅Avatar） | string  | circle   |

## Skeleton.Paragraph

| 属性      | 说明           | 类型    | 默认值   |
|-----------|----------------|---------|----------|
| className | 类名           | string  | -        |
| rows      | 行数           | number  | 4        |
| style     | 样式           | CSSProperties | -    |

### 设计变量

| 变量                                 | 默认值      | 用法                   |
|--------------------------------------|-------------|------------------------|
| $animation_duration-skeleton-highlight | 1400ms    | 骨架屏高亮动画时长     |
| $animation_function-skeleton-highlight | ease      | 暂无                   |
