package errcode

var ErrorUserNotExist = NewError(20010001, "不存在该用户")
var ErrorUserPasswordError = NewError(20010002, "密码错误")
var ErrorGetUserListFail = NewError(20010003, "获取用户列表失败")
var ErrorGetUserCountFail = NewError(20010004, "获取用户总数失败")
var ErrorGetThresholdFail = NewError(20010005, "获取阈值失败，请检查超影客户端程序是否启动")
var ErrorSetThresholdFail = NewError(20010006, "设置阈值失败，请检查超影客户端程序是否启动")
var ErrorGetReportTemplateFail = NewError(20010007, "获取报告模板失败，请检查超影客户端程序是否启动")
var ErrorSetReportTemplateFail = NewError(20010008, "设置报告模板失败，请检查超影客户端程序是否启动")
var ErrorGetLVEFThresholdFail = NewError(20010009, "获取LVEF阈值失败，请检查超影客户端程序是否启动")
var ErrorSetLVEFThresholdFail = NewError(20010010, "设置LVEF阈值失败，请检查超影客户端程序是否启动")
var ErrorExportFail = NewError(20010011, "导出数据失败")
var ErrorExportSelectFail = NewError(20010012, "查询数据失败")
var ErrorExportDataEmpty = NewError(20010013, "导出数据为空")
var DownloadFail = NewError(20010014, "下载失败")
var DownloadFileNotExist = NewError(20010015, "下载文件不存在")
var ErrorGetDownloadFileListFail = NewError(20010016, "获取下载文件列表失败")
var ErrorDeleteFileFail = NewError(20010017, "删除文件失败")
var ErrorGetDiskUsageFail = NewError(20010018, "获取磁盘使用状态失败")
var ErrorGetDownloadTimesFail = NewError(20010019, "获取待下载文件个数失败")
var ErrorGetPatientImageListFail = NewError(20010020, "获取图片列表失败")
var ErrorGetPatientLesionsListFail = NewError(20010021, "获取病灶列表失败")
var ErrorGetLastPatientInfoFail = NewError(20010022, "获取最新患者信息失败")
var ErrorGetPatientImageGroupFail = NewError(20010023, "获取图片组信息失败")
var ErrorGetPatientLesionsByImageInfoFail = NewError(20010024, "根据图片信息获取病灶信息失败")
var ErrorGetPatientImageListDescByTimeFail = NewError(20010025, "按时间倒序获取图片列表失败")
var ErrorGetPatientLesionsListDescByTimeFail = NewError(20010026, "按时间倒序获取病灶列表失败")
var ErrorGetGradingStandardFail = NewError(20010027, "获取评级标准失败")
var ErrorGetAiRecommendRetFail = NewError(20010028, "获取AI推荐结果失败")
var ErrorUploadFileFail = NewError(20010029, "上传文件失败")
var ErrorCreateFileFail = NewError(20010030, "创建文件失败")
var ErrorReadFileFail = NewError(20010031, "读取文件失败")
var ErrorUploadFileExist = NewError(20010032, "上传文件已存在")
var ErrorGetUploadFileCountFail = NewError(20010033, "获取上传文件总数失败")
var ErrorGetUploadFileListFail = NewError(20010034, "获取上传文件列表失败")
var ErrorGetPatientByFileIdFail = NewError(20010035, "获取患者信息失败")
var ErrorGetDisplayConfigFail = NewError(20010036, "获取显示配置失败")

var ErrorGetProductAccountListFail = NewError(20010037, "获取产品使用率列表失败")
var ErrorGetProductAccountCountFail = NewError(20010038, "获取产品使用率总数失败")
var ErrorGetAlgorithmStatListFail = NewError(20010039, "获取算法指标统计列表失败")
var ErrorGetAlgorithmStatDetailFail = NewError(20010040, "获取算法指标详细统计数据失败")
var ErrorExportProductUsageDataFail = NewError(20010041, "导出使用率数据失败")
var ErrorExportAlgorithmStatListFail = NewError(20010042, "导出算法统计指标数据失败")
