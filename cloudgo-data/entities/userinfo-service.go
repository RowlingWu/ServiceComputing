package entities

const tableName = "userinfo"

func Save(u *UserInfo) {
	_, err := mydb.Table(tableName).Insert(u)
	checkErr(err)
}

func FindByID(i int) *UserInfo {
	var user UserInfo
	_, err := mydb.Table(tableName).Where("uid=?", i).Get(&user)
	checkErr(err)
	return &user
}

func FindAll() *[]UserInfo {
	var users []UserInfo
	err := mydb.Table(tableName).Find(&users)
	checkErr(err)
	return &users
}
