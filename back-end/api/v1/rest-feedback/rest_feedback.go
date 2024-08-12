package restfeedback

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	errorcodes "wowcollector.io/internal/common/error-codes"
	httprequests "wowcollector.io/internal/entities/http-requests"
	errorresponse "wowcollector.io/internal/entities/response/error"
	githubhttp "wowcollector.io/internal/services/http/github-http"
)

func GetRoutes(r chi.Router) {
	r.Route("/feedback", func(r chi.Router) {
		r.Post("/", postFeedback)
	})
}

// @summary Submit a feedback form
// @description Submit a feedback form which will create a github issue
// @tags Feedback
// @accept multipart/form-data
// @produce json
// @Param description formData string true "Description" (multi-line text field)
// @Param attachments formData file false "Attachments (multiple files)"
// @Param email formData string false "Email (optional)"
// @Param battleTag formData string false "BattleTag (optional)"
// @Param rating formData int false "Rating"
// @Param type formData string true "bug or feedback"
// @success 200
// @failure 400 {object} errorresponse.ErrorResponse
// @router /api/v1/feedback [post]
func postFeedback(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Submitting feedback")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		zap.L().Error("Error parsing multipart form")
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	body := httprequests.GithubIssueBody{
		Owner:     "fyllekanin",
		Repo:      "wowcollector.io",
		Body:      r.FormValue("description"),
		IsBug:     r.FormValue("type") == "bug",
		Email:     r.FormValue("email"),
		BattleTag: r.FormValue("battleTag"),
		Rating:    r.FormValue("rating"),
	}

	if githubhttp.GetInstance().CreateIssue(&body) != nil {
		zap.L().Error("Error creating new issue")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, ""))
		return
	}

	w.WriteHeader(http.StatusOK)
}
