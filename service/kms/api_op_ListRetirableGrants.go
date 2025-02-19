// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all grants in which the specified principal is the RetiringPrincipal in
// the grant. You can specify any principal in your AWS account. The grants that
// are returned include grants for CMKs in your AWS account and other AWS accounts.
// You might use this operation to determine which grants you may retire. To retire
// a grant, use the RetireGrant operation. Cross-account use: You must specify a
// principal in your AWS account. However, this operation can return grants in any
// AWS account. You do not need kms:ListRetirableGrants permission (or any other
// additional permission) in any AWS account other than your own. Required
// permissions: kms:ListRetirableGrants
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (IAM policy) in your AWS account. Related operations:
//
// * CreateGrant
//
// *
// ListGrants
//
// * RetireGrant
//
// * RevokeGrant
func (c *Client) ListRetirableGrants(ctx context.Context, params *ListRetirableGrantsInput, optFns ...func(*Options)) (*ListRetirableGrantsOutput, error) {
	if params == nil {
		params = &ListRetirableGrantsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRetirableGrants", params, optFns, addOperationListRetirableGrantsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRetirableGrantsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRetirableGrantsInput struct {

	// The retiring principal for which to list grants. Enter a principal in your AWS
	// account. To specify the retiring principal, use the Amazon Resource Name (ARN)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// an AWS principal. Valid AWS principals include AWS accounts (root), IAM users,
	// federated users, and assumed role users. For examples of the ARN syntax for
	// specifying a principal, see AWS Identity and Access Management (IAM)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)
	// in the Example ARNs section of the Amazon Web Services General Reference.
	//
	// This member is required.
	RetiringPrincipal *string

	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, AWS KMS does not return more than the specified number of
	// items, but it might return fewer. This value is optional. If you include a
	// value, it must be between 1 and 100, inclusive. If you do not include a value,
	// it defaults to 50.
	Limit *int32

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated response
	// you just received.
	Marker *string
}

type ListRetirableGrantsOutput struct {

	// A list of grants.
	Grants []types.GrantListEntry

	// When Truncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent request.
	NextMarker *string

	// A flag that indicates whether there are more items in the list. When this value
	// is true, the list in this response is truncated. To get more items, pass the
	// value of the NextMarker element in thisresponse to the Marker parameter in a
	// subsequent request.
	Truncated bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRetirableGrantsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRetirableGrants{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRetirableGrants{}, middleware.After)
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
	if err = addOpListRetirableGrantsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRetirableGrants(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListRetirableGrants(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ListRetirableGrants",
	}
}
