/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/baselib/logger"
	"github.com/nebulaim/telegramd/baselib/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	webpage2 "github.com/nebulaim/telegramd/biz/core/webpage"
)

// messages.getWebPagePreview#25223e24 message:string = MessageMedia;
func (s *MessagesServiceImpl) MessagesGetWebPagePreview(ctx context.Context, request *mtproto.TLMessagesGetWebPagePreview) (*mtproto.MessageMedia, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getWebPagePreview#25223e24 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	webpage := webpage2.GetWebPagePreview(request.Message)
	media := &mtproto.TLMessageMediaWebPage{Data2: &mtproto.MessageMedia_Data{
		Webpage: webpage,
	}}

	glog.Infof("messages.getWebPagePreview#25223e24 - reply: %s\n", logger.JsonDebugData(media))
	return media.To_MessageMedia(), nil
}
