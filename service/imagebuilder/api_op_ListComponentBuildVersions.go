// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of component build versions for the specified semantic version.
func (c *Client) ListComponentBuildVersions(ctx context.Context, params *ListComponentBuildVersionsInput, optFns ...func(*Options)) (*ListComponentBuildVersionsOutput, error) {
	if params == nil {
		params = &ListComponentBuildVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComponentBuildVersions", params, optFns, addOperationListComponentBuildVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComponentBuildVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComponentBuildVersionsInput struct {

	// The component version Amazon Resource Name (ARN) whose versions you want to
	// list.
	//
	// This member is required.
	ComponentVersionArn *string

	// The maximum items to return in a request.
	MaxResults int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string
}

type ListComponentBuildVersionsOutput struct {

	// The list of component summaries for the specified semantic version.
	ComponentSummaryList []types.ComponentSummary

	// The next token used for paginated responses. When this is not empty, there are
	// additional elements that the service has not included in this request. Use this
	// token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListComponentBuildVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListComponentBuildVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListComponentBuildVersions{}, middleware.After)
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
	if err = addOpListComponentBuildVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComponentBuildVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListComponentBuildVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListComponentBuildVersions",
	}
}
