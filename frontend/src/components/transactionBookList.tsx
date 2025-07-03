import { Pencil } from "lucide-react";
import { Trash2 } from "lucide-react";
import { Button } from "./ui/button";

import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useEffect, useState } from "react";

interface TransactionBookProps {
  transaction_id: string;
  book_id: string;
  book_image: string;
  title: string;
  price: number;
  quantity: number;
  total: number;
}

function formattedPrice(x: number) {
  return "Rp" + x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
}

export function TransactionBookList({ transaction_id, book_id, book_image, title, price, quantity, total }: TransactionBookProps) {
  const [newQuantity, setQuantity] = useState<number>(quantity);
  const [subtotal, setSubtotal] = useState<number>(total);

  useEffect(() => {
    setSubtotal(price * (newQuantity));
  }, [newQuantity]);

  const handleQuantityChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = Number(e.target.value);
    setQuantity(Number(value));
  }

  const editQuantity = async (transaction_id: string, book_id: string, quantity: number) => {
    try {
      const response = await fetch(`http://localhost:5000/api/book_transaction/${transaction_id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify({ 
          book_id: book_id,
          quantity: quantity,
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to edit quantity");
      }

      window.location.reload();

      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error("Error editing quantity:", error);
    }
  }

  const deleteBook = async (transaction_id: string, book_id: string) => {
    try {
      const response = await fetch(`http://localhost:5000/api/book_transaction/${transaction_id}`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify({ 
          book_id: book_id,
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to delete book");
      }

      window.location.reload();

      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error("Error deleting book:", error);
    }
  }
  
  return (
    <div key={book_id} className="flex justify-between items-center p-2 border-b border-gray-200">
      <div className="flex items-center">
        <img src={book_image} alt={title} className="w-16 h-16 object-cover" />
        <div className="ml-4">
          <h3 className="font-bold">{title}</h3>
          <p className="text-sm text-gray-500">{formattedPrice(price)}</p>
        </div>
      </div>
      <div className="flex items-center">
        <p className="mr-4">Qty: {quantity}</p>
        <p className="font-bold mr-4">{formattedPrice(total)}</p>
        <Dialog>
          <DialogTrigger asChild>
            <Button size="icon">
                <Pencil />
            </Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[425px]">
            <DialogHeader>
              <DialogTitle>Edit Order</DialogTitle>
            </DialogHeader>
            <div className="grid gap-4 py-4">
              <div className="grid grid-cols-4 items-center gap-4">
                <Label htmlFor="quantity" className="text-right">
                  Quantity
                </Label>
                <Input
                  id="quantity"
                  name="quantity"
                  defaultValue={quantity}
                  className="col-span-3"
                  type="number"
                  min="1"
                  max="10"
                  onChange={handleQuantityChange}
                />
              </div>
              <div className="grid grid-cols-4 items-center gap-4">
                <Label htmlFor="subtotal" className="text-right">
                  Subtotal
                </Label>
                <Input
                  id="subtotal"
                  name="subtotal"
                  value={formattedPrice(subtotal)}
                  className="col-span-3"
                  type="text"
                  readOnly
                />
              </div>
            </div>
            <DialogFooter>
              <Button onClick={() => editQuantity(transaction_id, book_id, newQuantity)}>Save changes</Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
        <Button variant="destructive" className="mx-3" onClick={() => deleteBook(transaction_id, book_id)}>
          <Trash2 />
        </Button>
      </div>
    </div>
  );
}