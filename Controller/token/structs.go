package token

import "time"

type TokenST struct {
	Erro         bool
	Status       int
	Response     string
	DataToken    time.Time
	DataValidade time.Time
	CDE          string
	TPI          string
	IDT          string
	A1           TokenA1ST
	ToBytesJson  []byte
}

type token2ST struct {
	B1 *string
	B2 *string
	B3 *tokenB3ST
	B4 *string
	B5 *string
	B6 *string
}

type TokenSendST struct {
	CHVA string
	DVS  string
	CDE  string
	TPI  string
	IDT  string
}

type tokenB3ST struct {
	B1 *string
	B2 *string
	B3 *string
	B4 *string
	B5 *string
	B6 *string
}

type ReturnTokenDecodeST struct {
	Erro     bool
	Status   int
	Response string
	Data     time.Time
	Validade time.Time
	CDE      string
	TPI      string
	IDT      string
	Token    []byte
}

type TokenA1B3ST struct {
	B1 string
	B2 string
	B3 string
	B4 string
	B5 string
	B6 string
}

type TokenA1ST struct {
	B1 string
	B2 string
	B3 TokenA1B3ST
	B4 string
}

type ReturnTokenEncodeST struct {
	Erro     bool
	Status   int
	Response string
	A1       TokenA1ST
}

type BlocosST struct {
	B1 string
	B2 string
	B3 string
	B4 string
	B5 string
	B6 string
}

type TokenReciveST struct {
	CHVA *string
	DVS  *string
	CDE  *string
	TPI  *string
	IDT  *string
}
