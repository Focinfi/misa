package strings

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Focinfi/go-pipeline"
)

type Unquote struct{}

func (Unquote) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	//respRes := &pipeline.HandleRes{}
	//if reqRes != nil {
	//	respRes, err = reqRes.Copy()
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//}

	s, err := strconv.Unquote(`"\u003e\u003cCMBSDKPGK\u003e\u003cINFO\u003e\u003cFUNNAM\u003eNTIBCOPR\u003c/FUNNAM\u003e\u003cDATTYP\u003e2\u003c/DATTYP\u003e\u003cLGNNAM\u003e汇通大连复核\u003c/LGNNAM\u003e\u003c/INFO\u003e\u003cNTOPRMODX\u003e\u003cBUSMOD\u003e00001\u003c/BUSMOD\u003e\u003c/NTOPRMODX\u003e\u003cNTIBCOPRX\u003e\u003cSQRNBR\u003e1\u003c/SQRNBR\u003e\u003cBBKNBR\u003e10\u003c/BBKNBR\u003e\u003cACCNBR\u003e110925083710102\u003c/ACCNBR\u003e\u003cCNVNBR\u003e0000003127\u003c/CNVNBR\u003e\u003cYURREF\u003e1210096562567901191\u003c/YURREF\u003e\u003cCCYNBR\u003e10\u003c/CCYNBR\u003e\u003cTRSAMT\u003e0.50\u003c/TRSAMT\u003e\u003cCRTSQN\u003e\u003c/CRTSQN\u003e\u003cNTFCH1\u003e\u003c/NTFCH1\u003e\u003cNTFCH2\u003e\u003c/NTFCH2\u003e\u003cCDTNAM\u003e四川泰达天然气有限公司\u003c/CDTNAM\u003e\u003cCDTEAC\u003e51050175004600000353\u003c/CDTEAC\u003e\u003cCDTBRD\u003e105100000017\u003c/CDTBRD\u003e\u003cTRSTYP\u003eC202\u003c/TRSTYP\u003e\u003cTRSCAT\u003e02014\u003c/TRSCAT\u003e\u003cRMKTXT\u003e\u003c/RMKTXT\u003e\u003cRSV30Z\u003e\u003c/RSV30Z\u003e\u003c/NTIBCOPRX\u003e\u003c/CMBSDKPGK\u003e"`)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	return nil, nil
}
