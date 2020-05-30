package utility

import (
	"encoding/json"
	"net/http"
)

type ReturnJson struct {
	Code    string
	Message string
	Data    interface{}
}

const (
	ConstSuccessCode = "200"
	ConstFailCode    = "400"
	ConstSuccess     = "Success"
	ConstEMPTY       = ""
)

//SetReturnData... common function to set Return data
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

//ReturnResponse... common function to set response
func ReturnResponse(w http.ResponseWriter, data interface{}) {
	returnJson, _ := json.Marshal(data)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnJson)
}
