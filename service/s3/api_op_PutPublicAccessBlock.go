// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or modifies the PublicAccessBlock configuration for an Amazon S3 bucket.
// To use this operation, you must have the s3:PutBucketPublicAccessBlock
// permission. For more information about Amazon S3 permissions, see Specifying
// Permissions in a Policy
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html).
// When Amazon S3 evaluates the PublicAccessBlock configuration for a bucket or an
// object, it checks the PublicAccessBlock configuration for both the bucket (or
// the bucket that contains the object) and the bucket owner's account. If the
// PublicAccessBlock configurations are different between the bucket and the
// account, Amazon S3 uses the most restrictive combination of the bucket-level and
// account-level settings. For more information about when Amazon S3 considers a
// bucket or an object public, see The Meaning of "Public"
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status).
// Related Resources
//
// * GetPublicAccessBlock
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html)
//
// *
// DeletePublicAccessBlock
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html)
//
// *
// GetBucketPolicyStatus
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicyStatus.html)
//
// *
// Using Amazon S3 Block Public Access
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html)
func (c *Client) PutPublicAccessBlock(ctx context.Context, params *PutPublicAccessBlockInput, optFns ...func(*Options)) (*PutPublicAccessBlockOutput, error) {
	if params == nil {
		params = &PutPublicAccessBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPublicAccessBlock", params, optFns, addOperationPutPublicAccessBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPublicAccessBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPublicAccessBlockInput struct {

	// The name of the Amazon S3 bucket whose PublicAccessBlock configuration you want
	// to set.
	//
	// This member is required.
	Bucket *string

	// The PublicAccessBlock configuration that you want to apply to this Amazon S3
	// bucket. You can enable the configuration options in any combination. For more
	// information about when Amazon S3 considers a bucket or object public, see The
	// Meaning of "Public"
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	PublicAccessBlockConfiguration *types.PublicAccessBlockConfiguration

	// The MD5 hash of the PutPublicAccessBlock request body. For requests made using
	// the AWS Command Line Interface (CLI) or AWS SDKs, this field is calculated
	// automatically.
	ContentMD5 *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type PutPublicAccessBlockOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutPublicAccessBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutPublicAccessBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutPublicAccessBlock{}, middleware.After)
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
	if err = addOpPutPublicAccessBlockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPublicAccessBlock(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutPublicAccessBlockUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutPublicAccessBlock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutPublicAccessBlock",
	}
}

// getPutPublicAccessBlockBucketMember returns a pointer to string denoting a
// provided bucket member valueand a boolean indicating if the input has a modeled
// bucket name,
func getPutPublicAccessBlockBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutPublicAccessBlockInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutPublicAccessBlockUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutPublicAccessBlockBucketMember,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		SupportsAccelerate:      true,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
