package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrabNumbersFromLine(t *testing.T) {
	nums := grabNumbersFromLine("41 48  1 8 9")

	fmt.Println(nums)
	assert.True(t, reflect.DeepEqual(nums, []int{41, 48, 1, 8, 9}))
}

func TestCalculateCardScore(t *testing.T) {

	score := calculateCardScore([]int{83, 86, 6, 31, 17, 9, 48, 53}, []int{41, 48, 83, 86, 17})
	assert.Equal(t, 8, score)

	score = calculateCardScore([]int{47, 35, 45, 34, 36, 22, 52, 15}, []int{12, 15, 25, 16, 73})
	assert.Equal(t, 1, score)

	score = calculateCardScore([]int{12, 41, 51, 12, 31}, []int{89, 89, 78, 97, 68, 97, 68, 96})
	assert.Equal(t, 0, score)

}
