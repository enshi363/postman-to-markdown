package main

type ItemStructV2Point1 struct {
	Name    string               `json:"name"`
	Item    []ItemStructV2Point1 `json:"item"`
	Request struct {
		Method string `json:"method"`
		Header []struct {
			Key         string `json:"key"`
			Value       string `json:"value"`
			Description string `json:"description"`
		} `json:"header"`
		Body struct {
			Mode     string `json:"mode"`
			Raw      string `json:"raw"`
			Formdata []struct {
				Key         string `json:"key"`
				Value       string `json:"value"`
				Description string `json:"description"`
				Type        string `json:"type"`
			} `json:"formdata"`
			Urlencoded []struct {
				Key         string `json:"key"`
				Value       string `json:"value"`
				Description string `json:"description"`
				Type        string `json:"type"`
			} `json:"urlencoded"`
			Options struct {
				Raw struct {
					Language string `json:"language"`
				} `json:"raw"`
			} `json:"options"`
		} `json:"body"`
		URL struct {
			Raw   string   `json:"raw"`
			Host  []string `json:"host"`
			Port  string   `json:"port"`
			Path  []string `json:"path"`
			Query []struct {
				Key         string `json:"key"`
				Value       string `json:"value"`
				Description string `json:"description"`
			} `json:"query"`
			Variable []struct {
				Key         string `json:"key"`
				Value       string `json:"value"`
				Description string `json:"description"`
			} `json:"variable"`
		} `json:"url"`
		Description string `json:"description"`
	} `json:"request"`
	Response []struct {
		Name            string `json:"name"`
		OriginalRequest struct {
			Method string `json:"method"`
			Header []struct {
				Key         string `json:"key"`
				Value       string `json:"value"`
				Description string `json:"description"`
			} `json:"header"`
			Body struct {
				Mode     string `json:"mode"`
				Raw      string `json:"raw"`
				Formdata []struct {
					Key         string `json:"key"`
					Value       string `json:"value"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"formdata"`
				Urlencoded []struct {
					Key         string `json:"key"`
					Value       string `json:"value"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"urlencoded"`
				Options struct {
					Raw struct {
						Language string `json:"language"`
					} `json:"raw"`
				} `json:"options"`
			} `json:"body"`
			URL struct {
				Raw   string   `json:"raw"`
				Host  []string `json:"host"`
				Port  string   `json:"port"`
				Path  []string `json:"path"`
				Query []struct {
					Key         string `json:"key"`
					Value       string `json:"value"`
					Description string `json:"description"`
				} `json:"query"`
				Variable []struct {
					Key         string `json:"key"`
					Value       string `json:"value"`
					Description string `json:"description"`
				} `json:"variable"`
			} `json:"url"`
		} `json:"originalRequest"`
		Status                 string `json:"status"`
		Code                   int    `json:"code"`
		PostmanPreviewlanguage string `json:"_postman_previewlanguage"`
		Header                 []struct {
			Key         string `json:"key"`
			Value       string `json:"value"`
			Description string `json:"description"`
		} `json:"header"`
		Cookie []interface{} `json:"cookie"`
		Body   string        `json:"body"`
	} `json:"response"`
}

type PostManStructV2Point1 struct {
	Info struct {
		PostmanID   string `json:"_postman_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Schema      string `json:"schema"`
	} `json:"info"`
	Item  []ItemStructV2Point1 `json:"item"`
	Event []struct {
		Listen string `json:"listen"`
		Script struct {
			Type string   `json:"type"`
			Exec []string `json:"exec"`
		} `json:"script"`
	} `json:"event"`
}
