// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables automatic rotation of the key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) for the
// specified symmetric customer master key (CMK). You cannot enable automatic
// rotation of asymmetric CMKs, CMKs with imported key material, or CMKs in a
// custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
// The CMK that you use for this operation must be in a compatible key state. For
// details, see How Key State Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a CMK in a different AWS account. Required
// permissions: kms:EnableKeyRotation
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations:
//
// * DisableKeyRotation
//
// * GetKeyRotationStatus
func (c *Client) EnableKeyRotation(ctx context.Context, params *EnableKeyRotationInput, optFns ...func(*Options)) (*EnableKeyRotationOutput, error) {
	if params == nil {
		params = &EnableKeyRotationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableKeyRotation", params, optFns, addOperationEnableKeyRotationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableKeyRotationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableKeyRotationInput struct {

	// Identifies a symmetric customer master key (CMK). You cannot enable automatic
	// rotation of asymmetric CMKs, CMKs with imported key material, or CMKs in a
	// custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK. For example:
	//
	// *
	// Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
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

type EnableKeyRotationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableKeyRotationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableKeyRotation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableKeyRotation{}, middleware.After)
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
	if err = addOpEnableKeyRotationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableKeyRotation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableKeyRotation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "EnableKeyRotation",
	}
}
