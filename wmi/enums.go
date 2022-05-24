package wmi

// enum WbemImpersonationLevelEnum
var WbemImpersonationLevelEnum = struct {
	WbemImpersonationLevelAnonymous int32
	WbemImpersonationLevelIdentify int32
	WbemImpersonationLevelImpersonate int32
	WbemImpersonationLevelDelegate int32
}{
	WbemImpersonationLevelAnonymous: 1,
	WbemImpersonationLevelIdentify: 2,
	WbemImpersonationLevelImpersonate: 3,
	WbemImpersonationLevelDelegate: 4,
}

// enum WbemAuthenticationLevelEnum
var WbemAuthenticationLevelEnum = struct {
	WbemAuthenticationLevelDefault int32
	WbemAuthenticationLevelNone int32
	WbemAuthenticationLevelConnect int32
	WbemAuthenticationLevelCall int32
	WbemAuthenticationLevelPkt int32
	WbemAuthenticationLevelPktIntegrity int32
	WbemAuthenticationLevelPktPrivacy int32
}{
	WbemAuthenticationLevelDefault: 0,
	WbemAuthenticationLevelNone: 1,
	WbemAuthenticationLevelConnect: 2,
	WbemAuthenticationLevelCall: 3,
	WbemAuthenticationLevelPkt: 4,
	WbemAuthenticationLevelPktIntegrity: 5,
	WbemAuthenticationLevelPktPrivacy: 6,
}

// enum WbemPrivilegeEnum
var WbemPrivilegeEnum = struct {
	WbemPrivilegeCreateToken int32
	WbemPrivilegePrimaryToken int32
	WbemPrivilegeLockMemory int32
	WbemPrivilegeIncreaseQuota int32
	WbemPrivilegeMachineAccount int32
	WbemPrivilegeTcb int32
	WbemPrivilegeSecurity int32
	WbemPrivilegeTakeOwnership int32
	WbemPrivilegeLoadDriver int32
	WbemPrivilegeSystemProfile int32
	WbemPrivilegeSystemtime int32
	WbemPrivilegeProfileSingleProcess int32
	WbemPrivilegeIncreaseBasePriority int32
	WbemPrivilegeCreatePagefile int32
	WbemPrivilegeCreatePermanent int32
	WbemPrivilegeBackup int32
	WbemPrivilegeRestore int32
	WbemPrivilegeShutdown int32
	WbemPrivilegeDebug int32
	WbemPrivilegeAudit int32
	WbemPrivilegeSystemEnvironment int32
	WbemPrivilegeChangeNotify int32
	WbemPrivilegeRemoteShutdown int32
	WbemPrivilegeUndock int32
	WbemPrivilegeSyncAgent int32
	WbemPrivilegeEnableDelegation int32
	WbemPrivilegeManageVolume int32
}{
	WbemPrivilegeCreateToken: 1,
	WbemPrivilegePrimaryToken: 2,
	WbemPrivilegeLockMemory: 3,
	WbemPrivilegeIncreaseQuota: 4,
	WbemPrivilegeMachineAccount: 5,
	WbemPrivilegeTcb: 6,
	WbemPrivilegeSecurity: 7,
	WbemPrivilegeTakeOwnership: 8,
	WbemPrivilegeLoadDriver: 9,
	WbemPrivilegeSystemProfile: 10,
	WbemPrivilegeSystemtime: 11,
	WbemPrivilegeProfileSingleProcess: 12,
	WbemPrivilegeIncreaseBasePriority: 13,
	WbemPrivilegeCreatePagefile: 14,
	WbemPrivilegeCreatePermanent: 15,
	WbemPrivilegeBackup: 16,
	WbemPrivilegeRestore: 17,
	WbemPrivilegeShutdown: 18,
	WbemPrivilegeDebug: 19,
	WbemPrivilegeAudit: 20,
	WbemPrivilegeSystemEnvironment: 21,
	WbemPrivilegeChangeNotify: 22,
	WbemPrivilegeRemoteShutdown: 23,
	WbemPrivilegeUndock: 24,
	WbemPrivilegeSyncAgent: 25,
	WbemPrivilegeEnableDelegation: 26,
	WbemPrivilegeManageVolume: 27,
}

// enum WbemCimtypeEnum
var WbemCimtypeEnum = struct {
	WbemCimtypeSint8 int32
	WbemCimtypeUint8 int32
	WbemCimtypeSint16 int32
	WbemCimtypeUint16 int32
	WbemCimtypeSint32 int32
	WbemCimtypeUint32 int32
	WbemCimtypeSint64 int32
	WbemCimtypeUint64 int32
	WbemCimtypeReal32 int32
	WbemCimtypeReal64 int32
	WbemCimtypeBoolean int32
	WbemCimtypeString int32
	WbemCimtypeDatetime int32
	WbemCimtypeReference int32
	WbemCimtypeChar16 int32
	WbemCimtypeObject int32
}{
	WbemCimtypeSint8: 16,
	WbemCimtypeUint8: 17,
	WbemCimtypeSint16: 2,
	WbemCimtypeUint16: 18,
	WbemCimtypeSint32: 3,
	WbemCimtypeUint32: 19,
	WbemCimtypeSint64: 20,
	WbemCimtypeUint64: 21,
	WbemCimtypeReal32: 4,
	WbemCimtypeReal64: 5,
	WbemCimtypeBoolean: 11,
	WbemCimtypeString: 8,
	WbemCimtypeDatetime: 101,
	WbemCimtypeReference: 102,
	WbemCimtypeChar16: 103,
	WbemCimtypeObject: 13,
}

// enum WbemErrorEnum
var WbemErrorEnum = struct {
	WbemNoErr int32
	WbemErrFailed int32
	WbemErrNotFound int32
	WbemErrAccessDenied int32
	WbemErrProviderFailure int32
	WbemErrTypeMismatch int32
	WbemErrOutOfMemory int32
	WbemErrInvalidContext int32
	WbemErrInvalidParameter int32
	WbemErrNotAvailable int32
	WbemErrCriticalError int32
	WbemErrInvalidStream int32
	WbemErrNotSupported int32
	WbemErrInvalidSuperclass int32
	WbemErrInvalidNamespace int32
	WbemErrInvalidObject int32
	WbemErrInvalidClass int32
	WbemErrProviderNotFound int32
	WbemErrInvalidProviderRegistration int32
	WbemErrProviderLoadFailure int32
	WbemErrInitializationFailure int32
	WbemErrTransportFailure int32
	WbemErrInvalidOperation int32
	WbemErrInvalidQuery int32
	WbemErrInvalidQueryType int32
	WbemErrAlreadyExists int32
	WbemErrOverrideNotAllowed int32
	WbemErrPropagatedQualifier int32
	WbemErrPropagatedProperty int32
	WbemErrUnexpected int32
	WbemErrIllegalOperation int32
	WbemErrCannotBeKey int32
	WbemErrIncompleteClass int32
	WbemErrInvalidSyntax int32
	WbemErrNondecoratedObject int32
	WbemErrReadOnly int32
	WbemErrProviderNotCapable int32
	WbemErrClassHasChildren int32
	WbemErrClassHasInstances int32
	WbemErrQueryNotImplemented int32
	WbemErrIllegalNull int32
	WbemErrInvalidQualifierType int32
	WbemErrInvalidPropertyType int32
	WbemErrValueOutOfRange int32
	WbemErrCannotBeSingleton int32
	WbemErrInvalidCimType int32
	WbemErrInvalidMethod int32
	WbemErrInvalidMethodParameters int32
	WbemErrSystemProperty int32
	WbemErrInvalidProperty int32
	WbemErrCallCancelled int32
	WbemErrShuttingDown int32
	WbemErrPropagatedMethod int32
	WbemErrUnsupportedParameter int32
	WbemErrMissingParameter int32
	WbemErrInvalidParameterId int32
	WbemErrNonConsecutiveParameterIds int32
	WbemErrParameterIdOnRetval int32
	WbemErrInvalidObjectPath int32
	WbemErrOutOfDiskSpace int32
	WbemErrBufferTooSmall int32
	WbemErrUnsupportedPutExtension int32
	WbemErrUnknownObjectType int32
	WbemErrUnknownPacketType int32
	WbemErrMarshalVersionMismatch int32
	WbemErrMarshalInvalidSignature int32
	WbemErrInvalidQualifier int32
	WbemErrInvalidDuplicateParameter int32
	WbemErrTooMuchData int32
	WbemErrServerTooBusy int32
	WbemErrInvalidFlavor int32
	WbemErrCircularReference int32
	WbemErrUnsupportedClassUpdate int32
	WbemErrCannotChangeKeyInheritance int32
	WbemErrCannotChangeIndexInheritance int32
	WbemErrTooManyProperties int32
	WbemErrUpdateTypeMismatch int32
	WbemErrUpdateOverrideNotAllowed int32
	WbemErrUpdatePropagatedMethod int32
	WbemErrMethodNotImplemented int32
	WbemErrMethodDisabled int32
	WbemErrRefresherBusy int32
	WbemErrUnparsableQuery int32
	WbemErrNotEventClass int32
	WbemErrMissingGroupWithin int32
	WbemErrMissingAggregationList int32
	WbemErrPropertyNotAnObject int32
	WbemErrAggregatingByObject int32
	WbemErrUninterpretableProviderQuery int32
	WbemErrBackupRestoreWinmgmtRunning int32
	WbemErrQueueOverflow int32
	WbemErrPrivilegeNotHeld int32
	WbemErrInvalidOperator int32
	WbemErrLocalCredentials int32
	WbemErrCannotBeAbstract int32
	WbemErrAmendedObject int32
	WbemErrClientTooSlow int32
	WbemErrNullSecurityDescriptor int32
	WbemErrTimeout int32
	WbemErrInvalidAssociation int32
	WbemErrAmbiguousOperation int32
	WbemErrQuotaViolation int32
	WbemErrTransactionConflict int32
	WbemErrForcedRollback int32
	WbemErrUnsupportedLocale int32
	WbemErrHandleOutOfDate int32
	WbemErrConnectionFailed int32
	WbemErrInvalidHandleRequest int32
	WbemErrPropertyNameTooWide int32
	WbemErrClassNameTooWide int32
	WbemErrMethodNameTooWide int32
	WbemErrQualifierNameTooWide int32
	WbemErrRerunCommand int32
	WbemErrDatabaseVerMismatch int32
	WbemErrVetoPut int32
	WbemErrVetoDelete int32
	WbemErrInvalidLocale int32
	WbemErrProviderSuspended int32
	WbemErrSynchronizationRequired int32
	WbemErrNoSchema int32
	WbemErrProviderAlreadyRegistered int32
	WbemErrProviderNotRegistered int32
	WbemErrFatalTransportError int32
	WbemErrEncryptedConnectionRequired int32
	WbemErrRegistrationTooBroad int32
	WbemErrRegistrationTooPrecise int32
	WbemErrTimedout int32
	WbemErrResetToDefault int32
}{
	WbemNoErr: 0,
	WbemErrFailed: -2147217407,
	WbemErrNotFound: -2147217406,
	WbemErrAccessDenied: -2147217405,
	WbemErrProviderFailure: -2147217404,
	WbemErrTypeMismatch: -2147217403,
	WbemErrOutOfMemory: -2147217402,
	WbemErrInvalidContext: -2147217401,
	WbemErrInvalidParameter: -2147217400,
	WbemErrNotAvailable: -2147217399,
	WbemErrCriticalError: -2147217398,
	WbemErrInvalidStream: -2147217397,
	WbemErrNotSupported: -2147217396,
	WbemErrInvalidSuperclass: -2147217395,
	WbemErrInvalidNamespace: -2147217394,
	WbemErrInvalidObject: -2147217393,
	WbemErrInvalidClass: -2147217392,
	WbemErrProviderNotFound: -2147217391,
	WbemErrInvalidProviderRegistration: -2147217390,
	WbemErrProviderLoadFailure: -2147217389,
	WbemErrInitializationFailure: -2147217388,
	WbemErrTransportFailure: -2147217387,
	WbemErrInvalidOperation: -2147217386,
	WbemErrInvalidQuery: -2147217385,
	WbemErrInvalidQueryType: -2147217384,
	WbemErrAlreadyExists: -2147217383,
	WbemErrOverrideNotAllowed: -2147217382,
	WbemErrPropagatedQualifier: -2147217381,
	WbemErrPropagatedProperty: -2147217380,
	WbemErrUnexpected: -2147217379,
	WbemErrIllegalOperation: -2147217378,
	WbemErrCannotBeKey: -2147217377,
	WbemErrIncompleteClass: -2147217376,
	WbemErrInvalidSyntax: -2147217375,
	WbemErrNondecoratedObject: -2147217374,
	WbemErrReadOnly: -2147217373,
	WbemErrProviderNotCapable: -2147217372,
	WbemErrClassHasChildren: -2147217371,
	WbemErrClassHasInstances: -2147217370,
	WbemErrQueryNotImplemented: -2147217369,
	WbemErrIllegalNull: -2147217368,
	WbemErrInvalidQualifierType: -2147217367,
	WbemErrInvalidPropertyType: -2147217366,
	WbemErrValueOutOfRange: -2147217365,
	WbemErrCannotBeSingleton: -2147217364,
	WbemErrInvalidCimType: -2147217363,
	WbemErrInvalidMethod: -2147217362,
	WbemErrInvalidMethodParameters: -2147217361,
	WbemErrSystemProperty: -2147217360,
	WbemErrInvalidProperty: -2147217359,
	WbemErrCallCancelled: -2147217358,
	WbemErrShuttingDown: -2147217357,
	WbemErrPropagatedMethod: -2147217356,
	WbemErrUnsupportedParameter: -2147217355,
	WbemErrMissingParameter: -2147217354,
	WbemErrInvalidParameterId: -2147217353,
	WbemErrNonConsecutiveParameterIds: -2147217352,
	WbemErrParameterIdOnRetval: -2147217351,
	WbemErrInvalidObjectPath: -2147217350,
	WbemErrOutOfDiskSpace: -2147217349,
	WbemErrBufferTooSmall: -2147217348,
	WbemErrUnsupportedPutExtension: -2147217347,
	WbemErrUnknownObjectType: -2147217346,
	WbemErrUnknownPacketType: -2147217345,
	WbemErrMarshalVersionMismatch: -2147217344,
	WbemErrMarshalInvalidSignature: -2147217343,
	WbemErrInvalidQualifier: -2147217342,
	WbemErrInvalidDuplicateParameter: -2147217341,
	WbemErrTooMuchData: -2147217340,
	WbemErrServerTooBusy: -2147217339,
	WbemErrInvalidFlavor: -2147217338,
	WbemErrCircularReference: -2147217337,
	WbemErrUnsupportedClassUpdate: -2147217336,
	WbemErrCannotChangeKeyInheritance: -2147217335,
	WbemErrCannotChangeIndexInheritance: -2147217328,
	WbemErrTooManyProperties: -2147217327,
	WbemErrUpdateTypeMismatch: -2147217326,
	WbemErrUpdateOverrideNotAllowed: -2147217325,
	WbemErrUpdatePropagatedMethod: -2147217324,
	WbemErrMethodNotImplemented: -2147217323,
	WbemErrMethodDisabled: -2147217322,
	WbemErrRefresherBusy: -2147217321,
	WbemErrUnparsableQuery: -2147217320,
	WbemErrNotEventClass: -2147217319,
	WbemErrMissingGroupWithin: -2147217318,
	WbemErrMissingAggregationList: -2147217317,
	WbemErrPropertyNotAnObject: -2147217316,
	WbemErrAggregatingByObject: -2147217315,
	WbemErrUninterpretableProviderQuery: -2147217313,
	WbemErrBackupRestoreWinmgmtRunning: -2147217312,
	WbemErrQueueOverflow: -2147217311,
	WbemErrPrivilegeNotHeld: -2147217310,
	WbemErrInvalidOperator: -2147217309,
	WbemErrLocalCredentials: -2147217308,
	WbemErrCannotBeAbstract: -2147217307,
	WbemErrAmendedObject: -2147217306,
	WbemErrClientTooSlow: -2147217305,
	WbemErrNullSecurityDescriptor: -2147217304,
	WbemErrTimeout: -2147217303,
	WbemErrInvalidAssociation: -2147217302,
	WbemErrAmbiguousOperation: -2147217301,
	WbemErrQuotaViolation: -2147217300,
	WbemErrTransactionConflict: -2147217299,
	WbemErrForcedRollback: -2147217298,
	WbemErrUnsupportedLocale: -2147217297,
	WbemErrHandleOutOfDate: -2147217296,
	WbemErrConnectionFailed: -2147217295,
	WbemErrInvalidHandleRequest: -2147217294,
	WbemErrPropertyNameTooWide: -2147217293,
	WbemErrClassNameTooWide: -2147217292,
	WbemErrMethodNameTooWide: -2147217291,
	WbemErrQualifierNameTooWide: -2147217290,
	WbemErrRerunCommand: -2147217289,
	WbemErrDatabaseVerMismatch: -2147217288,
	WbemErrVetoPut: -2147217287,
	WbemErrVetoDelete: -2147217286,
	WbemErrInvalidLocale: -2147217280,
	WbemErrProviderSuspended: -2147217279,
	WbemErrSynchronizationRequired: -2147217278,
	WbemErrNoSchema: -2147217277,
	WbemErrProviderAlreadyRegistered: -2147217276,
	WbemErrProviderNotRegistered: -2147217275,
	WbemErrFatalTransportError: -2147217274,
	WbemErrEncryptedConnectionRequired: -2147217273,
	WbemErrRegistrationTooBroad: -2147213311,
	WbemErrRegistrationTooPrecise: -2147213310,
	WbemErrTimedout: -2147209215,
	WbemErrResetToDefault: -2147209214,
}

// enum WbemObjectTextFormatEnum
var WbemObjectTextFormatEnum = struct {
	WbemObjectTextFormatCIMDTD20 int32
	WbemObjectTextFormatWMIDTD20 int32
}{
	WbemObjectTextFormatCIMDTD20: 1,
	WbemObjectTextFormatWMIDTD20: 2,
}

// enum WbemChangeFlagEnum
var WbemChangeFlagEnum = struct {
	WbemChangeFlagCreateOrUpdate int32
	WbemChangeFlagUpdateOnly int32
	WbemChangeFlagCreateOnly int32
	WbemChangeFlagUpdateCompatible int32
	WbemChangeFlagUpdateSafeMode int32
	WbemChangeFlagUpdateForceMode int32
	WbemChangeFlagStrongValidation int32
	WbemChangeFlagAdvisory int32
}{
	WbemChangeFlagCreateOrUpdate: 0,
	WbemChangeFlagUpdateOnly: 1,
	WbemChangeFlagCreateOnly: 2,
	WbemChangeFlagUpdateCompatible: 0,
	WbemChangeFlagUpdateSafeMode: 32,
	WbemChangeFlagUpdateForceMode: 64,
	WbemChangeFlagStrongValidation: 128,
	WbemChangeFlagAdvisory: 65536,
}

// enum WbemFlagEnum
var WbemFlagEnum = struct {
	WbemFlagReturnImmediately int32
	WbemFlagReturnWhenComplete int32
	WbemFlagBidirectional int32
	WbemFlagForwardOnly int32
	WbemFlagNoErrorObject int32
	WbemFlagReturnErrorObject int32
	WbemFlagSendStatus int32
	WbemFlagDontSendStatus int32
	WbemFlagEnsureLocatable int32
	WbemFlagDirectRead int32
	WbemFlagSendOnlySelected int32
	WbemFlagUseAmendedQualifiers int32
	WbemFlagGetDefault int32
	WbemFlagSpawnInstance int32
	WbemFlagUseCurrentTime int32
}{
	WbemFlagReturnImmediately: 16,
	WbemFlagReturnWhenComplete: 0,
	WbemFlagBidirectional: 0,
	WbemFlagForwardOnly: 32,
	WbemFlagNoErrorObject: 64,
	WbemFlagReturnErrorObject: 0,
	WbemFlagSendStatus: 128,
	WbemFlagDontSendStatus: 0,
	WbemFlagEnsureLocatable: 256,
	WbemFlagDirectRead: 512,
	WbemFlagSendOnlySelected: 0,
	WbemFlagUseAmendedQualifiers: 131072,
	WbemFlagGetDefault: 0,
	WbemFlagSpawnInstance: 1,
	WbemFlagUseCurrentTime: 1,
}

// enum WbemQueryFlagEnum
var WbemQueryFlagEnum = struct {
	WbemQueryFlagDeep int32
	WbemQueryFlagShallow int32
	WbemQueryFlagPrototype int32
}{
	WbemQueryFlagDeep: 0,
	WbemQueryFlagShallow: 1,
	WbemQueryFlagPrototype: 2,
}

// enum WbemTextFlagEnum
var WbemTextFlagEnum = struct {
	WbemTextFlagNoFlavors int32
}{
	WbemTextFlagNoFlavors: 1,
}

// enum WbemTimeout
var WbemTimeout = struct {
	WbemTimeoutInfinite int32
}{
	WbemTimeoutInfinite: -1,
}

// enum WbemComparisonFlagEnum
var WbemComparisonFlagEnum = struct {
	WbemComparisonFlagIncludeAll int32
	WbemComparisonFlagIgnoreQualifiers int32
	WbemComparisonFlagIgnoreObjectSource int32
	WbemComparisonFlagIgnoreDefaultValues int32
	WbemComparisonFlagIgnoreClass int32
	WbemComparisonFlagIgnoreCase int32
	WbemComparisonFlagIgnoreFlavor int32
}{
	WbemComparisonFlagIncludeAll: 0,
	WbemComparisonFlagIgnoreQualifiers: 1,
	WbemComparisonFlagIgnoreObjectSource: 2,
	WbemComparisonFlagIgnoreDefaultValues: 4,
	WbemComparisonFlagIgnoreClass: 8,
	WbemComparisonFlagIgnoreCase: 16,
	WbemComparisonFlagIgnoreFlavor: 32,
}

// enum WbemConnectOptionsEnum
var WbemConnectOptionsEnum = struct {
	WbemConnectFlagUseMaxWait int32
}{
	WbemConnectFlagUseMaxWait: 128,
}

