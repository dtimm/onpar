package matchers

import (
	"fmt"
	"regexp"
)

type RegexpMatcher struct {
	pattern string
}

func Regexp(pattern string) RegexpMatcher {
	return RegexpMatcher{
		pattern: pattern,
	}
}

func (m RegexpMatcher) Match(actual interface{}) (interface{}, error) {
	r, err := regexp.Compile(m.pattern)
	if err != nil {
		return nil, err
	}

	s, ok := actual.(string)
	if !ok {
		return nil, fmt.Errorf("'%v' (%T) is not a string", actual, actual)
	}

	if !r.MatchString(s) {
		return nil, fmt.Errorf("%s does not match pattern %s", s, m.pattern)
	}

	return nil, nil
}
