package main

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

func sqlErr() error {
	return errors.Wrap(sql.ErrNoRows, "sql failed")
}

func dao() error {
	return errors.WithMessage(sqlErr(), "dao failed")
}

func main() {
	err := dao()
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("data not found, %v\n", err)
		log.Printf("%+v\n", err)
		return
	}

}
