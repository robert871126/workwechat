package workwechat

import (
	"sync"

	"github.com/JefferyWang/workwechat/base"
	"github.com/JefferyWang/workwechat/cache"
	"github.com/JefferyWang/workwechat/context"
	"github.com/JefferyWang/workwechat/crm"
	"github.com/JefferyWang/workwechat/media"
)

// WorkWechat struct
type WorkWechat struct {
	Context *context.Context
}

// Config 企业微信配置信息
type Config struct {
	CorpID     string // 企业ID
	CorpSecret string // 应用的凭证密钥

	Cache cache.Cache
}

// NewWorkWechat init
func NewWorkWechat(cfg *Config) *WorkWechat {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &WorkWechat{context}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.CorpID = cfg.CorpID
	context.CorpSecret = cfg.CorpSecret
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
}

// GetBase 通讯录管理接口
func (wwc *WorkWechat) GetBase() *base.Base {
	return base.NewBase(wwc.Context)
}

// GetMedia 素材管理接口
func (wwc *WorkWechat) GetMedia() *media.Media {
	return media.NewMedia(wwc.Context)
}

// GetCrm 外部联系人接口
func (wwc *WorkWechat) GetCrm() *crm.Crm {
	return crm.NewCrm(wwc.Context)
}
