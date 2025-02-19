// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a model in Amazon SageMaker. In the request, you name the model and
// describe a primary container. For the primary container, you specify the Docker
// image that contains inference code, artifacts (from prior training), and a
// custom environment map that the inference code uses when you deploy the model
// for predictions. Use this API to create a model if you want to use Amazon
// SageMaker hosting services or run a batch transform job. To host your model, you
// create an endpoint configuration with the CreateEndpointConfig API, and then
// create an endpoint with the CreateEndpoint API. Amazon SageMaker then deploys
// all of the containers that you defined for the model in the hosting environment.
// For an example that calls this method when deploying a model to Amazon SageMaker
// hosting services, see Deploy the Model to Amazon SageMaker Hosting Services (AWS
// SDK for Python (Boto 3)).
// (https://docs.aws.amazon.com/sagemaker/latest/dg/ex1-deploy-model.html#ex1-deploy-model-boto)
// To run a batch transform using your model, you start a job with the
// CreateTransformJob API. Amazon SageMaker uses your model and your dataset to get
// inferences which are then saved to a specified S3 location. In the CreateModel
// request, you must define a container with the PrimaryContainer parameter. In the
// request, you also provide an IAM role that Amazon SageMaker can assume to access
// model artifacts and docker image for deployment on ML compute hosting instances
// or for batch transform jobs. In addition, you also use the IAM role to manage
// permissions the inference code needs. For example, if the inference code access
// any other AWS resources, you grant necessary permissions via this role.
func (c *Client) CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) {
	if params == nil {
		params = &CreateModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModel", params, optFns, addOperationCreateModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelInput struct {

	// The Amazon Resource Name (ARN) of the IAM role that Amazon SageMaker can assume
	// to access model artifacts and docker image for deployment on ML compute
	// instances or for batch transform jobs. Deploying on ML compute instances is part
	// of model hosting. For more information, see Amazon SageMaker Roles
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html). To be
	// able to pass this role to Amazon SageMaker, the caller of this API must have the
	// iam:PassRole permission.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The name of the new model.
	//
	// This member is required.
	ModelName *string

	// Specifies the containers in the inference pipeline.
	Containers []types.ContainerDefinition

	// Isolates the model container. No inbound or outbound network calls can be made
	// to or from the model container.
	EnableNetworkIsolation bool

	// The location of the primary docker image containing inference code, associated
	// artifacts, and custom environment map that the inference code uses when the
	// model is deployed for predictions.
	PrimaryContainer *types.ContainerDefinition

	// An array of key-value pairs. You can use tags to categorize your AWS resources
	// in different ways, for example, by purpose, owner, or environment. For more
	// information, see Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	Tags []types.Tag

	// A VpcConfig object that specifies the VPC that you want your model to connect
	// to. Control access to and from your model container by configuring the VPC.
	// VpcConfig is used in hosting services and in batch transform. For more
	// information, see Protect Endpoints by Using an Amazon Virtual Private Cloud
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Data
	// in Batch Transform Jobs by Using an Amazon Virtual Private Cloud
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-vpc.html).
	VpcConfig *types.VpcConfig
}

type CreateModelOutput struct {

	// The ARN of the model created in Amazon SageMaker.
	//
	// This member is required.
	ModelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateModel{}, middleware.After)
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
	if err = addOpCreateModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateModel",
	}
}
