package problem

import (
	"encoding/json"
	"net/http"
)

var (
	ResponseKeyCasingConverter StringCasingConverter = SnakeCaseConverter
	DefaultErrorKey            string                = "reason"
)

type Problem struct {
	Type     string `json:"type"`
	Status   int    `json:"status"`
	Title    string `json:"title"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance"`

	Extensions []ProblemExtension `json:"-"`
}

type ProblemExtension struct {
	Key   string
	Value any
}

func (p *Problem) WithStatus(status int) *Problem {
	p.Status = status
	return p
}

func (p *Problem) WithType(problemType string) *Problem {
	p.Type = problemType
	return p
}

func (p *Problem) WithTitle(title string) *Problem {
	p.Title = title
	return p
}

func (p *Problem) WithDetail(detail string) *Problem {
	p.Detail = detail
	return p
}

func (p *Problem) WithInstance(instance string) *Problem {
	p.Instance = instance
	return p
}

func (p *Problem) WithExtension(key string, value any) *Problem {
	p.Extensions = append(p.Extensions, ProblemExtension{
		Key:   key,
		Value: value,
	})

	return p
}

func (p *Problem) BuildMap() map[string]any {
	problemMap := make(map[string]any)

	if p.Type != "" {
		problemMap[ResponseKeyCasingConverter("type")] = p.Type
	} else {
		problemMap[ResponseKeyCasingConverter("type")] = "about:blank"
	}

	problemMap[ResponseKeyCasingConverter("status")] = p.Status
	problemMap[ResponseKeyCasingConverter("title")] = p.Title

	if p.Detail != "" {
		problemMap[ResponseKeyCasingConverter("detail")] = p.Detail
	}

	problemMap[ResponseKeyCasingConverter("instance")] = p.Instance

	for _, extension := range p.Extensions {
		problemMap[ResponseKeyCasingConverter(extension.Key)] = extension.Value
	}

	return problemMap
}

func (p *Problem) AsJson() ([]byte, error) {
	jsonString, err := json.Marshal(p.BuildMap())
	if err != nil {
		return nil, err
	}

	return jsonString, nil
}

func (p *Problem) WriteToResponseWriter(w http.ResponseWriter) error {
	jsonString, err := p.AsJson()
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/problem+json")
	w.Write(jsonString)
	w.WriteHeader(p.Status)

	return nil
}

func RespondWithProblem(w http.ResponseWriter, r *http.Request, problem Problem) error {
	if problem.Instance == "" {
		problem.WithInstance(r.URL.Path)
	} 

	return problem.WriteToResponseWriter(w)
}

func From(problem *Problem) *Problem {
	p := *problem
	newP := p
	return &newP
}
