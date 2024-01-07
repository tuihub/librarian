package biz

import "github.com/Xuanwo/go-locale"

type Locale int

const (
	LocaleEn Locale = iota
	LocaleChs
	LocaleCht
)

func GetLocale() Locale {
	tag, err := locale.Detect()
	if err != nil {
		return LocaleEn
	}
	base, _, region := tag.Raw()
	if base.String() == "zh" {
		if region.String() == "TW" || region.String() == "HK" || region.String() == "MO" {
			return LocaleCht
		}
		return LocaleChs
	}
	return LocaleEn
}
