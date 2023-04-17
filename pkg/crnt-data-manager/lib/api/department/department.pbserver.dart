///
//  Generated code. Do not modify.
//  source: api/department/department.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'department.pb.dart' as $1;
import 'department.pbjson.dart';

export 'department.pb.dart';

abstract class DepartmentRegistryServiceBase extends $pb.GeneratedService {
  $async.Future<$1.DepartmentCreateResponse> createDepartment($pb.ServerContext ctx, $1.DepartmentCreateRequest request);
  $async.Future<$1.DepartmentAddProjectResponse> departmentAddProject($pb.ServerContext ctx, $1.DepartmentAddProjectRequest request);
  $async.Future<$1.DepartmentRemoveProjectResponse> departmentRemoveProject($pb.ServerContext ctx, $1.DepartmentRemoveProjectRequest request);
  $async.Future<$1.DepartmentGetResponse> getDepartments($pb.ServerContext ctx, $1.DepartmentGetRequest request);
  $async.Future<$1.DepartmentGetByIDResponse> getDepartmentByID($pb.ServerContext ctx, $1.DepartmentGetByIDRequest request);
  $async.Future<$1.DepartmentUpdateResponse> updateDepartment($pb.ServerContext ctx, $1.DepartmentUpdateRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'CreateDepartment': return $1.DepartmentCreateRequest();
      case 'DepartmentAddProject': return $1.DepartmentAddProjectRequest();
      case 'DepartmentRemoveProject': return $1.DepartmentRemoveProjectRequest();
      case 'GetDepartments': return $1.DepartmentGetRequest();
      case 'GetDepartmentByID': return $1.DepartmentGetByIDRequest();
      case 'UpdateDepartment': return $1.DepartmentUpdateRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'CreateDepartment': return this.createDepartment(ctx, request as $1.DepartmentCreateRequest);
      case 'DepartmentAddProject': return this.departmentAddProject(ctx, request as $1.DepartmentAddProjectRequest);
      case 'DepartmentRemoveProject': return this.departmentRemoveProject(ctx, request as $1.DepartmentRemoveProjectRequest);
      case 'GetDepartments': return this.getDepartments(ctx, request as $1.DepartmentGetRequest);
      case 'GetDepartmentByID': return this.getDepartmentByID(ctx, request as $1.DepartmentGetByIDRequest);
      case 'UpdateDepartment': return this.updateDepartment(ctx, request as $1.DepartmentUpdateRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => DepartmentRegistryServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => DepartmentRegistryServiceBase$messageJson;
}

