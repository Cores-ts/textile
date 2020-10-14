package client

import (
	"context"

	pb "github.com/textileio/textile/v2/api/billingd/pb"
	"google.golang.org/grpc"
)

// Client provides the client api.
type Client struct {
	c    pb.APIServiceClient
	conn *grpc.ClientConn
}

// NewClient starts the client.
func NewClient(target string, opts ...grpc.DialOption) (*Client, error) {
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		c:    pb.NewAPIServiceClient(conn),
		conn: conn,
	}, nil
}

// Close closes the client's grpc connection and cancels any active requests.
func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) CheckHealth(ctx context.Context) error {
	_, err := c.c.CheckHealth(ctx, &pb.CheckHealthRequest{})
	return err
}

func (c *Client) CreateCustomer(ctx context.Context, email string) (string, error) {
	res, err := c.c.CreateCustomer(ctx, &pb.CreateCustomerRequest{
		Email: email,
	})
	if err != nil {
		return "", err
	}
	return res.CustomerId, nil
}

func (c *Client) GetCustomer(ctx context.Context, customerID string) (*pb.GetCustomerResponse, error) {
	return c.c.GetCustomer(ctx, &pb.GetCustomerRequest{
		CustomerId: customerID,
	})
}

func (c *Client) DeleteCustomer(ctx context.Context, customerID string) error {
	_, err := c.c.DeleteCustomer(ctx, &pb.DeleteCustomerRequest{
		CustomerId: customerID,
	})
	return err
}

func (c *Client) AddCard(ctx context.Context, customerID, token string) error {
	_, err := c.c.AddCard(ctx, &pb.AddCardRequest{
		CustomerId: customerID,
		Token:      token,
	})
	return err
}

func (c *Client) IncStoredData(ctx context.Context, customerID string, incSize int64) (*pb.IncStoredDataResponse, error) {
	return c.c.IncStoredData(ctx, &pb.IncStoredDataRequest{
		CustomerId: customerID,
		IncSize:    incSize,
	})
}

func (c *Client) GetStoredData(ctx context.Context, customerID string) (*pb.GetStoredDataResponse, error) {
	return c.c.GetStoredData(ctx, &pb.GetStoredDataRequest{
		CustomerId: customerID,
	})
}

func (c *Client) IncNetworkEgress(ctx context.Context, customerID string, incSize int64) (*pb.IncNetworkEgressResponse, error) {
	return c.c.IncNetworkEgress(ctx, &pb.IncNetworkEgressRequest{
		CustomerId: customerID,
		IncSize:    incSize,
	})
}

func (c *Client) GetNetworkEgress(ctx context.Context, customerID string) (*pb.GetNetworkEgressResponse, error) {
	return c.c.GetNetworkEgress(ctx, &pb.GetNetworkEgressRequest{
		CustomerId: customerID,
	})
}

func (c *Client) IncInstanceReads(ctx context.Context, customerID string, incCount int64) (*pb.IncInstanceReadsResponse, error) {
	return c.c.IncInstanceReads(ctx, &pb.IncInstanceReadsRequest{
		CustomerId: customerID,
		IncCount:   incCount,
	})
}

func (c *Client) GetInstanceReads(ctx context.Context, customerID string) (*pb.GetInstanceReadsResponse, error) {
	return c.c.GetInstanceReads(ctx, &pb.GetInstanceReadsRequest{
		CustomerId: customerID,
	})
}

func (c *Client) IncInstanceWrites(ctx context.Context, customerID string, incCount int64) (*pb.IncInstanceWritesResponse, error) {
	return c.c.IncInstanceWrites(ctx, &pb.IncInstanceWritesRequest{
		CustomerId: customerID,
		IncCount:   incCount,
	})
}

func (c *Client) GetInstanceWrites(ctx context.Context, customerID string) (*pb.GetInstanceWritesResponse, error) {
	return c.c.GetInstanceWrites(ctx, &pb.GetInstanceWritesRequest{
		CustomerId: customerID,
	})
}