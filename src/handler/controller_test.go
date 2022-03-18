package controller

import (
	"testing"
)

func TestConverter1(t *testing.T) {
	converter := PlainConverter{}

	workdays, err := converter.Deserialize("03-22-2022\t07:45-16:15")

	if err != nil {
		t.Fatal(err)
	}

	if len(workdays) != 1 {
		t.Fatalf("Invalid workdays length %d", len(workdays))
	}

	if workdays[0].Date.Year() != 2022 {
		t.Fatalf("Invalid year %d", workdays[0].Date.Year())
	}

	if workdays[0].Date.Month() != 3 {
		t.Fatalf("Invalid month %d", workdays[0].Date.Month())
	}

	if workdays[0].Date.Day() != 22 {
		t.Fatalf("Invalid day %d", workdays[0].Date.Day())
	}

	if workdays[0].Begin.Hour != 7 {
		t.Fatalf("Invalid workdays begin hour %d", workdays[0].Begin.Hour)
	}

	if workdays[0].Begin.Min != 45 {
		t.Fatalf("Invalid workdays begin min %d", workdays[0].Begin.Min)
	}
	if workdays[0].End.Hour != 16 {
		t.Fatalf("Invalid workdays end hour %d", workdays[0].End.Hour)
	}

	if workdays[0].End.Min != 15 {
		t.Fatalf("Invalid workdays end min %d", workdays[0].End.Min)
	}
}

func TestConverter2(t *testing.T) {
	converter := PlainConverter{}

	workdays, err := converter.Deserialize("01-31-2022\t05:32-18:59")

	if err != nil {
		t.Fatal(err)
	}

	if len(workdays) != 1 {
		t.Fatalf("Invalid workdays length %d", len(workdays))
	}

	if workdays[0].Date.Year() != 2022 {
		t.Fatalf("Invalid year %d", workdays[0].Date.Year())
	}

	if workdays[0].Date.Month() != 1 {
		t.Fatalf("Invalid month %d", workdays[0].Date.Month())
	}

	if workdays[0].Date.Day() != 31 {
		t.Fatalf("Invalid day %d", workdays[0].Date.Day())
	}

	if workdays[0].Begin.Hour != 5 {
		t.Fatalf("Invalid workdays begin hour %d", workdays[0].Begin.Hour)
	}

	if workdays[0].Begin.Min != 32 {
		t.Fatalf("Invalid workdays begin min %d", workdays[0].Begin.Min)
	}
	if workdays[0].End.Hour != 18 {
		t.Fatalf("Invalid workdays end hour %d", workdays[0].End.Hour)
	}

	if workdays[0].End.Min != 59 {
		t.Fatalf("Invalid workdays end min %d", workdays[0].End.Min)
	}
}

func TestConverter3(t *testing.T) {
	converter := PlainConverter{}

	workdays, err := converter.Deserialize("01-31-2022\t05:32-18:59\r\n12-31-2023\t00:01-23:39")

	if err != nil {
		t.Fatal(err)
	}

	if len(workdays) != 2 {
		t.Fatalf("Invalid workdays length %d", len(workdays))
	}

	if workdays[1].Date.Year() != 2023 {
		t.Fatalf("Invalid year %d", workdays[0].Date.Year())
	}

	if workdays[1].Date.Month() != 12 {
		t.Fatalf("Invalid month %d", workdays[0].Date.Month())
	}

	if workdays[1].Date.Day() != 31 {
		t.Fatalf("Invalid day %d", workdays[0].Date.Day())
	}

	if workdays[1].Begin.Hour != 0 {
		t.Fatalf("Invalid workdays begin hour %d", workdays[0].Begin.Hour)
	}

	if workdays[1].Begin.Min != 1 {
		t.Fatalf("Invalid workdays begin min %d", workdays[0].Begin.Min)
	}
	if workdays[1].End.Hour != 23 {
		t.Fatalf("Invalid workdays end hour %d", workdays[0].End.Hour)
	}

	if workdays[1].End.Min != 39 {
		t.Fatalf("Invalid workdays end min %d", workdays[0].End.Min)
	}
}
