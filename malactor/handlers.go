/**
 * MIT License
 *
 * Copyright (c) 2018 - 2019 ccsdsmo
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package malactor

import (
	. "github.com/CNES/ccsdsmo-malgo/mal"
)

type handler interface {
}

type provider interface {
	handler
}

type consumer interface {
	handler
}

// ================================================================================
// MAL Send interaction handler

type ProviderSend interface {
	provider
	OnSend(endpoint *EndPoint, msg *Message)
}

// ================================================================================
// MAL Submit interaction handlers

type ProviderSubmit interface {
	provider
	OnSubmit(endpoint *EndPoint, msg *Message)
}

type ConsumerSubmit interface {
	consumer
	OnAck(endpoint *EndPoint, msg *Message)
}

// ================================================================================
// MAL Request interaction handlers

type ProviderRequest interface {
	provider
	OnRequest(endpoint *EndPoint, msg *Message)
}

type ConsumerRequest interface {
	consumer
	OnResponse(endpoint *EndPoint, msg *Message)
}

// ================================================================================
// MAL Invoke interaction handlers

type ProviderInvoke interface {
	provider
	OnInvoke(endpoint *EndPoint, msg *Message)
}

type ConsumerInvoke interface {
	consumer
	OnAck(endpoint *EndPoint, msg *Message)
	OnResponse(endpoint *EndPoint, msg *Message)
}

// ================================================================================
// MAL Progress interaction handlers

type ProviderProgress interface {
	provider
	OnProgress(endpoint *EndPoint, msg *Message)
}

type ConsumerProgress interface {
	consumer
	OnAck(endpoint *EndPoint, msg *Message)
	OnUpdate(endpoint *EndPoint, msg *Message)
	OnResponse(endpoint *EndPoint, msg *Message)
}

// ================================================================================
// MAL PubSub interaction handlers

type ProviderPubSub interface {
	provider
	OnPublishRegisterAck(endpoint *EndPoint, msg *Message)
	OnPublishDeregisterAck(endpoint *EndPoint, msg *Message)
	OnPublishError(endpoint *EndPoint, msg *Message)
}

type ConsumerPubSub interface {
	consumer
	OnRegisterAck(endpoint *EndPoint, msg *Message)
	OnDeregister(endpoint *EndPoint, msg *Message)
	OnNotify(endpoint *EndPoint, msg *Message)
}

type BrokerPubSub interface {
	handler
	OnRegister(endpoint *EndPoint, msg *Message)
	OnDeregister(endpoint *EndPoint, msg *Message)
	OnPublishRegister(endpoint *EndPoint, msg *Message)
	OnPublishDeregister(endpoint *EndPoint, msg *Message)
	OnPublish(endpoint *EndPoint, msg *Message)
	OnNotifyError(endpoint *EndPoint, msg *Message)
}
