package main

import (
	"fmt"
)

// // Person stub
// type Person struct {
// 	Drugs map[string]Drug
// }

// // Drug stub
// type Drug struct {
// 	TargetReceptors map[string]Receptor
// }

// // Receptor stub
// type Receptor struct {
// 	SideEffects map[string]SideEffect
// }

// // SideEffect stub
// type SideEffect struct {
// 	Description string
// 	Count       int
// }

// var m map[string]int

// // Joe stub
// var Joe Person

// // TmpDrug stub
// var TmpDrug Drug

// // TmpReceptor stub
// var TmpReceptor Receptor

// // TmpSideEffect stub
// var TmpSideEffect SideEffect

// func LoadSideEffect() {
// 	list := arraylist.New()

// }

// func newClient() *dgo.Dgraph {
// 	// Dial a gRPC connection. The address to dial to can be configured when
// 	// setting up the dgraph cluster.
// 	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return dgo.NewDgraphClient(
// 		api.NewDgraphClient(d),
// 	)
// }

// func setup(c *dgo.Dgraph) {
// 	// Install a schema into dgraph. Accounts have a `name` and a `balance`.
// 	err := c.Alter(context.Background(), &api.Operation{
// 		Schema: `
// 			name: string @index(term) .
// 			balance: int .
// 		`,
// 	})
// }

// func nuke_from_orbit(c *dgo.Dgraph) {
// 	err := c.Alter(context.Background(), &api.Operation{DropAll: true})
// }

// func runTxn(c *dgo.Dgraph) {
// 	txn := c.NewTxn()
// 	defer txn.Discard()
// }

func init() {
	// Joe = Person{
	// 	Drugs: make(map[string]Drug),
	// }

	// TmpDrug = Drug{
	// 	TargetReceptors: make(map[string]Receptor),
	// }

	// TmpReceptor = Receptor{
	// 	SideEffects: make(map[string]SideEffect),
	// }

	// TmpSideEffect.Description = "death"
	// TmpSideEffect.Count = 1

	// TmpReceptor["sideEffect1"] := TmpSideEffect

	// Joe["drug0"] = Drug0

	// Joe.Drugs = []Drug{
	// 	Drug{
	// 		TargetReceptors: []Receptor{
	// 			Receptor{
	// 				SideEffects: []SideEffect{
	// 					SideEffect{
	// 						Description: "boom",
	// 						Count:       1,
	// 					},
	// 					SideEffect{
	// 						Description: "boom",
	// 						Count:       1,
	// 					},
	// 				},
	// 			},
	// 			Receptor{
	// 				SideEffects: []SideEffect{
	// 					SideEffect{
	// 						Description: "boom",
	// 						Count:       1,
	// 					},
	// 					SideEffect{
	// 						Description: "boom",
	// 						Count:       1,
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}

func main() {
	fmt.Println("hello world")
}
