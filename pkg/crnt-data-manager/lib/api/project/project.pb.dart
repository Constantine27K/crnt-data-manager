///
//  Generated code. Do not modify.
//  source: api/project/project.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class Project extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Project', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'shortName')
    ..aOB(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'isArchived')
    ..p<$fixnum.Int64>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'responsibleTeamIds', $pb.PbFieldType.K6)
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'description')
    ..aInt64(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'department')
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'responsible')
    ..hasRequiredFields = false
  ;

  Project._() : super();
  factory Project({
    $fixnum.Int64? id,
    $core.String? name,
    $core.String? shortName,
    $core.bool? isArchived,
    $core.Iterable<$fixnum.Int64>? responsibleTeamIds,
    $core.String? description,
    $fixnum.Int64? department,
    $core.String? responsible,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (name != null) {
      _result.name = name;
    }
    if (shortName != null) {
      _result.shortName = shortName;
    }
    if (isArchived != null) {
      _result.isArchived = isArchived;
    }
    if (responsibleTeamIds != null) {
      _result.responsibleTeamIds.addAll(responsibleTeamIds);
    }
    if (description != null) {
      _result.description = description;
    }
    if (department != null) {
      _result.department = department;
    }
    if (responsible != null) {
      _result.responsible = responsible;
    }
    return _result;
  }
  factory Project.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Project.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Project clone() => Project()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Project copyWith(void Function(Project) updates) => super.copyWith((message) => updates(message as Project)) as Project; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Project create() => Project._();
  Project createEmptyInstance() => create();
  static $pb.PbList<Project> createRepeated() => $pb.PbList<Project>();
  @$core.pragma('dart2js:noInline')
  static Project getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Project>(create);
  static Project? _defaultInstance;

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
  $core.String get shortName => $_getSZ(2);
  @$pb.TagNumber(3)
  set shortName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasShortName() => $_has(2);
  @$pb.TagNumber(3)
  void clearShortName() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get isArchived => $_getBF(3);
  @$pb.TagNumber(4)
  set isArchived($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasIsArchived() => $_has(3);
  @$pb.TagNumber(4)
  void clearIsArchived() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$fixnum.Int64> get responsibleTeamIds => $_getList(4);

  @$pb.TagNumber(6)
  $core.String get description => $_getSZ(5);
  @$pb.TagNumber(6)
  set description($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasDescription() => $_has(5);
  @$pb.TagNumber(6)
  void clearDescription() => clearField(6);

  @$pb.TagNumber(7)
  $fixnum.Int64 get department => $_getI64(6);
  @$pb.TagNumber(7)
  set department($fixnum.Int64 v) { $_setInt64(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasDepartment() => $_has(6);
  @$pb.TagNumber(7)
  void clearDepartment() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get responsible => $_getSZ(7);
  @$pb.TagNumber(8)
  set responsible($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasResponsible() => $_has(7);
  @$pb.TagNumber(8)
  void clearResponsible() => clearField(8);
}

class ProjectCreateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectCreateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..aOM<Project>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'project', subBuilder: Project.create)
    ..hasRequiredFields = false
  ;

  ProjectCreateRequest._() : super();
  factory ProjectCreateRequest({
    Project? project,
  }) {
    final _result = create();
    if (project != null) {
      _result.project = project;
    }
    return _result;
  }
  factory ProjectCreateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectCreateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectCreateRequest clone() => ProjectCreateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectCreateRequest copyWith(void Function(ProjectCreateRequest) updates) => super.copyWith((message) => updates(message as ProjectCreateRequest)) as ProjectCreateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectCreateRequest create() => ProjectCreateRequest._();
  ProjectCreateRequest createEmptyInstance() => create();
  static $pb.PbList<ProjectCreateRequest> createRepeated() => $pb.PbList<ProjectCreateRequest>();
  @$core.pragma('dart2js:noInline')
  static ProjectCreateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectCreateRequest>(create);
  static ProjectCreateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  Project get project => $_getN(0);
  @$pb.TagNumber(1)
  set project(Project v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasProject() => $_has(0);
  @$pb.TagNumber(1)
  void clearProject() => clearField(1);
  @$pb.TagNumber(1)
  Project ensureProject() => $_ensure(0);
}

class ProjectCreateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectCreateResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  ProjectCreateResponse._() : super();
  factory ProjectCreateResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory ProjectCreateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectCreateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectCreateResponse clone() => ProjectCreateResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectCreateResponse copyWith(void Function(ProjectCreateResponse) updates) => super.copyWith((message) => updates(message as ProjectCreateResponse)) as ProjectCreateResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectCreateResponse create() => ProjectCreateResponse._();
  ProjectCreateResponse createEmptyInstance() => create();
  static $pb.PbList<ProjectCreateResponse> createRepeated() => $pb.PbList<ProjectCreateResponse>();
  @$core.pragma('dart2js:noInline')
  static ProjectCreateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectCreateResponse>(create);
  static ProjectCreateResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class ProjectAddTeamRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectAddTeamRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'teamId')
    ..hasRequiredFields = false
  ;

  ProjectAddTeamRequest._() : super();
  factory ProjectAddTeamRequest({
    $fixnum.Int64? id,
    $fixnum.Int64? teamId,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (teamId != null) {
      _result.teamId = teamId;
    }
    return _result;
  }
  factory ProjectAddTeamRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectAddTeamRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectAddTeamRequest clone() => ProjectAddTeamRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectAddTeamRequest copyWith(void Function(ProjectAddTeamRequest) updates) => super.copyWith((message) => updates(message as ProjectAddTeamRequest)) as ProjectAddTeamRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectAddTeamRequest create() => ProjectAddTeamRequest._();
  ProjectAddTeamRequest createEmptyInstance() => create();
  static $pb.PbList<ProjectAddTeamRequest> createRepeated() => $pb.PbList<ProjectAddTeamRequest>();
  @$core.pragma('dart2js:noInline')
  static ProjectAddTeamRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectAddTeamRequest>(create);
  static ProjectAddTeamRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get teamId => $_getI64(1);
  @$pb.TagNumber(2)
  set teamId($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTeamId() => $_has(1);
  @$pb.TagNumber(2)
  void clearTeamId() => clearField(2);
}

class ProjectAddTeamResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectAddTeamResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'projectId')
    ..hasRequiredFields = false
  ;

  ProjectAddTeamResponse._() : super();
  factory ProjectAddTeamResponse({
    $fixnum.Int64? projectId,
  }) {
    final _result = create();
    if (projectId != null) {
      _result.projectId = projectId;
    }
    return _result;
  }
  factory ProjectAddTeamResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectAddTeamResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectAddTeamResponse clone() => ProjectAddTeamResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectAddTeamResponse copyWith(void Function(ProjectAddTeamResponse) updates) => super.copyWith((message) => updates(message as ProjectAddTeamResponse)) as ProjectAddTeamResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectAddTeamResponse create() => ProjectAddTeamResponse._();
  ProjectAddTeamResponse createEmptyInstance() => create();
  static $pb.PbList<ProjectAddTeamResponse> createRepeated() => $pb.PbList<ProjectAddTeamResponse>();
  @$core.pragma('dart2js:noInline')
  static ProjectAddTeamResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectAddTeamResponse>(create);
  static ProjectAddTeamResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get projectId => $_getI64(0);
  @$pb.TagNumber(1)
  set projectId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasProjectId() => $_has(0);
  @$pb.TagNumber(1)
  void clearProjectId() => clearField(1);
}

class ProjectRemoveTeamRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectRemoveTeamRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'teamId')
    ..hasRequiredFields = false
  ;

  ProjectRemoveTeamRequest._() : super();
  factory ProjectRemoveTeamRequest({
    $fixnum.Int64? id,
    $fixnum.Int64? teamId,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (teamId != null) {
      _result.teamId = teamId;
    }
    return _result;
  }
  factory ProjectRemoveTeamRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectRemoveTeamRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectRemoveTeamRequest clone() => ProjectRemoveTeamRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectRemoveTeamRequest copyWith(void Function(ProjectRemoveTeamRequest) updates) => super.copyWith((message) => updates(message as ProjectRemoveTeamRequest)) as ProjectRemoveTeamRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectRemoveTeamRequest create() => ProjectRemoveTeamRequest._();
  ProjectRemoveTeamRequest createEmptyInstance() => create();
  static $pb.PbList<ProjectRemoveTeamRequest> createRepeated() => $pb.PbList<ProjectRemoveTeamRequest>();
  @$core.pragma('dart2js:noInline')
  static ProjectRemoveTeamRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectRemoveTeamRequest>(create);
  static ProjectRemoveTeamRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get teamId => $_getI64(1);
  @$pb.TagNumber(2)
  set teamId($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTeamId() => $_has(1);
  @$pb.TagNumber(2)
  void clearTeamId() => clearField(2);
}

class ProjectRemoveTeamResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectRemoveTeamResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'projectId')
    ..hasRequiredFields = false
  ;

  ProjectRemoveTeamResponse._() : super();
  factory ProjectRemoveTeamResponse({
    $fixnum.Int64? projectId,
  }) {
    final _result = create();
    if (projectId != null) {
      _result.projectId = projectId;
    }
    return _result;
  }
  factory ProjectRemoveTeamResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectRemoveTeamResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectRemoveTeamResponse clone() => ProjectRemoveTeamResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectRemoveTeamResponse copyWith(void Function(ProjectRemoveTeamResponse) updates) => super.copyWith((message) => updates(message as ProjectRemoveTeamResponse)) as ProjectRemoveTeamResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectRemoveTeamResponse create() => ProjectRemoveTeamResponse._();
  ProjectRemoveTeamResponse createEmptyInstance() => create();
  static $pb.PbList<ProjectRemoveTeamResponse> createRepeated() => $pb.PbList<ProjectRemoveTeamResponse>();
  @$core.pragma('dart2js:noInline')
  static ProjectRemoveTeamResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectRemoveTeamResponse>(create);
  static ProjectRemoveTeamResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get projectId => $_getI64(0);
  @$pb.TagNumber(1)
  set projectId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasProjectId() => $_has(0);
  @$pb.TagNumber(1)
  void clearProjectId() => clearField(1);
}

class ProjectUpdateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectUpdateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOM<Project>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'project', subBuilder: Project.create)
    ..hasRequiredFields = false
  ;

  ProjectUpdateRequest._() : super();
  factory ProjectUpdateRequest({
    $fixnum.Int64? id,
    Project? project,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (project != null) {
      _result.project = project;
    }
    return _result;
  }
  factory ProjectUpdateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectUpdateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectUpdateRequest clone() => ProjectUpdateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectUpdateRequest copyWith(void Function(ProjectUpdateRequest) updates) => super.copyWith((message) => updates(message as ProjectUpdateRequest)) as ProjectUpdateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectUpdateRequest create() => ProjectUpdateRequest._();
  ProjectUpdateRequest createEmptyInstance() => create();
  static $pb.PbList<ProjectUpdateRequest> createRepeated() => $pb.PbList<ProjectUpdateRequest>();
  @$core.pragma('dart2js:noInline')
  static ProjectUpdateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectUpdateRequest>(create);
  static ProjectUpdateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  Project get project => $_getN(1);
  @$pb.TagNumber(2)
  set project(Project v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasProject() => $_has(1);
  @$pb.TagNumber(2)
  void clearProject() => clearField(2);
  @$pb.TagNumber(2)
  Project ensureProject() => $_ensure(1);
}

class ProjectUpdateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectUpdateResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  ProjectUpdateResponse._() : super();
  factory ProjectUpdateResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory ProjectUpdateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectUpdateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectUpdateResponse clone() => ProjectUpdateResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectUpdateResponse copyWith(void Function(ProjectUpdateResponse) updates) => super.copyWith((message) => updates(message as ProjectUpdateResponse)) as ProjectUpdateResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectUpdateResponse create() => ProjectUpdateResponse._();
  ProjectUpdateResponse createEmptyInstance() => create();
  static $pb.PbList<ProjectUpdateResponse> createRepeated() => $pb.PbList<ProjectUpdateResponse>();
  @$core.pragma('dart2js:noInline')
  static ProjectUpdateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectUpdateResponse>(create);
  static ProjectUpdateResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class ProjectGetRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectGetRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..p<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ids', $pb.PbFieldType.K6)
    ..pPS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'names')
    ..pPS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'shortNames')
    ..aOB(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'isArchived')
    ..p<$fixnum.Int64>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'responsibleTeamIds', $pb.PbFieldType.K6)
    ..hasRequiredFields = false
  ;

  ProjectGetRequest._() : super();
  factory ProjectGetRequest({
    $core.Iterable<$fixnum.Int64>? ids,
    $core.Iterable<$core.String>? names,
    $core.Iterable<$core.String>? shortNames,
    $core.bool? isArchived,
    $core.Iterable<$fixnum.Int64>? responsibleTeamIds,
  }) {
    final _result = create();
    if (ids != null) {
      _result.ids.addAll(ids);
    }
    if (names != null) {
      _result.names.addAll(names);
    }
    if (shortNames != null) {
      _result.shortNames.addAll(shortNames);
    }
    if (isArchived != null) {
      _result.isArchived = isArchived;
    }
    if (responsibleTeamIds != null) {
      _result.responsibleTeamIds.addAll(responsibleTeamIds);
    }
    return _result;
  }
  factory ProjectGetRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectGetRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectGetRequest clone() => ProjectGetRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectGetRequest copyWith(void Function(ProjectGetRequest) updates) => super.copyWith((message) => updates(message as ProjectGetRequest)) as ProjectGetRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectGetRequest create() => ProjectGetRequest._();
  ProjectGetRequest createEmptyInstance() => create();
  static $pb.PbList<ProjectGetRequest> createRepeated() => $pb.PbList<ProjectGetRequest>();
  @$core.pragma('dart2js:noInline')
  static ProjectGetRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectGetRequest>(create);
  static ProjectGetRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$fixnum.Int64> get ids => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<$core.String> get names => $_getList(1);

  @$pb.TagNumber(3)
  $core.List<$core.String> get shortNames => $_getList(2);

  @$pb.TagNumber(4)
  $core.bool get isArchived => $_getBF(3);
  @$pb.TagNumber(4)
  set isArchived($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasIsArchived() => $_has(3);
  @$pb.TagNumber(4)
  void clearIsArchived() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$fixnum.Int64> get responsibleTeamIds => $_getList(4);
}

class ProjectGetResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectGetResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..pc<Project>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'projects', $pb.PbFieldType.PM, subBuilder: Project.create)
    ..hasRequiredFields = false
  ;

  ProjectGetResponse._() : super();
  factory ProjectGetResponse({
    $core.Iterable<Project>? projects,
  }) {
    final _result = create();
    if (projects != null) {
      _result.projects.addAll(projects);
    }
    return _result;
  }
  factory ProjectGetResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectGetResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectGetResponse clone() => ProjectGetResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectGetResponse copyWith(void Function(ProjectGetResponse) updates) => super.copyWith((message) => updates(message as ProjectGetResponse)) as ProjectGetResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectGetResponse create() => ProjectGetResponse._();
  ProjectGetResponse createEmptyInstance() => create();
  static $pb.PbList<ProjectGetResponse> createRepeated() => $pb.PbList<ProjectGetResponse>();
  @$core.pragma('dart2js:noInline')
  static ProjectGetResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectGetResponse>(create);
  static ProjectGetResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Project> get projects => $_getList(0);
}

class ProjectGetByIDRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectGetByIDRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  ProjectGetByIDRequest._() : super();
  factory ProjectGetByIDRequest({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory ProjectGetByIDRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectGetByIDRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectGetByIDRequest clone() => ProjectGetByIDRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectGetByIDRequest copyWith(void Function(ProjectGetByIDRequest) updates) => super.copyWith((message) => updates(message as ProjectGetByIDRequest)) as ProjectGetByIDRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectGetByIDRequest create() => ProjectGetByIDRequest._();
  ProjectGetByIDRequest createEmptyInstance() => create();
  static $pb.PbList<ProjectGetByIDRequest> createRepeated() => $pb.PbList<ProjectGetByIDRequest>();
  @$core.pragma('dart2js:noInline')
  static ProjectGetByIDRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectGetByIDRequest>(create);
  static ProjectGetByIDRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class ProjectGetByIDResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ProjectGetByIDResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.project'), createEmptyInstance: create)
    ..aOM<Project>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'project', subBuilder: Project.create)
    ..hasRequiredFields = false
  ;

  ProjectGetByIDResponse._() : super();
  factory ProjectGetByIDResponse({
    Project? project,
  }) {
    final _result = create();
    if (project != null) {
      _result.project = project;
    }
    return _result;
  }
  factory ProjectGetByIDResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectGetByIDResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectGetByIDResponse clone() => ProjectGetByIDResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectGetByIDResponse copyWith(void Function(ProjectGetByIDResponse) updates) => super.copyWith((message) => updates(message as ProjectGetByIDResponse)) as ProjectGetByIDResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ProjectGetByIDResponse create() => ProjectGetByIDResponse._();
  ProjectGetByIDResponse createEmptyInstance() => create();
  static $pb.PbList<ProjectGetByIDResponse> createRepeated() => $pb.PbList<ProjectGetByIDResponse>();
  @$core.pragma('dart2js:noInline')
  static ProjectGetByIDResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectGetByIDResponse>(create);
  static ProjectGetByIDResponse? _defaultInstance;

  @$pb.TagNumber(1)
  Project get project => $_getN(0);
  @$pb.TagNumber(1)
  set project(Project v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasProject() => $_has(0);
  @$pb.TagNumber(1)
  void clearProject() => clearField(1);
  @$pb.TagNumber(1)
  Project ensureProject() => $_ensure(0);
}

class ProjectRegistryApi {
  $pb.RpcClient _client;
  ProjectRegistryApi(this._client);

  $async.Future<ProjectCreateResponse> createProject($pb.ClientContext? ctx, ProjectCreateRequest request) {
    var emptyResponse = ProjectCreateResponse();
    return _client.invoke<ProjectCreateResponse>(ctx, 'ProjectRegistry', 'CreateProject', request, emptyResponse);
  }
  $async.Future<ProjectAddTeamResponse> addResponsibleTeam($pb.ClientContext? ctx, ProjectAddTeamRequest request) {
    var emptyResponse = ProjectAddTeamResponse();
    return _client.invoke<ProjectAddTeamResponse>(ctx, 'ProjectRegistry', 'AddResponsibleTeam', request, emptyResponse);
  }
  $async.Future<ProjectRemoveTeamResponse> removeResponsibleTeam($pb.ClientContext? ctx, ProjectRemoveTeamRequest request) {
    var emptyResponse = ProjectRemoveTeamResponse();
    return _client.invoke<ProjectRemoveTeamResponse>(ctx, 'ProjectRegistry', 'RemoveResponsibleTeam', request, emptyResponse);
  }
  $async.Future<ProjectUpdateResponse> updateProject($pb.ClientContext? ctx, ProjectUpdateRequest request) {
    var emptyResponse = ProjectUpdateResponse();
    return _client.invoke<ProjectUpdateResponse>(ctx, 'ProjectRegistry', 'UpdateProject', request, emptyResponse);
  }
  $async.Future<ProjectGetResponse> getProjects($pb.ClientContext? ctx, ProjectGetRequest request) {
    var emptyResponse = ProjectGetResponse();
    return _client.invoke<ProjectGetResponse>(ctx, 'ProjectRegistry', 'GetProjects', request, emptyResponse);
  }
  $async.Future<ProjectGetByIDResponse> getProjectByID($pb.ClientContext? ctx, ProjectGetByIDRequest request) {
    var emptyResponse = ProjectGetByIDResponse();
    return _client.invoke<ProjectGetByIDResponse>(ctx, 'ProjectRegistry', 'GetProjectByID', request, emptyResponse);
  }
}

