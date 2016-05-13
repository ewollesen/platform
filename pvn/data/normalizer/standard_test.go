package normalizer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/tidepool-org/platform/pvn/data"
	"github.com/tidepool-org/platform/pvn/data/context"
	"github.com/tidepool-org/platform/pvn/data/normalizer"
)

type TestDatum struct{}

func (t *TestDatum) Parse(parser data.ObjectParser)       {}
func (t *TestDatum) Validate(validator data.Validator)    {}
func (t *TestDatum) Normalize(normalizer data.Normalizer) {}

var _ = Describe("Standard", func() {

	It("New returns an error if context is nil", func() {
		standard, err := normalizer.NewStandard(nil)
		Expect(standard).To(BeNil())
		Expect(err).To(HaveOccurred())
	})

	Describe("new normalizer", func() {
		var standard *normalizer.Standard

		BeforeEach(func() {
			var err error
			standard, err = normalizer.NewStandard(context.NewStandard())
			Expect(err).ToNot(HaveOccurred())
		})

		It("exists", func() {
			Expect(standard).ToNot(BeNil())
		})

		It("has a contained Data that is empty", func() {
			Expect(standard.Data()).To(BeEmpty())
		})

		It("ignores sending a nil datum to AppendDatum", func() {
			standard.AppendDatum(nil)
			Expect(standard.Data()).To(BeEmpty())
		})

		Describe("AddDatum with a first datum", func() {
			var firstDatum *TestDatum

			BeforeEach(func() {
				firstDatum = &TestDatum{}
				standard.AppendDatum(firstDatum)
			})

			It("has data", func() {
				Expect(standard.Data()).ToNot(BeEmpty())
			})

			It("has the datum", func() {
				Expect(standard.Data()).To(ConsistOf(firstDatum))
			})

			Describe("and AddDatum with a second data", func() {
				var secondDatum *TestDatum

				BeforeEach(func() {
					secondDatum = &TestDatum{}
					standard.AppendDatum(secondDatum)
				})

				It("has data", func() {
					Expect(standard.Data()).ToNot(BeEmpty())
				})

				It("has both data", func() {
					Expect(standard.Data()).To(ConsistOf(firstDatum, secondDatum))
				})
			})
		})

		Describe("NewChildNormalizer", func() {
			var child data.Normalizer

			BeforeEach(func() {
				child = standard.NewChildNormalizer("child")
			})

			It("exists", func() {
				Expect(child).ToNot(BeNil())
			})

			Describe("AppendDatum with a first error", func() {
				var firstDatum *TestDatum

				BeforeEach(func() {
					firstDatum = &TestDatum{}
					child.AppendDatum(firstDatum)
				})

				It("has data", func() {
					Expect(standard.Data()).ToNot(BeEmpty())
				})

				It("has the data", func() {
					Expect(standard.Data()).To(ConsistOf(firstDatum))
				})

				Describe("and AppendDatum with a second error to the parent context", func() {
					var secondDatum *TestDatum

					BeforeEach(func() {
						secondDatum = &TestDatum{}
						standard.AppendDatum(secondDatum)
					})

					It("has data", func() {
						Expect(standard.Data()).ToNot(BeEmpty())
					})

					It("has both data", func() {
						Expect(standard.Data()).To(ConsistOf(firstDatum, secondDatum))
					})
				})
			})
		})
	})
})
