package admin

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/owncast/owncast/controllers"
	"github.com/owncast/owncast/core/data"
	"github.com/owncast/owncast/models"
	log "github.com/sirupsen/logrus"
)

type ConfigValue struct {
	Value interface{} `json:"value"`
}

func ChangeTags(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValues, success := getValuesFromRequest(w, r)
	if !success {
		return
	}

	var tagStrings []string
	for _, tag := range configValues {
		tagStrings = append(tagStrings, tag.Value.(string))
	}

	if err := data.SetServerMetadataTags(tagStrings); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "changed")
}

func ChangeStreamTitle(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetStreamTitle(configValue.Value.(string)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "changed")
}

func ChangeServerName(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetServerName(configValue.Value.(string)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "changed")
}

func ChangeServerTitle(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetServerTitle(configValue.Value.(string)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "changed")
}

func ChangeServerSummary(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetServerSummary(configValue.Value.(string)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "changed")
}

func ChangeExtraPageContent(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetExtraPageBodyContent(configValue.Value.(string)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "changed")
}

func ChangeStreamKey(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetStreamKey(configValue.Value.(string)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "changed")
}

func ChangeLogoPath(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetLogoPath(configValue.Value.(string)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "changed")
}

func ChangeNSFW(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetNSFW(configValue.Value.(bool)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "changed")
}

func ChangeFfmpegPath(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	path := configValue.Value.(string)
	if err := data.VerifyFFMpegPath(path); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	if err := data.SetFfmpegPath(configValue.Value.(string)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "changed")
}

func ChangeWebServerPort(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetHTTPPortNumber(configValue.Value.(int)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "http port set")
}

func ChangeRTMPServerPort(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetRTMPPortNumber(configValue.Value.(int)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "rtmp port set")
}

func ChangeServerURL(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetServerURL(configValue.Value.(string)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "server url set")
}

func ChangeDirectoryEnabled(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetDirectoryEnabled(configValue.Value.(bool)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}
	controllers.WriteSimpleResponse(w, true, "directory state changed")
}

func ChangeDisableUpgradeChecks(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetDisableUpgradeChecks(configValue.Value.(bool)); err != nil {
		controllers.WriteSimpleResponse(w, false, err.Error())
		return
	}
	controllers.WriteSimpleResponse(w, true, "disable upgrade checks changed")
}

func SetStreamLatencyLevel(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	configValue, success := getValueFromRequest(w, r)
	if !success {
		return
	}

	if err := data.SetStreamLatencyLevel(configValue.Value.(float64)); err != nil {
		controllers.WriteSimpleResponse(w, false, "error setting stream latency "+err.Error())
		return
	}

	controllers.WriteSimpleResponse(w, true, "set stream latency")
}

func SetS3Configuration(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	type s3ConfigurationRequest struct {
		Value models.S3 `json:"value"`
	}

	decoder := json.NewDecoder(r.Body)
	var newS3Config s3ConfigurationRequest
	if err := decoder.Decode(&newS3Config); err != nil {
		controllers.WriteSimpleResponse(w, false, "unable to update s3 config with provided values")
		return
	}

	if newS3Config.Value.Enabled {
		if newS3Config.Value.AccessKey == "" || newS3Config.Value.Secret == "" {
			controllers.WriteSimpleResponse(w, false, "s3 support requires an access key and secret")
			return
		}

		if newS3Config.Value.Region == "" || newS3Config.Value.Endpoint == "" {
			controllers.WriteSimpleResponse(w, false, "s3 support requires a region and endpoint")
			return
		}

		if newS3Config.Value.Bucket == "" {
			controllers.WriteSimpleResponse(w, false, "s3 support requires a bucket created for storing public video segments")
			return
		}
	}

	data.SetS3Config(newS3Config.Value)
	controllers.WriteSimpleResponse(w, true, "storage configuration changed")

}

func SetStreamOutputVariants(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	type streamOutputVariantRequest struct {
		Value []models.StreamOutputVariant `json:"value"`
	}

	decoder := json.NewDecoder(r.Body)
	var videoVariants streamOutputVariantRequest
	if err := decoder.Decode(&videoVariants); err != nil {
		controllers.WriteSimpleResponse(w, false, "unable to update video config with provided values")
		return
	}

	if err := data.SetStreamOutputVariants(videoVariants.Value); err != nil {
		controllers.WriteSimpleResponse(w, false, "unable to update video config with provided values")
		return
	}

	controllers.WriteSimpleResponse(w, true, "stream output variants updated")
}

func SetSocialHandles(w http.ResponseWriter, r *http.Request) {
	if !requirePOST(w, r) {
		return
	}

	type socialHandlesRequest struct {
		Value []models.SocialHandle `json:"value"`
	}

	decoder := json.NewDecoder(r.Body)
	var socialHandles socialHandlesRequest
	if err := decoder.Decode(&socialHandles); err != nil {
		controllers.WriteSimpleResponse(w, false, "unable to update social handles with provided values")
		return
	}

	if err := data.SetSocialHandles(socialHandles.Value); err != nil {
		controllers.WriteSimpleResponse(w, false, "unable to update social handles with provided values")
		return
	}

	controllers.WriteSimpleResponse(w, true, "social handles updated")
}

func requirePOST(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		controllers.WriteSimpleResponse(w, false, r.Method+" not supported")
		return false
	}

	return true
}

func getValueFromRequest(w http.ResponseWriter, r *http.Request) (ConfigValue, bool) {
	decoder := json.NewDecoder(r.Body)
	var configValue ConfigValue
	if err := decoder.Decode(&configValue); err != nil {
		log.Warnln(err)
		controllers.WriteSimpleResponse(w, false, "unable to parse new value")
		return configValue, false
	}

	return configValue, true
}

func getValuesFromRequest(w http.ResponseWriter, r *http.Request) ([]ConfigValue, bool) {
	var values []ConfigValue

	decoder := json.NewDecoder(r.Body)
	var configValue ConfigValue
	if err := decoder.Decode(&configValue); err != nil {
		controllers.WriteSimpleResponse(w, false, "unable to parse array of values")
		return values, false
	}

	object := reflect.ValueOf(configValue.Value)

	for i := 0; i < object.Len(); i++ {
		values = append(values, ConfigValue{Value: object.Index(i).Interface()})
	}

	return values, true
}