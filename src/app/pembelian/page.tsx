"use client";

import { Button } from "@/components/ui/button";
import React, { useState, useEffect } from "react";
import Link from "next/link";
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
import { Badge } from "@/components/ui/badge";

interface Transaction {
  id: string;
  created_at: string;
  grand_total: number;
  status: string;
}

function formattedPrice(x: number) {
  return "Rp" + x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
}

const PembelianPage = () => {
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  useEffect(() => {
    fetchTransactions();
  }, []);

  const fetchTransactions = async () => {
    try {
      const response = await fetch(
        "http://localhost:5000/api/transaction/list",
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
        }
      );

      if (!response.ok) {
        throw new Error("Failed to get transactions");
      }

      const data = await response.json();
      console.log("Transactions data:", data);
      setTransactions(data.data.transactions);
    } catch (error) {
      console.error("Error fetching transactions:", error);
    }
  };

  // Fungsi untuk menangani pembayaran
  const handlePay = async (transaction_id: string, amount: number) => {
    try {
      const response = await fetch(
        "http://localhost:5000/api/payment/standard",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            transaction_id,
            amount,
          }),
          credentials: "include",
        }
      );

      if (!response.ok) {
        throw new Error("Payment failed");
      }

      const data = await response.json();
      console.log("Payment response:", data);

      if (data.data) {
        window.open(`/pembelian/invoice/${data.data.transaction_id}`, "_blank");
      } else {
        alert("Pembayaran berhasil, tetapi tidak ada link.");
      }
    } catch (error) {
      console.error("Error during payment:", error);
      alert("Pembayaran gagal!");
    }
  };

  return (
    <div>
      <h1>Pembelian Saya</h1>
      <Table>
        <TableCaption>A list of your transactions.</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="font-bold text-center">Date</TableHead>
            <TableHead className="font-bold text-center">
              Total Amount
            </TableHead>
            <TableHead className="font-bold text-center">Status</TableHead>
            <TableHead className="text-right font-bold"></TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {transactions.map((transaction) => (
            <TableRow key={transaction.id}>
              <TableCell className="text-center">
                {new Date(transaction.created_at).toLocaleString()}
              </TableCell>
              <TableCell className="text-center">
                {formattedPrice(transaction.grand_total)}
              </TableCell>

              <TableCell className={`text-center font-bold `}>
                <Badge
                  variant="outline"
                  className={`text-center font-bold ${
                    transaction.status === "draft"
                      ? "bg-yellow-600 text-white"
                      : "bg-green-600 text-white"
                  }`}
                >
                  {transaction.status}
                </Badge>
              </TableCell>

              <TableCell className="text-right">
                {transaction.status === "draft" &&
                transaction.grand_total != 0 ? (
                  <Button
                    onClick={() =>
                      handlePay(transaction.id, transaction.grand_total)
                    }
                  >
                    Bayar
                  </Button>
                ) : transaction.grand_total === 0 ? (
                  <Button>
                    <Link href={"/"}>Order</Link>
                  </Button>
                ) : (
                  <Button>
                    <Link href={`/pembelian/invoice/${transaction.id}`}>
                      Lihat
                    </Link>
                  </Button>
                )}
                {/* <Button>
                  <Link href={`/pembelian/invoice/${transaction.id}`}>
                    Lihat
                  </Link>
                </Button> */}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
};

export default PembelianPage;

