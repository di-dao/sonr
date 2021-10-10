package api

import (
	"github.com/sonr-io/core/internal/common"
)

func NewInitialzeResponse(gpf common.GetProfileFunc, success bool) *InitializeResponse {
	resp := &InitializeResponse{Success: success}
	if !success || gpf == nil {
		return resp
	}
	p, err := gpf()
	if err != nil {
		logger.Error("Failed to get profile", err)
		return resp
	}
	resp.Profile = p
	return resp
}

// IsDev returns true if the node is running in development mode.
func (ir *InitializeRequest) IsDev() bool {
	return ir.GetEnvironment().IsDev()
}
