package problem

import (
	"fmt"
	"net/http"
)

// ###################################################
// #              Second Syntax Option
// ###################################################

type BuilderOption interface {
	apply(*Problem)
}

type BuilderOptionFunc func(*Problem)

func (f BuilderOptionFunc) apply(problem *Problem) { f(problem) }

func New(options ...BuilderOption) *Problem {
	problem := &Problem{}

	for _, option := range options {
		option.apply(problem)
	}

	return problem
}

func From(problem *Problem, options ...BuilderOption) *Problem {
	p := *problem
	newP := p

	for _, option := range options {
		option.apply(&newP)
	}

	return &newP
}

func Title(title string) BuilderOption {
	return BuilderOptionFunc(func(problem *Problem) {
		problem.WithTitle(title)
	})
}

func Titlef(title string, values ...any) BuilderOption {
	return BuilderOptionFunc(func(problem *Problem) {
		problem.WithTitle(fmt.Sprintf(title, values...))
	})
}

func Status(status int) BuilderOption {
	return BuilderOptionFunc(func(problem *Problem) {
		problem.WithStatus(status)
	})
}

func Detail(detail string) BuilderOption {
	return BuilderOptionFunc(func(problem *Problem) {
		problem.WithDetail(detail)
	})
}

func Detailf(detail string, values ...any) BuilderOption {
	return BuilderOptionFunc(func(problem *Problem) {
		problem.WithDetail(fmt.Sprintf(detail, values...))
	})
}

func Instance(instance string) BuilderOption {
	return BuilderOptionFunc(func(problem *Problem) {
		problem.WithInstance(instance)
	})
}

func InstanceFromRequest(r *http.Request) BuilderOption {
	return BuilderOptionFunc(func(problem *Problem) {
		problem.WithInstance(r.URL.Path)
	})
}

func Type(typeString string) BuilderOption {
	return BuilderOptionFunc(func(problem *Problem) {
		problem.WithType(typeString)
	})
}

func Typef(typeString string, values ...any) BuilderOption {
	return BuilderOptionFunc(func(problem *Problem) {
		problem.WithType(fmt.Sprintf(typeString, values...))
	})
}

func With(key string, value any) BuilderOption {
	return BuilderOptionFunc(func(problem *Problem) {
		problem.WithExtension(key, value)
	})
}

func Error(err error) BuilderOption {
	return BuilderOptionFunc(func(problem *Problem) {
		problem.WithExtension(DefaultErrorKey, err)
	})
}
