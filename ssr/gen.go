// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package ssr

import (
	"github.com/hinego/gfen/genx"
)

type (
	IGen interface {
		Execute(in *genx.Execute) (err error)
		GetModule() string
	}
)

var (
	localGen IGen
)

func Gen() IGen {
	if localGen == nil {
		panic("implement not found for interface IGen, forgot register?")
	}
	return localGen
}

func RegisterGen(i IGen) {
	localGen = i
}
