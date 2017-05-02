// Copyright 2017 Pulumi, Inc. All rights reserved.

package ec2

import (
	"github.com/pulumi/coconut/pkg/resource/idl"
)

// An Internet gateway enables your instances to connect to the Internet through the Amazon EC2 edge network.  See
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internet-gateway.html.
type InternetGateway struct {
	idl.NamedResource
}