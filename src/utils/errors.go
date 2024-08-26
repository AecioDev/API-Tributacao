package utils

import (
	"api-tributacao/src/globals"
	"log"
)

func FatalOnErr(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// erro para evitar leak de mensagens de erro privadas
type InternalError struct {
	originalErr error
}

func NewInternalError(err error) InternalError {
	return InternalError{originalErr: err}
}

func (e InternalError) Error() string {
	if globals.DEV {
		if e.originalErr != nil {
			return e.originalErr.Error()
		}

		// talvez aqui faça sentido simpesmente evitar o check originalErr != nil
		// e chamar .Error() nele, mesmo que isso possivelmente de panic pois esse
		// é exatamente o comportamento que um nil value de uma interface tem, ter
		// um fallback para uma mensagem default aqui parece esquisito
		return "internal error [CODE 500]"
	}

	return "internal error"
}

func (e InternalError) Unwrap() error {
	return e.originalErr
}
