// Code generated by goctl. DO NOT EDIT.
package types


type AddBookReq struct {
	Name  string `form:"name"`
	Price int64  `form:"price"`
}

type AddBookResp struct {
	Ok bool `json:"ok"`
}

type QueryByNameReq struct {
	Name string `form:"name"`
}

type QueryByNameResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type QueryAllReq struct {
}

type QueryAllResp struct {
	BookList []*QueryByNameResponse `json:"bookList"`
}
