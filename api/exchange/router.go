package exchangerate

import "net/http"

func RouterInit() {
	handler := NewHandler()

	http.HandleFunc("GET /exchange", handler.Get)
}
