/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * io/writer_test.go                                      *
 *                                                        *
 * hprose writer test for Go.                             *
 *                                                        *
 * LastModified: Aug 17, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"bytes"
	"strconv"
	"testing"
)

func testSerializeNil(t *testing.T, writer *Writer, b *bytes.Buffer) {
	b.Truncate(0)
	err := writer.Serialize(nil)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "n" {
		t.Error(b.String())
	}
}
func testSerializeTrue(t *testing.T, writer *Writer, b *bytes.Buffer) {
	b.Truncate(0)
	err := writer.Serialize(true)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "t" {
		t.Error(b.String())
	}
}
func testSerializeFalse(t *testing.T, writer *Writer, b *bytes.Buffer) {
	b.Truncate(0)
	err := writer.Serialize(false)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "f" {
		t.Error(b.String())
	}
}
func testSerializeDigit(t *testing.T, writer *Writer, b *bytes.Buffer) {
	for i := 0; i <= 9; i++ {
		b.Truncate(0)
		err := writer.Serialize(i)
		if err != nil {
			t.Error(err.Error())
		}
		if b.String() != strconv.Itoa(i) {
			t.Error(b.String())
		}
	}
}

func TestSerialize(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, false)
	testSerializeNil(t, writer, b)
	testSerializeTrue(t, writer, b)
	testSerializeFalse(t, writer, b)
	testSerializeDigit(t, writer, b)

}
