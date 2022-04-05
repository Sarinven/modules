// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/modules/assets/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type assetID struct {
	ClassificationID types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	HashID           types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*assetID)(nil)
var _ helpers.Key = (*assetID)(nil)

func (assetID assetID) String() string {
	var values []string
	values = append(values, assetID.ClassificationID.String())
	values = append(values, assetID.HashID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (assetID assetID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, assetID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, assetID.HashID.Bytes()...)

	return Bytes
}
func (assetID assetID) Compare(id types.ID) int {
	return bytes.Compare(assetID.Bytes(), id.Bytes())
}
func (assetID assetID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(assetID.Bytes())
}
func (assetID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, assetID{})
}
func (assetID assetID) IsPartial() bool {
	return len(assetID.HashID.Bytes()) == 0
}
func (assetID assetID) Equals(key helpers.Key) bool {
	return assetID.Compare(assetIDFromInterface(key)) == 0
}

func NewAssetID(classificationID types.ID, immutableProperties types.Properties) types.ID {
	return assetID{
		ClassificationID: classificationID,
		HashID:           base.HasImmutables{Properties: immutableProperties}.GenerateHashID(),
	}
}
