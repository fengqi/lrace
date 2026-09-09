package lrace

import "testing"

func TestStringSplitWith(t *testing.T) {
	type testCase struct {
		str     string
		sep     []string
		split   []string
		exclude []string
	}

	cases := []testCase{
		{
			str:     "a.b-c.d",
			sep:     []string{".", "-"},
			split:   []string{"a", "b-c", "d"},
			exclude: []string{"b-c"},
		},
		{
			str:     "a.b-c@d",
			sep:     []string{".", "-", "@"},
			split:   []string{"a", "b", "c", "d"},
			exclude: []string{},
		},
		{
			// 短语后半段自身也是其他短语的前缀时（HD 是 HDR 的前缀），合并不能被前缀检查抢占
			str:   "Midway 2019 2160p CAN UHD Blu-ray HEVC DTS-HD MA 5.1-THDBST@HDSky.nfo",
			sep:   []string{"-", ".", ",", "_", " ", "[", "]", "(", ")", "{", "}", "@", ":", "："},
			split: []string{"Midway", "2019", "2160p", "CAN", "UHD", "Blu-ray", "HEVC", "DTS-HD", "MA 5.1", "THDBST", "HDSky", "nfo"},
			exclude: []string{
				"WEB-DL", "DDP5.1", "DDP 5.1", "DDP.5.1", "H.265", "H265", "BLU-RAY",
				"MA5.1", "MA 5.1", "MA.5.1", "MA7.1", "MA 7.1", "MA.7.1", "DTS-HD", "HDR", "SDR", "DV",
			},
		},
		{
			// 片段恰好完整等于某个短语时（HDR），后续片段不能被吞掉
			str:   "Movie 2020 1080p HDR HEVC DTS-HD MA 5.1.mkv",
			sep:   []string{"-", ".", ",", "_", " ", "[", "]", "(", ")", "{", "}", "@", ":", "："},
			split: []string{"Movie", "2020", "1080p", "HDR", "HEVC", "DTS-HD", "MA 5.1", "mkv"},
			exclude: []string{
				"WEB-DL", "DDP5.1", "DDP 5.1", "DDP.5.1", "H.265", "H265", "BLU-RAY",
				"MA5.1", "MA 5.1", "MA.5.1", "MA7.1", "MA 7.1", "MA.7.1", "DTS-HD", "HDR", "SDR", "DV",
			},
		},
		{
			// 末尾剩余长度不足一个短语时不能越界
			str:     "Movie DTS AB",
			sep:     []string{"-", ".", ",", "_", " ", "[", "]", "(", ")", "{", "}", "@", ":", "："},
			split:   []string{"Movie", "DTS", "AB"},
			exclude: []string{"DTS-HD"},
		},
		{
			// 短语只匹配到半个片段时不应合并
			str:     "MA 5.12",
			sep:     []string{"-", ".", ",", "_", " ", "[", "]", "(", ")", "{", "}", "@", ":", "："},
			split:   []string{"MA", "5", "12"},
			exclude: []string{"MA 5.1"},
		},
	}
	for _, item := range cases {
		give := StringSplitWith(item.str, item.sep, item.exclude)
		if !ArrayCompare(item.split, give, true) {
			t.Errorf("SplitWith(%s, %v) give: %v, want: %v", item.str, item.sep, give, item.split)
		}
	}
}
