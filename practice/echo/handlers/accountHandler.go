type AccountHandler struct {
    au usecases.AccountUsecase
}


func NewAccountHandler(au usecases.AccountUsecase) *AccountHandler {
    return &AccountHandler{au}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
    err := h.au.CreateTodo(account)
}