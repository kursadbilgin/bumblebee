package Controllers

import (
	"bumblebee/Models"
	"fmt"
	"net/http"
	"strings"
)

func validation(url, linkType string) *Models.ErrorResponse {
	if !strings.Contains(url, Models.HOMEPAGE_LINK) && linkType == "webLink" {
		return &Models.ErrorResponse{
			Status:       http.StatusBadRequest,
			ErrorMessage: fmt.Sprintf("Wrong web url: %s", url),
		}
	}

	if !strings.Contains(url, Models.DEEPLINK_PREFIX) && linkType == "deepLink" {
		return &Models.ErrorResponse{
			Status:       http.StatusBadRequest,
			ErrorMessage: fmt.Sprintf("Wrong deep url: %s", url),
		}
	}

	return nil
}
