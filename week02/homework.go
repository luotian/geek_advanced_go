package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/pkg/errors"
)

type MyDao struct {
}

type DaoData struct {
	val int
}

func (*MyDao) fakeQuery() (*DaoData, error) {
	ret_val := rand.Int31()
	if ret_val%2 == 0 {
		return nil, sql.ErrNoRows
	}

	return &DaoData{int(ret_val)}, nil
}

func init() {
	fmt.Println("Homework For week02 : Error")
}

func LoadApp() error {
	var dao MyDao
	myval, err := dao.fakeQuery()
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}

		return errors.Errorf("query error: %w", err)
	}

	fmt.Printf("myval is :%v\n", myval.val)
	return nil
}

func main() {
	err := LoadApp()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("continue run app")
}
