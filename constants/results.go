package constants

const SUCCESSED int = 1
const FAILED int = 0
const RECORD_EXISTED int = 20001
const PASSWORD_NOT_SAME int = 20002

func GetResultMsgByCode(code int) (resultMsg string) {
	switch code {
	case SUCCESSED:
		resultMsg = "操作成功"
		break
	case FAILED:
		resultMsg = "操作失败"
		break
	case RECORD_EXISTED:
		resultMsg = "记录已存在"
		break
	case PASSWORD_NOT_SAME:
		resultMsg = "密码长度需要超过6位，并且两次密码输入需要一致"
	default:
		resultMsg = "操作成功"
	}
	return resultMsg
}
