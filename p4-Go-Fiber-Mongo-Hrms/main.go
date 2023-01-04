/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: main.go
  @author: Lido
  @time: 2023-01-04 10:49
  @description:
*/
package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017/" + dbName

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalln(" bad mongoURL !")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		panic(err)
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func main() {

	if err := Connect(); err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()

	app.Get("/employee", GetAllEmployee)
	app.Post("/employee",CreateEmployee)
	app.Put("/employee/:id",UpdateEmployee)
	app.Delete("/employee/:id",DeleteEmployee)

	log.Println("starting at server 8082 port !")

	log.Fatalln(app.Listen(":8082"))

}



func GetAllEmployee(c *fiber.Ctx) error {

	query := bson.D{{}}
	cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var employees []Employee = make([]Employee, 0)
	err = cursor.All(c.Context(), &employees)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employees)
}

func CreateEmployee(c *fiber.Ctx) error {
	collection := mg.Db.Collection("employees")

	employee := new(Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	employee.ID = ""

	insertResult, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	//返回插入的结果
	filter := bson.D{{Key: "_id",Value: insertResult.InsertedID}}
	createRecord := collection.FindOne(c.Context(),filter)

	createdEmployee := &Employee{}
	createRecord.Decode(createdEmployee)

	return c.Status(200).JSON(createdEmployee)

}

func UpdateEmployee(c *fiber.Ctx) error{
	idParam := c.Params("id")

	employeeID,err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	employee := new(Employee)

	if err := c.BodyParser(employee); err != nil{
		return c.Status(500).SendString(err.Error())
	}

	query := bson.D{{Key: "_id",Value: employeeID}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name",Value: employee.Name},
				{Key: "age",Value: employee.Age},
				{Key: "salary",Value: employee.Salary},
			},
		},
	}

	err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(),query,update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments{
			return c.SendStatus(400)
		}

		return c.SendStatus(500)
	}

	employee.ID = idParam

	return c.Status(200).JSON(employee)
}

func DeleteEmployee(c *fiber.Ctx) error{

	employeeID,err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	query := bson.D{{Key: "_id",Value: employeeID}}

	res,err := mg.Db.Collection("employees").DeleteOne(c.Context(),&query)

	if err != nil {
		return c.SendStatus(500)
	}

	if res.DeletedCount < 1{
		return c.SendStatus(400)
	}

	return c.Status(200).JSON("record delete")
}