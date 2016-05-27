package calculator

import (
	"github.com/tidepool-org/platform/data"
	"github.com/tidepool-org/platform/data/types/common/bloodglucose"
)

//NOTE: this is the matrix we are working to. Only animas at this stage
// animas: {`target`, `range`}
// insulet: {`target`, `high`}
// medtronic: {`low`, `high`}
// tandem: {`target`}

type BloodGlucoseTarget struct {
	Target *float64 `json:"target,omitempty" bson:"target,omitempty"`
	Range  *int     `json:"range,omitempty" bson:"range,omitempty"`
}

func NewBloodGlucoseTarget() *BloodGlucoseTarget {
	return &BloodGlucoseTarget{}
}

func (b *BloodGlucoseTarget) Parse(parser data.ObjectParser) {
	b.Target = parser.ParseFloat("target")
	b.Range = parser.ParseInteger("range")
}

func (b *BloodGlucoseTarget) Validate(validator data.Validator, units *string) {
	switch *units {
	case bloodglucose.Mmoll, bloodglucose.MmolL:
		validator.ValidateFloat("target", b.Target).InRange(bloodglucose.AllowedMmolLRange())
	default:
		validator.ValidateFloat("target", b.Target).InRange(bloodglucose.AllowedMgdLRange())
	}

	validator.ValidateInteger("range", b.Range).InRange(0, 50)
}

func (b *BloodGlucoseTarget) Normalize(normalizer data.Normalizer, units *string) {
	if b.Target != nil {
		b.Target = normalizer.NormalizeBloodGlucose("target", units).NormalizeValue(b.Target)
	}
}

func ParseBloodGlucoseTarget(parser data.ObjectParser) *BloodGlucoseTarget {
	var bloodGlucoseTarget *BloodGlucoseTarget
	if parser.Object() != nil {
		bloodGlucoseTarget = NewBloodGlucoseTarget()
		bloodGlucoseTarget.Parse(parser)
	}
	return bloodGlucoseTarget
}
