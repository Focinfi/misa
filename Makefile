reinstall_cmd_misa:
	go install ./cmd/misa && misa ls

test: modvendor
	go test ./... -v --cover

modvendor:
	GO111MODULE=on GOPROXY=https://goproxy.cn go mod vendor

gitpush: test
	git push

testHttp:
	misa run get-first-pipeline-id -d '{"method": "GET", "url": "https://raw.githubusercontent.com/Focinfi/misa/master/configs/conf.example.json"}'

testDownloadPic:
	misa run download-pic -d '{\
	    "method": "GET", \
	    "url": "http://service-i-cashdesk-oil--oilaccess.pay-platform-test.k8s.test.chinawayltd.com/cashdesk-oil/v2/oilAccess/rechargeOrder/FK191216473938/electronicBill?returnPic=true", \
	    "header": { \
	        "AccessKey": "fb2298f4-7792-47a0-bde4-308b0faf4076-e9d807e4-af85-4e95-abdf-905dc9223fbd" \
	    } \
	  }' -v

default: reinstall_cmd_misa
