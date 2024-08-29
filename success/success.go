package success

import (
	"encoding/json"
	"net/http"
)

var ResponseKeys map[string]string = map[string]string{
	"success":  "success",
	"type":     "type",
	"status":   "status",
	"title":    "title",
	"detail":   "detail",
	"instance": "instance",
}

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

	if p.Status == 0 {
		p.Status = 200
	}

	problemMap[ResponseKeys["success"]] = true
	problemMap[ResponseKeys["status"]] = p.Status

	if p.Type != "" {
		problemMap[ResponseKeys["type"]] = p.Type
	}

	if p.Title != "" {
		problemMap[ResponseKeys["title"]] = p.Title
	}

	if p.Detail != "" {
		problemMap[ResponseKeys["detail"]] = p.Detail
	}

	if p.Instance != "" {
		problemMap[ResponseKeys["instance"]] = p.Instance
	}

	for _, extension := range p.Extensions {
		problemMap[extension.Key] = extension.Value
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

	w.WriteHeader(p.Status)
	w.Header().Set("content-type", "application/json")
	w.Write(jsonString)

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
