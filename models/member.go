package models

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/lib/pq"
	u "go-member-api/utils"
	"log"
	"os"
	"strconv"
	"time"
)

type Member struct {
	first_name       string `csv:"first_name"`
	last_name        string `csv:"last_name"`
	phone_number     string `csv:"phone_number"`
	client_member_id int    `csv:"client_member_id"`
	account_id       int    `csv:"account_id"`
}

type MemberWithID struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PhoneNumber     string `json:"phone_number"`
	ClientMemberID  int    `json:"client_member_id"`
	AccountID       int    `json:"account_id"`
}

/*
Validate required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (member *Member) Validate() (map[string]interface{}, bool) {
	if member.first_name == "" {
		return u.Message(false, "Member first name should be on the payload"), false
	}

	if member.last_name == "" {
		return u.Message(false, "Member last name should be on the payload"), false
	}

	if member.phone_number == "" {
		return u.Message(false, "Member phone number should be on the payload"), false
	}

	if member.client_member_id == 0 {
		return u.Message(false, "Client member id should be on the payload"), false
	}

	if member.account_id == 0 {
		return u.Message(false, "Account id should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func BatchCreateMethodI(inputFile string) error {
	members := readCSV(inputFile)
	a := time.Now()
	tx, err := GetDB().Begin()
	if err != nil {
		handleError(err)
	}

	txOK := false
	defer func() {
		if !txOK {
			tx.Rollback()
		}
	}()

	// The ON COMMIT DROP clause at the end makes sure that the table
	// is cleaned up at the end of the transaction.
	// While the "for{..} state machine" goroutine in charge of delayed
	sqlFDataMakeTempTable := `CREATE TEMPORARY TABLE fstore_data_load
  (id serial, first_name VARCHAR(50),last_name VARCHAR(50),phone_number integer, client_member_id integer, account_id integer)
	ON COMMIT DROP`

	// saving ensures this function is not running twice at any given time.
	_, err = tx.Exec(sqlFDataMakeTempTable)

	stmt, err := tx.Prepare(pq.CopyIn("member", "first_name", "last_name", "phone_number", "client_member_id", "account_id"))
	for _, member := range members {
		_, err := stmt.Exec(member.first_name, member.last_name, member.phone_number, member.client_member_id, member.account_id)
		if err != nil {
			handleError(err)
		}
	}

	sqlFDataSetFromTemp := `INSERT INTO member(first_name, last_name, phone_number, client_member_id, account_id)
	VALUES %s ON CONFLICT (phone_number, client_member_id) DO UPDATE SET phone_number = excluded.phone_number, client_member_id=EXCLUDED.client_member_id`
	_, err = tx.Exec(sqlFDataSetFromTemp)
	if err != nil {
		handleError(err)
	}

	err = tx.Commit()
	if err != nil {
		handleError(err)
	}
	txOK = true
	delta := time.Now().Sub(a)
	fmt.Println("all rows inserted successfully in %d nano seconds", delta.Nanoseconds())
	return nil
}

func BatchCreateMethodII(inputFile string) error {
	members := readCSV(inputFile)

	a := time.Now()

	mList := make([]Member, 0)
	for _, m := range members {
		mList = append(mList, m)
		if len(mList) == 1000 {
			Transact(mList)
			mList = mList[:0]
		}
		// Don't forget the last batch.
	}

	if len(mList) > 0 {
		Transact(mList)
	}
	delta := time.Now().Sub(a)
	fmt.Println("all rows inserted successfully in %d nano seconds", delta.Nanoseconds())
	return nil
}

func GetMemberByID(id int) (*MemberWithID, error) {
	member := &MemberWithID{}
	sqlStatement := `SELECT * FROM member WHERE id=$1`
	row := GetDB().QueryRow(sqlStatement, id)
	err := row.Scan(&member.ID, &member.FirstName, &member.LastName, &member.PhoneNumber, &member.ClientMemberID, &member.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")

		}
		handleError(err)
		return member, err

	}
	return member, nil
}

func GetMemberByPhoneNumber(phone_number string) (*MemberWithID, error) {
	member := &MemberWithID{}
	sqlStatement := `SELECT * FROM member WHERE phone_number=$1`
	row := GetDB().QueryRow(sqlStatement, phone_number)
	err := row.Scan(&member.ID, &member.FirstName, &member.LastName, &member.PhoneNumber, &member.ClientMemberID, &member.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		}
		handleError(err)
		return member, err

	}
	return member, nil
}

func GetMemberByClientMemberID(client_id int) (*MemberWithID, error) {
	member := &MemberWithID{}
	sqlStatement := `SELECT * FROM member WHERE client_member_id=$1`
	row := GetDB().QueryRow(sqlStatement, client_id)
	err := row.Scan(&member.ID, &member.FirstName, &member.LastName, &member.PhoneNumber, &member.ClientMemberID, &member.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		}
		handleError(err)
		return member, err

	}
	return member, nil
}

func GetMembersByAccountID(account_id int) ([]*MemberWithID, error) {
  members := make([]*MemberWithID, 0)
	sqlStatement := `SELECT * FROM member WHERE account_id=$1`
	rows, err := GetDB().Query(sqlStatement, account_id)
	if err != nil {
		fmt.Println(err)
		return members, err
	}

	defer rows.Close()
	for rows.Next() {
		member := MemberWithID{}
		err := rows.Scan(&member.ID, &member.FirstName, &member.LastName, &member.PhoneNumber, &member.ClientMemberID, &member.AccountID)
		if err != nil {
			fmt.Println(err)
			return members, err
		}
		members = append(members, &member)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return members, err
	}

	return members, nil
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func readCSV(inputFile string) []Member {
	csvfile, err := os.Open(inputFile)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	reader := csv.NewReader(csvfile)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	members := make([]Member, 0)

	for i, record := range records {
		if i == 0 {
			// skip header line
			continue
		}
		account_id, err := strconv.Atoi(record[4])
		if err != nil {
			log.Fatal(err)
		}

		client_member_id, err := strconv.Atoi(record[5])
		if err != nil {
			log.Fatal(err)
		}

		members = append(members, Member{
			first_name:       record[1],
			last_name:        record[2],
			phone_number:     record[3],
			client_member_id: account_id,
			account_id:       client_member_id})
	}
	return members
}

func Transact(mList []Member) {
	txn, err := GetDB().Begin()
	handleError(err)
	stmt, _ := txn.Prepare(pq.CopyIn("member", "first_name", "last_name", "phone_number", "client_member_id", "account_id"))
	for _, member := range mList {
		_, err := stmt.Exec(member.first_name, member.last_name, member.phone_number, member.client_member_id, member.account_id)
		if err != nil {
			handleError(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		txn.Rollback()
		handleError(err)
	}
	err = stmt.Close()
	handleError(err)
	err = txn.Commit()
	handleError(err)
}
