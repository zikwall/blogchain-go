package clickhouse

import "github.com/zikwall/clickhouse-buffer/src/api"

type BufferAdapter struct {
	chBuffer api.Client
}

func NewClickhouseBufferAdapter(bufferClient api.Client) *BufferAdapter {
	return &BufferAdapter{
		chBuffer: bufferClient,
	}
}

func (cba *BufferAdapter) Client() api.Client {
	return cba.chBuffer
}

func (cba *BufferAdapter) Close() error {
	cba.chBuffer.Close()
	return nil
}

func (cba *BufferAdapter) CloseMessage() string {
	return "close clickhouse buffer adapter"
}
