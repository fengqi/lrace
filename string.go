package lrace

import "strings"

// splitSpan 原始字符串中一个片段的起止位置
type splitSpan struct {
	start int
	end   int
}

// StringSplitWith 以指定的多个字符串切割string，且跳过部分特殊字符
// 灵感来源：strings.FieldsFunc()
func StringSplitWith(s string, sep []string, exclude []string) []string {
	// 先按分隔符切出各片段的起止位置
	spans := make([]splitSpan, 0, 32)
	start := -1
	for end, item := range s {
		if InArray(sep, string(item)) {
			if start >= 0 {
				spans = append(spans, splitSpan{start, end})
				start = ^start
			}
		} else {
			if start < 0 {
				start = end
			}
		}
	}

	if start >= 0 {
		spans = append(spans, splitSpan{start, len(s)})
	}

	// 逐片段输出：片段起始处的原文命中 exclude 项时，整个短语（含内部包裹的分隔符）作为一项输出
	a := make([]string, 0, len(spans))
	for i := 0; i < len(spans); {
		matchLen := matchExclude(s, spans, i, exclude)
		if matchLen > 0 {
			end := spans[i].start + matchLen
			a = append(a, s[spans[i].start:end])
			for i < len(spans) && spans[i].start < end { // 跳过短语覆盖到的片段
				i++
			}
			continue
		}

		a = append(a, s[spans[i].start:spans[i].end])
		i++
	}

	return a
}

// matchExclude 查找从 spans[i] 起始的原文命中的 exclude 项，返回命中项的字节长度，未命中返回0
// 大小写不敏感；命中项的结束位置必须对齐片段边界，避免吞掉半个片段；多个命中时取最长的
func matchExclude(s string, spans []splitSpan, i int, exclude []string) int {
	matchLen := 0
	for _, item := range exclude {
		if len(item) <= matchLen {
			continue
		}

		end := spans[i].start + len(item)
		if end > len(s) || !strings.EqualFold(s[spans[i].start:end], item) {
			continue
		}

		// 结束位置需要落在某个片段的结尾上
		aligned := false
		for _, span := range spans[i:] {
			if span.end == end {
				aligned = true
				break
			}
		}
		if !aligned {
			continue
		}

		matchLen = len(item)
	}

	return matchLen
}
