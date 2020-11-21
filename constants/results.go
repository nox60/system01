package constants

const SUCCESSED int = 1
const FAILED int = 0
const RECORD_EXISTED int = 20001

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
	default:
		resultMsg = "操作成功"
	}
	return resultMsg
}
