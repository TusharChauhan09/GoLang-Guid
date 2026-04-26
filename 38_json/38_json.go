// 38_json.go
// Topic: encoding/json — marshal, unmarshal, struct tags
//
// MARSHAL (Go -> JSON)
//   data, err := json.Marshal(v)
//   data, err := json.MarshalIndent(v, "", "  ")
//
// UNMARSHAL (JSON -> Go)
//   var v T
//   err := json.Unmarshal(data, &v)        // pass POINTER
//
// STRUCT TAGS
//   `json:"name"`               rename
//   `json:"name,omitempty"`     omit zero values
//   `json:"-"`                  exclude from JSON
//   `json:",string"`            encode as string
//
// ONLY EXPORTED FIELDS get marshalled (capital first letter).
//
// MAPS / SLICES / PRIMITIVES marshal directly.
//   map keys must be string (or implement TextMarshaler).
//
// STREAMING
//   enc := json.NewEncoder(w); enc.Encode(v)
//   dec := json.NewDecoder(r); dec.Decode(&v)
//
// CUSTOM (un)MARSHAL
//   func (t T) MarshalJSON() ([]byte, error)
//   func (t *T) UnmarshalJSON(data []byte) error
//
// ARBITRARY JSON
//   var v map[string]any            // generic decode
//   var v any                       // even more generic
//
// Run: go run 38_json.go

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Address struct {
	City string `json:"city"`
	Zip  string `json:"zip,omitempty"`
}

type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Password string   `json:"-"`            // never emitted
	Tags     []string `json:"tags"`
	Addr     Address  `json:"addr"`
}

func main() {
	u := User{
		ID:       1,
		Name:     "Ada",
		Password: "secret",
		Tags:     []string{"admin", "go"},
		Addr:     Address{City: "London"},
	}

	// Marshal
	data, _ := json.Marshal(u)
	fmt.Println(string(data))

	// Pretty
	pretty, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(pretty))

	// Unmarshal
	raw := `{"id":2,"name":"Bob","email":"b@x.com","tags":["dev"],"addr":{"city":"Paris","zip":"75000"}}`
	var got User
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", got)

	// Generic decode
	var any1 map[string]any
	json.Unmarshal([]byte(raw), &any1)
	fmt.Println(any1["name"], any1["addr"].(map[string]any)["city"])

	// Stream encode
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(u)

	// Stream decode (line-delimited JSON)
	stream := `{"id":1}
{"id":2}
{"id":3}`
	dec := json.NewDecoder(strings.NewReader(stream))
	for dec.More() {
		var u User
		dec.Decode(&u)
		fmt.Println("decoded id:", u.ID)
	}
}
