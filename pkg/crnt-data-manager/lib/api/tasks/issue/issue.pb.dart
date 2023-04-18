///
//  Generated code. Do not modify.
//  source: api/tasks/issue/issue.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import '../../comments/comments.pb.dart' as $3;
import '../../../google/protobuf/timestamp.pb.dart' as $0;
import '../../state/status/status.pb.dart' as $4;

import 'issue.pbenum.dart';
import '../../state/template/template.pbenum.dart' as $5;
import '../../state/priority/priority.pbenum.dart' as $6;

export 'issue.pbenum.dart';

class Issue extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Issue', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'compositeName')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..e<IssueType>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: IssueType.ISSUE_TYPE_UNKNOWN, valueOf: IssueType.valueOf, enumValues: IssueType.values)
    ..aInt64(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'parentId')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'description')
    ..aOM<$3.Comments>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'comments', subBuilder: $3.Comments.create)
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'author')
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'assigned')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'qa')
    ..aOS(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'reviewer')
    ..e<$5.Template>(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'template', $pb.PbFieldType.OE, defaultOrMaker: $5.Template.TEMPLATE_UNKNOWN, valueOf: $5.Template.valueOf, enumValues: $5.Template.values)
    ..aOM<$0.Timestamp>(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'createdAt', subBuilder: $0.Timestamp.create)
    ..aOM<$0.Timestamp>(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'deadline', subBuilder: $0.Timestamp.create)
    ..aOM<$0.Timestamp>(15, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'updatedAt', subBuilder: $0.Timestamp.create)
    ..aOM<$4.IssueStatus>(16, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', subBuilder: $4.IssueStatus.create)
    ..e<$6.Priority>(17, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'priority', $pb.PbFieldType.OE, defaultOrMaker: $6.Priority.PRIORITY_UNKNOWN, valueOf: $6.Priority.valueOf, enumValues: $6.Priority.values)
    ..aInt64(18, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sprintId')
    ..aInt64(19, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'projectId')
    ..p<$fixnum.Int64>(20, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'componentIds', $pb.PbFieldType.K6)
    ..aInt64(21, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'storyPoints')
    ..a<$core.double>(22, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'payment', $pb.PbFieldType.OD)
    ..aInt64(23, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'timeSpent')
    ..pc<Issue>(24, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'children', $pb.PbFieldType.PM, subBuilder: Issue.create)
    ..hasRequiredFields = false
  ;

  Issue._() : super();
  factory Issue({
    $fixnum.Int64? id,
    $core.String? compositeName,
    $core.String? name,
    IssueType? type,
    $fixnum.Int64? parentId,
    $core.String? description,
    $3.Comments? comments,
    $core.String? author,
    $core.String? assigned,
    $core.String? qa,
    $core.String? reviewer,
    $5.Template? template,
    $0.Timestamp? createdAt,
    $0.Timestamp? deadline,
    $0.Timestamp? updatedAt,
    $4.IssueStatus? status,
    $6.Priority? priority,
    $fixnum.Int64? sprintId,
    $fixnum.Int64? projectId,
    $core.Iterable<$fixnum.Int64>? componentIds,
    $fixnum.Int64? storyPoints,
    $core.double? payment,
    $fixnum.Int64? timeSpent,
    $core.Iterable<Issue>? children,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (compositeName != null) {
      _result.compositeName = compositeName;
    }
    if (name != null) {
      _result.name = name;
    }
    if (type != null) {
      _result.type = type;
    }
    if (parentId != null) {
      _result.parentId = parentId;
    }
    if (description != null) {
      _result.description = description;
    }
    if (comments != null) {
      _result.comments = comments;
    }
    if (author != null) {
      _result.author = author;
    }
    if (assigned != null) {
      _result.assigned = assigned;
    }
    if (qa != null) {
      _result.qa = qa;
    }
    if (reviewer != null) {
      _result.reviewer = reviewer;
    }
    if (template != null) {
      _result.template = template;
    }
    if (createdAt != null) {
      _result.createdAt = createdAt;
    }
    if (deadline != null) {
      _result.deadline = deadline;
    }
    if (updatedAt != null) {
      _result.updatedAt = updatedAt;
    }
    if (status != null) {
      _result.status = status;
    }
    if (priority != null) {
      _result.priority = priority;
    }
    if (sprintId != null) {
      _result.sprintId = sprintId;
    }
    if (projectId != null) {
      _result.projectId = projectId;
    }
    if (componentIds != null) {
      _result.componentIds.addAll(componentIds);
    }
    if (storyPoints != null) {
      _result.storyPoints = storyPoints;
    }
    if (payment != null) {
      _result.payment = payment;
    }
    if (timeSpent != null) {
      _result.timeSpent = timeSpent;
    }
    if (children != null) {
      _result.children.addAll(children);
    }
    return _result;
  }
  factory Issue.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Issue.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Issue clone() => Issue()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Issue copyWith(void Function(Issue) updates) => super.copyWith((message) => updates(message as Issue)) as Issue; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Issue create() => Issue._();
  Issue createEmptyInstance() => create();
  static $pb.PbList<Issue> createRepeated() => $pb.PbList<Issue>();
  @$core.pragma('dart2js:noInline')
  static Issue getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Issue>(create);
  static Issue? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get compositeName => $_getSZ(1);
  @$pb.TagNumber(2)
  set compositeName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCompositeName() => $_has(1);
  @$pb.TagNumber(2)
  void clearCompositeName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(2);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);

  @$pb.TagNumber(4)
  IssueType get type => $_getN(3);
  @$pb.TagNumber(4)
  set type(IssueType v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasType() => $_has(3);
  @$pb.TagNumber(4)
  void clearType() => clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get parentId => $_getI64(4);
  @$pb.TagNumber(5)
  set parentId($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasParentId() => $_has(4);
  @$pb.TagNumber(5)
  void clearParentId() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get description => $_getSZ(5);
  @$pb.TagNumber(6)
  set description($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasDescription() => $_has(5);
  @$pb.TagNumber(6)
  void clearDescription() => clearField(6);

  @$pb.TagNumber(7)
  $3.Comments get comments => $_getN(6);
  @$pb.TagNumber(7)
  set comments($3.Comments v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasComments() => $_has(6);
  @$pb.TagNumber(7)
  void clearComments() => clearField(7);
  @$pb.TagNumber(7)
  $3.Comments ensureComments() => $_ensure(6);

  @$pb.TagNumber(8)
  $core.String get author => $_getSZ(7);
  @$pb.TagNumber(8)
  set author($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasAuthor() => $_has(7);
  @$pb.TagNumber(8)
  void clearAuthor() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get assigned => $_getSZ(8);
  @$pb.TagNumber(9)
  set assigned($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasAssigned() => $_has(8);
  @$pb.TagNumber(9)
  void clearAssigned() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get qa => $_getSZ(9);
  @$pb.TagNumber(10)
  set qa($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasQa() => $_has(9);
  @$pb.TagNumber(10)
  void clearQa() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get reviewer => $_getSZ(10);
  @$pb.TagNumber(11)
  set reviewer($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasReviewer() => $_has(10);
  @$pb.TagNumber(11)
  void clearReviewer() => clearField(11);

  @$pb.TagNumber(12)
  $5.Template get template => $_getN(11);
  @$pb.TagNumber(12)
  set template($5.Template v) { setField(12, v); }
  @$pb.TagNumber(12)
  $core.bool hasTemplate() => $_has(11);
  @$pb.TagNumber(12)
  void clearTemplate() => clearField(12);

  @$pb.TagNumber(13)
  $0.Timestamp get createdAt => $_getN(12);
  @$pb.TagNumber(13)
  set createdAt($0.Timestamp v) { setField(13, v); }
  @$pb.TagNumber(13)
  $core.bool hasCreatedAt() => $_has(12);
  @$pb.TagNumber(13)
  void clearCreatedAt() => clearField(13);
  @$pb.TagNumber(13)
  $0.Timestamp ensureCreatedAt() => $_ensure(12);

  @$pb.TagNumber(14)
  $0.Timestamp get deadline => $_getN(13);
  @$pb.TagNumber(14)
  set deadline($0.Timestamp v) { setField(14, v); }
  @$pb.TagNumber(14)
  $core.bool hasDeadline() => $_has(13);
  @$pb.TagNumber(14)
  void clearDeadline() => clearField(14);
  @$pb.TagNumber(14)
  $0.Timestamp ensureDeadline() => $_ensure(13);

  @$pb.TagNumber(15)
  $0.Timestamp get updatedAt => $_getN(14);
  @$pb.TagNumber(15)
  set updatedAt($0.Timestamp v) { setField(15, v); }
  @$pb.TagNumber(15)
  $core.bool hasUpdatedAt() => $_has(14);
  @$pb.TagNumber(15)
  void clearUpdatedAt() => clearField(15);
  @$pb.TagNumber(15)
  $0.Timestamp ensureUpdatedAt() => $_ensure(14);

  @$pb.TagNumber(16)
  $4.IssueStatus get status => $_getN(15);
  @$pb.TagNumber(16)
  set status($4.IssueStatus v) { setField(16, v); }
  @$pb.TagNumber(16)
  $core.bool hasStatus() => $_has(15);
  @$pb.TagNumber(16)
  void clearStatus() => clearField(16);
  @$pb.TagNumber(16)
  $4.IssueStatus ensureStatus() => $_ensure(15);

  @$pb.TagNumber(17)
  $6.Priority get priority => $_getN(16);
  @$pb.TagNumber(17)
  set priority($6.Priority v) { setField(17, v); }
  @$pb.TagNumber(17)
  $core.bool hasPriority() => $_has(16);
  @$pb.TagNumber(17)
  void clearPriority() => clearField(17);

  @$pb.TagNumber(18)
  $fixnum.Int64 get sprintId => $_getI64(17);
  @$pb.TagNumber(18)
  set sprintId($fixnum.Int64 v) { $_setInt64(17, v); }
  @$pb.TagNumber(18)
  $core.bool hasSprintId() => $_has(17);
  @$pb.TagNumber(18)
  void clearSprintId() => clearField(18);

  @$pb.TagNumber(19)
  $fixnum.Int64 get projectId => $_getI64(18);
  @$pb.TagNumber(19)
  set projectId($fixnum.Int64 v) { $_setInt64(18, v); }
  @$pb.TagNumber(19)
  $core.bool hasProjectId() => $_has(18);
  @$pb.TagNumber(19)
  void clearProjectId() => clearField(19);

  @$pb.TagNumber(20)
  $core.List<$fixnum.Int64> get componentIds => $_getList(19);

  @$pb.TagNumber(21)
  $fixnum.Int64 get storyPoints => $_getI64(20);
  @$pb.TagNumber(21)
  set storyPoints($fixnum.Int64 v) { $_setInt64(20, v); }
  @$pb.TagNumber(21)
  $core.bool hasStoryPoints() => $_has(20);
  @$pb.TagNumber(21)
  void clearStoryPoints() => clearField(21);

  @$pb.TagNumber(22)
  $core.double get payment => $_getN(21);
  @$pb.TagNumber(22)
  set payment($core.double v) { $_setDouble(21, v); }
  @$pb.TagNumber(22)
  $core.bool hasPayment() => $_has(21);
  @$pb.TagNumber(22)
  void clearPayment() => clearField(22);

  @$pb.TagNumber(23)
  $fixnum.Int64 get timeSpent => $_getI64(22);
  @$pb.TagNumber(23)
  set timeSpent($fixnum.Int64 v) { $_setInt64(22, v); }
  @$pb.TagNumber(23)
  $core.bool hasTimeSpent() => $_has(22);
  @$pb.TagNumber(23)
  void clearTimeSpent() => clearField(23);

  @$pb.TagNumber(24)
  $core.List<Issue> get children => $_getList(23);
}

class IssueInfo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueInfo', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'compositeName')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..e<IssueType>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: IssueType.ISSUE_TYPE_UNKNOWN, valueOf: IssueType.valueOf, enumValues: IssueType.values)
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'assigned')
    ..e<$6.Priority>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'priority', $pb.PbFieldType.OE, defaultOrMaker: $6.Priority.PRIORITY_UNKNOWN, valueOf: $6.Priority.valueOf, enumValues: $6.Priority.values)
    ..aOM<$4.IssueStatus>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', subBuilder: $4.IssueStatus.create)
    ..aInt64(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'storyPoints')
    ..hasRequiredFields = false
  ;

  IssueInfo._() : super();
  factory IssueInfo({
    $fixnum.Int64? id,
    $core.String? compositeName,
    $core.String? name,
    IssueType? type,
    $core.String? assigned,
    $6.Priority? priority,
    $4.IssueStatus? status,
    $fixnum.Int64? storyPoints,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (compositeName != null) {
      _result.compositeName = compositeName;
    }
    if (name != null) {
      _result.name = name;
    }
    if (type != null) {
      _result.type = type;
    }
    if (assigned != null) {
      _result.assigned = assigned;
    }
    if (priority != null) {
      _result.priority = priority;
    }
    if (status != null) {
      _result.status = status;
    }
    if (storyPoints != null) {
      _result.storyPoints = storyPoints;
    }
    return _result;
  }
  factory IssueInfo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueInfo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueInfo clone() => IssueInfo()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueInfo copyWith(void Function(IssueInfo) updates) => super.copyWith((message) => updates(message as IssueInfo)) as IssueInfo; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueInfo create() => IssueInfo._();
  IssueInfo createEmptyInstance() => create();
  static $pb.PbList<IssueInfo> createRepeated() => $pb.PbList<IssueInfo>();
  @$core.pragma('dart2js:noInline')
  static IssueInfo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueInfo>(create);
  static IssueInfo? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get compositeName => $_getSZ(1);
  @$pb.TagNumber(2)
  set compositeName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCompositeName() => $_has(1);
  @$pb.TagNumber(2)
  void clearCompositeName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(2);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);

  @$pb.TagNumber(4)
  IssueType get type => $_getN(3);
  @$pb.TagNumber(4)
  set type(IssueType v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasType() => $_has(3);
  @$pb.TagNumber(4)
  void clearType() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get assigned => $_getSZ(4);
  @$pb.TagNumber(5)
  set assigned($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasAssigned() => $_has(4);
  @$pb.TagNumber(5)
  void clearAssigned() => clearField(5);

  @$pb.TagNumber(6)
  $6.Priority get priority => $_getN(5);
  @$pb.TagNumber(6)
  set priority($6.Priority v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasPriority() => $_has(5);
  @$pb.TagNumber(6)
  void clearPriority() => clearField(6);

  @$pb.TagNumber(7)
  $4.IssueStatus get status => $_getN(6);
  @$pb.TagNumber(7)
  set status($4.IssueStatus v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasStatus() => $_has(6);
  @$pb.TagNumber(7)
  void clearStatus() => clearField(7);
  @$pb.TagNumber(7)
  $4.IssueStatus ensureStatus() => $_ensure(6);

  @$pb.TagNumber(8)
  $fixnum.Int64 get storyPoints => $_getI64(7);
  @$pb.TagNumber(8)
  set storyPoints($fixnum.Int64 v) { $_setInt64(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasStoryPoints() => $_has(7);
  @$pb.TagNumber(8)
  void clearStoryPoints() => clearField(8);
}

class IssueCreateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueCreateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aOM<Issue>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'issue', subBuilder: Issue.create)
    ..hasRequiredFields = false
  ;

  IssueCreateRequest._() : super();
  factory IssueCreateRequest({
    Issue? issue,
  }) {
    final _result = create();
    if (issue != null) {
      _result.issue = issue;
    }
    return _result;
  }
  factory IssueCreateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueCreateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueCreateRequest clone() => IssueCreateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueCreateRequest copyWith(void Function(IssueCreateRequest) updates) => super.copyWith((message) => updates(message as IssueCreateRequest)) as IssueCreateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueCreateRequest create() => IssueCreateRequest._();
  IssueCreateRequest createEmptyInstance() => create();
  static $pb.PbList<IssueCreateRequest> createRepeated() => $pb.PbList<IssueCreateRequest>();
  @$core.pragma('dart2js:noInline')
  static IssueCreateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueCreateRequest>(create);
  static IssueCreateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  Issue get issue => $_getN(0);
  @$pb.TagNumber(1)
  set issue(Issue v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasIssue() => $_has(0);
  @$pb.TagNumber(1)
  void clearIssue() => clearField(1);
  @$pb.TagNumber(1)
  Issue ensureIssue() => $_ensure(0);
}

class IssueCreateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueCreateResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  IssueCreateResponse._() : super();
  factory IssueCreateResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory IssueCreateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueCreateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueCreateResponse clone() => IssueCreateResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueCreateResponse copyWith(void Function(IssueCreateResponse) updates) => super.copyWith((message) => updates(message as IssueCreateResponse)) as IssueCreateResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueCreateResponse create() => IssueCreateResponse._();
  IssueCreateResponse createEmptyInstance() => create();
  static $pb.PbList<IssueCreateResponse> createRepeated() => $pb.PbList<IssueCreateResponse>();
  @$core.pragma('dart2js:noInline')
  static IssueCreateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueCreateResponse>(create);
  static IssueCreateResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class IssueCreateSubtaskRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueCreateSubtaskRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOM<Issue>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'child', subBuilder: Issue.create)
    ..hasRequiredFields = false
  ;

  IssueCreateSubtaskRequest._() : super();
  factory IssueCreateSubtaskRequest({
    $fixnum.Int64? id,
    Issue? child,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (child != null) {
      _result.child = child;
    }
    return _result;
  }
  factory IssueCreateSubtaskRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueCreateSubtaskRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueCreateSubtaskRequest clone() => IssueCreateSubtaskRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueCreateSubtaskRequest copyWith(void Function(IssueCreateSubtaskRequest) updates) => super.copyWith((message) => updates(message as IssueCreateSubtaskRequest)) as IssueCreateSubtaskRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueCreateSubtaskRequest create() => IssueCreateSubtaskRequest._();
  IssueCreateSubtaskRequest createEmptyInstance() => create();
  static $pb.PbList<IssueCreateSubtaskRequest> createRepeated() => $pb.PbList<IssueCreateSubtaskRequest>();
  @$core.pragma('dart2js:noInline')
  static IssueCreateSubtaskRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueCreateSubtaskRequest>(create);
  static IssueCreateSubtaskRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  Issue get child => $_getN(1);
  @$pb.TagNumber(2)
  set child(Issue v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasChild() => $_has(1);
  @$pb.TagNumber(2)
  void clearChild() => clearField(2);
  @$pb.TagNumber(2)
  Issue ensureChild() => $_ensure(1);
}

class IssueCreateSubtaskResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueCreateSubtaskResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  IssueCreateSubtaskResponse._() : super();
  factory IssueCreateSubtaskResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory IssueCreateSubtaskResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueCreateSubtaskResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueCreateSubtaskResponse clone() => IssueCreateSubtaskResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueCreateSubtaskResponse copyWith(void Function(IssueCreateSubtaskResponse) updates) => super.copyWith((message) => updates(message as IssueCreateSubtaskResponse)) as IssueCreateSubtaskResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueCreateSubtaskResponse create() => IssueCreateSubtaskResponse._();
  IssueCreateSubtaskResponse createEmptyInstance() => create();
  static $pb.PbList<IssueCreateSubtaskResponse> createRepeated() => $pb.PbList<IssueCreateSubtaskResponse>();
  @$core.pragma('dart2js:noInline')
  static IssueCreateSubtaskResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueCreateSubtaskResponse>(create);
  static IssueCreateSubtaskResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class IssueAddCommentRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueAddCommentRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOM<$3.Comment>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'comment', subBuilder: $3.Comment.create)
    ..hasRequiredFields = false
  ;

  IssueAddCommentRequest._() : super();
  factory IssueAddCommentRequest({
    $fixnum.Int64? id,
    $3.Comment? comment,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (comment != null) {
      _result.comment = comment;
    }
    return _result;
  }
  factory IssueAddCommentRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueAddCommentRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueAddCommentRequest clone() => IssueAddCommentRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueAddCommentRequest copyWith(void Function(IssueAddCommentRequest) updates) => super.copyWith((message) => updates(message as IssueAddCommentRequest)) as IssueAddCommentRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueAddCommentRequest create() => IssueAddCommentRequest._();
  IssueAddCommentRequest createEmptyInstance() => create();
  static $pb.PbList<IssueAddCommentRequest> createRepeated() => $pb.PbList<IssueAddCommentRequest>();
  @$core.pragma('dart2js:noInline')
  static IssueAddCommentRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueAddCommentRequest>(create);
  static IssueAddCommentRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $3.Comment get comment => $_getN(1);
  @$pb.TagNumber(2)
  set comment($3.Comment v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasComment() => $_has(1);
  @$pb.TagNumber(2)
  void clearComment() => clearField(2);
  @$pb.TagNumber(2)
  $3.Comment ensureComment() => $_ensure(1);
}

class IssueAddCommentResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueAddCommentResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  IssueAddCommentResponse._() : super();
  factory IssueAddCommentResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory IssueAddCommentResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueAddCommentResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueAddCommentResponse clone() => IssueAddCommentResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueAddCommentResponse copyWith(void Function(IssueAddCommentResponse) updates) => super.copyWith((message) => updates(message as IssueAddCommentResponse)) as IssueAddCommentResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueAddCommentResponse create() => IssueAddCommentResponse._();
  IssueAddCommentResponse createEmptyInstance() => create();
  static $pb.PbList<IssueAddCommentResponse> createRepeated() => $pb.PbList<IssueAddCommentResponse>();
  @$core.pragma('dart2js:noInline')
  static IssueAddCommentResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueAddCommentResponse>(create);
  static IssueAddCommentResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class IssueUpdateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueUpdateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOM<Issue>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'issue', subBuilder: Issue.create)
    ..hasRequiredFields = false
  ;

  IssueUpdateRequest._() : super();
  factory IssueUpdateRequest({
    $fixnum.Int64? id,
    Issue? issue,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (issue != null) {
      _result.issue = issue;
    }
    return _result;
  }
  factory IssueUpdateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueUpdateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueUpdateRequest clone() => IssueUpdateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueUpdateRequest copyWith(void Function(IssueUpdateRequest) updates) => super.copyWith((message) => updates(message as IssueUpdateRequest)) as IssueUpdateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueUpdateRequest create() => IssueUpdateRequest._();
  IssueUpdateRequest createEmptyInstance() => create();
  static $pb.PbList<IssueUpdateRequest> createRepeated() => $pb.PbList<IssueUpdateRequest>();
  @$core.pragma('dart2js:noInline')
  static IssueUpdateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueUpdateRequest>(create);
  static IssueUpdateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  Issue get issue => $_getN(1);
  @$pb.TagNumber(2)
  set issue(Issue v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasIssue() => $_has(1);
  @$pb.TagNumber(2)
  void clearIssue() => clearField(2);
  @$pb.TagNumber(2)
  Issue ensureIssue() => $_ensure(1);
}

class IssueUpdateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueUpdateResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  IssueUpdateResponse._() : super();
  factory IssueUpdateResponse({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory IssueUpdateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueUpdateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueUpdateResponse clone() => IssueUpdateResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueUpdateResponse copyWith(void Function(IssueUpdateResponse) updates) => super.copyWith((message) => updates(message as IssueUpdateResponse)) as IssueUpdateResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueUpdateResponse create() => IssueUpdateResponse._();
  IssueUpdateResponse createEmptyInstance() => create();
  static $pb.PbList<IssueUpdateResponse> createRepeated() => $pb.PbList<IssueUpdateResponse>();
  @$core.pragma('dart2js:noInline')
  static IssueUpdateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueUpdateResponse>(create);
  static IssueUpdateResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class IssueGetRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueGetRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..p<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ids', $pb.PbFieldType.K6)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aInt64(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'parentId')
    ..e<IssueType>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: IssueType.ISSUE_TYPE_UNKNOWN, valueOf: IssueType.valueOf, enumValues: IssueType.values)
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'author')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'assigned')
    ..aInt64(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sprintId')
    ..aInt64(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'projectId')
    ..aOM<$4.IssueStatus>(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', subBuilder: $4.IssueStatus.create)
    ..hasRequiredFields = false
  ;

  IssueGetRequest._() : super();
  factory IssueGetRequest({
    $core.Iterable<$fixnum.Int64>? ids,
    $core.String? name,
    $fixnum.Int64? parentId,
    IssueType? type,
    $core.String? author,
    $core.String? assigned,
    $fixnum.Int64? sprintId,
    $fixnum.Int64? projectId,
    $4.IssueStatus? status,
  }) {
    final _result = create();
    if (ids != null) {
      _result.ids.addAll(ids);
    }
    if (name != null) {
      _result.name = name;
    }
    if (parentId != null) {
      _result.parentId = parentId;
    }
    if (type != null) {
      _result.type = type;
    }
    if (author != null) {
      _result.author = author;
    }
    if (assigned != null) {
      _result.assigned = assigned;
    }
    if (sprintId != null) {
      _result.sprintId = sprintId;
    }
    if (projectId != null) {
      _result.projectId = projectId;
    }
    if (status != null) {
      _result.status = status;
    }
    return _result;
  }
  factory IssueGetRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueGetRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueGetRequest clone() => IssueGetRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueGetRequest copyWith(void Function(IssueGetRequest) updates) => super.copyWith((message) => updates(message as IssueGetRequest)) as IssueGetRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueGetRequest create() => IssueGetRequest._();
  IssueGetRequest createEmptyInstance() => create();
  static $pb.PbList<IssueGetRequest> createRepeated() => $pb.PbList<IssueGetRequest>();
  @$core.pragma('dart2js:noInline')
  static IssueGetRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueGetRequest>(create);
  static IssueGetRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$fixnum.Int64> get ids => $_getList(0);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get parentId => $_getI64(2);
  @$pb.TagNumber(3)
  set parentId($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasParentId() => $_has(2);
  @$pb.TagNumber(3)
  void clearParentId() => clearField(3);

  @$pb.TagNumber(4)
  IssueType get type => $_getN(3);
  @$pb.TagNumber(4)
  set type(IssueType v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasType() => $_has(3);
  @$pb.TagNumber(4)
  void clearType() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get author => $_getSZ(4);
  @$pb.TagNumber(5)
  set author($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasAuthor() => $_has(4);
  @$pb.TagNumber(5)
  void clearAuthor() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get assigned => $_getSZ(5);
  @$pb.TagNumber(6)
  set assigned($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasAssigned() => $_has(5);
  @$pb.TagNumber(6)
  void clearAssigned() => clearField(6);

  @$pb.TagNumber(7)
  $fixnum.Int64 get sprintId => $_getI64(6);
  @$pb.TagNumber(7)
  set sprintId($fixnum.Int64 v) { $_setInt64(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasSprintId() => $_has(6);
  @$pb.TagNumber(7)
  void clearSprintId() => clearField(7);

  @$pb.TagNumber(8)
  $fixnum.Int64 get projectId => $_getI64(7);
  @$pb.TagNumber(8)
  set projectId($fixnum.Int64 v) { $_setInt64(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasProjectId() => $_has(7);
  @$pb.TagNumber(8)
  void clearProjectId() => clearField(8);

  @$pb.TagNumber(9)
  $4.IssueStatus get status => $_getN(8);
  @$pb.TagNumber(9)
  set status($4.IssueStatus v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasStatus() => $_has(8);
  @$pb.TagNumber(9)
  void clearStatus() => clearField(9);
  @$pb.TagNumber(9)
  $4.IssueStatus ensureStatus() => $_ensure(8);
}

class IssueGetResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueGetResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..pc<Issue>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'issues', $pb.PbFieldType.PM, subBuilder: Issue.create)
    ..hasRequiredFields = false
  ;

  IssueGetResponse._() : super();
  factory IssueGetResponse({
    $core.Iterable<Issue>? issues,
  }) {
    final _result = create();
    if (issues != null) {
      _result.issues.addAll(issues);
    }
    return _result;
  }
  factory IssueGetResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueGetResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueGetResponse clone() => IssueGetResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueGetResponse copyWith(void Function(IssueGetResponse) updates) => super.copyWith((message) => updates(message as IssueGetResponse)) as IssueGetResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueGetResponse create() => IssueGetResponse._();
  IssueGetResponse createEmptyInstance() => create();
  static $pb.PbList<IssueGetResponse> createRepeated() => $pb.PbList<IssueGetResponse>();
  @$core.pragma('dart2js:noInline')
  static IssueGetResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueGetResponse>(create);
  static IssueGetResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Issue> get issues => $_getList(0);
}

class IssueInfoGetRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueInfoGetRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..p<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ids', $pb.PbFieldType.K6)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aInt64(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'parentId')
    ..e<IssueType>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: IssueType.ISSUE_TYPE_UNKNOWN, valueOf: IssueType.valueOf, enumValues: IssueType.values)
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'author')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'assigned')
    ..aInt64(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sprintId')
    ..aInt64(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'projectId')
    ..aOM<$4.IssueStatus>(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', subBuilder: $4.IssueStatus.create)
    ..hasRequiredFields = false
  ;

  IssueInfoGetRequest._() : super();
  factory IssueInfoGetRequest({
    $core.Iterable<$fixnum.Int64>? ids,
    $core.String? name,
    $fixnum.Int64? parentId,
    IssueType? type,
    $core.String? author,
    $core.String? assigned,
    $fixnum.Int64? sprintId,
    $fixnum.Int64? projectId,
    $4.IssueStatus? status,
  }) {
    final _result = create();
    if (ids != null) {
      _result.ids.addAll(ids);
    }
    if (name != null) {
      _result.name = name;
    }
    if (parentId != null) {
      _result.parentId = parentId;
    }
    if (type != null) {
      _result.type = type;
    }
    if (author != null) {
      _result.author = author;
    }
    if (assigned != null) {
      _result.assigned = assigned;
    }
    if (sprintId != null) {
      _result.sprintId = sprintId;
    }
    if (projectId != null) {
      _result.projectId = projectId;
    }
    if (status != null) {
      _result.status = status;
    }
    return _result;
  }
  factory IssueInfoGetRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueInfoGetRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueInfoGetRequest clone() => IssueInfoGetRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueInfoGetRequest copyWith(void Function(IssueInfoGetRequest) updates) => super.copyWith((message) => updates(message as IssueInfoGetRequest)) as IssueInfoGetRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueInfoGetRequest create() => IssueInfoGetRequest._();
  IssueInfoGetRequest createEmptyInstance() => create();
  static $pb.PbList<IssueInfoGetRequest> createRepeated() => $pb.PbList<IssueInfoGetRequest>();
  @$core.pragma('dart2js:noInline')
  static IssueInfoGetRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueInfoGetRequest>(create);
  static IssueInfoGetRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$fixnum.Int64> get ids => $_getList(0);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get parentId => $_getI64(2);
  @$pb.TagNumber(3)
  set parentId($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasParentId() => $_has(2);
  @$pb.TagNumber(3)
  void clearParentId() => clearField(3);

  @$pb.TagNumber(4)
  IssueType get type => $_getN(3);
  @$pb.TagNumber(4)
  set type(IssueType v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasType() => $_has(3);
  @$pb.TagNumber(4)
  void clearType() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get author => $_getSZ(4);
  @$pb.TagNumber(5)
  set author($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasAuthor() => $_has(4);
  @$pb.TagNumber(5)
  void clearAuthor() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get assigned => $_getSZ(5);
  @$pb.TagNumber(6)
  set assigned($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasAssigned() => $_has(5);
  @$pb.TagNumber(6)
  void clearAssigned() => clearField(6);

  @$pb.TagNumber(7)
  $fixnum.Int64 get sprintId => $_getI64(6);
  @$pb.TagNumber(7)
  set sprintId($fixnum.Int64 v) { $_setInt64(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasSprintId() => $_has(6);
  @$pb.TagNumber(7)
  void clearSprintId() => clearField(7);

  @$pb.TagNumber(8)
  $fixnum.Int64 get projectId => $_getI64(7);
  @$pb.TagNumber(8)
  set projectId($fixnum.Int64 v) { $_setInt64(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasProjectId() => $_has(7);
  @$pb.TagNumber(8)
  void clearProjectId() => clearField(8);

  @$pb.TagNumber(9)
  $4.IssueStatus get status => $_getN(8);
  @$pb.TagNumber(9)
  set status($4.IssueStatus v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasStatus() => $_has(8);
  @$pb.TagNumber(9)
  void clearStatus() => clearField(9);
  @$pb.TagNumber(9)
  $4.IssueStatus ensureStatus() => $_ensure(8);
}

class IssueInfoGetResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueInfoGetResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..pc<IssueInfo>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'issueInfo', $pb.PbFieldType.PM, subBuilder: IssueInfo.create)
    ..hasRequiredFields = false
  ;

  IssueInfoGetResponse._() : super();
  factory IssueInfoGetResponse({
    $core.Iterable<IssueInfo>? issueInfo,
  }) {
    final _result = create();
    if (issueInfo != null) {
      _result.issueInfo.addAll(issueInfo);
    }
    return _result;
  }
  factory IssueInfoGetResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueInfoGetResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueInfoGetResponse clone() => IssueInfoGetResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueInfoGetResponse copyWith(void Function(IssueInfoGetResponse) updates) => super.copyWith((message) => updates(message as IssueInfoGetResponse)) as IssueInfoGetResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueInfoGetResponse create() => IssueInfoGetResponse._();
  IssueInfoGetResponse createEmptyInstance() => create();
  static $pb.PbList<IssueInfoGetResponse> createRepeated() => $pb.PbList<IssueInfoGetResponse>();
  @$core.pragma('dart2js:noInline')
  static IssueInfoGetResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueInfoGetResponse>(create);
  static IssueInfoGetResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<IssueInfo> get issueInfo => $_getList(0);
}

class IssueGetByIDRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueGetByIDRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  IssueGetByIDRequest._() : super();
  factory IssueGetByIDRequest({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory IssueGetByIDRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueGetByIDRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueGetByIDRequest clone() => IssueGetByIDRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueGetByIDRequest copyWith(void Function(IssueGetByIDRequest) updates) => super.copyWith((message) => updates(message as IssueGetByIDRequest)) as IssueGetByIDRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueGetByIDRequest create() => IssueGetByIDRequest._();
  IssueGetByIDRequest createEmptyInstance() => create();
  static $pb.PbList<IssueGetByIDRequest> createRepeated() => $pb.PbList<IssueGetByIDRequest>();
  @$core.pragma('dart2js:noInline')
  static IssueGetByIDRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueGetByIDRequest>(create);
  static IssueGetByIDRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class IssueGetByIDResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueGetByIDResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aOM<Issue>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'issue', subBuilder: Issue.create)
    ..hasRequiredFields = false
  ;

  IssueGetByIDResponse._() : super();
  factory IssueGetByIDResponse({
    Issue? issue,
  }) {
    final _result = create();
    if (issue != null) {
      _result.issue = issue;
    }
    return _result;
  }
  factory IssueGetByIDResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueGetByIDResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueGetByIDResponse clone() => IssueGetByIDResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueGetByIDResponse copyWith(void Function(IssueGetByIDResponse) updates) => super.copyWith((message) => updates(message as IssueGetByIDResponse)) as IssueGetByIDResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueGetByIDResponse create() => IssueGetByIDResponse._();
  IssueGetByIDResponse createEmptyInstance() => create();
  static $pb.PbList<IssueGetByIDResponse> createRepeated() => $pb.PbList<IssueGetByIDResponse>();
  @$core.pragma('dart2js:noInline')
  static IssueGetByIDResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueGetByIDResponse>(create);
  static IssueGetByIDResponse? _defaultInstance;

  @$pb.TagNumber(1)
  Issue get issue => $_getN(0);
  @$pb.TagNumber(1)
  set issue(Issue v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasIssue() => $_has(0);
  @$pb.TagNumber(1)
  void clearIssue() => clearField(1);
  @$pb.TagNumber(1)
  Issue ensureIssue() => $_ensure(0);
}

class IssueInfoGetByIDRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueInfoGetByIDRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  IssueInfoGetByIDRequest._() : super();
  factory IssueInfoGetByIDRequest({
    $fixnum.Int64? id,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    return _result;
  }
  factory IssueInfoGetByIDRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueInfoGetByIDRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueInfoGetByIDRequest clone() => IssueInfoGetByIDRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueInfoGetByIDRequest copyWith(void Function(IssueInfoGetByIDRequest) updates) => super.copyWith((message) => updates(message as IssueInfoGetByIDRequest)) as IssueInfoGetByIDRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueInfoGetByIDRequest create() => IssueInfoGetByIDRequest._();
  IssueInfoGetByIDRequest createEmptyInstance() => create();
  static $pb.PbList<IssueInfoGetByIDRequest> createRepeated() => $pb.PbList<IssueInfoGetByIDRequest>();
  @$core.pragma('dart2js:noInline')
  static IssueInfoGetByIDRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueInfoGetByIDRequest>(create);
  static IssueInfoGetByIDRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class IssueInfoGetByIDResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueInfoGetByIDResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..aOM<IssueInfo>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'issueInfo', subBuilder: IssueInfo.create)
    ..hasRequiredFields = false
  ;

  IssueInfoGetByIDResponse._() : super();
  factory IssueInfoGetByIDResponse({
    IssueInfo? issueInfo,
  }) {
    final _result = create();
    if (issueInfo != null) {
      _result.issueInfo = issueInfo;
    }
    return _result;
  }
  factory IssueInfoGetByIDResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueInfoGetByIDResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueInfoGetByIDResponse clone() => IssueInfoGetByIDResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueInfoGetByIDResponse copyWith(void Function(IssueInfoGetByIDResponse) updates) => super.copyWith((message) => updates(message as IssueInfoGetByIDResponse)) as IssueInfoGetByIDResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueInfoGetByIDResponse create() => IssueInfoGetByIDResponse._();
  IssueInfoGetByIDResponse createEmptyInstance() => create();
  static $pb.PbList<IssueInfoGetByIDResponse> createRepeated() => $pb.PbList<IssueInfoGetByIDResponse>();
  @$core.pragma('dart2js:noInline')
  static IssueInfoGetByIDResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueInfoGetByIDResponse>(create);
  static IssueInfoGetByIDResponse? _defaultInstance;

  @$pb.TagNumber(1)
  IssueInfo get issueInfo => $_getN(0);
  @$pb.TagNumber(1)
  set issueInfo(IssueInfo v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasIssueInfo() => $_has(0);
  @$pb.TagNumber(1)
  void clearIssueInfo() => clearField(1);
  @$pb.TagNumber(1)
  IssueInfo ensureIssueInfo() => $_ensure(0);
}

class IssuePaymentGetRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssuePaymentGetRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  IssuePaymentGetRequest._() : super();
  factory IssuePaymentGetRequest() => create();
  factory IssuePaymentGetRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssuePaymentGetRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssuePaymentGetRequest clone() => IssuePaymentGetRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssuePaymentGetRequest copyWith(void Function(IssuePaymentGetRequest) updates) => super.copyWith((message) => updates(message as IssuePaymentGetRequest)) as IssuePaymentGetRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssuePaymentGetRequest create() => IssuePaymentGetRequest._();
  IssuePaymentGetRequest createEmptyInstance() => create();
  static $pb.PbList<IssuePaymentGetRequest> createRepeated() => $pb.PbList<IssuePaymentGetRequest>();
  @$core.pragma('dart2js:noInline')
  static IssuePaymentGetRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssuePaymentGetRequest>(create);
  static IssuePaymentGetRequest? _defaultInstance;
}

class IssuePaymentGetResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssuePaymentGetResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.tasks.issue'), createEmptyInstance: create)
    ..m<$core.String, $core.double>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'userToPayment', entryClassName: 'IssuePaymentGetResponse.UserToPaymentEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OD, packageName: const $pb.PackageName('github.constantine27k.crnt_data_manager.api.tasks.issue'))
    ..hasRequiredFields = false
  ;

  IssuePaymentGetResponse._() : super();
  factory IssuePaymentGetResponse({
    $core.Map<$core.String, $core.double>? userToPayment,
  }) {
    final _result = create();
    if (userToPayment != null) {
      _result.userToPayment.addAll(userToPayment);
    }
    return _result;
  }
  factory IssuePaymentGetResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssuePaymentGetResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssuePaymentGetResponse clone() => IssuePaymentGetResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssuePaymentGetResponse copyWith(void Function(IssuePaymentGetResponse) updates) => super.copyWith((message) => updates(message as IssuePaymentGetResponse)) as IssuePaymentGetResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssuePaymentGetResponse create() => IssuePaymentGetResponse._();
  IssuePaymentGetResponse createEmptyInstance() => create();
  static $pb.PbList<IssuePaymentGetResponse> createRepeated() => $pb.PbList<IssuePaymentGetResponse>();
  @$core.pragma('dart2js:noInline')
  static IssuePaymentGetResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssuePaymentGetResponse>(create);
  static IssuePaymentGetResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.Map<$core.String, $core.double> get userToPayment => $_getMap(0);
}

class IssueRegistryApi {
  $pb.RpcClient _client;
  IssueRegistryApi(this._client);

  $async.Future<IssueCreateResponse> createIssue($pb.ClientContext? ctx, IssueCreateRequest request) {
    var emptyResponse = IssueCreateResponse();
    return _client.invoke<IssueCreateResponse>(ctx, 'IssueRegistry', 'CreateIssue', request, emptyResponse);
  }
  $async.Future<IssueCreateSubtaskResponse> createSubtask($pb.ClientContext? ctx, IssueCreateSubtaskRequest request) {
    var emptyResponse = IssueCreateSubtaskResponse();
    return _client.invoke<IssueCreateSubtaskResponse>(ctx, 'IssueRegistry', 'CreateSubtask', request, emptyResponse);
  }
  $async.Future<IssueAddCommentResponse> addComment($pb.ClientContext? ctx, IssueAddCommentRequest request) {
    var emptyResponse = IssueAddCommentResponse();
    return _client.invoke<IssueAddCommentResponse>(ctx, 'IssueRegistry', 'AddComment', request, emptyResponse);
  }
  $async.Future<IssueUpdateResponse> updateIssue($pb.ClientContext? ctx, IssueUpdateRequest request) {
    var emptyResponse = IssueUpdateResponse();
    return _client.invoke<IssueUpdateResponse>(ctx, 'IssueRegistry', 'UpdateIssue', request, emptyResponse);
  }
  $async.Future<IssueGetResponse> getIssues($pb.ClientContext? ctx, IssueGetRequest request) {
    var emptyResponse = IssueGetResponse();
    return _client.invoke<IssueGetResponse>(ctx, 'IssueRegistry', 'GetIssues', request, emptyResponse);
  }
  $async.Future<IssueGetByIDResponse> getIssueByID($pb.ClientContext? ctx, IssueGetByIDRequest request) {
    var emptyResponse = IssueGetByIDResponse();
    return _client.invoke<IssueGetByIDResponse>(ctx, 'IssueRegistry', 'GetIssueByID', request, emptyResponse);
  }
  $async.Future<IssueInfoGetResponse> getIssueInfo($pb.ClientContext? ctx, IssueInfoGetRequest request) {
    var emptyResponse = IssueInfoGetResponse();
    return _client.invoke<IssueInfoGetResponse>(ctx, 'IssueRegistry', 'GetIssueInfo', request, emptyResponse);
  }
  $async.Future<IssueInfoGetByIDResponse> getIssueInfoByID($pb.ClientContext? ctx, IssueInfoGetByIDRequest request) {
    var emptyResponse = IssueInfoGetByIDResponse();
    return _client.invoke<IssueInfoGetByIDResponse>(ctx, 'IssueRegistry', 'GetIssueInfoByID', request, emptyResponse);
  }
  $async.Future<IssuePaymentGetResponse> getUserPayment($pb.ClientContext? ctx, IssuePaymentGetRequest request) {
    var emptyResponse = IssuePaymentGetResponse();
    return _client.invoke<IssuePaymentGetResponse>(ctx, 'IssueRegistry', 'GetUserPayment', request, emptyResponse);
  }
}

