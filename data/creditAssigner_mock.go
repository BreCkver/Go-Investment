package data

import (
	"errors"

	"github.com/BreCkver/Go-Investment/models"
)

type CreditAssignerDataMock struct {
	successfully bool
}

func NewCreditAssignerDataMock(successfully bool) *CreditAssignerDataMock {
	return &CreditAssignerDataMock{
		successfully: successfully,
	}
}

func (d *CreditAssignerDataMock) CreditAssignmentSummarySave(summary *models.CreditAssignmentSummary) (string, error) {
	if d.successfully {
		return "", nil
	} else {
		return "", errors.New("Invalid CreditAssignmentSummarySave")
	}

}

func (d *CreditAssignerDataMock) GetLastCreditAssignmentSummary() (*models.CreditAssignmentSummary, error) {

	summary := models.CreditAssignmentSummary{
		AssignmentTotal:      1,
		AssignmentWrong:      0,
		AssignmentSuccess:    1,
		InvestmentTotal:      100,
		InvestmentAvgSuccess: 100,
		InvestmentAvgWrong:   0,
	}

	if d.successfully {
		return &summary, nil
	} else {
		return nil, errors.New("Invalid GetLastCreditAssignmentSummary")
	}
}

func (d *CreditAssignerDataMock) UpdateLastCreditAssignmentSummary() error {

	if d.successfully {
		return nil
	} else {
		return errors.New("Invalid GetLastCreditAssignmentSummary")
	}
}
