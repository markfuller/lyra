package resource

import (
	"regexp"
	"testing"

	"github.com/lyraproj/pcore/px"
	"github.com/lyraproj/pcore/types"
	"github.com/stretchr/testify/require"
)

func TestConvertOutput(t *testing.T) {
	s := `{
  "value": {
		"rgId": "/subscriptions/c82736ee-c108-452b-8178-f548c95d18fe/resourceGroups/lyra-tf-test",
		"rgLocation": "westeurope",
		"rgName": "lyra-tf-test",
		"1": 2,
		"false": true
	}
}`

	m := convertOutput([]byte(s))
	h, ok := m.(px.OrderedMap)
	require.True(t, ok)
	require.Equal(t, 5, h.Len())
	require.Equal(t, "/subscriptions/c82736ee-c108-452b-8178-f548c95d18fe/resourceGroups/lyra-tf-test", h.Get5("rgId", px.EmptyString).String())
	require.Equal(t, "westeurope", h.Get5("rgLocation", px.EmptyString).String())
	require.Equal(t, "lyra-tf-test", h.Get5("rgName", px.EmptyString).String())
	require.Equal(t, int64(2), h.Get5("1", types.ZERO).(px.Integer).Int())
	require.Equal(t, true, h.Get5("false", types.BooleanFalse).(px.Boolean).Bool())
}

func TestConvertOutput_EmptyOutput(t *testing.T) {
	m := convertOutput([]byte(""))
	require.Equal(t, px.Undef, m)
}

func TestUniqueString(t *testing.T) {
	s := uniqueString()
	re := regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{8}")
	b := re.Match([]byte(s))
	require.True(t, b)

}

func TestEndToEnd(t *testing.T) {
	t.Skip("Only to be run as needed.  Will probably time out on deletion")
	ch := ConfigHandler{}
	c := Config{
		WorkingDir: "/Users/Mark/code/lyraproj/lyra/examples/go-samples/terraform_go/tfroot",
	}
	_, externalID, err := ch.Create(&c)
	require.NoError(t, err)
	require.Equal(t, c.WorkingDir, externalID)
	_, err = ch.Update(externalID, &c)
	require.NoError(t, err)
	err = ch.Delete(externalID)
	require.NoError(t, err)
}
