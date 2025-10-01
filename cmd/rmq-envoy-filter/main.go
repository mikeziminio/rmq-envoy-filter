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

func (tc *tcpContext) OnDownstreamData(size int, endStream bool) types.Action {
	data, err := proxywasm.GetDownstreamData(0, size)
	if err != nil {
		proxywasm.LogErrorf("failed to get downstream data: %v", err)
	}
	proxywasm.LogInfof("downstream data len: %d", len(data))
	return types.ActionContinue
}

func (tc *tcpContext) OnUpstreamData(size int, endStream bool) types.Action {
	data, err := proxywasm.GetUpstreamData(0, size)
	if err != nil {
		proxywasm.LogErrorf("failed to get upstream data: %v", err)
	}
	proxywasm.LogInfof("upstream data len: %d", len(data))
	return types.ActionContinue
}
