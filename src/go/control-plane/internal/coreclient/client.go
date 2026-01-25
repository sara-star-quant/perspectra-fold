package coreclient

import (
    "context"
    "errors"
    "fmt"
    "math"
    "time"

    corepb "github.com/peterz/multidimensional-transformation/control-plane/internal/corepb"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/credentials/insecure"
    "google.golang.org/grpc/status"
)

var (
    ErrEmptyVector      = errors.New("empty vector")
    ErrMismatchedLength = errors.New("mismatched length")
    ErrZeroNorm         = errors.New("zero norm")
)

type Client interface {
    CosineSimilarity(ctx context.Context, a []float64, b []float64) (float64, error)
    Close() error
}

func New(addr string, timeout time.Duration) (Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    conn, err := grpc.DialContext(
        ctx,
        addr,
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithBlock(),
    )
    if err != nil {
        return nil, err
    }

    return &grpcClient{
        conn:    conn,
        client: corepb.NewCoreComputeClient(conn),
        timeout: timeout,
    }, nil
}

func NewLocal() Client {
    return localClient{}
}

type grpcClient struct {
    conn    *grpc.ClientConn
    client  corepb.CoreComputeClient
    timeout time.Duration
}

func (c *grpcClient) CosineSimilarity(ctx context.Context, a []float64, b []float64) (float64, error) {
    if err := validateVectors(a, b); err != nil {
        return 0, err
    }

    ctx, cancel := withTimeout(ctx, c.timeout)
    defer cancel()

    resp, err := c.client.CosineSimilarity(ctx, &corepb.CosineSimilarityRequest{
        A: a,
        B: b,
    })
    if err != nil {
        return 0, mapRPCError(err)
    }

    return resp.Value, nil
}

func (c *grpcClient) Close() error {
    return c.conn.Close()
}

type localClient struct{}

func (localClient) CosineSimilarity(_ context.Context, a []float64, b []float64) (float64, error) {
    if err := validateVectors(a, b); err != nil {
        return 0, err
    }

    var dot float64
    var normA float64
    var normB float64
    for i := range a {
        dot += a[i] * b[i]
        normA += a[i] * a[i]
        normB += b[i] * b[i]
    }
    if normA == 0 || normB == 0 {
        return 0, ErrZeroNorm
    }

    return dot / (math.Sqrt(normA) * math.Sqrt(normB)), nil
}

func (localClient) Close() error {
    return nil
}

func validateVectors(a []float64, b []float64) error {
    if len(a) == 0 || len(b) == 0 {
        return ErrEmptyVector
    }
    if len(a) != len(b) {
        return ErrMismatchedLength
    }
    return nil
}

func withTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
    if ctx == nil {
        return context.WithTimeout(context.Background(), timeout)
    }
    if _, ok := ctx.Deadline(); ok {
        return ctx, func() {}
    }
    return context.WithTimeout(ctx, timeout)
}

func mapRPCError(err error) error {
    statusErr, ok := status.FromError(err)
    if !ok {
        return err
    }

    switch statusErr.Code() {
    case codes.InvalidArgument:
        return fmt.Errorf("invalid input: %w", err)
    case codes.FailedPrecondition:
        return ErrZeroNorm
    case codes.Unavailable:
        return fmt.Errorf("core compute unavailable: %w", err)
    default:
        return err
    }
}
