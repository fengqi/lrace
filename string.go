package lrace

import "strings"

// StringSplitWith 以指定的多个字符串切割string，且跳过部分特殊字符
// 灵感来源：strings.FieldsFunc()
func StringSplitWith(s string, sep []string, exclude []string) []string {
	type span struct {
		start int
		end   int
	}

	spans := make([]span, 0, 32)
	start := -1
	for end, item := range s {
		flag := false
		if InArray(sep, string(item)) {
			flag = true
		}

		if flag {
			if start >= 0 {
				spans = append(spans, span{start, end})
				start = ^start
			}
		} else {
			if start < 0 {
				start = end
			}
		}
	}

	if start >= 0 {
		spans = append(spans, span{start, len(s)})
	}

	exclude = ArrayToUpper(exclude)
	a := make([]string, len(spans))
	flag := false
	jump := 0
	for i, span := range spans {
		if span.start < jump {
			continue
		}

		a[i] = s[span.start:span.end]

		if HasArrayPrefix(strings.ToUpper(a[i]), exclude) {
			flag = true
			continue
		}

		if flag && exclude != nil {
			for _, item := range exclude {
				guest := s[spans[i-1].start : spans[i-1].start+len(item)]
				if InArray(exclude, strings.ToUpper(guest)) {
					a[i] = guest
					a[i-1] = ""
					jump = spans[i-1].start + len(item)
					break
				}
			}
			flag = false
		}
	}

	k := 0
	for _, v := range a {
		if v != "" {
			a[k] = v
			k++
		}
	}

	return a[:k]
}
