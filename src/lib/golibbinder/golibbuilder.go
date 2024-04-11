package main

type goLibBuilder struct {
	className        string
	methodReturnType string
	methodName       string
	methodParameter  string
}

func NewGoLibBuilder() *goLibBuilder {
	return &goLibBuilder{}
}

func (goLibBuilder *goLibBuilder) SetClassName(className *string) {
	goLibBuilder.className = *className
}

func (goLibBuilder *goLibBuilder) SetMethodReturnType(methodReturnType *string) {
	goLibBuilder.methodReturnType = *methodReturnType
}

func (goLibBuilder *goLibBuilder) SetMethodName(methodName *string) {
	goLibBuilder.methodName = *methodName
}

func (goLibBuilder *goLibBuilder) SetMethodParameter(methodParameter *string) {
	goLibBuilder.methodParameter = *methodParameter
}

func (goLibBuilder *goLibBuilder) GetClassName() *string {
	return &goLibBuilder.className
}

func (goLibBuilder *goLibBuilder) GetMethodReturnType() *string {
	return &goLibBuilder.methodReturnType
}

func (goLibBuilder *goLibBuilder) GetMethodName() *string {
	return &goLibBuilder.methodName
}

func (goLibBuilder *goLibBuilder) GetMethodParameter() *string {
	return &goLibBuilder.methodParameter
}

func (goLibBuilder *goLibBuilder) BuildGoLibJavaClass() *string {

	var classTemplate = `
	package com.jgo.framework.gorunner.golib;

	import com.sun.jna.Library;

	public interface ` + goLibBuilder.className + `extends Library {` +
		goLibBuilder.methodReturnType + ` ` + goLibBuilder.methodName + ` (` + goLibBuilder.methodParameter + `);` +
		`}`

	return &classTemplate
}
