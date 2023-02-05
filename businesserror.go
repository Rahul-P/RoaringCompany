package roaringcompany

import (
	"errors"
	"strings"
)

const DefaultErrorMessage string = "Apologies! An error has been encountered. Please try again in a while. Rest assured our army of monkeys is on this issue"

var (
	InvalidArgument            = RaiseError("Invalid arguments supplied")
	ErrUnknown                 = RaiseError(DefaultErrorMessage)
	DatabaseIssue              = RaiseError("Unable to interact with database")
	NoRecordFound              = RaiseError("Record not found")
	NoRecordsFound             = RaiseError("Records not found")
	DataNotFoundToUpdate       = RaiseError("Data not found to update")
	RecordExists               = RaiseError("Record exists. Cannot proceed. Please update or delete exisitng record")
	DataChangedBeforeOperation = RaiseError("Data has been changed before this operation could complete, please check changes by reloading data and try again")
	ErrPhonenumNotFound        = RaiseError("Phone num is not found")
	CRAP                       = RaiseError("CRAAAPPP")
)

func RaiseError(message string) error {
	if len(strings.TrimSpace(message)) == 0 {
		return errors.New(DefaultErrorMessage)
	}
	return errors.New(message)
}
