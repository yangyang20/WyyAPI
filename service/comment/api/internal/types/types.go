// Code generated by goctl. DO NOT EDIT.
package types

type CommentReq struct {
	T         string `form:"t"`
	Type      string `form:"type"`
	Id        int    `form:"id"`
	Content   string `form:"content,optional"`
	CommentId string `form:"commentId,optional"`
	ThreadId  string `form:"threadId,optional"`
}

type CommentReply struct {
	Code    int         `json:"code"`
	Comment interface{} `json:"comment,omitempty"`
	Msg     string      `json:"msg,omitempty"`
}

type CommentAlbumReq struct {
	Rid        string `form:"rid"`
	Limit      int    `form:"limit,default=20"`
	Offset     int    `form:"offset,default=0"`
	BeforeTime string `form:"beforeTime,default=0"`
}

type CommentAlbumReply struct {
	IsMusician    bool        `json:"isMusician"`
	Cnum          int         `json:"cnum"`
	UserId        int         `json:"userId"`
	TopComments   interface{} `json:"topComments"`
	MoreHot       bool        `json:"moreHot"`
	HotComments   interface{} `json:"hotComments"`
	CommentBanner interface{} `json:"commentBanner"`
	Code          int         `json:"code"`
	Comments      interface{} `json:"comments"`
	Total         int         `json:"total"`
	More          bool        `json:"more"`
}

type CommentDjReq struct {
	Rid        string `form:"rid"`
	Limit      int    `form:"limit,default=20"`
	Offset     int    `form:"offset,default=0"`
	BeforeTime string `form:"beforeTime,default=0"`
}

type CommentDjReply struct {
	IsMusician    bool        `json:"isMusician"`
	Cnum          int         `json:"cnum"`
	UserId        int         `json:"userId"`
	TopComments   interface{} `json:"topComments"`
	MoreHot       bool        `json:"moreHot"`
	HotComments   interface{} `json:"hotComments"`
	CommentBanner interface{} `json:"commentBanner"`
	Code          int         `json:"code"`
	Comments      interface{} `json:"comments"`
	Total         int         `json:"total"`
	More          bool        `json:"more"`
}

type CommentEventReq struct {
	Limit      int    `form:"limit,default=20"`
	Offset     int    `form:"offset,default=0"`
	BeforeTime string `form:"beforeTime,default=0"`
	ThreadId   string `form:"threadId"`
}

type CommentEventReply struct {
	IsMusician    bool        `json:"isMusician"`
	Cnum          int         `json:"cnum"`
	UserId        int         `json:"userId"`
	TopComments   interface{} `json:"topComments"`
	MoreHot       bool        `json:"moreHot"`
	HotComments   interface{} `json:"hotComments"`
	CommentBanner interface{} `json:"commentBanner"`
	Code          int         `json:"code"`
	Comments      interface{} `json:"comments"`
	Total         int         `json:"total"`
	More          bool        `json:"more"`
}

type CommentFloorReq struct {
	ParentCommentId string `form:"parentCommentId"`
	Type            string `form:"type"`
	Time            string `form:"time,default=-1"`
	Limit           int    `form:"limit,default=20"`
	Id              int    `form:"id"`
}

type CommentFloorReply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CommentHotReq struct {
	Type       string `form:"type"`
	Rid        string `form:"rid"`
	Limit      int    `form:"limit,default=20"`
	Offset     int    `form:"offset,default=0"`
	BeforeTime string `form:"beforeTime,default=0"`
}

type CommentHotReply struct {
	TopComments interface{} `json:"topComments"`
	HasMore     bool        `json:"hasMore"`
	HotComments interface{} `json:"hotComments"`
	Total       int         `json:"total"`
	Code        int         `json:"code"`
}

type CommentHugListReq struct {
	Type     string `form:"type,default=0"`
	Sid      string `form:"sid"`
	Uid      string `form:"uid"`
	Cid      string `form:"cid"`
	Cursor   string `form:"cursor,default=-1"`
	Page     string `form:"page,default=1"`
	IdCursor string `form:"idCursor,default=-1"`
	PageSize int    `form:"pageSize,default=100"`
}

type CommentHugListReply struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type CommentLikeReq struct {
	Id       string `form:"id"`
	T        string `form:"t,options=like,unlike"`
	Type     string `form:"type"`
	Cid      string `form:"cid"`
	ThreadId string `form:"threadId"`
}

type CommentLikeReply struct {
	Code int `json:"code"`
}

type CommentMusicReq struct {
	Type       int    `form:"type"`
	Rid        string `form:"rid"`
	Limit      int    `form:"limit,default=20"`
	Offset     int    `form:"offset,default=0"`
	BeforeTime string `form:"beforeTime,default=0"`
}

type CommentMusicReply struct {
	IsMusician    bool        `json:"isMusician"`
	Cnum          int         `json:"cnum"`
	UserId        int         `json:"userId"`
	TopComments   interface{} `json:"topComments"`
	MoreHot       bool        `json:"moreHot"`
	HotComments   interface{} `json:"hotComments"`
	CommentBanner interface{} `json:"commentBanner"`
	Code          int         `json:"code"`
	Comments      interface{} `json:"comments"`
	Total         int         `json:"total"`
	More          bool        `json:"more"`
}

type CommentMvReq struct {
	Type       int    `form:"type"`
	Rid        string `form:"rid"`
	Limit      int    `form:"limit,default=20"`
	Offset     int    `form:"offset,default=0"`
	BeforeTime string `form:"beforeTime,default=0"`
}

type CommentMvReply struct {
	IsMusician    bool        `json:"isMusician"`
	Cnum          int         `json:"cnum"`
	UserId        int         `json:"userId"`
	TopComments   interface{} `json:"topComments"`
	MoreHot       bool        `json:"moreHot"`
	HotComments   interface{} `json:"hotComments"`
	CommentBanner interface{} `json:"commentBanner"`
	Code          int         `json:"code"`
	Comments      interface{} `json:"comments"`
	Total         int         `json:"total"`
	More          bool        `json:"more"`
}

type CommentNewRep struct {
	Type      string `form:"type"`
	Id        string `form:"id"`
	PageSize  int    `form:"pageSize,default=20"`
	PageNo    int    `form:"pageNo,default=1"`
	SortType  int    `form:"sortType,default=99"`
	ShowInner bool   `form:"showInner,default=true"`
	Cursor    string `form:"cursor,omitempty,default=0"`
}

type CommentNewReply struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type CommentPlayListReq struct {
	Rid        string `form:"rid"`
	Limit      int    `form:"limit,default=20"`
	Offset     int    `form:"offset,default=0"`
	BeforeTime string `form:"beforeTime,default=0"`
}

type CommentPlayListReply struct {
	IsMusician    bool        `json:"isMusician"`
	Cnum          int         `json:"cnum"`
	UserId        int         `json:"userId"`
	TopComments   interface{} `json:"topComments"`
	MoreHot       bool        `json:"moreHot"`
	HotComments   interface{} `json:"hotComments"`
	CommentBanner interface{} `json:"commentBanner"`
	Code          int         `json:"code"`
	Comments      interface{} `json:"comments"`
	Total         int         `json:"total"`
	More          bool        `json:"more"`
}

type CommentVideoReq struct {
	Rid        string `form:"rid"`
	Limit      int    `form:"limit,default=20"`
	Offset     int    `form:"offset,default=0"`
	BeforeTime string `form:"beforeTime,default=0"`
}

type CommentVideoReply struct {
	IsMusician    bool        `json:"isMusician"`
	Cnum          int         `json:"cnum"`
	UserId        int         `json:"userId"`
	TopComments   interface{} `json:"topComments"`
	MoreHot       bool        `json:"moreHot"`
	HotComments   interface{} `json:"hotComments"`
	CommentBanner interface{} `json:"commentBanner"`
	Code          int         `json:"code"`
	Comments      interface{} `json:"comments"`
	Total         int         `json:"total"`
	More          bool        `json:"more"`
}
