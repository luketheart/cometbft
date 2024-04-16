package main

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	nums := []int{1062, 1063, 1072, 1062, 1082, 1074, 1072, 1067, 1070, 1063, 1069, 1068, 1085, 1065, 1068, 1068, 1069, 1066, 1068, 1065, 1066, 1068, 1069, 1064, 1074, 1127, 1064, 1054, 1067, 1066, 1060, 1066, 1063, 1071, 1066, 1068, 1067, 1068, 1066, 1066, 1058, 1064, 1056, 1068, 1069, 1071, 1066, 1070, 1065, 1065, 1067, 1064, 1058, 1068, 1061, 1065, 1057, 1054, 1067, 1071, 1065, 1066, 1065, 1054, 1195, 1067, 1066, 1066, 1069, 1065, 1059, 1064, 1072, 1065}
	s := 0
	for _, n := range nums {
		s += n
	}
	fmt.Println(float32(s) / float32(len(nums)))
}
