# Image 图片

## 组件说明
用于展示和预览图片。

## Image API

| 属性              | 说明                                         | 类型                                   | 默认值   | 版本      |
|------------------|--------------------------------------------|----------------------------------------|----------|-----------|
| alt              | 图像描述                                     | string                                 | -        |           |
| className        | 自定义样式类名                               | string                                 | -        |           |
| crossOrigin      | 透传给原生 img 标签的 crossorigin            | 'anonymous' \| 'use-credentials'       | -        |           |
| fallback         | 加载失败容错地址或自定义加载失败时的显示内容   | ReactNode                              | -        |           |
| height           | 图片显示高度                                 | number                                 | -        |           |
| imgCls           | 自定义样式类名，透传给 img 节点              | string                                 | -        |           |
| imgStyle         | 自定义样式，透传给 img 节点                  | CSSProperties                          | -        |           |
| onClick          | 点击图片的回调                               | (event: any) => void                   | -        |           |
| onError          | 加载错误回调                                 | (event: Event) => void                 | -        |           |
| onLoad           | 加载成功回调                                 | (event: Event) => void                 | -        |           |
| placeholder      | 图片未加载时的占位内容                       | ReactNode                              | -        |           |
| preview          | 预览参数，为 false 时禁用预览                | boolean \| ImagePreview                | -        |           |
| src              | 图片获取地址                                 | string                                 | -        |           |
| style            | 自定义样式                                   | CSSProperties                          | -        |           |
| width            | 图片显示宽度                                 | number                                 | -        |           |
| setDownloadName  | 设置图片下载名称                             | (src: string) => string                | -        | 2.40.0    |

> 其他支持的属性同 img，其他属性将透传至底层的 img 节点。

## ImagePreview API

| 属性                | 说明                                         | 类型                                   | 默认值       | 版本      |
|--------------------|--------------------------------------------|----------------------------------------|--------------|-----------|
| adaptiveTip        | 适应页面操作按钮提示                         | string                                 | "适应页面"   |           |
| className          | 自定义样式类名                               | string                                 | -            |           |
| closable           | 是否显示关闭按钮                             | boolean                                | true         |           |
| closeOnEsc         | 点击 esc 关闭预览                            | boolean                                | true         |           |
| crossOrigin        | 透传给预览图片的原生 img 标签的 crossorigin  | 'anonymous' \| 'use-credentials'       | -            |           |
| currentIndex       | 受控属性，当前预览图片下标                   | number                                 | -            |           |
| defaultCurrentIndex| 首次展示图片下标                             | number                                 | -            |           |
| defaultVisible     | 首次是否开启预览                             | boolean                                | -            |           |
| disableDownload    | 禁用下载                                     | boolean                                | false        |           |
| downloadTip        | 下载操作按钮提示                             | string                                 | "下载"       |           |
| getPopupContainer  | 指定父级 DOM                                 | () => HTMLElement                      | () => document.body |      |
| infinite           | 是否无限循环                                 | boolean                                | false        |           |
| lazyLoad           | 是否开启懒加载                               | boolean                                | true         |           |
| lazyLoadMargin     | Intersection Observer API 的 rootMargin 参数  | string                                 | "0px 100px 100px 0px" |    |
| maskClosable       | 点击遮罩是否可关闭                           | boolean                                | true         |           |
| nextTip            | 下一步操作按钮提示                           | string                                 | "下一步"     |           |
| originTip          | 原始尺寸操作按钮提示                         | string                                 | "原始尺寸"   |           |
| onChange           | 切换图片触发的事件                           | (index: number) => void                | -            |           |
| onClose            | 点击关闭按钮时的回调函数                     | () => void                             | -            |           |
| onDownload         | 图片下载回调函数                             | (src: string, index: number) => void   | -            |           |
| onDownloadError    | 图片下载错误回调函数                         | (src: string) => void                  | -            | v2.54.0   |
| onRotateLeft       | 旋转图片的回调                               | (angle: number) => void                | -            |           |
| onNext             | 向后切换图片的回调                           | (index: number) => void                | -            |           |
| onPrev             | 向前切换图片的回调                           | (index: number) => void                | -            |           |
| onZoomIn           | 图片放大时的回调函数                         | (zoom: number) => void                 | -            |           |
| onZoomOut          | 图片缩小时的回调函数                         | (zoom: number) => void                 | -            |           |
| onVisibleChange    | 切换可见状态触发的回调                       | (visible: boolean) => void             | -            |           |
| preLoad            | 是否开启预加载                               | boolean                                | true         |           |
| preLoadGap         | 预加载的步长                                 | number                                 | 2            |           |
| previewTitle       | 自定义预览 title                             | ReactNode                              | -            |           |
| previewCls         | 自定义预览样式类名                           | string                                 | -            |           |
| previewStyle       | 自定义预览样式                               | object                                 | -            |           |
| prevTip            | 上一步操作按钮提示                           | string                                 | "上一步"     |           |
| renderHeader       | 自定义渲染预览顶部信息                       | (info: ReactNode) => ReactNode         | -            |           |
| renderPreviewMenu  | 自定义渲染预览底部菜单信息                   | (props: MenuProps) => ReactNode        | -            |           |
| rotateTip          | 旋转操作按钮提示                             | string                                 | "旋转"       |           |
| showTooltip        | 是否展示底部操作区提示                       | boolean                                | false        |           |
| src                | 图片列表信息                                 | string \| string[]                     | -            |           |
| style              | 自定义样式                                   | CSSProperties                          | -            |           |
| viewerVisibleDelay | 隐藏预览操作按钮前的无操作时长               | number                                 | 10000        |           |
| visible            | 受控属性，是否预览                           | boolean                                | -            |           |
| zIndex             | 预览层层级                                   | number                                 | 1070         |           |
| zoomInTip          | 放大操作按钮提示                             | string                                 | "放大"       |           |
| zoomOutTip         | 缩小操作按钮提示                             | string                                 | "缩小"       |           |
| zoomStep           | 图片每次缩小/放大比例                        | number                                 | 0.1          |           |
| setDownloadName    | 设置图片下载名称                             | (src: string) => string                | -            | 2.40.0    |

### MenuProps

| 属性           | 说明                         | 类型                        | 版本      |
|---------------|----------------------------|-----------------------------|-----------|
| curPage       | 当前图片页下标               | number                      |           |
| disabledPrev  | 是否禁用向左切换按钮         | boolean                     |           |
| disabledNext  | 是否禁用向右切换按钮         | boolean                     |           |
| disableDownload| 是否禁用下载按钮            | boolean                     |           |
| max           | 图片缩放最大比例             | number                      |           |
| min           | 图片缩放最小比例             | number                      |           |
| onDownload    | 图片下载的调用函数           | () => void                  |           |
| onZoomIn      | 图片放大时的调用函数         | () => void                  |           |
| onZoomOut     | 图片缩小时的调用函数         | () => void                  |           |
| onPrev        | 向前切换图片的调用函数       | () => void                  |           |
| onNext        | 向后切换图片的调用函数       | () => void                  |           |
| onRotateLeft  | 逆时针旋转图片的调用函数     | () => void                  |           |
| onRotateRight | 顺时针旋转图片的调用函数     | () => void                  |           |
| ratio         | 原始尺寸或适应页面按钮状态   | "adaptation" \| "realSize"  |           |
| step          | 缩放的比例步长               | number                      |           |
| totalNum      | 可预览的总图片数             | number                      |           |
| zoom          | 当前图片缩放比例             | number                      |           |
| menuItems     | 默认底部预览操作区域功能按钮 | ReactNode[]                 | 2.40.0    |

## 设计变量

| 变量                                   | 默认值      | 用法                   |
|----------------------------------------|-------------|------------------------|
| $transition_duration-image_preview_image_img | 300ms      | 预览图像动画持续时间   |
| $transition_function-image_preview_image_img | cubic-bezier(0.215, 0.61, 0.355, 1) | 预览图片动画函数 |
| $transition_delay-image_preview_image_img    | 0s         | 预览图片延迟时间       |
| $transition_duration-image_preview           | 500ms      | 预览图片透明度变化动画时间 |
