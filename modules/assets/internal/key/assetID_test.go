package key

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_AssetID_Methods(t *testing.T) {
	classificationID := base.NewID("classificationID")
	immutables := base.NewImmutables(base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData")))))

	testAssetID := NewAssetID(classificationID, immutables).(assetID)
	require.Equal(t, assetID{ClassificationID: classificationID, HashID: immutables.GetHashID()}, testAssetID)
	require.Equal(t, strings.Join([]string{classificationID.String(), immutables.GetHashID().String()}, constants.FirstOrderCompositeIDSeparator), testAssetID.String())
	require.Equal(t, false, testAssetID.IsPartial())
	require.Equal(t, true, assetID{ClassificationID: classificationID, HashID: base.NewID("")}.IsPartial())
	require.Equal(t, true, testAssetID.Equals(testAssetID))
	require.Equal(t, false, testAssetID.Equals(assetID{ClassificationID: classificationID, HashID: base.NewID("")}))
	require.Equal(t, true, testAssetID.Matches(testAssetID))
	require.Equal(t, false, testAssetID.Matches(nil))
	require.Equal(t, false, testAssetID.Matches(assetID{ClassificationID: classificationID, HashID: base.NewID("")}))
	require.Equal(t, testAssetID, New(testAssetID))
	require.Equal(t, assetID{ClassificationID: base.NewID(""), HashID: base.NewID("")}, New(base.NewID("")))
	require.Equal(t, testAssetID, readAssetID(testAssetID.String()))
}
