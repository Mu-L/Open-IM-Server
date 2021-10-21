/*
** description("").
** copyright('Open_IM,www.Open_IM.io').
** author("fg,Gordon@tuoyun.net").
** time(2021/5/21 15:29).
 */
package gate

import (
	"Open_IM/src/common/constant"
	"Open_IM/src/common/log"
	"github.com/mitchellh/mapstructure"
)

type Req struct {
	ReqIdentifier int32       `json:"reqIdentifier" validate:"required"`
	Token         string      `json:"token" validate:"required"`
	SendID        string      `json:"sendID" validate:"required"`
	OperationID   string      `json:"operationID" validate:"required"`
	MsgIncr       string      `json:"msgIncr" validate:"required"`
	Data          interface{} `json:"data"`
}
type Resp struct {
	ReqIdentifier int32       `json:"reqIdentifier"`
	MsgIncr       string      `json:"msgIncr"`
	OperationID   string      `json:"operationID"`
	ErrCode       int32       `json:"errCode"`
	ErrMsg        string      `json:"errMsg"`
	Data          interface{} `json:"data"`
}

type SeqData struct {
	SeqBegin int64 `mapstructure:"seqBegin" validate:"required"`
	SeqEnd   int64 `mapstructure:"seqEnd" validate:"required"`
}
type MsgData struct {
	PlatformID  int32                  `mapstructure:"platformID" validate:"required"`
	SessionType int32                  `mapstructure:"sessionType" validate:"required"`
	MsgFrom     int32                  `mapstructure:"msgFrom" validate:"required"`
	ContentType int32                  `mapstructure:"contentType" validate:"required"`
	RecvID      string                 `mapstructure:"recvID" validate:"required"`
	ForceList   []string               `mapstructure:"forceList"`
	Content     string                 `mapstructure:"content" validate:"required"`
	Options     map[string]interface{} `mapstructure:"options" validate:"required"`
	ClientMsgID string                 `mapstructure:"clientMsgID" validate:"required"`
	OfflineInfo map[string]interface{} `mapstructure:"offlineInfo" validate:"required"`
	Ext         map[string]interface{} `mapstructure:"ext"`
}
type SeqListData struct {
	SeqList []int64 `mapstructure:"seqList" validate:"required"`
}

func (ws *WServer) argsValidate(m *Req, r int32) (isPass bool, errCode int32, errMsg string, data interface{}) {
	switch r {
	case constant.WSPullMsg:
		data = SeqData{}
	case constant.WSSendMsg:
		data = MsgData{}
	case constant.WSPullMsgBySeqList:
		data = SeqListData{}
	default:
	}
	if err := mapstructure.WeakDecode(m.Data, &data); err != nil {
		log.ErrorByKv("map to Data struct  err", "", "err", err.Error(), "reqIdentifier", r)
		return false, 203, err.Error(), nil
	} else if err := validate.Struct(data); err != nil {
		log.ErrorByKv("data args validate  err", "", "err", err.Error(), "reqIdentifier", r)
		return false, 204, err.Error(), nil

	} else {
		return true, 0, "", data
	}

}
