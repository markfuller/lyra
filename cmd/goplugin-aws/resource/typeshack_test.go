package resource

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/lyraproj/puppet-evaluator/eval"
	"github.com/lyraproj/semver/semver"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"

	// Initialize pcore
	_ "github.com/lyraproj/puppet-evaluator/pcore"
)

func TestGeneratePuppetTypes(t *testing.T) {
	fmt.Println("Started")
	instance := ec2.Vpc{}
	typ := reflect.TypeOf(instance)
	found := make(map[reflect.Type]string)
	m := getAllNestedTypes(typ, found)

	//convert to slice
	keys := make([]reflect.Type, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	require.NotEmpty(t, keys)
	eval.Puppet.Do(func(c eval.Context) {
		// Create a TypeSet from a list of Go structs
		typeSet := c.Reflector().TypeSetFromReflect(`Aws`, semver.MustParseVersion(`0.1.0`), nil,
			keys...)
		// Make the types known to the current loader
		c.AddTypes(typeSet)

		b := bytes.NewBufferString("")

		// Print the TypeSet in human readable form
		typeSet.ToString(b, eval.PRETTY_EXPANDED, nil)
		output := b.String()
		t.Log(output)

	})

}

func TestGetAllNestedTypes(t *testing.T) {
	typ := reflect.TypeOf(ec2.Vpc{})
	found := make(map[reflect.Type]string)
	allTypes := getAllNestedTypes(typ, found)
	require.Contains(t, allTypes, reflect.TypeOf(ec2.Vpc{}))
	require.Contains(t, allTypes, reflect.TypeOf(ec2.VpcCidrBlockAssociation{}))
	require.Contains(t, allTypes, reflect.TypeOf(ec2.VpcCidrBlockState{}))
}

func getAllNestedTypes(typ reflect.Type, found map[reflect.Type]string) map[reflect.Type]string {
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		//get the child type for a slice or array
		childType := field.Type

		if field.Type.Kind() == reflect.Slice || field.Type.Kind() == reflect.Interface || field.Type.Kind() == reflect.Ptr || typ.Kind() == reflect.Array {
			childType = field.Type.Elem()

			//this handles where we have an array of slices, probably proper recursion here (e.g. would this cover *[]*ec2.ssda)
			if childType.Kind() == reflect.Ptr {
				childType = childType.Elem()
			}
		}

		//recurse only if we didn't already find this type
		if _, isPresent := found[childType]; !isPresent {
			if childType.Kind() == reflect.Struct {
				childFound := getAllNestedTypes(childType, found)
				for k, v := range childFound {
					found[k] = v
				}
			}
		}
	}

	v := fmt.Sprintf("%v", typ)
	if v != "struct {}" {
		found[typ] = v
	}

	return found
}
