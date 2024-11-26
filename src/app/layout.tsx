import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";

import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/app-sidebar";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <html>
      <body>
        <SidebarProvider>
          <AppSidebar />
          <SidebarTrigger />
          <main className={"m-8 w-full"}>{children}</main>
        </SidebarProvider>
      </body>
    </html>
  );
}

