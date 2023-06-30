package keeper

import (
	"errors"
	"strconv"

	"planet/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
)

// TransmitUpdatePostPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitUpdatePostPacket(
	ctx sdk.Context,
	packetData types.UpdatePostPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %w", err)
	}

	return k.channelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvUpdatePostPacket processes packet reception
func (k Keeper) OnRecvUpdatePostPacket(ctx sdk.Context, packet channeltypes.Packet, data types.UpdatePostPacketData) (packetAck types.UpdatePostPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic // Done
	postID, perr := strconv.ParseUint(data.PostID, 10, 64);

	if perr != nil {
		packetAck.IsSuccess = false;
		return packetAck, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot parse postID: %w", perr)
	}

	post, is_find := k.GetPost(ctx, postID);

	if !is_find {
		packetAck.IsSuccess = false;
		return packetAck, nil
	}

	post.Content = data.Content;
	post.Title = data.Title;

	k.SetPost(
		ctx,
		post,
	)

	packetAck.IsSuccess = true

	return packetAck, nil
}

// OnAcknowledgementUpdatePostPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementUpdatePostPacket(ctx sdk.Context, packet channeltypes.Packet, data types.UpdatePostPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.UpdatePostPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic 	// Done
		// Find the post by postID
		postID, perr := strconv.ParseUint(data.PostID, 10, 64);

		if perr != nil {
			return nil
		}

		if packetAck.IsSuccess {
			sentPost, is_find := k.GetSentPost(ctx, postID);

			if !is_find {
				return nil
			}
			
			// update SentPost and save
			sentPost.Title = data.Title;
			k.SetSentPost(ctx, sentPost);
		}

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutUpdatePostPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutUpdatePostPacket(ctx sdk.Context, packet channeltypes.Packet, data types.UpdatePostPacketData) error {

	// TODO: packet timeout logic

	return nil
}
