package main

import (
	// "time"
	"database/sql"
)

type wallet struct {
	ID    int     `json:"user_id"`
	Balance float64  `json:"ruble_balance"`
}

// для бухгалтерии
// type reservation struct {
// 	ID int
// 	UserID int
// 	ServiceID int
// 	Cost int
// 	Reservation_Time time.Time
// }

type Service struct {
	ID int
	Name string
}

type reserveQuery struct {
	userId int `json:"userId"`
	serviceId int `json:"serviceId"`
	orderId int `json:"orderId"`
	cost float64 `json:"cost"`
}

// для бухгалтерии
// type StatementElem struct {
// 	RecordTime   time.Time       `json:"recordTime"`
// 	TransferType string          `json:"transferType"`
// 	Amount       decimal.Decimal `json:"amount"`
// 	Description  string          `json:"description"`
// }

// получение баланса
func (w *wallet) getWallet(db *sql.DB) error {
	return db.QueryRow("SELECT ruble_balance FROM balance WHERE id=$1",
		w.ID).Scan(&w.Balance)
}

// обновление баланса
func (w *wallet) updateBalance(db *sql.DB) error {
	_, err :=
		db.Exec("INSERT INTO balance(ruble_balance) VALUES($1) ON DUPLICATE KEY UPDATE balance SET ruble_balance= ruble_balance + $1 WHERE id=$2",
				w.Balance, w.ID)
	return err
}

func (rq *reserveQuery) makeReservation(db * sql.DB) error {
	_, err := db.Exec(`INSERT INTO reservations(reservation_id, user_id, service_id, cost) 
						VALUES ($1, $2, $3, $4)`, rq.orderId, rq.userId, rq.serviceId, rq.cost)
	return err
}

func (rq *reserveQuery) confirmReservation(db * sql.DB) error {
	_, err := db.Exec(`UPDATE balance SET ruble_balance = ruble_balance - $1 WHERE id = $2;
					   DELETE FROM reservations WHERE reservation_id = $3`,
					   rq.cost, rq.userId, rq.orderId)  
	return err;
}
