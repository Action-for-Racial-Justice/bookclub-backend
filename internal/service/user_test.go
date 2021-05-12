package service_test

// var (
// 	user = &models.Book{
// 		EntryID:  "1",
// 		Name:     "A Farewell to Arms",
// 		Author:   "Ernest Hemingway",
// 		IsActive: true,
// 	}
// )

// func TestGetSSOToken(t *testing.T) {
// 	ts := createTestSuite(t)

// 	ts.mockMysql.EXPECT().GetLoginResponse("1").Return(book, nil).Times(1)
// 	bookData := ts.svc.GetBookData("1")

// 	assert.Equal(t, book, bookData)
// }

// func TestGetBookDataForEntryIDFail(t *testing.T) {
// 	ts := createTestSuite(t)

// 	ts.mockMysql.EXPECT().GetBookDataForEntryID("1").Return(nil, errors.New("GetBookDataForEntryID error")).Times(1)
// 	bookData := ts.svc.GetBookData("1")

// 	assert.NotEqual(t, book, bookData)
// }
