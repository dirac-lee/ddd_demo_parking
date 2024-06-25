package E

import (
	"fmt"
	"testing"

	"code.byted.org/oec/status_code/Status"
)

func TestUsecase(t *testing.T) {
	usecase_simple()
	usecase_with_message()
	usecase_nil_exception()
}

func usecase_simple() {
	Try(func() {
		Throw(New(Status.AccountEmailsIllegalErr))
	}).Catch(func(e *Exception) {
		fmt.Printf(e.Sprintf("usecase_simple failed"))
	})
}

func usecase_with_message() {
	Try(func() {
		media()
	}).Catch(func(e *Exception) {
		e.AppendDetail("usecase catch the exception")
		fmt.Printf(e.Sprintf("[usecase_with_message] fialed"))
	})
}

func media() {
	Try(func() {
		inner()
	}).Catch(func(e *Exception) {
		e.AppendDetail("media catch then throw")
		Throw(e)
	})
}

func inner() {
	Throw(New(Status.AccountEmailsIllegalErr, WithDetail("cause by inner")))
}

func usecase_nil_exception() {
	Try(func() {
		Throw(nil)
	}).Catch(func(e *Exception) {
		e.AppendDetail("usecase catch the exception")
		fmt.Printf(e.Sprintf("[usecase_nil_exception] fialed"))
	})
}
