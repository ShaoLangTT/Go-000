package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")
)
