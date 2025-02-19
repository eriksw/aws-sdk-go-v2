// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a change request for Change Manager. The runbooks (Automation documents)
// specified in the change request run only after all required approvals for the
// change request have been received.
func (c *Client) StartChangeRequestExecution(ctx context.Context, params *StartChangeRequestExecutionInput, optFns ...func(*Options)) (*StartChangeRequestExecutionOutput, error) {
	if params == nil {
		params = &StartChangeRequestExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartChangeRequestExecution", params, optFns, addOperationStartChangeRequestExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartChangeRequestExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartChangeRequestExecutionInput struct {

	// The name of the change template document to run during the runbook workflow.
	//
	// This member is required.
	DocumentName *string

	// Information about the Automation runbooks (Automation documents) that are run
	// during the runbook workflow. The Automation runbooks specified for the runbook
	// workflow can't run until all required approvals for the change request have been
	// received.
	//
	// This member is required.
	Runbooks []types.Runbook

	// The name of the change request associated with the runbook workflow to be run.
	ChangeRequestName *string

	// The user-provided idempotency token. The token must be unique, is case
	// insensitive, enforces the UUID format, and can't be reused.
	ClientToken *string

	// The version of the change template document to run during the runbook workflow.
	DocumentVersion *string

	// A key-value map of parameters that match the declared parameters in the change
	// template document.
	Parameters map[string][]string

	// The date and time specified in the change request to run the Automation
	// runbooks. The Automation runbooks specified for the runbook workflow can't run
	// until all required approvals for the change request have been received.
	ScheduledTime *time.Time

	// Optional metadata that you assign to a resource. You can specify a maximum of
	// five tags for a change request. Tags enable you to categorize a resource in
	// different ways, such as by purpose, owner, or environment. For example, you
	// might want to tag a change request to identify an environment or target AWS
	// Region. In this case, you could specify the following key-value pairs:
	//
	// *
	// Key=Environment,Value=Production
	//
	// * Key=Region,Value=us-east-2
	Tags []types.Tag
}

type StartChangeRequestExecutionOutput struct {

	// The unique ID of a runbook workflow operation. (A runbook workflow is a type of
	// Automation operation.)
	AutomationExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartChangeRequestExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartChangeRequestExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartChangeRequestExecution{}, middleware.After)
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
	if err = addOpStartChangeRequestExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartChangeRequestExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartChangeRequestExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "StartChangeRequestExecution",
	}
}
