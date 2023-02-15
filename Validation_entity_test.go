package example_final

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Customer struct{
	gorm.Model
	User 	string		`valid:"requried"`
}

func TestValidation(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("check status not blank" , func(t *testing.T) {

		c := Customer{
			User: "nun",
		}
			ok,err := govalidator.ValidateStruct(c)

			g.Expect(ok).NotTo(BeTrue())
			g.Expect(err).NotTo(BeNil())
			g.Expect(err.Error()).To(Equal("User: The following validator is invalid or can't be applied to the field: \"requried\""))
		

	})
}

	
