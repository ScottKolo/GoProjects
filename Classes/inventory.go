// Create an application which manages an inventory of products. Create a
// product class which has a price, id, and quantity on hand. Then create an
// *inventory* class which keeps track of various products and can sum up
// the inventory value.

package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Product "class" that contains a price, id, and quantity
type Product struct {
	price    float64
	id       int
	quantity int
}

// Inventory "class" that contains a list of Products
type Inventory struct {
	products []*Product
}

// Inventory "class" method that prints a formatted inventory list
func (inv *Inventory) printInventory() {
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with right alignment
	w.Init(os.Stdout, 6, 0, 4, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "ID\tPRICE\tQUANTITY\t")
	for _, p := range inv.products {
		fmt.Fprintf(w, "%d\t$%.2f\t%d\t\n", p.id, p.price, p.quantity)
	}
	fmt.Fprintln(w)
	w.Flush()
}

// Inventory "class" method that adds a product to the inventory
func (inv *Inventory) add(p *Product) {
	inv.products = append(inv.products, p)
}

func main() {
	// Create some Products
	p1 := Product{0.99, 1, 16}
	p2 := Product{19.99, 2, 5}
	p3 := Product{4.99, 3, 8}

	// Create an Inventory and add the products to it
	var inv Inventory
	inv.add(&p1)
	inv.add(&p2)
	inv.add(&p3)

	fmt.Println("Inventory List Before Changes")
	inv.printInventory()

	// Make a few changes
	p2.quantity += 4
	p1.price += 2
	p3.quantity -= 5

	fmt.Println("Inventory List After Changes")
	inv.printInventory()
}
