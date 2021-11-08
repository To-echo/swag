package api

import (
	"net/http"

	_ "github.com/To-echo/swag/testdata/conflict_name/model"
	_ "github.com/To-echo/swag/testdata/conflict_name/model3"
)

// @Tags Health
// @Description  Check if Health  of service it's OK!
// @ID health
// @Accept  json
// @Produce  json
// @Success 200 {object} model3.Response{data=model.ErrorsResponse}
// @Router /health [get]
func Get1(w http.ResponseWriter, r *http.Request) {

}
