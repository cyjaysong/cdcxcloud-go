package smsv2

// SmsSubmitResult 短信提交结果
type SmsSubmitResult int

const (
	SmsSubmitSuccess             SmsSubmitResult = 0   // 提交成功
	SmsSubmitExtnoInvalid        SmsSubmitResult = 10  // 原发号码错误，即extno错误
	SmsSubmitBalanceInsufficient SmsSubmitResult = 15  // 余额不足
	SmsSubmitServerErr           SmsSubmitResult = 100 // 系统内部错误，请联系管理员
)

var SmsStat = map[string]string{
	"DELIVRD":  "短信投递成功",
	"EXPIRED":  "消息有效期已过期",
	"DELETED":  "消息已被删除",
	"REJECTED": "消息处于拒绝状态",
	"MA:0001":  "全局黑名单号码",
	"MA:0002":  "内容非法",
	"MA:0003":  "无法找到下级路由",
	"MA:0004":  "错误的内容长度",
	"MA:0005":  "目的号码格式错误",
	"MA:0006":  "系统拒绝",
	"MA:0009":  "未定义错误",
	"MA:0011":  "未知系统内部错误",
	"MA:0015":  "余额不足",
	"MA:0017":  "签名无效",
	"MA:0021":  "号码格式错误",
	"MA:0022":  "下发次数限制",
	"MA:0023":  "客户黑名单号码",
	"MA:0024":  "内容未报备",
	"MA:0025":  "不支持该短信",
	"MA:0026":  "分条发送，组包超时",
	"MA:0027":  "通道黑名单",
	"MA:0028":  "全局黑名单号段",
	"MA:0029":  "通道黑名单号段",
	"MA:0030":  "直接产生拒绝报告",
	"MA:0033":  "地区拦截",
	"MA:0044":  "号段拦截",
	"MO:200":   "不支持分条短信",
	"MO:0254":  "转发提交超时",
	"MO:0255":  "转发提交过程中，连接断开",
}

type ApiStatus int

const (
	ApiSuccess                    ApiStatus = 0   // 成功
	ApiInvalidIp                  ApiStatus = 2   // IP错误
	ApiAccountVerifyFailed        ApiStatus = 3   // 账号密码错误
	ApiOtherError                 ApiStatus = 5   // 其它错误
	ApiAPNError                   ApiStatus = 6   // 接入点错误（如账户本身开的是CMPP接入）
	ApiAccountAbnormality         ApiStatus = 7   // 账号状态异常（账号已停用）
	ApiServerErr                  ApiStatus = 11  // 系统内部错误，请联系管理员
	ApiInvalidDuration            ApiStatus = 30  // action=statis时:结束时间不能小于开始时间,间隔不能大于180天; action=select时:date大于当天日期,date不在过去180天内
	ApiFrequentTemplateOperations ApiStatus = 31  // action=templateAdd或action=templateDelete时；操作过于频繁，建议每10秒操作一次
	ApiConditionInvalid           ApiStatus = 32  // condition错，只能为APMID或MOBILE
	ApiListOutRange               ApiStatus = 33  // 值列表过多，最多只能1000
	ApiRequestParamInvalid        ApiStatus = 34  // 请求参数有误
	ApiNoMobiles                  ApiStatus = 35  // mobileEencryptionMode=aes时，解析后的正确号码数为0
	ApiTemplateNotExist           ApiStatus = 36  // 未找到模板
	ApiTemplateStatusErr          ApiStatus = 37  // 模板状态有误
	ApiTemplateTimePeriodInvalid  ApiStatus = 38  // 模板时间不允许
	ApiNotSupported               ApiStatus = 39  // 接口不支持
	ApiRequestExpiration          ApiStatus = 40  // 请求超过时效
	ApiRetry                      ApiStatus = 41  // 重放失败
	ApiServerError                ApiStatus = 100 // 系统内部错误，请联系管理员
)
