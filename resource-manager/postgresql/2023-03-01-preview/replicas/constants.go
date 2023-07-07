package replicas

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveDirectoryAuthEnum string

const (
	ActiveDirectoryAuthEnumDisabled ActiveDirectoryAuthEnum = "Disabled"
	ActiveDirectoryAuthEnumEnabled  ActiveDirectoryAuthEnum = "Enabled"
)

func PossibleValuesForActiveDirectoryAuthEnum() []string {
	return []string{
		string(ActiveDirectoryAuthEnumDisabled),
		string(ActiveDirectoryAuthEnumEnabled),
	}
}

func parseActiveDirectoryAuthEnum(input string) (*ActiveDirectoryAuthEnum, error) {
	vals := map[string]ActiveDirectoryAuthEnum{
		"disabled": ActiveDirectoryAuthEnumDisabled,
		"enabled":  ActiveDirectoryAuthEnumEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActiveDirectoryAuthEnum(input)
	return &out, nil
}

type ArmServerKeyType string

const (
	ArmServerKeyTypeAzureKeyVault ArmServerKeyType = "AzureKeyVault"
	ArmServerKeyTypeSystemManaged ArmServerKeyType = "SystemManaged"
)

func PossibleValuesForArmServerKeyType() []string {
	return []string{
		string(ArmServerKeyTypeAzureKeyVault),
		string(ArmServerKeyTypeSystemManaged),
	}
}

func parseArmServerKeyType(input string) (*ArmServerKeyType, error) {
	vals := map[string]ArmServerKeyType{
		"azurekeyvault": ArmServerKeyTypeAzureKeyVault,
		"systemmanaged": ArmServerKeyTypeSystemManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ArmServerKeyType(input)
	return &out, nil
}

type AzureManagedDiskPerformanceTiers string

const (
	AzureManagedDiskPerformanceTiersPEightZero AzureManagedDiskPerformanceTiers = "P80"
	AzureManagedDiskPerformanceTiersPFiveZero  AzureManagedDiskPerformanceTiers = "P50"
	AzureManagedDiskPerformanceTiersPFour      AzureManagedDiskPerformanceTiers = "P4"
	AzureManagedDiskPerformanceTiersPFourZero  AzureManagedDiskPerformanceTiers = "P40"
	AzureManagedDiskPerformanceTiersPOne       AzureManagedDiskPerformanceTiers = "P1"
	AzureManagedDiskPerformanceTiersPOneFive   AzureManagedDiskPerformanceTiers = "P15"
	AzureManagedDiskPerformanceTiersPOneZero   AzureManagedDiskPerformanceTiers = "P10"
	AzureManagedDiskPerformanceTiersPSevenZero AzureManagedDiskPerformanceTiers = "P70"
	AzureManagedDiskPerformanceTiersPSix       AzureManagedDiskPerformanceTiers = "P6"
	AzureManagedDiskPerformanceTiersPSixZero   AzureManagedDiskPerformanceTiers = "P60"
	AzureManagedDiskPerformanceTiersPThree     AzureManagedDiskPerformanceTiers = "P3"
	AzureManagedDiskPerformanceTiersPThreeZero AzureManagedDiskPerformanceTiers = "P30"
	AzureManagedDiskPerformanceTiersPTwo       AzureManagedDiskPerformanceTiers = "P2"
	AzureManagedDiskPerformanceTiersPTwoZero   AzureManagedDiskPerformanceTiers = "P20"
)

func PossibleValuesForAzureManagedDiskPerformanceTiers() []string {
	return []string{
		string(AzureManagedDiskPerformanceTiersPEightZero),
		string(AzureManagedDiskPerformanceTiersPFiveZero),
		string(AzureManagedDiskPerformanceTiersPFour),
		string(AzureManagedDiskPerformanceTiersPFourZero),
		string(AzureManagedDiskPerformanceTiersPOne),
		string(AzureManagedDiskPerformanceTiersPOneFive),
		string(AzureManagedDiskPerformanceTiersPOneZero),
		string(AzureManagedDiskPerformanceTiersPSevenZero),
		string(AzureManagedDiskPerformanceTiersPSix),
		string(AzureManagedDiskPerformanceTiersPSixZero),
		string(AzureManagedDiskPerformanceTiersPThree),
		string(AzureManagedDiskPerformanceTiersPThreeZero),
		string(AzureManagedDiskPerformanceTiersPTwo),
		string(AzureManagedDiskPerformanceTiersPTwoZero),
	}
}

func parseAzureManagedDiskPerformanceTiers(input string) (*AzureManagedDiskPerformanceTiers, error) {
	vals := map[string]AzureManagedDiskPerformanceTiers{
		"p80": AzureManagedDiskPerformanceTiersPEightZero,
		"p50": AzureManagedDiskPerformanceTiersPFiveZero,
		"p4":  AzureManagedDiskPerformanceTiersPFour,
		"p40": AzureManagedDiskPerformanceTiersPFourZero,
		"p1":  AzureManagedDiskPerformanceTiersPOne,
		"p15": AzureManagedDiskPerformanceTiersPOneFive,
		"p10": AzureManagedDiskPerformanceTiersPOneZero,
		"p70": AzureManagedDiskPerformanceTiersPSevenZero,
		"p6":  AzureManagedDiskPerformanceTiersPSix,
		"p60": AzureManagedDiskPerformanceTiersPSixZero,
		"p3":  AzureManagedDiskPerformanceTiersPThree,
		"p30": AzureManagedDiskPerformanceTiersPThreeZero,
		"p2":  AzureManagedDiskPerformanceTiersPTwo,
		"p20": AzureManagedDiskPerformanceTiersPTwoZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureManagedDiskPerformanceTiers(input)
	return &out, nil
}

type CreateMode string

const (
	CreateModeCreate             CreateMode = "Create"
	CreateModeDefault            CreateMode = "Default"
	CreateModeGeoRestore         CreateMode = "GeoRestore"
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	CreateModeReplica            CreateMode = "Replica"
	CreateModeReviveDropped      CreateMode = "ReviveDropped"
	CreateModeUpdate             CreateMode = "Update"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeCreate),
		string(CreateModeDefault),
		string(CreateModeGeoRestore),
		string(CreateModePointInTimeRestore),
		string(CreateModeReplica),
		string(CreateModeReviveDropped),
		string(CreateModeUpdate),
	}
}

func parseCreateMode(input string) (*CreateMode, error) {
	vals := map[string]CreateMode{
		"create":             CreateModeCreate,
		"default":            CreateModeDefault,
		"georestore":         CreateModeGeoRestore,
		"pointintimerestore": CreateModePointInTimeRestore,
		"replica":            CreateModeReplica,
		"revivedropped":      CreateModeReviveDropped,
		"update":             CreateModeUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMode(input)
	return &out, nil
}

type GeoRedundantBackupEnum string

const (
	GeoRedundantBackupEnumDisabled GeoRedundantBackupEnum = "Disabled"
	GeoRedundantBackupEnumEnabled  GeoRedundantBackupEnum = "Enabled"
)

func PossibleValuesForGeoRedundantBackupEnum() []string {
	return []string{
		string(GeoRedundantBackupEnumDisabled),
		string(GeoRedundantBackupEnumEnabled),
	}
}

func parseGeoRedundantBackupEnum(input string) (*GeoRedundantBackupEnum, error) {
	vals := map[string]GeoRedundantBackupEnum{
		"disabled": GeoRedundantBackupEnumDisabled,
		"enabled":  GeoRedundantBackupEnumEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GeoRedundantBackupEnum(input)
	return &out, nil
}

type HighAvailabilityMode string

const (
	HighAvailabilityModeDisabled      HighAvailabilityMode = "Disabled"
	HighAvailabilityModeSameZone      HighAvailabilityMode = "SameZone"
	HighAvailabilityModeZoneRedundant HighAvailabilityMode = "ZoneRedundant"
)

func PossibleValuesForHighAvailabilityMode() []string {
	return []string{
		string(HighAvailabilityModeDisabled),
		string(HighAvailabilityModeSameZone),
		string(HighAvailabilityModeZoneRedundant),
	}
}

func parseHighAvailabilityMode(input string) (*HighAvailabilityMode, error) {
	vals := map[string]HighAvailabilityMode{
		"disabled":      HighAvailabilityModeDisabled,
		"samezone":      HighAvailabilityModeSameZone,
		"zoneredundant": HighAvailabilityModeZoneRedundant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HighAvailabilityMode(input)
	return &out, nil
}

type KeyStatusEnum string

const (
	KeyStatusEnumInvalid KeyStatusEnum = "Invalid"
	KeyStatusEnumValid   KeyStatusEnum = "Valid"
)

func PossibleValuesForKeyStatusEnum() []string {
	return []string{
		string(KeyStatusEnumInvalid),
		string(KeyStatusEnumValid),
	}
}

func parseKeyStatusEnum(input string) (*KeyStatusEnum, error) {
	vals := map[string]KeyStatusEnum{
		"invalid": KeyStatusEnumInvalid,
		"valid":   KeyStatusEnumValid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyStatusEnum(input)
	return &out, nil
}

type PasswordAuthEnum string

const (
	PasswordAuthEnumDisabled PasswordAuthEnum = "Disabled"
	PasswordAuthEnumEnabled  PasswordAuthEnum = "Enabled"
)

func PossibleValuesForPasswordAuthEnum() []string {
	return []string{
		string(PasswordAuthEnumDisabled),
		string(PasswordAuthEnumEnabled),
	}
}

func parsePasswordAuthEnum(input string) (*PasswordAuthEnum, error) {
	vals := map[string]PasswordAuthEnum{
		"disabled": PasswordAuthEnumDisabled,
		"enabled":  PasswordAuthEnumEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PasswordAuthEnum(input)
	return &out, nil
}

type ReplicationRole string

const (
	ReplicationRoleAsyncReplica    ReplicationRole = "AsyncReplica"
	ReplicationRoleGeoAsyncReplica ReplicationRole = "GeoAsyncReplica"
	ReplicationRoleNone            ReplicationRole = "None"
	ReplicationRolePrimary         ReplicationRole = "Primary"
)

func PossibleValuesForReplicationRole() []string {
	return []string{
		string(ReplicationRoleAsyncReplica),
		string(ReplicationRoleGeoAsyncReplica),
		string(ReplicationRoleNone),
		string(ReplicationRolePrimary),
	}
}

func parseReplicationRole(input string) (*ReplicationRole, error) {
	vals := map[string]ReplicationRole{
		"asyncreplica":    ReplicationRoleAsyncReplica,
		"geoasyncreplica": ReplicationRoleGeoAsyncReplica,
		"none":            ReplicationRoleNone,
		"primary":         ReplicationRolePrimary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicationRole(input)
	return &out, nil
}

type ServerHAState string

const (
	ServerHAStateCreatingStandby ServerHAState = "CreatingStandby"
	ServerHAStateFailingOver     ServerHAState = "FailingOver"
	ServerHAStateHealthy         ServerHAState = "Healthy"
	ServerHAStateNotEnabled      ServerHAState = "NotEnabled"
	ServerHAStateRemovingStandby ServerHAState = "RemovingStandby"
	ServerHAStateReplicatingData ServerHAState = "ReplicatingData"
)

func PossibleValuesForServerHAState() []string {
	return []string{
		string(ServerHAStateCreatingStandby),
		string(ServerHAStateFailingOver),
		string(ServerHAStateHealthy),
		string(ServerHAStateNotEnabled),
		string(ServerHAStateRemovingStandby),
		string(ServerHAStateReplicatingData),
	}
}

func parseServerHAState(input string) (*ServerHAState, error) {
	vals := map[string]ServerHAState{
		"creatingstandby": ServerHAStateCreatingStandby,
		"failingover":     ServerHAStateFailingOver,
		"healthy":         ServerHAStateHealthy,
		"notenabled":      ServerHAStateNotEnabled,
		"removingstandby": ServerHAStateRemovingStandby,
		"replicatingdata": ServerHAStateReplicatingData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerHAState(input)
	return &out, nil
}

type ServerPublicNetworkAccessState string

const (
	ServerPublicNetworkAccessStateDisabled ServerPublicNetworkAccessState = "Disabled"
	ServerPublicNetworkAccessStateEnabled  ServerPublicNetworkAccessState = "Enabled"
)

func PossibleValuesForServerPublicNetworkAccessState() []string {
	return []string{
		string(ServerPublicNetworkAccessStateDisabled),
		string(ServerPublicNetworkAccessStateEnabled),
	}
}

func parseServerPublicNetworkAccessState(input string) (*ServerPublicNetworkAccessState, error) {
	vals := map[string]ServerPublicNetworkAccessState{
		"disabled": ServerPublicNetworkAccessStateDisabled,
		"enabled":  ServerPublicNetworkAccessStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerPublicNetworkAccessState(input)
	return &out, nil
}

type ServerState string

const (
	ServerStateDisabled ServerState = "Disabled"
	ServerStateDropping ServerState = "Dropping"
	ServerStateReady    ServerState = "Ready"
	ServerStateStarting ServerState = "Starting"
	ServerStateStopped  ServerState = "Stopped"
	ServerStateStopping ServerState = "Stopping"
	ServerStateUpdating ServerState = "Updating"
)

func PossibleValuesForServerState() []string {
	return []string{
		string(ServerStateDisabled),
		string(ServerStateDropping),
		string(ServerStateReady),
		string(ServerStateStarting),
		string(ServerStateStopped),
		string(ServerStateStopping),
		string(ServerStateUpdating),
	}
}

func parseServerState(input string) (*ServerState, error) {
	vals := map[string]ServerState{
		"disabled": ServerStateDisabled,
		"dropping": ServerStateDropping,
		"ready":    ServerStateReady,
		"starting": ServerStateStarting,
		"stopped":  ServerStateStopped,
		"stopping": ServerStateStopping,
		"updating": ServerStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerState(input)
	return &out, nil
}

type ServerVersion string

const (
	ServerVersionOneFive  ServerVersion = "15"
	ServerVersionOneFour  ServerVersion = "14"
	ServerVersionOneOne   ServerVersion = "11"
	ServerVersionOneThree ServerVersion = "13"
	ServerVersionOneTwo   ServerVersion = "12"
)

func PossibleValuesForServerVersion() []string {
	return []string{
		string(ServerVersionOneFive),
		string(ServerVersionOneFour),
		string(ServerVersionOneOne),
		string(ServerVersionOneThree),
		string(ServerVersionOneTwo),
	}
}

func parseServerVersion(input string) (*ServerVersion, error) {
	vals := map[string]ServerVersion{
		"15": ServerVersionOneFive,
		"14": ServerVersionOneFour,
		"11": ServerVersionOneOne,
		"13": ServerVersionOneThree,
		"12": ServerVersionOneTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerVersion(input)
	return &out, nil
}

type SkuTier string

const (
	SkuTierBurstable       SkuTier = "Burstable"
	SkuTierGeneralPurpose  SkuTier = "GeneralPurpose"
	SkuTierMemoryOptimized SkuTier = "MemoryOptimized"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBurstable),
		string(SkuTierGeneralPurpose),
		string(SkuTierMemoryOptimized),
	}
}

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"burstable":       SkuTierBurstable,
		"generalpurpose":  SkuTierGeneralPurpose,
		"memoryoptimized": SkuTierMemoryOptimized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}

type StorageAutoGrow string

const (
	StorageAutoGrowDisabled StorageAutoGrow = "Disabled"
	StorageAutoGrowEnabled  StorageAutoGrow = "Enabled"
)

func PossibleValuesForStorageAutoGrow() []string {
	return []string{
		string(StorageAutoGrowDisabled),
		string(StorageAutoGrowEnabled),
	}
}

func parseStorageAutoGrow(input string) (*StorageAutoGrow, error) {
	vals := map[string]StorageAutoGrow{
		"disabled": StorageAutoGrowDisabled,
		"enabled":  StorageAutoGrowEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAutoGrow(input)
	return &out, nil
}
