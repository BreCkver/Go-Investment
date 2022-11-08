package models

type CreditAssignment struct {
	CreditType300 int32 `json:"credit_type_300"`
	CreditType500 int32 `bson:"credit_type_500" json:"credit_type_500"`
	CreditType700 int32 `bson:"credit_type_700" json:"credit_type_700"`
}

func NewCreditAssignment(creditType300, creditType500, creditType700 int32) *CreditAssignment {
	return &CreditAssignment{
		CreditType300: creditType300,
		CreditType500: creditType500,
		CreditType700: creditType700,
	}
}
