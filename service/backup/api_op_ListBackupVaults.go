// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of recovery point storage containers along with information about
// them.
func (c *Client) ListBackupVaults(ctx context.Context, params *ListBackupVaultsInput, optFns ...func(*Options)) (*ListBackupVaultsOutput, error) {
	if params == nil {
		params = &ListBackupVaultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackupVaults", params, optFns, addOperationListBackupVaultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackupVaultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBackupVaultsInput struct {

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string
}

type ListBackupVaultsOutput struct {

	// An array of backup vault list members containing vault metadata, including
	// Amazon Resource Name (ARN), display name, creation date, number of saved
	// recovery points, and encryption information if the resources saved in the backup
	// vault are encrypted.
	BackupVaultList []types.BackupVaultListMember

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBackupVaultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackupVaults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackupVaults{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBackupVaults(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListBackupVaultsAPIClient is a client that implements the ListBackupVaults
// operation.
type ListBackupVaultsAPIClient interface {
	ListBackupVaults(context.Context, *ListBackupVaultsInput, ...func(*Options)) (*ListBackupVaultsOutput, error)
}

var _ ListBackupVaultsAPIClient = (*Client)(nil)

// ListBackupVaultsPaginatorOptions is the paginator options for ListBackupVaults
type ListBackupVaultsPaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBackupVaultsPaginator is a paginator for ListBackupVaults
type ListBackupVaultsPaginator struct {
	options   ListBackupVaultsPaginatorOptions
	client    ListBackupVaultsAPIClient
	params    *ListBackupVaultsInput
	nextToken *string
	firstPage bool
}

// NewListBackupVaultsPaginator returns a new ListBackupVaultsPaginator
func NewListBackupVaultsPaginator(client ListBackupVaultsAPIClient, params *ListBackupVaultsInput, optFns ...func(*ListBackupVaultsPaginatorOptions)) *ListBackupVaultsPaginator {
	options := ListBackupVaultsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListBackupVaultsInput{}
	}

	return &ListBackupVaultsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBackupVaultsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListBackupVaults page.
func (p *ListBackupVaultsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBackupVaultsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListBackupVaults(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListBackupVaults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListBackupVaults",
	}
}