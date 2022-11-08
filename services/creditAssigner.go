package services

import (
	"errors"
	"log"
	"math"

	"github.com/BreCkver/Go-Investment/data"
	"github.com/BreCkver/Go-Investment/models"
)

const (
	threeHundred = 300
	fiveHundred  = 500
	sevenHundred = 700
)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type CreditAssignerService struct {
	da *data.CreditAssignerData
}

func NewCreditAssignmentService(data *data.CreditAssignerData) *CreditAssignerService {
	return &CreditAssignerService{
		da: data,
	}
}

func (s *CreditAssignerService) Assign(investment int32) (typeCount300, typeCount500, typeCount700 int32, errG error) {

	var err error
	var isValid bool

	typeCount300, typeCount500, typeCount700 = 0, 0, 0

	if investment <= 10 {
		errG = errors.New("Invalid param")
	}

	for investment > 10 {
		isValid = false
		_, err = isValidDecrease(investment, fiveHundred)
		if err == nil {
			investment -= fiveHundred
			typeCount500 += 1
			isValid = true
		}

		_, err = isValidDecrease(investment, threeHundred)
		if err == nil {
			investment -= threeHundred
			typeCount300 += 1
			isValid = true
		}

		_, err = isValidDecrease(investment, sevenHundred)
		if err == nil {
			investment -= sevenHundred
			typeCount700 += 1
			isValid = true
		}

		if !isValid {
			errG = errors.New("Invalid param")
			break
		}
	}

	return
}

func isValidDecrease(inv int32, value int32) (int32, error) {
	var err error
	if inv > 0 && inv >= value {
		inv -= value
		if isMod(inv, fiveHundred) || isMod(inv, sevenHundred) || isMod(inv, threeHundred) {
			err = nil
		} else {
			err = errors.New("Invalid result")
		}

	} else {
		err = errors.New("Invalid value")
	}

	return inv, err
}

func isMod(inv int32, value int32) bool {
	return inv%value == 0
}

func (s *CreditAssignerService) SaveStatistics(investment int32, successfully bool) {

	summaryLast, err := s.GetStatistics()
	if err != nil {
		log.Fatal("Error GetStatistics: ", err)
	}

	summaryLast.AssignmentTotal += 1
	summaryLast.IsActive = true
	summaryLast.InvestmentTotal += float64(investment)
	if successfully {
		summaryLast.AssignmentSuccess += 1
		summaryLast.InvestmentAvgSuccess = (summaryLast.InvestmentTotal / float64(summaryLast.AssignmentSuccess))
	} else {
		summaryLast.AssignmentWrong += 1
		summaryLast.InvestmentAvgWrong = (summaryLast.InvestmentTotal / float64(summaryLast.AssignmentWrong))
	}

	err = s.da.UpdateLastCreditAssignmentSummary()
	if err != nil {
		log.Fatal("Error UpdateLastCreditAssignmentSummary: ", err)
	}

	_, err = s.da.CreditAssignmentSummarySave(&summaryLast)
	if err != nil {
		log.Fatal("Error CreditAssignmentSummarySave: ", err)
	}
}

func (s *CreditAssignerService) GetStatistics() (models.CreditAssignmentSummary, error) {

	summaryLast, err := s.da.GetLastCreditAssignmentSummary()
	if err != nil {
		log.Fatal("Error GetLastCreditAssignmentSummary: ", err)
		return summaryLast, err

	} else {
		summaryLast.InvestmentTotal = (math.Round(summaryLast.InvestmentTotal*100) / 100)
		summaryLast.InvestmentAvgSuccess = (math.Round(summaryLast.InvestmentAvgSuccess*100) / 100)
		summaryLast.InvestmentAvgWrong = (math.Round(summaryLast.InvestmentAvgWrong*100) / 100)
		return summaryLast, nil
	}
}
