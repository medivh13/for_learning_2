package books

import (
	validation "github.com/go-ozzo/ozzo-validation"
	booksProto "for_learning_2/src/app/proto/books"
)

type BookDTOInterface interface {
	Validate() error
}

type BookReqDTO struct {
	Subject string `json:"subject"`
}

func (dto *BookReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Subject, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type GetBooksRespDTO struct {
	Name        string     `json:"name"`
	SubjectType string     `json:"subject_type"`
	Works       []*WorkDTO `json:"works"`
}

type WorkDTO struct {
	Title        string       `json:"title"`
	CoverID      int64        `json:"cover_id"`
	EditionCount int64        `json:"edition_count"`
	Authors      []*AuthorDTO `json:"authors"`
}

type AuthorDTO struct {
	Name string `json:"name`
}

func TransformDTOToProto(dto *GetBooksRespDTO) (*booksProto.BookResp, error) {
	// Buat instansi BookResp protobuf
	protoResp := &booksProto.BookResp{
		Name:        dto.Name,
		SubjectType: dto.SubjectType,
	}

	// Iterasi melalui Works DTO dan transformasi ke Work protobuf
	for _, work := range dto.Works {
		protoWork := &booksProto.Work{
			Title:        work.Title,
			CoverId:      work.CoverID,
			EditionCount: work.EditionCount,
		}

		// Iterasi melalui Authors dalam Work dan transformasi ke Author protobuf
		for _, author := range work.Authors {
			protoAuthor := &booksProto.Author{
				Name: author.Name,
			}
			protoWork.Authors = append(protoWork.Authors, protoAuthor)
		}

		protoResp.Works = append(protoResp.Works, protoWork)
	}

	return protoResp, nil
}