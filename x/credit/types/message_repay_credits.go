package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRepayCredits = "repay_credits"

var _ sdk.Msg = &MsgRepayCredits{}

func NewMsgRepayCredits(creator string, amount uint64) *MsgRepayCredits {
	return &MsgRepayCredits{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgRepayCredits) Route() string {
	return RouterKey
}

func (msg *MsgRepayCredits) Type() string {
	return TypeMsgRepayCredits
}

func (msg *MsgRepayCredits) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRepayCredits) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRepayCredits) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
