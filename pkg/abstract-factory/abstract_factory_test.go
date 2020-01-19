package abstract_factory

import (
	"testing"

	"github.com/jayhrat/intern/pkg/abstract-factory/modern"
	"github.com/jayhrat/intern/pkg/abstract-factory/victorian"
)

func TestFactory(t *testing.T) {
	modernFactory := modern.NewModernFurnitureFactory()
	victorianFactory := victorian.NewVictorianFurnitureFactory()

	wantModernChair := "Sitting on the modern chair"
	wantModernCoffeeTable := "Put item on the modernCoffeeTable"
	wantModernSofa := "Lying on the modernSofa"

	modernChair := modernFactory.CreateChair()
	modernCoffeeTable := modernFactory.CreateCoffeeTable()
	modernSofa := modernFactory.CreateSofa()

	wantVictorianChair := "Sitting on the victorian chair"
	wantVictorianCoffeeTable := "Put item on the modernCoffeeTable"
	wantVictorianSofa := "Lying on the victorianSofa"

	victorianChair := victorianFactory.CreateChair()
	victorianCoffeeTable := victorianFactory.CreateCoffeeTable()
	victorianSofa := victorianFactory.CreateSofa()

	if got := modernChair.SitDown(); wantModernChair != got {
		t.Errorf("expected %q, got %q", wantModernChair, got)
	}
	if got := modernCoffeeTable.Put(); wantModernCoffeeTable != got {
		t.Errorf("expected %q, got %q", wantModernCoffeeTable, got)
	}
	if got := modernSofa.LieDown(); wantModernSofa != got {
		t.Errorf("expected %q, got %q", wantModernSofa, got)
	}

	if got := victorianChair.SitDown(); wantVictorianChair != got {
		t.Errorf("expected %q, got %q", wantVictorianChair, got)
	}
	if got := victorianCoffeeTable.Put(); wantVictorianCoffeeTable != got {
		t.Errorf("expected %q, got %q", wantVictorianCoffeeTable, got)
	}
	if got := victorianSofa.LieDown(); wantVictorianSofa != got {
		t.Errorf("expected %q, got %q", wantVictorianSofa, got)
	}

}
