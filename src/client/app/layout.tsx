import "./globals.css";
import { Frank_Ruhl_Libre } from "next/font/google";
import {} from "next/font/local";
import { cn } from "@/lib/utils";
import { Toaster } from "react-hot-toast";
import { getLocale, getMessages } from "next-intl/server";
import { NextIntlClientProvider } from "next-intl";
const inter = Frank_Ruhl_Libre({
  subsets: ["latin"],
});

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const messages = await getMessages({
    locale: "en-US",
  });
  const locale = await getLocale();

  return (
    <html lang={locale}>
      <NextIntlClientProvider messages={messages}>
        <body className={cn(inter.className, "bg-white", "h-screen")}>
          <Toaster position="top-center" />
          {children}
        </body>
      </NextIntlClientProvider>
    </html>
  );
}
