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

// ################################################################################
// Defines MAL URIList type
// ################################################################################

type URIList []*URI

var (
	NullURIList *URIList = nil
)

func NewURIList(size int) *URIList {
	var list URIList = URIList(make([]*URI, size))
	return &list
}

// ================================================================================
// Defines MAL URIList type as an ElementList

func (list *URIList) Size() int {
	if list != nil {
		return len(*list)
	}
	return -1
}

func (list *URIList) GetElementAt(i int) Element {
	if list != nil {
		if i < list.Size() {
			return (*list)[i]
		}
		return nil
	}
	return nil
}

func (list *URIList) AppendElement(element Element) {
	if list != nil {
		*list = append(*list, element.(*URI))
	}
}

// ================================================================================
// Defines MAL URIList type as a MAL Composite

func (list *URIList) Composite() Composite {
	return list
}

// ================================================================================
// Defines MAL URIList type as a MAL Element

const MAL_URI_LIST_TYPE_SHORT_FORM Integer = -0x12
const MAL_URI_LIST_SHORT_FORM Long = 0x1000001FFFFEE

// Registers MAL URIList type for polymorpsism handling
func init() {
	RegisterMALElement(MAL_URI_LIST_SHORT_FORM, NullURIList)
}

// Returns the absolute short form of the element type.
func (*URIList) GetShortForm() Long {
	return MAL_URI_LIST_SHORT_FORM
}

// Returns the number of the area this element type belongs to.
func (*URIList) GetAreaNumber() UShort {
	return MAL_ATTRIBUTE_AREA_NUMBER
}

// Returns the version of the area this element type belongs to.
func (*URIList) GetAreaVersion() UOctet {
	return MAL_ATTRIBUTE_AREA_VERSION
}

// Returns the number of the service this element type belongs to.
func (*URIList) GetServiceNumber() UShort {
	return MAL_ATTRIBUTE_AREA_SERVICE_NUMBER
}

// Returns the relative short form of the element type.
func (*URIList) GetTypeShortForm() Integer {
	//	return MAL_URI_TYPE_SHORT_FORM & 0x01FFFF00
	return MAL_URI_LIST_TYPE_SHORT_FORM
}

// Encodes this element using the supplied encoder.
// @param encoder The encoder to use, must not be null.
func (list *URIList) Encode(encoder Encoder) error {
	err := encoder.EncodeUInteger(NewUInteger(uint32(len([]*URI(*list)))))
	if err != nil {
		return err
	}
	for _, e := range []*URI(*list) {
		err = encoder.EncodeNullableURI(e)
		if err != nil {
			return err
		}
	}
	return nil
}

// Decodes an instance of this element type using the supplied decoder.
// @param decoder The decoder to use, must not be null.
// @return the decoded instance, may be not the same instance as this Element.
func (list *URIList) Decode(decoder Decoder) (Element, error) {
	return DecodeURIList(decoder)
}

// Decodes an instance of URIList using the supplied decoder.
// @param decoder The decoder to use, must not be null.
// @return the decoded URIList instance.
func DecodeURIList(decoder Decoder) (*URIList, error) {
	size, err := decoder.DecodeUInteger()
	if err != nil {
		return nil, err
	}
	list := URIList(make([]*URI, int(*size)))
	for i := 0; i < len(list); i++ {
		list[i], err = decoder.DecodeNullableURI()
		if err != nil {
			return nil, err
		}
	}
	return &list, nil
}

// The method allows the creation of an element in a generic way, i.e., using the MAL Element polymorphism.
func (list *URIList) CreateElement() Element {
	return NewURIList(0)
}

func (list *URIList) IsNull() bool {
	return list == nil
}

func (*URIList) Null() Element {
	return NullURIList
}
