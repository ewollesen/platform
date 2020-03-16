package test

import (
	"math/rand"

	"github.com/tidepool-org/platform/data/types"
	dataTypesBasalTest "github.com/tidepool-org/platform/data/types/basal/test"
	"github.com/tidepool-org/platform/data/types/settings/pump"
	dataTypesTest "github.com/tidepool-org/platform/data/types/test"
	"github.com/tidepool-org/platform/pointer"
	"github.com/tidepool-org/platform/test"
)

func NewMeta() interface{} {
	return &types.Meta{
		Type: "pumpSettings",
	}
}

func NewManufacturer(minimumLength int, maximumLength int) string {
	return test.RandomStringFromRange(minimumLength, maximumLength)
}

func NewManufacturers(minimumLength int, maximumLength int) []string {
	result := make([]string, minimumLength+rand.Intn(maximumLength-minimumLength+1))
	for index := range result {
		result[index] = NewManufacturer(1, 100)
	}
	return result
}

func NewPump(unitsBloodGlucose *string) *pump.Pump {
	scheduleName := dataTypesBasalTest.NewScheduleName()
	datum := pump.New()
	datum.Base = *dataTypesTest.NewBase()
	datum.Type = "pumpSettings"
	datum.ActiveScheduleName = pointer.FromString(scheduleName)
	datum.Basal = NewBasal()
	datum.BasalRateSchedules = pump.NewBasalRateStartArrayMap()
	datum.BasalRateSchedules.Set(scheduleName, NewBasalRateStartArray())
	datum.BloodGlucoseTargetSchedules = pump.NewBloodGlucoseTargetStartArrayMap()
	datum.BloodGlucoseTargetSchedules.Set(scheduleName, NewBloodGlucoseTargetStartArray(unitsBloodGlucose))
	datum.Bolus = NewBolus()
	datum.CarbohydrateRatioSchedules = pump.NewCarbohydrateRatioStartArrayMap()
	datum.CarbohydrateRatioSchedules.Set(scheduleName, NewCarbohydrateRatioStartArray())
	datum.Display = NewDisplay()
	datum.InsulinSensitivitySchedules = pump.NewInsulinSensitivityStartArrayMap()
	datum.InsulinSensitivitySchedules.Set(scheduleName, NewInsulinSensitivityStartArray(unitsBloodGlucose))
	datum.Manufacturers = pointer.FromStringArray(NewManufacturers(1, 10))
	datum.Model = pointer.FromString(test.RandomStringFromRange(1, 100))
	datum.SerialNumber = pointer.FromString(test.RandomStringFromRange(1, 100))
	datum.Units = NewUnits(unitsBloodGlucose)
	return datum
}

func ClonePump(datum *pump.Pump) *pump.Pump {
	if datum == nil {
		return nil
	}
	clone := pump.New()
	clone.Base = *dataTypesTest.CloneBase(&datum.Base)
	clone.ActiveScheduleName = pointer.CloneString(datum.ActiveScheduleName)
	clone.Basal = CloneBasal(datum.Basal)
	clone.BasalRateSchedule = CloneBasalRateStartArray(datum.BasalRateSchedule)
	clone.BasalRateSchedules = CloneBasalRateStartArrayMap(datum.BasalRateSchedules)
	clone.BloodGlucoseTargetSchedule = CloneBloodGlucoseTargetStartArray(datum.BloodGlucoseTargetSchedule)
	clone.BloodGlucoseTargetSchedules = CloneBloodGlucoseTargetStartArrayMap(datum.BloodGlucoseTargetSchedules)
	clone.Bolus = CloneBolus(datum.Bolus)
	clone.CarbohydrateRatioSchedule = CloneCarbohydrateRatioStartArray(datum.CarbohydrateRatioSchedule)
	clone.CarbohydrateRatioSchedules = CloneCarbohydrateRatioStartArrayMap(datum.CarbohydrateRatioSchedules)
	clone.Display = CloneDisplay(datum.Display)
	clone.InsulinSensitivitySchedule = CloneInsulinSensitivityStartArray(datum.InsulinSensitivitySchedule)
	clone.InsulinSensitivitySchedules = CloneInsulinSensitivityStartArrayMap(datum.InsulinSensitivitySchedules)
	clone.Manufacturers = pointer.CloneStringArray(datum.Manufacturers)
	clone.Model = pointer.CloneString(datum.Model)
	clone.SerialNumber = pointer.CloneString(datum.SerialNumber)
	clone.Units = CloneUnits(datum.Units)
	return clone
}
