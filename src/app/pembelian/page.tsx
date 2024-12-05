"use client";

import { Button } from "@/components/ui/button";
import React, { useState, useEffect } from "react";
import Link from "next/link";


interface Transaction {
  id: string;
  created_at: string;
  grand_total: number;
  status: string;
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
      console.log("Transactions data:", data.data.transactions);
      setTransactions(data.data.transactions);
    } catch (error) {
      console.error("Error fetching transactions:", error);
    }
  };

  // Fungsi untuk menangani pembayaran
  const handlePay = async (transaction_id: string, amount: number) => {
    try {
      const response = await fetch("http://localhost:5000/api/payment/standard/", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          transaction_id,
          amount,
        }),
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error("Payment failed");
      }

      const data = await response.json();
      console.log("Payment response:", data);

      if (data.data) {
        window.open(data.data, "_blank");
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
      <table>
        <thead>
          <tr>
            <th>Tanggal</th>
            <th>Status</th>
            <th>Total</th>
            <th>Aksi</th>
          </tr>
        </thead>
        <tbody>
          {transactions.map((transaction, index) => (
            <tr key={index}>
              <td>{new Date(transaction.created_at).toLocaleString()}</td>
              <td>{transaction.status}</td>
              <td>{transaction.grand_total}</td>
              <td>
                {/* Menampilkan tombol berdasarkan status */}
                {transaction.status === "draft" ? (
                  <Button
                    onClick={() =>
                      handlePay(transaction.id, transaction.grand_total)
                    }
                  >
                    Bayar
                  </Button>
                ) : (
                  <Button>
                    <Link href={"pembelian/invoice/[id]"} as={`/pembelian/invoice/${transaction.id}`}>
                    Lihat
                    </Link>
                    </Button>
                )}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default PembelianPage;

