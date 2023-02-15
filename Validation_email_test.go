package example_final

import (

	"testing"
"gorm.io/gorm"
."github.com/onsi/gomega"
"github.com/asaskevich/govalidator"

)


type Music struct{
	gorm.Model
	Email 	string   `valid:"email"`
}

func TestValidationMusic(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("check Email", func(t *testing.T) {

		m := Music{
			Email: "nun",
		}
		ok, err := govalidator.ValidateStruct(m)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email: nun does not validate as email"))
		 
	})

}