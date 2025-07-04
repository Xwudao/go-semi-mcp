# Icon 图标

## 如何引入

```js
import Icon, { IconHome } from '@douyinfe/semi-icons';
import { IconAvatar, IconCard } from '@douyinfe/semi-icons-lab';
```

## 基础使用

```js
import React from 'react';
import { IconHome } from '@douyinfe/semi-icons';

() => <IconHome />;
```

## 旋转

```js
import React from 'react';
import { IconHome, IconEmoji, IconSpin } from '@douyinfe/semi-icons';

() => (
    <div>
        <IconHome size="small" />
        <IconEmoji rotate={180} />
        <IconSpin spin />
    </div>
);
```

## 尺寸

Icon 组件封装了 size 属性，支持 `extra-small` (8x8)，`small` (12x12)，`default` (16x16)，`large` (20x20)，`extra-large` (24x24)，`inherit` 继承字体大小。

## 颜色

图标会自动继承外部容器 CSS 的 color 属性，也可通过 style props 修改颜色。

```js
import React from 'react';
import { IconLikeHeart, IconFlag } from '@douyinfe/semi-icons';

() => (
    <div style={{ color: '#E91E63' }}>
        <IconLikeHeart size="extra-large"/>
        <IconFlag size="extra-large"/>
    </div>
);
```

## 自定义图标

可传入自定义 svg。

```js
import React from 'react';
import { Icon } from '@douyinfe/semi-ui';

() => {
    function CustomIcon() {
        return <svg width="1em" height="1em" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <circle cx="12" cy="12" r="11" fill="#FBCD2C"/>
        </svg>;
    }
    return <Icon svg={<CustomIcon />} />;
};
```

## 使用 svgr 转 ReactComponent

```js
// webpack.config.js
{
  test: /\.svg$/,
  use: ['@svgr/webpack'],
}

import { Icon } from '@douyinfe/semi-ui';
import StarIcon from './star.svg';

<Icon svg={<StarIcon />} />
```

## API 参考

| 属性          | 说明                       | 类型                                  | 默认值      |
|---------------|----------------------------|---------------------------------------|-------------|
| className     | 类名                       | string                                | 无          |
| onClick       | 单击图标回调               | (e: Event) => void                    | 无          |
| onMouseDown   | 鼠标按钮按下回调 >=v1.21   | (e: Event) => void                    | 无          |
| onMouseEnter  | 鼠标进入回调               | (e: Event) => void                    | 无          |
| onMouseLeave  | 鼠标离开回调               | (e: Event) => void                    | 无          |
| onMouseMove   | 鼠标移动回调 >=v1.21       | (e: Event) => void                    | 无          |
| onMouseUp     | 鼠标按钮抬起回调 >=v1.21   | (e: Event) => void                    | 无          |
| rotate        | 旋转度数                   | number                                |             |
| size          | 尺寸                       | string (`inherit`/`extra-small`/`small`/`default`/`large`/`extra-large`) | default     |
| spin          | 旋转动画                   | boolean                               |             |
| style         | 图标样式                   | CSSProperties                         | 无          |
| svg           | 图标内容                   | ReactNode                             | 无          |

## Accessibility

- Icon 组件 role 为 `img`，aria-label 默认为组件文件名，可自定义。
- 内部 svg 元素默认 aria-hidden。
