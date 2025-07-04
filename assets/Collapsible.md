# Collapsible 折叠

## API 参考

| 属性           | 说明                                                         | 类型                    | 默认值    | 版本      |
|----------------|--------------------------------------------------------------|-------------------------|-----------|-----------|
| className      | 类名                                                         | string                  | -         | 0.34.0    |
| collapseHeight | 折叠高度                                                     | number                  | 0         | 1.0.0     |
| duration       | 动画执行的时间                                               | number                  | 250       | -         |
| fade           | 是否开启淡入淡出                                             | boolean                 | false     | 2.21.0    |
| isOpen         | 是否展开内容区域                                             | boolean                 | false     | -         |
| keepDOM        | 是否保留隐藏的面板 DOM 树，默认销毁                          | boolean                 | false     | 0.25.0    |
| lazyRender     | 配合 keepDOM 使用，为 true 时挂载时不会渲染组件              | boolean                 | false     | 2.54.0    |
| motion         | 是否开启动画                                                 | boolean                 | true      | -         |
| onMotionEnd    | 动画结束的回调                                               | () => void              | -         | -         |
| reCalcKey      | 当 reCalcKey 改变时，将重新计算子节点的高度                  | number \| string        | -         | 1.5.0     |
| style          | 样式                                                         | CSSProperties           | -         | 0.34.0    |
| id             | id，html id string type                                      | string                  | -         | 2.3.0     |

### Accessibility

- 支持 `id` 属性，可配合 `aria-controls` 实现可访问性。

