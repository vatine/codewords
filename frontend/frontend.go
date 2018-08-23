package frontend

import (
	"errors"
	"fmt"
	"sync"

	cw "github.com/vatine/codewords"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Some pre-fetched codewords, for quicker dispatch
var cache []string
var cacheLock sync.Mutex

// Global client variable
var client cw.CodewordsServiceClient

// Various settings
var cacheMinSize int
var cacheMaxSize int

// Set the minimum size for the cache, if it ever gets smaller,
// request enough to fill it fully as we dispatch the next code word.
func SetMinimumCacheSize(n int) {
	cacheMinSize = n
}

// Set the maximum size for the cache, we should never try to fill it
// fuller than this
func SetMaximumCacheSize(n int) {
	cacheMaxSize = n
}

func init() {
	SetMinimumCacheSize(1)
	SetMaximumCacheSize(5)
	cache = []string{}
}

// Refill the cache, expects to be called with the cacheLock held,
// unsafe to call without that being the case.
func refillCache(ctx context.Context) error {
	current := len(cache)
	needed := cacheMaxSize - current
	if needed <= 0 {
		needed = 1
	}
	req := cw.CodewordsRequest{Count: int32(needed)}
	r, err := client.GetCodewords(ctx, &req)

	if err != nil {
		fmt.Printf("refillCache error: %s\n", err)
		return err
	}
	for _, word := range r.Words {
		// Despite errors, we may still have words
		cache = append(cache, word)
	}
	return err
}

func NextCodeword(ctx context.Context) (string, error) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	if len(cache) <= cacheMinSize {
		// The cache will be too small after this fetch,
		// refill it now, avoids having to deal with an empty
		// cache in a nice and consistent manner.
		err := refillCache(ctx)
		if err != nil {
			// Strictly speaking, we should log this, but...
		}
	}
	if len(cache) == 0 {
		// Should never happen, but...
		return "", errors.New("Nothing in the cache.")
	}

	rv := cache[0]
	cache = cache[1:]
	return rv, nil
}

// Initiate the connection to the backend, set the client internally,
// return the grpc.ClientConn, so the main server can close everythong
// down.
func Connect(endpoint string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	client = cw.NewCodewordsServiceClient(conn)

	return conn, nil
}
