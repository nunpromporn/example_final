package example_final

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required"`
	Url  string `valid:"url"`
}

func TestValidationVideO(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("check Name", func(t *testing.T) {

		v := Video{
			Name: "",
			Url:  "https://www.youtube.com/watch?v=YjSPnhLy5N0",
		}
		ok, err := govalidator.ValidateStruct(v)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name: non zero value required"))

	})

}
