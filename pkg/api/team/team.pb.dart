///
//  Generated code. Do not modify.
//  source: api/team/team.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class Team extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Team', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.team'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..pPS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'members')
    ..p<$fixnum.Int64>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'projects', $pb.PbFieldType.K6)
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'techOwner')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'businessOwner')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'department')
    ..hasRequiredFields = false
  ;

  Team._() : super();
  factory Team({
    $fixnum.Int64? id,
    $core.String? name,
    $core.Iterable<$core.String>? members,
    $core.Iterable<$fixnum.Int64>? projects,
    $core.String? techOwner,
    $core.String? businessOwner,
    $core.String? department,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (name != null) {
      _result.name = name;
    }
    if (members != null) {
      _result.members.addAll(members);
    }
    if (projects != null) {
      _result.projects.addAll(projects);
    }
    if (techOwner != null) {
      _result.techOwner = techOwner;
    }
    if (businessOwner != null) {
      _result.businessOwner = businessOwner;
    }
    if (department != null) {
      _result.department = department;
    }
    return _result;
  }
  factory Team.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Team.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Team clone() => Team()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Team copyWith(void Function(Team) updates) => super.copyWith((message) => updates(message as Team)) as Team; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Team create() => Team._();
  Team createEmptyInstance() => create();
  static $pb.PbList<Team> createRepeated() => $pb.PbList<Team>();
  @$core.pragma('dart2js:noInline')
  static Team getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Team>(create);
  static Team? _defaultInstance;

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
  $core.List<$core.String> get members => $_getList(2);

  @$pb.TagNumber(4)
  $core.List<$fixnum.Int64> get projects => $_getList(3);

  @$pb.TagNumber(5)
  $core.String get techOwner => $_getSZ(4);
  @$pb.TagNumber(5)
  set techOwner($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasTechOwner() => $_has(4);
  @$pb.TagNumber(5)
  void clearTechOwner() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get businessOwner => $_getSZ(5);
  @$pb.TagNumber(6)
  set businessOwner($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasBusinessOwner() => $_has(5);
  @$pb.TagNumber(6)
  void clearBusinessOwner() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get department => $_getSZ(6);
  @$pb.TagNumber(7)
  set department($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasDepartment() => $_has(6);
  @$pb.TagNumber(7)
  void clearDepartment() => clearField(7);
}

class TeamCreateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TeamCreateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.team'), createEmptyInstance: create)
    ..aOM<Team>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'team', subBuilder: Team.create)
    ..hasRequiredFields = false
  ;

  TeamCreateRequest._() : super();
  factory TeamCreateRequest({
    Team? team,
  }) {
    final _result = create();
    if (team != null) {
      _result.team = team;
    }
    return _result;
  }
  factory TeamCreateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TeamCreateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TeamCreateRequest clone() => TeamCreateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TeamCreateRequest copyWith(void Function(TeamCreateRequest) updates) => super.copyWith((message) => updates(message as TeamCreateRequest)) as TeamCreateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TeamCreateRequest create() => TeamCreateRequest._();
  TeamCreateRequest createEmptyInstance() => create();
  static $pb.PbList<TeamCreateRequest> createRepeated() => $pb.PbList<TeamCreateRequest>();
  @$core.pragma('dart2js:noInline')
  static TeamCreateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TeamCreateRequest>(create);
  static TeamCreateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  Team get team => $_getN(0);
  @$pb.TagNumber(1)
  set team(Team v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTeam() => $_has(0);
  @$pb.TagNumber(1)
  void clearTeam() => clearField(1);
  @$pb.TagNumber(1)
  Team ensureTeam() => $_ensure(0);
}

class TeamCreateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TeamCreateResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.team'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  TeamCreateResponse._() : super();
  factory TeamCreateResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory TeamCreateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TeamCreateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TeamCreateResponse clone() => TeamCreateResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TeamCreateResponse copyWith(void Function(TeamCreateResponse) updates) => super.copyWith((message) => updates(message as TeamCreateResponse)) as TeamCreateResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TeamCreateResponse create() => TeamCreateResponse._();
  TeamCreateResponse createEmptyInstance() => create();
  static $pb.PbList<TeamCreateResponse> createRepeated() => $pb.PbList<TeamCreateResponse>();
  @$core.pragma('dart2js:noInline')
  static TeamCreateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TeamCreateResponse>(create);
  static TeamCreateResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class TeamUpdateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TeamUpdateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.team'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOM<Team>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'team', subBuilder: Team.create)
    ..hasRequiredFields = false
  ;

  TeamUpdateRequest._() : super();
  factory TeamUpdateRequest({
    $fixnum.Int64? id,
    Team? team,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (team != null) {
      _result.team = team;
    }
    return _result;
  }
  factory TeamUpdateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TeamUpdateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TeamUpdateRequest clone() => TeamUpdateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TeamUpdateRequest copyWith(void Function(TeamUpdateRequest) updates) => super.copyWith((message) => updates(message as TeamUpdateRequest)) as TeamUpdateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TeamUpdateRequest create() => TeamUpdateRequest._();
  TeamUpdateRequest createEmptyInstance() => create();
  static $pb.PbList<TeamUpdateRequest> createRepeated() => $pb.PbList<TeamUpdateRequest>();
  @$core.pragma('dart2js:noInline')
  static TeamUpdateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TeamUpdateRequest>(create);
  static TeamUpdateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  Team get team => $_getN(1);
  @$pb.TagNumber(2)
  set team(Team v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasTeam() => $_has(1);
  @$pb.TagNumber(2)
  void clearTeam() => clearField(2);
  @$pb.TagNumber(2)
  Team ensureTeam() => $_ensure(1);
}

class TeamUpdateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TeamUpdateResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.team'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  TeamUpdateResponse._() : super();
  factory TeamUpdateResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory TeamUpdateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TeamUpdateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TeamUpdateResponse clone() => TeamUpdateResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TeamUpdateResponse copyWith(void Function(TeamUpdateResponse) updates) => super.copyWith((message) => updates(message as TeamUpdateResponse)) as TeamUpdateResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TeamUpdateResponse create() => TeamUpdateResponse._();
  TeamUpdateResponse createEmptyInstance() => create();
  static $pb.PbList<TeamUpdateResponse> createRepeated() => $pb.PbList<TeamUpdateResponse>();
  @$core.pragma('dart2js:noInline')
  static TeamUpdateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TeamUpdateResponse>(create);
  static TeamUpdateResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class TeamGetRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TeamGetRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.team'), createEmptyInstance: create)
    ..p<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ids', $pb.PbFieldType.K6)
    ..pPS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'names')
    ..pPS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'departments')
    ..hasRequiredFields = false
  ;

  TeamGetRequest._() : super();
  factory TeamGetRequest({
    $core.Iterable<$fixnum.Int64>? ids,
    $core.Iterable<$core.String>? names,
    $core.Iterable<$core.String>? departments,
  }) {
    final _result = create();
    if (ids != null) {
      _result.ids.addAll(ids);
    }
    if (names != null) {
      _result.names.addAll(names);
    }
    if (departments != null) {
      _result.departments.addAll(departments);
    }
    return _result;
  }
  factory TeamGetRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TeamGetRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TeamGetRequest clone() => TeamGetRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TeamGetRequest copyWith(void Function(TeamGetRequest) updates) => super.copyWith((message) => updates(message as TeamGetRequest)) as TeamGetRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TeamGetRequest create() => TeamGetRequest._();
  TeamGetRequest createEmptyInstance() => create();
  static $pb.PbList<TeamGetRequest> createRepeated() => $pb.PbList<TeamGetRequest>();
  @$core.pragma('dart2js:noInline')
  static TeamGetRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TeamGetRequest>(create);
  static TeamGetRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$fixnum.Int64> get ids => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<$core.String> get names => $_getList(1);

  @$pb.TagNumber(3)
  $core.List<$core.String> get departments => $_getList(2);
}

class TeamGetResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TeamGetResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.team'), createEmptyInstance: create)
    ..pc<Team>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'teams', $pb.PbFieldType.PM, subBuilder: Team.create)
    ..hasRequiredFields = false
  ;

  TeamGetResponse._() : super();
  factory TeamGetResponse({
    $core.Iterable<Team>? teams,
  }) {
    final _result = create();
    if (teams != null) {
      _result.teams.addAll(teams);
    }
    return _result;
  }
  factory TeamGetResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TeamGetResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TeamGetResponse clone() => TeamGetResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TeamGetResponse copyWith(void Function(TeamGetResponse) updates) => super.copyWith((message) => updates(message as TeamGetResponse)) as TeamGetResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TeamGetResponse create() => TeamGetResponse._();
  TeamGetResponse createEmptyInstance() => create();
  static $pb.PbList<TeamGetResponse> createRepeated() => $pb.PbList<TeamGetResponse>();
  @$core.pragma('dart2js:noInline')
  static TeamGetResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TeamGetResponse>(create);
  static TeamGetResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Team> get teams => $_getList(0);
}

class TeamGetByIDRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TeamGetByIDRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.team'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  TeamGetByIDRequest._() : super();
  factory TeamGetByIDRequest({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory TeamGetByIDRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TeamGetByIDRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TeamGetByIDRequest clone() => TeamGetByIDRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TeamGetByIDRequest copyWith(void Function(TeamGetByIDRequest) updates) => super.copyWith((message) => updates(message as TeamGetByIDRequest)) as TeamGetByIDRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TeamGetByIDRequest create() => TeamGetByIDRequest._();
  TeamGetByIDRequest createEmptyInstance() => create();
  static $pb.PbList<TeamGetByIDRequest> createRepeated() => $pb.PbList<TeamGetByIDRequest>();
  @$core.pragma('dart2js:noInline')
  static TeamGetByIDRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TeamGetByIDRequest>(create);
  static TeamGetByIDRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class TeamGetByIDResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TeamGetByIDResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.team'), createEmptyInstance: create)
    ..aOM<Team>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'team', subBuilder: Team.create)
    ..hasRequiredFields = false
  ;

  TeamGetByIDResponse._() : super();
  factory TeamGetByIDResponse({
    Team? team,
  }) {
    final _result = create();
    if (team != null) {
      _result.team = team;
    }
    return _result;
  }
  factory TeamGetByIDResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TeamGetByIDResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TeamGetByIDResponse clone() => TeamGetByIDResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TeamGetByIDResponse copyWith(void Function(TeamGetByIDResponse) updates) => super.copyWith((message) => updates(message as TeamGetByIDResponse)) as TeamGetByIDResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TeamGetByIDResponse create() => TeamGetByIDResponse._();
  TeamGetByIDResponse createEmptyInstance() => create();
  static $pb.PbList<TeamGetByIDResponse> createRepeated() => $pb.PbList<TeamGetByIDResponse>();
  @$core.pragma('dart2js:noInline')
  static TeamGetByIDResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TeamGetByIDResponse>(create);
  static TeamGetByIDResponse? _defaultInstance;

  @$pb.TagNumber(1)
  Team get team => $_getN(0);
  @$pb.TagNumber(1)
  set team(Team v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTeam() => $_has(0);
  @$pb.TagNumber(1)
  void clearTeam() => clearField(1);
  @$pb.TagNumber(1)
  Team ensureTeam() => $_ensure(0);
}

class TeamRegistryApi {
  $pb.RpcClient _client;
  TeamRegistryApi(this._client);

  $async.Future<TeamCreateResponse> createTeam($pb.ClientContext? ctx, TeamCreateRequest request) {
    var emptyResponse = TeamCreateResponse();
    return _client.invoke<TeamCreateResponse>(ctx, 'TeamRegistry', 'CreateTeam', request, emptyResponse);
  }
  $async.Future<TeamUpdateResponse> updateTeam($pb.ClientContext? ctx, TeamUpdateRequest request) {
    var emptyResponse = TeamUpdateResponse();
    return _client.invoke<TeamUpdateResponse>(ctx, 'TeamRegistry', 'UpdateTeam', request, emptyResponse);
  }
  $async.Future<TeamGetResponse> getTeams($pb.ClientContext? ctx, TeamGetRequest request) {
    var emptyResponse = TeamGetResponse();
    return _client.invoke<TeamGetResponse>(ctx, 'TeamRegistry', 'GetTeams', request, emptyResponse);
  }
  $async.Future<TeamGetByIDResponse> getTeamByID($pb.ClientContext? ctx, TeamGetByIDRequest request) {
    var emptyResponse = TeamGetByIDResponse();
    return _client.invoke<TeamGetByIDResponse>(ctx, 'TeamRegistry', 'GetTeamByID', request, emptyResponse);
  }
}

