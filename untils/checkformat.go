package utils

import "regexp"

// 邮箱格式校验
func ValidateEmail(email string) bool {
	result, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
	return result
}

// 手机号格式校验
// 检查字符串是否符合中国大陆手机号码格式 中国大陆手机号码的一般格式，即以1开头，第二位通常是3、4、5、6、7、8、9中的一个数字，后面跟着9位数字
func ValidChineseMobileNumber(phone string) bool {
	// 正则表达式，匹配中国大陆手机号码格式
	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}
