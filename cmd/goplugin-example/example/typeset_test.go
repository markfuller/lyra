package example

import (
	"bytes"
	"fmt"
	"github.com/lyraproj/puppet-evaluator/eval"
	"github.com/lyraproj/servicesdk/service"
	"time"

	// Initialize pcore
	_ "github.com/lyraproj/puppet-evaluator/pcore"
)

func Example_typeSetChecker() {
	output := typeSetChecker()
	fmt.Println(output)
	//output:
	// type Example = TypeSet[{
	//   pcore_uri => 'http://puppet.com/2016.1/pcore',
	//   pcore_version => '1.0.0',
	//   name_authority => 'http://puppet.com/2016.1/runtime',
	//   name => 'Testing',
	//   version => '0.1.0',
	//   types => {
	//     Bar => {
	//       attributes => {
	//         'abc' => String,
	//         'the_time' => Timestamp
	//       }
	//     },
	//     Baz => {
	//       attributes => {
	//         'xyz' => String
	//       }
	//     },
	//     Foo => {
	//       attributes => {
	//         'bar_pointers' => Array[Optional[Bar]],
	//         'baz_pointers' => Array[Optional[Baz]],
	//         'bazs' => Array[Baz]
	//       }
	//     }
	//   }
	// }]
}

func typeSetChecker() string {
	output := ""
	eval.Puppet.Do(func(c eval.Context) {
		sb := service.NewServerBuilder(c, `Testing`)
		_ = sb.RegisterTypes("Testing",
			Foo{},
			Bar{},
			Baz{})
		s := sb.Server()
		typeSet, _ := s.Metadata(c)
		b := bytes.NewBufferString("type Example = ")

		typeSet.ToString(b, eval.PRETTY_EXPANDED, nil)

		output = b.String()
	})
	return output
}

type Foo struct {
	BarPointers []*Bar //Array[Optional] instead of Array[Optional[Bar]]
	BazPointers []*Baz //Array[Baz] instead of Array[Optional[Baz]]
	Bazs        []Baz
}

type Bar struct {
	Abc     string
	TheTime time.Time
}

type Baz struct {
	Xyz string
}
