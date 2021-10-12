package Controllers

import (
	"bumblebee/Models"
	"fmt"
	"net/url"
	"strings"
)

func DeepToWeb(deepUrl string) (string, error) {
	parseUrl, err := url.Parse(deepUrl)
	if err != nil {
		logError := fmt.Sprintf("[DeepToWeb] Url Parse Error")
		LogError(logError, err.Error())

		return "", err
	}

	hostReplace := strings.Replace(deepUrl, Models.DEEPLINK_PREFIX, Models.HOMEPAGE_LINK, -1)
	rawQuerySplit := strings.Split(parseUrl.RawQuery, "&")

	if strings.Contains(parseUrl.String(), Models.PAGE_SEARCH_PARAM) {
		searchReplace := strings.Replace(hostReplace, rawQuerySplit[0], Models.SEARCH_PREFIX+"?", -1)
		searchQuery := strings.Replace(searchReplace, "&"+Models.QUERY_PARAM, Models.Q_PARAM, -1)

		logMessage := fmt.Sprintf("[ConvertDeepToWebLinkSearch] Deep Link: %s, Web Link: %s", deepUrl, searchQuery)
		LogInfo(logMessage)

		return searchQuery, nil
	}

	if strings.Contains(parseUrl.String(), Models.PAGE_PRODUCT_PARAM) {
		productReplace := strings.Replace(hostReplace, rawQuerySplit[0], "/brand/", -1)
		contentIDReplace := strings.Replace(productReplace, Models.IDList[0], "name-p-", -1)
		boutiqueIDReplace := strings.Replace(contentIDReplace, Models.IDList[1], Models.BOUTIQUE_ID, -1)
		merchantIDReplace := strings.Replace(boutiqueIDReplace, Models.IDList[2], Models.MERCHANT_ID_LOWER, -1)
		productQuery := strings.Replace(merchantIDReplace, "&", "?", 1)

		logMessage := fmt.Sprintf("[ConvertDeepToWebLinkProduct] Deep Link: %s, Web Link: %s", deepUrl, productQuery)
		LogInfo(logMessage)

		return productQuery, nil
	}

	return Models.HOMEPAGE_LINK, nil
}
