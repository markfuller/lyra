type Aws = TypeSet[{
  pcore_uri => 'http://puppet.com/2016.1/pcore',
  pcore_version => '1.0.0',
  name_authority => 'http://puppet.com/2016.1/runtime',
  name => 'Aws',
  version => '0.1.0',
  types => {
    VpcCidrBlockAssociation => {
      attributes => {
        'association_id' => {
          'type' => Optional[String],
          'value' => undef
        },
        'cidr_block' => {
          'type' => Optional[String],
          'value' => undef
        },
        'cidr_block_state' => {
          'type' => Optional,
          'value' => undef
        }
      },
      functions => {
        'go_string' => Callable[
          [0, 0],
          String],
        'string' => Callable[
          [0, 0],
          String]
      }
    },
    VpcIpv6CidrBlockAssociation => {
      attributes => {
        'association_id' => {
          'type' => Optional[String],
          'value' => undef
        },
        'ipv6_cidr_block' => {
          'type' => Optional[String],
          'value' => undef
        },
        'ipv6_cidr_block_state' => {
          'type' => Optional,
          'value' => undef
        }
      },
      functions => {
        'go_string' => Callable[
          [0, 0],
          String],
        'string' => Callable[
          [0, 0],
          String]
      }
    },
    Tag => {
      attributes => {
        'key' => {
          'type' => Optional[String],
          'value' => undef
        },
        'value' => {
          'type' => Optional[String],
          'value' => undef
        }
      },
      functions => {
        'go_string' => Callable[
          [0, 0],
          String],
        'string' => Callable[
          [0, 0],
          String]
      }
    },
    Vpc2 => {
      attributes => {
        'cidr_block' => {
          'type' => Optional[String],
          'value' => undef
        },
        'cidr_block_association_set' => Array[VpcCidrBlockAssociation],
        'dhcp_options_id' => {
          'type' => Optional[String],
          'value' => undef
        },
        'instance_tenancy' => {
          'type' => Optional[String],
          'value' => undef
        },
        'ipv6_cidr_block_association_set' => Array[VpcIpv6CidrBlockAssociation],
        'is_default' => {
          'type' => Optional[Boolean],
          'value' => undef
        },
        'owner_id' => {
          'type' => Optional[String],
          'value' => undef
        },
        'state' => {
          'type' => Optional[String],
          'value' => undef
        },
        'tags' => Array[Tag],
        'vpc_id' => {
          'type' => Optional[String],
          'value' => undef
        }
      },
      functions => {
        'go_string' => Callable[
          [0, 0],
          String],
        'string' => Callable[
          [0, 0],
          String]
      }
    },
    VpcCidrBlockState => {
      attributes => {
        'state' => {
          'type' => Optional[String],
          'value' => undef
        },
        'status_message' => {
          'type' => Optional[String],
          'value' => undef
        }
      },
      functions => {
        'go_string' => Callable[
          [0, 0],
          String],
        'string' => Callable[
          [0, 0],
          String]
      }
    }
  }
}]