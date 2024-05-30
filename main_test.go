package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pincodeapp/controlers"
	"pincodeapp/models"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetPincodes(t *testing.T) {
	models.InitDatabase()

	r := gin.New()
	r.GET("/pincodes", controlers.GetPincodes)
	req, err := http.NewRequest("GET", "/pincodes", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreatePincode(t *testing.T) {
	// Initialize the database
	models.InitDatabase()

	r := gin.Default()
	r.POST("/pincodes", controlers.CreatePincode)

	newPincode := models.Pincode{
		Officename:     "Test Office",
		Pincode:        "123456",
		OfficeType:     "Test Type",
		Deliverystatus: "Yes",
	}
	jsonBody, err := json.Marshal(newPincode)
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}
	req, err := http.NewRequest("POST", "/pincodes", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}
func TestUpdatePincode(t *testing.T) {
	// Initialize the test database
	models.InitDatabase()

	// Create a pincode to update
	initialPincode := models.Pincode{
		Officename:     "Initial Office",
		Pincode:        "123456",
		OfficeType:     "Initial Type",
		Deliverystatus: "No",
	}
	models.DB.Create(&initialPincode)

	r := gin.Default()
	r.PUT("/pincodes/:id", controlers.UpdatePincode)

	updatedPincode := models.Pincode{
		Officename:     "Updated Office",
		OfficeType:     "Updated Type",
		Deliverystatus: "Yes",
	}
	jsonBody, err := json.Marshal(updatedPincode)
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}
	req, err := http.NewRequest("PUT", "/pincodes/"+fmt.Sprint(initialPincode.ID), bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var updatedPincodeFromDB models.Pincode
	if err := models.DB.First(&updatedPincodeFromDB, initialPincode.ID).Error; err != nil {
		t.Fatalf("failed to retrieve updated pincode: %v", err)
	}

	if updatedPincodeFromDB.Officename != updatedPincode.Officename ||
		updatedPincodeFromDB.OfficeType != updatedPincode.OfficeType ||
		updatedPincodeFromDB.Deliverystatus != updatedPincode.Deliverystatus {
		t.Errorf("pincode not updated correctly: got %+v want %+v", updatedPincodeFromDB, updatedPincode)
	}
}
