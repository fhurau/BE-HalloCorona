package usersdto

type CreateUserRequest struct {
	Name     string `gorm:"type: varchar(255)" json:"name" validate:"required"`
	Username string `gorm:"type: varchar(255)" json:"username" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	ListAs   string `gorm:"type: varchar(100)" json:"listAs" validate:"required"`
	Gender   string `gorm:"type: varchar(100)" json:"gender" validate:"required"`
	Phone    string `gorm:"type: varchar(100)" json:"phone" validate:"required"`
	Address  string `gorm:"type: text" json:"address" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Username string `gorm:"type: varchar(255)" json:"username"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	ListAs   string `gorm:"type: varchar(100)" json:"listAs"`
	Gender   string `gorm:"type: varchar(100)" json:"gender"`
	Phone    string `gorm:"type: varchar(100)" json:"phone"`
	Address  string `gorm:"type: text" json:"address"`
}
