---
name: tui-theme
description: 在 hongmengchen 的命令行工具里使用 github.com/hongmengchen/tui/theme 共享配色和卡片样式；改配色、调整 TUI 样式时读它。
---

# tui theme

`github.com/hongmengchen/tui` 是 kube-token / npm-token / hmck 共用的终端样式库。**不要**在各工具仓库里重新定义配色或 `card`、`truncate`、`hints` 这类原语。

## 用法

```go
import "github.com/hongmengchen/tui/theme"

var styles = theme.Default // 中性 Slate，默认

styles.Title.Render("my-tool")
styles.Good.Render("✓ 有效")
styles.CardOn.Render(body)   // 选中卡片
styles.CardOff.Render(body)  // 未选中
styles.Hints("↑↓", "选择", "q", "退出")
theme.Truncate(s, width)
theme.Pad(s, width)
```

配色预设：`theme.Slate()`（默认）、`theme.Purple()`、`theme.Mono()`；自定义用 `theme.New(palette)`。

## 规矩

- 只把**跨工具共用**的样式放进本仓库；某个工具独有的布局/组件留在它自己仓库。
- 改颜色或样式语义时，同时更新 `theme_test.go`，打新 tag（如 `v0.1.1`），再去三个工具里 `go get` + 发版。
- 不要在工具仓库里复制调色板来“临时改一下”，需要新预设就在本仓库加。
- 中文截断一律用 `theme.Truncate`，不要用字节截断或 `runewidth` 自己算。

## 开发

```bash
go test ./...
```
