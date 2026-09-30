package store

import (
	"testing"

	"github.com/deathcore1998/earthquake-cli/models"
)

func createDB(t *testing.T) *DB {
	t.Helper()

	db, err := NewDB(":memory:")
	if err != nil {
		t.Fatalf("NewDB failed: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return db
}

func createTestFeature(id string, mag float64, time int64) models.Feature {
	return models.Feature{
		ID: id,
		Properties: models.Properties{
			Mag:     mag,
			Place:   "Test save place",
			Time:    time,
			Tsunami: 0,
			URL:     "https://test.com",
			Status:  "reviewed",
		},
		Geometry: models.Geometry{
			Coordinates: []float64{5.0, 10.0, 15.0},
		},
	}
}

func TestSaveGetAll(t *testing.T) {
	db := createDB(t)
	feature := createTestFeature("testSave", 4.5, 1790765930586)

	err := db.Save(feature)
	if err != nil {
		t.Errorf("Save failed: %v", err)
	}

	features, err := db.GetAll("time", "desc")
	if err != nil {
		t.Errorf("Failed GetAll: %v", err)
	}

	if len(features) != 1 {
		t.Errorf("expected 1, got %d", len(features))
	}

	if feature.ID != features[0].ID {
		t.Errorf("ID: got %q, want %q", features[0].ID, feature.ID)
	}
}

func TestSaveDuplicate(t *testing.T) {
	db := createDB(t)
	feature := createTestFeature("testSave", 4.5, 1790765930586)

	err := db.Save(feature)
	if err != nil {
		t.Errorf("Failed first save: %v", err)
	}

	err = db.Save(feature)
	if err != nil {
		t.Errorf("Failed second save: %v", err)
	}

	features, err := db.GetAll("time", "desc")
	if err != nil {
		t.Errorf("Failed GetAll: %v", err)
	}

	if len(features) != 1 {
		t.Errorf("expected 1, got %d", len(features))
	}
}

func TestSorting(t *testing.T) {
	db := createDB(t)

	testFeatures := []models.Feature{
		createTestFeature("test-1", 3.0, 1000),
		createTestFeature("test-2", 4.0, 2000),
		createTestFeature("test-3", 5.0, 3000),
	}

	for index, feature := range testFeatures {
		err := db.Save(feature)
		if err != nil {
			t.Fatalf("Save %d failed: %v", index, err)
		}
	}

	features, err := db.GetAll("mag", "desc")
	if err != nil {
		t.Errorf("Failed GetAll: %v", err)
	}

	lenFeatures := len(features)
	if lenFeatures != 3 {
		t.Fatalf("expected 3 features, got %d", lenFeatures)
	}

	want := []float64{5.0, 4.0, 3.0}
	for i, f := range features {
		if f.Properties.Mag != want[i] {
			t.Errorf("features[%d].Mag: got %v, want %v", i, f.Properties.Mag, want[i])
		}
	}
}
