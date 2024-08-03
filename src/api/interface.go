package api

import (
	v1_0 "github.com/be-freelancemitra-web/src/api/v1"
)

type ApiV1_0Layer interface {
	GetV1_0ApiLayer() v1_0.ApiLayer
}

type apiV1_0Object struct {
	v1_0 v1_0.ApiLayer
}

func NewApiV1Obj() ApiV1_0Layer {
	return &apiV1_0Object{v1_0: v1_0.NewApiObj()}
}

func (s *apiV1_0Object) GetV1_0ApiLayer() v1_0.ApiLayer {
	return s.v1_0
}
