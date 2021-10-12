package Controllers

import (
	"bumblebee/Models"
	"fmt"
	"net/url"
	"strings"
)

func WebToDeep(webUrl string) (string, error) {
	parseUrl, err := url.Parse(webUrl)
	if err != nil {
		logError := fmt.Sprintf("[WebToDeep] Url Parse Error")
		LogError(logError, err.Error())

		return "", err
	}
	replaceUrl := strings.Replace(webUrl, "?", "&", -1)
	hostReplace := strings.Replace(replaceUrl, Models.HOMEPAGE_LINK, Models.DEEPLINK_PREFIX, -1)

	if strings.Contains(parseUrl.String(), Models.SEARCH_PREFIX) {
		searchReplace := strings.Replace(hostReplace, Models.SEARCH_PREFIX, Models.PAGE_SEARCH_PARAM, -1)
		searchQuery := strings.Replace(searchReplace, Models.Q_PARAM, Models.QUERY_PARAM, -1)

		logMessage := fmt.Sprintf("[ConvertWebToDeepLinkSearch] Web Link: %s, Deep Link: %s", webUrl, searchQuery)
		LogInfo(logMessage)

		return searchQuery, nil
	}

	if strings.Contains(parseUrl.String(), Models.PRODUCT_PREFIX) {
		pathSplit := strings.Split(parseUrl.Path, Models.PRODUCT_PREFIX)
		productReplace := strings.Replace(hostReplace, pathSplit[0], Models.PAGE_PRODUCT_PARAM, -1)
		boutiqueIDReplace := strings.Replace(productReplace, Models.BOUTIQUE_ID, Models.CAMPAIGN_ID, -1)
		merchantIDReplace := strings.Replace(boutiqueIDReplace, Models.MERCHANT_ID_LOWER, Models.MERCHANT_ID, -1)
		queryReplace := strings.Replace(merchantIDReplace, Models.PRODUCT_PREFIX, "&"+Models.CONTENT_ID+"=", -1)

		logMessage := fmt.Sprintf("[ConvertWebToDeepLinkProduct] Web Link: %s, Deep Link: %s", webUrl, queryReplace)
		LogInfo(logMessage)

		return queryReplace, nil
	}

	return Models.HOMEPAGE_DEEPLINK, nil
}
