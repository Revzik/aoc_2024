package day3

import "testing"

func TestExtractMultiplications(t *testing.T) {
	multiplicationString := "mul(123,2)"

	a, b := extractNumbers(multiplicationString)

	if a != 123 || b != 2 {
		t.Fatalf("Expected 123 and 2 to be found %s, got %d and %d", multiplicationString, a, b)
	}
}

func TestComputeMultiplications(t *testing.T) {
	multiplicationString := "mul(46,856)"
	expectedResult := 39376

	result := computeMultiplications(multiplicationString)

	if result != expectedResult {
		t.Fatalf("Expected %d to be result of computing %s, got %d", expectedResult, multiplicationString, result)
	}
}

func TestComputeAllMultiplications(t *testing.T) {
	instruction := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expectedResult := 161

	result := computeMultiplications(instruction)

	if result != expectedResult {
		t.Fatalf("Expected %d to be result of computing %s, got %d", expectedResult, instruction, result)
	}
}

func TestComputeMultiplicationsDoEnabler(t *testing.T) {
	instruction := "do();56-mul(2,6)"
	expectedResult := 12

	result, _ := computeWithEnablers(instruction, DoEnabler)

	if result != expectedResult {
		t.Fatalf("Expected %d to be result of computing %s, got %d", expectedResult, instruction, result)
	}
}

func TestComputeMultiplicationsDontEnabler(t *testing.T) {
	instruction := "don't();56-mul(2,6)"
	expectedResult := 0

	result, _ := computeWithEnablers(instruction, DoEnabler)

	if result != expectedResult {
		t.Fatalf("Expected %d to be result of computing %s, got %d", expectedResult, instruction, result)
	}
}

func TestComputeAllMultiplicationsWithEnablers(t *testing.T) {
	instruction := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}
	expectedResult := 48

	result := computeAllMultiplicationsWithEnablers(instruction)

	if result != expectedResult {
		t.Fatalf("Expected %d to be result of computing %s, got %d", expectedResult, instruction, result)
	}
}
