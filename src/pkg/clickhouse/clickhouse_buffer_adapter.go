package clickhouse

import clickhousebuffer "github.com/zikwall/clickhouse-buffer"

type BufferAdapter struct {
	chBuffer clickhousebuffer.Client
}

func NewClickhouseBufferAdapter(bufferClient clickhousebuffer.Client) *BufferAdapter {
	return &BufferAdapter{
		chBuffer: bufferClient,
	}
}

func (cba *BufferAdapter) Client() clickhousebuffer.Client {
	return cba.chBuffer
}

func (cba *BufferAdapter) Drop() error {
	cba.chBuffer.Close()
	return nil
}

func (cba *BufferAdapter) DropMsg() string {
	return "close clickhouse buffer adapter"
}
