package optional_test

import (
	"github.com/CameronHonis/optional"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("EmptyOptional", func() {
	It("returns an empty optional", func() {
		opt := optional.EmptyOptional[int]()
		val, err := opt.Get()
		Expect(err).To(HaveOccurred())
		Expect(val).To(BeZero())
	})
})
var _ = Describe("NewOptional", func() {
	It("returns an optional with the value", func() {
		opt := optional.NewOptional[int](12)
		val, err := opt.Get()
		Expect(err).ToNot(HaveOccurred())
		Expect(val).To(Equal(12))
	})
})
var _ = Describe("Optional", func() {
	var opt *optional.Optional[int]
	BeforeEach(func() {
		opt = optional.EmptyOptional[int]()
	})
	Describe("::IsPresent", func() {
		When("the optional is empty", func() {
			It("returns false", func() {
				Expect(opt.IsPresent()).To(BeFalse())
			})
		})
		When("the optional is not empty", func() {
			BeforeEach(func() {
				opt = optional.NewOptional[int](12)
			})
			It("returns true", func() {
				Expect(opt.IsPresent()).To(BeTrue())
			})
		})
	})
	Describe("::GetOrElse", func() {
		When("the optional is empty", func() {
			It("returns the default value", func() {
				val := opt.GetOrElse(7)
				Expect(val).To(Equal(7))
			})
		})
		When("the optional is not empty", func() {
			BeforeEach(func() {
				opt = optional.NewOptional[int](12)
			})
			It("returns the value", func() {
				val := opt.GetOrElse(7)
				Expect(val).To(Equal(12))
			})
		})
	})
	Describe("::Get", func() {
		When("the optional is empty", func() {
			It("returns an error", func() {
				val, err := opt.Get()
				Expect(err).To(HaveOccurred())
				Expect(val).To(BeZero())
			})
			When("the type is a pointer", func() {
				var optPointer *optional.Optional[*int]
				BeforeEach(func() {
					optPointer = optional.EmptyOptional[*int]()
				})
				It("returns nil (?) and an error", func() {
					val, err := optPointer.Get()
					Expect(err).To(HaveOccurred())
					Expect(val).To(BeNil())
				})
			})
		})
		When("the optional is not empty", func() {
			BeforeEach(func() {
				opt = optional.NewOptional[int](12)
			})
			It("returns the value", func() {
				val, err := opt.Get()
				Expect(err).ToNot(HaveOccurred())
				Expect(val).To(Equal(12))
			})
		})
	})
})
