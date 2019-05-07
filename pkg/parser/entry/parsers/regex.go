package parsers

import (
	"regexp"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/prometheus/common/model"

	"github.com/grafana/loki/pkg/parser"
)

type Config struct {
	Expr   string
	Labels []parser.Label
}

type Regex struct {
	expr *regexp.Regexp
}

func NewRegex(config map[interface{}]interface{}) Regex {

	err := mapstructure.Decode(rg, &cfg2)
	return Regex{
		expr: regexp.MustCompile(config.Expr),
	}
}

func (r *Regex) Parse(labels model.LabelSet, time time.Time, entry string) (time.Time, string, error) {

}
