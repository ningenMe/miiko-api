// @generated by protoc-gen-es v1.2.0 with parameter "target=ts"
// @generated from file v1/health.proto (package miiko.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message miiko.v1.HealthServiceCheckRequest
 */
export class HealthServiceCheckRequest extends Message<HealthServiceCheckRequest> {
  /**
   * @generated from field: string service = 1;
   */
  service = "";

  constructor(data?: PartialMessage<HealthServiceCheckRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.HealthServiceCheckRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "service", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HealthServiceCheckRequest {
    return new HealthServiceCheckRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HealthServiceCheckRequest {
    return new HealthServiceCheckRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HealthServiceCheckRequest {
    return new HealthServiceCheckRequest().fromJsonString(jsonString, options);
  }

  static equals(a: HealthServiceCheckRequest | PlainMessage<HealthServiceCheckRequest> | undefined, b: HealthServiceCheckRequest | PlainMessage<HealthServiceCheckRequest> | undefined): boolean {
    return proto3.util.equals(HealthServiceCheckRequest, a, b);
  }
}

/**
 * @generated from message miiko.v1.HealthServiceCheckResponse
 */
export class HealthServiceCheckResponse extends Message<HealthServiceCheckResponse> {
  /**
   * @generated from field: miiko.v1.HealthServiceCheckResponse.ServingStatus status = 1;
   */
  status = HealthServiceCheckResponse_ServingStatus.UNSPECIFIED;

  constructor(data?: PartialMessage<HealthServiceCheckResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.HealthServiceCheckResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "status", kind: "enum", T: proto3.getEnumType(HealthServiceCheckResponse_ServingStatus) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HealthServiceCheckResponse {
    return new HealthServiceCheckResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HealthServiceCheckResponse {
    return new HealthServiceCheckResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HealthServiceCheckResponse {
    return new HealthServiceCheckResponse().fromJsonString(jsonString, options);
  }

  static equals(a: HealthServiceCheckResponse | PlainMessage<HealthServiceCheckResponse> | undefined, b: HealthServiceCheckResponse | PlainMessage<HealthServiceCheckResponse> | undefined): boolean {
    return proto3.util.equals(HealthServiceCheckResponse, a, b);
  }
}

/**
 * @generated from enum miiko.v1.HealthServiceCheckResponse.ServingStatus
 */
export enum HealthServiceCheckResponse_ServingStatus {
  /**
   * @generated from enum value: SERVING_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: SERVING_STATUS_SERVING = 1;
   */
  SERVING = 1,

  /**
   * @generated from enum value: SERVING_STATUS_NOT_SERVING = 2;
   */
  NOT_SERVING = 2,

  /**
   * Used only by the Watch method.
   *
   * @generated from enum value: SERVING_STATUS_SERVICE_UNKNOWN = 3;
   */
  SERVICE_UNKNOWN = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(HealthServiceCheckResponse_ServingStatus)
proto3.util.setEnumType(HealthServiceCheckResponse_ServingStatus, "miiko.v1.HealthServiceCheckResponse.ServingStatus", [
  { no: 0, name: "SERVING_STATUS_UNSPECIFIED" },
  { no: 1, name: "SERVING_STATUS_SERVING" },
  { no: 2, name: "SERVING_STATUS_NOT_SERVING" },
  { no: 3, name: "SERVING_STATUS_SERVICE_UNKNOWN" },
]);

