package usecase_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"peanut/domain"
)

var _ = Describe("Content", func() {

	var listContent []domain.Content
	var itemContent1 domain.Content
	var itemContent2 domain.Content

	BeforeEach(func() {
		itemContent1 = domain.Content{
			Name:      "name 1",
			Thumbnail: "Content 1",
			Category:  "1",
			Tag:       "1",
		}
		itemContent1.ID = uint(1)

		itemContent1 = domain.Content{
			Name:      "name 1",
			Thumbnail: "Content 1",
			Category:  "1",
			Tag:       "1",
		}
		itemContent1.ID = uint(1)
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
				content, err := contentUc.GetContents()
				// check
				Expect(err).To(BeNil())
				Expect(content).To(Equal(listContent))
			})
		})
	})

	Describe("API Create content", func() {

		Context("- create content", func() {
			It("should be return content", func() {
				// prepare
				contentRepo.EXPECT().CreateContent(itemContent1).Return(itemContent1, nil)
				// do
				content, err := contentUc.CreateContent(itemContent1)
				// check
				Expect(err).To(BeNil())
				Expect(content).To(Equal(itemContent1))
			})
		})
	})
})
