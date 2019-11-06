package bor

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	borTypes "github.com/maticnetwork/heimdall/bor/types"
	"github.com/maticnetwork/heimdall/types"
)

var cdc = codec.New()

//
// Propose Span Msg
//

var _ sdk.Msg = &MsgProposeSpan{}

// MsgProposeSpan creates msg propose span
type MsgProposeSpan struct {
	ID         uint64                `json:"span_id"`
	Proposer   types.HeimdallAddress `json:"proposer"`
	StartBlock uint64                `json:"start_block"`
	EndBlock   uint64                `json:"end_block"`
	ChainID    string                `json:"bor_chain_id"`
}

// NewMsgProposeSpan creates new propose span message
func NewMsgProposeSpan(
	id uint64,
	proposer types.HeimdallAddress,
	startBlock uint64,
	endBlock uint64,
	chainID string,
) MsgProposeSpan {
	return MsgProposeSpan{
		ID:         id,
		Proposer:   proposer,
		StartBlock: startBlock,
		EndBlock:   endBlock,
		ChainID:    chainID,
	}
}

// Type returns message type
func (msg MsgProposeSpan) Type() string {
	return "propose-span"
}

// Route returns route for message
func (msg MsgProposeSpan) Route() string {
	return borTypes.RouterKey
}

// GetSigners returns address of the signer
func (msg MsgProposeSpan) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{types.HeimdallAddressToAccAddress(msg.Proposer)}
}

// GetSignBytes returns sign bytes for proposeSpan message type
func (msg MsgProposeSpan) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// ValidateBasic validates the message and returns error
func (msg MsgProposeSpan) ValidateBasic() sdk.Error {
	if msg.Proposer.Empty() {
		return sdk.ErrInvalidAddress(msg.Proposer.String())
	}

	return nil
}