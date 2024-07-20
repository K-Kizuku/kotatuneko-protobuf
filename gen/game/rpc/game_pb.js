// @generated by protoc-gen-es v1.10.0
// @generated from file game/rpc/game.proto (package game.rpc, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { GameState, Hand, Object$, Player, Stat } from "../resources/game_pb.js";

/**
 * @generated from message game.rpc.GameStatusResponse
 */
export const GameStatusResponse = /*@__PURE__*/ proto3.makeMessageType(
  "game.rpc.GameStatusResponse",
  () => [
    { no: 1, name: "sender_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "room_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "state", kind: "enum", T: proto3.getEnumType(GameState) },
    { no: 4, name: "players", kind: "message", T: Player, repeated: true },
    { no: 5, name: "objects", kind: "message", T: Object$, repeated: true },
    { no: 6, name: "stats", kind: "message", T: Stat, repeated: true },
    { no: 7, name: "hands", kind: "message", T: Hand, repeated: true },
  ],
);

/**
 * @generated from message game.rpc.GameStatusRequest
 */
export const GameStatusRequest = /*@__PURE__*/ proto3.makeMessageType(
  "game.rpc.GameStatusRequest",
  () => [
    { no: 1, name: "sender_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "room_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "state", kind: "enum", T: proto3.getEnumType(GameState) },
    { no: 4, name: "players", kind: "message", T: Player, repeated: true },
    { no: 5, name: "objects", kind: "message", T: Object$, repeated: true },
    { no: 6, name: "stats", kind: "message", T: Stat, repeated: true },
    { no: 7, name: "hands", kind: "message", T: Hand, repeated: true },
  ],
);

/**
 * @generated from message game.rpc.PhysicsInitRequest
 */
export const PhysicsInitRequest = /*@__PURE__*/ proto3.makeMessageType(
  "game.rpc.PhysicsInitRequest",
  () => [
    { no: 1, name: "sender_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "room_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "objects", kind: "message", T: Object$, repeated: true },
  ],
);

/**
 * @generated from message game.rpc.PhysicsInitResponse
 */
export const PhysicsInitResponse = /*@__PURE__*/ proto3.makeMessageType(
  "game.rpc.PhysicsInitResponse",
  () => [
    { no: 1, name: "sender_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "room_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "state", kind: "enum", T: proto3.getEnumType(GameState) },
    { no: 4, name: "objects", kind: "message", T: Object$, repeated: true },
    { no: 5, name: "hands", kind: "message", T: Hand, repeated: true },
  ],
);

/**
 * @generated from message game.rpc.PhysicsRequest
 */
export const PhysicsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "game.rpc.PhysicsRequest",
  () => [
    { no: 1, name: "sender_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "room_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "hands", kind: "message", T: Hand },
  ],
);

/**
 * @generated from message game.rpc.PhysicsResponse
 */
export const PhysicsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "game.rpc.PhysicsResponse",
  () => [
    { no: 1, name: "sender_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "room_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "objects", kind: "message", T: Object$, repeated: true },
  ],
);
