// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates a contact flow to start a new task.
func (c *Client) StartTaskContact(ctx context.Context, params *StartTaskContactInput, optFns ...func(*Options)) (*StartTaskContactOutput, error) {
	if params == nil {
		params = &StartTaskContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTaskContact", params, optFns, addOperationStartTaskContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTaskContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTaskContactInput struct {

	// The identifier of the contact flow for initiating the tasks. To see the
	// ContactFlowId in the Amazon Connect console user interface, on the navigation
	// menu go to Routing, Contact Flows. Choose the contact flow. On the contact flow
	// page, under the name of the contact flow, choose Show additional flow
	// information. The ContactFlowId is the last part of the ARN, shown here in bold:
	// arn:aws:connect:us-west-2:xxxxxxxxxxxx:instance/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/contact-flow/846ec553-a005-41c0-8341-xxxxxxxxxxxx
	//
	// This member is required.
	ContactFlowId *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The name of a task that is shown to an agent in the Contact Control Panel (CCP).
	//
	// This member is required.
	Name *string

	// A custom key-value pair using an attribute map. The attributes are standard
	// Amazon Connect attributes, and can be accessed in contact flows just like any
	// other contact attributes. There can be up to 32,768 UTF-8 bytes across all
	// key-value pairs per contact. Attribute keys can include only alphanumeric, dash,
	// and underscore characters.
	Attributes map[string]string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// A description of the task that is shown to an agent in the Contact Control Panel
	// (CCP).
	Description *string

	// The identifier of the previous chat, voice, or task contact.
	PreviousContactId *string

	// A formatted URL that is shown to an agent in the Contact Control Panel (CCP).
	References map[string]types.Reference
}

type StartTaskContactOutput struct {

	// The identifier of this contact within the Amazon Connect instance.
	ContactId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartTaskContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartTaskContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartTaskContact{}, middleware.After)
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
	if err = addIdempotencyToken_opStartTaskContactMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartTaskContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTaskContact(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartTaskContact struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartTaskContact) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartTaskContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartTaskContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartTaskContactInput ")
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
func addIdempotencyToken_opStartTaskContactMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartTaskContact{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartTaskContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "StartTaskContact",
	}
}
