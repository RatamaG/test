package calculadora

import "testing"

func TestSuma(t *testing.T) {

	valor := Suma(7, 23)
	
	if valor != 30 {
	
		t.Errorf("Se esperaba %d y tiene %d",10,valor)
	
	}
}
