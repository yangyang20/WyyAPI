package config

import (
	"github.com/zeromicro/go-zero/rest"
	"strconv"
)

type Config struct {
	rest.RestConf
}

const (
	R_SO_4_ = iota
	R_MV_5_
	A_PL_0_
	R_AL_3_
	A_DJ_1_
	R_VI_62_
	A_EV_2_
)

func ResourceTypeMap(p string) string {
	switch p {
	case strconv.Itoa(R_SO_4_):
		return "R_SO_4_"
	case strconv.Itoa(R_MV_5_):
		return "R_MV_5_"
	case strconv.Itoa(A_PL_0_):
		return "A_PL_0_"
	case strconv.Itoa(R_AL_3_):
		return "R_AL_3_"
	case strconv.Itoa(A_DJ_1_):
		return "A_DJ_1_"
	case strconv.Itoa(R_VI_62_):
		return "R_VI_62_"
	case strconv.Itoa(A_EV_2_):
		return "A_EV_2_"
	default:
		return "UNKNOWN"
	}
}
