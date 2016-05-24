package validator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/tidepool-org/platform/data"
	"github.com/tidepool-org/platform/data/context"
	"github.com/tidepool-org/platform/data/validator"
)

var _ = Describe("StandardInterfaceArray", func() {
	It("New returns nil if context is nil", func() {
		value := []interface{}{}
		Expect(validator.NewStandardInterfaceArray(nil, "werewolf", &value)).To(BeNil())
	})

	Context("with context", func() {
		var standardContext *context.Standard

		BeforeEach(func() {
			standardContext = context.NewStandard()
		})

		Context("new validator with nil reference and nil value", func() {
			var standardInterfaceArray *validator.StandardInterfaceArray
			var result data.InterfaceArray

			BeforeEach(func() {
				standardInterfaceArray = validator.NewStandardInterfaceArray(standardContext, nil, nil)
			})

			It("exists", func() {
				Expect(standardInterfaceArray).ToNot(BeNil())
			})

			Context("Exists", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.Exists()
				})

				It("adds the expected error", func() {
					Expect(standardContext.Errors()).To(HaveLen(1))
					Expect(standardContext.Errors()[0]).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Code).To(Equal("value-does-not-exist"))
					Expect(standardContext.Errors()[0].Title).To(Equal("value does not exist"))
					Expect(standardContext.Errors()[0].Detail).To(Equal("Value does not exist"))
					Expect(standardContext.Errors()[0].Source).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Source.Pointer).To(Equal("/<nil>"))
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthEqualTo(1)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthNotEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthNotEqualTo(4)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthLessThan", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthLessThan(3)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthLessThanOrEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthLessThan(1)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthGreaterThan", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthGreaterThan(3)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthGreaterThanOrEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthGreaterThanOrEqualTo(4)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthInRange", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthInRange(0, 1)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})
		})

		Context("new validator with valid reference and value with length of 1", func() {
			var standardInterfaceArray *validator.StandardInterfaceArray
			var result data.InterfaceArray

			BeforeEach(func() {
				value := []interface{}{"one"}
				standardInterfaceArray = validator.NewStandardInterfaceArray(standardContext, "werewolf", &value)
			})

			It("exists", func() {
				Expect(standardInterfaceArray).ToNot(BeNil())
			})

			Context("Exists", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.Exists()
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthEqualTo(1)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthNotEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthNotEqualTo(4)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthLessThan", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthLessThan(3)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthLessThanOrEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthLessThanOrEqualTo(1)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthGreaterThan", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthGreaterThan(3)
				})

				It("adds the expected error", func() {
					Expect(standardContext.Errors()).To(HaveLen(1))
					Expect(standardContext.Errors()[0]).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Code).To(Equal("length-out-of-range"))
					Expect(standardContext.Errors()[0].Title).To(Equal("length is out of range"))
					Expect(standardContext.Errors()[0].Detail).To(Equal("Length 1 is not greater than 3"))
					Expect(standardContext.Errors()[0].Source).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Source.Pointer).To(Equal("/werewolf"))
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthGreaterThanOrEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthGreaterThanOrEqualTo(4)
				})

				It("adds the expected error", func() {
					Expect(standardContext.Errors()).To(HaveLen(1))
					Expect(standardContext.Errors()[0]).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Code).To(Equal("length-out-of-range"))
					Expect(standardContext.Errors()[0].Title).To(Equal("length is out of range"))
					Expect(standardContext.Errors()[0].Detail).To(Equal("Length 1 is not greater than or equal to 4"))
					Expect(standardContext.Errors()[0].Source).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Source.Pointer).To(Equal("/werewolf"))
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthInRange", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthInRange(0, 3)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})
		})

		Context("new validator with valid reference and value with length of 4", func() {
			var standardInterfaceArray *validator.StandardInterfaceArray
			var result data.InterfaceArray

			BeforeEach(func() {
				value := []interface{}{"one", "two", "three", "four"}
				standardInterfaceArray = validator.NewStandardInterfaceArray(standardContext, "werewolf", &value)
			})

			It("exists", func() {
				Expect(standardInterfaceArray).ToNot(BeNil())
			})

			Context("Exists", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.Exists()
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthEqualTo(1)
				})

				It("adds the expected error", func() {
					Expect(standardContext.Errors()).To(HaveLen(1))
					Expect(standardContext.Errors()[0]).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Code).To(Equal("length-out-of-range"))
					Expect(standardContext.Errors()[0].Title).To(Equal("length is out of range"))
					Expect(standardContext.Errors()[0].Detail).To(Equal("Length 4 is not equal to 1"))
					Expect(standardContext.Errors()[0].Source).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Source.Pointer).To(Equal("/werewolf"))
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthNotEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthNotEqualTo(4)
				})

				It("adds the expected error", func() {
					Expect(standardContext.Errors()).To(HaveLen(1))
					Expect(standardContext.Errors()[0]).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Code).To(Equal("length-out-of-range"))
					Expect(standardContext.Errors()[0].Title).To(Equal("length is out of range"))
					Expect(standardContext.Errors()[0].Detail).To(Equal("Length 4 is equal to 4"))
					Expect(standardContext.Errors()[0].Source).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Source.Pointer).To(Equal("/werewolf"))
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthLessThan", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthLessThan(3)
				})

				It("adds the expected error", func() {
					Expect(standardContext.Errors()).To(HaveLen(1))
					Expect(standardContext.Errors()[0]).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Code).To(Equal("length-out-of-range"))
					Expect(standardContext.Errors()[0].Title).To(Equal("length is out of range"))
					Expect(standardContext.Errors()[0].Detail).To(Equal("Length 4 is not less than 3"))
					Expect(standardContext.Errors()[0].Source).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Source.Pointer).To(Equal("/werewolf"))
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthLessThanOrEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthLessThanOrEqualTo(1)
				})

				It("adds the expected error", func() {
					Expect(standardContext.Errors()).To(HaveLen(1))
					Expect(standardContext.Errors()[0]).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Code).To(Equal("length-out-of-range"))
					Expect(standardContext.Errors()[0].Title).To(Equal("length is out of range"))
					Expect(standardContext.Errors()[0].Detail).To(Equal("Length 4 is not less than or equal to 1"))
					Expect(standardContext.Errors()[0].Source).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Source.Pointer).To(Equal("/werewolf"))
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthGreaterThan", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthGreaterThan(3)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthGreaterThanOrEqualTo", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthGreaterThanOrEqualTo(4)
				})

				It("does not add an error", func() {
					Expect(standardContext.Errors()).To(BeEmpty())
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})

			Context("LengthInRange", func() {
				BeforeEach(func() {
					result = standardInterfaceArray.LengthInRange(0, 3)
				})

				It("adds the expected error", func() {
					Expect(standardContext.Errors()).To(HaveLen(1))
					Expect(standardContext.Errors()[0]).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Code).To(Equal("length-out-of-range"))
					Expect(standardContext.Errors()[0].Title).To(Equal("length is out of range"))
					Expect(standardContext.Errors()[0].Detail).To(Equal("Length 4 is not between 0 and 3"))
					Expect(standardContext.Errors()[0].Source).ToNot(BeNil())
					Expect(standardContext.Errors()[0].Source.Pointer).To(Equal("/werewolf"))
				})

				It("returns self", func() {
					Expect(result).To(BeIdenticalTo(standardInterfaceArray))
				})
			})
		})
	})
})