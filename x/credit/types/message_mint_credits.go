package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintCredits = "mint_credits"

var _ sdk.Msg = &MsgMintCredits{}

func NewMsgMintCredits(creator string, amount uint64) *MsgMintCredits {
	return &MsgMintCredits{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgMintCredits) Route() string {
	return RouterKey
}

func (msg *MsgMintCredits) Type() string {
	return TypeMsgMintCredits
}

func (msg *MsgMintCredits) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintCredits) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintCredits) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
