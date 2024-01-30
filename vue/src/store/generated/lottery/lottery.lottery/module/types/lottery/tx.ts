/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "lottery.lottery";

export interface MsgCreateLotteryItem {
  creator: string;
  amount: number;
  participants: string[];
}

export interface MsgCreateLotteryItemResponse {}

export interface MsgUpdateLotteryItem {
  creator: string;
  amount: number;
  participants: string[];
}

export interface MsgUpdateLotteryItemResponse {}

export interface MsgDeleteLotteryItem {
  creator: string;
}

export interface MsgDeleteLotteryItemResponse {}

const baseMsgCreateLotteryItem: object = {
  creator: "",
  amount: 0,
  participants: "",
};

export const MsgCreateLotteryItem = {
  encode(
    message: MsgCreateLotteryItem,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== 0) {
      writer.uint32(24).int32(message.amount);
    }
    for (const v of message.participants) {
      writer.uint32(34).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateLotteryItem {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateLotteryItem } as MsgCreateLotteryItem;
    message.participants = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 3:
          message.amount = reader.int32();
          break;
        case 4:
          message.participants.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateLotteryItem {
    const message = { ...baseMsgCreateLotteryItem } as MsgCreateLotteryItem;
    message.participants = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount);
    } else {
      message.amount = 0;
    }
    if (object.participants !== undefined && object.participants !== null) {
      for (const e of object.participants) {
        message.participants.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: MsgCreateLotteryItem): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    if (message.participants) {
      obj.participants = message.participants.map((e) => e);
    } else {
      obj.participants = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateLotteryItem>): MsgCreateLotteryItem {
    const message = { ...baseMsgCreateLotteryItem } as MsgCreateLotteryItem;
    message.participants = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = 0;
    }
    if (object.participants !== undefined && object.participants !== null) {
      for (const e of object.participants) {
        message.participants.push(e);
      }
    }
    return message;
  },
};

const baseMsgCreateLotteryItemResponse: object = {};

export const MsgCreateLotteryItemResponse = {
  encode(
    _: MsgCreateLotteryItemResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateLotteryItemResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateLotteryItemResponse,
    } as MsgCreateLotteryItemResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateLotteryItemResponse {
    const message = {
      ...baseMsgCreateLotteryItemResponse,
    } as MsgCreateLotteryItemResponse;
    return message;
  },

  toJSON(_: MsgCreateLotteryItemResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateLotteryItemResponse>
  ): MsgCreateLotteryItemResponse {
    const message = {
      ...baseMsgCreateLotteryItemResponse,
    } as MsgCreateLotteryItemResponse;
    return message;
  },
};

const baseMsgUpdateLotteryItem: object = {
  creator: "",
  amount: 0,
  participants: "",
};

export const MsgUpdateLotteryItem = {
  encode(
    message: MsgUpdateLotteryItem,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== 0) {
      writer.uint32(24).int32(message.amount);
    }
    for (const v of message.participants) {
      writer.uint32(34).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateLotteryItem {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateLotteryItem } as MsgUpdateLotteryItem;
    message.participants = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 3:
          message.amount = reader.int32();
          break;
        case 4:
          message.participants.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateLotteryItem {
    const message = { ...baseMsgUpdateLotteryItem } as MsgUpdateLotteryItem;
    message.participants = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount);
    } else {
      message.amount = 0;
    }
    if (object.participants !== undefined && object.participants !== null) {
      for (const e of object.participants) {
        message.participants.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: MsgUpdateLotteryItem): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    if (message.participants) {
      obj.participants = message.participants.map((e) => e);
    } else {
      obj.participants = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateLotteryItem>): MsgUpdateLotteryItem {
    const message = { ...baseMsgUpdateLotteryItem } as MsgUpdateLotteryItem;
    message.participants = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = 0;
    }
    if (object.participants !== undefined && object.participants !== null) {
      for (const e of object.participants) {
        message.participants.push(e);
      }
    }
    return message;
  },
};

const baseMsgUpdateLotteryItemResponse: object = {};

export const MsgUpdateLotteryItemResponse = {
  encode(
    _: MsgUpdateLotteryItemResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateLotteryItemResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateLotteryItemResponse,
    } as MsgUpdateLotteryItemResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateLotteryItemResponse {
    const message = {
      ...baseMsgUpdateLotteryItemResponse,
    } as MsgUpdateLotteryItemResponse;
    return message;
  },

  toJSON(_: MsgUpdateLotteryItemResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateLotteryItemResponse>
  ): MsgUpdateLotteryItemResponse {
    const message = {
      ...baseMsgUpdateLotteryItemResponse,
    } as MsgUpdateLotteryItemResponse;
    return message;
  },
};

const baseMsgDeleteLotteryItem: object = { creator: "" };

export const MsgDeleteLotteryItem = {
  encode(
    message: MsgDeleteLotteryItem,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteLotteryItem {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteLotteryItem } as MsgDeleteLotteryItem;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteLotteryItem {
    const message = { ...baseMsgDeleteLotteryItem } as MsgDeleteLotteryItem;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteLotteryItem): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteLotteryItem>): MsgDeleteLotteryItem {
    const message = { ...baseMsgDeleteLotteryItem } as MsgDeleteLotteryItem;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgDeleteLotteryItemResponse: object = {};

export const MsgDeleteLotteryItemResponse = {
  encode(
    _: MsgDeleteLotteryItemResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteLotteryItemResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteLotteryItemResponse,
    } as MsgDeleteLotteryItemResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteLotteryItemResponse {
    const message = {
      ...baseMsgDeleteLotteryItemResponse,
    } as MsgDeleteLotteryItemResponse;
    return message;
  },

  toJSON(_: MsgDeleteLotteryItemResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteLotteryItemResponse>
  ): MsgDeleteLotteryItemResponse {
    const message = {
      ...baseMsgDeleteLotteryItemResponse,
    } as MsgDeleteLotteryItemResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateLotteryItem(
    request: MsgCreateLotteryItem
  ): Promise<MsgCreateLotteryItemResponse>;
  UpdateLotteryItem(
    request: MsgUpdateLotteryItem
  ): Promise<MsgUpdateLotteryItemResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteLotteryItem(
    request: MsgDeleteLotteryItem
  ): Promise<MsgDeleteLotteryItemResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateLotteryItem(
    request: MsgCreateLotteryItem
  ): Promise<MsgCreateLotteryItemResponse> {
    const data = MsgCreateLotteryItem.encode(request).finish();
    const promise = this.rpc.request(
      "lottery.lottery.Msg",
      "CreateLotteryItem",
      data
    );
    return promise.then((data) =>
      MsgCreateLotteryItemResponse.decode(new Reader(data))
    );
  }

  UpdateLotteryItem(
    request: MsgUpdateLotteryItem
  ): Promise<MsgUpdateLotteryItemResponse> {
    const data = MsgUpdateLotteryItem.encode(request).finish();
    const promise = this.rpc.request(
      "lottery.lottery.Msg",
      "UpdateLotteryItem",
      data
    );
    return promise.then((data) =>
      MsgUpdateLotteryItemResponse.decode(new Reader(data))
    );
  }

  DeleteLotteryItem(
    request: MsgDeleteLotteryItem
  ): Promise<MsgDeleteLotteryItemResponse> {
    const data = MsgDeleteLotteryItem.encode(request).finish();
    const promise = this.rpc.request(
      "lottery.lottery.Msg",
      "DeleteLotteryItem",
      data
    );
    return promise.then((data) =>
      MsgDeleteLotteryItemResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
