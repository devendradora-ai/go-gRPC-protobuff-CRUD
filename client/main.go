package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/go-gRPC-protobuff-CRUD/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	host = "127.0.0.1:50052"
)

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	crudClient := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()


	fmt.Println("Product Inventory CRUD")
	fmt.Print(" 1 : create \n 2 : read \n 3 : update \n 4: delete \n Enter choice :  ")

	choice := bufio.NewReader(os.Stdin)
	text, _ := choice.ReadString('\n')

	switch text {
	case "1\n":
		fmt.Print("\nbarcode: ")
		barcodeR := bufio.NewReader(os.Stdin)
		barcode, _ := barcodeR.ReadString('\n')
		barcode = strings.Trim(barcode, "\n")

		fmt.Print("\nstore: ")
		storeR := bufio.NewReader(os.Stdin)
		store, _ := storeR.ReadString('\n')
		store = strings.Trim(store, "\n")

		fmt.Print("\nproduct name: ")
		nameR := bufio.NewReader(os.Stdin)
		name, _ := nameR.ReadString('\n')
		name = strings.Trim(name, "\n")


		fmt.Print("\ncategory: ")
		categoryR := bufio.NewReader(os.Stdin)
		category, _ := categoryR.ReadString('\n')
		category = strings.Trim(category, "\n")


		fmt.Print("\ndescription: ")
		descriptionR := bufio.NewReader(os.Stdin)
		description, _ := descriptionR.ReadString('\n')
		description = strings.Trim(description, "\n")


		fmt.Print("\nimage: ")
		imageR := bufio.NewReader(os.Stdin)
		image, _ := imageR.ReadString('\n')
		image = strings.Trim(image, "\n")


		fmt.Print("\nis restricted: ")
		restrictedR := bufio.NewReader(os.Stdin)
		restricted, _ := restrictedR.ReadString('\n')
		restricted = strings.Trim(restricted, "\n")
		restrictedBool, err := strconv.ParseBool(restricted)


		fmt.Print("\ninventory: ")
		inventoryR := bufio.NewReader(os.Stdin)
		inventory, _ := inventoryR.ReadString('\n')
		inventory = strings.Trim(store, "\n")
		inventoryInt, err := strconv.Atoi(inventory)


		fmt.Print("\nmsrp: ")
		msrpR := bufio.NewReader(os.Stdin)
		msrp, _ := msrpR.ReadString('\n')
		msrp = strings.Trim(msrp, "\n")
		msrpFloat, err := strconv.ParseFloat(msrp,32)


		fmt.Print("\nprice: ")
		priceR := bufio.NewReader(os.Stdin)
		price, _ := priceR.ReadString('\n')
		price = strings.Trim(price, "\n")
		priceFloat, err := strconv.ParseFloat(price,32)


		product, err := crudClient.CreateProduct(ctx, &pb.InventoryProduct{Barcode: barcode, Store: store, Name: name, Category: category,Description: description,Image: image,
			Restricted: bool(restrictedBool) ,Inventory: int32(inventoryInt), Msrp: float32(msrpFloat) , Price : float32(priceFloat)})
		if err != nil {
			log.Fatalf("Could not create a new Product: %v", err)
		}
		fmt.Println("\nInserted product with ", product.Id)

	case "2\n":
		fmt.Print("\nbarcode : ")
		barcodeR := bufio.NewReader(os.Stdin)
		barcode, _ := barcodeR.ReadString('\n')
		barcode = strings.Trim(barcode, "\n")

		product, err := crudClient.ReadProduct(ctx, &pb.ID{Id: barcode})
		if err != nil {
			log.Fatalf("Error in getting product: %v", err)
		}
		fmt.Println("\n Product found!")
		fmt.Println("name:", product.Barcode, " store: ",product.Store, "name : ",product.Name)
	

	case "3\n":
		
		fmt.Print("\nbarcode: ")
		barcodeR := bufio.NewReader(os.Stdin)
		barcode, _ := barcodeR.ReadString('\n')
		barcode = strings.Trim(barcode, "\n")

		fmt.Print("\nstore: ")
		storeR := bufio.NewReader(os.Stdin)
		store, _ := storeR.ReadString('\n')
		store = strings.Trim(store, "\n")

		fmt.Print("\nproduct name: ")
		nameR := bufio.NewReader(os.Stdin)
		name, _ := nameR.ReadString('\n')
		name = strings.Trim(name, "\n")


		fmt.Print("\ncategory: ")
		categoryR := bufio.NewReader(os.Stdin)
		category, _ := categoryR.ReadString('\n')
		category = strings.Trim(category, "\n")


		fmt.Print("\ndescription: ")
		descriptionR := bufio.NewReader(os.Stdin)
		description, _ := descriptionR.ReadString('\n')
		description = strings.Trim(description, "\n")


		fmt.Print("\nimage: ")
		imageR := bufio.NewReader(os.Stdin)
		image, _ := imageR.ReadString('\n')
		image = strings.Trim(image, "\n")


		fmt.Print("\nis restricted: ")
		restrictedR := bufio.NewReader(os.Stdin)
		restricted, _ := restrictedR.ReadString('\n')
		restricted = strings.Trim(restricted, "\n")
		restrictedBool, err := strconv.ParseBool(restricted)


		fmt.Print("\ninventory: ")
		inventoryR := bufio.NewReader(os.Stdin)
		inventory, _ := inventoryR.ReadString('\n')
		inventory = strings.Trim(store, "\n")
		inventoryInt, err := strconv.Atoi(inventory)


		fmt.Print("\nmsrp: ")
		msrpR := bufio.NewReader(os.Stdin)
		msrp, _ := msrpR.ReadString('\n')
		msrp = strings.Trim(msrp, "\n")
		msrpFloat, err := strconv.ParseFloat(msrp,32)


		fmt.Print("\nprice: ")
		priceR := bufio.NewReader(os.Stdin)
		price, _ := priceR.ReadString('\n')
		price = strings.Trim(price, "\n")
		priceFloat, err := strconv.ParseFloat(price,32)


		product, err := crudClient.UpdateProduct(ctx, &pb.InventoryProduct{Barcode: barcode, Store: store, Name: name, Category: category,Description: description,Image: image,
			Restricted: bool(restrictedBool) ,Inventory: int32(inventoryInt), Msrp: float32(msrpFloat) , Price : float32(priceFloat)})
		if err != nil {
			log.Fatalf("Error updating product: %v", err)
		}
		fmt.Println("\nupdated product with ", product.Id)


	case "4\n":
		
		fmt.Print("\nbarcode: ")
		barcodeR := bufio.NewReader(os.Stdin)
		barcode, _ := barcodeR.ReadString('\n')
		barcode = strings.Trim(barcode, "\n")
		product, _ := crudClient.DeleteProduct(ctx, &pb.ID{Id: barcode})
		if product != nil {
			log.Printf("\nProduct %s deleted", product.Id)
		} else {
			log.Printf("Successful delete (even though ID didn't exist)")
		}

	default:
		fmt.Println("\nWrong option!")
	}
}
