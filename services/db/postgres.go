package db

import (
	"example/graphql/graph/model"
	"example/graphql/ports"
	"example/graphql/utils"
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type DbService struct {
	dB *sqlx.DB
}

func NewDbService() ports.HackerDB {

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		db.Close()
		panic(err)
	}
	slog.Info("Connected to database")
	return &DbService{
		dB: db,
	}
}

func (d *DbService) Close() {
	d.dB.Close()
}

func (d *DbService) Migrate() error {
	driver, err := postgres.WithInstance(d.dB.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (d *DbService) CreateLink(link *model.Link) (*model.Link, error) {
	var id string
	err := d.dB.Get(&id, `INSERT into links(Title,Address,UserID) VALUES ($1,$2,$3) RETURNING ID`, link.Title, link.Address, link.User.ID)
	link.ID = id
	return link, err
}

func (d *DbService) AllLinks(userid int) ([]*model.Link, error) {

	var info []struct {
		Id       string `db:"id"`
		Title    string `db:"title"`
		Addr     string `db:"address"`
		Userid   string `db:"userid"`
		UserName string `db:"username"`
	}
	err := d.dB.Select(&info, `
		SELECT links.id, title, address ,users.id as userid, users.Username as username 
		FROM links inner join users 	on links.userid = users.id
		where users.id = $1`, userid)

	if err != nil {
		slog.Error("failed to select links", slog.Any("err", err))
		return nil, err
	}
	links := make([]*model.Link, 0, len(info))
	for _, v := range info {
		links = append(links, &model.Link{ID: v.Id, Title: v.Title, Address: v.Addr, User: &model.User{
			ID:   v.Userid,
			Name: v.UserName,
		}})
	}
	return links, err
}
func (d *DbService) CreateUser(user model.NewUser) (*model.User, error) {
	var id string
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	err = d.dB.Get(&id, `INSERT into users(Username,Password) VALUES ($1,$2) RETURNING ID`, user.Username, password)
	return &model.User{ID: id, Name: user.Username}, err
}

func (d *DbService) UserByName(name string) (int, error) {
	var id int
	err := d.dB.Get(&id, "SELECT id FROM users WHERE username = $1", name)
	return id, err
}

func (d *DbService) AuthenticateUser(username string, password string) bool {
	var hashedPassword string
	err := d.dB.Get(&hashedPassword, "SELECT password FROM users WHERE username = $1 ", username)
	if err != nil {
		return false
	}
	return utils.PasswordMatch(password, hashedPassword)
}
