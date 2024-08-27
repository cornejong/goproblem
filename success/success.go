package success

import (
	"encoding/json"
	"net/http"

	"github.com/cornejong/goproblem/caseconverter"
)

type Success struct {
	Type     string `json:"type,omitempty"`
	Status   int    `json:"status"`
	Title    string `json:"title,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance"`

	Extensions []SuccessExtension `json:"-"`
}

type SuccessExtension struct {
	Key   string
	Value any
}

func (p *Success) WithStatus(status int) *Success {
	p.Status = status
	return p
}

func (p *Success) WithType(problemType string) *Success {
	p.Type = problemType
	return p
}

func (p *Success) WithTitle(title string) *Success {
	p.Title = title
	return p
}

func (p *Success) WithDetail(detail string) *Success {
	p.Detail = detail
	return p
}

func (p *Success) WithInstance(instance string) *Success {
	p.Instance = instance
	return p
}

func (p *Success) WithExtension(key string, value any) *Success {
	p.Extensions = append(p.Extensions, SuccessExtension{
		Key:   key,
		Value: value,
	})

	return p
}

func (p *Success) BuildMap() map[string]any {
	problemMap := make(map[string]any)

	if p.Type != "" {
		problemMap[caseconverter.ResponseKeyCasingConverter("type")] = p.Type
	} else {
		problemMap[caseconverter.ResponseKeyCasingConverter("type")] = "about:blank"
	}

	problemMap[caseconverter.ResponseKeyCasingConverter("status")] = p.Status
	problemMap[caseconverter.ResponseKeyCasingConverter("title")] = p.Title

	if p.Detail != "" {
		problemMap[caseconverter.ResponseKeyCasingConverter("detail")] = p.Detail
	}

	problemMap[caseconverter.ResponseKeyCasingConverter("instance")] = p.Instance

	for _, extension := range p.Extensions {
		problemMap[caseconverter.ResponseKeyCasingConverter(extension.Key)] = extension.Value
	}

	return problemMap
}

func (p *Success) AsJson() ([]byte, error) {
	jsonString, err := json.Marshal(p.BuildMap())
	if err != nil {
		return nil, err
	}

	return jsonString, nil
}

func (p *Success) WriteToResponseWriter(w http.ResponseWriter) error {
	jsonString, err := p.AsJson()
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/problem")
	w.Write(jsonString)
	w.WriteHeader(p.Status)

	return nil
}

func (p *Success) RespondWithSuccess(w http.ResponseWriter, r *http.Request) error {
	if p.Instance == "" {
		p.WithInstance(r.URL.Path)
	}

	return p.WriteToResponseWriter(w)
}

func RespondWithSuccess(w http.ResponseWriter, r *http.Request, problem *Success) error {
	if problem.Instance == "" {
		problem.WithInstance(r.URL.Path)
	}

	return problem.WriteToResponseWriter(w)
}
