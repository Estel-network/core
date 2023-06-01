package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCloseTransaction = "close_transaction"

var _ sdk.Msg = &MsgCloseTransaction{}

func NewMsgCloseTransaction(creator string, id string) *MsgCloseTransaction {
	return &MsgCloseTransaction{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgCloseTransaction) Route() string {
	return RouterKey
}

func (msg *MsgCloseTransaction) Type() string {
	return TypeMsgCloseTransaction
}

func (msg *MsgCloseTransaction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCloseTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCloseTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
