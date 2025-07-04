# Cropper 图片裁切

通过设定裁切框的宽高比例，自由裁切图片。

## API

| 属性                  | 说明                           | 类型                                      | 默认值                  |
|----------------------|------------------------------|------------------------------------------|-----------------------|
| aspectRatio          | 裁切框比例                      | number                                   | -                     |
| className            | 类名                           | string                                   | -                     |
| cropperBoxClassName  | 裁切框类名                      | string                                   | -                     |
| cropperBoxStyle      | 裁切框样式                      | CSSProperties                            | -                     |
| defaultAspectRatio   | 初始裁切框比例                   | number                                   | 1                     |
| imgProps             | 透传给 img 标签的属性             | object                                   | -                     |
| fill                 | 裁切结果中非图片部分的填充色        | string                                   | 'rgba(0, 0, 0, 0)'    |
| maxZoom              | 最大缩放倍数                      | number                                   | 3                     |
| minZoom              | 最小缩放倍数                      | number                                   | 0.1                   |
| onZoomChange         | 缩放回调                         | (zoom: number) => void                   | -                     |
| preview              | 指定预览容器                      | () => HTMLElement                        | -                     |
| rotate               | 旋转角度                         | number                                   | -                     |
| shape                | 裁切框形状                        | 'rect' \| 'round' \| 'roundRect'         | 'rect'                |
| src                  | 图片地址                         | string                                   | -                     |
| showResizeBox        | 是否展示调整块                    | boolean                                  | true                  |
| style                | 样式                             | CSSProperties                            | -                     |
| zoom                 | 缩放比例                         | number                                   | -                     |
| zoomStep             | 缩放步长                         | number                                   | 0.1                   |

## Methods

| Name              | Description                |
|-------------------|---------------------------|
| getCropperCanvas  | 获取裁剪图片的 canvas      |

## 设计变量

| 变量                          | 默认值                        | 用法                   |
|------------------------------|------------------------------|----------------------|
| $color-cropper_mask-bg       | var(--semi-color-overlay-bg) | 裁切框遮罩背景颜色      |
| $color-cropper_box-outline   | var(--semi-color-primary)    | 裁切框边框颜色          |
| $color-cropper_box_corner-bg | var(--semi-color-primary)    | 裁切框调整块背景色       |
