///
//  Generated code. Do not modify.
//  source: api/department/department.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class Department extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Department', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..p<$fixnum.Int64>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'projects', $pb.PbFieldType.K6)
    ..pPS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'members')
    ..hasRequiredFields = false
  ;

  Department._() : super();
  factory Department({
    $fixnum.Int64? id,
    $core.String? name,
    $core.Iterable<$fixnum.Int64>? projects,
    $core.Iterable<$core.String>? members,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (name != null) {
      _result.name = name;
    }
    if (projects != null) {
      _result.projects.addAll(projects);
    }
    if (members != null) {
      _result.members.addAll(members);
    }
    return _result;
  }
  factory Department.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Department.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Department clone() => Department()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Department copyWith(void Function(Department) updates) => super.copyWith((message) => updates(message as Department)) as Department; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Department create() => Department._();
  Department createEmptyInstance() => create();
  static $pb.PbList<Department> createRepeated() => $pb.PbList<Department>();
  @$core.pragma('dart2js:noInline')
  static Department getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Department>(create);
  static Department? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<$fixnum.Int64> get projects => $_getList(2);

  @$pb.TagNumber(4)
  $core.List<$core.String> get members => $_getList(3);
}

class DepartmentCreateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentCreateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..aOM<Department>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'department', subBuilder: Department.create)
    ..hasRequiredFields = false
  ;

  DepartmentCreateRequest._() : super();
  factory DepartmentCreateRequest({
    Department? department,
  }) {
    final _result = create();
    if (department != null) {
      _result.department = department;
    }
    return _result;
  }
  factory DepartmentCreateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentCreateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentCreateRequest clone() => DepartmentCreateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentCreateRequest copyWith(void Function(DepartmentCreateRequest) updates) => super.copyWith((message) => updates(message as DepartmentCreateRequest)) as DepartmentCreateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentCreateRequest create() => DepartmentCreateRequest._();
  DepartmentCreateRequest createEmptyInstance() => create();
  static $pb.PbList<DepartmentCreateRequest> createRepeated() => $pb.PbList<DepartmentCreateRequest>();
  @$core.pragma('dart2js:noInline')
  static DepartmentCreateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentCreateRequest>(create);
  static DepartmentCreateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  Department get department => $_getN(0);
  @$pb.TagNumber(1)
  set department(Department v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasDepartment() => $_has(0);
  @$pb.TagNumber(1)
  void clearDepartment() => clearField(1);
  @$pb.TagNumber(1)
  Department ensureDepartment() => $_ensure(0);
}

class DepartmentCreateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentCreateResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  DepartmentCreateResponse._() : super();
  factory DepartmentCreateResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory DepartmentCreateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentCreateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentCreateResponse clone() => DepartmentCreateResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentCreateResponse copyWith(void Function(DepartmentCreateResponse) updates) => super.copyWith((message) => updates(message as DepartmentCreateResponse)) as DepartmentCreateResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentCreateResponse create() => DepartmentCreateResponse._();
  DepartmentCreateResponse createEmptyInstance() => create();
  static $pb.PbList<DepartmentCreateResponse> createRepeated() => $pb.PbList<DepartmentCreateResponse>();
  @$core.pragma('dart2js:noInline')
  static DepartmentCreateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentCreateResponse>(create);
  static DepartmentCreateResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class DepartmentAddProjectRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentAddProjectRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'projectId')
    ..hasRequiredFields = false
  ;

  DepartmentAddProjectRequest._() : super();
  factory DepartmentAddProjectRequest({
    $fixnum.Int64? id,
    $fixnum.Int64? projectId,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (projectId != null) {
      _result.projectId = projectId;
    }
    return _result;
  }
  factory DepartmentAddProjectRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentAddProjectRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentAddProjectRequest clone() => DepartmentAddProjectRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentAddProjectRequest copyWith(void Function(DepartmentAddProjectRequest) updates) => super.copyWith((message) => updates(message as DepartmentAddProjectRequest)) as DepartmentAddProjectRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentAddProjectRequest create() => DepartmentAddProjectRequest._();
  DepartmentAddProjectRequest createEmptyInstance() => create();
  static $pb.PbList<DepartmentAddProjectRequest> createRepeated() => $pb.PbList<DepartmentAddProjectRequest>();
  @$core.pragma('dart2js:noInline')
  static DepartmentAddProjectRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentAddProjectRequest>(create);
  static DepartmentAddProjectRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get projectId => $_getI64(1);
  @$pb.TagNumber(2)
  set projectId($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasProjectId() => $_has(1);
  @$pb.TagNumber(2)
  void clearProjectId() => clearField(2);
}

class DepartmentAddProjectResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentAddProjectResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  DepartmentAddProjectResponse._() : super();
  factory DepartmentAddProjectResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory DepartmentAddProjectResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentAddProjectResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentAddProjectResponse clone() => DepartmentAddProjectResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentAddProjectResponse copyWith(void Function(DepartmentAddProjectResponse) updates) => super.copyWith((message) => updates(message as DepartmentAddProjectResponse)) as DepartmentAddProjectResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentAddProjectResponse create() => DepartmentAddProjectResponse._();
  DepartmentAddProjectResponse createEmptyInstance() => create();
  static $pb.PbList<DepartmentAddProjectResponse> createRepeated() => $pb.PbList<DepartmentAddProjectResponse>();
  @$core.pragma('dart2js:noInline')
  static DepartmentAddProjectResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentAddProjectResponse>(create);
  static DepartmentAddProjectResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class DepartmentRemoveProjectRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentRemoveProjectRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'projectId')
    ..hasRequiredFields = false
  ;

  DepartmentRemoveProjectRequest._() : super();
  factory DepartmentRemoveProjectRequest({
    $fixnum.Int64? id,
    $fixnum.Int64? projectId,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (projectId != null) {
      _result.projectId = projectId;
    }
    return _result;
  }
  factory DepartmentRemoveProjectRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentRemoveProjectRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentRemoveProjectRequest clone() => DepartmentRemoveProjectRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentRemoveProjectRequest copyWith(void Function(DepartmentRemoveProjectRequest) updates) => super.copyWith((message) => updates(message as DepartmentRemoveProjectRequest)) as DepartmentRemoveProjectRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentRemoveProjectRequest create() => DepartmentRemoveProjectRequest._();
  DepartmentRemoveProjectRequest createEmptyInstance() => create();
  static $pb.PbList<DepartmentRemoveProjectRequest> createRepeated() => $pb.PbList<DepartmentRemoveProjectRequest>();
  @$core.pragma('dart2js:noInline')
  static DepartmentRemoveProjectRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentRemoveProjectRequest>(create);
  static DepartmentRemoveProjectRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get projectId => $_getI64(1);
  @$pb.TagNumber(2)
  set projectId($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasProjectId() => $_has(1);
  @$pb.TagNumber(2)
  void clearProjectId() => clearField(2);
}

class DepartmentRemoveProjectResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentRemoveProjectResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  DepartmentRemoveProjectResponse._() : super();
  factory DepartmentRemoveProjectResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory DepartmentRemoveProjectResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentRemoveProjectResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentRemoveProjectResponse clone() => DepartmentRemoveProjectResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentRemoveProjectResponse copyWith(void Function(DepartmentRemoveProjectResponse) updates) => super.copyWith((message) => updates(message as DepartmentRemoveProjectResponse)) as DepartmentRemoveProjectResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentRemoveProjectResponse create() => DepartmentRemoveProjectResponse._();
  DepartmentRemoveProjectResponse createEmptyInstance() => create();
  static $pb.PbList<DepartmentRemoveProjectResponse> createRepeated() => $pb.PbList<DepartmentRemoveProjectResponse>();
  @$core.pragma('dart2js:noInline')
  static DepartmentRemoveProjectResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentRemoveProjectResponse>(create);
  static DepartmentRemoveProjectResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class DepartmentGetRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentGetRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..p<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ids', $pb.PbFieldType.K6)
    ..pPS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'names')
    ..hasRequiredFields = false
  ;

  DepartmentGetRequest._() : super();
  factory DepartmentGetRequest({
    $core.Iterable<$fixnum.Int64>? ids,
    $core.Iterable<$core.String>? names,
  }) {
    final _result = create();
    if (ids != null) {
      _result.ids.addAll(ids);
    }
    if (names != null) {
      _result.names.addAll(names);
    }
    return _result;
  }
  factory DepartmentGetRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentGetRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentGetRequest clone() => DepartmentGetRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentGetRequest copyWith(void Function(DepartmentGetRequest) updates) => super.copyWith((message) => updates(message as DepartmentGetRequest)) as DepartmentGetRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentGetRequest create() => DepartmentGetRequest._();
  DepartmentGetRequest createEmptyInstance() => create();
  static $pb.PbList<DepartmentGetRequest> createRepeated() => $pb.PbList<DepartmentGetRequest>();
  @$core.pragma('dart2js:noInline')
  static DepartmentGetRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentGetRequest>(create);
  static DepartmentGetRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$fixnum.Int64> get ids => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<$core.String> get names => $_getList(1);
}

class DepartmentGetResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentGetResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..pc<Department>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'departments', $pb.PbFieldType.PM, subBuilder: Department.create)
    ..hasRequiredFields = false
  ;

  DepartmentGetResponse._() : super();
  factory DepartmentGetResponse({
    $core.Iterable<Department>? departments,
  }) {
    final _result = create();
    if (departments != null) {
      _result.departments.addAll(departments);
    }
    return _result;
  }
  factory DepartmentGetResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentGetResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentGetResponse clone() => DepartmentGetResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentGetResponse copyWith(void Function(DepartmentGetResponse) updates) => super.copyWith((message) => updates(message as DepartmentGetResponse)) as DepartmentGetResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentGetResponse create() => DepartmentGetResponse._();
  DepartmentGetResponse createEmptyInstance() => create();
  static $pb.PbList<DepartmentGetResponse> createRepeated() => $pb.PbList<DepartmentGetResponse>();
  @$core.pragma('dart2js:noInline')
  static DepartmentGetResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentGetResponse>(create);
  static DepartmentGetResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Department> get departments => $_getList(0);
}

class DepartmentGetByIDRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentGetByIDRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  DepartmentGetByIDRequest._() : super();
  factory DepartmentGetByIDRequest({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory DepartmentGetByIDRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentGetByIDRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentGetByIDRequest clone() => DepartmentGetByIDRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentGetByIDRequest copyWith(void Function(DepartmentGetByIDRequest) updates) => super.copyWith((message) => updates(message as DepartmentGetByIDRequest)) as DepartmentGetByIDRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentGetByIDRequest create() => DepartmentGetByIDRequest._();
  DepartmentGetByIDRequest createEmptyInstance() => create();
  static $pb.PbList<DepartmentGetByIDRequest> createRepeated() => $pb.PbList<DepartmentGetByIDRequest>();
  @$core.pragma('dart2js:noInline')
  static DepartmentGetByIDRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentGetByIDRequest>(create);
  static DepartmentGetByIDRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class DepartmentGetByIDResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentGetByIDResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..aOM<Department>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'department', subBuilder: Department.create)
    ..hasRequiredFields = false
  ;

  DepartmentGetByIDResponse._() : super();
  factory DepartmentGetByIDResponse({
    Department? department,
  }) {
    final _result = create();
    if (department != null) {
      _result.department = department;
    }
    return _result;
  }
  factory DepartmentGetByIDResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentGetByIDResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentGetByIDResponse clone() => DepartmentGetByIDResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentGetByIDResponse copyWith(void Function(DepartmentGetByIDResponse) updates) => super.copyWith((message) => updates(message as DepartmentGetByIDResponse)) as DepartmentGetByIDResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentGetByIDResponse create() => DepartmentGetByIDResponse._();
  DepartmentGetByIDResponse createEmptyInstance() => create();
  static $pb.PbList<DepartmentGetByIDResponse> createRepeated() => $pb.PbList<DepartmentGetByIDResponse>();
  @$core.pragma('dart2js:noInline')
  static DepartmentGetByIDResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentGetByIDResponse>(create);
  static DepartmentGetByIDResponse? _defaultInstance;

  @$pb.TagNumber(1)
  Department get department => $_getN(0);
  @$pb.TagNumber(1)
  set department(Department v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasDepartment() => $_has(0);
  @$pb.TagNumber(1)
  void clearDepartment() => clearField(1);
  @$pb.TagNumber(1)
  Department ensureDepartment() => $_ensure(0);
}

class DepartmentUpdateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentUpdateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOM<Department>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'department', subBuilder: Department.create)
    ..hasRequiredFields = false
  ;

  DepartmentUpdateRequest._() : super();
  factory DepartmentUpdateRequest({
    $fixnum.Int64? id,
    Department? department,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (department != null) {
      _result.department = department;
    }
    return _result;
  }
  factory DepartmentUpdateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentUpdateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentUpdateRequest clone() => DepartmentUpdateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentUpdateRequest copyWith(void Function(DepartmentUpdateRequest) updates) => super.copyWith((message) => updates(message as DepartmentUpdateRequest)) as DepartmentUpdateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentUpdateRequest create() => DepartmentUpdateRequest._();
  DepartmentUpdateRequest createEmptyInstance() => create();
  static $pb.PbList<DepartmentUpdateRequest> createRepeated() => $pb.PbList<DepartmentUpdateRequest>();
  @$core.pragma('dart2js:noInline')
  static DepartmentUpdateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentUpdateRequest>(create);
  static DepartmentUpdateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  Department get department => $_getN(1);
  @$pb.TagNumber(2)
  set department(Department v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasDepartment() => $_has(1);
  @$pb.TagNumber(2)
  void clearDepartment() => clearField(2);
  @$pb.TagNumber(2)
  Department ensureDepartment() => $_ensure(1);
}

class DepartmentUpdateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DepartmentUpdateResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.department'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  DepartmentUpdateResponse._() : super();
  factory DepartmentUpdateResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory DepartmentUpdateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DepartmentUpdateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DepartmentUpdateResponse clone() => DepartmentUpdateResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DepartmentUpdateResponse copyWith(void Function(DepartmentUpdateResponse) updates) => super.copyWith((message) => updates(message as DepartmentUpdateResponse)) as DepartmentUpdateResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DepartmentUpdateResponse create() => DepartmentUpdateResponse._();
  DepartmentUpdateResponse createEmptyInstance() => create();
  static $pb.PbList<DepartmentUpdateResponse> createRepeated() => $pb.PbList<DepartmentUpdateResponse>();
  @$core.pragma('dart2js:noInline')
  static DepartmentUpdateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DepartmentUpdateResponse>(create);
  static DepartmentUpdateResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class DepartmentRegistryApi {
  $pb.RpcClient _client;
  DepartmentRegistryApi(this._client);

  $async.Future<DepartmentCreateResponse> createDepartment($pb.ClientContext? ctx, DepartmentCreateRequest request) {
    var emptyResponse = DepartmentCreateResponse();
    return _client.invoke<DepartmentCreateResponse>(ctx, 'DepartmentRegistry', 'CreateDepartment', request, emptyResponse);
  }
  $async.Future<DepartmentAddProjectResponse> departmentAddProject($pb.ClientContext? ctx, DepartmentAddProjectRequest request) {
    var emptyResponse = DepartmentAddProjectResponse();
    return _client.invoke<DepartmentAddProjectResponse>(ctx, 'DepartmentRegistry', 'DepartmentAddProject', request, emptyResponse);
  }
  $async.Future<DepartmentRemoveProjectResponse> departmentRemoveProject($pb.ClientContext? ctx, DepartmentRemoveProjectRequest request) {
    var emptyResponse = DepartmentRemoveProjectResponse();
    return _client.invoke<DepartmentRemoveProjectResponse>(ctx, 'DepartmentRegistry', 'DepartmentRemoveProject', request, emptyResponse);
  }
  $async.Future<DepartmentGetResponse> getDepartments($pb.ClientContext? ctx, DepartmentGetRequest request) {
    var emptyResponse = DepartmentGetResponse();
    return _client.invoke<DepartmentGetResponse>(ctx, 'DepartmentRegistry', 'GetDepartments', request, emptyResponse);
  }
  $async.Future<DepartmentGetByIDResponse> getDepartmentByID($pb.ClientContext? ctx, DepartmentGetByIDRequest request) {
    var emptyResponse = DepartmentGetByIDResponse();
    return _client.invoke<DepartmentGetByIDResponse>(ctx, 'DepartmentRegistry', 'GetDepartmentByID', request, emptyResponse);
  }
  $async.Future<DepartmentUpdateResponse> updateDepartment($pb.ClientContext? ctx, DepartmentUpdateRequest request) {
    var emptyResponse = DepartmentUpdateResponse();
    return _client.invoke<DepartmentUpdateResponse>(ctx, 'DepartmentRegistry', 'UpdateDepartment', request, emptyResponse);
  }
}

