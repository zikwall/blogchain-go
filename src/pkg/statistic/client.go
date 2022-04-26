package statistic

import (
	"context"
	"time"

	"github.com/zikwall/blogchain/src/protobuf/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type StorageClientImp struct {
	conn   *grpc.ClientConn
	client storage.StorageClient
}

func NewClient(ctx context.Context, listenAddress string) (*StorageClientImp, error) {
	impl := &StorageClientImp{}
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	conn, err := grpc.DialContext(ctx, listenAddress, opts...)
	if err != nil {
		return nil, err
	}

	impl.client = storage.NewStorageClient(conn)
	impl.conn = conn
	return impl, nil
}

func (s *StorageClientImp) Client() storage.StorageClient {
	return s.client
}

func (s *StorageClientImp) Drop() error {
	return s.conn.Close()
}

func (s *StorageClientImp) DropMsg() string {
	return "close storage client connection"
}
