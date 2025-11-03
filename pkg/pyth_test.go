package pkg

import "testing"

func TestPyth(t *testing.T) {
	id := "0xeaa020c61cc479712813461ce153894a96a6c00b21ed0cfc2798d1f9a9e9c94a"
	data, err := GetPythPrice(id, 1751628726)
	if err != nil {
		t.Fatal(err)
	}
	if data.Price.String() != "99995525" {
		t.Fatalf("expected price 99995525, got %s", data.Price.String())
	}
	if data.F != .99995525 {
		t.Fatalf("expected float 0.99995525, got %f", data.F)
	}
	if data.Id != id {
		t.Fatalf("expected id, got %s", data.Id)
	}
	if data.PublishTime == 0 {
		t.Fatalf("expected publish time, got %d", data.PublishTime)
	}
	if len(data.Data) == 1472 {
		t.Fatal("expected data length 1472, got", len(data.Data))
	}
}
func TestPyth2(t *testing.T) {
	id := "0x25a9be2a62a2269ce401b0ec5d5ae4a7e567a536cefb3153faae949316cbb7e6"
	data, err := GetPythPrice(id, 1762163141)
	if err != nil {
		t.Fatal(err)
	}
	if data.Price.String() != "3774163" {
		t.Fatalf("expected price 3774163, got %s", data.Price.String())
	}
	if data.F != .03774163 {
		t.Fatalf("expected float 0.3774163, got %f", data.F)
	}
	if data.Id != id {
		t.Fatalf("expected id, got %s", data.Id)
	}
	if data.PublishTime == 0 {
		t.Fatalf("expected publish time, got %d", data.PublishTime)
	}
	if len(data.Data) == 2622 {
		t.Fatal("expected data length 1472, got", len(data.Data))
	}
}
