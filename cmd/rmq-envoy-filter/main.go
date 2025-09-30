package main

import (
	"github.com/proxy-wasm/proxy-wasm-go-sdk/proxywasm"
	"github.com/proxy-wasm/proxy-wasm-go-sdk/proxywasm/types"
)

func init() {
	proxywasm.SetVMContext(&vmContext{})
}

func main() {}

type vmContext struct {
	types.DefaultVMContext
}

func (_ *vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	types.DefaultPluginContext
}

func (_ *pluginContext) NewTcpContext(contextID uint32) types.TcpContext {
	return &tcpContext{}
}

type tcpContext struct {
	types.DefaultTcpContext
}

func (_ *tcpContext) OnNewConnection() types.Action {
	proxywasm.LogInfof("new connection...")
	return types.ActionContinue
}
