///
//  Generated code. Do not modify.
//  source: api/tasks/issue/issue.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'issue.pb.dart' as $7;
import 'issue.pbjson.dart';

export 'issue.pb.dart';

abstract class IssueRegistryServiceBase extends $pb.GeneratedService {
  $async.Future<$7.IssueCreateResponse> createIssue($pb.ServerContext ctx, $7.IssueCreateRequest request);
  $async.Future<$7.IssueCreateSubtaskResponse> createSubtask($pb.ServerContext ctx, $7.IssueCreateSubtaskRequest request);
  $async.Future<$7.IssueAddCommentResponse> addComment($pb.ServerContext ctx, $7.IssueAddCommentRequest request);
  $async.Future<$7.IssueUpdateResponse> updateIssue($pb.ServerContext ctx, $7.IssueUpdateRequest request);
  $async.Future<$7.IssueGetResponse> getIssues($pb.ServerContext ctx, $7.IssueGetRequest request);
  $async.Future<$7.IssueGetByIDResponse> getIssueByID($pb.ServerContext ctx, $7.IssueGetByIDRequest request);
  $async.Future<$7.IssueInfoGetResponse> getIssueInfo($pb.ServerContext ctx, $7.IssueInfoGetRequest request);
  $async.Future<$7.IssueInfoGetByIDResponse> getIssueInfoByID($pb.ServerContext ctx, $7.IssueInfoGetByIDRequest request);
  $async.Future<$7.IssuePaymentGetResponse> getUserPayment($pb.ServerContext ctx, $7.IssuePaymentGetRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'CreateIssue': return $7.IssueCreateRequest();
      case 'CreateSubtask': return $7.IssueCreateSubtaskRequest();
      case 'AddComment': return $7.IssueAddCommentRequest();
      case 'UpdateIssue': return $7.IssueUpdateRequest();
      case 'GetIssues': return $7.IssueGetRequest();
      case 'GetIssueByID': return $7.IssueGetByIDRequest();
      case 'GetIssueInfo': return $7.IssueInfoGetRequest();
      case 'GetIssueInfoByID': return $7.IssueInfoGetByIDRequest();
      case 'GetUserPayment': return $7.IssuePaymentGetRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'CreateIssue': return this.createIssue(ctx, request as $7.IssueCreateRequest);
      case 'CreateSubtask': return this.createSubtask(ctx, request as $7.IssueCreateSubtaskRequest);
      case 'AddComment': return this.addComment(ctx, request as $7.IssueAddCommentRequest);
      case 'UpdateIssue': return this.updateIssue(ctx, request as $7.IssueUpdateRequest);
      case 'GetIssues': return this.getIssues(ctx, request as $7.IssueGetRequest);
      case 'GetIssueByID': return this.getIssueByID(ctx, request as $7.IssueGetByIDRequest);
      case 'GetIssueInfo': return this.getIssueInfo(ctx, request as $7.IssueInfoGetRequest);
      case 'GetIssueInfoByID': return this.getIssueInfoByID(ctx, request as $7.IssueInfoGetByIDRequest);
      case 'GetUserPayment': return this.getUserPayment(ctx, request as $7.IssuePaymentGetRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => IssueRegistryServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => IssueRegistryServiceBase$messageJson;
}

