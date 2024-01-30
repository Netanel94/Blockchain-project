// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateLotteryItem } from "./types/lottery/tx";
import { MsgUpdateLotteryItem } from "./types/lottery/tx";
import { MsgDeleteLotteryItem } from "./types/lottery/tx";


const types = [
  ["/lottery.lottery.MsgCreateLotteryItem", MsgCreateLotteryItem],
  ["/lottery.lottery.MsgUpdateLotteryItem", MsgUpdateLotteryItem],
  ["/lottery.lottery.MsgDeleteLotteryItem", MsgDeleteLotteryItem],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateLotteryItem: (data: MsgCreateLotteryItem): EncodeObject => ({ typeUrl: "/lottery.lottery.MsgCreateLotteryItem", value: MsgCreateLotteryItem.fromPartial( data ) }),
    msgUpdateLotteryItem: (data: MsgUpdateLotteryItem): EncodeObject => ({ typeUrl: "/lottery.lottery.MsgUpdateLotteryItem", value: MsgUpdateLotteryItem.fromPartial( data ) }),
    msgDeleteLotteryItem: (data: MsgDeleteLotteryItem): EncodeObject => ({ typeUrl: "/lottery.lottery.MsgDeleteLotteryItem", value: MsgDeleteLotteryItem.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
