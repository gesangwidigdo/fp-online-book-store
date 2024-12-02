"use client";
import { BookCard } from '@/components/book-card';
import { useEffect, useState } from 'react';

interface Book {
  id: string;
  book_image: string;
  title: string;
  slug: string;
  author: string;
  summary: string;
  price: number;
}

export default function HomePage() {
  const [books, setBooks] = useState<Book[]>([]);
  const loadBooks = async () => {
    try {
      const response = await fetch("http://localhost:5000/api/book/", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error("Failed to get books");
      }

      const booksData = await response.json();
      console.log(booksData.data);
      setBooks(booksData.data);
    } catch (error) {
      console.error("Error fetching books:", error);
    }
  };

  useEffect(() => {
    loadBooks();
  }, []); 

  if (books === null) {
    return <div>Something wrong...</div>;
  }

  return (
    <div>
      <h1>Home Page</h1>
        <div className="flex flex-wrap justify-center">
          {books.map((book) => (
            <BookCard 
              key={book.id}
              slug={book.slug}
              imgUrl={book.book_image}
              title={book.title}
              author={book.author}
              price={book.price}
            />
          ))}
        </div>
    </div>
  );
}
