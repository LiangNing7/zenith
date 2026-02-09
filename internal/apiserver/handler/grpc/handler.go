// Copyright 2025 LiangNing7 <LiangNing7@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/LiangNing7/zenith.

package grpc

import (
	"github.com/LiangNing7/zenith/internal/apiserver/biz"
	apiv1 "github.com/LiangNing7/zenith/pkg/api/apiserver/v1"
)

// Handler 负责处理博客模块的请求.
type Handler struct {
	apiv1.UnimplementedZenithServer

	biz biz.IBiz
}

// NewHandler 创建一个新的 Handler 实例.
func NewHandler(biz biz.IBiz) *Handler {
	return &Handler{
		biz: biz,
	}
}
