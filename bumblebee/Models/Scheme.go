package Models

const HOMEPAGE_LINK = "https://www.trendyol.com"
const HOMEPAGE_DEEPLINK = "ty://?Page=Home"
const DEEPLINK_PREFIX = "ty://?"
const SEARCH_PREFIX = "/sr"
const PRODUCT_PREFIX = "-p-"
const Q_PARAM = "q"
const PAGE_SEARCH_PARAM = "Page=Search"
const PAGE_PRODUCT_PARAM = "Page=Product"
const QUERY_PARAM = "Query"
const MERCHANT_ID_LOWER = "merchantId"
const CONTENT_ID = "ContentId"
const BOUTIQUE_ID = "boutiqueId"
const CAMPAIGN_ID = "CampaignId"
const MERCHANT_ID = "MerchantId"

var IDList = []string{"&ContentId=", "CampaignId", "MerchantId"}

type ConverterWebToDeepLinkRequest struct {
	WebURL string `json:"web_url"`
}

type ConverterDeepLinkToWebRequest struct {
	DeepLink string `json:"deep_link"`
}

type ResponseWebToDeep struct {
	Status   int    `json:"status"`
	DeepLink string `json:"deep_link"`
}

type ResponseDeepToWeb struct {
	Status  int    `json:"status"`
	WebLink string `json:"web_link"`
}

type ErrorResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message"`
}
