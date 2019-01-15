package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/lyraproj/lyra/cmd/goplugin-aws/resource"
	"github.com/lyraproj/puppet-evaluator/eval"
	"github.com/lyraproj/servicesdk/grpc"
	"github.com/lyraproj/servicesdk/service"
)

const (
	providerName      = "lyra-aws-ec2"
	providerNamespace = "Lyra::Aws"
	logLevel          = "info"
)

// Start this provider
func Start() {

	eval.Puppet.Do(func(c eval.Context) {

		sb := service.NewServerBuilder(c, `Aws`)

		var evs []eval.Type
		evs = sb.RegisterTypes("Aws",
			ec2.Vpc{},
			ec2.VpcCidrBlockState{},
			ec2.VpcCidrBlockAssociation{},
			ec2.VpcIpv6CidrBlockAssociation{},
			ec2.Tag{},
		)
		sb.RegisterHandler("Aws::NativeVpcHandler", &resource.NativeVPCHandler{}, evs[0])

		s := sb.Server()
		grpc.Serve(c, s)
	})
}
