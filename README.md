# tui

几个命令行小工具（[kube-token](https://github.com/hongmengchen/kube-token)、[npm-token](https://github.com/hongmengchen/npm-token)、[hmck](https://github.com/hongmengchen/hmck)）共用的终端样式。

只放样式原语：配色、标题/正文/语义色、选中卡片、键位提示、按显示宽度截断和补齐。页面布局和交互留在各工具自己的仓库里。

## 用法

```bash
go get github.com/hongmengchen/tui@v0.1.0
```

```go
import "github.com/hongmengchen/tui/theme"

var styles = theme.Default // Slate 配色

func View() string {
	return styles.Title.Render("my-tool") + styles.Dim.Render("  说明") + "\n" +
		styles.CardOn.Render("选中的一行") + "\n" +
		styles.Hints("↑↓", "选择", "q", "退出")
}
```

### 自带配色

| 函数 | 说明 |
| --- | --- |
| `theme.Slate()` | 中性灰蓝，默认 |
| `theme.Purple()` | 以前的紫色主题 |
| `theme.Mono()` | 只有灰阶 |

换一套配色不用改工具代码里的任何样式名：

```go
var styles = theme.New(theme.Purple())
```

### API

- `theme.Default`：`Styles`，包含 `Title / Accent / Dim / Text / Good / Bad / Warn / Key / Label / CardOn / CardOff`，都是 `lipgloss.Style`。
- `theme.New(p Palette) Styles`：从自定义调色板生成样式。
- `theme.Truncate(s string, w int) string`：按终端显示宽度截断，中文不会算错。
- `theme.Pad(s string, w int) string`：按显示宽度补空格。
- `styles.Hints("键", "说明", ...) string`：底部键位提示。

## 约定

- 这里只放跨工具共用的样式。某个工具独有的组件（进度条、输入框布局等）留在它自己那边。
- 改配色属于公共变更：改完发新 tag，各工具升级依赖后重新发版。

## 开发

```bash
go test ./...
```
