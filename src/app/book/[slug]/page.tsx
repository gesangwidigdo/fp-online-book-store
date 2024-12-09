"use client";

import { useEffect, useState } from "react";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

interface BookDetail {
  id: string;
  isbn: string;
  title: string;
  slug: string;
  author: string;
  summary: string;
  book_image: string;
  publication_year: number;
  price: number;
}

function formattedPrice(x: number) {
  return "Rp" + x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
}

export default function BookDetail({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  const router = useRouter();
  const [book, setBook] = useState<BookDetail | null>(null);
  const [slug, setSlug] = useState<string | null>(null);
  const [quantity, setQuantity] = useState<number>(1);
  const [subtotal, setSubtotal] = useState<number>(0);

  useEffect(() => {
    const unwrapParams = async () => {
      const { slug } = await params;
      setSlug(slug);
    };
    unwrapParams();
  }, [params]);

  useEffect(() => {
    const fetchBook = async () => {
      if (slug) {
        const fetchedBook = await fetchBookBySlug(slug);
        setBook(fetchedBook);
      }
    };
    fetchBook();
  }, [slug]);

  useEffect(() => {
    setSubtotal(book ? book.price * quantity : 0);
  }, [quantity, book]);

  const handleQuantityChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = Number(e.target.value);
    setQuantity(Number(value));
  };

  if (!book) {
    return <div>Loading...</div>;
  }

  const AddToCart = async (id: string, quantity: number) => {
    try {
      const response = await fetch(
        `http://localhost:5000/api/book_transaction/`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
          body: JSON.stringify({
            book_id: id,
            quantity: quantity,
          }),
        }
      );

      if (!response.ok) {
        if (response.status === 401) {
          router.push("/login");
          return;
        }
        return {
          notFound: true,
        };
      }

      const cartData = await response.json();
      console.log(cartData.data);
      router.push("/keranjang");
      return cartData.data;
    } catch (error) {
      console.error("Error adding to cart:", error);
      router.push("/keranjang");
    }
  };

  return (
    <div className="flex">
      {/* image and details */}
      <div className="mx-5">
        <img
          className="w-64 max-w-lg mb-8 border-4 border-black p-1"
          src={book.book_image}
          alt={book.title}
        />
        <div className="details">
          <p className="text-md mb-3 font-light">ISBN: {book.isbn}</p>
          <p className="text-md mb-3 font-light">Author: {book.author}</p>
          <p className="text-md mb-3 font-light">
            Publication Year: {book.publication_year}
          </p>
          <p className="text-md mb-3 font-light">
            Price: {formattedPrice(book.price)}
          </p>
        </div>
      </div>

      {/* summary */}
      <div className="mx-5 w-2/5">
        <div className="header">
          <h2 className="text-3xl font-semibold">{book.title}</h2>
          <p className="text-md font-light text-gray-500">{book.author}</p>
        </div>
        <div className="description my-10 w-auto">
          <p className="text-md font-semibold">Summary</p>
          <p className="text-md font-light text-justify">{book.summary}</p>
        </div>
      </div>

      {/* add to cart */}
      <div className="w-auto">
        <Card className="w-[300px]">
          <CardHeader>
            <CardTitle>Add to Cart</CardTitle>
          </CardHeader>
          <CardContent>
            <form>
              <div className="grid w-full items-center gap-4">
                <div className="flex flex-col space-y-1.5">
                  <Label htmlFor="quantity">Quantity</Label>
                  <Input
                    type="number"
                    id="quantity"
                    name="quantity"
                    placeholder="0"
                    min="1"
                    max="10"
                    value={quantity}
                    onChange={handleQuantityChange}
                  />
                </div>
                <div className="flex flex-col space-y-1.5">
                  <Label htmlFor="subtotal">Subtotal</Label>
                  <Input
                    type="text"
                    id="subtotal"
                    name="subtotal"
                    placeholder={formattedPrice(0)}
                    value={formattedPrice(subtotal)}
                    readOnly
                  />
                </div>
              </div>
            </form>
          </CardContent>
          <CardFooter className="flex justify-between">
            <Button onClick={() => book && AddToCart(book.id, quantity)}>
              Add to Cart
            </Button>
          </CardFooter>
        </Card>
      </div>
    </div>
  );
}

export const fetchBooks = async () => {
  const response = await fetch("http://localhost:5000/api/book/", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });

  if (!response.ok) {
    return {
      notFound: true,
    };
  }

  const booksData = await response.json();
  return booksData.data;
};

export const fetchBookBySlug = async (slug: string) => {
  const response = await fetch(`http://localhost:5000/api/book/${slug}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });

  if (!response.ok) {
    return {
      notFound: true,
    };
  }

  const bookData = await response.json();
  return bookData.data;
};

export const getStaticParams = async () => {
  const data = await fetchBooks();
  return data.map((book: BookDetail) => ({
    slug: book.slug,
  }));
};

