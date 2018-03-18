package httpjsonrpc

import (
	"DNA/net/httpjsonrpc/common"
	"encoding/json"
	"fmt"
	log "github.com/alecthomas/log4go"
	"io/ioutil"
	"net/http"
)

var DefJsonRpcSvr = NewJsonRpcServer()

type JsonRpcServer struct {
	address  string
	handlers map[string]func(req *common.JsonRpcRequest, resp *common.JsonRpcResponse)
	httpSvr  *http.Server
}

func NewJsonRpcServer() *JsonRpcServer {
	return &JsonRpcServer{
		handlers: make(map[string]func(req *common.JsonRpcRequest, resp *common.JsonRpcResponse)),
	}
}

func (this *JsonRpcServer) Start(address string) {
	this.address = address
	this.httpSvr = &http.Server{
		Addr:    address,
		Handler: http.DefaultServeMux,
	}
	log.Info("JsonRpcServer start at:%s", this.address)
	defer log.Info("JsonRpcServer stop")

	http.HandleFunc("/", this.Handler)
	err := this.httpSvr.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			return
		}
		panic(fmt.Sprintf("httpSvr.ListenAndServe error:%s", err))
	}
}

func (this *JsonRpcServer) RegHandler(method string, handler func(req *common.JsonRpcRequest, resp *common.JsonRpcResponse)) {
	this.handlers[method] = handler
}

func (this *JsonRpcServer) GetHandler(method string) func(req *common.JsonRpcRequest, resp *common.JsonRpcResponse) {
	handler, ok := this.handlers[method]
	if !ok {
		return nil
	}
	return handler
}

func (this *JsonRpcServer) Handler(w http.ResponseWriter, r *http.Request) {
	resp := &common.JsonRpcResponse{}
	defer func() {
		w.WriteHeader(http.StatusOK)

		if resp.ErrorInfo == "" {
			resp.ErrorInfo = common.GetRPCErrorDesc(resp.ErrorCode)
		}
		data, err := json.Marshal(resp)
		if err != nil {
			log.Error("JsonRpcServer josn.Marshal JsonRpcResponse:%+v error:%s", resp, err)
			return
		}
		_, err = w.Write(data)
		if err != nil {
			log.Error("JsonRpcServer Write:%s error %s", data, err)
			return
		}
		log.Info("[JsonRpcResponse]%s", data)
	}()

	if r.Method != http.MethodPost {
		resp.ErrorCode = common.RPCERR_HTTP_METHOD_INVALID
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("JsonRpcServer read body error:%s", err)
		resp.ErrorCode = common.RPCERR_INVALID_REQUEST
		resp.ErrorInfo = "invalid body"
		return
	}
	defer r.Body.Close()

	log.Info("[JosnRpcRequest]%s", data)
	req := &common.JsonRpcRequest{}
	err = json.Unmarshal(data, req)
	if err != nil {
		log.Error("JsonRpcServer json.Unmarshal JsonRpcRequest:%s error:%s", data, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		resp.ErrorInfo = "invalid params"
		return
	}

	resp.Method = req.Method
	resp.Qid = req.Qid

	handler := this.GetHandler(req.Method)
	if handler == nil {
		resp.ErrorCode = common.RPCERR_UNSUPPORT_METHOD
		resp.ErrorInfo = "unsupport method"
		return
	}

	handler(req, resp)
}

func (this *JsonRpcServer) Close() {
	err := this.httpSvr.Close()
	if err != nil {
		log.Error("httpSvr close error:%s", err)
	}
}
