package utility

type ReturnJson struct {
	Code    string
	Message string
	Data    interface{}
}

const (
	ConstSuccessCode = "200"
	ConstFailCode    = "400"
	ConstSuccess     = "Success"
)

func SetReturnData(err error, data interface{}) (returnData ReturnJson) {

	returnData.Code = ConstSuccessCode
	returnData.Message = ConstSuccess
	if err != nil {
		returnData.Code = ConstFailCode
		returnData.Message = err.Error()
	}
	returnData.Data = data

	return
}
