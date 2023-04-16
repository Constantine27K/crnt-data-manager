///
//  Generated code. Do not modify.
//  source: api/sprint/sprint.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'sprint.pb.dart' as $8;
import 'sprint.pbjson.dart';

export 'sprint.pb.dart';

abstract class SprintRegistryServiceBase extends $pb.GeneratedService {
  $async.Future<$8.SprintCreateResponse> createSprint($pb.ServerContext ctx, $8.SprintCreateRequest request);
  $async.Future<$8.AddIssueResponse> addIssue($pb.ServerContext ctx, $8.AddIssueRequest request);
  $async.Future<$8.RemoveIssueResponse> removeIssue($pb.ServerContext ctx, $8.RemoveIssueRequest request);
  $async.Future<$8.SprintUpdateResponse> updateSprint($pb.ServerContext ctx, $8.SprintUpdateRequest request);
  $async.Future<$8.SprintChangeStatusResponse> sprintChangeStatus($pb.ServerContext ctx, $8.SprintChangeStatusRequest request);
  $async.Future<$8.SprintGetResponse> getSprint($pb.ServerContext ctx, $8.SprintGetRequest request);
  $async.Future<$8.SprintGetByIDResponse> getSprintByID($pb.ServerContext ctx, $8.SprintGetByIDRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'CreateSprint': return $8.SprintCreateRequest();
      case 'AddIssue': return $8.AddIssueRequest();
      case 'RemoveIssue': return $8.RemoveIssueRequest();
      case 'UpdateSprint': return $8.SprintUpdateRequest();
      case 'SprintChangeStatus': return $8.SprintChangeStatusRequest();
      case 'GetSprint': return $8.SprintGetRequest();
      case 'GetSprintByID': return $8.SprintGetByIDRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'CreateSprint': return this.createSprint(ctx, request as $8.SprintCreateRequest);
      case 'AddIssue': return this.addIssue(ctx, request as $8.AddIssueRequest);
      case 'RemoveIssue': return this.removeIssue(ctx, request as $8.RemoveIssueRequest);
      case 'UpdateSprint': return this.updateSprint(ctx, request as $8.SprintUpdateRequest);
      case 'SprintChangeStatus': return this.sprintChangeStatus(ctx, request as $8.SprintChangeStatusRequest);
      case 'GetSprint': return this.getSprint(ctx, request as $8.SprintGetRequest);
      case 'GetSprintByID': return this.getSprintByID(ctx, request as $8.SprintGetByIDRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => SprintRegistryServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => SprintRegistryServiceBase$messageJson;
}

