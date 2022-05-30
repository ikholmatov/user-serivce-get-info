package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	pb "github.com/venomuz/project2/genproto"
	"log"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	UserQuery := `INSERT INTO users VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := r.db.Exec(UserQuery, user.ID, user.FirstName, user.LastName, pq.Array(user.Email), user.Bio, pq.Array(user.PhoneNumber), user.TypeID, user.Status)
	if err != nil {
		log.Panicf("%s\n%s", "Error while users to table addresses", err)
	}
	AddressQuery := `INSERT INTO addresses VALUES($1,$2,$3,$4,$5,$6)`
	_, err = r.db.Exec(AddressQuery, user.Addr.ID, user.ID, user.Addr.Country, user.Addr.City, user.Addr.District, user.Addr.PostalCode)
	if err != nil {
		log.Panicf("%s\n%s", "Error while inserting to table addresses", err)
	}

	return user, nil
}
func (r *userRepo) GetByID(ID string) (*pb.User, error) {
	a := pb.User{}
	GetUsers := `SELECT * FROM users WHERE id = $1`
	err := r.db.QueryRow(GetUsers, ID).Scan(&a.ID, &a.FirstName, &a.LastName, pq.Array(&a.Email), &a.Bio, pq.Array(&a.PhoneNumber), &a.TypeID, &a.Status)
	if err != nil {
		log.Panicf("%s\n%s", "Error while geting data from table users", err)
	}

	GetAddresses := `SELECT id FROM addresses WHERE userid = $1`
	err = r.db.QueryRow(GetAddresses, ID).Scan(&a.Addr.ID)
	fmt.Println(a)
	if err != nil {
		log.Panicf("%s\n%s", "Error while geting data from table addresses", err)
	}

	return &a, err
}
