// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a Boolean value that indicates whether automatic rotation of the key
// material
// (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) is
// enabled for the specified customer master key (CMK). You cannot enable automatic
// rotation of asymmetric CMKs, CMKs with imported key material, or CMKs in a
// custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
// The key rotation status for these CMKs is always false. The CMK that you use for
// this operation must be in a compatible key state. For details, see How Key State
// Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
//
// * Disabled: The key rotation status
// does not change when you disable a CMK. However, while the CMK is disabled, AWS
// KMS does not rotate the backing key.
//
// * Pending deletion: While a CMK is pending
// deletion, its key rotation status is false and AWS KMS does not rotate the
// backing key. If you cancel the deletion, the original key rotation status is
// restored.
//
// Cross-account use: Yes. To perform this operation on a CMK in a
// different AWS account, specify the key ARN in the value of the KeyId parameter.
// Required permissions: kms:GetKeyRotationStatus
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations:
//
// * DisableKeyRotation
//
// * EnableKeyRotation
func (c *Client) GetKeyRotationStatus(ctx context.Context, params *GetKeyRotationStatusInput, optFns ...func(*Options)) (*GetKeyRotationStatusOutput, error) {
	if params == nil {
		params = &GetKeyRotationStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetKeyRotationStatus", params, optFns, addOperationGetKeyRotationStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetKeyRotationStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetKeyRotationStatusInput struct {

	// A unique identifier for the customer master key (CMK). Specify the key ID or the
	// Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS
	// account, you must use the key ARN. For example:
	//
	// * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string
}

type GetKeyRotationStatusOutput struct {

	// A Boolean value that specifies whether key rotation is enabled.
	KeyRotationEnabled bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetKeyRotationStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetKeyRotationStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetKeyRotationStatus{}, middleware.After)
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
	if err = addOpGetKeyRotationStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetKeyRotationStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetKeyRotationStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "GetKeyRotationStatus",
	}
}
