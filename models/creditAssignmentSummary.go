package models

type CreditAssignmentSummary struct {
	AssignmentTotal      int32   `bson:"assigment_total" json:"assigment_total"`
	AssignmentSuccess    int32   `bson:"assigment_success" json:"assigment_success"`
	AssignmentWrong      int32   `bson:"assigment_wrong" json:"assigment_wrong"`
	InvestmentTotal      float64 `bson:"investment_total" json:"investment_total"`
	InvestmentAvgSuccess float64 `bson:"investment_avg_success" json:"investment_avg_success"`
	InvestmentAvgWrong   float64 `bson:"investment_avg_wrong" json:"investment_avg_wrong"`
	IsActive             bool    `bson:"is_active" json:"is_active"`
}
