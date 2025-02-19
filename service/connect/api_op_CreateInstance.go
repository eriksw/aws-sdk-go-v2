// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Initiates an Amazon Connect instance with all the supported channels enabled. It
// does not attach any storage (such as Amazon S3, or Kinesis) or allow for any
// configurations on features such as Contact Lens for Amazon Connect.
func (c *Client) CreateInstance(ctx context.Context, params *CreateInstanceInput, optFns ...func(*Options)) (*CreateInstanceOutput, error) {
	if params == nil {
		params = &CreateInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInstance", params, optFns, addOperationCreateInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstanceInput struct {

	// The type of identity management for your Amazon Connect users.
	//
	// This member is required.
	IdentityManagementType types.DirectoryType

	// Whether your contact center handles incoming contacts.
	//
	// This member is required.
	InboundCallsEnabled *bool

	// Whether your contact center allows outbound calls.
	//
	// This member is required.
	OutboundCallsEnabled *bool

	// The idempotency token.
	ClientToken *string

	// The identifier for the directory.
	DirectoryId *string

	// The name for your instance.
	InstanceAlias *string
}

type CreateInstanceOutput struct {

	// The Amazon Resource Name (ARN) of the instance.
	Arn *string

	// The identifier for the instance.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateInstance{}, middleware.After)
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
	if err = addOpCreateInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "CreateInstance",
	}
}
