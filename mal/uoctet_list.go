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
// Defines MAL UOctetList type
// ################################################################################

type UOctetList []*UOctet

var (
	NullUOctetList *UOctetList = nil
)

func NewUOctetList(size int) *UOctetList {
	var list UOctetList = UOctetList(make([]*UOctet, size))
	return &list
}

// ================================================================================
// Defines MAL UOctetList type as an ElementList

func (list *UOctetList) Size() int {
	if list != nil {
		return len(*list)
	}
	return -1
}

func (list *UOctetList) GetElementAt(i int) Element {
	if list != nil {
		if i < list.Size() {
			return (*list)[i]
		}
		return nil
	}
	return nil
}

func (list *UOctetList) AppendElement(element Element) {
	if list != nil {
		*list = append(*list, element.(*UOctet))
	}
}

// ================================================================================
// Defines MAL UOctetList type as a MAL Composite

func (list *UOctetList) Composite() Composite {
	return list
}

// ================================================================================
// Defines MAL UOctet type as a MAL Element

const MAL_UOCTET_LIST_TYPE_SHORT_FORM Integer = -0x08
const MAL_UOCTET_LIST_SHORT_FORM Long = 0x1000001FFFFF8

// Registers MAL UOctetList type for polymorpsism handling
func init() {
	RegisterMALElement(MAL_UOCTET_LIST_SHORT_FORM, NullUOctetList)
}

// Returns the absolute short form of the element type.
func (*UOctetList) GetShortForm() Long {
	return MAL_UOCTET_LIST_SHORT_FORM
}

// Returns the number of the area this element type belongs to.
func (*UOctetList) GetAreaNumber() UShort {
	return MAL_ATTRIBUTE_AREA_NUMBER
}

// Returns the version of the area this element type belongs to.
func (*UOctetList) GetAreaVersion() UOctet {
	return MAL_ATTRIBUTE_AREA_VERSION
}

// Returns the number of the service this element type belongs to.
func (*UOctetList) GetServiceNumber() UShort {
	return MAL_ATTRIBUTE_AREA_SERVICE_NUMBER
}

// Return the relative short form of the element type.
func (*UOctetList) GetTypeShortForm() Integer {
	//	return MAL_UOCTET_TYPE_SHORT_FORM & 0x01FFFF00
	return MAL_UOCTET_LIST_TYPE_SHORT_FORM
}

// Encodes this element using the supplied encoder.
// @param encoder The encoder to use, must not be null.
func (list *UOctetList) Encode(encoder Encoder) error {
	err := encoder.EncodeUInteger(NewUInteger(uint32(len([]*UOctet(*list)))))
	if err != nil {
		return err
	}
	for _, e := range []*UOctet(*list) {
		err = encoder.EncodeNullableUOctet(e)
		if err != nil {
			return err
		}
	}
	return nil
}

// Decodes an instance of this element type using the supplied decoder.
// @param decoder The decoder to use, must not be null.
// @return the decoded instance, may be not the same instance as this Element.
func (list *UOctetList) Decode(decoder Decoder) (Element, error) {
	return DecodeUOctetList(decoder)
}

// Decodes an instance of UOctetList using the supplied decoder.
// @param decoder The decoder to use, must not be null.
// @return the decoded UOctetList instance.
func DecodeUOctetList(decoder Decoder) (*UOctetList, error) {
	size, err := decoder.DecodeUInteger()
	if err != nil {
		return nil, err
	}
	list := UOctetList(make([]*UOctet, int(*size)))
	for i := 0; i < len(list); i++ {
		list[i], err = decoder.DecodeNullableUOctet()
		if err != nil {
			return nil, err
		}
	}
	return &list, nil
}

// The method allows the creation of an element in a generic way, i.e., using the MAL Element polymorphism.
func (list *UOctetList) CreateElement() Element {
	return NewUOctetList(0)
}

func (list *UOctetList) IsNull() bool {
	return list == nil
}

func (*UOctetList) Null() Element {
	return NullUOctetList
}
