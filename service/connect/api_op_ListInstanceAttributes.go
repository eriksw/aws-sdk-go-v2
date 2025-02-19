// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Returns a paginated list of all attribute types for the given instance.
func (c *Client) ListInstanceAttributes(ctx context.Context, params *ListInstanceAttributesInput, optFns ...func(*Options)) (*ListInstanceAttributesOutput, error) {
	if params == nil {
		params = &ListInstanceAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstanceAttributes", params, optFns, addOperationListInstanceAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstanceAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstanceAttributesInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The maximimum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string
}

type ListInstanceAttributesOutput struct {

	// The attribute types.
	Attributes []types.Attribute

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListInstanceAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInstanceAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInstanceAttributes{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addOpListInstanceAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInstanceAttributes(options.Region), middleware.Before); err != nil {
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

// ListInstanceAttributesAPIClient is a client that implements the
// ListInstanceAttributes operation.
type ListInstanceAttributesAPIClient interface {
	ListInstanceAttributes(context.Context, *ListInstanceAttributesInput, ...func(*Options)) (*ListInstanceAttributesOutput, error)
}

var _ ListInstanceAttributesAPIClient = (*Client)(nil)

// ListInstanceAttributesPaginatorOptions is the paginator options for
// ListInstanceAttributes
type ListInstanceAttributesPaginatorOptions struct {
	// The maximimum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInstanceAttributesPaginator is a paginator for ListInstanceAttributes
type ListInstanceAttributesPaginator struct {
	options   ListInstanceAttributesPaginatorOptions
	client    ListInstanceAttributesAPIClient
	params    *ListInstanceAttributesInput
	nextToken *string
	firstPage bool
}

// NewListInstanceAttributesPaginator returns a new ListInstanceAttributesPaginator
func NewListInstanceAttributesPaginator(client ListInstanceAttributesAPIClient, params *ListInstanceAttributesInput, optFns ...func(*ListInstanceAttributesPaginatorOptions)) *ListInstanceAttributesPaginator {
	options := ListInstanceAttributesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListInstanceAttributesInput{}
	}

	return &ListInstanceAttributesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInstanceAttributesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListInstanceAttributes page.
func (p *ListInstanceAttributesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInstanceAttributesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListInstanceAttributes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInstanceAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListInstanceAttributes",
	}
}
