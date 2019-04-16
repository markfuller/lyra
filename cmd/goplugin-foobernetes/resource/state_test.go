package resource

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeployments(t *testing.T) {
	if true {
		return
	}
	deployment := loadFakeCloudData()
	require.NotEmpty(t, deployment)

	instances := deployment.Instances
	require.Len(t, instances, 3)

	dbInstanceID := ""
	appServerIDs := []string{}
	for _, v := range instances {
		if v.Image == "lyra::database" {
			require.Equal(t, "eu1", *v.Location)
			require.Equal(t, "64G", v.Memory)
			dbInstanceID = *v.InstanceID
		} else {
			require.Equal(t, "lyra::application", v.Image)
			appServerIDs = append(appServerIDs, *v.InstanceID)
		}
	}
	require.NotEmpty(t, dbInstanceID)
	require.Len(t, appServerIDs, 2)
	lbs := deployment.LoadBalancers
	require.Len(t, lbs, 2)

	webServers := deployment.WebServers
	webServerIDs := []string{}
	require.Len(t, webServers, 2)
	for _, v := range webServers {
		require.ElementsMatch(t, appServerIDs, v.AppServers)
		webServerIDs = append(webServerIDs, *v.WebServerID)
	}
	require.Len(t, webServerIDs, 2)

	var foundPrimaryLB, foundSecondaryLB bool
	for _, v := range lbs {
		if *v.Location == "eu1" {
			require.Equal(t, false, *v.Replica)
			require.Equal(t, "10.0.0.1", *v.LoadBalancerIP)
			tags := *v.Tags
			require.Equal(t, "primary", tags["role"])
			require.Equal(t, "lyra team", tags["team"])
			require.ElementsMatch(t, webServerIDs, v.WebServerIDs)
			foundPrimaryLB = true
		} else if *v.Location == "eu2" {
			require.Equal(t, true, *v.Replica)
			require.Equal(t, "10.0.0.2", *v.LoadBalancerIP)
			tags := *v.Tags
			require.Equal(t, "secondary", tags["role"])
			require.ElementsMatch(t, webServerIDs, v.WebServerIDs)
			foundSecondaryLB = true
		}
	}
	require.True(t, foundSecondaryLB)
	require.True(t, foundPrimaryLB)

}
