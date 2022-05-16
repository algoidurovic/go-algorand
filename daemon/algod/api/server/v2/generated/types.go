// Package generated provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

import (
	"encoding/json"
	"time"
)

// Account defines model for Account.
type Account struct {

	// the account public key
	Address string `json:"address"`

	// \[algo\] total number of MicroAlgos in the account
	Amount uint64 `json:"amount"`

	// specifies the amount of MicroAlgos in the account, without the pending rewards.
	AmountWithoutPendingRewards uint64 `json:"amount-without-pending-rewards"`

	// \[appl\] applications local data stored in this account.
	//
	// Note the raw object uses `map[int] -> AppLocalState` for this type.
	AppsLocalState *[]ApplicationLocalState `json:"apps-local-state,omitempty"`

	// \[teap\] the sum of all extra application program pages for this account.
	AppsTotalExtraPages *uint64 `json:"apps-total-extra-pages,omitempty"`

	// Specifies maximums on the number of each type that may be stored.
	AppsTotalSchema *ApplicationStateSchema `json:"apps-total-schema,omitempty"`

	// \[asset\] assets held by this account.
	//
	// Note the raw object uses `map[int] -> AssetHolding` for this type.
	Assets *[]AssetHolding `json:"assets,omitempty"`

	// \[spend\] the address against which signing should be checked. If empty, the address of the current account is used. This field can be updated in any transaction by setting the RekeyTo field.
	AuthAddr *string `json:"auth-addr,omitempty"`

	// \[appp\] parameters of applications created by this account including app global data.
	//
	// Note: the raw account uses `map[int] -> AppParams` for this type.
	CreatedApps *[]Application `json:"created-apps,omitempty"`

	// \[apar\] parameters of assets created by this account.
	//
	// Note: the raw account uses `map[int] -> Asset` for this type.
	CreatedAssets *[]Asset `json:"created-assets,omitempty"`

	// MicroAlgo balance required by the account.
	//
	// The requirement grows based on asset and application usage.
	MinBalance uint64 `json:"min-balance"`

	// AccountParticipation describes the parameters used by this account in consensus protocol.
	Participation *AccountParticipation `json:"participation,omitempty"`

	// amount of MicroAlgos of pending rewards in this account.
	PendingRewards uint64 `json:"pending-rewards"`

	// \[ebase\] used as part of the rewards computation. Only applicable to accounts which are participating.
	RewardBase *uint64 `json:"reward-base,omitempty"`

	// \[ern\] total rewards of MicroAlgos the account has received, including pending rewards.
	Rewards uint64 `json:"rewards"`

	// The round for which this information is relevant.
	Round uint64 `json:"round"`

	// Indicates what type of signature is used by this account, must be one of:
	// * sig
	// * msig
	// * lsig
	SigType *string `json:"sig-type,omitempty"`

	// \[onl\] delegation status of the account's MicroAlgos
	// * Offline - indicates that the associated account is delegated.
	// *  Online  - indicates that the associated account used as part of the delegation pool.
	// *   NotParticipating - indicates that the associated account is neither a delegator nor a delegate.
	Status string `json:"status"`

	// The count of all applications that have been opted in, equivalent to the count of application local data (AppLocalState objects) stored in this account.
	TotalAppsOptedIn uint64 `json:"total-apps-opted-in"`

	// The count of all assets that have been opted in, equivalent to the count of AssetHolding objects held by this account.
	TotalAssetsOptedIn uint64 `json:"total-assets-opted-in"`

	// The count of all apps (AppParams objects) created by this account.
	TotalCreatedApps uint64 `json:"total-created-apps"`

	// The count of all assets (AssetParams objects) created by this account.
	TotalCreatedAssets uint64 `json:"total-created-assets"`
}

// AccountParticipation defines model for AccountParticipation.
type AccountParticipation struct {

	// \[sel\] Selection public key (if any) currently registered for this round.
	SelectionParticipationKey []byte `json:"selection-participation-key"`

	// \[stprf\] Root of the state proof key (if any)
	StateProofKey *[]byte `json:"state-proof-key,omitempty"`

	// \[voteFst\] First round for which this participation is valid.
	VoteFirstValid uint64 `json:"vote-first-valid"`

	// \[voteKD\] Number of subkeys in each batch of participation keys.
	VoteKeyDilution uint64 `json:"vote-key-dilution"`

	// \[voteLst\] Last round for which this participation is valid.
	VoteLastValid uint64 `json:"vote-last-valid"`

	// \[vote\] root participation public key (if any) currently registered for this round.
	VoteParticipationKey []byte `json:"vote-participation-key"`
}

// AccountStateDelta defines model for AccountStateDelta.
type AccountStateDelta struct {
	Address string `json:"address"`

	// Application state delta.
	Delta StateDelta `json:"delta"`
}

// Application defines model for Application.
type Application struct {

	// \[appidx\] application index.
	Id uint64 `json:"id"`

	// Stores the global information associated with an application.
	Params ApplicationParams `json:"params"`
}

// ApplicationLocalState defines model for ApplicationLocalState.
type ApplicationLocalState struct {

	// The application which this local state is for.
	Id uint64 `json:"id"`

	// Represents a key-value store for use in an application.
	KeyValue *TealKeyValueStore `json:"key-value,omitempty"`

	// Specifies maximums on the number of each type that may be stored.
	Schema ApplicationStateSchema `json:"schema"`
}

// ApplicationParams defines model for ApplicationParams.
type ApplicationParams struct {

	// \[approv\] approval program.
	ApprovalProgram []byte `json:"approval-program"`

	// \[clearp\] approval program.
	ClearStateProgram []byte `json:"clear-state-program"`

	// The address that created this application. This is the address where the parameters and global state for this application can be found.
	Creator string `json:"creator"`

	// \[epp\] the amount of extra program pages available to this app.
	ExtraProgramPages *uint64 `json:"extra-program-pages,omitempty"`

	// Represents a key-value store for use in an application.
	GlobalState *TealKeyValueStore `json:"global-state,omitempty"`

	// Specifies maximums on the number of each type that may be stored.
	GlobalStateSchema *ApplicationStateSchema `json:"global-state-schema,omitempty"`

	// Specifies maximums on the number of each type that may be stored.
	LocalStateSchema *ApplicationStateSchema `json:"local-state-schema,omitempty"`
}

// ApplicationStateSchema defines model for ApplicationStateSchema.
type ApplicationStateSchema struct {

	// \[nbs\] num of byte slices.
	NumByteSlice uint64 `json:"num-byte-slice"`

	// \[nui\] num of uints.
	NumUint uint64 `json:"num-uint"`
}

// Asset defines model for Asset.
type Asset struct {

	// unique asset identifier
	Index uint64 `json:"index"`

	// AssetParams specifies the parameters for an asset.
	//
	// \[apar\] when part of an AssetConfig transaction.
	//
	// Definition:
	// data/transactions/asset.go : AssetParams
	Params AssetParams `json:"params"`
}

// AssetHolding defines model for AssetHolding.
type AssetHolding struct {

	// \[a\] number of units held.
	Amount uint64 `json:"amount"`

	// Asset ID of the holding.
	AssetId uint64 `json:"asset-id"`

	// \[f\] whether or not the holding is frozen.
	IsFrozen bool `json:"is-frozen"`
}

// AssetParams defines model for AssetParams.
type AssetParams struct {

	// \[c\] Address of account used to clawback holdings of this asset.  If empty, clawback is not permitted.
	Clawback *string `json:"clawback,omitempty"`

	// The address that created this asset. This is the address where the parameters for this asset can be found, and also the address where unwanted asset units can be sent in the worst case.
	Creator string `json:"creator"`

	// \[dc\] The number of digits to use after the decimal point when displaying this asset. If 0, the asset is not divisible. If 1, the base unit of the asset is in tenths. If 2, the base unit of the asset is in hundredths, and so on. This value must be between 0 and 19 (inclusive).
	Decimals uint64 `json:"decimals"`

	// \[df\] Whether holdings of this asset are frozen by default.
	DefaultFrozen *bool `json:"default-frozen,omitempty"`

	// \[f\] Address of account used to freeze holdings of this asset.  If empty, freezing is not permitted.
	Freeze *string `json:"freeze,omitempty"`

	// \[m\] Address of account used to manage the keys of this asset and to destroy it.
	Manager *string `json:"manager,omitempty"`

	// \[am\] A commitment to some unspecified asset metadata. The format of this metadata is up to the application.
	MetadataHash *[]byte `json:"metadata-hash,omitempty"`

	// \[an\] Name of this asset, as supplied by the creator. Included only when the asset name is composed of printable utf-8 characters.
	Name *string `json:"name,omitempty"`

	// Base64 encoded name of this asset, as supplied by the creator.
	NameB64 *[]byte `json:"name-b64,omitempty"`

	// \[r\] Address of account holding reserve (non-minted) units of this asset.
	Reserve *string `json:"reserve,omitempty"`

	// \[t\] The total number of units of this asset.
	Total uint64 `json:"total"`

	// \[un\] Name of a unit of this asset, as supplied by the creator. Included only when the name of a unit of this asset is composed of printable utf-8 characters.
	UnitName *string `json:"unit-name,omitempty"`

	// Base64 encoded name of a unit of this asset, as supplied by the creator.
	UnitNameB64 *[]byte `json:"unit-name-b64,omitempty"`

	// \[au\] URL where more information about the asset can be retrieved. Included only when the URL is composed of printable utf-8 characters.
	Url *string `json:"url,omitempty"`

	// Base64 encoded URL where more information about the asset can be retrieved.
	UrlB64 *[]byte `json:"url-b64,omitempty"`
}

// BuildVersion defines model for BuildVersion.
type BuildVersion struct {
	Branch      string `json:"branch"`
	BuildNumber uint64 `json:"build_number"`
	Channel     string `json:"channel"`
	CommitHash  string `json:"commit_hash"`
	Major       uint64 `json:"major"`
	Minor       uint64 `json:"minor"`
}

// DryrunRequest defines model for DryrunRequest.
type DryrunRequest struct {
	Accounts []Account     `json:"accounts"`
	Apps     []Application `json:"apps"`

	// LatestTimestamp is available to some TEAL scripts. Defaults to the latest confirmed timestamp this algod is attached to.
	LatestTimestamp uint64 `json:"latest-timestamp"`

	// ProtocolVersion specifies a specific version string to operate under, otherwise whatever the current protocol of the network this algod is running in.
	ProtocolVersion string `json:"protocol-version"`

	// Round is available to some TEAL scripts. Defaults to the current round on the network this algod is attached to.
	Round   uint64            `json:"round"`
	Sources []DryrunSource    `json:"sources"`
	Txns    []json.RawMessage `json:"txns"`
}

// DryrunSource defines model for DryrunSource.
type DryrunSource struct {
	AppIndex uint64 `json:"app-index"`

	// FieldName is what kind of sources this is. If lsig then it goes into the transactions[this.TxnIndex].LogicSig. If approv or clearp it goes into the Approval Program or Clear State Program of application[this.AppIndex].
	FieldName string `json:"field-name"`
	Source    string `json:"source"`
	TxnIndex  uint64 `json:"txn-index"`
}

// DryrunState defines model for DryrunState.
type DryrunState struct {

	// Evaluation error if any
	Error *string `json:"error,omitempty"`

	// Line number
	Line uint64 `json:"line"`

	// Program counter
	Pc      uint64       `json:"pc"`
	Scratch *[]TealValue `json:"scratch,omitempty"`
	Stack   []TealValue  `json:"stack"`
}

// DryrunTxnResult defines model for DryrunTxnResult.
type DryrunTxnResult struct {
	AppCallMessages *[]string      `json:"app-call-messages,omitempty"`
	AppCallTrace    *[]DryrunState `json:"app-call-trace,omitempty"`

	// Budget added during execution of app call transaction.
	BudgetAdded *uint64 `json:"budget-added,omitempty"`

	// Budget consumed during execution of app call transaction.
	BudgetConsumed *uint64 `json:"budget-consumed,omitempty"`

	// Net cost of app execution. Will be DEPRECATED.
	Cost *uint64 `json:"cost,omitempty"`

	// Disassembled program line by line.
	Disassembly []string `json:"disassembly"`

	// Application state delta.
	GlobalDelta *StateDelta          `json:"global-delta,omitempty"`
	LocalDeltas *[]AccountStateDelta `json:"local-deltas,omitempty"`

	// Disassembled lsig program line by line.
	LogicSigDisassembly *[]string      `json:"logic-sig-disassembly,omitempty"`
	LogicSigMessages    *[]string      `json:"logic-sig-messages,omitempty"`
	LogicSigTrace       *[]DryrunState `json:"logic-sig-trace,omitempty"`
	Logs                *[][]byte      `json:"logs,omitempty"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message string                  `json:"message"`
}

// EvalDelta defines model for EvalDelta.
type EvalDelta struct {

	// \[at\] delta action.
	Action uint64 `json:"action"`

	// \[bs\] bytes value.
	Bytes *string `json:"bytes,omitempty"`

	// \[ui\] uint value.
	Uint *uint64 `json:"uint,omitempty"`
}

// EvalDeltaKeyValue defines model for EvalDeltaKeyValue.
type EvalDeltaKeyValue struct {
	Key string `json:"key"`

	// Represents a TEAL value delta.
	Value EvalDelta `json:"value"`
}

// ParticipationKey defines model for ParticipationKey.
type ParticipationKey struct {

	// Address the key was generated for.
	Address string `json:"address"`

	// When registered, this is the first round it may be used.
	EffectiveFirstValid *uint64 `json:"effective-first-valid,omitempty"`

	// When registered, this is the last round it may be used.
	EffectiveLastValid *uint64 `json:"effective-last-valid,omitempty"`

	// The key's ParticipationID.
	Id string `json:"id"`

	// AccountParticipation describes the parameters used by this account in consensus protocol.
	Key AccountParticipation `json:"key"`

	// Round when this key was last used to propose a block.
	LastBlockProposal *uint64 `json:"last-block-proposal,omitempty"`

	// Round when this key was last used to generate a state proof.
	LastStateProof *uint64 `json:"last-state-proof,omitempty"`

	// Round when this key was last used to vote.
	LastVote *uint64 `json:"last-vote,omitempty"`
}

// PendingTransactionResponse defines model for PendingTransactionResponse.
type PendingTransactionResponse struct {

	// The application index if the transaction was found and it created an application.
	ApplicationIndex *uint64 `json:"application-index,omitempty"`

	// The number of the asset's unit that were transferred to the close-to address.
	AssetClosingAmount *uint64 `json:"asset-closing-amount,omitempty"`

	// The asset index if the transaction was found and it created an asset.
	AssetIndex *uint64 `json:"asset-index,omitempty"`

	// Rewards in microalgos applied to the close remainder to account.
	CloseRewards *uint64 `json:"close-rewards,omitempty"`

	// Closing amount for the transaction.
	ClosingAmount *uint64 `json:"closing-amount,omitempty"`

	// The round where this transaction was confirmed, if present.
	ConfirmedRound *uint64 `json:"confirmed-round,omitempty"`

	// Application state delta.
	GlobalStateDelta *StateDelta `json:"global-state-delta,omitempty"`

	// Inner transactions produced by application execution.
	InnerTxns *[]PendingTransactionResponse `json:"inner-txns,omitempty"`

	// \[ld\] Local state key/value changes for the application being executed by this transaction.
	LocalStateDelta *[]AccountStateDelta `json:"local-state-delta,omitempty"`

	// \[lg\] Logs for the application being executed by this transaction.
	Logs *[][]byte `json:"logs,omitempty"`

	// Indicates that the transaction was kicked out of this node's transaction pool (and specifies why that happened).  An empty string indicates the transaction wasn't kicked out of this node's txpool due to an error.
	PoolError string `json:"pool-error"`

	// Rewards in microalgos applied to the receiver account.
	ReceiverRewards *uint64 `json:"receiver-rewards,omitempty"`

	// Rewards in microalgos applied to the sender account.
	SenderRewards *uint64 `json:"sender-rewards,omitempty"`

	// The raw signed transaction.
	Txn map[string]interface{} `json:"txn"`
}

// StateDelta defines model for StateDelta.
type StateDelta []EvalDeltaKeyValue

// TealKeyValue defines model for TealKeyValue.
type TealKeyValue struct {
	Key string `json:"key"`

	// Represents a TEAL value.
	Value TealValue `json:"value"`
}

// TealKeyValueStore defines model for TealKeyValueStore.
type TealKeyValueStore []TealKeyValue

// TealValue defines model for TealValue.
type TealValue struct {

	// \[tb\] bytes value.
	Bytes string `json:"bytes"`

	// \[tt\] value type. Value `1` refers to **bytes**, value `2` refers to **uint**
	Type uint64 `json:"type"`

	// \[ui\] uint value.
	Uint uint64 `json:"uint"`
}

// Version defines model for Version.
type Version struct {
	Build          BuildVersion `json:"build"`
	GenesisHashB64 []byte       `json:"genesis_hash_b64"`
	GenesisId      string       `json:"genesis_id"`
	Versions       []string     `json:"versions"`
}

// AccountId defines model for account-id.
type AccountId string

// Address defines model for address.
type Address string

// AddressRole defines model for address-role.
type AddressRole string

// AfterTime defines model for after-time.
type AfterTime time.Time

// AssetId defines model for asset-id.
type AssetId uint64

// BeforeTime defines model for before-time.
type BeforeTime time.Time

// Catchpoint defines model for catchpoint.
type Catchpoint string

// CurrencyGreaterThan defines model for currency-greater-than.
type CurrencyGreaterThan uint64

// CurrencyLessThan defines model for currency-less-than.
type CurrencyLessThan uint64

// ExcludeCloseTo defines model for exclude-close-to.
type ExcludeCloseTo bool

// Format defines model for format.
type Format string

// Limit defines model for limit.
type Limit uint64

// Max defines model for max.
type Max uint64

// MaxRound defines model for max-round.
type MaxRound uint64

// MinRound defines model for min-round.
type MinRound uint64

// Next defines model for next.
type Next string

// NotePrefix defines model for note-prefix.
type NotePrefix string

// Round defines model for round.
type Round uint64

// RoundNumber defines model for round-number.
type RoundNumber uint64

// SigType defines model for sig-type.
type SigType string

// TxId defines model for tx-id.
type TxId string

// TxType defines model for tx-type.
type TxType string

// AccountApplicationResponse defines model for AccountApplicationResponse.
type AccountApplicationResponse struct {

	// Stores local state associated with an application.
	AppLocalState *ApplicationLocalState `json:"app-local-state,omitempty"`

	// Stores the global information associated with an application.
	CreatedApp *ApplicationParams `json:"created-app,omitempty"`

	// The round for which this information is relevant.
	Round uint64 `json:"round"`
}

// AccountAssetResponse defines model for AccountAssetResponse.
type AccountAssetResponse struct {

	// Describes an asset held by an account.
	//
	// Definition:
	// data/basics/userBalance.go : AssetHolding
	AssetHolding *AssetHolding `json:"asset-holding,omitempty"`

	// AssetParams specifies the parameters for an asset.
	//
	// \[apar\] when part of an AssetConfig transaction.
	//
	// Definition:
	// data/transactions/asset.go : AssetParams
	CreatedAsset *AssetParams `json:"created-asset,omitempty"`

	// The round for which this information is relevant.
	Round uint64 `json:"round"`
}

// AccountResponse defines model for AccountResponse.
type AccountResponse Account

// ApplicationResponse defines model for ApplicationResponse.
type ApplicationResponse Application

// AssetResponse defines model for AssetResponse.
type AssetResponse Asset

// BlockResponse defines model for BlockResponse.
type BlockResponse struct {

	// Block header data.
	Block map[string]interface{} `json:"block"`

	// Optional certificate object. This is only included when the format is set to message pack.
	Cert *map[string]interface{} `json:"cert,omitempty"`
}

// CatchpointAbortResponse defines model for CatchpointAbortResponse.
type CatchpointAbortResponse struct {

	// Catchup abort response string
	CatchupMessage string `json:"catchup-message"`
}

// CatchpointStartResponse defines model for CatchpointStartResponse.
type CatchpointStartResponse struct {

	// Catchup start response string
	CatchupMessage string `json:"catchup-message"`
}

// CompileResponse defines model for CompileResponse.
type CompileResponse struct {

	// base32 SHA512_256 of program bytes (Address style)
	Hash string `json:"hash"`

	// base64 encoded program bytes
	Result string `json:"result"`

	// JSON of the source map
	Sourcemap *map[string]interface{} `json:"sourcemap,omitempty"`
}

// DisassembleResponse defines model for DisassembleResponse.
type DisassembleResponse struct {

	// disassembled Teal code
	Result string `json:"result"`
}

// DryrunResponse defines model for DryrunResponse.
type DryrunResponse struct {
	Error string `json:"error"`

	// Protocol version is the protocol version Dryrun was operated under.
	ProtocolVersion string            `json:"protocol-version"`
	Txns            []DryrunTxnResult `json:"txns"`
}

// NodeStatusResponse defines model for NodeStatusResponse.
type NodeStatusResponse struct {

	// The current catchpoint that is being caught up to
	Catchpoint *string `json:"catchpoint,omitempty"`

	// The number of blocks that have already been obtained by the node as part of the catchup
	CatchpointAcquiredBlocks *uint64 `json:"catchpoint-acquired-blocks,omitempty"`

	// The number of accounts from the current catchpoint that have been processed so far as part of the catchup
	CatchpointProcessedAccounts *uint64 `json:"catchpoint-processed-accounts,omitempty"`

	// The total number of accounts included in the current catchpoint
	CatchpointTotalAccounts *uint64 `json:"catchpoint-total-accounts,omitempty"`

	// The total number of blocks that are required to complete the current catchpoint catchup
	CatchpointTotalBlocks *uint64 `json:"catchpoint-total-blocks,omitempty"`

	// The number of accounts from the current catchpoint that have been verified so far as part of the catchup
	CatchpointVerifiedAccounts *uint64 `json:"catchpoint-verified-accounts,omitempty"`

	// CatchupTime in nanoseconds
	CatchupTime uint64 `json:"catchup-time"`

	// The last catchpoint seen by the node
	LastCatchpoint *string `json:"last-catchpoint,omitempty"`

	// LastRound indicates the last round seen
	LastRound uint64 `json:"last-round"`

	// LastVersion indicates the last consensus version supported
	LastVersion string `json:"last-version"`

	// NextVersion of consensus protocol to use
	NextVersion string `json:"next-version"`

	// NextVersionRound is the round at which the next consensus version will apply
	NextVersionRound uint64 `json:"next-version-round"`

	// NextVersionSupported indicates whether the next consensus version is supported by this node
	NextVersionSupported bool `json:"next-version-supported"`

	// StoppedAtUnsupportedRound indicates that the node does not support the new rounds and has stopped making progress
	StoppedAtUnsupportedRound bool `json:"stopped-at-unsupported-round"`

	// TimeSinceLastRound in nanoseconds
	TimeSinceLastRound uint64 `json:"time-since-last-round"`
}

// ParticipationKeyResponse defines model for ParticipationKeyResponse.
type ParticipationKeyResponse ParticipationKey

// ParticipationKeysResponse defines model for ParticipationKeysResponse.
type ParticipationKeysResponse []ParticipationKey

// PendingTransactionsResponse defines model for PendingTransactionsResponse.
type PendingTransactionsResponse struct {

	// An array of signed transaction objects.
	TopTransactions []map[string]interface{} `json:"top-transactions"`

	// Total number of transactions in the pool.
	TotalTransactions uint64 `json:"total-transactions"`
}

// PostParticipationResponse defines model for PostParticipationResponse.
type PostParticipationResponse struct {

	// encoding of the participation ID.
	PartId string `json:"partId"`
}

// PostTransactionsResponse defines model for PostTransactionsResponse.
type PostTransactionsResponse struct {

	// encoding of the transaction hash.
	TxId string `json:"txId"`
}

// ProofResponse defines model for ProofResponse.
type ProofResponse struct {

	// The type of hash function used to create the proof, must be one of:
	// * sha512_256
	// * sha256
	Hashtype string `json:"hashtype"`

	// Index of the transaction in the block's payset.
	Idx uint64 `json:"idx"`

	// Merkle proof of transaction membership.
	Proof []byte `json:"proof"`

	// Hash of SignedTxnInBlock for verifying proof.
	Stibhash []byte `json:"stibhash"`

	// Represents the depth of the tree that is being proven, i.e. the number of edges from a leaf to the root.
	Treedepth uint64 `json:"treedepth"`
}

// SupplyResponse defines model for SupplyResponse.
type SupplyResponse struct {

	// Round
	CurrentRound uint64 `json:"current_round"`

	// OnlineMoney
	OnlineMoney uint64 `json:"online-money"`

	// TotalMoney
	TotalMoney uint64 `json:"total-money"`
}

// TransactionParametersResponse defines model for TransactionParametersResponse.
type TransactionParametersResponse struct {

	// ConsensusVersion indicates the consensus protocol version
	// as of LastRound.
	ConsensusVersion string `json:"consensus-version"`

	// Fee is the suggested transaction fee
	// Fee is in units of micro-Algos per byte.
	// Fee may fall to zero but transactions must still have a fee of
	// at least MinTxnFee for the current network protocol.
	Fee uint64 `json:"fee"`

	// GenesisHash is the hash of the genesis block.
	GenesisHash []byte `json:"genesis-hash"`

	// GenesisID is an ID listed in the genesis block.
	GenesisId string `json:"genesis-id"`

	// LastRound indicates the last round seen
	LastRound uint64 `json:"last-round"`

	// The minimum transaction fee (not per byte) required for the
	// txn to validate for the current network protocol.
	MinFee uint64 `json:"min-fee"`
}

// VersionsResponse defines model for VersionsResponse.
type VersionsResponse Version

// AccountInformationParams defines parameters for AccountInformation.
type AccountInformationParams struct {

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`

	// When set to `all` will exclude asset holdings, application local state, created asset parameters, any created application parameters. Defaults to `none`.
	Exclude *string `json:"exclude,omitempty"`
}

// AccountApplicationInformationParams defines parameters for AccountApplicationInformation.
type AccountApplicationInformationParams struct {

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// AccountAssetInformationParams defines parameters for AccountAssetInformation.
type AccountAssetInformationParams struct {

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// GetPendingTransactionsByAddressParams defines parameters for GetPendingTransactionsByAddress.
type GetPendingTransactionsByAddressParams struct {

	// Truncated number of transactions to display. If max=0, returns all pending txns.
	Max *uint64 `json:"max,omitempty"`

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// GetBlockParams defines parameters for GetBlock.
type GetBlockParams struct {

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// GetProofParams defines parameters for GetProof.
type GetProofParams struct {

	// The type of hash function used to create the proof, must be one of:
	// * sha512_256
	// * sha256
	Hashtype *string `json:"hashtype,omitempty"`

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// TealCompileParams defines parameters for TealCompile.
type TealCompileParams struct {

	// When set to `true`, returns the source map of the program as a JSON. Defaults to `false`.
	Sourcemap *bool `json:"sourcemap,omitempty"`
}

// TealDryrunJSONBody defines parameters for TealDryrun.
type TealDryrunJSONBody DryrunRequest

// GetPendingTransactionsParams defines parameters for GetPendingTransactions.
type GetPendingTransactionsParams struct {

	// Truncated number of transactions to display. If max=0, returns all pending txns.
	Max *uint64 `json:"max,omitempty"`

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// PendingTransactionInformationParams defines parameters for PendingTransactionInformation.
type PendingTransactionInformationParams struct {

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// TealDryrunRequestBody defines body for TealDryrun for application/json ContentType.
type TealDryrunJSONRequestBody TealDryrunJSONBody
