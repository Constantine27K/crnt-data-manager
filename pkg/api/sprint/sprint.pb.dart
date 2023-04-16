///
//  Generated code. Do not modify.
//  source: api/sprint/sprint.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/timestamp.pb.dart' as $0;

import '../state/status/status.pbenum.dart' as $4;

class Sprint extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Sprint', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aInt64(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'project')
    ..aOM<$0.Timestamp>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'startedAt', subBuilder: $0.Timestamp.create)
    ..aOM<$0.Timestamp>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'finishedAt', subBuilder: $0.Timestamp.create)
    ..e<$4.Sprint>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: $4.Sprint.STATUS_SPRINT_UNKNOWN, valueOf: $4.Sprint.valueOf, enumValues: $4.Sprint.values)
    ..aInt64(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'storiesBacklog')
    ..aInt64(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'storiesInProgress')
    ..aInt64(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'storiesDone')
    ..p<$fixnum.Int64>(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'issues', $pb.PbFieldType.K6)
    ..hasRequiredFields = false
  ;

  Sprint._() : super();
  factory Sprint({
    $fixnum.Int64? id,
    $core.String? name,
    $fixnum.Int64? project,
    $0.Timestamp? startedAt,
    $0.Timestamp? finishedAt,
    $4.Sprint? status,
    $fixnum.Int64? storiesBacklog,
    $fixnum.Int64? storiesInProgress,
    $fixnum.Int64? storiesDone,
    $core.Iterable<$fixnum.Int64>? issues,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (name != null) {
      _result.name = name;
    }
    if (project != null) {
      _result.project = project;
    }
    if (startedAt != null) {
      _result.startedAt = startedAt;
    }
    if (finishedAt != null) {
      _result.finishedAt = finishedAt;
    }
    if (status != null) {
      _result.status = status;
    }
    if (storiesBacklog != null) {
      _result.storiesBacklog = storiesBacklog;
    }
    if (storiesInProgress != null) {
      _result.storiesInProgress = storiesInProgress;
    }
    if (storiesDone != null) {
      _result.storiesDone = storiesDone;
    }
    if (issues != null) {
      _result.issues.addAll(issues);
    }
    return _result;
  }
  factory Sprint.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Sprint.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Sprint clone() => Sprint()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Sprint copyWith(void Function(Sprint) updates) => super.copyWith((message) => updates(message as Sprint)) as Sprint; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Sprint create() => Sprint._();
  Sprint createEmptyInstance() => create();
  static $pb.PbList<Sprint> createRepeated() => $pb.PbList<Sprint>();
  @$core.pragma('dart2js:noInline')
  static Sprint getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Sprint>(create);
  static Sprint? _defaultInstance;

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
  $fixnum.Int64 get project => $_getI64(2);
  @$pb.TagNumber(3)
  set project($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasProject() => $_has(2);
  @$pb.TagNumber(3)
  void clearProject() => clearField(3);

  @$pb.TagNumber(4)
  $0.Timestamp get startedAt => $_getN(3);
  @$pb.TagNumber(4)
  set startedAt($0.Timestamp v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasStartedAt() => $_has(3);
  @$pb.TagNumber(4)
  void clearStartedAt() => clearField(4);
  @$pb.TagNumber(4)
  $0.Timestamp ensureStartedAt() => $_ensure(3);

  @$pb.TagNumber(5)
  $0.Timestamp get finishedAt => $_getN(4);
  @$pb.TagNumber(5)
  set finishedAt($0.Timestamp v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasFinishedAt() => $_has(4);
  @$pb.TagNumber(5)
  void clearFinishedAt() => clearField(5);
  @$pb.TagNumber(5)
  $0.Timestamp ensureFinishedAt() => $_ensure(4);

  @$pb.TagNumber(6)
  $4.Sprint get status => $_getN(5);
  @$pb.TagNumber(6)
  set status($4.Sprint v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasStatus() => $_has(5);
  @$pb.TagNumber(6)
  void clearStatus() => clearField(6);

  @$pb.TagNumber(7)
  $fixnum.Int64 get storiesBacklog => $_getI64(6);
  @$pb.TagNumber(7)
  set storiesBacklog($fixnum.Int64 v) { $_setInt64(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasStoriesBacklog() => $_has(6);
  @$pb.TagNumber(7)
  void clearStoriesBacklog() => clearField(7);

  @$pb.TagNumber(8)
  $fixnum.Int64 get storiesInProgress => $_getI64(7);
  @$pb.TagNumber(8)
  set storiesInProgress($fixnum.Int64 v) { $_setInt64(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasStoriesInProgress() => $_has(7);
  @$pb.TagNumber(8)
  void clearStoriesInProgress() => clearField(8);

  @$pb.TagNumber(9)
  $fixnum.Int64 get storiesDone => $_getI64(8);
  @$pb.TagNumber(9)
  set storiesDone($fixnum.Int64 v) { $_setInt64(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasStoriesDone() => $_has(8);
  @$pb.TagNumber(9)
  void clearStoriesDone() => clearField(9);

  @$pb.TagNumber(10)
  $core.List<$fixnum.Int64> get issues => $_getList(9);
}

class SprintCreateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SprintCreateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aOM<Sprint>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sprint', subBuilder: Sprint.create)
    ..hasRequiredFields = false
  ;

  SprintCreateRequest._() : super();
  factory SprintCreateRequest({
    Sprint? sprint,
  }) {
    final _result = create();
    if (sprint != null) {
      _result.sprint = sprint;
    }
    return _result;
  }
  factory SprintCreateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SprintCreateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SprintCreateRequest clone() => SprintCreateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SprintCreateRequest copyWith(void Function(SprintCreateRequest) updates) => super.copyWith((message) => updates(message as SprintCreateRequest)) as SprintCreateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SprintCreateRequest create() => SprintCreateRequest._();
  SprintCreateRequest createEmptyInstance() => create();
  static $pb.PbList<SprintCreateRequest> createRepeated() => $pb.PbList<SprintCreateRequest>();
  @$core.pragma('dart2js:noInline')
  static SprintCreateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SprintCreateRequest>(create);
  static SprintCreateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  Sprint get sprint => $_getN(0);
  @$pb.TagNumber(1)
  set sprint(Sprint v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasSprint() => $_has(0);
  @$pb.TagNumber(1)
  void clearSprint() => clearField(1);
  @$pb.TagNumber(1)
  Sprint ensureSprint() => $_ensure(0);
}

class SprintCreateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SprintCreateResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  SprintCreateResponse._() : super();
  factory SprintCreateResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory SprintCreateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SprintCreateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SprintCreateResponse clone() => SprintCreateResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SprintCreateResponse copyWith(void Function(SprintCreateResponse) updates) => super.copyWith((message) => updates(message as SprintCreateResponse)) as SprintCreateResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SprintCreateResponse create() => SprintCreateResponse._();
  SprintCreateResponse createEmptyInstance() => create();
  static $pb.PbList<SprintCreateResponse> createRepeated() => $pb.PbList<SprintCreateResponse>();
  @$core.pragma('dart2js:noInline')
  static SprintCreateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SprintCreateResponse>(create);
  static SprintCreateResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class AddIssueRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AddIssueRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'issueId')
    ..hasRequiredFields = false
  ;

  AddIssueRequest._() : super();
  factory AddIssueRequest({
    $fixnum.Int64? id,
    $fixnum.Int64? issueId,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (issueId != null) {
      _result.issueId = issueId;
    }
    return _result;
  }
  factory AddIssueRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AddIssueRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AddIssueRequest clone() => AddIssueRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AddIssueRequest copyWith(void Function(AddIssueRequest) updates) => super.copyWith((message) => updates(message as AddIssueRequest)) as AddIssueRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AddIssueRequest create() => AddIssueRequest._();
  AddIssueRequest createEmptyInstance() => create();
  static $pb.PbList<AddIssueRequest> createRepeated() => $pb.PbList<AddIssueRequest>();
  @$core.pragma('dart2js:noInline')
  static AddIssueRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AddIssueRequest>(create);
  static AddIssueRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get issueId => $_getI64(1);
  @$pb.TagNumber(2)
  set issueId($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasIssueId() => $_has(1);
  @$pb.TagNumber(2)
  void clearIssueId() => clearField(2);
}

class AddIssueResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AddIssueResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sprintId')
    ..hasRequiredFields = false
  ;

  AddIssueResponse._() : super();
  factory AddIssueResponse({
    $fixnum.Int64? sprintId,
  }) {
    final _result = create();
    if (sprintId != null) {
      _result.sprintId = sprintId;
    }
    return _result;
  }
  factory AddIssueResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AddIssueResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AddIssueResponse clone() => AddIssueResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AddIssueResponse copyWith(void Function(AddIssueResponse) updates) => super.copyWith((message) => updates(message as AddIssueResponse)) as AddIssueResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AddIssueResponse create() => AddIssueResponse._();
  AddIssueResponse createEmptyInstance() => create();
  static $pb.PbList<AddIssueResponse> createRepeated() => $pb.PbList<AddIssueResponse>();
  @$core.pragma('dart2js:noInline')
  static AddIssueResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AddIssueResponse>(create);
  static AddIssueResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get sprintId => $_getI64(0);
  @$pb.TagNumber(1)
  set sprintId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSprintId() => $_has(0);
  @$pb.TagNumber(1)
  void clearSprintId() => clearField(1);
}

class RemoveIssueRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RemoveIssueRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'issueId')
    ..hasRequiredFields = false
  ;

  RemoveIssueRequest._() : super();
  factory RemoveIssueRequest({
    $fixnum.Int64? id,
    $fixnum.Int64? issueId,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (issueId != null) {
      _result.issueId = issueId;
    }
    return _result;
  }
  factory RemoveIssueRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RemoveIssueRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RemoveIssueRequest clone() => RemoveIssueRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RemoveIssueRequest copyWith(void Function(RemoveIssueRequest) updates) => super.copyWith((message) => updates(message as RemoveIssueRequest)) as RemoveIssueRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RemoveIssueRequest create() => RemoveIssueRequest._();
  RemoveIssueRequest createEmptyInstance() => create();
  static $pb.PbList<RemoveIssueRequest> createRepeated() => $pb.PbList<RemoveIssueRequest>();
  @$core.pragma('dart2js:noInline')
  static RemoveIssueRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RemoveIssueRequest>(create);
  static RemoveIssueRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get issueId => $_getI64(1);
  @$pb.TagNumber(2)
  set issueId($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasIssueId() => $_has(1);
  @$pb.TagNumber(2)
  void clearIssueId() => clearField(2);
}

class RemoveIssueResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RemoveIssueResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sprintId')
    ..hasRequiredFields = false
  ;

  RemoveIssueResponse._() : super();
  factory RemoveIssueResponse({
    $fixnum.Int64? sprintId,
  }) {
    final _result = create();
    if (sprintId != null) {
      _result.sprintId = sprintId;
    }
    return _result;
  }
  factory RemoveIssueResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RemoveIssueResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RemoveIssueResponse clone() => RemoveIssueResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RemoveIssueResponse copyWith(void Function(RemoveIssueResponse) updates) => super.copyWith((message) => updates(message as RemoveIssueResponse)) as RemoveIssueResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RemoveIssueResponse create() => RemoveIssueResponse._();
  RemoveIssueResponse createEmptyInstance() => create();
  static $pb.PbList<RemoveIssueResponse> createRepeated() => $pb.PbList<RemoveIssueResponse>();
  @$core.pragma('dart2js:noInline')
  static RemoveIssueResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RemoveIssueResponse>(create);
  static RemoveIssueResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get sprintId => $_getI64(0);
  @$pb.TagNumber(1)
  set sprintId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSprintId() => $_has(0);
  @$pb.TagNumber(1)
  void clearSprintId() => clearField(1);
}

class SprintUpdateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SprintUpdateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOM<Sprint>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sprint', subBuilder: Sprint.create)
    ..hasRequiredFields = false
  ;

  SprintUpdateRequest._() : super();
  factory SprintUpdateRequest({
    $fixnum.Int64? id,
    Sprint? sprint,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (sprint != null) {
      _result.sprint = sprint;
    }
    return _result;
  }
  factory SprintUpdateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SprintUpdateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SprintUpdateRequest clone() => SprintUpdateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SprintUpdateRequest copyWith(void Function(SprintUpdateRequest) updates) => super.copyWith((message) => updates(message as SprintUpdateRequest)) as SprintUpdateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SprintUpdateRequest create() => SprintUpdateRequest._();
  SprintUpdateRequest createEmptyInstance() => create();
  static $pb.PbList<SprintUpdateRequest> createRepeated() => $pb.PbList<SprintUpdateRequest>();
  @$core.pragma('dart2js:noInline')
  static SprintUpdateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SprintUpdateRequest>(create);
  static SprintUpdateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  Sprint get sprint => $_getN(1);
  @$pb.TagNumber(2)
  set sprint(Sprint v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasSprint() => $_has(1);
  @$pb.TagNumber(2)
  void clearSprint() => clearField(2);
  @$pb.TagNumber(2)
  Sprint ensureSprint() => $_ensure(1);
}

class SprintUpdateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SprintUpdateResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  SprintUpdateResponse._() : super();
  factory SprintUpdateResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory SprintUpdateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SprintUpdateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SprintUpdateResponse clone() => SprintUpdateResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SprintUpdateResponse copyWith(void Function(SprintUpdateResponse) updates) => super.copyWith((message) => updates(message as SprintUpdateResponse)) as SprintUpdateResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SprintUpdateResponse create() => SprintUpdateResponse._();
  SprintUpdateResponse createEmptyInstance() => create();
  static $pb.PbList<SprintUpdateResponse> createRepeated() => $pb.PbList<SprintUpdateResponse>();
  @$core.pragma('dart2js:noInline')
  static SprintUpdateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SprintUpdateResponse>(create);
  static SprintUpdateResponse? _defaultInstance;

  @$pb.TagNumber(2)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(2)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(2)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(2)
  void clearId() => clearField(2);
}

class SprintChangeStatusRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SprintChangeStatusRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..e<$4.Sprint>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: $4.Sprint.STATUS_SPRINT_UNKNOWN, valueOf: $4.Sprint.valueOf, enumValues: $4.Sprint.values)
    ..hasRequiredFields = false
  ;

  SprintChangeStatusRequest._() : super();
  factory SprintChangeStatusRequest({
    $fixnum.Int64? id,
    $4.Sprint? status,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (status != null) {
      _result.status = status;
    }
    return _result;
  }
  factory SprintChangeStatusRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SprintChangeStatusRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SprintChangeStatusRequest clone() => SprintChangeStatusRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SprintChangeStatusRequest copyWith(void Function(SprintChangeStatusRequest) updates) => super.copyWith((message) => updates(message as SprintChangeStatusRequest)) as SprintChangeStatusRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SprintChangeStatusRequest create() => SprintChangeStatusRequest._();
  SprintChangeStatusRequest createEmptyInstance() => create();
  static $pb.PbList<SprintChangeStatusRequest> createRepeated() => $pb.PbList<SprintChangeStatusRequest>();
  @$core.pragma('dart2js:noInline')
  static SprintChangeStatusRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SprintChangeStatusRequest>(create);
  static SprintChangeStatusRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $4.Sprint get status => $_getN(1);
  @$pb.TagNumber(2)
  set status($4.Sprint v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasStatus() => $_has(1);
  @$pb.TagNumber(2)
  void clearStatus() => clearField(2);
}

class SprintChangeStatusResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SprintChangeStatusResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  SprintChangeStatusResponse._() : super();
  factory SprintChangeStatusResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory SprintChangeStatusResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SprintChangeStatusResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SprintChangeStatusResponse clone() => SprintChangeStatusResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SprintChangeStatusResponse copyWith(void Function(SprintChangeStatusResponse) updates) => super.copyWith((message) => updates(message as SprintChangeStatusResponse)) as SprintChangeStatusResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SprintChangeStatusResponse create() => SprintChangeStatusResponse._();
  SprintChangeStatusResponse createEmptyInstance() => create();
  static $pb.PbList<SprintChangeStatusResponse> createRepeated() => $pb.PbList<SprintChangeStatusResponse>();
  @$core.pragma('dart2js:noInline')
  static SprintChangeStatusResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SprintChangeStatusResponse>(create);
  static SprintChangeStatusResponse? _defaultInstance;

  @$pb.TagNumber(2)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(2)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(2)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(2)
  void clearId() => clearField(2);
}

class SprintGetRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SprintGetRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..p<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ids', $pb.PbFieldType.K6)
    ..pPS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'names')
    ..p<$fixnum.Int64>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'projects', $pb.PbFieldType.K6)
    ..aOM<$0.Timestamp>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'startedAtAfter', subBuilder: $0.Timestamp.create)
    ..aOM<$0.Timestamp>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'finishedAtBefore', subBuilder: $0.Timestamp.create)
    ..e<$4.Sprint>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: $4.Sprint.STATUS_SPRINT_UNKNOWN, valueOf: $4.Sprint.valueOf, enumValues: $4.Sprint.values)
    ..hasRequiredFields = false
  ;

  SprintGetRequest._() : super();
  factory SprintGetRequest({
    $core.Iterable<$fixnum.Int64>? ids,
    $core.Iterable<$core.String>? names,
    $core.Iterable<$fixnum.Int64>? projects,
    $0.Timestamp? startedAtAfter,
    $0.Timestamp? finishedAtBefore,
    $4.Sprint? status,
  }) {
    final _result = create();
    if (ids != null) {
      _result.ids.addAll(ids);
    }
    if (names != null) {
      _result.names.addAll(names);
    }
    if (projects != null) {
      _result.projects.addAll(projects);
    }
    if (startedAtAfter != null) {
      _result.startedAtAfter = startedAtAfter;
    }
    if (finishedAtBefore != null) {
      _result.finishedAtBefore = finishedAtBefore;
    }
    if (status != null) {
      _result.status = status;
    }
    return _result;
  }
  factory SprintGetRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SprintGetRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SprintGetRequest clone() => SprintGetRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SprintGetRequest copyWith(void Function(SprintGetRequest) updates) => super.copyWith((message) => updates(message as SprintGetRequest)) as SprintGetRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SprintGetRequest create() => SprintGetRequest._();
  SprintGetRequest createEmptyInstance() => create();
  static $pb.PbList<SprintGetRequest> createRepeated() => $pb.PbList<SprintGetRequest>();
  @$core.pragma('dart2js:noInline')
  static SprintGetRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SprintGetRequest>(create);
  static SprintGetRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$fixnum.Int64> get ids => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<$core.String> get names => $_getList(1);

  @$pb.TagNumber(3)
  $core.List<$fixnum.Int64> get projects => $_getList(2);

  @$pb.TagNumber(4)
  $0.Timestamp get startedAtAfter => $_getN(3);
  @$pb.TagNumber(4)
  set startedAtAfter($0.Timestamp v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasStartedAtAfter() => $_has(3);
  @$pb.TagNumber(4)
  void clearStartedAtAfter() => clearField(4);
  @$pb.TagNumber(4)
  $0.Timestamp ensureStartedAtAfter() => $_ensure(3);

  @$pb.TagNumber(5)
  $0.Timestamp get finishedAtBefore => $_getN(4);
  @$pb.TagNumber(5)
  set finishedAtBefore($0.Timestamp v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasFinishedAtBefore() => $_has(4);
  @$pb.TagNumber(5)
  void clearFinishedAtBefore() => clearField(5);
  @$pb.TagNumber(5)
  $0.Timestamp ensureFinishedAtBefore() => $_ensure(4);

  @$pb.TagNumber(6)
  $4.Sprint get status => $_getN(5);
  @$pb.TagNumber(6)
  set status($4.Sprint v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasStatus() => $_has(5);
  @$pb.TagNumber(6)
  void clearStatus() => clearField(6);
}

class SprintGetResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SprintGetResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..pc<Sprint>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sprints', $pb.PbFieldType.PM, subBuilder: Sprint.create)
    ..hasRequiredFields = false
  ;

  SprintGetResponse._() : super();
  factory SprintGetResponse({
    $core.Iterable<Sprint>? sprints,
  }) {
    final _result = create();
    if (sprints != null) {
      _result.sprints.addAll(sprints);
    }
    return _result;
  }
  factory SprintGetResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SprintGetResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SprintGetResponse clone() => SprintGetResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SprintGetResponse copyWith(void Function(SprintGetResponse) updates) => super.copyWith((message) => updates(message as SprintGetResponse)) as SprintGetResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SprintGetResponse create() => SprintGetResponse._();
  SprintGetResponse createEmptyInstance() => create();
  static $pb.PbList<SprintGetResponse> createRepeated() => $pb.PbList<SprintGetResponse>();
  @$core.pragma('dart2js:noInline')
  static SprintGetResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SprintGetResponse>(create);
  static SprintGetResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Sprint> get sprints => $_getList(0);
}

class SprintGetByIDRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SprintGetByIDRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  SprintGetByIDRequest._() : super();
  factory SprintGetByIDRequest({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory SprintGetByIDRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SprintGetByIDRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SprintGetByIDRequest clone() => SprintGetByIDRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SprintGetByIDRequest copyWith(void Function(SprintGetByIDRequest) updates) => super.copyWith((message) => updates(message as SprintGetByIDRequest)) as SprintGetByIDRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SprintGetByIDRequest create() => SprintGetByIDRequest._();
  SprintGetByIDRequest createEmptyInstance() => create();
  static $pb.PbList<SprintGetByIDRequest> createRepeated() => $pb.PbList<SprintGetByIDRequest>();
  @$core.pragma('dart2js:noInline')
  static SprintGetByIDRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SprintGetByIDRequest>(create);
  static SprintGetByIDRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class SprintGetByIDResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SprintGetByIDResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.sprint'), createEmptyInstance: create)
    ..aOM<Sprint>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sprint', subBuilder: Sprint.create)
    ..hasRequiredFields = false
  ;

  SprintGetByIDResponse._() : super();
  factory SprintGetByIDResponse({
    Sprint? sprint,
  }) {
    final _result = create();
    if (sprint != null) {
      _result.sprint = sprint;
    }
    return _result;
  }
  factory SprintGetByIDResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SprintGetByIDResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SprintGetByIDResponse clone() => SprintGetByIDResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SprintGetByIDResponse copyWith(void Function(SprintGetByIDResponse) updates) => super.copyWith((message) => updates(message as SprintGetByIDResponse)) as SprintGetByIDResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SprintGetByIDResponse create() => SprintGetByIDResponse._();
  SprintGetByIDResponse createEmptyInstance() => create();
  static $pb.PbList<SprintGetByIDResponse> createRepeated() => $pb.PbList<SprintGetByIDResponse>();
  @$core.pragma('dart2js:noInline')
  static SprintGetByIDResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SprintGetByIDResponse>(create);
  static SprintGetByIDResponse? _defaultInstance;

  @$pb.TagNumber(1)
  Sprint get sprint => $_getN(0);
  @$pb.TagNumber(1)
  set sprint(Sprint v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasSprint() => $_has(0);
  @$pb.TagNumber(1)
  void clearSprint() => clearField(1);
  @$pb.TagNumber(1)
  Sprint ensureSprint() => $_ensure(0);
}

class SprintRegistryApi {
  $pb.RpcClient _client;
  SprintRegistryApi(this._client);

  $async.Future<SprintCreateResponse> createSprint($pb.ClientContext? ctx, SprintCreateRequest request) {
    var emptyResponse = SprintCreateResponse();
    return _client.invoke<SprintCreateResponse>(ctx, 'SprintRegistry', 'CreateSprint', request, emptyResponse);
  }
  $async.Future<AddIssueResponse> addIssue($pb.ClientContext? ctx, AddIssueRequest request) {
    var emptyResponse = AddIssueResponse();
    return _client.invoke<AddIssueResponse>(ctx, 'SprintRegistry', 'AddIssue', request, emptyResponse);
  }
  $async.Future<RemoveIssueResponse> removeIssue($pb.ClientContext? ctx, RemoveIssueRequest request) {
    var emptyResponse = RemoveIssueResponse();
    return _client.invoke<RemoveIssueResponse>(ctx, 'SprintRegistry', 'RemoveIssue', request, emptyResponse);
  }
  $async.Future<SprintUpdateResponse> updateSprint($pb.ClientContext? ctx, SprintUpdateRequest request) {
    var emptyResponse = SprintUpdateResponse();
    return _client.invoke<SprintUpdateResponse>(ctx, 'SprintRegistry', 'UpdateSprint', request, emptyResponse);
  }
  $async.Future<SprintChangeStatusResponse> sprintChangeStatus($pb.ClientContext? ctx, SprintChangeStatusRequest request) {
    var emptyResponse = SprintChangeStatusResponse();
    return _client.invoke<SprintChangeStatusResponse>(ctx, 'SprintRegistry', 'SprintChangeStatus', request, emptyResponse);
  }
  $async.Future<SprintGetResponse> getSprint($pb.ClientContext? ctx, SprintGetRequest request) {
    var emptyResponse = SprintGetResponse();
    return _client.invoke<SprintGetResponse>(ctx, 'SprintRegistry', 'GetSprint', request, emptyResponse);
  }
  $async.Future<SprintGetByIDResponse> getSprintByID($pb.ClientContext? ctx, SprintGetByIDRequest request) {
    var emptyResponse = SprintGetByIDResponse();
    return _client.invoke<SprintGetByIDResponse>(ctx, 'SprintRegistry', 'GetSprintByID', request, emptyResponse);
  }
}

