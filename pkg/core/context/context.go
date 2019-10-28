package context

//ModuleContext is interface for context module management
type ModuleContext interface {
	AddModule(module string)
	AddModuleGroup(module, group string)
	Cleanup(module string)
}

//MessageContext is interface for message syncing
type MessageContext interface {
	// async mode
	Send(module string, message interface{})
	Receive(module string) (interface{}, error)
}

// Context is global context object
type Context struct {
	moduleContext  ModuleContext
	messageContext MessageContext
}
