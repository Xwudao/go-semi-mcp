# Progress 进度条

用于展示用户操作的当前进度和状态，一般在操作耗时较长时使用。也可用来表示任务/对象的完成度。

## API 参考

| 属性             | 说明                                                                 | 类型                                               | 默认值                           |
|------------------|----------------------------------------------------------------------|----------------------------------------------------|----------------------------------|
| aria-label       | aria-label属性，用来给当前元素加上的标签描述, 用于提升可访问性         | string                                             |                                  |
| aria-labelledby  | aria-labelledby属性，表明某些元素的 id 是当前元素的标签                | string                                             |                                  |
| aria-valuetext   | aria-valuetext属性，用于提升可访问性                                  | string                                             |                                  |
| className        | 样式类名                                                              | string                                             |                                  |
| direction        | 条状进度条方向 horizontal、vertical                                   | string                                             | 'horizontal'                     |
| format           | 格式化函数，入参为当前百分比，return 的结果将会直接渲染在圆形进度条中心 | (percent: number) => ReactNode                     | (percent) => percent + '%'       |
| id               | id 标识                                                               | string                                             |                                  |
| orbitStroke      | 进度条轨道填充色                                                      | string                                             | 'var(--semi-color-fill-0)'       |
| percent          | 进度百分比                                                            | number                                             |                                  |
| showInfo         | 是否显示文本                                                          | boolean                                            | false                            |
| size             | 尺寸, 可选 default、small、large                                      | string                                             | 'default'                        |
| stroke           | 进度条填充色，支持数组                                                | string \| Array<{percent:number; color:string }>   | 'var(--semi-color-success)'      |
| strokeGradient   | 是否自动生成渐变色补齐区间颜色                                        | boolean                                            | false                            |
| strokeLinecap    | 圆角round/方角square(仅 type='circle'模式下生效)                      | string                                             | 'round'                          |
| strokeWidth      | type 为circle时，该属性控制进度条宽度                                  | number                                             | 4                                |
| style            | 样式                                                                  | CSSProperties                                      |                                  |
| type             | 类型，可选line、circle                                                | string                                             | 'line'                           |
| width            | 环形进度条宽度                                                        | number                                             | size='default'为72，'small'为24  |

### Accessibility

- Progress 具有 progressbar role。
- 自动设置 aria-valuenow。
- 支持 aria-label、aria-labelledby、aria-valuetext。

### 设计变量

| 变量                              | 默认值                        | 用法                 |
|-----------------------------------|-------------------------------|----------------------|
| $color-progress_default-bg        | var(--semi-color-fill-0)      | 默认背景色           |
| $color-progress_track_inner-bg    | var(--semi-color-success)     | 默认进度色           |
| $color-progress_line_text-text    | var(--semi-color-text-0)      | 百分比文本文字颜色   |
| $color-progress_circle-text       | var(--semi-color-mode-minor-text) | 暂无             |
