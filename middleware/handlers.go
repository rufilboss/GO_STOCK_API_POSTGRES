package middleware

import(

)

type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}