# UserGuide 用户引导

## API 参考

| 属性              | 说明                                                         | 类型                                  | 默认值      |
|-------------------|--------------------------------------------------------------|---------------------------------------|-------------|
| className         | 自定义类名                                                   | string                                | -           |
| current           | 当前步骤的索引                                               | number                                | 0           |
| finishText        | 最后一步完成按钮的文本                                       | string                                | '完成'      |
| mask              | 是否显示蒙层                                                 | boolean                               | true        |
| mode              | 引导模式，可选值：popup（气泡卡片）或 modal（弹窗式）        | string                                | popup       |
| nextButtonProps   | 下一步按钮的属性                                             | ButtonProps                           | {}          |
| onChange          | 步骤改变时的回调                                             | function(current: number)             | () => void  |
| onFinish          | 完成所有步骤时的回调                                         | function()                            | () => void  |
| onNext            | 点击下一步按钮时的回调                                       | function(current: number)             | () => void  |
| onPrev            | 点击上一步按钮时的回调                                       | function(current: number)             | () => void  |
| onSkip            | 点击跳过按钮时的回调                                         | function()                            | () => void  |
| position          | 弹出层相对于目标元素的位置                                   | string                                | bottom      |
| prevButtonProps   | 上一步按钮的属性                                             | ButtonProps                           | {}          |
| showPrevButton    | 是否显示上一步按钮                                           | boolean                               | true        |
| showSkipButton    | 是否显示跳过按钮                                             | boolean                               | true        |
| spotlightPadding  | 高亮区域的内边距，单位为像素                                 | number                                | -           |
| steps             | 引导步骤配置，必填                                           | StepItem[]                            | []          |
| style             | 自定义样式                                                   | React.CSSProperties                   | -           |
| theme             | 主题样式，可选值：default 或 primary                         | string                                | default     |
| visible           | 是否显示引导                                                 | boolean                               | false       |
| getPopupContainer | 指定父级 DOM，弹层将会渲染至该 DOM 中                        | () => HTMLElement                     | -           |
| zIndex            | 弹层层级                                                     | number                                | 1030        |

### Steps.Step

| 属性             | 说明                                   | 类型                                  | 默认值  |
|------------------|----------------------------------------|---------------------------------------|---------|
| className        | 步骤的自定义类名                       | string                                | -       |
| cover            | 步骤的封面图                           | ReactNode                             | -       |
| target           | 目标元素，高亮区域会聚焦到这个元素上   | (() => Element) \| Element            | -       |
| title            | 步骤标题                               | string \| ReactNode                   | -       |
| description      | 步骤描述                               | ReactNode                             | -       |
| mask             | 是否显示此步骤的蒙层，会覆盖全局配置   | boolean                               | -       |
| showArrow        | 是否显示箭头（仅在 mode=popup 时有效） | boolean                               | true    |
| spotlightPadding | 此步骤高亮区域区域的内边距             | number                                | -       |
| theme            | 此步骤的主题，会覆盖全局配置           | default \| primary                    | -       |
| position         | 此步骤弹出层的位置，会覆盖全局配置     | Position                              | -       |
