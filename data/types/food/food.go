package food

import (
	"github.com/tidepool-org/platform/data"
	"github.com/tidepool-org/platform/data/types"
)

type Food struct {
	types.Base `bson:",inline"`

	Nutrition *Nutrition `json:"nutrition,omitempty" bson:"nutrition,omitempty"`
}

func Type() string {
	return "food"
}

func NewDatum() data.Datum {
	return New()
}

func New() *Food {
	return &Food{}
}

func Init() *Food {
	food := New()
	food.Init()
	return food
}

func (f *Food) Init() {
	f.Base.Init()
	f.Type = Type()

	f.Nutrition = nil
}

func (f *Food) Parse(parser data.ObjectParser) error {
	parser.SetMeta(f.Meta())

	if err := f.Base.Parse(parser); err != nil {
		return err
	}

	f.Nutrition = ParseNutrition(parser.NewChildObjectParser("nutrition"))

	return nil
}

func (f *Food) Validate(validator data.Validator) error {
	validator.SetMeta(f.Meta())

	if err := f.Base.Validate(validator); err != nil {
		return err
	}

	validator.ValidateString("type", &f.Type).EqualTo(Type())
	if f.Nutrition != nil {
		f.Nutrition.Validate(validator.NewChildValidator("nutrition"))
	}

	return nil
}

func (f *Food) Normalize(normalizer data.Normalizer) error {
	normalizer.SetMeta(f.Meta())

	if err := f.Base.Normalize(normalizer); err != nil {
		return err
	}

	if f.Nutrition != nil {
		f.Nutrition.Normalize(normalizer.NewChildNormalizer("nutrition"))
	}

	return nil
}