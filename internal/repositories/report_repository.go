package repositories

import (
	"database/sql"

	"kasir-api/internal/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetTodaySummary() (*models.SalesSummary, error) {
	var summary models.SalesSummary

	// total revenue & total transaksi
	err := r.db.QueryRow(`
		SELECT
			COALESCE(SUM(total_amount), 0),
			COUNT(*)
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&summary.TotalRevenue, &summary.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// produk terlaris
	err = r.db.QueryRow(`
		SELECT
			p.name,
			SUM(td.quantity)
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY SUM(td.quantity) DESC
		LIMIT 1
	`).Scan(&summary.ProdukTerlaris.Nama, &summary.ProdukTerlaris.QtyTerjual)

	// kalau belum ada transaksi hari ini
	if err == sql.ErrNoRows {
		summary.ProdukTerlaris = models.BestProduct{}
		return &summary, nil
	}

	if err != nil {
		return nil, err
	}

	return &summary, nil
}