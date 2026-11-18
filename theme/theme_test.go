package theme

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestTruncate(t *testing.T) {
	tests := []struct {
		in   string
		w    int
		want string
	}{
		{"short", 10, "short"},
		{"0123456789", 10, "0123456789"},
		{"0123456789abc", 10, "012345678…"},
		{"中文也不能按字节截断", 8, "中文也…"},
		{"abc", 0, "abc"},
	}
	for _, tt := range tests {
		if got := Truncate(tt.in, tt.w); got != tt.want {
			t.Errorf("Truncate(%q, %d) = %q, want %q", tt.in, tt.w, got, tt.want)
		}
		if lipgloss.Width(Truncate(tt.in, tt.w)) > tt.w && tt.w > 0 {
			t.Errorf("Truncate(%q, %d) 超宽", tt.in, tt.w)
		}
	}
}

func TestPad(t *testing.T) {
	if got := Pad("abc", 5); got != "abc  " {
		t.Errorf("Pad = %q", got)
	}
	if got := Pad("中文", 5); lipgloss.Width(got) != 5 {
		t.Errorf("Pad 中文宽度 = %d, want 5", lipgloss.Width(got))
	}
	if got := Pad("abcdef", 3); got != "abcdef" {
		t.Errorf("已够宽不应截断: %q", got)
	}
}

func TestHints(t *testing.T) {
	out := Default.Hints("↑↓", "选择", "q", "退出")
	for _, want := range []string{"↑↓", "选择", "q", "退出", "·"} {
		if !strings.Contains(out, want) {
			t.Errorf("Hints 里缺少 %q: %q", want, out)
		}
	}
}

func TestPresets(t *testing.T) {
	for name, p := range map[string]Palette{"slate": Slate(), "purple": Purple(), "mono": Mono()} {
		s := New(p)
		if s.Title.Render("x") == "" || s.CardOn.Render("x") == "" {
			t.Errorf("%s 预设样式为空", name)
		}
	}
}
