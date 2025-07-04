# Steps 步骤

## Steps

| 参数         | 说明                                                                 | 类型                        | 默认值    | 版本      |
| ------------ | -------------------------------------------------------------------- | --------------------------- | --------- | --------- |
| className    | 类名                                                                 | string                      |           |           |
| current      | 指定当前步骤，从 0 开始记数。在子 Step 元素中，可通过 status 覆盖状态 | number                      | 0         |           |
| direction    | 指定步骤条方向，horizontal/vertical                                  | string                      | horizontal|           |
| hasLine      | type=basic 时，控制是否显示连接线                                    | boolean                     | true      | 1.18.0    |
| initial      | 起始序号，从 0 开始记数                                              | number                      | 0         |           |
| status       | 当前步骤状态，wait/process/finish/error/warning                      | string                      | process   |           |
| size         | 尺寸，small/default                                                  | string                      | default   | 1.18.0    |
| style        | 样式                                                                 | CSSProperties               |           |           |
| type         | 步骤条类型，fill/basic/nav                                           | string                      | fill      | 1.18.0    |
| onChange     | 步骤切换回调                                                         | (index: number) => void     |           | 1.29.0    |

## Steps.Step

| 参数         | 说明                                   | 类型                                  | 默认值 |
| ------------ | -------------------------------------- | ------------------------------------- | ------ |
| aria-label   | aria-label 描述                        | React.AriaAttributes["aria-label"]    |        |
| className    | 类名                                   | string                                |        |
| description  | 步骤详情描述                           | ReactNode                             | -      |
| icon         | 步骤图标                               | ReactNode                             | -      |
| role         | 容器 role                              | React.AriaRole                        | -      |
| status       | 指定状态，wait/process/finish/error/warning | string                           | wait   |
| style        | 样式                                   | CSSProperties                         |        |
| title        | 标题                                   | ReactNode                             | -      |
| onClick      | 点击回调                               | function                              | -      |
| onKeyDown    | 回车事件回调                           | function                              | -      |
