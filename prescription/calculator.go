package prescription

import (
	"github.com/tidepool-org/platform/structure"
)

const (
	MethodWeight                  = "weight"
	MethodTotalDailyDose          = "totalDailyDose"
	MethodTotalDailyDoseAndWeight = "totalDailyDoseAndWeight"

	TotalDailyDoseScaleFactorMinimum = 0
	TotalDailyDoseScaleFactorMaximum = 1
)

type Calculator struct {
	Method                        string   `json:"method,omitempty" bson:"method,omitempty"`
	RecommendedBasalRate          *float64 `json:"recommendedBasalRate,omitempty" bson:"recommendedBasalRate,omitempty"`
	RecommendedCarbohydrateRatio  *float64 `json:"recommendedCarbohydrateRatio,omitempty" bson:"recommendedCarbohydrateRatio,omitempty"`
	RecommendedInsulinSensitivity *float64 `json:"recommendedInsulinSensitivity,omitempty" bson:"recommendedInsulinSensitivity,omitempty"`
	TotalDailyDose                *float64 `json:"totalDailyDose,omitempty" bson:"totalDailyDose,omitempty"`
	TotalDailyDoseScaleFactor     *float64 `json:"totalDailyDoseScaleFactor,omitempty" bson:"totalDailyDoseScaleFactor,omitempty"`
	Weight                        *float64 `json:"weight,omitempty" bson:"weight,omitempty"`
	WeightUnits                   *string  `json:"weightUnits,omitempty" bson:"weightUnits,omitempty"`
}

func AllowedCalculatorWeightUnits() []string {
	return []string{
		UnitKg,
		UnitLbs,
	}
}

func AllowedCalculatorMethods() []string {
	return []string{
		MethodTotalDailyDose,
		MethodTotalDailyDoseAndWeight,
		MethodWeight,
	}
}

func (c *Calculator) ValidateWeightInputs(validator structure.Validator) {
	validator.Float64("weight", c.Weight).GreaterThan(0)
	validator.String("weightUnits", c.WeightUnits).OneOf(AllowedCalculatorWeightUnits()...)
}

func (c *Calculator) ValidateTotalDailyDoseInputs(validator structure.Validator) {
	validator.Float64("totalDailyDose", c.TotalDailyDose).GreaterThan(0)
	validator.Float64("totalDailyDoseScaleFactor", c.TotalDailyDoseScaleFactor).InRange(TotalDailyDoseScaleFactorMinimum, TotalDailyDoseScaleFactorMaximum)
}

func (c *Calculator) Validate(validator structure.Validator) {
	if c.Method != "" {
		validator.String("method", &c.Method).OneOf(AllowedCalculatorMethods()...)
		if c.Method == MethodTotalDailyDose {
			c.ValidateTotalDailyDoseInputs(validator)
		}
		if c.Method == MethodWeight {
			c.ValidateWeightInputs(validator)
		}
		if c.Method == MethodTotalDailyDoseAndWeight {
			c.ValidateTotalDailyDoseInputs(validator)
			c.ValidateWeightInputs(validator)
		}
		validator.Float64("recommendedBasalRate", c.RecommendedBasalRate).GreaterThan(0)
		validator.Float64("recommendedCarbohydrateRatio", c.RecommendedCarbohydrateRatio).GreaterThan(0)
		validator.Float64("recommendedInsulinSensitivity", c.RecommendedInsulinSensitivity).GreaterThan(0)
	}
}
