package modelyesod

import "github.com/invopop/jsonschema"

type SimpleKeywordFilterActionConfig struct {
	TitleInclude   []string `json:"title_include,omitempty"   jsonschema:"title=Title include"`
	TitleExclude   []string `json:"title_exclude,omitempty"   jsonschema:"title=Title exclude"`
	ContentInclude []string `json:"content_include,omitempty" jsonschema:"title=Content include"`
	ContentExclude []string `json:"content_exclude,omitempty" jsonschema:"title=Content exclude"`
}

func GetSimpleKeywordFilterActionConfigSchema() (string, error) {
	return reflectJSONSchema(new(SimpleKeywordFilterActionConfig))
}

type KeywordFilterActionConfig struct {
	OrList []*KeywordFilterAndList `json:"or_list" jsonschema:"title=OR list"`
}

func GetKeywordFilterActionConfigSchema() (string, error) {
	return reflectJSONSchema(new(KeywordFilterActionConfig))
}

type KeywordFilterAndList struct {
	AndList []*KeywordFilter `json:"and_list" jsonschema:"title=AND list"`
}

type KeywordFilter struct {
	Field    string `json:"field"    jsonschema:"title=field,enum=author,enum=title,enum=description,enum=content"`
	Equation string `json:"equation" jsonschema:"title=equation,enum=equal,enum=not_equal,enum=contain,enum=not_contain,enum=start_with,enum=not_start_with,enum=end_with,enum=not_end_with"`
	Value    string `json:"value"    jsonschema:"title=value"`
}

type KeywordFilterField string

const (
	KeywordFilterFieldAuthor      KeywordFilterField = "author"
	KeywordFilterFieldTitle       KeywordFilterField = "title"
	KeywordFilterFieldDescription KeywordFilterField = "description"
	KeywordFilterFieldContent     KeywordFilterField = "content"
)

type KeywordFilterEquation string

const (
	KeywordFilterEquationEqual        KeywordFilterEquation = "equal"
	KeywordFilterEquationNotEqual     KeywordFilterEquation = "not_equal"
	KeywordFilterEquationContain      KeywordFilterEquation = "contain"
	KeywordFilterEquationNotContain   KeywordFilterEquation = "not_contain"
	KeywordFilterEquationStartWith    KeywordFilterEquation = "start_with"
	KeywordFilterEquationNotStartWith KeywordFilterEquation = "not_start_with"
	KeywordFilterEquationEndWith      KeywordFilterEquation = "end_with"
	KeywordFilterEquationNotEndWith   KeywordFilterEquation = "not_end_with"
)

type DescriptionGeneratorActionConfig struct{}

func GetDescriptionGeneratorActionConfigSchema() (string, error) {
	return reflectJSONSchema(new(DescriptionGeneratorActionConfig))
}

func reflectJSONSchema(v interface{}) (string, error) {
	r := new(jsonschema.Reflector)
	r.ExpandedStruct = true
	r.DoNotReference = true
	j, err := r.Reflect(v).MarshalJSON()
	if err != nil {
		return "", err
	}
	return string(j), nil
}
