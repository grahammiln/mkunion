package ast

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestScoreCalculation_Calculate(t *testing.T) {
	ast := HumanFriendlyRules{
		AtLeastOneOf: []FiledBoostRule{
			{
				"question.thanks": BoostRuleOneOf{
					ConstBoost: &ConstBoost{
						Boost: 3.0,
						RuleOneOf: RuleOneOf{
							Gt: 10,
						},
					},
				},
			},
		},
		MustMatch: []FiledRule{
			{
				"question.similarity": RuleOneOf{Gt: 0.98},
			},
		},
	}

	data := map[string]interface{}{
		"question": map[string]interface{}{
			"thanks":     22,
			"similarity": 0.99,
		},
	}

	calc := NewScoreCalculator()
	res := calc.Calculate(ast, data)
	assert.Equal(t, 3.0, res)
}

func TestCalculationForListOfResults(t *testing.T) {
	ast := HumanFriendlyRules{
		AtLeastOneOf: []FiledBoostRule{
			{
				"question.thanks": BoostRuleOneOf{
					ConstBoost: &ConstBoost{
						Boost: 3.0,
						RuleOneOf: RuleOneOf{
							Gt: 10,
						},
					},
				},
				"question.verified": BoostRuleOneOf{
					ConstBoost: &ConstBoost{
						Boost: 100.0,
						RuleOneOf: RuleOneOf{
							Eq: true,
						},
					},
				},
			},
		},
		MustMatch: []FiledRule{
			{
				"question.similarity": RuleOneOf{Gt: 0.98},
			},
		},
	}

	data := []map[string]interface{}{
		{
			"question": map[string]interface{}{
				"thanks":     22,
				"similarity": 0.99,
			},
		},
		{
			"question": map[string]interface{}{
				"thanks":     2,
				"similarity": 0.99,
			},
		},
		{
			"question": map[string]interface{}{
				"thanks":     22,
				"similarity": 0.7,
			},
		},
		{
			"question": map[string]interface{}{
				"thanks":     2,
				"similarity": 0.99,
				"verified":   true,
			},
		},
	}

	calc := NewScoreCalculator()
	for i, d := range data {
		score := calc.Calculate(ast, d)
		data[i]["score"] = score
	}

	assert.Equal(t, 3.0, data[0]["score"])
	assert.Equal(t, 0.0, data[1]["score"])
	assert.Equal(t, 0.0, data[2]["score"])
	assert.Equal(t, 100.0, data[3]["score"])

	// now sort by score
	sort.SliceStable(data, func(i, j int) bool {
		return data[i]["score"].(float64) > data[j]["score"].(float64)
	})

	// pick first result
	// notice that score 100 comes form last element, that is verified
	assert.Equal(t, 100.0, data[0]["score"])
}
