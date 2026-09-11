package pre_order

import (
	"mrcoco/internal/models"
	"time"
)

type PreOrderResponse struct {
	ID         uint        `json:"id"`
	Tanggal    models.Date `json:"tanggal"`
	KonsumenID uint        `json:"konsumen_id"`
	Konsumen   string      `json:"konsumen"`
	TotalQty   int         `json:"total_qty"`
	GrandTotal int         `json:"grand_total"`
	Catatan    string      `json:"catatan"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type PreOrderItemResponse struct {
	ID       uint   `json:"id"`
	ItemID   uint   `json:"item_id"`
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	Satuan   string `json:"satuan"`
	Qty      int    `json:"qty"`
	Harga    int    `json:"harga"`
	Subtotal int    `json:"subtotal"`
}

type PreOrderDetailResponse struct {
	PreOrderResponse
	Items []PreOrderItemResponse `json:"items"`
}

type PreOrderCreateResponse struct {
	ID         uint   `json:"id"`
	Tanggal    string `json:"tanggal"`
	GrandTotal int    `json:"grand_total"`
	Catatan    string `json:"catatan"`
}

func ToPreOrderResponse(p *DataPreOrder) PreOrderResponse {
	resp := PreOrderResponse{
		ID:         p.ID,
		Tanggal:    p.Tanggal,
		KonsumenID: p.KonsumenID,
		GrandTotal: p.GrandTotal,
		Catatan:    p.Catatan,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
	if p.Konsumen != nil {
		resp.Konsumen = p.Konsumen.Nama
	}
	for _, item := range p.Items {
		resp.TotalQty += item.Qty
	}
	return resp
}

func ToPreOrderResponseList(preOrders []DataPreOrder) []PreOrderResponse {
	result := make([]PreOrderResponse, len(preOrders))
	for i, p := range preOrders {
		result[i] = ToPreOrderResponse(&p)
	}
	return result
}

func ToPreOrderItemResponse(item *DataPreOrderItem) PreOrderItemResponse {
	resp := PreOrderItemResponse{
		ID:       item.ID,
		ItemID:   item.ItemID,
		Qty:      item.Qty,
		Harga:    item.Harga,
		Subtotal: item.Subtotal,
	}
	if item.Item != nil {
		resp.Kode = item.Item.Kode
		resp.Nama = item.Item.Nama
		if item.Item.Satuan != nil {
			resp.Satuan = item.Item.Satuan.Nama
		}
	}
	return resp
}

func ToPreOrderDetailResponse(p *DataPreOrder) PreOrderDetailResponse {
	resp := PreOrderDetailResponse{
		PreOrderResponse: ToPreOrderResponse(p),
		Items:            make([]PreOrderItemResponse, len(p.Items)),
	}
	for i, item := range p.Items {
		resp.Items[i] = ToPreOrderItemResponse(&item)
	}
	return resp
}

func ToPreOrderCreateResponse(p *DataPreOrder) PreOrderCreateResponse {
	return PreOrderCreateResponse{
		ID:         p.ID,
		Tanggal:    p.Tanggal.Format("2006-01-02"),
		GrandTotal: p.GrandTotal,
		Catatan:    p.Catatan,
	}
}
