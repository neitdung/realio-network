/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'realiotech.network.asset';
const baseToken = { index: '', name: '', symbol: '', total: 0, decimals: '', authorizationRequired: false, creator: '' };
export const Token = {
    encode(message, writer = Writer.create()) {
        if (message.index !== '') {
            writer.uint32(10).string(message.index);
        }
        if (message.name !== '') {
            writer.uint32(18).string(message.name);
        }
        if (message.symbol !== '') {
            writer.uint32(26).string(message.symbol);
        }
        if (message.total !== 0) {
            writer.uint32(32).int64(message.total);
        }
        if (message.decimals !== '') {
            writer.uint32(42).string(message.decimals);
        }
        if (message.authorizationRequired === true) {
            writer.uint32(48).bool(message.authorizationRequired);
        }
        if (message.creator !== '') {
            writer.uint32(58).string(message.creator);
        }
        Object.entries(message.authorized).forEach(([key, value]) => {
            Token_AuthorizedEntry.encode({ key: key, value }, writer.uint32(66).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseToken };
        message.authorized = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.index = reader.string();
                    break;
                case 2:
                    message.name = reader.string();
                    break;
                case 3:
                    message.symbol = reader.string();
                    break;
                case 4:
                    message.total = longToNumber(reader.int64());
                    break;
                case 5:
                    message.decimals = reader.string();
                    break;
                case 6:
                    message.authorizationRequired = reader.bool();
                    break;
                case 7:
                    message.creator = reader.string();
                    break;
                case 8:
                    const entry8 = Token_AuthorizedEntry.decode(reader, reader.uint32());
                    if (entry8.value !== undefined) {
                        message.authorized[entry8.key] = entry8.value;
                    }
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseToken };
        message.authorized = {};
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = '';
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = '';
        }
        if (object.symbol !== undefined && object.symbol !== null) {
            message.symbol = String(object.symbol);
        }
        else {
            message.symbol = '';
        }
        if (object.total !== undefined && object.total !== null) {
            message.total = Number(object.total);
        }
        else {
            message.total = 0;
        }
        if (object.decimals !== undefined && object.decimals !== null) {
            message.decimals = String(object.decimals);
        }
        else {
            message.decimals = '';
        }
        if (object.authorizationRequired !== undefined && object.authorizationRequired !== null) {
            message.authorizationRequired = Boolean(object.authorizationRequired);
        }
        else {
            message.authorizationRequired = false;
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.authorized !== undefined && object.authorized !== null) {
            Object.entries(object.authorized).forEach(([key, value]) => {
                message.authorized[key] = TokenAuthorization.fromJSON(value);
            });
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.index !== undefined && (obj.index = message.index);
        message.name !== undefined && (obj.name = message.name);
        message.symbol !== undefined && (obj.symbol = message.symbol);
        message.total !== undefined && (obj.total = message.total);
        message.decimals !== undefined && (obj.decimals = message.decimals);
        message.authorizationRequired !== undefined && (obj.authorizationRequired = message.authorizationRequired);
        message.creator !== undefined && (obj.creator = message.creator);
        obj.authorized = {};
        if (message.authorized) {
            Object.entries(message.authorized).forEach(([k, v]) => {
                obj.authorized[k] = TokenAuthorization.toJSON(v);
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseToken };
        message.authorized = {};
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = '';
        }
        if (object.symbol !== undefined && object.symbol !== null) {
            message.symbol = object.symbol;
        }
        else {
            message.symbol = '';
        }
        if (object.total !== undefined && object.total !== null) {
            message.total = object.total;
        }
        else {
            message.total = 0;
        }
        if (object.decimals !== undefined && object.decimals !== null) {
            message.decimals = object.decimals;
        }
        else {
            message.decimals = '';
        }
        if (object.authorizationRequired !== undefined && object.authorizationRequired !== null) {
            message.authorizationRequired = object.authorizationRequired;
        }
        else {
            message.authorizationRequired = false;
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.authorized !== undefined && object.authorized !== null) {
            Object.entries(object.authorized).forEach(([key, value]) => {
                if (value !== undefined) {
                    message.authorized[key] = TokenAuthorization.fromPartial(value);
                }
            });
        }
        return message;
    }
};
const baseToken_AuthorizedEntry = { key: '' };
export const Token_AuthorizedEntry = {
    encode(message, writer = Writer.create()) {
        if (message.key !== '') {
            writer.uint32(10).string(message.key);
        }
        if (message.value !== undefined) {
            TokenAuthorization.encode(message.value, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseToken_AuthorizedEntry };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.key = reader.string();
                    break;
                case 2:
                    message.value = TokenAuthorization.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseToken_AuthorizedEntry };
        if (object.key !== undefined && object.key !== null) {
            message.key = String(object.key);
        }
        else {
            message.key = '';
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = TokenAuthorization.fromJSON(object.value);
        }
        else {
            message.value = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.key !== undefined && (obj.key = message.key);
        message.value !== undefined && (obj.value = message.value ? TokenAuthorization.toJSON(message.value) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseToken_AuthorizedEntry };
        if (object.key !== undefined && object.key !== null) {
            message.key = object.key;
        }
        else {
            message.key = '';
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = TokenAuthorization.fromPartial(object.value);
        }
        else {
            message.value = undefined;
        }
        return message;
    }
};
const baseTokenAuthorization = { tokenIndex: '', address: '', authorized: false };
export const TokenAuthorization = {
    encode(message, writer = Writer.create()) {
        if (message.tokenIndex !== '') {
            writer.uint32(10).string(message.tokenIndex);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
        }
        if (message.authorized === true) {
            writer.uint32(24).bool(message.authorized);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseTokenAuthorization };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.tokenIndex = reader.string();
                    break;
                case 2:
                    message.address = reader.string();
                    break;
                case 3:
                    message.authorized = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseTokenAuthorization };
        if (object.tokenIndex !== undefined && object.tokenIndex !== null) {
            message.tokenIndex = String(object.tokenIndex);
        }
        else {
            message.tokenIndex = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.authorized !== undefined && object.authorized !== null) {
            message.authorized = Boolean(object.authorized);
        }
        else {
            message.authorized = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.tokenIndex !== undefined && (obj.tokenIndex = message.tokenIndex);
        message.address !== undefined && (obj.address = message.address);
        message.authorized !== undefined && (obj.authorized = message.authorized);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseTokenAuthorization };
        if (object.tokenIndex !== undefined && object.tokenIndex !== null) {
            message.tokenIndex = object.tokenIndex;
        }
        else {
            message.tokenIndex = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.authorized !== undefined && object.authorized !== null) {
            message.authorized = object.authorized;
        }
        else {
            message.authorized = false;
        }
        return message;
    }
};
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}