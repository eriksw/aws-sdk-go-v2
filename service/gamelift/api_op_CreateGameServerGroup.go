// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is used with the Amazon GameLift FleetIQ solution and game server
// groups. Creates a GameLift FleetIQ game server group for managing game hosting
// on a collection of Amazon EC2 instances for game hosting. This operation creates
// the game server group, creates an Auto Scaling group in your AWS account, and
// establishes a link between the two groups. You can view the status of your game
// server groups in the GameLift console. Game server group metrics and events are
// emitted to Amazon CloudWatch. Before creating a new game server group, you must
// have the following:
//
// * An Amazon EC2 launch template that specifies how to
// launch Amazon EC2 instances with your game server build. For more information,
// see  Launching an Instance from a Launch Template
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html)
// in the Amazon EC2 User Guide.
//
// * An IAM role that extends limited access to your
// AWS account to allow GameLift FleetIQ to create and interact with the Auto
// Scaling group. For more information, see Create IAM roles for cross-service
// interaction
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-iam-permissions-roles.html)
// in the GameLift FleetIQ Developer Guide.
//
// To create a new game server group,
// specify a unique group name, IAM role and Amazon EC2 launch template, and
// provide a list of instance types that can be used in the group. You must also
// set initial maximum and minimum limits on the group's instance count. You can
// optionally set an Auto Scaling policy with target tracking based on a GameLift
// FleetIQ metric. Once the game server group and corresponding Auto Scaling group
// are created, you have full access to change the Auto Scaling group's
// configuration as needed. Several properties that are set when creating a game
// server group, including maximum/minimum size and auto-scaling policy settings,
// must be updated directly in the Auto Scaling group. Keep in mind that some Auto
// Scaling group properties are periodically updated by GameLift FleetIQ as part of
// its balancing activities to optimize for availability and cost. Learn more
// GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
// Related operations
//
// * CreateGameServerGroup
//
// * ListGameServerGroups
//
// *
// DescribeGameServerGroup
//
// * UpdateGameServerGroup
//
// * DeleteGameServerGroup
//
// *
// ResumeGameServerGroup
//
// * SuspendGameServerGroup
//
// * DescribeGameServerInstances
func (c *Client) CreateGameServerGroup(ctx context.Context, params *CreateGameServerGroupInput, optFns ...func(*Options)) (*CreateGameServerGroupOutput, error) {
	if params == nil {
		params = &CreateGameServerGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGameServerGroup", params, optFns, addOperationCreateGameServerGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGameServerGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGameServerGroupInput struct {

	// An identifier for the new game server group. This value is used to generate
	// unique ARN identifiers for the EC2 Auto Scaling group and the GameLift FleetIQ
	// game server group. The name must be unique per Region per AWS account.
	//
	// This member is required.
	GameServerGroupName *string

	// The EC2 instance types and sizes to use in the Auto Scaling group. The instance
	// definitions must specify at least two different instance types that are
	// supported by GameLift FleetIQ. For more information on instance types, see EC2
	// Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the
	// Amazon EC2 User Guide. You can optionally specify capacity weighting for each
	// instance type. If no weight value is specified for an instance type, it is set
	// to the default value "1". For more information about capacity weighting, see
	// Instance Weighting for Amazon EC2 Auto Scaling
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-weighting.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	//
	// This member is required.
	InstanceDefinitions []types.InstanceDefinition

	// The EC2 launch template that contains configuration settings and game server
	// code to be deployed to all instances in the game server group. You can specify
	// the template using either the template name or ID. For help with creating a
	// launch template, see Creating a Launch Template for an Auto Scaling Group
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html)
	// in the Amazon EC2 Auto Scaling User Guide. After the Auto Scaling group is
	// created, update this value directly in the Auto Scaling group using the AWS
	// console or APIs.
	//
	// This member is required.
	LaunchTemplate *types.LaunchTemplateSpecification

	// The maximum number of instances allowed in the EC2 Auto Scaling group. During
	// automatic scaling events, GameLift FleetIQ and EC2 do not scale up the group
	// above this maximum. After the Auto Scaling group is created, update this value
	// directly in the Auto Scaling group using the AWS console or APIs.
	//
	// This member is required.
	MaxSize *int32

	// The minimum number of instances allowed in the EC2 Auto Scaling group. During
	// automatic scaling events, GameLift FleetIQ and EC2 do not scale down the group
	// below this minimum. In production, this value should be set to at least 1. After
	// the Auto Scaling group is created, update this value directly in the Auto
	// Scaling group using the AWS console or APIs.
	//
	// This member is required.
	MinSize *int32

	// The Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) for an IAM
	// role that allows Amazon GameLift to access your EC2 Auto Scaling groups.
	//
	// This member is required.
	RoleArn *string

	// Configuration settings to define a scaling policy for the Auto Scaling group
	// that is optimized for game hosting. The scaling policy uses the metric
	// "PercentUtilizedGameServers" to maintain a buffer of idle game servers that can
	// immediately accommodate new games and players. After the Auto Scaling group is
	// created, update this value directly in the Auto Scaling group using the AWS
	// console or APIs.
	AutoScalingPolicy *types.GameServerGroupAutoScalingPolicy

	// Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand
	// Instances in the game server group. Method options include the following:
	//
	// *
	// SPOT_ONLY - Only Spot Instances are used in the game server group. If Spot
	// Instances are unavailable or not viable for game hosting, the game server group
	// provides no hosting capacity until Spot Instances can again be used. Until then,
	// no new instances are started, and the existing nonviable Spot Instances are
	// terminated (after current gameplay ends) and are not replaced.
	//
	// * SPOT_PREFERRED
	// - (default value) Spot Instances are used whenever available in the game server
	// group. If Spot Instances are unavailable, the game server group continues to
	// provide hosting capacity by falling back to On-Demand Instances. Existing
	// nonviable Spot Instances are terminated (after current gameplay ends) and are
	// replaced with new On-Demand Instances.
	//
	// * ON_DEMAND_ONLY - Only On-Demand
	// Instances are used in the game server group. No Spot Instances are used, even
	// when available, while this balancing strategy is in force.
	BalancingStrategy types.BalancingStrategy

	// A flag that indicates whether instances in the game server group are protected
	// from early termination. Unprotected instances that have active game servers
	// running might be terminated during a scale-down event, causing players to be
	// dropped from the game. Protected instances cannot be terminated while there are
	// active game servers running except in the event of a forced game server group
	// deletion (see ). An exception to this is with Spot Instances, which can be
	// terminated by AWS regardless of protection status. This property is set to
	// NO_PROTECTION by default.
	GameServerProtectionPolicy types.GameServerProtectionPolicy

	// A list of labels to assign to the new game server group resource. Tags are
	// developer-defined key-value pairs. Tagging AWS resources is useful for resource
	// management, access management, and cost allocation. For more information, see
	// Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the AWS
	// General Reference. Once the resource is created, you can use TagResource,
	// UntagResource, and ListTagsForResource to add, remove, and view tags,
	// respectively. The maximum tag limit may be lower than stated. See the AWS
	// General Reference for actual tagging limits.
	Tags []types.Tag

	// A list of virtual private cloud (VPC) subnets to use with instances in the game
	// server group. By default, all GameLift FleetIQ-supported Availability Zones are
	// used. You can use this parameter to specify VPCs that you've set up. This
	// property cannot be updated after the game server group is created, and the
	// corresponding Auto Scaling group will always use the property value that is set
	// with this request, even if the Auto Scaling group is updated directly.
	VpcSubnets []string
}

type CreateGameServerGroupOutput struct {

	// The newly created game server group object, including the new ARN value for the
	// GameLift FleetIQ game server group and the object's status. The EC2 Auto Scaling
	// group ARN is initially null, since the group has not yet been created. This
	// value is added once the game server group status reaches ACTIVE.
	GameServerGroup *types.GameServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateGameServerGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateGameServerGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateGameServerGroup{}, middleware.After)
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
	if err = addOpCreateGameServerGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGameServerGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGameServerGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateGameServerGroup",
	}
}
