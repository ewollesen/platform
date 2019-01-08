package association

import (
	"strconv"

	"github.com/tidepool-org/platform/data"
	"github.com/tidepool-org/platform/net"
	"github.com/tidepool-org/platform/structure"
	structureValidator "github.com/tidepool-org/platform/structure/validator"
)

const (
	AssociationArrayLengthMaximum = 100
	ReasonLengthMaximum           = 1000
	TypeDatum                     = "datum"
	TypeURL                       = "url"
)

func Types() []string {
	return []string{
		TypeDatum,
		TypeURL,
	}
}

type Association struct {
	ID     *string `json:"id,omitempty" bson:"id,omitempty"`
	Reason *string `json:"reason,omitempty" bson:"reason,omitempty"`
	Type   *string `json:"type,omitempty" bson:"type,omitempty"`
	URL    *string `json:"url,omitempty" bson:"url,omitempty"`
}

func ParseAssociation(parser structure.ObjectParser) *Association {
	if !parser.Exists() {
		return nil
	}
	datum := NewAssociation()
	parser.Parse(datum)
	return datum
}

func NewAssociation() *Association {
	return &Association{}
}

func (a *Association) Parse(parser structure.ObjectParser) {
	a.ID = parser.String("id")
	a.Reason = parser.String("reason")
	a.Type = parser.String("type")
	a.URL = parser.String("url")
}

func (a *Association) Validate(validator structure.Validator) {
	if a.Type != nil {
		switch *a.Type {
		case TypeDatum:
			validator.String("id", a.ID).Exists().Using(data.IDValidator)
		case TypeURL:
			validator.String("id", a.ID).NotExists()
		}
	}
	validator.String("reason", a.Reason).NotEmpty().LengthLessThanOrEqualTo(ReasonLengthMaximum)
	validator.String("type", a.Type).Exists().OneOf(Types()...)
	if a.Type != nil {
		switch *a.Type {
		case TypeDatum:
			validator.String("url", a.URL).NotExists()
		case TypeURL:
			validator.String("url", a.URL).Exists().Using(net.URLValidator)
		}
	}
}

func (a *Association) Normalize(normalizer data.Normalizer) {}

type AssociationArray []*Association

func ParseAssociationArray(parser structure.ArrayParser) *AssociationArray {
	if !parser.Exists() {
		return nil
	}
	datum := NewAssociationArray()
	parser.Parse(datum)
	return datum
}

func NewAssociationArray() *AssociationArray {
	return &AssociationArray{}
}

func (a *AssociationArray) Parse(parser structure.ArrayParser) {
	for _, reference := range parser.References() {
		*a = append(*a, ParseAssociation(parser.WithReferenceObjectParser(reference)))
	}
}

func (a *AssociationArray) Validate(validator structure.Validator) {
	if length := len(*a); length == 0 {
		validator.ReportError(structureValidator.ErrorValueEmpty())
	} else if length > AssociationArrayLengthMaximum {
		validator.ReportError(structureValidator.ErrorLengthNotLessThanOrEqualTo(length, AssociationArrayLengthMaximum))
	}

	for index, datum := range *a {
		if datumValidator := validator.WithReference(strconv.Itoa(index)); datum != nil {
			datum.Validate(datumValidator)
		} else {
			datumValidator.ReportError(structureValidator.ErrorValueNotExists())
		}
	}
}

func (a *AssociationArray) Normalize(normalizer data.Normalizer) {
	for index, datum := range *a {
		if datum != nil {
			datum.Normalize(normalizer.WithReference(strconv.Itoa(index)))
		}
	}
}
