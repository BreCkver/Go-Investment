package services

import (
	"testing"

	"github.com/BreCkver/Go-Investment/data"
)

func TestAssign_Success(t *testing.T) {
	service := NewCreditAssignmentService(nil)

	tables := []struct {
		invest int32
	}{
		{1500},
		{300},
		{500},
		{700},
		{900},
	}

	for _, item := range tables {
		typeCount300, typeCount500, typeCount700, err := service.Assign(item.invest)
		if err != nil {
			t.Errorf("Error in Assign: %s", err.Error())
		} else {
			sum := (typeCount300 * 300) + (typeCount500 * 500) + (typeCount700 * 700)
			if sum != item.invest {
				t.Errorf("Sum was incorrect, got: %d, expected: %d", sum, item.invest)
			}
		}
	}
}

func TestAssign_Error(t *testing.T) {
	service := NewCreditAssignmentService(nil)

	tables := []struct {
		invest int32
	}{
		{400},
		{100},
		{50},
		{380},
		{8570},
	}

	for _, item := range tables {
		_, _, _, err := service.Assign(item.invest)
		if err == nil {
			t.Errorf("Investment invalid: %d", item.invest)
		}
	}

}

func TestGetStatistics_Success(t *testing.T) {
	var mock = data.NewCreditAssignerDataMock(true)
	service := NewCreditAssignmentService(mock)

	summary, err := service.GetStatistics()
	if err != nil {
		t.Errorf("Summary exception: %s", err.Error())
	}

	if summary.AssignmentTotal <= 0 {
		t.Errorf("Summary invalid: %d", summary.AssignmentTotal)
	}
}

func TestGetStatistics_Error(t *testing.T) {
	var mock = data.NewCreditAssignerDataMock(false)
	service := NewCreditAssignmentService(mock)

	_, err := service.GetStatistics()
	if err == nil {
		t.Error("GetStatistics retrieved successfully in error test")
	}
}

func TestSaveStatisticsOk_Success(t *testing.T) {
	var mock = data.NewCreditAssignerDataMock(true)
	service := NewCreditAssignmentService(mock)

	err := service.SaveStatistics(1500, true)
	if err != nil {
		t.Errorf("SaveStatistics exception: %s", err.Error())
	}
}

func TestSaveStatisticsWrong_Success(t *testing.T) {
	var mock = data.NewCreditAssignerDataMock(true)
	service := NewCreditAssignmentService(mock)

	err := service.SaveStatistics(1500, false)
	if err != nil {
		t.Errorf("SaveStatistics exception: %s", err.Error())
	}
}

func TestSaveStatistics_Error(t *testing.T) {
	var mock = data.NewCreditAssignerDataMock(false)
	service := NewCreditAssignmentService(mock)

	err := service.SaveStatistics(1500, true)
	if err == nil {
		t.Error("SaveStatistics retrieved successfully in error test")
	}
}
