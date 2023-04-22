// @generated by protoc-gen-connect-es v0.8.6 with parameter "target=ts"
// @generated from file v1/compro_category.proto (package miiko.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodKind } from "@bufbuild/protobuf";
import { CategoryGetResponse, CategoryPostRequest, TopicGetRequest, TopicGetResponse } from "./compro_category_pb.js";

/**
 * @generated from service miiko.v1.ComproCategoryService
 */
export const ComproCategoryService = {
  typeName: "miiko.v1.ComproCategoryService",
  methods: {
    /**
     * @generated from rpc miiko.v1.ComproCategoryService.CategoryGet
     */
    categoryGet: {
      name: "CategoryGet",
      I: Empty,
      O: CategoryGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc miiko.v1.ComproCategoryService.CategoryPost
     */
    categoryPost: {
      name: "CategoryPost",
      I: CategoryPostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc miiko.v1.ComproCategoryService.TopicGet
     */
    topicGet: {
      name: "TopicGet",
      I: TopicGetRequest,
      O: TopicGetResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

