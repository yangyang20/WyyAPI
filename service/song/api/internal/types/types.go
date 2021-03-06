// Code generated by goctl. DO NOT EDIT.
package types

type SongDetailReq struct {
	Ids string `form:"ids"`
}

type SongDetailReply struct {
	Songs      interface{} `json:"songs"`
	Privileges interface{} `json:"privileges"`
	Code       int         `json:"code"`
}

type SongDownloadUrlReq struct {
	Id int `form:"id"`
	Br int `form:"br,default=999000"`
}

type SongDownloadUrlReply struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}

type SongOrderUpdateReq struct {
	Pid      int    `form:"pid"`
	TrackIds string `form:"trackIds"`
	Op       string `form:"op,default=update"`
}

type SongOrderUpdateReply struct {
}

type SongPurchasedReq struct {
	Limit  int `form:"limit,default=20"`
	Offset int `form:"offset,default=0"`
}

type SongPurchasedReply struct {
}

type SongUrlReq struct {
	Id int `form:"id"`
	Br int `form:"br,default=999000"`
}

type SongUrlReply struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}
