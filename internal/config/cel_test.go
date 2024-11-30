package config

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCelAstMap(t *testing.T) {
	t.Run("invalid cel ast map", func(t *testing.T) {
		invalid_cel_filter := make([]interface{}, 1)
		invalid_cel_filter = append(invalid_cel_filter, map[string]string{"operation": "insert", "condition": "'hello'"})

		c := FilterStruct{Tables: map[string][]interface{}{"users": invalid_cel_filter}}

		_, err := NewCelAstMap(c)

		assert.EqualError(t, err, errors.New("output must be boolean").Error())
	})

	t.Run("valid cel ast map", func(t *testing.T) {
		valid_cel_filter := make([]interface{}, 1)
		valid_cel_filter = append(valid_cel_filter, map[string]string{"operation": "insert", "condition": "row.status == 'active'"})

		c := FilterStruct{Tables: map[string][]interface{}{"users": valid_cel_filter}}

		_, err := NewCelAstMap(c)

		assert.Nil(t, err)
	})
}
