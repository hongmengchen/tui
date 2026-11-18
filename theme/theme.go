// Package theme 是几个命令行小工具共用的终端配色和基础样式。
//
// 它只放“长什么样”：颜色、标题/正文/语义色、选中卡片、键位提示这些原语。
// 每个工具自己的页面布局、交互和文案仍然留在各自仓库里。
package theme

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Palette 是一套完整的颜色。浅色/深色终端各取一个值，lipgloss 会自己挑。
type Palette struct {
	Accent lipgloss.AdaptiveColor // 标题、选中项、按键
	Dim    lipgloss.AdaptiveColor // 次要说明、路径
	Text   lipgloss.AdaptiveColor // 正文
	Line   lipgloss.AdaptiveColor // 未选中卡片的左侧竖线
	Good   lipgloss.AdaptiveColor // 成功、有效
	Bad    lipgloss.AdaptiveColor // 失败、失效
	Warn   lipgloss.AdaptiveColor // 可更新、需要处理
}

// Slate 是默认的中性灰蓝配色，适合长时间盯着的终端界面。
func Slate() Palette {
	return Palette{
		Accent: lipgloss.AdaptiveColor{Light: "#334155", Dark: "#CBD5E1"},
		Dim:    lipgloss.AdaptiveColor{Light: "#64748B", Dark: "#94A3B8"},
		Text:   lipgloss.AdaptiveColor{Light: "#0F172A", Dark: "#E2E8F0"},
		Line:   lipgloss.AdaptiveColor{Light: "#E2E8F0", Dark: "#334155"},
		Good:   lipgloss.AdaptiveColor{Light: "#047857", Dark: "#34D399"},
		Bad:    lipgloss.AdaptiveColor{Light: "#B91C1C", Dark: "#F87171"},
		Warn:   lipgloss.AdaptiveColor{Light: "#B45309", Dark: "#FBBF24"},
	}
}

// Purple 是以前各工具用过的紫色主题，留着给想换回去的人。
func Purple() Palette {
	return Palette{
		Accent: lipgloss.AdaptiveColor{Light: "#6D28D9", Dark: "#C4B5FD"},
		Dim:    lipgloss.AdaptiveColor{Light: "#6B7280", Dark: "#9CA3AF"},
		Text:   lipgloss.AdaptiveColor{Light: "#1F2937", Dark: "#E5E7EB"},
		Line:   lipgloss.AdaptiveColor{Light: "#DDD6FE", Dark: "#4C1D95"},
		Good:   lipgloss.AdaptiveColor{Light: "#047857", Dark: "#34D399"},
		Bad:    lipgloss.AdaptiveColor{Light: "#B91C1C", Dark: "#F87171"},
		Warn:   lipgloss.AdaptiveColor{Light: "#B45309", Dark: "#FBBF24"},
	}
}

// Mono 只用灰阶，适合把输出贴到不支持颜色的地方。
func Mono() Palette {
	return Palette{
		Accent: lipgloss.AdaptiveColor{Light: "#111827", Dark: "#F3F4F6"},
		Dim:    lipgloss.AdaptiveColor{Light: "#6B7280", Dark: "#9CA3AF"},
		Text:   lipgloss.AdaptiveColor{Light: "#111827", Dark: "#F9FAFB"},
		Line:   lipgloss.AdaptiveColor{Light: "#D1D5DB", Dark: "#4B5563"},
		Good:   lipgloss.AdaptiveColor{Light: "#374151", Dark: "#D1D5DB"},
		Bad:    lipgloss.AdaptiveColor{Light: "#111827", Dark: "#F9FAFB"},
		Warn:   lipgloss.AdaptiveColor{Light: "#4B5563", Dark: "#9CA3AF"},
	}
}

// Styles 是调色板派生出来的一组常用样式。
type Styles struct {
	Title  lipgloss.Style // 界面标题
	Accent lipgloss.Style // 只需要主色时
	Dim    lipgloss.Style // 弱化说明
	Text   lipgloss.Style // 正文
	Good   lipgloss.Style // 成功
	Bad    lipgloss.Style // 失败
	Warn   lipgloss.Style // 提醒
	Key    lipgloss.Style // 快捷键
	Label  lipgloss.Style // 表单标签（固定 9 格宽）
	// CardOn / CardOff 是列表卡片的两种左侧竖线：选中 / 未选中。
	CardOn  lipgloss.Style
	CardOff lipgloss.Style
}

// New 从调色板生成样式。
func New(p Palette) Styles {
	border := lipgloss.NormalBorder()
	return Styles{
		Title:  lipgloss.NewStyle().Bold(true).Foreground(p.Accent),
		Accent: lipgloss.NewStyle().Foreground(p.Accent),
		Dim:    lipgloss.NewStyle().Foreground(p.Dim),
		Text:   lipgloss.NewStyle().Foreground(p.Text),
		Good:   lipgloss.NewStyle().Foreground(p.Good).Bold(true),
		Bad:    lipgloss.NewStyle().Foreground(p.Bad),
		Warn:   lipgloss.NewStyle().Foreground(p.Warn),
		Key:    lipgloss.NewStyle().Foreground(p.Accent).Bold(true),
		Label:  lipgloss.NewStyle().Foreground(p.Dim).Width(9),
		CardOn: lipgloss.NewStyle().
			Border(border, false, false, false, true).
			BorderForeground(p.Accent).
			PaddingLeft(1),
		CardOff: lipgloss.NewStyle().
			Border(border, false, false, false, true).
			BorderForeground(p.Line).
			PaddingLeft(1),
	}
}

// Default 是各工具默认使用的样式（Slate 配色）。
var Default = New(Slate())

// Truncate 按终端显示宽度截断，补一个省略号。中文和 emoji 都不会算错。
func Truncate(s string, w int) string {
	if w <= 0 || lipgloss.Width(s) <= w {
		return s
	}
	r := []rune(s)
	for len(r) > 0 && lipgloss.Width(string(r)) > w-1 {
		r = r[:len(r)-1]
	}
	return string(r) + "…"
}

// Pad 用空格把 s 补齐到显示宽度 w；已经够宽就原样返回。
func Pad(s string, w int) string {
	if n := lipgloss.Width(s); n < w {
		return s + strings.Repeat(" ", w-n)
	}
	return s
}

// Hints 把成对的「键、说明」渲染成一行底部提示。
// 例：styles.Hints("↑↓", "选择", "q", "退出")
func (s Styles) Hints(pairs ...string) string {
	var parts []string
	for i := 0; i+1 < len(pairs); i += 2 {
		parts = append(parts, s.Key.Render(pairs[i])+s.Dim.Render(" "+pairs[i+1]))
	}
	return "  " + strings.Join(parts, s.Dim.Render("  ·  "))
}
