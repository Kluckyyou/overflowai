package service

import (
	"github.com/Kluckyyou/overflowai/internal/config"
	"github.com/Kluckyyou/overflowai/internal/util/responseheaders"
)

func compileResponseHeaderFilter(cfg *config.Config) *responseheaders.CompiledHeaderFilter {
	if cfg == nil {
		return nil
	}
	return responseheaders.CompileHeaderFilter(cfg.Security.ResponseHeaders)
}
