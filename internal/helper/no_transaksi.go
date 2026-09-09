package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

func GenerateNoTransaksi(db *gorm.DB, tablePrefix string, tanggal time.Time) (string, error) {
	ym := fmt.Sprintf("%d%02d", tanggal.Year(), int(tanggal.Month()))
	searchPattern := tablePrefix + "-" + ym + "%"

	var lastNo string
	err := db.Raw(
		"SELECT no_transaksi FROM "+getTableForPrefix(tablePrefix)+
			" WHERE no_transaksi LIKE ? ORDER BY no_transaksi DESC LIMIT 1",
		searchPattern,
	).Scan(&lastNo).Error

	var nextNum int
	if err == nil && len(lastNo) > 0 {
		parts := strings.Split(lastNo, "-")
		if len(parts) == 2 {
			numStr := parts[1]
			if len(numStr) > 6 {
				numStr = numStr[6:]
			}
			if n, parseErr := strconv.Atoi(numStr); parseErr == nil {
				nextNum = n
			}
		}
	}

	nextNum++
	noTransaksi := fmt.Sprintf("%s-%s%04d", tablePrefix, ym, nextNum)

	var exists int
	db.Raw(
		"SELECT COUNT(*) FROM "+getTableForPrefix(tablePrefix)+
			" WHERE no_transaksi = ?",
		noTransaksi,
	).Scan(&exists)

	if exists > 0 {
		nextNum++
		noTransaksi = fmt.Sprintf("%s-%s%04d", tablePrefix, ym, nextNum)
	}

	return noTransaksi, nil
}

func getTableForPrefix(prefix string) string {
	switch prefix {
	case "P":
		return "data_pembelian"
	case "H":
		return "data_bayar_hutang"
	default:
		return "data_pembelian"
	}
}
