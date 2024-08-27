package success

import (
	"fmt"
	"net/http"
)

// ###################################################
// #              Second Syntax Option
// ###################################################

type BuilderOption interface {
	apply(*Success)
}

type BuilderOptionFunc func(*Success)

func (f BuilderOptionFunc) apply(success *Success) { f(success) }

func New(options ...BuilderOption) *Success {
	success := &Success{}

	for _, option := range options {
		option.apply(success)
	}

	return success
}

func Title(title string) BuilderOption {
	return BuilderOptionFunc(func(success *Success) {
		success.WithTitle(title)
	})
}

func Titlef(title string, values ...any) BuilderOption {
	return BuilderOptionFunc(func(success *Success) {
		success.WithTitle(fmt.Sprintf(title, values...))
	})
}

func Status(status int) BuilderOption {
	return BuilderOptionFunc(func(success *Success) {
		success.WithStatus(status)
	})
}

func Detail(detail string) BuilderOption {
	return BuilderOptionFunc(func(success *Success) {
		success.WithDetail(detail)
	})
}

func Detailf(detail string, values ...any) BuilderOption {
	return BuilderOptionFunc(func(success *Success) {
		success.WithDetail(fmt.Sprintf(detail, values...))
	})
}

func Instance(instance string) BuilderOption {
	return BuilderOptionFunc(func(success *Success) {
		success.WithInstance(instance)
	})
}

func InstanceFromRequest(r *http.Request) BuilderOption {
	return BuilderOptionFunc(func(success *Success) {
		success.WithInstance(r.URL.Path)
	})
}

func Type(typeString string) BuilderOption {
	return BuilderOptionFunc(func(success *Success) {
		success.WithType(typeString)
	})
}

func Typef(typeString string, values ...any) BuilderOption {
	return BuilderOptionFunc(func(success *Success) {
		success.WithType(fmt.Sprintf(typeString, values...))
	})
}

func With(key string, value any) BuilderOption {
	return BuilderOptionFunc(func(success *Success) {
		success.WithExtension(key, value)
	})
}
