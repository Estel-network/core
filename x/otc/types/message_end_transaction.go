package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEndTransaction = "end_transaction"

var _ sdk.Msg = &MsgEndTransaction{}

func NewMsgEndTransaction(creator string, id string) *MsgEndTransaction {
	return &MsgEndTransaction{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgEndTransaction) Route() string {
	return RouterKey
}

func (msg *MsgEndTransaction) Type() string {
	return TypeMsgEndTransaction
}

func (msg *MsgEndTransaction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEndTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEndTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
