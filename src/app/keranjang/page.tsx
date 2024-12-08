"use client";

import { useEffect, useState } from "react";
import { TransactionBookList } from "@/components/transactionBookList";
import { Button } from "@/components/ui/button";

interface Books {
  id: string,
  title: string,
  book_image: string,
  price: number,
  quantity: number,
  total: number,
}

interface TransactionBooks {
  id: string;
  grand_total: number;
  books: Books[];
}

function formattedPrice(x: number) {
  return "Rp" + x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
}

export default function KeranjangPage() {
  const [transactions, setTransactions] = useState<TransactionBooks | null>(null);
  const loadTransactions = async () => {
    try {
      const response = await fetch("http://localhost:5000/api/transaction/", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error("Failed to get transactions");
      }

      const transactionsData = await response.json();
      console.log(transactionsData.data)
      setTransactions(transactionsData.data);
    } catch (error) {
      console.error("Error fetching transactions:", error);
    }
  };

  useEffect(() => {
    loadTransactions();
  }, []);

  if (transactions === null) {
    return <div>Something wrong...</div>;
  }

  const grandTotal = transactions.books.reduce((acc, book) => acc + book.total, 0);

  return (
    <div className="">
      <h1>Keranjang Page</h1>
      <div className="border-b-2 pb-4 mt-5">
        {transactions.books.map((book) => (
          <TransactionBookList 
            key={book.id}
            transaction_id={transactions.id}
            book_id={book.id}
            title={book.title}
            book_image={book.book_image}
            price={book.price}
            quantity={book.quantity}
            total={book.total}
          />
        ))}

      </div>
      <div className="mt-5 flex justify-between font-bold">
        <p>Total: {formattedPrice(grandTotal)}</p>
        <Button>Buat Pesanan</Button>
      </div>
    </div>
  );
}

