# PinCode 验证码输入

用于便捷直观地输入验证码。

## API 参考

| 属性          | 说明                                 | 类型                                                      | 默认值     |
| ------------- | ------------------------------------ | --------------------------------------------------------- | ---------- |
| autoFocus     | 是否自动聚焦到第一个元素             | boolean                                                   | true       |
| className     | 类名                                 | string                                                    | -          |
| count         | 验证码位数                           | number                                                    | 6          |
| defaultValue  | 输入框内容默认值                     | string                                                    | -          |
| disabled      | 禁用                                 | boolean                                                   | false      |
| format        | 验证码单个字符格式限制               | 'number' \| 'mixed' \| RegExp \| (char: string) => boolean| 'number'   |
| size          | 输入框大小，large、default、small    | string                                                    | 'default'  |
| style         | 样式                                 | object                                                    | -          |
| value         | 输入框内容                           | string                                                    | -          |
| onChange      | 输入回调                             | (value: string) => void                                   | -          |
| onComplete    | 验证码所有位数输入完毕回调           | (value: string) => void                                   | -          |

## Methods

通过 ref 调用实现特殊交互：

| 方法   | 说明                 |
| ------ | -------------------- |
| focus  | 聚焦，入参为验证码第几位 |
| blur   | 移出焦点，入参为验证码第几位 |
