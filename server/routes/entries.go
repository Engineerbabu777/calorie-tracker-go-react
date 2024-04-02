package routes

import (
	"awais-go-react-calorie/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

func AddEntry(c *gin.Context) {
	// USING TIME OUT!
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

	// CREATING NEW VARIABALE!
	var entry models.Entry

	// BINDING DATA TO THE VARIABLE FROM BODY!
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	// VALIDATING OUR STRUCT!
	validationerr := validate.Struct(entry)
	if validationerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationerr.Error()})
		fmt.Println(validationerr)
		return
	}

	// CREATING NEW ID!
	entry.ID = primitive.NewObjectID()

	// INSERTED NEW ENTRY!
	inserted, err := entryCollection.InsertOne(ctx, entry)

	// HANDLING ERROR!
	if err != nil {
		msg := fmt.Sprintf("Entry item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(err)
		return
	}

	// CANCELLING THE REQUEST
	defer cancel()
	// RETURN BACK THE INSERTED ID!
	c.JSOn(http.StatusOK, inserted)

}

func GetEntries(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entries []bson.M

	cursor, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Err()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Err()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)
}

func GetEntriesByIngredient(c *gin.Context) {

}

func GetEntryById(c *gin.Context) {

	EntryID := c.Params("id")
	docID, _ := primitive.ObjectIDFromHex(EntryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entry bson.M
	err := entryCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)

	defer cancel()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Err()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(entry)
	c.JSON(http.StatusOK, entry)
}

func UpdateIngredient(c *gin.Context) {

}

func UpdateEntry(c *gin.context) {

}

func DeleteEntry(c *gin.Context) {
	// GETTING ID FROM PARAMS USING CONTEXT!
	entryID := c.Params("id")

	// CONVERTING TO OBJECT ID!
	docID, _ := primitive.ObjectIDFromHex(entryID)

	// CREATING CONTEXT TIMEOUT SO THAT OUR SERVER DOES NOT GET STUCK!
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	// WILL EXECUTE BY THE END OF THE FUNCTION!
	defer cancel()

	// DELETING THE ENTRY IN THE COLLECTION USING ID!
	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})

	// HANDLING ERROR!
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while deleting the entry"})
		fmt.Println(err)
		return
	}

	// RETURNING RESPONSE BACK TO API CALL!
	c.JSON(http.StatusOK, result.DeletedCount)
}
