package dao

import (
	"github.com/globalsign/mgo"
	. "supplyservice/models"
	"github.com/globalsign/mgo/bson"
	"log"
)

type DataObject struct {
	Server string
	Database string
}

var db *mgo.Database

const (
	PARTNERS = "parters"
	PRODUCTS = "products"
	CONSUMERS = "consumers"
	TRANSACTIONS = "transactions"
	TICKETS = "tickets"
)

func (d *DataObject) Connect() {
	session, err := mgo.Dial(d.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(d.Database)
}

func (d *DataObject) FindAllTickets() ([]Ticket, error) {
	var tickets []Ticket
	err := db.C(TICKETS).Find(bson.M{}).All(&tickets)
	return tickets, err
}

func (d *DataObject) FindTicketById(id string) (Ticket, error) {
	var ticket Ticket
	err := db.C(TICKETS).FindId(bson.ObjectIdHex(id)).One(&ticket)
	return ticket, err
}

func (d *DataObject) InsertTicket(ticket Ticket) error {
	err := db.C(TICKETS).Insert(&ticket)
	return err
}

func (d *DataObject) DeleteTicket(ticket Ticket) error{
	err := db.C(TICKETS).Remove(&ticket)
	return err
}

func (d *DataObject) UpdateTicket(ticket Ticket) error {
	err := db.C(TICKETS).UpdateId(ticket.ID, &ticket)
	return err
}

func (d *DataObject) FindAllTransactions() ([]Transaction, error) {
	var transactions []Transaction
	err := db.C(TRANSACTIONS).Find(bson.M{}).All(&transactions)
	return transactions, err
}

func (d *DataObject) FindTransactionById(id string) (Transaction, error) {
	var transaction Transaction
	err := db.C(TRANSACTIONS).FindId(bson.ObjectIdHex(id)).One(&transaction)
	return transaction, err
}

func (d *DataObject) InsertTransaction(transaction Transaction) error {
	err := db.C(TRANSACTIONS).Insert(&transaction)
	return err
}

func (d *DataObject) DeleteTransaction(transaction Transaction) error{
	err := db.C(TRANSACTIONS).Remove(&transaction)
	return err
}

func (d *DataObject) UpdateTransaction(transaction Transaction) error {
	err := db.C(TRANSACTIONS).UpdateId(transaction.ID, &transaction)
	return err
}

func (d *DataObject) FindAllConsumers() ([]Consumer, error) {
	var consumers []Consumer
	err := db.C(CONSUMERS).Find(bson.M{}).All(&consumers)
	return consumers, err
}

func (d *DataObject) FindConsumerById(id string) (Consumer, error) {
	var consumer Consumer
	err := db.C(CONSUMERS).FindId(bson.ObjectIdHex(id)).One(&consumer)
	return consumer, err
}

func (d *DataObject) InsertConsumer(consumer Consumer) error {
	err := db.C(CONSUMERS).Insert(&consumer)
	return err
}

func (d *DataObject) DeleteConsumer(consumer Consumer) error{
	err := db.C(CONSUMERS).Remove(&consumer)
	return err
}

func (d *DataObject) UpdateConsumer(consumer Consumer) error {
	err := db.C(CONSUMERS).UpdateId(consumer.ID, &consumer)
	return err
}

func (d *DataObject) FindAllProducts() ([]Product, error) {
	var products []Product
	err := db.C(PRODUCTS).Find(bson.M{}).All(&products)
	return products, err
}

func (d *DataObject) FindProductById(id string) (Product, error) {
	var product Product
	err := db.C(PRODUCTS).FindId(bson.ObjectIdHex(id)).One(&product)
	return product, err
}

func (d *DataObject) InsertProduct(product Product) error {
	err := db.C(PRODUCTS).Insert(&product)
	return err
}

func (d *DataObject) DeleteProduct(product Product) error{
	err := db.C(PRODUCTS).Remove(&product)
	return err
}

func (d *DataObject) UpdateProduct(product Product) error {
	err := db.C(PRODUCTS).UpdateId(product.ID, &product)
	return err
}

func (d *DataObject) FindAllPartners() ([]Partner, error) {
	var partners []Partner
	err := db.C(PARTNERS).Find(bson.M{}).All(&partners)
	return partners, err
}

func (d *DataObject) FindPartnerById(id string) (Partner, error) {
	var partner Partner
	err := db.C(PARTNERS).FindId(bson.ObjectIdHex(id)).One(&partner)
	return partner, err
}

func (d *DataObject) InsertPartner(partner Partner) error {
	err := db.C(PARTNERS).Insert(&partner)
	return err
}

func (d *DataObject) DeletePartner(partner Partner) error{
	err := db.C(PARTNERS).Remove(&partner)
	return err
}

func (d *DataObject) UpdatePartner(partner Partner) error {
	err := db.C(PARTNERS).UpdateId(partner.ID, &partner)
	return err
}
