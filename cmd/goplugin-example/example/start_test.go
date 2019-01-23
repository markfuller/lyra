package example

import (
	"bytes"
	"fmt"
	"github.com/lyraproj/puppet-evaluator/eval"

	// Initialize pcore
	_ "github.com/lyraproj/puppet-evaluator/pcore"
)

func ExampleServer() {
	eval.Puppet.Do(func(c eval.Context) {
		s := Server(c)
		typeSet, _ := s.Metadata(c)
		b := bytes.NewBufferString("type Example = ")

		typeSet.ToString(b, eval.PRETTY_EXPANDED, nil)
		output := b.String()

		fmt.Println(output)
	})
	//output:
	// type Example = TypeSet[{
	//   pcore_uri => 'http://puppet.com/2016.1/pcore',
	//   pcore_version => '1.0.0',
	//   name_authority => 'http://puppet.com/2016.1/runtime',
	//   name => 'Example',
	//   version => '0.1.0',
	//   types => {
	//     Address => {
	//       attributes => {
	//         'line_one' => {
	//           'type' => String,
	//           'value' => ''
	//         }
	//       }
	//     },
	//     ContainedRes => {
	//       annotations => {
	//         Lyra::Resource => {
	//           'provided_attributes' => ['id'],
	//           'relationships' => {
	//             'owner' => {
	//               'type' => OwnerRes,
	//               'kind' => 'container',
	//               'cardinality' => 'one',
	//               'keys' => ['owner_id', 'id']
	//             }
	//           }
	//         }
	//       },
	//       attributes => {
	//         'id' => {
	//           'type' => Optional[String],
	//           'value' => undef
	//         },
	//         'owner_id' => String,
	//         'stuff' => String
	//       }
	//     },
	//     OwnerRes => {
	//       annotations => {
	//         Lyra::Resource => {
	//           'provided_attributes' => ['id'],
	//           'relationships' => {
	//             'mine' => {
	//               'type' => ContainedRes,
	//               'kind' => 'contained',
	//               'cardinality' => 'many',
	//               'keys' => ['id', 'owner_id']
	//             }
	//           }
	//         }
	//       },
	//       attributes => {
	//         'id' => {
	//           'type' => Optional[String],
	//           'value' => undef
	//         },
	//         'phone' => String
	//       }
	//     },
	//     Person => {
	//       attributes => {
	//         'name' => {
	//           'type' => String,
	//           'value' => ''
	//         },
	//         'age' => {
	//           'type' => Integer,
	//           'value' => 0
	//         },
	//         'human' => {
	//           'type' => Boolean,
	//           'value' => false
	//         },
	//         'address' => {
	//           'type' => Optional[Address],
	//           'value' => undef
	//         },
	//         'created' => {
	//           'type' => Optional[Timestamp],
	//           'value' => undef
	//         }
	//       }
	//     },
	//     PersonHandler => {
	//       functions => {
	//         'create' => Callable[
	//           [Person],
	//           Tuple[Person, String]],
	//         'delete' => Callable[String],
	//         'read' => Callable[
	//           [String],
	//           Person],
	//         'update' => Callable[
	//           [String, Person],
	//           Person]
	//       }
	//     }
	//   }
	// }]
}
