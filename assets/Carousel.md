# Carousel 轮播图

## 属性

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| --- | --- | --- | --- | --- |
| activeIndex | 受控属性 | number | - | 2.10.0 |
| animation | 切换动画，可选值：fade，slide | "fade" \| "slide" | "slide" | 2.10.0 |
| arrowProps | 箭头参数，用于自定义箭头样式和点击事件 | () => { leftArrow: ArrowButton, rightArrow: ArrowButton } | - | 2.10.0 |
| autoPlay | 是否自动循环展示，或传入参数对象 | boolean \| { interval?: number, hoverToPause?: boolean } | true | 2.10.0 |
| className | 样式类名 | string | - | 2.10.0 |
| defaultActiveIndex | 初始化时默认展示的索引 | number | 0 | 2.10.0 |
| indicatorPosition | 指示器位置：left、center、right | "left" \| "center" \| "right" | "center" | 2.10.0 |
| indicatorSize | 指示器尺寸：small、medium | "small" \| "medium" | "small" | 2.10.0 |
| indicatorType | 指示器类型：dot、line、columnar | "dot" \| "line" \| "columnar" | "dot" | 2.10.0 |
| theme | 指示器和箭头主题：primary、light、dark | "primary" \| "light" \| "dark" | "light" | 2.10.0 |
| onChange | 图片切换时的回调 | (index: number, preIndex: number) => void | - | 2.10.0 |
| arrowType | 箭头展示时机：hover、always | "hover" \| "always" | always | 2.10.0 |
| showArrow | 是否展示箭头 | boolean | true | 2.10.0 |
| showIndicator | 是否展示指示器 | boolean | true | 2.10.0 |
| slideDirection | slide 动画方向：left、right | "left" \| "right" | "left" | 2.10.0 |
| speed | 切换速度，单位 ms | number | 300 | 2.10.0 |
| style | 内联样式 | CSSProperties | - | 2.10.0 |
| trigger | 指示器触发时机：hover、click | "hover" \| "click" | - | 2.10.0 |

### ArrowButton

| 属性 | 说明 | 类型 | 默认值 | 版本 |
| --- | --- | --- | --- | --- |
| props | 箭头 div 上的参数，包括 style, onClick 等 | React.DetailedHTMLProps<React.HTMLAttributes<HTMLDivElement>, HTMLDivElement> | - | 2.10.0 |
| children | 箭头自定义 Icon | React.ReactNode | - | 2.10.0 |

## Methods

| 方法 | 说明 | 版本 |
| --- | --- | --- |
| play() | 播放 | 2.10.0 |
| stop() | 停止播放 | 2.10.0 |
| goTo(targetIndex) | 切换到指定位置 | 2.10.0 |
| prev() | 切换到上一位置 | 2.10.0 |
| next() | 切换到下一位置 | 2.10.0 |

## 设计变量

| 变量 | 默认值 | 用法 |
| --- | --- | --- |
| $transition_duration_carousel_indicator-bg | var(--semi-transition_delay-none) | 指示器-背景色-动画持续时间 |
| $transition_function_carousel_indicator-bg | var(--semi-transition_function-easeOut) | 指示器-背景色-过渡曲线 |
| $transition_delay_carousel_indicator-bg | var(--semi-transition_delay-none) | 指示器-背景色-延迟时间 |
| $transition_duration_carousel_arrow-bg | var(--semi-transition_delay-none) | 箭头-背景色-动画持续时间 |
| $transition_funciton_carousel_arrow-bg | var(--semi-transition_function-easeOut) | 箭头-背景色-过渡曲线 |
| $transition_delay_carousel_arrow-bg | var(--semi-transition_delay-none) | 箭头-背景色-延迟时间 |
