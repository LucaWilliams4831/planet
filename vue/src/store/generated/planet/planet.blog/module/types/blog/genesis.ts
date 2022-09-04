/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../blog/params";
import { Post } from "../blog/post";
import { SendPost } from "../blog/send_post";
import { TimedoutPost } from "../blog/timedout_post";

export const protobufPackage = "planet.blog";

/** GenesisState defines the blog module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  port_id: string;
  postList: Post[];
  postCount: number;
  sendPostList: SendPost[];
  sendPostCount: number;
  timedoutPostList: TimedoutPost[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  timedoutPostCount: number;
}

const baseGenesisState: object = {
  port_id: "",
  postCount: 0,
  sendPostCount: 0,
  timedoutPostCount: 0,
};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.port_id !== "") {
      writer.uint32(18).string(message.port_id);
    }
    for (const v of message.postList) {
      Post.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.postCount !== 0) {
      writer.uint32(32).uint64(message.postCount);
    }
    for (const v of message.sendPostList) {
      SendPost.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.sendPostCount !== 0) {
      writer.uint32(48).uint64(message.sendPostCount);
    }
    for (const v of message.timedoutPostList) {
      TimedoutPost.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    if (message.timedoutPostCount !== 0) {
      writer.uint32(64).uint64(message.timedoutPostCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.postList = [];
    message.sendPostList = [];
    message.timedoutPostList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.port_id = reader.string();
          break;
        case 3:
          message.postList.push(Post.decode(reader, reader.uint32()));
          break;
        case 4:
          message.postCount = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.sendPostList.push(SendPost.decode(reader, reader.uint32()));
          break;
        case 6:
          message.sendPostCount = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.timedoutPostList.push(
            TimedoutPost.decode(reader, reader.uint32())
          );
          break;
        case 8:
          message.timedoutPostCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.postList = [];
    message.sendPostList = [];
    message.timedoutPostList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.port_id !== undefined && object.port_id !== null) {
      message.port_id = String(object.port_id);
    } else {
      message.port_id = "";
    }
    if (object.postList !== undefined && object.postList !== null) {
      for (const e of object.postList) {
        message.postList.push(Post.fromJSON(e));
      }
    }
    if (object.postCount !== undefined && object.postCount !== null) {
      message.postCount = Number(object.postCount);
    } else {
      message.postCount = 0;
    }
    if (object.sendPostList !== undefined && object.sendPostList !== null) {
      for (const e of object.sendPostList) {
        message.sendPostList.push(SendPost.fromJSON(e));
      }
    }
    if (object.sendPostCount !== undefined && object.sendPostCount !== null) {
      message.sendPostCount = Number(object.sendPostCount);
    } else {
      message.sendPostCount = 0;
    }
    if (
      object.timedoutPostList !== undefined &&
      object.timedoutPostList !== null
    ) {
      for (const e of object.timedoutPostList) {
        message.timedoutPostList.push(TimedoutPost.fromJSON(e));
      }
    }
    if (
      object.timedoutPostCount !== undefined &&
      object.timedoutPostCount !== null
    ) {
      message.timedoutPostCount = Number(object.timedoutPostCount);
    } else {
      message.timedoutPostCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.port_id !== undefined && (obj.port_id = message.port_id);
    if (message.postList) {
      obj.postList = message.postList.map((e) =>
        e ? Post.toJSON(e) : undefined
      );
    } else {
      obj.postList = [];
    }
    message.postCount !== undefined && (obj.postCount = message.postCount);
    if (message.sendPostList) {
      obj.sendPostList = message.sendPostList.map((e) =>
        e ? SendPost.toJSON(e) : undefined
      );
    } else {
      obj.sendPostList = [];
    }
    message.sendPostCount !== undefined &&
      (obj.sendPostCount = message.sendPostCount);
    if (message.timedoutPostList) {
      obj.timedoutPostList = message.timedoutPostList.map((e) =>
        e ? TimedoutPost.toJSON(e) : undefined
      );
    } else {
      obj.timedoutPostList = [];
    }
    message.timedoutPostCount !== undefined &&
      (obj.timedoutPostCount = message.timedoutPostCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.postList = [];
    message.sendPostList = [];
    message.timedoutPostList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.port_id !== undefined && object.port_id !== null) {
      message.port_id = object.port_id;
    } else {
      message.port_id = "";
    }
    if (object.postList !== undefined && object.postList !== null) {
      for (const e of object.postList) {
        message.postList.push(Post.fromPartial(e));
      }
    }
    if (object.postCount !== undefined && object.postCount !== null) {
      message.postCount = object.postCount;
    } else {
      message.postCount = 0;
    }
    if (object.sendPostList !== undefined && object.sendPostList !== null) {
      for (const e of object.sendPostList) {
        message.sendPostList.push(SendPost.fromPartial(e));
      }
    }
    if (object.sendPostCount !== undefined && object.sendPostCount !== null) {
      message.sendPostCount = object.sendPostCount;
    } else {
      message.sendPostCount = 0;
    }
    if (
      object.timedoutPostList !== undefined &&
      object.timedoutPostList !== null
    ) {
      for (const e of object.timedoutPostList) {
        message.timedoutPostList.push(TimedoutPost.fromPartial(e));
      }
    }
    if (
      object.timedoutPostCount !== undefined &&
      object.timedoutPostCount !== null
    ) {
      message.timedoutPostCount = object.timedoutPostCount;
    } else {
      message.timedoutPostCount = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
