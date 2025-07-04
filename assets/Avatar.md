# Avatar 头像

头像，支持图片或字符展示。

## API 参考

### Avatar

| 属性         | 说明 | 类型 | 默认值 |
| ------------ | ---- | ---- | ------ |
| alt          | 图像的替代文本描述 | string | - |
| border       | 额外边框（>=2.52.0） | `{color?:string, motion?:boolean}` or boolean | - |
| bottomSlot   | 底部 Slot 配置（>=2.52.0） | `{render?: () => React.ReactNode, shape?: "circle" or "square", text: React.ReactNode, bgColor:string, textColor:string, className:string, style?:CSSProperties}` | - |
| className    | 类名 | string | - |
| color        | 指定头像的颜色，支持 amber、blue、cyan、green、grey、indigo、light-blue、light-green、lime、orange、pink、purple、red、teal、violet、yellow | string | grey |
| contentMotion| 头像内容区域动效（>=2.xx.0） | boolean | - |
| hoverMask    | hover 时头像内容覆盖层 | ReactNode | - |
| gap          | 字符头像距离左右两侧的像素大小 | number | 3 |
| imgAttr      | 原生 img 属性（>=1.5.0） | React.ImgHTMLAttributes<HTMLImageElement> | - |
| shape        | 指定头像的形状，支持 circle、square | string | circle |
| size         | 设置头像的大小，支持 extra-extra-small、extra-small、small、default、medium、large、extra-large 和合法的 width 属性值如 "10px" | string | medium |
| src          | 图片类头像的资源地址 | string | - |
| srcSet       | 设置图片类头像响应式资源地址 | string | - |
| style        | 样式名 | CSSProperties | - |
| topSlot      | 顶部 Slot 配置（>=2.52.0） | `{render?: () => React.ReactNode, gradientStart?: string, gradientEnd?: string, text: React.ReactNode, textColor:string, className:string, style?:CSSProperties}` | - |
| onClick      | 单击头像的回调 | (e: Event) => void | - |
| onError      | 图片加载失败事件，返回 false 会关闭组件默认的 fallback 行为 | (e: Event) => boolean | - |
| onMouseEnter | MouseEnter 事件的回调 | (e: Event) => void | - |
| onMouseLeave | MouseLeave 事件的回调 | (e: Event) => void | - |

### AvatarGroup

| 属性         | 说明 | 类型 | 默认值 |
| ------------ | ---- | ---- | ------ |
| maxCount     | 最大数量限制，超出后显示+N | number | - |
| overlapFrom  | 设置头像覆盖方向，支持 start, end | string | start |
| renderMore   | 自定义渲染 more 标签 | (restNumber: number, restAvatars: ReactNode[]) => ReactNode | - |
| shape        | 指定头像的形状，支持 circle、square | string | circle |
| size         | 设置头像的大小，支持 extra-extra-small, extra-small、small、default、medium、large、extra-large 和合法的 width 属性值 | string | medium |

## Accessibility

- Avatar 一般不用于操作，不需要被获取焦点。但当 Avatar 可以被点击操作时需要被聚焦，并响应键盘 Enter 事件。
- Avatar 的 alt 属性可以被屏幕阅读器读取，使用头像组件时，请使用 alt 属性解释头像的内容。

## 设计变量

| 变量 | 默认值 | 用法 |
| ---- | ------ | ---- |
| $animation_duration-additionalBorder | 800ms | 边框动画持续时间 |
| $animation_duration-content | 1000ms | 内容动画持续时间 |
| $animation_opacity-additionalBorder-start | 1 | 边框动画起始透明度 |
| $animation_opacity-additionalBorder-end | 0 | 边框动画结束透明度 |
| $animation_width-additionalBorder-end | 0 | 边框动画结束宽度 |
| $animation_scale-additionalBorder-start | 1 | 边框动画起始缩放比例 |
| $animation_scale-additionalBorder-end | 1.15 | 边框动画结束缩放比例 |
| $animation_scale-content-start | 1 | 边框动画起始缩放比例 |
| $animation_scale-content-middle | 0.9 | 边框动画中间态缩放比例 |
| $animation_scale-content-end | 1 | 边框动画结束缩放比例 |
