package fortune

import (
	"math/rand"
	"time"
)

type Luck int

const (
	Great      Luck = iota // 大吉
	Middle                 // 中吉
	Small                  // 小吉
	Normal                 // 吉
	Uncertain              // 末吉
	Curse                  // 凶
	GreatCurse             // 大凶
	Kind                   // 種類数
)

func (c Luck) String() string {
	switch c {
	case Great:
		return "大吉"
	case Middle:
		return "中吉"
	case Small:
		return "小吉"
	case Normal:
		return "吉"
	case Uncertain:
		return "末吉"
	case Curse:
		return "凶"
	case GreatCurse:
		return "大凶"
	}
	panic("Unknown value")
}

func Draw(t time.Time) Luck {
	if t.Month() == time.January && t.Day() >= 1 && t.Day() <= 3 {
		return Great
	}
	return Luck(rand.Intn(int(Kind)))
}
