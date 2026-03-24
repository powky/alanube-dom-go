package alanube

// ErrorCode represents an Alanube API error code.
type ErrorCode string

// --- Company validation errors (async) ---

const (
	ErrCertificateExpired          ErrorCode = "AP1013" // The certificate is expired
	ErrCertificateRequired         ErrorCode = "AP1003" // The certificate is required
	ErrCompanyHasBeenCertified     ErrorCode = "AP1008" // Company has been certified
	ErrCompanyInfoIncomplete       ErrorCode = "AP1006" // Company info incomplete for set test
	ErrCompanyNotFound             ErrorCode = "AP1002" // Company not found
	ErrExpiredCertificate          ErrorCode = "AP1005" // The certificate is expired
	ErrHasMainCompany              ErrorCode = "AP1004" // User already has a main company
	ErrInvalidCertificate          ErrorCode = "AP1001" // The certificate is invalid
	ErrInvalidCertificateFormat    ErrorCode = "AP1011" // Invalid certificate format (must be PKCS12)
	ErrInvalidCertificateLength    ErrorCode = "AP1012" // Invalid certificate length (check base64)
	ErrInvalidCertificatePassword  ErrorCode = "AP1010" // Invalid certificate password
	ErrLogoSizeExceeded            ErrorCode = "AP1009" // Logo size exceeded 150KB limit
	ErrMustHaveMainCompany         ErrorCode = "AP1015" // User must own a main company
	ErrSenderRNCNotMatchCompany    ErrorCode = "AP1016" // Sender RNC not match with company identification
	ErrSignatureTypeNotSupported   ErrorCode = "AP1007" // Signature type not supported
	ErrWebhookURLInvalid           ErrorCode = "AP1014" // Webhook URL is invalid
)

// --- User errors ---

const (
	ErrUserAlreadyExists ErrorCode = "AP0001" // User already exists
	ErrUserNotFound      ErrorCode = "AP0002" // User not found
)

// --- DGII errors ---

const (
	ErrDGIIAuthenticationError     ErrorCode = "AEP2003" // Error authenticating with DGII
	ErrDGIIBadRequestError         ErrorCode = "AEP2009" // DGII returned Bad Request
	ErrDGIIConnResetError          ErrorCode = "AEP2005" // Connection unexpectedly closed
	ErrDGIIEnvironmentNotSupported ErrorCode = "AEP2002" // Environment not supported
	ErrDGIIGeneralError            ErrorCode = "AEP2004" // Error communicating with DGII
	ErrDGIIHostUnreachable         ErrorCode = "AEP2012" // DGII host unreachable
	ErrDGIILeafSignatureError      ErrorCode = "AEP2013" // DGII leaf signature verification failed
	ErrDGIINotFoundError           ErrorCode = "AEP2008" // DGII returned Not Found
	ErrDGIIParseError              ErrorCode = "AEP2011" // DGII returned parsing error
	ErrDGIIServiceUnavailable      ErrorCode = "AEP2007" // DGII service unavailable
	ErrDGIITimeoutError            ErrorCode = "AEP2006" // Connection with DGII timed out
	ErrDGIIUnauthorizedError       ErrorCode = "AEP2010" // DGII returned Unauthorized
	ErrDGIIValueNotFound           ErrorCode = "AEP2001" // DGII key/value not found
)

// --- Document errors ---

const (
	ErrBuyerRNCOrForeignIDRequired     ErrorCode = "AP3006" // Buyer RNC or foreign identifier required (invoice > DOP$250K)
	ErrCNTextModifyCodeTypeToBeZero    ErrorCode = "AP3015" // Credit note modification code 2 must have total amount 0
	ErrCreditNoteCancelDifferentAmount ErrorCode = "AP3014" // Cancellation note total must match referenced document
	ErrDocumentInProcess               ErrorCode = "AP3010" // Document is in process
	ErrDocumentReferencedRejected      ErrorCode = "AP3012" // Referenced document is rejected or failed
	ErrDocumentReferencedInsufficient  ErrorCode = "AP3013" // Referenced document remainder less than credit note value
	ErrENCFInProcess                   ErrorCode = "AP3011" // ENCF document is in process
	ErrENCFUsed                        ErrorCode = "AP3001" // ENCF document was used
	ErrExtraInfoNotFound               ErrorCode = "AP3007" // Extra info reference not found
	ErrPrimaryUsed                     ErrorCode = "AP3005" // Primary ID was used
	ErrSchemaNotValid                  ErrorCode = "AP3008" // XML Schema is not valid
	ErrSenderRNCNotFound               ErrorCode = "AP3004" // Sender RNC not found
	ErrSignatureNotFound               ErrorCode = "AP3003" // Signature not found
	ErrTrackIDNotFound                 ErrorCode = "AP3009" // Track ID not found
	ErrXMLNotFound                     ErrorCode = "AP3002" // XML not found
)

// --- Acknowledgment errors ---

const (
	ErrAcknowledgmentNotFound ErrorCode = "AP7001" // Acknowledgment not found
)

// --- Set test errors ---

const (
	ErrSetTestDuplicatedRetryNumber ErrorCode = "AP4005" // Set test duplicated retry number
	ErrSetTestIDRequired            ErrorCode = "AP4002" // Set test ID is required
	ErrSetTestInProgress            ErrorCode = "AP4001" // Set test already in progress
	ErrSetTestNotFound              ErrorCode = "AP4004" // Set test not found
	ErrSetTestRetryNumberNotInteger ErrorCode = "AP4003" // Set test retry number is not integer
)

// --- General errors ---

const (
	ErrDateRangeTooLong ErrorCode = "AP9003" // Date range too long (max 90 days)
	ErrInvalidRequest   ErrorCode = "AP9002" // Invalid request structure
	ErrUnexpectedError  ErrorCode = "EPR500" // Unexpected error
)

// --- Common validation errors ---

const (
	ErrENCFPattern         ErrorCode = "AP10104" // ENCF invalid format (must start with E31-E47, 13 digits)
	ErrIDPattern           ErrorCode = "AP10100" // ID must be valid ULID format (26 chars)
	ErrInvalidEmailFormat  ErrorCode = "AP10101" // Invalid email format
	ErrInvalidNumberFormat ErrorCode = "AP10102" // Invalid string number format
	ErrMunicipalityEnum    ErrorCode = "AP10105" // Invalid municipality
	ErrPhoneNumberPattern  ErrorCode = "AP10107" // Phone number invalid format (XXX-XXX-XXXX)
	ErrProvinceEnum        ErrorCode = "AP10106" // Invalid province
	ErrRNCPattern          ErrorCode = "AP10103" // RNC must be 9-11 digits
)

// --- Type validation errors ---

const (
	ErrTypeArray                              ErrorCode = "AP10008" // Value must be an array
	ErrTypeArrayOrNull                        ErrorCode = "AP10024" // Value must be either an array or null
	ErrTypeBoolean                            ErrorCode = "AP10006" // Value must be a boolean
	ErrTypeBooleanOrArray                     ErrorCode = "AP10020" // Value must be either a boolean or an array
	ErrTypeBooleanArrayOrNull                 ErrorCode = "AP10043" // Value must be either a boolean, array, or null
	ErrTypeBooleanOrNull                      ErrorCode = "AP10021" // Value must be either a boolean or null
	ErrTypeBooleanOrObject                    ErrorCode = "AP10019" // Value must be either a boolean or an object
	ErrTypeBooleanObjectOrArray               ErrorCode = "AP10041" // Value must be either a boolean, object, or array
	ErrTypeBooleanObjectArrayOrNull           ErrorCode = "AP10059" // Value must be either a boolean, object, array, or null
	ErrTypeBooleanObjectOrNull                ErrorCode = "AP10042" // Value must be either a boolean, object, or null
	ErrTypeNull                               ErrorCode = "AP10009" // Value must be null
	ErrTypeNumber                             ErrorCode = "AP10005" // Value must be a number
	ErrTypeNumberOrArray                      ErrorCode = "AP10017" // Value must be either a number or an array
	ErrTypeNumberArrayOrNull                  ErrorCode = "AP10040" // Value must be either a number, array, or null
	ErrTypeNumberOrBoolean                    ErrorCode = "AP10015" // Value must be either a number or a boolean
	ErrTypeNumberBooleanOrArray               ErrorCode = "AP10036" // Value must be either a number, boolean, or array
	ErrTypeNumberBooleanArrayOrNull           ErrorCode = "AP10057" // Value must be either a number, boolean, array, or null
	ErrTypeNumberBooleanOrNull                ErrorCode = "AP10037" // Value must be either a number, boolean, or null
	ErrTypeNumberBooleanOrObject              ErrorCode = "AP10035" // Value must be either a number, boolean, or object
	ErrTypeNumberBooleanObjectOrArray         ErrorCode = "AP10055" // Value must be either a number, boolean, object, or array
	ErrTypeNumberBooleanObjectArrayOrNull     ErrorCode = "AP10065" // Value must be either a number, boolean, object, array, or null
	ErrTypeNumberBooleanObjectOrNull          ErrorCode = "AP10056" // Value must be either a number, boolean, object, or null
	ErrTypeNumberOrNull                       ErrorCode = "AP10018" // Value must be either a number or null
	ErrTypeNumberOrObject                     ErrorCode = "AP10016" // Value must be either a number or an object
	ErrTypeNumberObjectOrArray                ErrorCode = "AP10038" // Value must be either a number, object, or array
	ErrTypeNumberObjectArrayOrNull            ErrorCode = "AP10058" // Value must be either a number, object, array, or null
	ErrTypeNumberObjectOrNull                 ErrorCode = "AP10039" // Value must be either a number, object, or null
	ErrTypeObject                             ErrorCode = "AP10007" // Value must be an object
	ErrTypeObjectOrArray                      ErrorCode = "AP10022" // Value must be either an object or an array
	ErrTypeObjectArrayOrNull                  ErrorCode = "AP10044" // Value must be either an object, array, or null
	ErrTypeObjectOrNull                       ErrorCode = "AP10023" // Value must be either an object or null
	ErrTypeString                             ErrorCode = "AP10004" // Value must be a string
	ErrTypeStringOrArray                      ErrorCode = "AP10013" // Value must be either a string or an array
	ErrTypeStringArrayOrNull                  ErrorCode = "AP10034" // Value must be either a string, array, or null
	ErrTypeStringOrBoolean                    ErrorCode = "AP10011" // Value must be either a string or a boolean
	ErrTypeStringBooleanOrArray               ErrorCode = "AP10030" // Value must be either a string, boolean, or array
	ErrTypeStringBooleanArrayOrNull           ErrorCode = "AP10053" // Value must be either a string, boolean, array, or null
	ErrTypeStringBooleanOrNull                ErrorCode = "AP10031" // Value must be either a string, boolean, or null
	ErrTypeStringBooleanOrObject              ErrorCode = "AP10029" // Value must be either a string, boolean, or object
	ErrTypeStringBooleanObjectOrArray         ErrorCode = "AP10051" // Value must be either a string, boolean, object, or array
	ErrTypeStringBooleanObjectArrayOrNull     ErrorCode = "AP10064" // Value must be either a string, boolean, object, array, or null
	ErrTypeStringBooleanObjectOrNull          ErrorCode = "AP10052" // Value must be either a string, boolean, object, or null
	ErrTypeStringOrNull                       ErrorCode = "AP10014" // Value must be either a string or null
	ErrTypeStringOrNumber                     ErrorCode = "AP10010" // Value must be either a string or a number
	ErrTypeStringNumberOrArray                ErrorCode = "AP10027" // Value must be either a string, number, or array
	ErrTypeStringNumberArrayOrNull            ErrorCode = "AP10050" // Value must be either a string, number, array, or null
	ErrTypeStringNumberOrBoolean              ErrorCode = "AP10025" // Value must be either a string, number, or boolean
	ErrTypeStringNumberBooleanOrArray         ErrorCode = "AP10046" // Value must be either a string, number, boolean, or array
	ErrTypeStringNumberBooleanArrayOrNull     ErrorCode = "AP10062" // Value must be either a string, number, boolean, array, or null
	ErrTypeStringNumberBooleanOrNull          ErrorCode = "AP10047" // Value must be either a string, number, boolean, or null
	ErrTypeStringNumberBooleanOrObject        ErrorCode = "AP10045" // Value must be either a string, number, boolean, or object
	ErrTypeStringNumberBooleanObjectOrArray   ErrorCode = "AP10060" // Value must be either a string, number, boolean, object, or array
	ErrTypeStringNumberBooleanObjectArrayNull ErrorCode = "AP10066" // Value must be one of: string, number, boolean, object, array, or null
	ErrTypeStringNumberBooleanObjectOrNull    ErrorCode = "AP10061" // Value must be either a string, number, boolean, object, or null
	ErrTypeStringNumberOrNull                 ErrorCode = "AP10028" // Value must be either a string, number, or null
	ErrTypeStringNumberOrObject               ErrorCode = "AP10026" // Value must be either a string, number, or object
	ErrTypeStringNumberObjectOrArray          ErrorCode = "AP10048" // Value must be either a string, number, object, or array
	ErrTypeStringNumberObjectArrayOrNull      ErrorCode = "AP10063" // Value must be either a string, number, object, array, or null
	ErrTypeStringNumberObjectOrNull           ErrorCode = "AP10049" // Value must be either a string, number, object, or null
	ErrTypeStringOrObject                     ErrorCode = "AP10012" // Value must be either a string or an object
	ErrTypeStringObjectOrArray                ErrorCode = "AP10032" // Value must be either a string, object, or array
	ErrTypeStringObjectArrayOrNull            ErrorCode = "AP10054" // Value must be either a string, object, array, or null
	ErrTypeStringObjectOrNull                 ErrorCode = "AP10033" // Value must be either a string, object, or null
)

// --- Schema validation errors ---

const (
	ErrSchemaAdditionalProperties ErrorCode = "AP10068" // Additional properties not allowed
	ErrSchemaConst                ErrorCode = "AP10080" // Value must be equal to constant
	ErrSchemaContains             ErrorCode = "AP10081" // Value must contain value
	ErrSchemaDependencies         ErrorCode = "AP10086" // Dependency failed for property
	ErrSchemaElse                 ErrorCode = "AP10088" // Invalid value (else)
	ErrSchemaEnum                 ErrorCode = "AP10079" // Value must be one of allowed values
	ErrSchemaExclusiveMaximum     ErrorCode = "AP10076" // Value must be less than limit
	ErrSchemaExclusiveMinimum     ErrorCode = "AP10075" // Value must be greater than limit
	ErrSchemaFormat               ErrorCode = "AP10070" // Invalid format value
	ErrSchemaItems                ErrorCode = "AP10085" // Invalid item at index
	ErrSchemaMaximum              ErrorCode = "AP10074" // Value must be less than or equal to limit
	ErrSchemaMaxItems             ErrorCode = "AP10084" // Array must have at most N items
	ErrSchemaMaxLength            ErrorCode = "AP10072" // Value must be at most N characters
	ErrSchemaMinimum              ErrorCode = "AP10073" // Value must be greater than or equal to limit
	ErrSchemaMinItems             ErrorCode = "AP10083" // Array must have at least N items
	ErrSchemaMinLength            ErrorCode = "AP10071" // Value must be at least N characters
	ErrSchemaMultipleOf           ErrorCode = "AP10077" // Value must be a multiple of N
	ErrSchemaPattern              ErrorCode = "AP10078" // Value must match pattern
	ErrSchemaRequired             ErrorCode = "AP10067" // Missing required property
	ErrSchemaThen                 ErrorCode = "AP10087" // Invalid value (then)
	ErrSchemaType                 ErrorCode = "AP10069" // Invalid type value
	ErrSchemaUniqueItems          ErrorCode = "AP10082" // Array must have unique items
	ErrSchemaUnknown              ErrorCode = "AP10000" // Unknown type error
)

// --- Config validation errors ---

const (
	ErrConfigPDFNote ErrorCode = "AP11002" // Note PDF config invalid (max 500 chars, UTF-8 only)
	ErrConfigPDFType ErrorCode = "AP11001" // Type PDF config invalid (must be "generic" or "pos")
)

// --- Discounts or surcharges validation errors ---

const (
	ErrFitTypeDiscountsOrSurcharges   ErrorCode = "AP12001" // Fit type invalid (must be "D" or "R")
	ErrTypeValueDiscountsOrSurcharges ErrorCode = "AP12002" // Type value invalid (must be "%" or "$")
)

// --- Document identification (idDoc) validation errors ---

const (
	ErrCreditNoteIndicatorIdDoc ErrorCode = "AP13006" // Credit note indicator invalid (0 or 1)
	ErrIncomeTypeIdDoc          ErrorCode = "AP13002" // Income type invalid (1-6)
	ErrPaymentAccountTypeIdDoc  ErrorCode = "AP13005" // Payment account type invalid
	ErrPaymentMethodIdDoc       ErrorCode = "AP13004" // Payment method invalid (1-8)
	ErrPaymentTypeIdDoc         ErrorCode = "AP13003" // Payment type invalid (1-3)
	ErrTaxAmountIndicatorIdDoc  ErrorCode = "AP13001" // Tax amount indicator invalid (0 or 1)
)

// --- Information reference (document reference) validation errors ---

const (
	ErrModificationCodeInfoRef     ErrorCode = "AP14001" // Modification code invalid (4 or 5)
	ErrNoENCFModifiedPattern       ErrorCode = "AP14002" // NCF modified invalid (cannot modify with other credit/debit notes)
	ErrNotesModificationCodeEnum   ErrorCode = "AP14003" // Modification code invalid for notes (1-4)
)

// --- Currency validation errors ---

const (
	ErrCurrencyTypeOtherCurrency ErrorCode = "AP15001" // Currency type invalid (check documentation)
)

// --- Company sync validation errors ---

const (
	ErrCertificateContentBase64   ErrorCode = "AP16004" // Certificate content must be base64
	ErrCertificateExtensionPattern ErrorCode = "AP16003" // Certificate extension invalid (no dot)
	ErrCompanyTypeEnum            ErrorCode = "AP16001" // Company type invalid (must be "main" or "associated")
	ErrLogoBase64                 ErrorCode = "AP16002" // Logo must be base64
	ErrWebhookAuthStatusEnum      ErrorCode = "AP16006" // Webhook auth status invalid ("active" or "inactive")
	ErrWebhookCacheStatusEnum     ErrorCode = "AP16007" // Webhook cache status invalid ("active" or "inactive")
	ErrWebhookStatusEnum          ErrorCode = "AP16005" // Webhook status invalid ("active" or "inactive")
)

// IsErrorCode checks if an API error matches a specific error code.
func IsErrorCode(err error, code ErrorCode) bool {
	apiErr, ok := err.(*Error)
	if !ok {
		return false
	}
	return apiErr.Code == string(code)
}
