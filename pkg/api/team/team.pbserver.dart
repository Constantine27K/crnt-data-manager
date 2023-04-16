///
//  Generated code. Do not modify.
//  source: api/team/team.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'team.pb.dart' as $9;
import 'team.pbjson.dart';

export 'team.pb.dart';

abstract class TeamRegistryServiceBase extends $pb.GeneratedService {
  $async.Future<$9.TeamCreateResponse> createTeam($pb.ServerContext ctx, $9.TeamCreateRequest request);
  $async.Future<$9.TeamUpdateResponse> updateTeam($pb.ServerContext ctx, $9.TeamUpdateRequest request);
  $async.Future<$9.TeamGetResponse> getTeams($pb.ServerContext ctx, $9.TeamGetRequest request);
  $async.Future<$9.TeamGetByIDResponse> getTeamByID($pb.ServerContext ctx, $9.TeamGetByIDRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'CreateTeam': return $9.TeamCreateRequest();
      case 'UpdateTeam': return $9.TeamUpdateRequest();
      case 'GetTeams': return $9.TeamGetRequest();
      case 'GetTeamByID': return $9.TeamGetByIDRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'CreateTeam': return this.createTeam(ctx, request as $9.TeamCreateRequest);
      case 'UpdateTeam': return this.updateTeam(ctx, request as $9.TeamUpdateRequest);
      case 'GetTeams': return this.getTeams(ctx, request as $9.TeamGetRequest);
      case 'GetTeamByID': return this.getTeamByID(ctx, request as $9.TeamGetByIDRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => TeamRegistryServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => TeamRegistryServiceBase$messageJson;
}

