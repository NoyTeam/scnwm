package scnwm

import (
	"strings"

	"github.com/liuzl/gocc"
	"github.com/mozillazg/go-pinyin"
)

var words map[string]int = map[string]int{
	// 3
	"差不多": 3,
	"狗":   3,
	"他妈":  3,
	// 5
	"得了":  5,
	"死":   5,
	"脑血栓": 5,
	"所谓":  5,
	"骂":   5,
	// 8
	"猴子": 8,
	"你妈": 8,
	"原批": 8,
	"农批": 8,
	"op": 8,
	"粥批": 8,
	"滚":  8,
	"恶心": 8,
	"魔怔": 8,
	// 10
	"傻逼": 10,
	"脑瘫": 10,
	"弱智": 10,
	"残疾": 10,
	"😅":  10,
	"自由": 10,
	// 15
	"革命":  15,
	"退党":  15,
	"共产党": 15,
	"近平":  15,
	"政府":  15,
	"独裁":  15,
	"中共":  15,
	"法轮":  15,
}

var pinyins []string = []string{
	"ni ma",
	"ma de",
	"tmd",
	"sb",
	"sha bi",
	"nt",
	"gong chan",
	"du cai",
	"xjp",
	"jin ping",
}

func Test(text string) (int, error) {
	var num int = 0

	// GOCC
	t2s, err := gocc.New("t2s")
	if err != nil {
		return -1, err
	}
	out, err := t2s.Convert(text)
	if err != nil {
		return -1, err
	}
	// ToLower
	out = strings.ToLower(out)

	// Words
	for i, j := range words {
		if strings.Contains(out, i) {
			num = num + j
		}
	}

	// Pinyin
	if num < 8 {
		pa := pinyin.NewArgs()
		pinyinlist := pinyin.Pinyin(out, pa)

		var pinyintext string
		for _, j := range pinyinlist {
			pinyintext = pinyintext + " " + j[0]
		}

		for _, j := range pinyins {
			if strings.Contains(pinyintext, j) {
				num = num + 8
			}
		}
	}

	return num, nil
}
