package usecase_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"peanut/config"
	"peanut/domain"
	"time"
)

var _ = Describe("Content", func() {

	var listContent []domain.Content
	var itemContent1 domain.Content
	var itemContent2 domain.Content

	var createContent domain.CreateContent

	BeforeEach(func() {
		itemContent1 = domain.Content{
			Name:        "name 1",
			Thumbnail:   "tmp/13ae27dd-9bdd-40dd-a22c-ab1985e45290_8R9OBJBQP9XQAEPBLZ2AJA.jpeg",
			Description: "",
			Resolution:  "",
			AspectRatio: "",
			Tag:         "tag",
			Category:    "category",
		}
		itemContent1.PlayTime, _ = time.Parse(config.TimeFormatDefault, "2006-01-02T22:04:05+07:00")
		itemContent1.ID = uint(1)

		itemContent2 = domain.Content{
			Name:        "name 2",
			Thumbnail:   "tmp/446e8237-aae1-4d4e-a826-d3213f746e7c_8R9OBJBQP9XQAEPBLZ2AJA.jpeg",
			Description: "",
			Resolution:  "",
			AspectRatio: "",
			Tag:         "tag",
			Category:    "category",
		}
		itemContent2.PlayTime, _ = time.Parse(config.TimeFormatDefault, "2023-01-02T22:04:05+07:00")
		itemContent2.ID = uint(2)

		createContent = domain.CreateContent{
			Name:        "new name",
			PlayTime:    "2023-01-02 15:04:05",
			Description: "",
			Resolution:  "",
			AspectRatio: "",
			Tag:         "tag",
			Category:    "category",
		}
	})

	Describe("API List content", func() {

		listContent = []domain.Content{
			itemContent1,
			itemContent2,
		}
		Context("- Get list content", func() {
			It("should be return list", func() {
				// prepare
				contentRepo.EXPECT().GetContents().Return(listContent, nil)
				// do
				contents, err := contentUc.GetContents(ctx)
				// check
				Expect(err).To(BeNil())
				Expect(contents).To(Equal(listContent))
			})
		})
	})

	Describe("API Create content", func() {

		Context("- create content", func() {
			It("should be success", func() {
				// prepare
				contentRepo.EXPECT().CreateContent(gomock.Any()).Return(nil)
				// do
				err := contentUc.CreateContent(ctx, createContent, "tmp/446e8237-aae1-4d4e-a826-d3213f746e7c_8R9OBJBQP9XQAEPBLZ2AJA.jpeg")
				// check
				Expect(err).To(BeNil())
			})
		})
	})
})
