/**
 * MIT License
 *
 * Copyright (c) 2017 - 2018 CNES
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
package mal

import (
	"net/url"
	"strconv"
)

var (
	NULL_URI *URI = nil
)

// ################################################################################
// Defines MAL URI type
// ################################################################################

type URI string

var (
	NullURI *URI = nil
)

func NewURI(s string) *URI {
	var val URI = URI(s)
	return &val
}

// ================================================================================
// Defines MAL URI type as a MAL Attribute

func (uri *URI) attribute() Attribute {
	return uri
}

// ================================================================================
// Defines MAL URI type as a MAL Element

const MAL_URI_TYPE_SHORT_FORM Integer = 0x12
const MAL_URI_SHORT_FORM Long = 0x1000001000012

// Registers MAL URI type for polymorpsism handling
func init() {
	RegisterMALElement(MAL_URI_SHORT_FORM, NullURI)
}

// Returns the absolute short form of the element type.
func (*URI) GetShortForm() Long {
	return MAL_URI_SHORT_FORM
}

// Returns the number of the area this element type belongs to.
func (*URI) GetAreaNumber() UShort {
	return MAL_ATTRIBUTE_AREA_NUMBER
}

// Returns the version of the area this element type belongs to.
func (*URI) GetAreaVersion() UOctet {
	return MAL_ATTRIBUTE_AREA_VERSION
}

// Returns the number of the service this element type belongs to.
func (*URI) GetServiceNumber() UShort {
	return MAL_ATTRIBUTE_AREA_SERVICE_NUMBER
}

// Returns the relative short form of the element type.
func (*URI) GetTypeShortForm() Integer {
	return MAL_URI_TYPE_SHORT_FORM
}

// Encodes this element using the supplied encoder.
// @param encoder The encoder to use, must not be null.
func (uri *URI) Encode(encoder Encoder) error {
	return encoder.EncodeURI(uri)
}

// Decodes an instance of this element type using the supplied decoder.
// @param decoder The decoder to use, must not be null.
// @return the decoded instance, may be not the same instance as this Element.
func (uri *URI) Decode(decoder Decoder) (Element, error) {
	return decoder.DecodeURI()
}

// The method allows the creation of an element in a generic way, i.e., using the MAL Element polymorphism.
func (uri *URI) CreateElement() Element {
	return NewURI("")
}

func (u *URI) IsNull() bool {
	if u == nil {
		return true
	} else {
		return false
	}
}

func (*URI) Null() Element {
	return NullURI
}

func (uri *URI) ToURL() (*url.URL, error) {
	return url.Parse(string(*uri))
}

func (uri *URI) GetHostname() *String {
	u, err := uri.ToURL()
	if err != nil {
		logger.Errorf("URI.GetHostname: cannot parse %s", uri)
		return NullString
	}
	return NewString(u.Hostname())

}

func (uri *URI) GetPort() int {
	u, err := uri.ToURL()
	if err != nil {
		logger.Errorf("URI.GetPort: cannot parse %s", uri)
		return -1
	}
	port, err := strconv.Atoi(u.Port())
	return port
}

func (uri *URI) GetTransport() *String {
	u, err := uri.ToURL()
	if err != nil {
		logger.Errorf("URI.GetTransport: cannot parse %s", uri)
		return NullString
	}
	return NewString(u.Scheme)
}

func (uri *URI) GetService() *String {
	u, err := uri.ToURL()
	if err != nil {
		logger.Errorf("URI.GetService: cannot parse %s", uri)
		return NullString
	}
	return NewString(u.Path) // RawPath ?
}
