package usecase

import (
	bookUC "for_learning_2/src/app/usecase/books"
	pickUpUC "for_learning_2/src/app/usecase/pickup"
)

type AllUseCases struct {
	BookUC   bookUC.BooksUCInterface
	PickUpUC pickUpUC.PickUpUCInterface
}
