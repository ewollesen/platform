package base_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/tidepool-org/platform/errors"
	testErrors "github.com/tidepool-org/platform/errors/test"
	"github.com/tidepool-org/platform/pointer"
	"github.com/tidepool-org/platform/structure"
	structureBase "github.com/tidepool-org/platform/structure/base"
	testStructure "github.com/tidepool-org/platform/structure/test"
)

var _ = Describe("Base", func() {
	Context("New", func() {
		It("returns successfully", func() {
			Expect(structureBase.New()).ToNot(BeNil())
		})
	})

	Context("with source, meta, and new base", func() {
		var src *testStructure.Source
		var meta interface{}
		var base *structureBase.Base

		BeforeEach(func() {
			src = testStructure.NewSource()
			src.ParameterOutput = pointer.String(testErrors.NewSourceParameter())
			src.PointerOutput = pointer.String(testErrors.NewSourcePointer())
			meta = testErrors.NewMeta()
			base = structureBase.New()
			Expect(base).ToNot(BeNil())
		})

		AfterEach(func() {
			src.Expectations()
		})

		Context("Error", func() {
			It("returns nil if no error", func() {
				Expect(base.Error()).ToNot(HaveOccurred())
			})

			It("returns errors if any errors", func() {
				err1 := testErrors.NewError()
				base.ReportError(err1)
				err2 := testErrors.NewError()
				base.ReportError(err2)
				err3 := testErrors.NewError()
				base.ReportError(err3)
				Expect(base.Error()).To(Equal(errors.Append(err1, err2, err3)))
			})
		})

		Context("ReportError", func() {
			It("does not add error if nil", func() {
				base.ReportError(nil)
				Expect(base.Error()).ToNot(HaveOccurred())
			})

			It("reports the error", func() {
				err := testErrors.NewError()
				base.ReportError(err)
				Expect(base.Error()).To(Equal(errors.Append(err)))
			})

			It("reports the error with source", func() {
				err := testErrors.NewError()
				base.WithSource(src).ReportError(err)
				Expect(base.Error()).To(Equal(errors.WithSource(err, src)))
			})

			It("reports the error with meta", func() {
				err := testErrors.NewError()
				base.WithMeta(meta).ReportError(err)
				Expect(base.Error()).To(Equal(errors.WithMeta(err, meta)))
			})

			It("reports the error with source and meta", func() {
				err := testErrors.NewError()
				base.WithSource(src).WithMeta(meta).ReportError(err)
				Expect(base.Error()).To(Equal(errors.WithMeta(errors.WithSource(err, src), meta)))
			})

			It("reports the error on a offspring and the ancestor has it", func() {
				err := testErrors.NewError()
				result := base.WithMeta(meta).WithMeta(meta).WithMeta(meta)
				result.ReportError(err)
				Expect(result.Error()).To(Equal(errors.WithMeta(err, meta)))
				Expect(base.Error()).To(Equal(result.Error()))
			})
		})

		Context("WithSource", func() {
			It("returns a new base with source", func() {
				result := base.WithSource(src)
				Expect(result).ToNot(BeNil())
				Expect(result).ToNot(BeIdenticalTo(base))
				Expect(result.Error()).ToNot(HaveOccurred())
			})
		})

		Context("WithMeta", func() {
			It("returns a new base with meta", func() {
				result := base.WithMeta(meta)
				Expect(result).ToNot(BeNil())
				Expect(result).ToNot(BeIdenticalTo(base))
				Expect(result.Error()).ToNot(HaveOccurred())
			})
		})

		Context("WithReference", func() {
			It("returns a new base without change if no source", func() {
				reference := testStructure.NewReference()
				result := base.WithReference(reference)
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(base))
				Expect(result).ToNot(BeIdenticalTo(base))
				Expect(result.Error()).ToNot(HaveOccurred())
			})

			It("returns a new base with new source if source", func() {
				src.WithReferenceOutputs = []structure.Source{testStructure.NewSource()}
				reference := testStructure.NewReference()
				result := base.WithSource(src).WithReference(reference)
				Expect(result).ToNot(BeNil())
				Expect(result).ToNot(Equal(base))
				Expect(result.Error()).ToNot(HaveOccurred())
				Expect(src.WithReferenceInputs).To(Equal([]string{reference}))
			})
		})
	})
})