package gen

import (
	"flag"
)

type Manipulator func(stack NodeStack, context TransformContext) error
type ManipulationCallback func() error

type TransformContext interface {
	GetService(string) interface{}
	GetOptions() *flag.FlagSet
	TransformationCount() int
}

type TransformationPlan interface {
	TransformContext
	Run() error
}

func CreateTransformationPlan(handlers []NodeHandler, options flag.FlagSet) TransformationPlan {
	return contextImplemenation{options: options, handlers: handlers}
}

type contextImplemenation struct {
	options flag.FlagSet
	handlers []NodeHandler
}

func (context contextImplemenation) GetService(name string) interface{} {
	for _, handler := range context.handlers {
		service := handler.getService(name)
		if service != nil {
			return service
		}
	}
	return nil
}

func (context contextImplemenation) GetOptions() *flag.FlagSet {
	return &context.options
}

func run1(context TransformContext, handler *NodeHandler, options *flag.FlagSet) error {
	return handler.manipulate(*handler.stack, context)
}

func (context contextImplemenation) Run() error {
	for _, handler := range context.handlers {
		e := run1(context, &handler, &context.options)
		if e != nil {
			return e
		}
	}
	return nil
}

func (context contextImplemenation) TransformationCount() int {
	return len(context.handlers)
}