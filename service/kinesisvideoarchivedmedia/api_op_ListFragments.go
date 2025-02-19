// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Fragment objects from the specified stream and timestamp range
// within the archived data. Listing fragments is eventually consistent. This means
// that even if the producer receives an acknowledgment that a fragment is
// persisted, the result might not be returned immediately from a request to
// ListFragments. However, results are typically available in less than one second.
// You must first call the GetDataEndpoint API to get an endpoint. Then send the
// ListFragments requests to this endpoint using the --endpoint-url parameter
// (https://docs.aws.amazon.com/cli/latest/reference/). If an error is thrown after
// invoking a Kinesis Video Streams archived media API, in addition to the HTTP
// status code and the response body, it includes the following pieces of
// information:
//
// * x-amz-ErrorType HTTP header – contains a more specific error
// type in addition to what the HTTP status code provides.
//
// * x-amz-RequestId HTTP
// header – if you want to report an issue to AWS, the support team can better
// diagnose the problem if given the Request Id.
//
// Both the HTTP status code and the
// ErrorType header can be utilized to make programmatic decisions about whether
// errors are retry-able and under what conditions, as well as provide information
// on what actions the client programmer might need to take in order to
// successfully try again. For more information, see the Errors section at the
// bottom of this topic, as well as Common Errors
// (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/CommonErrors.html).
func (c *Client) ListFragments(ctx context.Context, params *ListFragmentsInput, optFns ...func(*Options)) (*ListFragmentsOutput, error) {
	if params == nil {
		params = &ListFragmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFragments", params, optFns, addOperationListFragmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFragmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFragmentsInput struct {

	// The name of the stream from which to retrieve a fragment list.
	//
	// This member is required.
	StreamName *string

	// Describes the timestamp range and timestamp origin for the range of fragments to
	// return.
	FragmentSelector *types.FragmentSelector

	// The total number of fragments to return. If the total number of fragments
	// available is more than the value specified in max-results, then a
	// ListFragmentsOutput$NextToken is provided in the output that you can use to
	// resume pagination.
	MaxResults *int32

	// A token to specify where to start paginating. This is the
	// ListFragmentsOutput$NextToken from a previously truncated response.
	NextToken *string
}

type ListFragmentsOutput struct {

	// A list of archived Fragment objects from the stream that meet the selector
	// criteria. Results are in no specific order, even across pages.
	Fragments []types.Fragment

	// If the returned list is truncated, the operation returns this token to use to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFragmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFragments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFragments{}, middleware.After)
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
	if err = addOpListFragmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFragments(options.Region), middleware.Before); err != nil {
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

// ListFragmentsAPIClient is a client that implements the ListFragments operation.
type ListFragmentsAPIClient interface {
	ListFragments(context.Context, *ListFragmentsInput, ...func(*Options)) (*ListFragmentsOutput, error)
}

var _ ListFragmentsAPIClient = (*Client)(nil)

// ListFragmentsPaginatorOptions is the paginator options for ListFragments
type ListFragmentsPaginatorOptions struct {
	// The total number of fragments to return. If the total number of fragments
	// available is more than the value specified in max-results, then a
	// ListFragmentsOutput$NextToken is provided in the output that you can use to
	// resume pagination.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFragmentsPaginator is a paginator for ListFragments
type ListFragmentsPaginator struct {
	options   ListFragmentsPaginatorOptions
	client    ListFragmentsAPIClient
	params    *ListFragmentsInput
	nextToken *string
	firstPage bool
}

// NewListFragmentsPaginator returns a new ListFragmentsPaginator
func NewListFragmentsPaginator(client ListFragmentsAPIClient, params *ListFragmentsInput, optFns ...func(*ListFragmentsPaginatorOptions)) *ListFragmentsPaginator {
	options := ListFragmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListFragmentsInput{}
	}

	return &ListFragmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFragmentsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListFragments page.
func (p *ListFragmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFragmentsOutput, error) {
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

	result, err := p.client.ListFragments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFragments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "ListFragments",
	}
}
