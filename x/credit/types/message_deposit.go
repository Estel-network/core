package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgDeposit = "deposit"

var _ sdk.Msg = &MsgDeposit{}

func NewMsgDeposit(creator string, id string, denom string, amount string) *MsgDeposit {
	return &MsgDeposit{
		Creator: creator,
		Id:      id,
		Denom:   denom,
		Amount:  amount,
	}
}

func (msg *MsgDeposit) Route() string {
	return RouterKey
}

func (msg *MsgDeposit) Type() string {
	return TypeMsgDeposit
}

func (msg *MsgDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeposit) ValidateBasic() error {
	/*_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	amount, _ := sdk.ParseCoinsNormalized(msg.Amount)
	if !amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coins object")
	}
	if amount.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is empty")
	}*/
	return nil
}
