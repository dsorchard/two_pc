package common

const MasterPort = "localhost:7170"
const ReplicaPortStart = 7171

var KilledSelfMarker = "::justkilledself::"
var FirstRestartAfterSuicideMarker = "::firstrestartaftersuicide::"

//----------------------------------------------------------------------

type MasterDeath int

const (
	MasterDontDie MasterDeath = iota
	MasterDieBeforeLoggingCommitted
	MasterDieAfterLoggingCommitted
)

//----------------------------------------------------------------------

type ReplicaDeath int

const (
	ReplicaDontDie ReplicaDeath = iota

	// During mutation
	ReplicaDieBeforeProcessingMutateRequest
	ReplicaDieAfterLoggingPrepared

	// During commit
	ReplicaDieBeforeProcessingCommit
	ReplicaDieAfterDeletingFromTempStore
	ReplicaDieAfterLoggingCommitted
)
