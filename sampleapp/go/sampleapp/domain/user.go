type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // フィールドサイズを255にセットします
	MemberNumber *string `gorm:"unique;not null"` // MemberNumberをuniqueかつnot nullにセットします
	Num          int     `gorm:"AUTO_INCREMENT"` // Numを自動インクリメントにセットします
	Address      string  `gorm:"index:addr"` // `addr`という名前のインデックスを作ります
	IgnoreMe     int     `gorm:"-"` // このフィールドは無視します
}