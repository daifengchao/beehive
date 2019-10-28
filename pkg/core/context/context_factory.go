package context

import (
	"sync"
	"k8s.io/klog"
)

//define channel type
const (
	MsgCtxTypeChannel = "channel"
)

var (
	// singleton
	context *Context
	once    sync.Once
)

// GetContext gets global context instance
func GetContext(contextType string) *Context {
	once.Do(func() {
		context = &Context{}
		switch contextType {
		case MsgCtxTypeChannel:
			channelContext := NewChannelContext()
			context.messageContext = channelContext
			context.moduleContext = channelContext
		default:
			klog.Warningf("Do not support context type:%s", contextType)
		}
	})
	return context
}

// AddModule adds module into module context
func (ctx *Context) AddModule(module string) {
	ctx.moduleContext.AddModule(module)
}

// AddModuleGroup adds module into module context group
func (ctx *Context) AddModuleGroup(module, group string) {
	ctx.moduleContext.AddModuleGroup(module, group)
}

// Cleanup cleans up module
func (ctx *Context) Cleanup(module string) {
	ctx.moduleContext.Cleanup(module)
}

// Send the message
func (ctx *Context) Send(module string, message interface{}) {
	ctx.messageContext.Send(module, message)
}

// Receive the message
// module : local module name
func (ctx *Context) Receive(module string) (interface{}, error) {
	message, err := ctx.messageContext.Receive(module)
	if err == nil {
		return message, nil
	}
	klog.Warning("Receive: failed to receive message")
	return message, err
}
