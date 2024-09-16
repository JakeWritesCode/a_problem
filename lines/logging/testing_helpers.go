package logging

import "bytes"

func NewTestBufferedLogger() (*LogrusHandler, *bytes.Buffer) {
	var buf bytes.Buffer
	handler := NewLogrusHandler("debug")
	handler.Logrus.Out = &buf
	return handler, &buf
}
