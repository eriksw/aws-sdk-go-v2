// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a thesaurus for an index. The thesaurus contains a list of synonyms in
// Solr format.
func (c *Client) CreateThesaurus(ctx context.Context, params *CreateThesaurusInput, optFns ...func(*Options)) (*CreateThesaurusOutput, error) {
	if params == nil {
		params = &CreateThesaurusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateThesaurus", params, optFns, addOperationCreateThesaurusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateThesaurusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateThesaurusInput struct {

	// The unique identifier of the index for the new thesaurus.
	//
	// This member is required.
	IndexId *string

	// The name for the new thesaurus.
	//
	// This member is required.
	Name *string

	// An AWS Identity and Access Management (IAM) role that gives Amazon Kendra
	// permissions to access thesaurus file specified in SourceS3Path.
	//
	// This member is required.
	RoleArn *string

	// The thesaurus file Amazon S3 source path.
	//
	// This member is required.
	SourceS3Path *types.S3Path

	// A token that you provide to identify the request to create a thesaurus. Multiple
	// calls to the CreateThesaurus operation with the same client token will create
	// only one index.
	ClientToken *string

	// The description for the new thesaurus.
	Description *string

	// A list of key-value pairs that identify the thesaurus. You can use the tags to
	// identify and organize your resources and to control access to resources.
	Tags []types.Tag
}

type CreateThesaurusOutput struct {

	// The unique identifier of the thesaurus.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateThesaurusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateThesaurus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateThesaurus{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateThesaurusMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateThesaurusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateThesaurus(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateThesaurus struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateThesaurus) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateThesaurus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateThesaurusInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateThesaurusInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateThesaurusMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateThesaurus{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateThesaurus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "CreateThesaurus",
	}
}
