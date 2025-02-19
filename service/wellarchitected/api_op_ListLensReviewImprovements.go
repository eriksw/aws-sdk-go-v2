// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List lens review improvements.
func (c *Client) ListLensReviewImprovements(ctx context.Context, params *ListLensReviewImprovementsInput, optFns ...func(*Options)) (*ListLensReviewImprovementsOutput, error) {
	if params == nil {
		params = &ListLensReviewImprovementsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLensReviewImprovements", params, optFns, addOperationListLensReviewImprovementsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLensReviewImprovementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to list lens review improvements.
type ListLensReviewImprovementsInput struct {

	// The alias of the lens, for example, serverless. Each lens is identified by its
	// LensSummary$LensAlias.
	//
	// This member is required.
	LensAlias *string

	// The ID assigned to the workload. This ID is unique within an AWS Region.
	//
	// This member is required.
	WorkloadId *string

	// The maximum number of results to return for this request.
	MaxResults int32

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The ID used to identify a pillar, for example, security. A pillar is identified
	// by its PillarReviewSummary$PillarId.
	PillarId *string
}

// Output of a list lens review improvements call.
type ListLensReviewImprovementsOutput struct {

	// List of improvement summaries of lens review in a workload.
	ImprovementSummaries []types.ImprovementSummary

	// The alias of the lens, for example, serverless. Each lens is identified by its
	// LensSummary$LensAlias.
	LensAlias *string

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The ID assigned to the workload. This ID is unique within an AWS Region.
	WorkloadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListLensReviewImprovementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLensReviewImprovements{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLensReviewImprovements{}, middleware.After)
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
	if err = addOpListLensReviewImprovementsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLensReviewImprovements(options.Region), middleware.Before); err != nil {
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

// ListLensReviewImprovementsAPIClient is a client that implements the
// ListLensReviewImprovements operation.
type ListLensReviewImprovementsAPIClient interface {
	ListLensReviewImprovements(context.Context, *ListLensReviewImprovementsInput, ...func(*Options)) (*ListLensReviewImprovementsOutput, error)
}

var _ ListLensReviewImprovementsAPIClient = (*Client)(nil)

// ListLensReviewImprovementsPaginatorOptions is the paginator options for
// ListLensReviewImprovements
type ListLensReviewImprovementsPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLensReviewImprovementsPaginator is a paginator for
// ListLensReviewImprovements
type ListLensReviewImprovementsPaginator struct {
	options   ListLensReviewImprovementsPaginatorOptions
	client    ListLensReviewImprovementsAPIClient
	params    *ListLensReviewImprovementsInput
	nextToken *string
	firstPage bool
}

// NewListLensReviewImprovementsPaginator returns a new
// ListLensReviewImprovementsPaginator
func NewListLensReviewImprovementsPaginator(client ListLensReviewImprovementsAPIClient, params *ListLensReviewImprovementsInput, optFns ...func(*ListLensReviewImprovementsPaginatorOptions)) *ListLensReviewImprovementsPaginator {
	options := ListLensReviewImprovementsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListLensReviewImprovementsInput{}
	}

	return &ListLensReviewImprovementsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLensReviewImprovementsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListLensReviewImprovements page.
func (p *ListLensReviewImprovementsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLensReviewImprovementsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListLensReviewImprovements(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLensReviewImprovements(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "ListLensReviewImprovements",
	}
}
