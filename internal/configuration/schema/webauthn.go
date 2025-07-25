package schema

import (
	"time"

	"github.com/go-webauthn/webauthn/metadata"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/google/uuid"
)

// WebAuthn represents the webauthn config.
type WebAuthn struct {
	Disable              bool   `koanf:"disable" yaml:"disable" toml:"disable" json:"disable" jsonschema:"default=false,title=Disable" jsonschema_description:"Disables the WebAuthn 2FA functionality."`
	EnablePasskeyLogin   bool   `koanf:"enable_passkey_login" yaml:"enable_passkey_login" toml:"enable_passkey_login" json:"enable_passkey_login" jsonschema:"default=false,title=Enable Passkey Logins" jsonschema_description:"Allows users to sign in via Passkeys."`
	EnablePasskey2FA     bool   `koanf:"experimental_enable_passkey_uv_two_factors" yaml:"experimental_enable_passkey_uv_two_factors" toml:"experimental_enable_passkey_uv_two_factors" json:"experimental_enable_passkey_uv_two_factors" jsonschema:"default=false,title=Experimental: Enable User Verified Passkey Second Factor" jsonschema_description:"This option is experimental and WILL be removed. When true this will consider a Passkey login where the WebAuthn user verification requirement is met as being a single event 2FA login."`
	EnablePasskeyUpgrade bool   `koanf:"experimental_enable_passkey_upgrade" yaml:"experimental_enable_passkey_upgrade" toml:"experimental_enable_passkey_upgrade" json:"experimental_enable_passkey_upgrade" jsonschema:"default=false,title=Experimental: Enable User Verified Passkey Second Factor" jsonschema_description:"When true this will attempt to automatically upgrade a WebAuthn credential to a Passkey when the browser or authenticator does not report a newly registered WebAuthn credential to be a Passkey."`
	DisplayName          string `koanf:"display_name" yaml:"display_name,omitempty" toml:"display_name,omitempty" json:"display_name,omitempty" jsonschema:"default=Authelia,title=Display Name" jsonschema_description:"The display name attribute for the WebAuthn relying party."`

	ConveyancePreference protocol.ConveyancePreference `koanf:"attestation_conveyance_preference" yaml:"attestation_conveyance_preference,omitempty" toml:"attestation_conveyance_preference,omitempty" json:"attestation_conveyance_preference,omitempty" jsonschema:"default=indirect,enum=none,enum=indirect,enum=direct,title=Conveyance Preference" jsonschema_description:"The default conveyance preference for all WebAuthn credentials."`

	Timeout time.Duration `koanf:"timeout" yaml:"timeout,omitempty" toml:"timeout,omitempty" json:"timeout,omitempty" jsonschema:"default=60 seconds,title=Timeout" jsonschema_description:"The default timeout for all WebAuthn ceremonies."`

	Filtering         WebAuthnFiltering         `koanf:"filtering" yaml:"filtering,omitempty" toml:"filtering,omitempty" json:"filtering,omitempty" jsonschema:"title=Filtering" jsonschema_description:"WebAuthn Authenticator filtering settings."`
	SelectionCriteria WebAuthnSelectionCriteria `koanf:"selection_criteria" yaml:"selection_criteria,omitempty" toml:"selection_criteria,omitempty" json:"selection_criteria,omitempty" jsonschema_description:"WebAuthn Authenticator selection criteria settings."`
	Metadata          WebAuthnMetadata          `koanf:"metadata" yaml:"metadata,omitempty" toml:"metadata,omitempty" json:"metadata,omitempty" jsonschema_description:"WebAuthn Metadata Service settings."`
}

type WebAuthnMetadata struct {
	Enabled bool `koanf:"enabled" yaml:"enabled" toml:"enabled" json:"enabled" jsonschema:"default=false,title=Enabled" jsonschema_description:"Enables the use of the WebAuthn Metadata Service."`

	CachePolicy string `koanf:"cache_policy" yaml:"cache_policy" toml:"cache_policy" json:"cache_policy" jsonschema:"default=strict,enum=strict,enum=relaxed,title=Cache Policy" jsonschema_description:"WebAuthn Authenticator metadata cache policy."`

	ValidateTrustAnchor           bool `koanf:"validate_trust_anchor" yaml:"validate_trust_anchor" toml:"validate_trust_anchor" json:"validate_trust_anchor" jsonschema:"default=true,title=Validate Trust Anchor" jsonschema_description:"WebAuthn Authenticator metadata entry trust anchor validation."`
	ValidateEntry                 bool `koanf:"validate_entry" yaml:"validate_entry" toml:"validate_entry" json:"validate_entry" jsonschema:"default=true,title=Filtering" jsonschema_description:"WebAuthn Authenticator metadata entry validation requires the AAGUID exists as a MDS3 registered entry."`
	ValidateEntryPermitZeroAAGUID bool `koanf:"validate_entry_permit_zero_aaguid" yaml:"validate_entry_permit_zero_aaguid" toml:"validate_entry_permit_zero_aaguid" json:"validate_entry_permit_zero_aaguid" jsonschema:"default=true,title=Filtering" jsonschema_description:"WebAuthn Authenticator metadata entry validation zero AAGUID's can be skipped'."`

	ValidateStatus           bool                           `koanf:"validate_status" yaml:"validate_status" toml:"validate_status" json:"validate_status" jsonschema:"default=true,title=Validate Status" jsonschema_description:"WebAuthn Authenticator status validation."`
	ValidateStatusPermitted  []metadata.AuthenticatorStatus `koanf:"validate_status_permitted" yaml:"validate_status_permitted,omitempty" toml:"validate_status_permitted,omitempty" json:"validate_status_permitted,omitempty" jsonschema:"enum=FIDO_NOT_CERTIFIED,enum=FIDO_CERTIFIED,enum=USER_VERIFICATION_BYPASS,enum=ATTESTATION_KEY_COMPROMISE,enum=USER_KEY_REMOTE_COMPROMISE,enum=USER_KEY_PHYSICAL_COMPROMISE,enum=UPDATE_AVAILABLE,enum=REVOKED,enum=SELF_ASSERTION_SUBMITTED,enum=FIDO_CERTIFIED_L1,enum=FIDO_CERTIFIED_L1plus,enum=FIDO_CERTIFIED_L2,enum=FIDO_CERTIFIED_L2plus,enum=FIDO_CERTIFIED_L3,enum=FIDO_CERTIFIED_L3plus,title=Validate Status (Permitted Statuses)" jsonschema_description:"WebAuthn Authenticator status validation can be configured to permit certain statuses. Generally this is discouraged."`
	ValidateStatusProhibited []metadata.AuthenticatorStatus `koanf:"validate_status_prohibited" yaml:"validate_status_prohibited,omitempty" toml:"validate_status_prohibited,omitempty" json:"validate_status_prohibited,omitempty" jsonschema:"enum=FIDO_NOT_CERTIFIED,enum=FIDO_CERTIFIED,enum=USER_VERIFICATION_BYPASS,enum=ATTESTATION_KEY_COMPROMISE,enum=USER_KEY_REMOTE_COMPROMISE,enum=USER_KEY_PHYSICAL_COMPROMISE,enum=UPDATE_AVAILABLE,enum=REVOKED,enum=SELF_ASSERTION_SUBMITTED,enum=FIDO_CERTIFIED_L1,enum=FIDO_CERTIFIED_L1plus,enum=FIDO_CERTIFIED_L2,enum=FIDO_CERTIFIED_L2plus,enum=FIDO_CERTIFIED_L3,enum=FIDO_CERTIFIED_L3plus,title=Validate Status (Prohibited Statuses)" jsonschema_description:"WebAuthn Authenticator status validation can prohibit certain statuses. Generally this is discouraged as the defaults are safe."`
}

type WebAuthnSelectionCriteria struct {
	Attachment       protocol.AuthenticatorAttachment     `koanf:"attachment" yaml:"attachment,omitempty" toml:"attachment,omitempty" json:"attachment,omitempty" jsonschema:"default=,enum=,enum=platform,enum=cross-platform,title=Attachment" jsonschema_description:"WebAuthn Authenticator attachment preference."`
	Discoverability  protocol.ResidentKeyRequirement      `koanf:"discoverability" yaml:"discoverability,omitempty" toml:"discoverability,omitempty" json:"discoverability,omitempty" jsonschema:"default=preferred,enum=discouraged,enum=preferred,enum=required,title=Discoverability Selection" jsonschema_description:"The default discoverable preference when registering WebAuthn credentials."`
	UserVerification protocol.UserVerificationRequirement `koanf:"user_verification" yaml:"user_verification,omitempty" toml:"user_verification,omitempty" json:"user_verification,omitempty" jsonschema:"default=preferred,enum=discouraged,enum=preferred,enum=required,title=User Verification" jsonschema_description:"The default user verification preference for all WebAuthn credentials."`
}

type WebAuthnFiltering struct {
	ProhibitBackupEligibility bool        `koanf:"prohibit_backup_eligibility" yaml:"prohibit_backup_eligibility" toml:"prohibit_backup_eligibility" json:"prohibit_backup_eligibility" jsonschema:"default=false,title=Prohibit Backup Eligibility" jsonschema_description:"Prohibits registering authenticators which claim backup eligibility i.e. exporting credentials off of the device."`
	PermittedAAGUIDs          []uuid.UUID `koanf:"permitted_aaguids" yaml:"permitted_aaguids,omitempty" toml:"permitted_aaguids,omitempty" json:"permitted_aaguids,omitempty" jsonschema:"title=Permitted AAGUIDs" jsonschema_description:"List of allowed WebAuthn AAGUIDs. No other authenticator can be registered."`
	ProhibitedAAGUIDs         []uuid.UUID `koanf:"prohibited_aaguids" yaml:"prohibited_aaguids,omitempty" toml:"prohibited_aaguids,omitempty" json:"prohibited_aaguids,omitempty" jsonschema:"title=Prohibited AAGUIDs" jsonschema_description:"List of prohibited WebAuthn AAGUIDs. Authenticators with these AAGUIDs cannot be registered."`
}

// DefaultWebAuthnConfiguration describes the default values for the WebAuthn.
var DefaultWebAuthnConfiguration = WebAuthn{
	DisplayName: "Authelia",
	Timeout:     time.Second * 60,

	ConveyancePreference: protocol.PreferIndirectAttestation,
	Metadata: WebAuthnMetadata{
		Enabled:                       false,
		ValidateTrustAnchor:           true,
		ValidateEntry:                 true,
		ValidateEntryPermitZeroAAGUID: false,
		ValidateStatus:                true,
		ValidateStatusPermitted:       []metadata.AuthenticatorStatus{},
		ValidateStatusProhibited: []metadata.AuthenticatorStatus{
			metadata.AttestationKeyCompromise,
			metadata.UserVerificationBypass,
			metadata.UserKeyRemoteCompromise,
			metadata.UserKeyPhysicalCompromise,
			metadata.Revoked,
		},
	},
}
