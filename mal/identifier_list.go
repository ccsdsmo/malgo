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
// Defines MAL IdentifierList type
// ################################################################################

type IdentifierList []*Identifier

var (
	NullIdentifierList *IdentifierList = nil
)

func NewIdentifierList(size int) *IdentifierList {
	var list IdentifierList = IdentifierList(make([]*Identifier, size))
	return &list
}

// ================================================================================
// Defines MAL IdentifierList type as an ElementList

func (list *IdentifierList) Size() int {
	if list != nil {
		return len(*list)
	}
	return -1
}

func (list *IdentifierList) GetElementAt(i int) Element {
	if list != nil {
		if i <= list.Size() {
			return (*list)[i]
		}
		return nil
	}
	return nil
}

// ================================================================================
// Defines MAL IdentifierList type as a MAL Composite

func (list *IdentifierList) Composite() Composite {
	return list
}

// ================================================================================
// Defines MAL IdentifierList type as a MAL Element

const MAL_IDENTIFIER_LIST_TYPE_SHORT_FORM Integer = -0x06
const MAL_IDENTIFIER_LIST_SHORT_FORM Long = 0x1000001FFFFFA

// Registers MAL IdentifierList type for polymorpsism handling
func init() {
	RegisterMALElement(MAL_IDENTIFIER_LIST_SHORT_FORM, NullIdentifierList)
}

// Returns the absolute short form of the element type.
func (*IdentifierList) GetShortForm() Long {
	return MAL_IDENTIFIER_LIST_SHORT_FORM
}

// Returns the number of the area this element type belongs to.
func (*IdentifierList) GetAreaNumber() UShort {
	return MAL_ATTRIBUTE_AREA_NUMBER
}

// Returns the version of the area this element type belongs to.
func (*IdentifierList) GetAreaVersion() UOctet {
	return MAL_ATTRIBUTE_AREA_VERSION
}

// Returns the number of the service this element type belongs to.
func (*IdentifierList) GetServiceNumber() UShort {
	return MAL_ATTRIBUTE_AREA_SERVICE_NUMBER
}

// Returns the relative short form of the element type.
func (*IdentifierList) GetTypeShortForm() Integer {
	//	return MAL_IDENTIFIER_TYPE_SHORT_FORM & 0x01FFFF00
	return MAL_IDENTIFIER_LIST_TYPE_SHORT_FORM
}

// Encodes this element using the supplied encoder.
// @param encoder The encoder to use, must not be null.
func (list *IdentifierList) Encode(encoder Encoder) error {
	err := encoder.EncodeUInteger(NewUInteger(uint32(len([]*Identifier(*list)))))
	if err != nil {
		return err
	}
	for _, e := range []*Identifier(*list) {
		encoder.EncodeNullableIdentifier(e)
	}
	return nil
}

// Decodes an instance of this element type using the supplied decoder.
// @param decoder The decoder to use, must not be null.
// @return the decoded instance, may be not the same instance as this Element.
func (*IdentifierList) Decode(decoder Decoder) (Element, error) {
	return DecodeIdentifierList(decoder)
}

// Decodes an instance of IdentifierList using the supplied decoder.
// @param decoder The decoder to use, must not be null.
// @return the decoded IdentifierList instance.
func DecodeIdentifierList(decoder Decoder) (*IdentifierList, error) {
	size, err := decoder.DecodeUInteger()
	if err != nil {
		return nil, err
	}
	list := IdentifierList(make([]*Identifier, int(*size)))
	for i := 0; i < len(list); i++ {
		list[i], err = decoder.DecodeNullableIdentifier()
		if err != nil {
			return nil, err
		}
	}
	return &list, nil
}

// The method allows the creation of an element in a generic way, i.e., using the MAL Element polymorphism.
func (*IdentifierList) CreateElement() Element {
	return NewIdentifierList(0)
}

func (list *IdentifierList) IsNull() bool {
	return list == nil
}

func (*IdentifierList) Null() Element {
	return NullIdentifierList
}
