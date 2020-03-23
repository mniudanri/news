package model

type NewsPostResponse struct {
    Id int64
    Author string
    Body string
    CreatedTime string
}

type NewsGetResponse struct {
    Id int64
    CreatedTime string
}