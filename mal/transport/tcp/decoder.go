/**
 * MIT License
 *
 * Copyright (c) 2017 - 2019 CNES
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
package tcp

import (
	"errors"
	. "github.com/CNES/ccsdsmo-malgo/mal"
	"github.com/CNES/ccsdsmo-malgo/mal/encoding/binary"
	"strings"
)

func (transport *TCPTransport) decode(buf []byte, from string) (*Message, error) {
	decoder := binary.NewBinaryDecoder(buf, false)

	b, err := decoder.Read()
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot read magic: %s", err.Error())
		return nil, err
	}
	if ((b >> 5) & 0x07) != transport.version {
		return nil, errors.New("TCPTransport.decode, MAL/TCP Version incompatible")
	}
	sdu := b & 0x1F

	interactionType, interactionStage, err := decodeSDU(sdu)
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot decode SDU: %s", err.Error())
		return nil, err
	}

	serviceArea, err := decoder.DecodeUShort()
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot decode serviceArea: %s", err.Error())
		return nil, err
	}

	service, err := decoder.DecodeUShort()
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot decode service: %s", err.Error())
		return nil, err
	}

	operation, err := decoder.DecodeUShort()
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot decode operation: %s", err.Error())
		return nil, err
	}

	areaVersion, err := decoder.DecodeUOctet()
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot decode areaVersion: %s", err.Error())
		return nil, err
	}

	b, err = decoder.Read()
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot decode flags: %s", err.Error())
		return nil, err
	}
	isError := ((b >> 7) & 0x01) == binary.TRUE
	qos, err := QoSLevelFromOrdinalValue(uint32((b >> 4) & 0x07))
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot decode qos level: %s", err.Error())
		return nil, err
	}
	session, err := SessionTypeFromOrdinalValue(uint32(b & 0xF))
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot decode session type: %s", err.Error())
		return nil, err
	}

	transactionId, err := decoder.DecodeULong()
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot decode transactionId: %s", err.Error())
		return nil, err
	}

	b, err = decoder.Read()
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot decode Transport flags: %s", err.Error())
		return nil, err
	}
	source_flag := ((b >> 7) & 0x01) == binary.TRUE
	destination_flag := ((b >> 6) & 0x01) == binary.TRUE
	priority_flag := ((b >> 5) & 0x01) == binary.TRUE
	timestamp_flag := ((b >> 4) & 0x01) == binary.TRUE
	network_zone_flag := ((b >> 3) & 0x01) == binary.TRUE
	session_name_flag := ((b >> 2) & 0x01) == binary.TRUE
	domain_flag := ((b >> 1) & 0x01) == binary.TRUE
	authentication_id_flag := (b & 0x01) == binary.TRUE

	encodingId, err := decoder.DecodeUOctet()
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot decode encodingId: %s", err.Error())
		return nil, err
	}

	// Skips variable length field
	_, err = decoder.ReadUInt32()
	if err != nil {
		logger.Errorf("TCPTransport.decode, cannot skip variable length: %s", err.Error())
		return nil, err
	}

	// Remaining data are now decoded from PDU using varint.
	decoder.Varint = true

	var urifrom *URI = nil
	if source_flag {
		urifrom, err = decoder.DecodeURI()
		logger.Debugf("TCPTransport.decode, sourceId= %s", *urifrom)
		if err != nil {
			logger.Errorf("TCPTransport.decode, cannot decode sourceId: %s", err.Error())
			return nil, err
		}
		if !strings.HasPrefix(string(*urifrom), MALTCP) {
			// Handle optimized sourceUri transport
			var uri URI = URI(MALTCP_URI + from + "/" + string(*urifrom))
			urifrom = &uri
			logger.Debugf("TCPTransport.decode, sourceId= %s", *urifrom)
		}
	} else {
		urifrom = transport.sourceId
	}

	var urito *URI = nil
	if destination_flag {
		urito, err = decoder.DecodeURI()
		logger.Debugf("TCPTransport.decode, destinationId= %s", *urito)
		if err != nil {
			logger.Errorf("TCPTransport.decode, cannot decode destinationId: %s", err.Error())
			return nil, err
		}
		if !strings.HasPrefix(string(*urito), MALTCP) {
			// Handle optimized destinationUri transport
			var uri URI = URI(string(transport.uri) + "/" + string(*urito))
			urito = &uri
			logger.Debugf("TCPTransport.decode, destinationId= %s", *urito)
		}
	} else {
		urito = transport.destinationId
	}

	var priority *UInteger = nil
	if priority_flag {
		priority, err = decoder.DecodeUInteger()
		if err != nil {
			logger.Errorf("TCPTransport.decode, cannot decode priority: %s", err.Error())
			return nil, err
		}
	} else {
		priority = &transport.dfltPriority
	}

	var timestamp *Time = nil
	if timestamp_flag {
		timestamp, err = decoder.DecodeTime()
		if err != nil {
			logger.Errorf("TCPTransport.decode, cannot decode timestamp: %s", err.Error())
			return nil, err
		}
	} else {
		timestamp = TimeNow()
	}

	var networkZone *Identifier = nil
	if network_zone_flag {
		networkZone, err = decoder.DecodeIdentifier()
		if err != nil {
			logger.Errorf("TCPTransport.decode, cannot decode networkZone: %s", err.Error())
			return nil, err
		}
	} else {
		networkZone = &transport.dfltNetworkZone
	}

	var sessionName *Identifier = nil
	if session_name_flag {
		sessionName, err = decoder.DecodeIdentifier()
		if err != nil {
			logger.Errorf("TCPTransport.decode, cannot decode sessionName: %s", err.Error())
			return nil, err
		}
	} else {
		sessionName = &transport.dfltSessionName
	}

	var domain *IdentifierList = nil
	if domain_flag {
		domain, err = DecodeIdentifierList(decoder)
		if err != nil {
			logger.Errorf("TCPTransport.decode, cannot decode domain: %s", err.Error())
			return nil, err
		}
	} else {
		domain = &transport.dfltDomain
	}

	var authenticationId *Blob = nil
	if authentication_id_flag {
		authenticationId, err = decoder.DecodeBlob()
		if err != nil {
			logger.Errorf("TCPTransport.decode, cannot decode authenticationId: %s", err.Error())
			return nil, err
		}
	} else {
		// Makes a copy to avoid modification of default value
		authenticationId = &transport.dfltAuthenticationId
	}

	// The remaining part of the buffer corresponds to the body part
	// of the message.
	body := NewTCPBody(decoder.Remaining(), false)

	var msg *Message = &Message{
		UriFrom:          urifrom,
		UriTo:            urito,
		AuthenticationId: *authenticationId,
		EncodingId:       *encodingId,
		Timestamp:        *timestamp,
		QoSLevel:         QoSLevel(qos),
		Priority:         *priority,
		Domain:           *domain,
		NetworkZone:      *networkZone,
		Session:          SessionType(session),
		SessionName:      *sessionName,
		InteractionType:  interactionType,
		InteractionStage: interactionStage,
		TransactionId:    *transactionId,
		ServiceArea:      *serviceArea,
		Service:          *service,
		Operation:        *operation,
		AreaVersion:      *areaVersion,
		IsErrorMessage:   Boolean(isError),
		Body:             body,
	}

	return msg, nil
}
