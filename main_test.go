package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

// Aquí va tu definición de SuccintProof y otras estructuras

func TestUnmarshalSuccintProof(t *testing.T) {
	// JSON de prueba
	// Leer el archivo JSON
	body, err := ioutil.ReadFile("test.json")
	if err != nil {
		t.Fatalf("Error reading JSON file: %v", err)
	}

	// Deserializar JSON
	var proofs []SuccintProof
	if err := json.Unmarshal(body, &proofs); err != nil {
		t.Fatalf("Error parsing JSON: %v", err)
	}

	fmt.Printf("%+v\n", proofs)

	// Verificar el resultado
	if len(proofs) != 10 {
		t.Fatalf("Expected 10 proof, got %d", len(proofs))
	}

	// if proof.ProofRequest.Type != "req_bytes" {
	// 	t.Errorf("Expected proof_request type 'req_bytes', got '%s'", proof.ProofRequest.Type)
	// }
}
