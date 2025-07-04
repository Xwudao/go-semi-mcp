# Spin 加载器

加载器组件用于告知用户内容正在加载且需要一段不确定的时长。

## API 参考

| 属性               | 说明                                         | 类型             | 默认值   |
|--------------------|----------------------------------------------|------------------|----------|
| childStyle         | 内部子元素的样式 v>=1.0.0                    | CSSProperties    | -        |
| delay              | 延迟显示加载效果的时间                        | number(ms)       | 0        |
| indicator          | 加载指示符                                   | ReactNode        | 无       |
| size               | 组件大小，可选值为 small, middle, large       | string           | middle   |
| spinning           | 是否处于加载中的状态                          | boolean          | true     |
| style              | 内联样式                                     | CSSProperties    | -        |
| tip                | 当 spin 作为包裹元素时，可以自定义描述文字    | ReactNode        | 无       |
| wrapperClassName   | 包裹元素的类名                               | string           | 无       |

## 设计变量

| 变量                                   | 默认值   | 用法                         |
|----------------------------------------|----------|------------------------------|
| $animation_duration-spin_wrapper-spin  | 600ms    | 加载图标容器旋转一周时长      |
| $animation_duration-spin_customChildren-spin | 1600ms | 自定义指示器时旋转一周时长    |

## 文案规范

- 准确地说明加载状态，使用比如“Loading”, “Submitting”, “Processing”等词
- 使用尽量少的词汇去描述状态

## FAQ

**怎么修改 icon 的颜色？**  
可以通过给 `.semi-spin-wrapper` 类添加 color 属性覆盖原有的颜色。
