package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*

sql.Open()
opens a registered database driver. The Go-MySQL-Driver registered the mysql driver here.
The second argument is the DSN (Data Source Name) that defines information pertaining to the database connection.
It supports following formats:

  user@unix(/path/to/socket)/dbname?charset=utf8
  user:password@tcp(localhost:5555)/dbname?charset=utf8
  user:password@/dbname
  user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

db.Prepare()
returns a SQL operation that is going to be executed. It also returns the execution status after executing SQL.

db.Query()
executes SQL and returns a Rows result.

stmt.Exec()
executes SQL that has been prepared and stored in Stmt.

Note that we use the format =? to pass arguments.
This is necessary for preventing SQL injection attacks.

*/

func main() {
	db, err := sql.Open("mysql", "astaxie:astaxie@/test?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
