// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the description of an endpoint.
func (c *Client) DescribeEndpoint(ctx context.Context, params *DescribeEndpointInput, optFns ...func(*Options)) (*DescribeEndpointOutput, error) {
	if params == nil {
		params = &DescribeEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpoint", params, optFns, addOperationDescribeEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointInput struct {

	// The name of the endpoint.
	//
	// This member is required.
	EndpointName *string
}

type DescribeEndpointOutput struct {

	// A timestamp that shows when the endpoint was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the endpoint.
	//
	// This member is required.
	EndpointArn *string

	// The name of the endpoint configuration associated with this endpoint.
	//
	// This member is required.
	EndpointConfigName *string

	// Name of the endpoint.
	//
	// This member is required.
	EndpointName *string

	// The status of the endpoint.
	//
	// * OutOfService: Endpoint is not available to take
	// incoming requests.
	//
	// * Creating: CreateEndpoint is executing.
	//
	// * Updating:
	// UpdateEndpoint or UpdateEndpointWeightsAndCapacities is executing.
	//
	// *
	// SystemUpdating: Endpoint is undergoing maintenance and cannot be updated or
	// deleted or re-scaled until it has completed. This maintenance operation does not
	// change any customer-specified values such as VPC config, KMS encryption, model,
	// instance type, or instance count.
	//
	// * RollingBack: Endpoint fails to scale up or
	// down or change its variant weight and is in the process of rolling back to its
	// previous configuration. Once the rollback completes, endpoint returns to an
	// InService status. This transitional status only applies to an endpoint that has
	// autoscaling enabled and is undergoing variant weight or capacity changes as part
	// of an UpdateEndpointWeightsAndCapacities call or when the
	// UpdateEndpointWeightsAndCapacities operation is called explicitly.
	//
	// * InService:
	// Endpoint is available to process incoming requests.
	//
	// * Deleting: DeleteEndpoint
	// is executing.
	//
	// * Failed: Endpoint could not be created, updated, or re-scaled.
	// Use DescribeEndpointOutput$FailureReason for information about the failure.
	// DeleteEndpoint is the only operation that can be performed on a failed endpoint.
	//
	// This member is required.
	EndpointStatus types.EndpointStatus

	// A timestamp that shows when the endpoint was last modified.
	//
	// This member is required.
	LastModifiedTime *time.Time

	//
	DataCaptureConfig *types.DataCaptureConfigSummary

	// If the status of the endpoint is Failed, the reason why it failed.
	FailureReason *string

	// The most recent deployment configuration for the endpoint.
	LastDeploymentConfig *types.DeploymentConfig

	// An array of ProductionVariantSummary objects, one for each model hosted behind
	// this endpoint.
	ProductionVariants []types.ProductionVariantSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEndpoint{}, middleware.After)
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
	if err = addOpDescribeEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeEndpoint",
	}
}
