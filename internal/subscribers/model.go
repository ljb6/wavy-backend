package subscribers

type Subscribers struct {
	ID      int    `jsond:"id"`
	User_ID int    `json:"user_id"`
	Name    string `jsond:"name"`
	Email   string `jsond:"email"`
}

type SubRequest struct {
	Name  string `jsond:"name"`
	Email string `jsond:"email"`
}
