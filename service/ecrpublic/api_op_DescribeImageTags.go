// Code generated by smithy-go-codegen DO NOT EDIT.

package ecrpublic

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the image tag details for a repository in a public registry.
func (c *Client) DescribeImageTags(ctx context.Context, params *DescribeImageTagsInput, optFns ...func(*Options)) (*DescribeImageTagsOutput, error) {
	if params == nil {
		params = &DescribeImageTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImageTags", params, optFns, addOperationDescribeImageTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImageTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImageTagsInput struct {

	// The name of the repository that contains the image tag details to describe.
	//
	// This member is required.
	RepositoryName *string

	// The maximum number of repository results returned by DescribeImageTags in
	// paginated output. When this parameter is used, DescribeImageTags only returns
	// maxResults results in a single page along with a nextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// DescribeImageTags request with the returned nextToken value. This value can be
	// between 1 and 1000. If this parameter is not used, then DescribeImageTags
	// returns up to 100 results and a nextToken value, if applicable. This option
	// cannot be used when you specify images with imageIds.
	MaxResults *int32

	// The nextToken value returned from a previous paginated DescribeImageTags request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	// This option cannot be used when you specify images with imageIds.
	NextToken *string

	// The AWS account ID associated with the public registry that contains the
	// repository in which to describe images. If you do not specify a registry, the
	// default public registry is assumed.
	RegistryId *string
}

type DescribeImageTagsOutput struct {

	// The image tag details for the images in the requested repository.
	ImageTagDetails []types.ImageTagDetail

	// The nextToken value to include in a future DescribeImageTags request. When the
	// results of a DescribeImageTags request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeImageTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImageTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImageTags{}, middleware.After)
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
	if err = addOpDescribeImageTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImageTags(options.Region), middleware.Before); err != nil {
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

// DescribeImageTagsAPIClient is a client that implements the DescribeImageTags
// operation.
type DescribeImageTagsAPIClient interface {
	DescribeImageTags(context.Context, *DescribeImageTagsInput, ...func(*Options)) (*DescribeImageTagsOutput, error)
}

var _ DescribeImageTagsAPIClient = (*Client)(nil)

// DescribeImageTagsPaginatorOptions is the paginator options for DescribeImageTags
type DescribeImageTagsPaginatorOptions struct {
	// The maximum number of repository results returned by DescribeImageTags in
	// paginated output. When this parameter is used, DescribeImageTags only returns
	// maxResults results in a single page along with a nextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// DescribeImageTags request with the returned nextToken value. This value can be
	// between 1 and 1000. If this parameter is not used, then DescribeImageTags
	// returns up to 100 results and a nextToken value, if applicable. This option
	// cannot be used when you specify images with imageIds.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeImageTagsPaginator is a paginator for DescribeImageTags
type DescribeImageTagsPaginator struct {
	options   DescribeImageTagsPaginatorOptions
	client    DescribeImageTagsAPIClient
	params    *DescribeImageTagsInput
	nextToken *string
	firstPage bool
}

// NewDescribeImageTagsPaginator returns a new DescribeImageTagsPaginator
func NewDescribeImageTagsPaginator(client DescribeImageTagsAPIClient, params *DescribeImageTagsInput, optFns ...func(*DescribeImageTagsPaginatorOptions)) *DescribeImageTagsPaginator {
	options := DescribeImageTagsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeImageTagsInput{}
	}

	return &DescribeImageTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeImageTagsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeImageTags page.
func (p *DescribeImageTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeImageTagsOutput, error) {
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

	result, err := p.client.DescribeImageTags(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeImageTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr-public",
		OperationName: "DescribeImageTags",
	}
}
