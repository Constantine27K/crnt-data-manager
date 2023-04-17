///
//  Generated code. Do not modify.
//  source: api/project/project.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'project.pb.dart' as $2;
import 'project.pbjson.dart';

export 'project.pb.dart';

abstract class ProjectRegistryServiceBase extends $pb.GeneratedService {
  $async.Future<$2.ProjectCreateResponse> createProject($pb.ServerContext ctx, $2.ProjectCreateRequest request);
  $async.Future<$2.ProjectAddTeamResponse> addResponsibleTeam($pb.ServerContext ctx, $2.ProjectAddTeamRequest request);
  $async.Future<$2.ProjectRemoveTeamResponse> removeResponsibleTeam($pb.ServerContext ctx, $2.ProjectRemoveTeamRequest request);
  $async.Future<$2.ProjectUpdateResponse> updateProject($pb.ServerContext ctx, $2.ProjectUpdateRequest request);
  $async.Future<$2.ProjectGetResponse> getProjects($pb.ServerContext ctx, $2.ProjectGetRequest request);
  $async.Future<$2.ProjectGetByIDResponse> getProjectByID($pb.ServerContext ctx, $2.ProjectGetByIDRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'CreateProject': return $2.ProjectCreateRequest();
      case 'AddResponsibleTeam': return $2.ProjectAddTeamRequest();
      case 'RemoveResponsibleTeam': return $2.ProjectRemoveTeamRequest();
      case 'UpdateProject': return $2.ProjectUpdateRequest();
      case 'GetProjects': return $2.ProjectGetRequest();
      case 'GetProjectByID': return $2.ProjectGetByIDRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'CreateProject': return this.createProject(ctx, request as $2.ProjectCreateRequest);
      case 'AddResponsibleTeam': return this.addResponsibleTeam(ctx, request as $2.ProjectAddTeamRequest);
      case 'RemoveResponsibleTeam': return this.removeResponsibleTeam(ctx, request as $2.ProjectRemoveTeamRequest);
      case 'UpdateProject': return this.updateProject(ctx, request as $2.ProjectUpdateRequest);
      case 'GetProjects': return this.getProjects(ctx, request as $2.ProjectGetRequest);
      case 'GetProjectByID': return this.getProjectByID(ctx, request as $2.ProjectGetByIDRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => ProjectRegistryServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => ProjectRegistryServiceBase$messageJson;
}

