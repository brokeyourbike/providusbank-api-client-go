package providusbank

type Code string

func (c Code) IsSuccesful() bool {
	return c == CodeCompleted
}

func (c Code) IsFailed() bool {
	switch c {
	case CodeInvalidSender,
		CodeDoNotHonor,
		CodeInvalidAccount,
		CodeAccountNameMismatch,
		CodeInvalidAmount,
		CodeUnknownBankCode,
		CodeNoActionTaken,
		CodeDuplicateRecord,
		CodeFormatError,
		CodeTransferNotSuccessful,
		CodeServiceUnavailable,
		CodeTransactionNotPermittedToSender,
		CodeTransactionNotPermittedToChannel,
		CodeTransferLimitExceeded,
		CodeSecurityViolation,
		CodeWithdrawalFrequencyExceeded,
		CodeDebitAccountInvalid,
		CodeDuplicateReference,
		CodeMethodNotAllowed,
		CodeCreditAccountNotPermitted,
		CodeRecipientBankNotAvailable,
		CodeRoutingError,
		CodeDuplicateTransaction,
		CodeSustemMalfunction,
		CodeDestinationResponseTimeout,
		CodeTransactionNotExists,
		CodeCreditAccountInvalid,
		CodeCreditAccountDormant,
		CodeInsufficientBalance,
		CodeCurrencyMismatch,
		CodeOtpFailInvalidAccountNumber,
		CodeCustomerNotEnrolled,
		CodeCustomerNotActive,
		CodeCustomerNotValidated,
		CodeCustomerAlreadyEnrolled,
		CodeCustomerAlreadyValidated,
		CodeOtpMismatch,
		CodeCbaSystemError:
		return true
	}
	return false
}

const (
	CodeCompleted                        Code = "00"
	CodeDormantAccount                   Code = "06"
	CodeTransactionNotExists             Code = "01"
	CodeInvalidSender                    Code = "03"
	CodeDoNotHonor                       Code = "05"
	CodeInvalidAccount                   Code = "07"
	CodeAccountNameMismatch              Code = "08"
	CodeInProgress                       Code = "09"
	CodeReversalNotSuccessful            Code = "11"
	CodeInvalidTransaction               Code = "12"
	CodeInvalidAmount                    Code = "13"
	CodeInvalidBatchNumber               Code = "14"
	CodeInvalidSession                   Code = "15"
	CodeUnknownBankCode                  Code = "16"
	CodeInvalidChannel                   Code = "17"
	CodeInvalidMethodCall                Code = "18"
	CodeNoActionTaken                    Code = "21"
	CodeUnableToLocateRecord             Code = "25"
	CodeReversalCompleted                Code = "25"
	CodeDuplicateRecord                  Code = "26"
	CodeFormatError                      Code = "30"
	CodeTransferNotSuccessful            Code = "32"
	CodeSuspectedFraud                   Code = "34"
	CodeContactSendingBank               Code = "35"
	CodeTransferCompleted                Code = "36"
	CodeInsufficientFunds                Code = "51"
	CodeTransactionNotPermittedToSender  Code = "57"
	CodeTransactionNotPermittedToChannel Code = "58"
	CodeServiceUnavailable               Code = "505"
	CodeTransferLimitExceeded            Code = "61"
	CodeSecurityViolation                Code = "63"
	CodeWithdrawalFrequencyExceeded      Code = "65"
	CodeResponseReceivedTooLate          Code = "68"
	CodeCustomerDetailsNotValidated      Code = "69"
	CodeNotificationNotReceived          Code = "70"
	CodeDebitAccountInvalid              Code = "7701"
	CodeCreditAccountInvalid             Code = "7702"
	CodeCreditAccountDormant             Code = "7703"
	CodeInsufficientBalance              Code = "7704"
	CodeInvalidAmount2                   Code = "7706"
	CodeCurrencyMismatch                 Code = "7708"
	CodeDuplicateReference               Code = "7709"
	CodeCustomerNotValidated             Code = "7710"
	CodeCustomerNotActive                Code = "7711"
	CodeCustomerAlreadyEnrolled          Code = "7712"
	CodeOtpFailInvalidAccountNumber      Code = "7713"
	CodeCustomerNotEnrolled              Code = "7714"
	CodeOtpMismatch                      Code = "7715"
	CodeCustomerAlreadyValidated         Code = "7716"
	CodeCbaSystemError                   Code = "7799"
	CodeAuthFailed                       Code = "8004"
	CodeMethodNotAllowed                 Code = "8005"
	CodeNoConenction                     Code = "8803"
	CodeCreditAccountNotPermitted        Code = "8888"
	CodeRecipientBankNotAvailable        Code = "91"
	CodeRoutingError                     Code = "92"
	CodeDuplicateTransaction             Code = "94"
	CodeSustemMalfunction                Code = "96"
	CodeDestinationResponseTimeout       Code = "97"
	CodeUnknownNfpResponse               Code = "909"
	CodeNullNfpResponse                  Code = "919"
	CodeLocalTimeout                     Code = "999"
	CodeAuthFailed2                      Code = "9999"
)
