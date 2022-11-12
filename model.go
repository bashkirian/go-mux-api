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
	OrderId int `json:"orderId"`
	UserId int `json:"userId"`
	ServiceId int `json:"serviceId"`
	Cost float64 `json:"cost"`
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
	return db.QueryRow("SELECT ruble_balance FROM balance WHERE user_id=$1",
		w.ID).Scan(&w.Balance)
}

// обновление баланса
func (w *wallet) updateBalance(db *sql.DB) error {
	_, err :=
		db.Exec(`INSERT INTO balance(user_id, ruble_balance) VALUES($2, $1) 
				ON CONFLICT(user_id) DO UPDATE 
					SET ruble_balance = balance.ruble_balance + $1;`,
				w.Balance, w.ID)
	return err
}

// создание резервации
func (rq *reserveQuery) makeReservation(db * sql.DB) error {
	_, err := db.Exec(`INSERT INTO reservations VALUES ($1, $2, $3, $4)`, rq.OrderId, rq.UserId, rq.ServiceId, rq.Cost)
	return err
}

// подтверждение резервации при наличии средств
func (rq *reserveQuery) confirmReservation(db * sql.DB) error {
	// _, err2 := db.Exec(`DELETE FROM reservations WHERE reservation_id = $1 AND service_id = $2`, 
	// 					rq.OrderId, rq.ServiceId)
	err2 := db.QueryRow(`SELECT cost FROM reservations r WHERE r.reservation_id = $1 AND r.service_id = $2 AND r.cost = $3;`, 
					    rq.OrderId, rq.ServiceId, rq.Cost).Scan(&rq.Cost)
	if err2 != nil {
		return err2;
	}
	_, err1 := db.Exec(`UPDATE balance SET ruble_balance = balance.ruble_balance - $1 WHERE user_id = $2`,
					   rq.Cost, rq.UserId)  
    return err1;
}
