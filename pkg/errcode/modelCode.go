package errcode

// var ErrorUserNotExist = NewError(20010001, "不存在该用户")
// var ErrorUserPasswordError = NewError(20010002, "密码错误")
// var ErrorGetUserListFail = NewError(20010003, "获取用户列表失败")
// var ErrorGetUserCountFail = NewError(20010004, "获取用户总数失败")

// const (
// 	SuccessCode = iota
// 	ErrorList   = 500
// 	ErrorCount  = 501
// 	ErrorCreate = 502
// 	ErrorUpdate = 503
// 	ErrorDel    = 504
// )

// func GetErrorMsg(it int, str string) *Error {
// 	switch it {
// 	case ErrorList:
// 		return NewError(ErrorList, fmt.Sprintf(`获取%v错误`, str))
// 	case ErrorCount:
// 		return NewError(ErrorCount, fmt.Sprintf(`获取%v总数错误`, str))
// 	case ErrorCreate:
// 		return NewError(ErrorCreate, fmt.Sprintf(`创建%v失败`, str))
// 	case ErrorUpdate:
// 		return NewError(ErrorUpdate, fmt.Sprintf(`更新%v失败`, str))
// 	case ErrorDel:
// 		return NewError(ErrorDel, fmt.Sprintf(`删除%v失败`, str))
// 	}
// 	return nil
// }

// 错误码
var ErrorUserNotExist = NewError(500101, "不存在该用户")
var ErrorUserPasswordError = NewError(500102, "密码错误")
var ErrorUserListFail = NewError(500103, "获取用户列表失败")
var ErrorUserCountFail = NewError(500104, "获取用户列表总数失败")
var ErrorRoleListFail = NewError(500105, "获取角色列表失败")
var ErrorRoleCountFail = NewError(500106, "获取角色列表总数失败")
var ErrorRoleUpdateFail = NewError(500107, "更新角色失败")
var ErrorRoleNameValidFail = NewError(500108, "角色名称不能为空")
