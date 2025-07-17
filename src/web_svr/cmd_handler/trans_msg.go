package cmd_handler

import (
	"github.com/Iori372552686/GoOne/lib/util/convert"
	define "github.com/Iori372552686/GoOne/src/web_svr/common"
	"github.com/Iori372552686/GoOne/src/web_svr/web_service"
	g1_protocol "github.com/Iori372552686/game_protocol/protocol"
	"github.com/gin-gonic/gin"
)

func MsgSecCheck(ctx *gin.Context, data []byte) gin.H {
	req := &define.MsgSecCheckReq{}
	err := convert.JsonToStruct(convert.Bytes2str(data), req)
	if err != nil {
		return gin.H{"code": g1_protocol.ErrorCode_ERR_MARSHAL, "data": nil, "msg": g1_protocol.ErrorCode_ERR_MARSHAL.String()}
	}

	ret := web_service.MsgSecCheck(req)
	return gin.H{"code": ret.Code, "data": ret.Argv, "msg": ret.Msg}
}
