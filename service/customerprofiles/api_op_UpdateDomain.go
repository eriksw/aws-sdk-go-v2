// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the properties of a domain, including creating or selecting a dead
// letter queue or an encryption key. Once a domain is created, the name can’t be
// changed.
func (c *Client) UpdateDomain(ctx context.Context, params *UpdateDomainInput, optFns ...func(*Options)) (*UpdateDomainOutput, error) {
	if params == nil {
		params = &UpdateDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomain", params, optFns, addOperationUpdateDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDomainInput struct {

	// The unique name for the domain.
	//
	// This member is required.
	DomainName *string

	// The URL of the SQS dead letter queue, which is used for reporting errors
	// associated with ingesting data from third party applications. If specified as an
	// empty string, it will clear any existing value. You must set up a policy on the
	// DeadLetterQueue for the SendMessage operation to enable Amazon Connect Customer
	// Profiles to send messages to the DeadLetterQueue.
	DeadLetterQueueUrl *string

	// The default encryption key, which is an AWS managed key, is used when no
	// specific type of encryption key is specified. It is used to encrypt all data
	// before it is placed in permanent or semi-permanent storage. If specified as an
	// empty string, it will clear any existing value.
	DefaultEncryptionKey *string

	// The default number of days until the data within the domain expires.
	DefaultExpirationDays *int32

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string
}

type UpdateDomainOutput struct {

	// The timestamp of when the domain was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The unique name for the domain.
	//
	// This member is required.
	DomainName *string

	// The timestamp of when the domain was most recently edited.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The URL of the SQS dead letter queue, which is used for reporting errors
	// associated with ingesting data from third party applications.
	DeadLetterQueueUrl *string

	// The default encryption key, which is an AWS managed key, is used when no
	// specific type of encryption key is specified. It is used to encrypt all data
	// before it is placed in permanent or semi-permanent storage.
	DefaultEncryptionKey *string

	// The default number of days until the data within the domain expires.
	DefaultExpirationDays *int32

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDomain{}, middleware.After)
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
	if err = addOpUpdateDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "UpdateDomain",
	}
}
