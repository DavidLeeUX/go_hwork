/*
 * @Descripttion:
 * @version: xv1.0
 * @Author: changwei5
 * @Date: 2021-04-24 20:16:45
 * @LastEditors: changwei5
 * @LastEditTime: 2021-04-24 20:30:29
 */
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
