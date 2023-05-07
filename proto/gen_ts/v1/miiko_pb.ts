// @generated by protoc-gen-es v1.2.0 with parameter "target=ts"
// @generated from file v1/miiko.proto (package miiko.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message miiko.v1.Category
 */
export class Category extends Message<Category> {
  /**
   * @generated from field: string category_id = 1;
   */
  categoryId = "";

  /**
   * @generated from field: string category_display_name = 2;
   */
  categoryDisplayName = "";

  /**
   * @generated from field: string category_system_name = 3;
   */
  categorySystemName = "";

  /**
   * @generated from field: int32 category_order = 4;
   */
  categoryOrder = 0;

  /**
   * @generated from field: int32 topic_size = 5;
   */
  topicSize = 0;

  /**
   * @generated from field: int32 problem_size = 6;
   */
  problemSize = 0;

  /**
   * @generated from field: repeated miiko.v1.Topic topic_list = 7;
   */
  topicList: Topic[] = [];

  constructor(data?: PartialMessage<Category>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.Category";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "category_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "category_display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "category_system_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "category_order", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "topic_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "problem_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "topic_list", kind: "message", T: Topic, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Category {
    return new Category().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Category {
    return new Category().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Category {
    return new Category().fromJsonString(jsonString, options);
  }

  static equals(a: Category | PlainMessage<Category> | undefined, b: Category | PlainMessage<Category> | undefined): boolean {
    return proto3.util.equals(Category, a, b);
  }
}

/**
 * @generated from message miiko.v1.Topic
 */
export class Topic extends Message<Topic> {
  /**
   * @generated from field: string topic_id = 1;
   */
  topicId = "";

  /**
   * @generated from field: string topic_display_name = 2;
   */
  topicDisplayName = "";

  /**
   * @generated from field: int32 topic_order = 3;
   */
  topicOrder = 0;

  /**
   * @generated from field: repeated miiko.v1.Problem problem_list = 4;
   */
  problemList: Problem[] = [];

  constructor(data?: PartialMessage<Topic>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.Topic";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "topic_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "topic_display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "topic_order", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "problem_list", kind: "message", T: Problem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Topic {
    return new Topic().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Topic {
    return new Topic().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Topic {
    return new Topic().fromJsonString(jsonString, options);
  }

  static equals(a: Topic | PlainMessage<Topic> | undefined, b: Topic | PlainMessage<Topic> | undefined): boolean {
    return proto3.util.equals(Topic, a, b);
  }
}

/**
 * @generated from message miiko.v1.Problem
 */
export class Problem extends Message<Problem> {
  /**
   * @generated from field: string problem_id = 1;
   */
  problemId = "";

  /**
   * @generated from field: string url = 2;
   */
  url = "";

  /**
   * @generated from field: string problem_display_name = 3;
   */
  problemDisplayName = "";

  /**
   * @generated from field: int32 estimation = 4;
   */
  estimation = 0;

  /**
   * @generated from field: repeated miiko.v1.Tag tag_list = 5;
   */
  tagList: Tag[] = [];

  constructor(data?: PartialMessage<Problem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.Problem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "problem_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "problem_display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "estimation", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "tag_list", kind: "message", T: Tag, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Problem {
    return new Problem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Problem {
    return new Problem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Problem {
    return new Problem().fromJsonString(jsonString, options);
  }

  static equals(a: Problem | PlainMessage<Problem> | undefined, b: Problem | PlainMessage<Problem> | undefined): boolean {
    return proto3.util.equals(Problem, a, b);
  }
}

/**
 * @generated from message miiko.v1.Tag
 */
export class Tag extends Message<Tag> {
  /**
   * @generated from field: string category_id = 1;
   */
  categoryId = "";

  /**
   * @generated from field: string topic_id = 2;
   */
  topicId = "";

  /**
   * @generated from field: string topic_display_name = 3;
   */
  topicDisplayName = "";

  constructor(data?: PartialMessage<Tag>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.Tag";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "category_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "topic_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "topic_display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Tag {
    return new Tag().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Tag {
    return new Tag().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Tag {
    return new Tag().fromJsonString(jsonString, options);
  }

  static equals(a: Tag | PlainMessage<Tag> | undefined, b: Tag | PlainMessage<Tag> | undefined): boolean {
    return proto3.util.equals(Tag, a, b);
  }
}

/**
 * req/res
 *
 * @generated from message miiko.v1.CategoryListGetRequest
 */
export class CategoryListGetRequest extends Message<CategoryListGetRequest> {
  /**
   * @generated from field: bool is_required_topic = 1;
   */
  isRequiredTopic = false;

  constructor(data?: PartialMessage<CategoryListGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.CategoryListGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "is_required_topic", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CategoryListGetRequest {
    return new CategoryListGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CategoryListGetRequest {
    return new CategoryListGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CategoryListGetRequest {
    return new CategoryListGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CategoryListGetRequest | PlainMessage<CategoryListGetRequest> | undefined, b: CategoryListGetRequest | PlainMessage<CategoryListGetRequest> | undefined): boolean {
    return proto3.util.equals(CategoryListGetRequest, a, b);
  }
}

/**
 * @generated from message miiko.v1.CategoryListGetResponse
 */
export class CategoryListGetResponse extends Message<CategoryListGetResponse> {
  /**
   * @generated from field: repeated miiko.v1.Category category_list = 1;
   */
  categoryList: Category[] = [];

  constructor(data?: PartialMessage<CategoryListGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.CategoryListGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "category_list", kind: "message", T: Category, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CategoryListGetResponse {
    return new CategoryListGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CategoryListGetResponse {
    return new CategoryListGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CategoryListGetResponse {
    return new CategoryListGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CategoryListGetResponse | PlainMessage<CategoryListGetResponse> | undefined, b: CategoryListGetResponse | PlainMessage<CategoryListGetResponse> | undefined): boolean {
    return proto3.util.equals(CategoryListGetResponse, a, b);
  }
}

/**
 * @generated from message miiko.v1.CategoryPostRequest
 */
export class CategoryPostRequest extends Message<CategoryPostRequest> {
  /**
   * @generated from field: string category_id = 1;
   */
  categoryId = "";

  /**
   * @generated from field: miiko.v1.Category category = 2;
   */
  category?: Category;

  constructor(data?: PartialMessage<CategoryPostRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.CategoryPostRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "category_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "category", kind: "message", T: Category },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CategoryPostRequest {
    return new CategoryPostRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CategoryPostRequest {
    return new CategoryPostRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CategoryPostRequest {
    return new CategoryPostRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CategoryPostRequest | PlainMessage<CategoryPostRequest> | undefined, b: CategoryPostRequest | PlainMessage<CategoryPostRequest> | undefined): boolean {
    return proto3.util.equals(CategoryPostRequest, a, b);
  }
}

/**
 * @generated from message miiko.v1.CategoryPostResponse
 */
export class CategoryPostResponse extends Message<CategoryPostResponse> {
  /**
   * @generated from field: string category_id = 1;
   */
  categoryId = "";

  constructor(data?: PartialMessage<CategoryPostResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.CategoryPostResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "category_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CategoryPostResponse {
    return new CategoryPostResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CategoryPostResponse {
    return new CategoryPostResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CategoryPostResponse {
    return new CategoryPostResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CategoryPostResponse | PlainMessage<CategoryPostResponse> | undefined, b: CategoryPostResponse | PlainMessage<CategoryPostResponse> | undefined): boolean {
    return proto3.util.equals(CategoryPostResponse, a, b);
  }
}

/**
 * @generated from message miiko.v1.TopicListGetRequest
 */
export class TopicListGetRequest extends Message<TopicListGetRequest> {
  /**
   * @generated from field: string category_system_name = 1;
   */
  categorySystemName = "";

  constructor(data?: PartialMessage<TopicListGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.TopicListGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "category_system_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TopicListGetRequest {
    return new TopicListGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TopicListGetRequest {
    return new TopicListGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TopicListGetRequest {
    return new TopicListGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: TopicListGetRequest | PlainMessage<TopicListGetRequest> | undefined, b: TopicListGetRequest | PlainMessage<TopicListGetRequest> | undefined): boolean {
    return proto3.util.equals(TopicListGetRequest, a, b);
  }
}

/**
 * @generated from message miiko.v1.TopicListGetResponse
 */
export class TopicListGetResponse extends Message<TopicListGetResponse> {
  /**
   * @generated from field: miiko.v1.Category category = 1;
   */
  category?: Category;

  /**
   * @generated from field: repeated miiko.v1.Topic topic_list = 2;
   */
  topicList: Topic[] = [];

  constructor(data?: PartialMessage<TopicListGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.TopicListGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "category", kind: "message", T: Category },
    { no: 2, name: "topic_list", kind: "message", T: Topic, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TopicListGetResponse {
    return new TopicListGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TopicListGetResponse {
    return new TopicListGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TopicListGetResponse {
    return new TopicListGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: TopicListGetResponse | PlainMessage<TopicListGetResponse> | undefined, b: TopicListGetResponse | PlainMessage<TopicListGetResponse> | undefined): boolean {
    return proto3.util.equals(TopicListGetResponse, a, b);
  }
}

/**
 * @generated from message miiko.v1.TopicPostRequest
 */
export class TopicPostRequest extends Message<TopicPostRequest> {
  /**
   * @generated from field: string topic_id = 1;
   */
  topicId = "";

  /**
   * @generated from field: miiko.v1.Topic topic = 2;
   */
  topic?: Topic;

  /**
   * @generated from field: string category_id = 3;
   */
  categoryId = "";

  constructor(data?: PartialMessage<TopicPostRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.TopicPostRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "topic_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "topic", kind: "message", T: Topic },
    { no: 3, name: "category_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TopicPostRequest {
    return new TopicPostRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TopicPostRequest {
    return new TopicPostRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TopicPostRequest {
    return new TopicPostRequest().fromJsonString(jsonString, options);
  }

  static equals(a: TopicPostRequest | PlainMessage<TopicPostRequest> | undefined, b: TopicPostRequest | PlainMessage<TopicPostRequest> | undefined): boolean {
    return proto3.util.equals(TopicPostRequest, a, b);
  }
}

/**
 * @generated from message miiko.v1.TopicPostResponse
 */
export class TopicPostResponse extends Message<TopicPostResponse> {
  /**
   * @generated from field: string topic_id = 1;
   */
  topicId = "";

  constructor(data?: PartialMessage<TopicPostResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.TopicPostResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "topic_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TopicPostResponse {
    return new TopicPostResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TopicPostResponse {
    return new TopicPostResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TopicPostResponse {
    return new TopicPostResponse().fromJsonString(jsonString, options);
  }

  static equals(a: TopicPostResponse | PlainMessage<TopicPostResponse> | undefined, b: TopicPostResponse | PlainMessage<TopicPostResponse> | undefined): boolean {
    return proto3.util.equals(TopicPostResponse, a, b);
  }
}

/**
 * @generated from message miiko.v1.ProblemListGetRequest
 */
export class ProblemListGetRequest extends Message<ProblemListGetRequest> {
  /**
   * @generated from field: int32 offset = 1;
   */
  offset = 0;

  /**
   * @generated from field: int32 limit = 2;
   */
  limit = 0;

  /**
   * @generated from field: miiko.v1.ProblemListGetRequest.SortType sort_type = 3;
   */
  sortType = ProblemListGetRequest_SortType.UNSPECIFIED;

  constructor(data?: PartialMessage<ProblemListGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.ProblemListGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "offset", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "sort_type", kind: "enum", T: proto3.getEnumType(ProblemListGetRequest_SortType) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProblemListGetRequest {
    return new ProblemListGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProblemListGetRequest {
    return new ProblemListGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProblemListGetRequest {
    return new ProblemListGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ProblemListGetRequest | PlainMessage<ProblemListGetRequest> | undefined, b: ProblemListGetRequest | PlainMessage<ProblemListGetRequest> | undefined): boolean {
    return proto3.util.equals(ProblemListGetRequest, a, b);
  }
}

/**
 * @generated from enum miiko.v1.ProblemListGetRequest.SortType
 */
export enum ProblemListGetRequest_SortType {
  /**
   * @generated from enum value: SORT_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: SORT_TYPE_CREATED_TIME = 1;
   */
  CREATED_TIME = 1,

  /**
   * @generated from enum value: SORT_TYPE_ESTIMATION = 2;
   */
  ESTIMATION = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(ProblemListGetRequest_SortType)
proto3.util.setEnumType(ProblemListGetRequest_SortType, "miiko.v1.ProblemListGetRequest.SortType", [
  { no: 0, name: "SORT_TYPE_UNSPECIFIED" },
  { no: 1, name: "SORT_TYPE_CREATED_TIME" },
  { no: 2, name: "SORT_TYPE_ESTIMATION" },
]);

/**
 * @generated from message miiko.v1.ProblemListGetResponse
 */
export class ProblemListGetResponse extends Message<ProblemListGetResponse> {
  /**
   * @generated from field: repeated miiko.v1.Problem problem_list = 1;
   */
  problemList: Problem[] = [];

  constructor(data?: PartialMessage<ProblemListGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.ProblemListGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "problem_list", kind: "message", T: Problem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProblemListGetResponse {
    return new ProblemListGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProblemListGetResponse {
    return new ProblemListGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProblemListGetResponse {
    return new ProblemListGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ProblemListGetResponse | PlainMessage<ProblemListGetResponse> | undefined, b: ProblemListGetResponse | PlainMessage<ProblemListGetResponse> | undefined): boolean {
    return proto3.util.equals(ProblemListGetResponse, a, b);
  }
}

/**
 * @generated from message miiko.v1.ProblemGetRequest
 */
export class ProblemGetRequest extends Message<ProblemGetRequest> {
  /**
   * @generated from field: string problem_id = 1;
   */
  problemId = "";

  constructor(data?: PartialMessage<ProblemGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.ProblemGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "problem_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProblemGetRequest {
    return new ProblemGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProblemGetRequest {
    return new ProblemGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProblemGetRequest {
    return new ProblemGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ProblemGetRequest | PlainMessage<ProblemGetRequest> | undefined, b: ProblemGetRequest | PlainMessage<ProblemGetRequest> | undefined): boolean {
    return proto3.util.equals(ProblemGetRequest, a, b);
  }
}

/**
 * @generated from message miiko.v1.ProblemGetResponse
 */
export class ProblemGetResponse extends Message<ProblemGetResponse> {
  /**
   * @generated from field: miiko.v1.Problem problem = 1;
   */
  problem?: Problem;

  constructor(data?: PartialMessage<ProblemGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.ProblemGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "problem", kind: "message", T: Problem },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProblemGetResponse {
    return new ProblemGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProblemGetResponse {
    return new ProblemGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProblemGetResponse {
    return new ProblemGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ProblemGetResponse | PlainMessage<ProblemGetResponse> | undefined, b: ProblemGetResponse | PlainMessage<ProblemGetResponse> | undefined): boolean {
    return proto3.util.equals(ProblemGetResponse, a, b);
  }
}

/**
 * @generated from message miiko.v1.ProblemPostRequest
 */
export class ProblemPostRequest extends Message<ProblemPostRequest> {
  /**
   * @generated from field: string problem_id = 1;
   */
  problemId = "";

  /**
   * @generated from field: miiko.v1.Problem problem = 2;
   */
  problem?: Problem;

  constructor(data?: PartialMessage<ProblemPostRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.ProblemPostRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "problem_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "problem", kind: "message", T: Problem },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProblemPostRequest {
    return new ProblemPostRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProblemPostRequest {
    return new ProblemPostRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProblemPostRequest {
    return new ProblemPostRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ProblemPostRequest | PlainMessage<ProblemPostRequest> | undefined, b: ProblemPostRequest | PlainMessage<ProblemPostRequest> | undefined): boolean {
    return proto3.util.equals(ProblemPostRequest, a, b);
  }
}

/**
 * @generated from message miiko.v1.ProblemPostResponse
 */
export class ProblemPostResponse extends Message<ProblemPostResponse> {
  /**
   * @generated from field: string problem_id = 1;
   */
  problemId = "";

  constructor(data?: PartialMessage<ProblemPostResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "miiko.v1.ProblemPostResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "problem_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProblemPostResponse {
    return new ProblemPostResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProblemPostResponse {
    return new ProblemPostResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProblemPostResponse {
    return new ProblemPostResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ProblemPostResponse | PlainMessage<ProblemPostResponse> | undefined, b: ProblemPostResponse | PlainMessage<ProblemPostResponse> | undefined): boolean {
    return proto3.util.equals(ProblemPostResponse, a, b);
  }
}

