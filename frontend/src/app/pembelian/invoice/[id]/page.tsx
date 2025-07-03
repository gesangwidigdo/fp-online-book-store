"use client";

import { useEffect, useState } from "react";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

interface Book {
  book_image: string;
  title: string;
  price: number;
  quantity: number;
  total: number;
}

function formattedPrice(x: number) {
  return "Rp" + x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
}

const InvoicePage = ({ params }: { params: Promise<{ id: string }> }) => {
  const [id, setId] = useState<string | null>(null);
  const [books, setBooks] = useState<Book[]>([]);
  const [totalAmount, setTotalAmount] = useState<number>(0);

  useEffect(() => {
    const unwrapParams = async () => {
      const { id } = await params;
      setId(id);
      console.log("ID:", id);
    };

    unwrapParams();
  }, [params]);

  useEffect(() => {
    if (id) {
      const loadCurTransaction = async () => {
        try {
          const response = await fetch(
            `http://localhost:5000/api/transaction/${id}`,
            {
              method: "GET",
              headers: {
                "Content-Type": "application/json",
              },
              credentials: "include",
            }
          );

          if (!response.ok) {
            throw new Error("Failed to get books");
          }

          const booksData = await response.json();
          console.log(booksData.data);
          setBooks(booksData.data.books);
          setTotalAmount(booksData.data.grand_total);
        } catch (error) {
          console.error("Error fetching books:", error);
        }
      };

      loadCurTransaction();
    }
  }, [id]);
  if (!id) {
    return <div>Loading...</div>;
  }

  return (
    <Table>
      <TableCaption>A list of your recent invoices.</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="font-bold text-center">Invoice</TableHead>
          <TableHead className="font-bold text-center">Title</TableHead>
          <TableHead className="font-bold text-center">Price</TableHead>
          <TableHead className="font-bold text-center">Quantity</TableHead>
          <TableHead className="text-right font-bold">Total Amount</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {books.map((book) => (
          <TableRow key={book.title}>
            <TableCell className="font-medium flex justify-center">
              <img
                src={book.book_image}
                alt={book.title}
                className="object-cover flex"
              />
            </TableCell>
            <TableCell className="text-center">{book.title}</TableCell>
            <TableCell className="text-center">
              {formattedPrice(book.price)}
            </TableCell>
            <TableCell className="text-center">{book.quantity}</TableCell>
            <TableCell className="text-right">
              {formattedPrice(book.total)}
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
      <TableFooter>
        <TableRow>
          <TableCell colSpan={4} className="text-center font-extrabold">
            Total
          </TableCell>
          <TableCell className="text-right font-extrabold">
            {" "}
            {formattedPrice(totalAmount)}{" "}
          </TableCell>
        </TableRow>
      </TableFooter>
    </Table>
  );
};

export default InvoicePage;

